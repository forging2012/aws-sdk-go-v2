// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacemetering

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// MarketplaceMetering provides the API operation methods for making requests to
// AWSMarketplace Metering. See this package's package overview docs
// for details on the service.
//
// MarketplaceMetering methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type MarketplaceMetering struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*MarketplaceMetering)

// Used for custom request initialization logic
var initRequest func(*MarketplaceMetering, *aws.Request)

// Service information constants
const (
	ServiceName = "metering.marketplace" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName            // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the MarketplaceMetering client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a MarketplaceMetering client from just a config.
//     svc := marketplacemetering.New(myConfig)
//
//     // Create a MarketplaceMetering client with additional configuration
//     svc := marketplacemetering.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *MarketplaceMetering {
	var signingName string
	signingName = "aws-marketplace"
	signingRegion := config.Region

	svc := &MarketplaceMetering{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2016-01-14",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSMPMeteringService",
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

// newRequest creates a new request for a MarketplaceMetering operation and runs any
// custom request initialization.
func (c *MarketplaceMetering) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
