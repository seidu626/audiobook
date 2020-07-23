module github.com/seidu626/audiobook/backend

go 1.14

// for local development, you can repoint go-micro to local development go-micro workspace
// replace github.com/micro/go-micro/v2 => /Users/schintha/Developer/Work/go/3rd-party/go-micro

require (
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/getsentry/sentry-go v0.6.1
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.4.2
	github.com/infobloxopen/atlas-app-toolkit v0.22.0
	github.com/jinzhu/gorm v1.9.12
	github.com/markbates/pkger v0.17.0
	github.com/micro/go-micro/v2 v2.6.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.19.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.5.1
	github.com/thoas/go-funk v0.6.0
	github.com/xmlking/configor v0.2.1
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/tools v0.0.0-20200626171337-aa94e735be7f // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)
