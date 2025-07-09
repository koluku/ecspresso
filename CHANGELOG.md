# Changelog

## [v0.0.1](https://github.com/koluku/ecspresso/commits/v0.0.1) - 2025-07-09

## [v2.5.0](https://github.com/kayac/ecspresso/compare/v2.4.6...v2.5.0) - 2025-05-10
- Pin external actions to mitigate supply chain attacks by @mashiike in https://github.com/kayac/ecspresso/pull/816
- Bump github.com/golang-jwt/jwt/v4 from 4.5.1 to 4.5.2 by @dependabot in https://github.com/kayac/ecspresso/pull/819
- Bump github.com/golang-jwt/jwt/v5 from 5.2.1 to 5.2.2 by @dependabot in https://github.com/kayac/ecspresso/pull/818
- Bump golang.org/x/crypto from 0.32.0 to 0.35.0 by @dependabot in https://github.com/kayac/ecspresso/pull/826
- Feature Request: new option: wait for Service Deployment, not for Service Stable by @sugitak in https://github.com/kayac/ecspresso/pull/821
- Add deploy --wait-until option. by @fujiwara in https://github.com/kayac/ecspresso/pull/828
- set WaitUntil for internal use. by @fujiwara in https://github.com/kayac/ecspresso/pull/830
- Bump the aws-sdk-go-v2 group across 1 directory with 16 updates by @dependabot in https://github.com/kayac/ecspresso/pull/812
- Add external plugin by @fujiwara in https://github.com/kayac/ecspresso/pull/760
- Add deployment_config_name to config of CodeDeploy. by @fujiwara in https://github.com/kayac/ecspresso/pull/829
- switch a logger to log/slog. by @fujiwara in https://github.com/kayac/ecspresso/pull/834
- Add JSON format support for logs and outputs by @fujiwara in https://github.com/kayac/ecspresso/pull/835
- replace time.Sleep with sleepContext for handling context cancellation by @fujiwara in https://github.com/kayac/ecspresso/pull/836
- refactor: improve verify module to use contextualized state by @fujiwara in https://github.com/kayac/ecspresso/pull/837
- feat: add JSON support for AppSpec by @fujiwara in https://github.com/kayac/ecspresso/pull/838

## [v2.4.6](https://github.com/kayac/ecspresso/compare/v2.4.5...v2.4.6) - 2025-03-01
- Bump golang.org/x/crypto from 0.24.0 to 0.31.0 by @dependabot in https://github.com/kayac/ecspresso/pull/779
- update golang.org/x/net v0.34.0 by @fujiwara in https://github.com/kayac/ecspresso/pull/782
- Bump github.com/fujiwara/tfstate-lookup from 1.3.2 to 1.4.2 by @dependabot in https://github.com/kayac/ecspresso/pull/753
- Bump the aws-sdk-go-v2 group across 1 directory with 16 updates by @dependabot in https://github.com/kayac/ecspresso/pull/786
- use hosted arm runner. by @fujiwara in https://github.com/kayac/ecspresso/pull/791
- Bump github.com/hashicorp/go-slug from 0.15.0 to 0.16.3 by @dependabot in https://github.com/kayac/ecspresso/pull/790
- update modules by @fujiwara in https://github.com/kayac/ecspresso/pull/793
- fix nightly build by @fujiwara in https://github.com/kayac/ecspresso/pull/801
- fix: add support for AWS China ECR URLs by @litanyofmadness in https://github.com/kayac/ecspresso/pull/806
- test: add test case for AWS China ECR image by @fujiwara in https://github.com/kayac/ecspresso/pull/808
- remove healthCheckGracePeriodSeconds validation by @ijin in https://github.com/kayac/ecspresso/pull/805
- use Songmu/tagpr for release management. by @fujiwara in https://github.com/kayac/ecspresso/pull/809

## [v2.4.5](https://github.com/kayac/ecspresso/compare/v2.4.4...v2.4.5) - 2024-12-09
- Update aws-sdk-go-v2/service/ecs to v1.51.0 and go mod tidy by @t-kikuc in https://github.com/kayac/ecspresso/pull/774
- Supports AvailabilityZoneRebalancing. by @fujiwara in https://github.com/kayac/ecspresso/pull/778

## [v2.4.4](https://github.com/kayac/ecspresso/compare/v2.4.3...v2.4.4) - 2024-11-19
- Supports VPC Lattice integration. by @fujiwara in https://github.com/kayac/ecspresso/pull/773

## [v2.4.3](https://github.com/kayac/ecspresso/compare/v2.4.2...v2.4.3) - 2024-11-06
- fix panic on run --dry-run by @fujiwara in https://github.com/kayac/ecspresso/pull/767
- Bump github.com/golang-jwt/jwt/v4 from 4.5.0 to 4.5.1 by @dependabot in https://github.com/kayac/ecspresso/pull/766
- fix null pointer exception by @fujiwara in https://github.com/kayac/ecspresso/pull/768

## [v2.4.2](https://github.com/kayac/ecspresso/compare/v2.4.1...v2.4.2) - 2024-10-21
- returns error if --watch-container is not found. by @fujiwara in https://github.com/kayac/ecspresso/pull/750
- Bump github.com/aws/aws-sdk-go-v2/service/ecs to v1.47.4 by @fujiwara in https://github.com/kayac/ecspresso/pull/758

## [v2.4.1](https://github.com/kayac/ecspresso/compare/v2.4.0...v2.4.1) - 2024-08-30
- Fix typo: `task` -> `tasks` by @t-kikuc in https://github.com/kayac/ecspresso/pull/743
- fix resolving name of service connect namespace. by @fujiwara in https://github.com/kayac/ecspresso/pull/745

## [v2.4.0](https://github.com/kayac/ecspresso/compare/v2.3.6...v2.4.0) - 2024-08-06
- Bump github.com/aws/smithy-go from 1.20.2 to 1.20.3 by @dependabot in https://github.com/kayac/ecspresso/pull/711
- Add Jsonnet native functions by @fujiwara in https://github.com/kayac/ecspresso/pull/702
- run --revision and --latest-task-definition are exclusive by @fujiwara in https://github.com/kayac/ecspresso/pull/719
- remove fallback to ssm.GetParmater API by @fujiwara in https://github.com/kayac/ecspresso/pull/720
- bump versions by @fujiwara in https://github.com/kayac/ecspresso/pull/721
- add disables colorized output option by @ch1aki in https://github.com/kayac/ecspresso/pull/718
- adds test for #718 by @fujiwara in https://github.com/kayac/ecspresso/pull/725
- Add diff --external. Runs external diff command. by @fujiwara in https://github.com/kayac/ecspresso/pull/727
- Add ignore.tags into a configuration. by @fujiwara in https://github.com/kayac/ecspresso/pull/728
- Fix/retry registry by @fujiwara in https://github.com/kayac/ecspresso/pull/729
- Clarify README by @ijin in https://github.com/kayac/ecspresso/pull/731
- Exit non-zero status when deployment is rolled back. by @fujiwara in https://github.com/kayac/ecspresso/pull/733
- Bump the aws-sdk-go-v2 group across 1 directory with 15 updates by @dependabot in https://github.com/kayac/ecspresso/pull/735
- Bump github.com/schollz/progressbar/v3 from 3.13.1 to 3.14.6 by @dependabot in https://github.com/kayac/ecspresso/pull/734
- Bump github.com/goccy/go-yaml from 1.9.5 to 1.12.0 by @dependabot in https://github.com/kayac/ecspresso/pull/724
- Bump github.com/opencontainers/image-spec from 1.0.2 to 1.1.0 by @dependabot in https://github.com/kayac/ecspresso/pull/686
- Bump github.com/kayac/go-config from 0.6.0 to 0.7.0 by @dependabot in https://github.com/kayac/ecspresso/pull/648

## [v2.3.6](https://github.com/kayac/ecspresso/compare/v2.3.5...v2.3.6) - 2024-07-17
- --rm-dist has been deprecated in favor of --clean by @shogo82148 in https://github.com/kayac/ecspresso/pull/712
- fix: verify ssm secrets from SSM parameters. by @fujiwara in https://github.com/kayac/ecspresso/pull/713
- Add pidMode mapping to tdToTaskDefinitionInput function by @ch1aki in https://github.com/kayac/ecspresso/pull/715
- Bump goreleaser/goreleaser-action from 5 to 6 by @dependabot in https://github.com/kayac/ecspresso/pull/709
- Bump github.com/hashicorp/go-retryablehttp from 0.7.1 to 0.7.7 by @dependabot in https://github.com/kayac/ecspresso/pull/707
- Add IpcMode into TaskDefinitionInput. by @fujiwara in https://github.com/kayac/ecspresso/pull/716
- Bump github.com/Azure/azure-sdk-for-go/sdk/azidentity from 1.3.1 to 1.6.0 by @dependabot in https://github.com/kayac/ecspresso/pull/701

## [v2.3.5](https://github.com/kayac/ecspresso/compare/v2.3.4...v2.3.5) - 2024-06-21
- update aws-sdk-go-v2/service/ecs to v1.43.1 by @stkhr in https://github.com/kayac/ecspresso/pull/704
- Bump the aws-sdk-go-v2 group across 1 directory with 13 updates by @dependabot in https://github.com/kayac/ecspresso/pull/705

## [v2.3.4](https://github.com/kayac/ecspresso/compare/v2.3.3...v2.3.4) - 2024-05-23
- Add exec -L flag. by @fujiwara in https://github.com/kayac/ecspresso/pull/690
- Bump golang.org/x/net from 0.17.0 to 0.23.0 by @dependabot in https://github.com/kayac/ecspresso/pull/689
- Supports arm64 by actions.yml by @fujiwara in https://github.com/kayac/ecspresso/pull/693
- fix: panic when secrets.name or valueFrom is missing by @fujiwara in https://github.com/kayac/ecspresso/pull/697
- fix: add missing space between words in run command log messages by @nao23 in https://github.com/kayac/ecspresso/pull/698

## [v2.3.3](https://github.com/kayac/ecspresso/compare/v2.3.2...v2.3.3) - 2024-03-29
- Allow to specify os,arch of binary to install in CircleCI's orb by @tomiyan in https://github.com/kayac/ecspresso/pull/666
- Add revision option to deploy command by @tksx1227 in https://github.com/kayac/ecspresso/pull/672
- Fix wait for rollbacked deployment with CodeDeploy. by @fujiwara in https://github.com/kayac/ecspresso/pull/673
- refactor test/ci by @fujiwara in https://github.com/kayac/ecspresso/pull/674
- Set shorten waiter max delay. by @fujiwara in https://github.com/kayac/ecspresso/pull/675
- Bump google.golang.org/protobuf from 1.30.0 to 1.33.0 by @dependabot in https://github.com/kayac/ecspresso/pull/676
- Bump actions/setup-go from 4 to 5 by @dependabot in https://github.com/kayac/ecspresso/pull/653
- fix: use GetParameters instead of GetParameter to simulate actual ECS' behavior by @aereal in https://github.com/kayac/ecspresso/pull/678
- do rollbackTaskDefinition if a rollbacking deployment was completed. by @fujiwara in https://github.com/kayac/ecspresso/pull/679
- Go 1.22 by @fujiwara in https://github.com/kayac/ecspresso/pull/680
- fallback to ssm.GetParameter if failed to ssm.GetParameters by @fujiwara in https://github.com/kayac/ecspresso/pull/681

## [v2.3.2](https://github.com/kayac/ecspresso/compare/v2.3.1...v2.3.2) - 2024-01-19
- update aws-sdk-go-v2/service/ecs to v1.37.0 by @fujiwara in https://github.com/kayac/ecspresso/pull/656
- Fix typo by @itkq in https://github.com/kayac/ecspresso/pull/654
- fix verify ECR images in another region. by @fujiwara in https://github.com/kayac/ecspresso/pull/660
- Supports managing EBS Volumes by ECS services/tasks. by @fujiwara in https://github.com/kayac/ecspresso/pull/659
- verify multiple tag specifications for ManagedEBSVolume by @fujiwara in https://github.com/kayac/ecspresso/pull/661
- Add caching for secretsmanager_arn function by @fujiwara in https://github.com/kayac/ecspresso/pull/662

## [v2.3.1](https://github.com/kayac/ecspresso/compare/v2.3.0...v2.3.1) - 2023-12-25
- update tfstate-lookup v1.1.6 by @fujiwara in https://github.com/kayac/ecspresso/pull/646
- add tfstate testing by @fujiwara in https://github.com/kayac/ecspresso/pull/647

## [v2.3.0](https://github.com/kayac/ecspresso/compare/v2.2.4...v2.3.0) - 2023-12-21
- docs: add the installation guide with aqua by @suzuki-shunsuke in https://github.com/kayac/ecspresso/pull/616
- Bump golang.org/x/net from 0.14.0 to 0.17.0 by @dependabot in https://github.com/kayac/ecspresso/pull/617
- update aws-sdk-go-v2/service/ecs v1.33.0 by @fujiwara in https://github.com/kayac/ecspresso/pull/624
- Bump goreleaser/goreleaser-action from 4 to 5 by @dependabot in https://github.com/kayac/ecspresso/pull/610
- add run --client-token by @fujiwara in https://github.com/kayac/ecspresso/pull/631
- add secretsmanager plugin by @fujiwara in https://github.com/kayac/ecspresso/pull/618
- diff command works whenever a remote service or a task definition are not found. by @fujiwara in https://github.com/kayac/ecspresso/pull/632
- Enables to override timeout in a configuration file by --timeout. by @fujiwara in https://github.com/kayac/ecspresso/pull/633
- ECSPRESSO_FILTER_COMMAND moves to cli flag. by @fujiwara in https://github.com/kayac/ecspresso/pull/634
- fix diff output compared with nil by @fujiwara in https://github.com/kayac/ecspresso/pull/635
- Bump github.com/alecthomas/kong from 0.7.0 to 0.8.1 by @dependabot in https://github.com/kayac/ecspresso/pull/623
- Bump github.com/fatih/color from 1.13.0 to 1.16.0 by @dependabot in https://github.com/kayac/ecspresso/pull/629
- Bump google.golang.org/grpc from 1.49.0 to 1.56.3 by @dependabot in https://github.com/kayac/ecspresso/pull/619
- update aws-sdk-go-v2 and ecsta by @fujiwara in https://github.com/kayac/ecspresso/pull/636
- go 1.21 by @fujiwara in https://github.com/kayac/ecspresso/pull/640
- Default plugins (ssm and secretsmanager) by @fujiwara in https://github.com/kayac/ecspresso/pull/641
- Refactoring options by @fujiwara in https://github.com/kayac/ecspresso/pull/643
- V2.3 by @fujiwara in https://github.com/kayac/ecspresso/pull/642
- Bump golang.org/x/crypto from 0.14.0 to 0.17.0 by @dependabot in https://github.com/kayac/ecspresso/pull/644

## [v2.2.4](https://github.com/kayac/ecspresso/compare/v2.2.3...v2.2.4) - 2023-10-06
- fix: conversion typo by @testwill in https://github.com/kayac/ecspresso/pull/607
- Bump actions/checkout from 3 to 4 by @dependabot in https://github.com/kayac/ecspresso/pull/609
- update tfstate-lookup v1.1.4 by @fujiwara in https://github.com/kayac/ecspresso/pull/613
- fix typo escpresso -> ecspresso by @Kiryuanzu in https://github.com/kayac/ecspresso/pull/612

## [v2.2.3](https://github.com/kayac/ecspresso/compare/v2.2.2...v2.2.3) - 2023-08-04
- add init --task-definition flag. by @fujiwara in https://github.com/kayac/ecspresso/pull/593
- Remove runningCount and pendingCount from generated service definition. by @fujiwara in https://github.com/kayac/ecspresso/pull/594
- Add next step and template syntax section by @fujiwara in https://github.com/kayac/ecspresso/pull/600
- An output of task-definition of init command to stable. by @fujiwara in https://github.com/kayac/ecspresso/pull/601

## [v2.2.2](https://github.com/kayac/ecspresso/compare/v2.2.1...v2.2.2) - 2023-07-19
- Bump goreleaser/goreleaser-action from 3 to 4 by @dependabot in https://github.com/kayac/ecspresso/pull/577
- add "Supported tfstate URL format" by @fujiwara in https://github.com/kayac/ecspresso/pull/588
- update to ecsta v0.3.2 by @fujiwara in https://github.com/kayac/ecspresso/pull/589
- update to ecsta v0.3.3 by @fujiwara in https://github.com/kayac/ecspresso/pull/590
- Fix DesiredCount ignoring by @HASHIMOTO-Takafumi in https://github.com/kayac/ecspresso/pull/591
- Add tests for #591 by @fujiwara in https://github.com/kayac/ecspresso/pull/592

## [v2.2.1](https://github.com/kayac/ecspresso/compare/v2.2.0...v2.2.1) - 2023-06-19
- fix typo `--autos-caling` -> `--auto-scaling` by @sinsoku in https://github.com/kayac/ecspresso/pull/570
- use t.Setenv() in tests. by @fujiwara in https://github.com/kayac/ecspresso/pull/576
- update ecsta v0.3.1 by @fujiwara in https://github.com/kayac/ecspresso/pull/583
- Fix/verify errors on create log group by @fujiwara in https://github.com/kayac/ecspresso/pull/584
- Bump actions/setup-go from 3 to 4 by @dependabot in https://github.com/kayac/ecspresso/pull/578

## [v2.2.0](https://github.com/kayac/ecspresso/compare/v2.1.0...v2.2.0) - 2023-05-26
- Update README for asdf plugin by @koluku in https://github.com/kayac/ecspresso/pull/544
- Bump golang.org/x/text from 0.3.7 to 0.3.8 by @dependabot in https://github.com/kayac/ecspresso/pull/517
- create a log group when awslogs-create-group=="true" on verify by @fujiwara in https://github.com/kayac/ecspresso/pull/541
- Bump golang.org/x/net from 0.0.0-20220909164309-bea034e7d591 to 0.7.0 by @dependabot in https://github.com/kayac/ecspresso/pull/518
- Add example with terraform by @fujiwara in https://github.com/kayac/ecspresso/pull/556
- Fix typo: latst -> latest by @KOBA789 in https://github.com/kayac/ecspresso/pull/558
- Add deregister --revision=latest and --delete flag. by @fujiwara in https://github.com/kayac/ecspresso/pull/560
- Fix/tests/terraform by @fujiwara in https://github.com/kayac/ecspresso/pull/557
- Fix: Unable to update service with tags. by @fujiwara in https://github.com/kayac/ecspresso/pull/551
- implements deploy/scale --auto-scaling-(min|max) option. by @fujiwara in https://github.com/kayac/ecspresso/pull/550
- Install Ecspresso into toolchain cache directory by @goruha in https://github.com/kayac/ecspresso/pull/566
- Install Ecspresso into toolchain cache directory, show installed versions by @fujiwara in https://github.com/kayac/ecspresso/pull/567
- Refactor options type by @fujiwara in https://github.com/kayac/ecspresso/pull/565
- Bump github.com/schollz/progressbar/v3 from 3.11.0 to 3.13.1 by @dependabot in https://github.com/kayac/ecspresso/pull/535
- Bump github.com/aws/aws-sdk-go-v2/credentials from 1.13.15 to 1.13.24 by @dependabot in https://github.com/kayac/ecspresso/pull/561
- Bump github.com/aws/aws-sdk-go-v2/service/ecr from 1.17.17 to 1.18.11 by @dependabot in https://github.com/kayac/ecspresso/pull/564
- Bump github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 from 1.18.18 to 1.19.11 by @dependabot in https://github.com/kayac/ecspresso/pull/563
- fix nil pointer dereference by @fujiwara in https://github.com/kayac/ecspresso/pull/569

## [v2.1.0](https://github.com/kayac/ecspresso/compare/v2.0.5...v2.1.0) - 2023-03-31
- change log level in verifyLogConfiguration by @fujiwara in https://github.com/kayac/ecspresso/pull/529
- update to ecsta@v0.3.0 by @fujiwara in https://github.com/kayac/ecspresso/pull/531
- fix: verify SecretsManager JSON Key by @m22r in https://github.com/kayac/ecspresso/pull/533
- Add `--assume-role-arn` option to the top level of the CLI by @moznion in https://github.com/kayac/ecspresso/pull/530
- exit 1 when run a task failed to start. by @fujiwara in https://github.com/kayac/ecspresso/pull/537
- Add --assume-role-arn option by @fujiwara in https://github.com/kayac/ecspresso/pull/538
- fix map2str to stable. by @fujiwara in https://github.com/kayac/ecspresso/pull/542
- fix documents links by @fujiwara in https://github.com/kayac/ecspresso/pull/543
- ecspresso revisons --revision (current|latest|[number]) by @fujiwara in https://github.com/kayac/ecspresso/pull/539

## [v2.0.5](https://github.com/kayac/ecspresso/compare/v2.0.3...v2.0.5) - 2023-03-03
- nightly branch based on v2 by @fujiwara in https://github.com/kayac/ecspresso/pull/501
- Wait a service stable after create service. by @fujiwara in https://github.com/kayac/ecspresso/pull/502
- fix(action): authenticate the API calls for increasing API rate limit by @aereal in https://github.com/kayac/ecspresso/pull/499
- fix v2 action testing by @fujiwara in https://github.com/kayac/ecspresso/pull/503
- Fix supend / resume autoscaling by @fujiwara in https://github.com/kayac/ecspresso/pull/497
- returns ErrNotFound when CodeDeploy resources are not found. by @fujiwara in https://github.com/kayac/ecspresso/pull/504
- Fix for confusing flags of run. by @fujiwara in https://github.com/kayac/ecspresso/pull/505
- add delete --terminate flag. by @fujiwara in https://github.com/kayac/ecspresso/pull/512
- Fujiwara/logutils@v1.1.1 by @fujiwara in https://github.com/kayac/ecspresso/pull/516
- update aws-sdk-go-v2/config to fix SSO configuration issue by @enm10k in https://github.com/kayac/ecspresso/pull/520
- Output JSON message without using logger. by @fujiwara in https://github.com/kayac/ecspresso/pull/514
- Cache results of verified resources. by @fujiwara in https://github.com/kayac/ecspresso/pull/515
- add delete --terminate to test cases. by @fujiwara in https://github.com/kayac/ecspresso/pull/526
- bump go version 1.20 by @fujiwara in https://github.com/kayac/ecspresso/pull/527

## [v2.0.4](https://github.com/kayac/ecspresso/compare/v2.0.3...v2.0.4) - 2023-02-03
- nightly branch based on v2 by @fujiwara in https://github.com/kayac/ecspresso/pull/501
- Wait a service stable after create service. by @fujiwara in https://github.com/kayac/ecspresso/pull/502
- fix(action): authenticate the API calls for increasing API rate limit by @aereal in https://github.com/kayac/ecspresso/pull/499
- fix v2 action testing by @fujiwara in https://github.com/kayac/ecspresso/pull/503
- Fix supend / resume autoscaling by @fujiwara in https://github.com/kayac/ecspresso/pull/497
- returns ErrNotFound when CodeDeploy resources are not found. by @fujiwara in https://github.com/kayac/ecspresso/pull/504
- Fix for confusing flags of run. by @fujiwara in https://github.com/kayac/ecspresso/pull/505
- add delete --terminate flag. by @fujiwara in https://github.com/kayac/ecspresso/pull/512
