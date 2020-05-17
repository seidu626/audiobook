# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<a name="unreleased"></a>
## [Unreleased]


<a name="v0.3.1"></a>
## [v0.3.1] - 2020-02-16
### Chore
- **changelog:** fix changlog template
- **deploy:** using cbws/etcd-operator image
- **deps:** updated deps
- **makefile:** testing proto specific Makefile

### Docs
- **commitizen:** adding docs for commitizen

### Feat
- **logger:** now using micro's logger with zerolog provider

### Refactor
- **logger:** loger fields changed to map[string]interface{}
- **zerolog:** mode based config

### Test
- **logger:** fix unit tests


<a name="v0.3.0"></a>
## [v0.3.0] - 2020-02-11
### Build
- **changelog:** updated template

### Chore
- **changelog:** update changelog
- **deps:** zerolog v1.18.0
- **deps:** upgrade micro v2
- **deps:** updated deps
- **docs:** kube tools
- **proto:** dedicated Makefile for proto and dedicated CI/CD for proto
- **proto:** excludes e2e dir from Buf

### Docs
- **gitbook:** updated docs
- **gitbook:** updated docs
- **gitbook:** adding docs
- **gitbook:** adding docs
- **gitflow:** adding changelog

### Feat
- **account:** using GenderType enum in proto
- **account:** refactor account protos
- **deps:** lock grpc v1.26.0
- **makefile:** adding zerolog plugin of go-micro logger  & GORM
- **makefile:** format proto with vscode plugin : zxh404.vscode-proto3

### Fix
- **micro:** replace Deprecated: NewPublisher with NewEvent
- **proto:** refactor protobuf files. now using Buf instead of prototool
- **proto:** replace prototool with Buf


<a name="v0.2.9"></a>
## [v0.2.9] - 2019-11-29

<a name="v0.2.8"></a>
## [v0.2.8] - 2019-11-14
### Fix
- pkger support for makefile & Dockerfile


<a name="v0.2.7"></a>
## [v0.2.7] - 2019-11-14
### Deps
- lib update


<a name="v0.2.6"></a>
## [v0.2.6] - 2019-11-04
### Actions
- polish
- polish
- polish
- polish
- cleanup
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- testing
- refactor github actions [skip ci]
- fix Release Management pipeline
- adding Release Management pipeline
- test Cache go modules

### Deploy
- fix update EtcdCluster name dynamically
- grouped microservices together
- fix remove imagePullSecrets
- fix micro version [skip ci]

### Deps
- update base image busybox:1.31.1 [skip ci]
- update base image busybox:1.31.1 [skip ci]
- update deps
- update deps
- update di-> v2.3.0  [skip ci]

### Docs
- polish
- adding k9s
- adding k9s
- using github actions badge


<a name="v0.2.4"></a>
## [v0.2.4] - 2019-10-26
### Deps
- go micro v1.14.0
- go micro v1.14.0, testing --metadata

### Docs
- prereq  doc updated [skip ci]

### Googlepubsub
- changing PROJECT_ID -> GOOGLEPUBSUB_PROJECT_ID

### Makefile
- using DOCKER_BUILDKIT and BUILDKIT_INLINE_CACHE for fast builds

### Wrapper
- adding transaction wrapper
- adding transaction wrapper

### Wrappers
- add log wrappers for server, client, pubsub


<a name="v0.2.3"></a>
## [v0.2.3] - 2019-10-22
### Broker
- experimenting with googlepubsub plugin

### Docs
- update go-micro -> v1.13.2
- polish docs [skip ci]
- polish docs [skip ci]
- polish docs [skip ci]
- polish docs [skip ci]
- fix docs
- added docs for e2e testing

### E2e
- adding emailer e2e test


<a name="v0.2.2"></a>
## [v0.2.2] - 2019-10-19
### Ci
- release v0.2.2

### Consul
- removing consul and replacing with etcd

### Deps
- go-micro upgraded to v1.13.1

### Docs
- how to debug in VS Code
- how to debug in VS Code


<a name="v0.2.1"></a>
## [v0.2.1] - 2019-10-18
### Deps
- upgrade o-micro to v1.13.0

### Docs
- deployment doc updated [skip ci]
- deployment doc updated [skip ci]
- deployment doc updated [skip ci]


<a name="v0.2.0"></a>
## [v0.2.0] - 2019-10-17
### Deploy
- adding NATS and etcd Operator

### Deps
- upgraded to micro v1.12.0 and switching consul -> etcd

### Docs
- polish

### Shared
- moved go-plugins/client/selector/static to codebase, to avoid pulling tons of deps from micro's go-plugins


<a name="v0.1.5"></a>
## [v0.1.5] - 2019-10-14
### Cmd
- removing micro and moving to https://github.com/xmlking/micro
- removing micro and moving to https://github.com/xmlking/micro

### Plugin
- adding go-plugins/client/selector/static back. this allow us to use same way as we do with mDNS


<a name="v0.1.4"></a>
## [v0.1.4] - 2019-10-13
### Deploy
- consul: using non-root account with CONSUL_DISABLE_PERM_MGMT
- polish docs
- polish docs
- using CONSUL_DISABLE_PERM_MGMT
- adding istio testing
- adjust port names as per Instio conventions https://istio.io/docs/setup/additional-setup/requirements/
- fix makefile action: e2e  [skip ci]
- fix makefile action: e2e  [skip ci]
- fix makefile action: e2e  [skip ci]

### Deps
- micro updated to v1.11.1

### Docs
- polish docs
- polish deploy docs
- testing with BloomRPC UI Client
- testing with BloomRPC UI Client

### E2e
- exploring MICRO_SELECTOR=static MICRO_REGISTRY=memory options for e2e testing

### Make
- now support test-e2e, test-unit, test-inte, test-cover, test-bench actions

### Test
- adding docs for kubeval tool
- adding first e2e test case  [skip ci]
- adding first e2e test case  [skip ci]


<a name="v0.1.3"></a>
## [v0.1.3] - 2019-10-06
### Deploy
- GitHib Actions driven e2e tests on KinD
- renamed development overlay to e2e


<a name="v0.1.2"></a>
## [v0.1.2] - 2019-10-06
### Chore
- **deps:** update module jinzhu/gorm to v1.9.11
- **deps:** update google.golang.org/genproto commit hash to c459b9c

### Cicd
- using statefulset for consul, using docker.pkg.github.com for images
- prune release pipeline

### Deploy
- now you can generate release.yaml with # [@kustomize](https://github.com/kustomize) build deploy/overlays/production/ | sed -e "s|\$(NS)|default|g" -e "s|\$(IMAGE_VERSION)|v0.1.1-15-gbf7ad14|g" | kubectl apply -f -
- removing deploy/kustomization.yaml sothat we can use 'kustomize build ./deploy/overlays/staging'
- moved CONSUL_SERVICE_NAME to correct location
- mounting config.yaml with subPath

### Docs
- polish


<a name="v0.1.1"></a>
## [v0.1.1] - 2019-09-24
### Ci
- fuzzit
- moving consul to bases folder, sothat we can share this config with different overlays
- testing overlays/k8s
- more kustomization
- increase deadline for golangci-lint
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- setting GO111MODULE=on for Setup Tools step
- adding golangci-lint
- adding Setup Tools step
- fix Upload coverage step

### Cicd
- fix makefile
- adding tag and release action for Makefile
- adding docker build
- adding docker build

### Config
- InitConfig now takes configDir and configFile

### Deploy
- testing --address=:8088
- using health-checks from latest master
- deploying to k8s with consul

### Deployment
- make docker now support GCR registry
- make docker now support GCR registry

### Deps
- update deps to master
- go-micro -> v1.10.0
- libs update

### Docs
- updated
- updated Kustomize docs
- updated Kustomize docs
- updated Kustomize docs
- updated Kustomize docs
- updated Kustomize docs
- protoc-gen-gorm installation info
- updated main README
- adding Renovate dependency Status badge
- adding GitHub Actions Build Status badge
- prune
- fix - Remove all untagged images
- update README
- update README
- update README

### Emailer
- fix DI issues with Interface

### Gorm
- enabled auto_preload for profiles, adding index for deletedAt

### Kustomize
- updated docs for SopsSecret plugin
- adding health-sidecar
- switching to grpc
- tested and working  production overlay
- tested and working  production overlay
- trying with production overlay
- using env-vars-common
- moved vars

### Micro
- updated deps to v1.9.1

### Plugins
- moving  shared logs, congig plugins to main.go, and keep runtime only plugins main.go

### Test
- email Integration fixed
- using new strategy for Integration Tests
- using new strategy for Integration Tests
- using new strategy for Integration Tests


<a name="v0.1.0"></a>
## v0.1.0 - 2019-07-05
### Docs
- added steps when behind VPN


[Unreleased]: https://github.com/xmlking/micro-starter-kit/compare/v0.3.1...HEAD
[v0.3.1]: https://github.com/xmlking/micro-starter-kit/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.9...v0.3.0
[v0.2.9]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.8...v0.2.9
[v0.2.8]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.7...v0.2.8
[v0.2.7]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.6...v0.2.7
[v0.2.6]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.4...v0.2.6
[v0.2.4]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.3...v0.2.4
[v0.2.3]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.2...v0.2.3
[v0.2.2]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/xmlking/micro-starter-kit/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.5...v0.2.0
[v0.1.5]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.4...v0.1.5
[v0.1.4]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.3...v0.1.4
[v0.1.3]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.2...v0.1.3
[v0.1.2]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/xmlking/micro-starter-kit/compare/v0.1.0...v0.1.1
