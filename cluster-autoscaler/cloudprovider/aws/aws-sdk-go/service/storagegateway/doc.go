// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package storagegateway provides the client and types for making API
// requests to AWS Storage Gateway.
//
// Storage Gateway is the service that connects an on-premises software appliance
// with cloud-based storage to provide seamless and secure integration between
// an organization's on-premises IT environment and the Amazon Web Services
// storage infrastructure. The service enables you to securely upload data to
// the Amazon Web Services Cloud for cost effective backup and rapid disaster
// recovery.
//
// Use the following links to get started using the Storage Gateway Service
// API Reference:
//
//    * Storage Gateway required request headers (https://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPI.html#AWSStorageGatewayHTTPRequestsHeaders):
//    Describes the required headers that you must send with every POST request
//    to Storage Gateway.
//
//    * Signing requests (https://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPI.html#AWSStorageGatewaySigningRequests):
//    Storage Gateway requires that you authenticate every request you send;
//    this topic describes how sign such a request.
//
//    * Error responses (https://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPI.html#APIErrorResponses):
//    Provides reference information about Storage Gateway errors.
//
//    * Operations in Storage Gateway (https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_Operations.html):
//    Contains detailed descriptions of all Storage Gateway operations, their
//    request parameters, response elements, possible errors, and examples of
//    requests and responses.
//
//    * Storage Gateway endpoints and quotas (https://docs.aws.amazon.com/general/latest/gr/sg.html):
//    Provides a list of each Amazon Web Services Region and the endpoints available
//    for use with Storage Gateway.
//
// Storage Gateway resource IDs are in uppercase. When you use these resource
// IDs with the Amazon EC2 API, EC2 expects resource IDs in lowercase. You must
// change your resource ID to lowercase to use it with the EC2 API. For example,
// in Storage Gateway the ID for a volume might be vol-AA22BB012345DAF670. When
// you use this ID with the EC2 API, you must change it to vol-aa22bb012345daf670.
// Otherwise, the EC2 API might not behave as expected.
//
// IDs for Storage Gateway volumes and Amazon EBS snapshots created from gateway
// volumes are changing to a longer format. Starting in December 2016, all new
// volumes and snapshots will be created with a 17-character string. Starting
// in April 2016, you will be able to use these longer IDs so you can test your
// systems with the new format. For more information, see Longer EC2 and EBS
// resource IDs (http://aws.amazon.com/ec2/faqs/#longer-ids).
//
// For example, a volume Amazon Resource Name (ARN) with the longer volume ID
// format looks like the following:
//
// arn:aws:storagegateway:us-west-2:111122223333:gateway/sgw-12A3456B/volume/vol-1122AABBCCDDEEFFG.
//
// A snapshot ID with the longer ID format looks like the following: snap-78e226633445566ee.
//
// For more information, see Announcement: Heads-up – Longer Storage Gateway
// volume and snapshot IDs coming in 2016 (http://forums.aws.amazon.com/ann.jspa?annID=3557).
//
// See https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30 for more information on this service.
//
// See storagegateway package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/storagegateway/
//
// Using the Client
//
// To contact AWS Storage Gateway with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Storage Gateway client StorageGateway for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/storagegateway/#New
package storagegateway
