module app

go 1.21.0

replace dslapp => ../app

require (
	dslapp v0.0.0
	github.com/PaddleHQ/go-aws-ssm v0.8.0
	github.com/aws/aws-lambda-go v1.19.1
	github.com/aws/aws-sdk-go-v2 v1.11.2
	github.com/aws/aws-sdk-go-v2/config v1.11.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.10.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.21.0
	github.com/aws/smithy-go v1.9.0
	github.com/megaproaktiv/awsmock v0.1.1
	gotest.tools v2.2.0+incompatible
)

require (
	github.com/aws/aws-sdk-go v1.16.24 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.6.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.3.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.5.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.9.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.11.1 // indirect
	github.com/gabyx/githooks/githooks v1.1.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
)
