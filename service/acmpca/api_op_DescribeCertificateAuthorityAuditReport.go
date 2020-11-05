// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists information about a specific audit report created by calling the
// CreateCertificateAuthorityAuditReport
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html)
// action. Audit information is created every time the certificate authority (CA)
// private key is used. The private key is used when you call the IssueCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_IssueCertificate.html)
// action or the RevokeCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_RevokeCertificate.html)
// action.
func (c *Client) DescribeCertificateAuthorityAuditReport(ctx context.Context, params *DescribeCertificateAuthorityAuditReportInput, optFns ...func(*Options)) (*DescribeCertificateAuthorityAuditReportOutput, error) {
	if params == nil {
		params = &DescribeCertificateAuthorityAuditReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCertificateAuthorityAuditReport", params, optFns, addOperationDescribeCertificateAuthorityAuditReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCertificateAuthorityAuditReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCertificateAuthorityAuditReportInput struct {

	// The report ID returned by calling the CreateCertificateAuthorityAuditReport
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html)
	// action.
	//
	// This member is required.
	AuditReportId *string

	// The Amazon Resource Name (ARN) of the private CA. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type DescribeCertificateAuthorityAuditReportOutput struct {

	// Specifies whether report creation is in progress, has succeeded, or has failed.
	AuditReportStatus types.AuditReportStatus

	// The date and time at which the report was created.
	CreatedAt *time.Time

	// Name of the S3 bucket that contains the report.
	S3BucketName *string

	// S3 key that uniquely identifies the report file in your S3 bucket.
	S3Key *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCertificateAuthorityAuditReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCertificateAuthorityAuditReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCertificateAuthorityAuditReport{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpDescribeCertificateAuthorityAuditReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificateAuthorityAuditReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCertificateAuthorityAuditReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "DescribeCertificateAuthorityAuditReport",
	}
}
