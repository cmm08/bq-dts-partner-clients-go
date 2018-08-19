// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package datatransfer

import (
	"math"
	"time"

	"cloud.google.com/go/internal/version"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	datatransferpb "google.golang.org/genproto/googleapis/cloud/bigquery/datatransfer/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// DataSourceCallOptions contains the retry settings for each method of DataSourceClient.
type DataSourceCallOptions struct {
	UpdateTransferRun          []gax.CallOption
	LogTransferRunMessages     []gax.CallOption
	StartBigQueryJobs          []gax.CallOption
	GetCredentials             []gax.CallOption
	FinishRun                  []gax.CallOption
	CreateDataSourceDefinition []gax.CallOption
	UpdateDataSourceDefinition []gax.CallOption
	DeleteDataSourceDefinition []gax.CallOption
	GetDataSourceDefinition    []gax.CallOption
	ListDataSourceDefinitions  []gax.CallOption
}

func defaultDataSourceClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("bigquerydatatransfer.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultDataSourceCallOptions() *DataSourceCallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &DataSourceCallOptions{
		UpdateTransferRun:          retry[[2]string{"default", "non_idempotent"}],
		LogTransferRunMessages:     retry[[2]string{"default", "non_idempotent"}],
		StartBigQueryJobs:          retry[[2]string{"default", "non_idempotent"}],
		GetCredentials:             retry[[2]string{"default", "idempotent"}],
		FinishRun:                  retry[[2]string{"default", "non_idempotent"}],
		CreateDataSourceDefinition: retry[[2]string{"default", "non_idempotent"}],
		UpdateDataSourceDefinition: retry[[2]string{"default", "non_idempotent"}],
		DeleteDataSourceDefinition: retry[[2]string{"default", "idempotent"}],
		GetDataSourceDefinition:    retry[[2]string{"default", "idempotent"}],
		ListDataSourceDefinitions:  retry[[2]string{"default", "idempotent"}],
	}
}

// DataSourceClient is a client for interacting with BigQuery Data Transfer API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type DataSourceClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	dataSourceClient datatransferpb.DataSourceServiceClient

	// The call options for this service.
	CallOptions *DataSourceCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDataSourceClient creates a new data source service client.
//
// The Google BigQuery Data Transfer API allows BigQuery users to
// configure transfer of their data from other Google Products into BigQuery.
// This service exposes methods that should be used by data source backend.
func NewDataSourceClient(ctx context.Context, opts ...option.ClientOption) (*DataSourceClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultDataSourceClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &DataSourceClient{
		conn:        conn,
		CallOptions: defaultDataSourceCallOptions(),

		dataSourceClient: datatransferpb.NewDataSourceServiceClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *DataSourceClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DataSourceClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DataSourceClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", version.Go()}, keyval...)
	kv = append(kv, "gapic", version.Repo, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// UpdateTransferRun update a transfer run. If successful, resets
// data_source.update_deadline_seconds timer.
func (c *DataSourceClient) UpdateTransferRun(ctx context.Context, req *datatransferpb.UpdateTransferRunRequest, opts ...gax.CallOption) (*datatransferpb.TransferRun, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateTransferRun[0:len(c.CallOptions.UpdateTransferRun):len(c.CallOptions.UpdateTransferRun)], opts...)
	var resp *datatransferpb.TransferRun
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dataSourceClient.UpdateTransferRun(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LogTransferRunMessages log messages for a transfer run. If successful (at least 1 message), resets
// data_source.update_deadline_seconds timer.
func (c *DataSourceClient) LogTransferRunMessages(ctx context.Context, req *datatransferpb.LogTransferRunMessagesRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.LogTransferRunMessages[0:len(c.CallOptions.LogTransferRunMessages):len(c.CallOptions.LogTransferRunMessages)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.dataSourceClient.LogTransferRunMessages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// StartBigQueryJobs notify the Data Transfer Service that data is ready for loading.
// The Data Transfer Service will start and monitor multiple BigQuery Load
// jobs for a transfer run. Monitored jobs will be automatically retried
// and produce log messages when starting and finishing a job.
// Can be called multiple times for the same transfer run.
func (c *DataSourceClient) StartBigQueryJobs(ctx context.Context, req *datatransferpb.StartBigQueryJobsRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.StartBigQueryJobs[0:len(c.CallOptions.StartBigQueryJobs):len(c.CallOptions.StartBigQueryJobs)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.dataSourceClient.StartBigQueryJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// GetCredentials get user authentication token so that the data source can perform
// operations on behalf of the user.
// Can only be called by configured data source service account.
// In all other cases - it will fail with permission denied error.
func (c *DataSourceClient) GetCredentials(ctx context.Context, req *datatransferpb.GetCredentialsRequest, opts ...gax.CallOption) (*datatransferpb.Credentials, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetCredentials[0:len(c.CallOptions.GetCredentials):len(c.CallOptions.GetCredentials)], opts...)
	var resp *datatransferpb.Credentials
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dataSourceClient.GetCredentials(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FinishRun notify the Data Transfer Service that the data source is done processing
// the run. No more status updates or requests to start/monitor jobs will be
// accepted. The run will be finalized by the Data Transfer Service when all
// monitored jobs are completed.
// Does not need to be called if the run is set to FAILED.
func (c *DataSourceClient) FinishRun(ctx context.Context, req *datatransferpb.FinishRunRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.FinishRun[0:len(c.CallOptions.FinishRun):len(c.CallOptions.FinishRun)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.dataSourceClient.FinishRun(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CreateDataSourceDefinition creates a data source definition.  Calling this method will automatically
// use your credentials to create the following Google Cloud resources in
// YOUR Google Cloud project.
// 1. OAuth client
// 2. Pub/Sub Topics and Subscriptions in each supported_location_ids. e.g.,
// projects/{project_id}/{topics|subscriptions}/bigquerydatatransfer.{data_source_id}.{location_id}.run
// The field data_source.client_id should be left empty in the input request,
// as the API will create a new OAuth client on behalf of the caller. On the
// other hand data_source.scopes usually need to be set when there are OAuth
// scopes that need to be granted by end users.
func (c *DataSourceClient) CreateDataSourceDefinition(ctx context.Context, req *datatransferpb.CreateDataSourceDefinitionRequest, opts ...gax.CallOption) (*datatransferpb.DataSourceDefinition, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateDataSourceDefinition[0:len(c.CallOptions.CreateDataSourceDefinition):len(c.CallOptions.CreateDataSourceDefinition)], opts...)
	var resp *datatransferpb.DataSourceDefinition
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dataSourceClient.CreateDataSourceDefinition(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateDataSourceDefinition updates an existing data source definition. If changing
// supported_location_ids, triggers same effects as mentioned in "Create a
// data source definition."
func (c *DataSourceClient) UpdateDataSourceDefinition(ctx context.Context, req *datatransferpb.UpdateDataSourceDefinitionRequest, opts ...gax.CallOption) (*datatransferpb.DataSourceDefinition, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateDataSourceDefinition[0:len(c.CallOptions.UpdateDataSourceDefinition):len(c.CallOptions.UpdateDataSourceDefinition)], opts...)
	var resp *datatransferpb.DataSourceDefinition
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dataSourceClient.UpdateDataSourceDefinition(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteDataSourceDefinition deletes a data source definition, all of the transfer configs associated
// with this data source definition (if any) must be deleted first by the user
// in ALL regions, in order to delete the data source definition.
// This method is primarily meant for deleting data sources created during
// testing stage.
// If the data source is referenced by transfer configs in the region
// specified in the request URL, the method will fail immediately. If in the
// current region (e.g., US) it's not used by any transfer configs, but in
// another region (e.g., EU) it is, then although the method will succeed in
// region US, but it will fail when the deletion operation is replicated to
// region EU. And eventually, the system will replicate the data source
// definition back from EU to US, in order to bring all regions to
// consistency. The final effect is that the data source appears to be
// 'undeleted' in the US region.
func (c *DataSourceClient) DeleteDataSourceDefinition(ctx context.Context, req *datatransferpb.DeleteDataSourceDefinitionRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.DeleteDataSourceDefinition[0:len(c.CallOptions.DeleteDataSourceDefinition):len(c.CallOptions.DeleteDataSourceDefinition)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.dataSourceClient.DeleteDataSourceDefinition(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// GetDataSourceDefinition retrieves an existing data source definition.
func (c *DataSourceClient) GetDataSourceDefinition(ctx context.Context, req *datatransferpb.GetDataSourceDefinitionRequest, opts ...gax.CallOption) (*datatransferpb.DataSourceDefinition, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetDataSourceDefinition[0:len(c.CallOptions.GetDataSourceDefinition):len(c.CallOptions.GetDataSourceDefinition)], opts...)
	var resp *datatransferpb.DataSourceDefinition
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dataSourceClient.GetDataSourceDefinition(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListDataSourceDefinitions lists supported data source definitions.
func (c *DataSourceClient) ListDataSourceDefinitions(ctx context.Context, req *datatransferpb.ListDataSourceDefinitionsRequest, opts ...gax.CallOption) *DataSourceDefinitionIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListDataSourceDefinitions[0:len(c.CallOptions.ListDataSourceDefinitions):len(c.CallOptions.ListDataSourceDefinitions)], opts...)
	it := &DataSourceDefinitionIterator{}
	req = proto.Clone(req).(*datatransferpb.ListDataSourceDefinitionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*datatransferpb.DataSourceDefinition, string, error) {
		var resp *datatransferpb.ListDataSourceDefinitionsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.dataSourceClient.ListDataSourceDefinitions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.DataSourceDefinitions, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// DataSourceDefinitionIterator manages a stream of *datatransferpb.DataSourceDefinition.
type DataSourceDefinitionIterator struct {
	items    []*datatransferpb.DataSourceDefinition
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*datatransferpb.DataSourceDefinition, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DataSourceDefinitionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DataSourceDefinitionIterator) Next() (*datatransferpb.DataSourceDefinition, error) {
	var item *datatransferpb.DataSourceDefinition
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DataSourceDefinitionIterator) bufLen() int {
	return len(it.items)
}

func (it *DataSourceDefinitionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
