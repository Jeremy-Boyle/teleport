// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the permissions for your network interfaces.
func (c *Client) DescribeNetworkInterfacePermissions(ctx context.Context, params *DescribeNetworkInterfacePermissionsInput, optFns ...func(*Options)) (*DescribeNetworkInterfacePermissionsOutput, error) {
	if params == nil {
		params = &DescribeNetworkInterfacePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNetworkInterfacePermissions", params, optFns, c.addOperationDescribeNetworkInterfacePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNetworkInterfacePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeNetworkInterfacePermissions.
type DescribeNetworkInterfacePermissionsInput struct {

	// One or more filters.
	//
	// *
	// network-interface-permission.network-interface-permission-id - The ID of the
	// permission.
	//
	// * network-interface-permission.network-interface-id - The ID of the
	// network interface.
	//
	// * network-interface-permission.aws-account-id - The Amazon
	// Web Services account ID.
	//
	// * network-interface-permission.aws-service - The
	// Amazon Web Service.
	//
	// * network-interface-permission.permission - The type of
	// permission (INSTANCE-ATTACH | EIP-ASSOCIATE).
	Filters []types.Filter

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. If this
	// parameter is not specified, up to 50 results are returned by default.
	MaxResults *int32

	// One or more network interface permission IDs.
	NetworkInterfacePermissionIds []string

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Contains the output for DescribeNetworkInterfacePermissions.
type DescribeNetworkInterfacePermissionsOutput struct {

	// The network interface permissions.
	NetworkInterfacePermissions []types.NetworkInterfacePermission

	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNetworkInterfacePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeNetworkInterfacePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNetworkInterfacePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNetworkInterfacePermissions(options.Region), middleware.Before); err != nil {
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
	return nil
}

// DescribeNetworkInterfacePermissionsAPIClient is a client that implements the
// DescribeNetworkInterfacePermissions operation.
type DescribeNetworkInterfacePermissionsAPIClient interface {
	DescribeNetworkInterfacePermissions(context.Context, *DescribeNetworkInterfacePermissionsInput, ...func(*Options)) (*DescribeNetworkInterfacePermissionsOutput, error)
}

var _ DescribeNetworkInterfacePermissionsAPIClient = (*Client)(nil)

// DescribeNetworkInterfacePermissionsPaginatorOptions is the paginator options for
// DescribeNetworkInterfacePermissions
type DescribeNetworkInterfacePermissionsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. If this
	// parameter is not specified, up to 50 results are returned by default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNetworkInterfacePermissionsPaginator is a paginator for
// DescribeNetworkInterfacePermissions
type DescribeNetworkInterfacePermissionsPaginator struct {
	options   DescribeNetworkInterfacePermissionsPaginatorOptions
	client    DescribeNetworkInterfacePermissionsAPIClient
	params    *DescribeNetworkInterfacePermissionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeNetworkInterfacePermissionsPaginator returns a new
// DescribeNetworkInterfacePermissionsPaginator
func NewDescribeNetworkInterfacePermissionsPaginator(client DescribeNetworkInterfacePermissionsAPIClient, params *DescribeNetworkInterfacePermissionsInput, optFns ...func(*DescribeNetworkInterfacePermissionsPaginatorOptions)) *DescribeNetworkInterfacePermissionsPaginator {
	if params == nil {
		params = &DescribeNetworkInterfacePermissionsInput{}
	}

	options := DescribeNetworkInterfacePermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNetworkInterfacePermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNetworkInterfacePermissionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeNetworkInterfacePermissions page.
func (p *DescribeNetworkInterfacePermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNetworkInterfacePermissionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeNetworkInterfacePermissions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeNetworkInterfacePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNetworkInterfacePermissions",
	}
}
