// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// MachineLearning provides the API operation methods for making requests to
// Amazon Machine Learning. See this package's package overview docs
// for details on the service.
//
// MachineLearning methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type MachineLearning struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*MachineLearning)

// Used for custom request initialization logic
var initRequest func(*MachineLearning, *aws.Request)

// Service information constants
const (
	ServiceName = "machinelearning" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName       // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the MachineLearning client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a MachineLearning client from just a config.
//     svc := machinelearning.New(myConfig)
//
//     // Create a MachineLearning client with additional configuration
//     svc := machinelearning.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *MachineLearning {
	var signingName string
	signingRegion := config.Region

	svc := &MachineLearning{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2014-12-12",
				JSONVersion:   "1.1",
				TargetPrefix:  "AmazonML_20141212",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a MachineLearning operation and runs any
// custom request initialization.
func (c *MachineLearning) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
