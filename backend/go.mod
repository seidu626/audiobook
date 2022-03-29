module github.com/seidu626/audiobook/backend

go 1.14

require google.golang.org/grpc v1.26.0

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

require (
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/getsentry/sentry-go v0.7.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.5 // indirect
	github.com/infobloxopen/atlas-app-toolkit v0.22.0
	github.com/jinzhu/gorm v1.9.15
	github.com/markbates/pkger v0.17.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.19.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	github.com/xmlking/configor v0.2.1
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/protobuf v1.23.0
)
