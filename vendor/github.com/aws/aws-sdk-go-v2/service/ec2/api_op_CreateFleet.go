// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an EC2 Fleet that contains the configuration information for On-Demand
// Instances and Spot Instances. Instances are launched immediately if there is
// available capacity.
//
// A single EC2 Fleet can include multiple launch specifications that vary by
// instance type, AMI, Availability Zone, or subnet.
//
// For more information, see [EC2 Fleet] in the Amazon EC2 User Guide.
//
// [EC2 Fleet]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet.html
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	if params == nil {
		params = &CreateFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleet", params, optFns, c.addOperationCreateFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetInput struct {

	// The configuration for the EC2 Fleet.
	//
	// This member is required.
	LaunchTemplateConfigs []types.FleetLaunchTemplateConfigRequest

	// The number of units to request.
	//
	// This member is required.
	TargetCapacitySpecification *types.TargetCapacitySpecificationRequest

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see [Ensuring idempotency].
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// Reserved.
	Context *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Indicates whether running instances should be terminated if the total target
	// capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.
	//
	// Supported only for fleets of type maintain .
	ExcessCapacityTerminationPolicy types.FleetExcessCapacityTerminationPolicy

	// Describes the configuration of On-Demand Instances in an EC2 Fleet.
	OnDemandOptions *types.OnDemandOptionsRequest

	// Indicates whether EC2 Fleet should replace unhealthy Spot Instances. Supported
	// only for fleets of type maintain . For more information, see [EC2 Fleet health checks] in the Amazon EC2
	// User Guide.
	//
	// [EC2 Fleet health checks]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#ec2-fleet-health-checks
	ReplaceUnhealthyInstances *bool

	// Describes the configuration of Spot Instances in an EC2 Fleet.
	SpotOptions *types.SpotOptionsRequest

	// The key-value pair for tagging the EC2 Fleet request on creation. For more
	// information, see [Tag your resources].
	//
	// If the fleet type is instant , specify a resource type of fleet to tag the
	// fleet or instance to tag the instances at launch.
	//
	// If the fleet type is maintain or request , specify a resource type of fleet to
	// tag the fleet. You cannot specify a resource type of instance . To tag instances
	// at launch, specify the tags in a [launch template].
	//
	// [launch template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template
	// [Tag your resources]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-resources
	TagSpecifications []types.TagSpecification

	// Indicates whether running instances should be terminated when the EC2 Fleet
	// expires.
	TerminateInstancesWithExpiration *bool

	// The fleet type. The default value is maintain .
	//
	//   - maintain - The EC2 Fleet places an asynchronous request for your desired
	//   capacity, and continues to maintain your desired Spot capacity by replenishing
	//   interrupted Spot Instances.
	//
	//   - request - The EC2 Fleet places an asynchronous one-time request for your
	//   desired capacity, but does submit Spot requests in alternative capacity pools if
	//   Spot capacity is unavailable, and does not maintain Spot capacity if Spot
	//   Instances are interrupted.
	//
	//   - instant - The EC2 Fleet places a synchronous one-time request for your
	//   desired capacity, and returns errors for any instances that could not be
	//   launched.
	//
	// For more information, see [EC2 Fleet request types] in the Amazon EC2 User Guide.
	//
	// [EC2 Fleet request types]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-request-type.html
	Type types.FleetType

	// The start date and time of the request, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request
	// immediately.
	ValidFrom *time.Time

	// The end date and time of the request, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ). At this point, no new EC2 Fleet requests are placed or
	// able to fulfill the request. If no value is specified, the request remains until
	// you cancel it.
	ValidUntil *time.Time

	noSmithyDocumentSerde
}

type CreateFleetOutput struct {

	// Information about the instances that could not be launched by the fleet.
	// Supported only for fleets of type instant .
	Errors []types.CreateFleetError

	// The ID of the EC2 Fleet.
	FleetId *string

	// Information about the instances that were launched by the fleet. Supported only
	// for fleets of type instant .
	Instances []types.CreateFleetInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFleet"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFleet",
	}
}