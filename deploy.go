package ecspresso

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kayac/ecspresso/v2/appspec"
	"github.com/samber/lo"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	cdTypes "github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	isatty "github.com/mattn/go-isatty"
)

const (
	CodeDeployConsoleURLFmt = "https://%s.console.aws.amazon.com/codesuite/codedeploy/deployments/%s?region=%s"
)

type DeployOption struct {
	DryRun               *bool   `help:"dry run" default:"false"`
	DesiredCount         *int32  `name:"tasks" help:"desired count of tasks" default:"-1"`
	SkipTaskDefinition   *bool   `help:"skip register a new task definition" default:"false"`
	ForceNewDeployment   *bool   `help:"force a new deployment of the service" default:"false"`
	NoWait               *bool   `help:"exit ecspresso immediately after just deployed without waiting for service stable" default:"false"`
	SuspendAutoScaling   *bool   `help:"suspend application auto-scaling attached with the ECS service"`
	ResumeAutoScaling    *bool   `help:"resume application auto-scaling attached with the ECS service"`
	RollbackEvents       *string `help:"roll back when specified events happened (DEPLOYMENT_FAILURE,DEPLOYMENT_STOP_ON_ALARM,DEPLOYMENT_STOP_ON_REQUEST,...) CodeDeploy only." default:""`
	UpdateService        *bool   `help:"update service attributes by service definition" default:"true" negatable:""`
	LatestTaskDefinition *bool   `help:"deploy with the latest task definition without registering a new task definition" default:"false"`
}

func (opt DeployOption) DryRunString() string {
	if *opt.DryRun {
		return dryRunStr
	}
	return ""
}

func calcDesiredCount(sv *Service, opt DeployOption) *int32 {
	if sv.SchedulingStrategy == types.SchedulingStrategyDaemon {
		return nil
	}
	if oc := opt.DesiredCount; oc != nil {
		if *oc == DefaultDesiredCount {
			return sv.DesiredCount
		}
		return oc // --tasks
	}
	return nil
}

func (d *App) Deploy(ctx context.Context, opt DeployOption) error {
	d.Log("[DEBUG] deploy")
	d.LogJSON(opt)
	ctx, cancel := d.Start(ctx)
	defer cancel()

	var sv *Service
	d.Log("Starting deploy %s", opt.DryRunString())
	sv, err := d.DescribeServiceStatus(ctx, 0)
	if err != nil {
		if errors.As(err, &errNotFound) {
			d.Log("Service %s not found. Creating a new service %s", d.Service, opt.DryRunString())
			return d.createService(ctx, opt)
		}
		return err
	}

	deployFunc := d.UpdateServiceTasks // default
	waitFunc := d.WaitServiceStable    // default
	// detect controller
	if dc := sv.DeploymentController; dc != nil {
		switch dc.Type {
		case types.DeploymentControllerTypeCodeDeploy:
			deployFunc = d.DeployByCodeDeploy
			waitFunc = d.WaitForCodeDeploy
		case types.DeploymentControllerTypeEcs:
			deployFunc = d.UpdateServiceTasks
			waitFunc = d.WaitServiceStable
		default:
			return fmt.Errorf("unsupported deployment controller type: %s", dc.Type)
		}
	}

	var tdArn string
	if *opt.LatestTaskDefinition {
		family := strings.Split(arnToName(*sv.TaskDefinition), ":")[0]
		var err error
		tdArn, err = d.findLatestTaskDefinitionArn(ctx, family)
		if err != nil {
			return err
		}
	} else if *opt.SkipTaskDefinition {
		tdArn = *sv.TaskDefinition
	} else {
		td, err := d.LoadTaskDefinition(d.config.TaskDefinitionPath)
		if err != nil {
			return err
		}
		if *opt.DryRun {
			d.Log("[INFO] task definition: %s", MustMarshalJSONStringForAPI(td))
		} else {
			newTd, err := d.RegisterTaskDefinition(ctx, td)
			if err != nil {
				return err
			}
			tdArn = *newTd.TaskDefinitionArn
		}
	}

	var count *int32
	if d.config.ServiceDefinitionPath != "" && aws.ToBool(opt.UpdateService) {
		newSv, err := d.LoadServiceDefinition(d.config.ServiceDefinitionPath)
		if err != nil {
			return err
		}
		ds, err := diffServices(newSv, sv, "", d.config.ServiceDefinitionPath, true)
		if err != nil {
			return fmt.Errorf("failed to diff of service definitions: %w", err)
		}
		if ds != "" {
			if err = d.UpdateServiceAttributes(ctx, newSv, tdArn, opt); err != nil {
				return err
			}
			sv = newSv // updated
		} else {
			d.Log("service attributes will not change")
		}
		count = calcDesiredCount(newSv, opt)
	} else {
		count = calcDesiredCount(sv, opt)
	}
	if count != nil {
		d.Log("desired count: %d", *count)
	} else {
		d.Log("desired count: unchanged")
	}

	if *opt.DryRun {
		d.Log("DRY RUN OK")
		return nil
	}

	// manage auto scaling only when set option --suspend-auto-scaling or --no-suspend-auto-scaling explicitly
	if suspendState := opt.SuspendAutoScaling; suspendState != nil {
		if err := d.suspendAutoScaling(ctx, *suspendState); err != nil {
			return err
		}
	}

	if err := deployFunc(ctx, tdArn, count, sv, opt); err != nil {
		return err
	}

	if *opt.NoWait {
		d.Log("Service is deployed.")
		return nil
	}

	if err := waitFunc(ctx, sv); err != nil {
		return err
	}

	d.Log("Service is stable now. Completed!")
	return nil
}

func (d *App) UpdateServiceTasks(ctx context.Context, taskDefinitionArn string, count *int32, sv *Service, opt DeployOption) error {
	in := &ecs.UpdateServiceInput{
		Service:            sv.ServiceName,
		Cluster:            aws.String(d.Cluster),
		TaskDefinition:     aws.String(taskDefinitionArn),
		DesiredCount:       count,
		ForceNewDeployment: *opt.ForceNewDeployment,
	}
	msg := "Updating service tasks"
	if *opt.ForceNewDeployment {
		msg = msg + " with force new deployment"
	}
	msg = msg + "..."
	d.Log(msg)
	d.LogJSON(in)

	_, err := d.ecs.UpdateService(ctx, in)
	if err != nil {
		return fmt.Errorf("failed to update service tasks: %w", err)
	}
	time.Sleep(delayForServiceChanged) // wait for service updated
	return nil
}

func svToUpdateServiceInput(sv *Service) *ecs.UpdateServiceInput {
	in := &ecs.UpdateServiceInput{
		CapacityProviderStrategy:      sv.CapacityProviderStrategy,
		DeploymentConfiguration:       sv.DeploymentConfiguration,
		DesiredCount:                  sv.DesiredCount,
		EnableECSManagedTags:          &sv.EnableECSManagedTags,
		EnableExecuteCommand:          &sv.EnableExecuteCommand,
		HealthCheckGracePeriodSeconds: sv.HealthCheckGracePeriodSeconds,
		LoadBalancers:                 sv.LoadBalancers,
		NetworkConfiguration:          sv.NetworkConfiguration,
		PlacementConstraints:          sv.PlacementConstraints,
		PlacementStrategy:             sv.PlacementStrategy,
		PlatformVersion:               sv.PlatformVersion,
		PropagateTags:                 sv.PropagateTags,
		ServiceRegistries:             sv.ServiceRegistries,
	}
	if sv.SchedulingStrategy == types.SchedulingStrategyDaemon {
		in.PlacementStrategy = nil
	}
	return in
}

func (d *App) UpdateServiceAttributes(ctx context.Context, sv *Service, taskDefinitionArn string, opt DeployOption) error {
	in := svToUpdateServiceInput(sv)
	if sv.isCodeDeploy() {
		d.Log("[INFO] deployment by CodeDeploy")
		// unable to update attributes below with a CODE_DEPLOY deployment controller.
		in.NetworkConfiguration = nil
		in.PlatformVersion = nil
		in.ForceNewDeployment = false
		in.LoadBalancers = nil
		in.ServiceRegistries = nil
		in.TaskDefinition = nil
	} else {
		d.Log("[INFO] deployment by ECS rolling update")
		in.ForceNewDeployment = aws.ToBool(opt.ForceNewDeployment)
		in.TaskDefinition = aws.String(taskDefinitionArn)
	}
	in.Service = aws.String(d.Service)
	in.Cluster = aws.String(d.Cluster)

	if *opt.DryRun {
		d.Log("[INFO] update service input: %s", MustMarshalJSONStringForAPI(in))
		return nil
	}
	d.Log("Updating service attributes...")

	if _, err := d.ecs.UpdateService(ctx, in); err != nil {
		return fmt.Errorf("failed to update service attributes: %w", err)
	}
	time.Sleep(delayForServiceChanged) // wait for service updated
	return nil
}

func (d *App) DeployByCodeDeploy(ctx context.Context, taskDefinitionArn string, count *int32, sv *Service, opt DeployOption) error {
	if count != nil {
		d.Log("updating desired count to %d", *count)
	}
	_, err := d.ecs.UpdateService(
		ctx,
		&ecs.UpdateServiceInput{
			Service:      aws.String(d.Service),
			Cluster:      aws.String(d.Cluster),
			DesiredCount: count,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}
	if aws.ToBool(opt.SkipTaskDefinition) && !aws.ToBool(opt.UpdateService) && !aws.ToBool(opt.ForceNewDeployment) {
		// no need to create new deployment.
		return nil
	}

	return d.createDeployment(ctx, sv, taskDefinitionArn, opt.RollbackEvents)
}

func (d *App) findDeploymentInfo(ctx context.Context) (*cdTypes.DeploymentInfo, error) {
	// search deploymentGroup in CodeDeploy
	d.Log("[DEBUG] find applications in CodeDeploy")
	apps, err := d.findCodeDeployApplications(ctx)
	if err != nil {
		return nil, err
	}

	for _, app := range apps {
		groups, err := d.findCodeDeployDeploymentGroups(ctx, *app.ApplicationName)
		if err != nil {
			return nil, err
		}
		for _, dg := range groups {
			d.Log("[DEBUG] deploymentGroup %v", dg)
			for _, ecsService := range dg.EcsServices {
				if *ecsService.ClusterName == d.config.Cluster && *ecsService.ServiceName == d.config.Service {
					return &cdTypes.DeploymentInfo{
						ApplicationName:      app.ApplicationName,
						DeploymentGroupName:  dg.DeploymentGroupName,
						DeploymentConfigName: dg.DeploymentConfigName,
					}, nil
				}
			}
		}
	}
	return nil, fmt.Errorf(
		"failed to find CodeDeploy Application/DeploymentGroup for ECS service %s on cluster %s",
		d.config.Service,
		d.config.Cluster,
	)
}

func (d *App) findCodeDeployApplications(ctx context.Context) ([]cdTypes.ApplicationInfo, error) {
	var appNames []string
	if d.config.CodeDeploy.ApplicationName != "" {
		appNames = []string{d.config.CodeDeploy.ApplicationName}
	} else {
		pager := codedeploy.NewListApplicationsPaginator(d.codedeploy, &codedeploy.ListApplicationsInput{})
		for pager.HasMorePages() {
			p, err := pager.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to list applications: %w", err)
			}
			appNames = append(appNames, p.Applications...)
		}
	}
	d.Log("[DEBUG] found CodeDeploy applications: %v", appNames)

	var apps []cdTypes.ApplicationInfo
	// BatchGetApplications accepts applications less than 100
	for _, names := range lo.Chunk(appNames, 100) {
		res, err := d.codedeploy.BatchGetApplications(ctx, &codedeploy.BatchGetApplicationsInput{
			ApplicationNames: names,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to batch get applications in CodeDeploy: %w", err)
		}
		for _, info := range res.ApplicationsInfo {
			d.Log("[DEBUG] application %s compute platform %s", *info.ApplicationName, info.ComputePlatform)
			if info.ComputePlatform != cdTypes.ComputePlatformEcs {
				continue
			}
			apps = append(apps, info)
		}
	}
	return apps, nil
}

func (d *App) findCodeDeployDeploymentGroups(ctx context.Context, appName string) ([]cdTypes.DeploymentGroupInfo, error) {
	var groupNames []string
	if d.config.CodeDeploy.DeploymentGroupName != "" {
		groupNames = []string{d.config.CodeDeploy.DeploymentGroupName}
	} else {
		pager := codedeploy.NewListDeploymentGroupsPaginator(d.codedeploy, &codedeploy.ListDeploymentGroupsInput{
			ApplicationName: &appName,
		})
		for pager.HasMorePages() {
			p, err := pager.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to list deployment groups in CodeDeploy: %w", err)
			}
			groupNames = append(groupNames, p.DeploymentGroups...)
		}
	}
	d.Log("[DEBUG] CodeDeploy found deploymentGroups: %v", groupNames)

	var groups []cdTypes.DeploymentGroupInfo
	// BatchGetDeploymentGroups accepts applications less than 100
	for _, names := range lo.Chunk(groupNames, 100) {
		gs, err := d.codedeploy.BatchGetDeploymentGroups(ctx, &codedeploy.BatchGetDeploymentGroupsInput{
			ApplicationName:      &appName,
			DeploymentGroupNames: names,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to batch get deployment groups in CodeDeploy: %w", err)
		}
		groups = append(groups, gs.DeploymentGroupsInfo...)
	}
	return groups, nil
}

func (d *App) createDeployment(ctx context.Context, sv *Service, taskDefinitionArn string, rollbackEvents *string) error {
	spec, err := appspec.NewWithService(&sv.Service, taskDefinitionArn)
	if err != nil {
		return fmt.Errorf("failed to create appspec: %w", err)
	}
	if d.config.AppSpec != nil {
		spec.Hooks = d.config.AppSpec.Hooks
	}
	d.Log("[DEBUG] appSpecContent: %s", spec.String())

	// deployment
	dp, err := d.findDeploymentInfo(ctx)
	if err != nil {
		return err
	}
	dd := &codedeploy.CreateDeploymentInput{
		ApplicationName:      dp.ApplicationName,
		DeploymentGroupName:  dp.DeploymentGroupName,
		DeploymentConfigName: dp.DeploymentConfigName,
		Revision: &cdTypes.RevisionLocation{
			RevisionType: cdTypes.RevisionLocationTypeAppSpecContent,
			AppSpecContent: &cdTypes.AppSpecContent{
				Content: aws.String(spec.String()),
			},
		},
	}
	if ev := aws.ToString(rollbackEvents); ev != "" {
		var events []cdTypes.AutoRollbackEvent
		for _, ev := range strings.Split(ev, ",") {
			switch ev {
			case "DEPLOYMENT_FAILURE":
				events = append(events, cdTypes.AutoRollbackEventDeploymentFailure)
			case "DEPLOYMENT_STOP_ON_ALARM":
				events = append(events, cdTypes.AutoRollbackEventDeploymentStopOnAlarm)
			case "DEPLOYMENT_STOP_ON_REQUEST":
				events = append(events, cdTypes.AutoRollbackEventDeploymentStopOnRequest)
			default:
				return fmt.Errorf("invalid rollback event: %s", ev)
			}
		}
		dd.AutoRollbackConfiguration = &cdTypes.AutoRollbackConfiguration{
			Enabled: true,
			Events:  events,
		}
	}

	d.Log("[DEBUG] creating a deployment to CodeDeploy %v", dd)

	res, err := d.codedeploy.CreateDeployment(ctx, dd)
	if err != nil {
		return fmt.Errorf("failed to create deployment: %w", err)
	}
	id := *res.DeploymentId
	u := fmt.Sprintf(
		CodeDeployConsoleURLFmt,
		d.config.Region,
		id,
		d.config.Region,
	)
	d.Log("Deployment %s is created on CodeDeploy:", id)
	d.Log(u)

	if isatty.IsTerminal(os.Stdout.Fd()) {
		if err := exec.Command("open", u).Start(); err != nil {
			d.Log("Couldn't open URL %s", u)
		}
	}
	return nil
}
