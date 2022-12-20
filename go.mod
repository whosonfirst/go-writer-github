module github.com/whosonfirst/go-writer-github/v3

go 1.18

// Important: Until things settle down we need to peg ourselves to:
// golang.org/x/oauth2 v0.0.0-20220808172628-8227340efae7
// Because if we don't we run it these errors:
// go: found github.com/sfomuseum/runtimevar in github.com/sfomuseum/runtimevar v1.0.4
//	github.com/whosonfirst/go-writer-github/v3 imports
//	github.com/sfomuseum/runtimevar imports
//	gocloud.dev/runtimevar/awsparamstore tested by
//	gocloud.dev/runtimevar/awsparamstore.test imports
//	gocloud.dev/internal/testing/setup imports
//	golang.org/x/oauth2/google imports
//	cloud.google.com/go/compute/metadata: ambiguous import: found package cloud.google.com/go/compute/metadata in multiple modules:
//	cloud.google.com/go/compute v1.7.0 (~/go/pkg/mod/cloud.google.com/go/compute@v1.7.0/metadata)
//	cloud.google.com/go/compute/metadata v0.2.0 (~/go/pkg/mod/cloud.google.com/go/compute/metadata@v0.2.0)

require (
	github.com/google/go-github/v48 v48.2.0
	github.com/sfomuseum/runtimevar v1.0.4
	github.com/whosonfirst/go-ioutil v1.0.2
	github.com/whosonfirst/go-writer/v3 v3.1.0
	gocloud.dev v0.27.0
	golang.org/x/oauth2 v0.0.0-20220808172628-8227340efae7
)

require (
	github.com/aaronland/go-aws-session v0.1.0 // indirect
	github.com/aaronland/go-roster v1.0.0 // indirect
	github.com/aaronland/go-string v1.0.0 // indirect
	github.com/aws/aws-sdk-go v1.44.163 // indirect
	github.com/aws/aws-sdk-go-v2 v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.15.15 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v1.27.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.10 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/g8rswimmer/error-chain v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/natefinch/atomic v1.0.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/api v0.91.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220802133213-ce4fa296bf78 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
