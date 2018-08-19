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
	emptypb "github.com/golang/protobuf/ptypes/empty"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	datatransferpb "google.golang.org/genproto/googleapis/cloud/bigquery/datatransfer/v1"
	field_maskpb "google.golang.org/genproto/protobuf/field_mask"
)

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockDataTransferServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	datatransferpb.DataTransferServiceServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockDataTransferServer) GetDataSource(ctx context.Context, req *datatransferpb.GetDataSourceRequest) (*datatransferpb.DataSource, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.DataSource), nil
}

func (s *mockDataTransferServer) ListDataSources(ctx context.Context, req *datatransferpb.ListDataSourcesRequest) (*datatransferpb.ListDataSourcesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.ListDataSourcesResponse), nil
}

func (s *mockDataTransferServer) CreateTransferConfig(ctx context.Context, req *datatransferpb.CreateTransferConfigRequest) (*datatransferpb.TransferConfig, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.TransferConfig), nil
}

func (s *mockDataTransferServer) UpdateTransferConfig(ctx context.Context, req *datatransferpb.UpdateTransferConfigRequest) (*datatransferpb.TransferConfig, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.TransferConfig), nil
}

func (s *mockDataTransferServer) DeleteTransferConfig(ctx context.Context, req *datatransferpb.DeleteTransferConfigRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockDataTransferServer) GetTransferConfig(ctx context.Context, req *datatransferpb.GetTransferConfigRequest) (*datatransferpb.TransferConfig, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.TransferConfig), nil
}

func (s *mockDataTransferServer) ListTransferConfigs(ctx context.Context, req *datatransferpb.ListTransferConfigsRequest) (*datatransferpb.ListTransferConfigsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.ListTransferConfigsResponse), nil
}

func (s *mockDataTransferServer) ScheduleTransferRuns(ctx context.Context, req *datatransferpb.ScheduleTransferRunsRequest) (*datatransferpb.ScheduleTransferRunsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.ScheduleTransferRunsResponse), nil
}

func (s *mockDataTransferServer) GetTransferRun(ctx context.Context, req *datatransferpb.GetTransferRunRequest) (*datatransferpb.TransferRun, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.TransferRun), nil
}

func (s *mockDataTransferServer) DeleteTransferRun(ctx context.Context, req *datatransferpb.DeleteTransferRunRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockDataTransferServer) ListTransferRuns(ctx context.Context, req *datatransferpb.ListTransferRunsRequest) (*datatransferpb.ListTransferRunsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.ListTransferRunsResponse), nil
}

func (s *mockDataTransferServer) ListTransferLogs(ctx context.Context, req *datatransferpb.ListTransferLogsRequest) (*datatransferpb.ListTransferLogsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.ListTransferLogsResponse), nil
}

func (s *mockDataTransferServer) CheckValidCreds(ctx context.Context, req *datatransferpb.CheckValidCredsRequest) (*datatransferpb.CheckValidCredsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.CheckValidCredsResponse), nil
}

func (s *mockDataTransferServer) EnableDataTransferService(ctx context.Context, req *datatransferpb.EnableDataTransferServiceRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockDataTransferServer) IsDataTransferServiceEnabled(ctx context.Context, req *datatransferpb.IsDataTransferServiceEnabledRequest) (*datatransferpb.IsDataTransferServiceEnabledResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.IsDataTransferServiceEnabledResponse), nil
}

type mockDataSourceServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	datatransferpb.DataSourceServiceServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockDataSourceServer) UpdateTransferRun(ctx context.Context, req *datatransferpb.UpdateTransferRunRequest) (*datatransferpb.TransferRun, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.TransferRun), nil
}

func (s *mockDataSourceServer) LogTransferRunMessages(ctx context.Context, req *datatransferpb.LogTransferRunMessagesRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockDataSourceServer) StartBigQueryJobs(ctx context.Context, req *datatransferpb.StartBigQueryJobsRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockDataSourceServer) GetCredentials(ctx context.Context, req *datatransferpb.GetCredentialsRequest) (*datatransferpb.Credentials, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.Credentials), nil
}

func (s *mockDataSourceServer) FinishRun(ctx context.Context, req *datatransferpb.FinishRunRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockDataSourceServer) CreateDataSourceDefinition(ctx context.Context, req *datatransferpb.CreateDataSourceDefinitionRequest) (*datatransferpb.DataSourceDefinition, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.DataSourceDefinition), nil
}

func (s *mockDataSourceServer) UpdateDataSourceDefinition(ctx context.Context, req *datatransferpb.UpdateDataSourceDefinitionRequest) (*datatransferpb.DataSourceDefinition, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.DataSourceDefinition), nil
}

func (s *mockDataSourceServer) DeleteDataSourceDefinition(ctx context.Context, req *datatransferpb.DeleteDataSourceDefinitionRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockDataSourceServer) GetDataSourceDefinition(ctx context.Context, req *datatransferpb.GetDataSourceDefinitionRequest) (*datatransferpb.DataSourceDefinition, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.DataSourceDefinition), nil
}

func (s *mockDataSourceServer) ListDataSourceDefinitions(ctx context.Context, req *datatransferpb.ListDataSourceDefinitionsRequest) (*datatransferpb.ListDataSourceDefinitionsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*datatransferpb.ListDataSourceDefinitionsResponse), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockDataTransfer mockDataTransferServer
	mockDataSource   mockDataSourceServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	datatransferpb.RegisterDataTransferServiceServer(serv, &mockDataTransfer)
	datatransferpb.RegisterDataSourceServiceServer(serv, &mockDataSource)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestDataTransferServiceGetDataSource(t *testing.T) {
	var name2 string = "name2-1052831874"
	var dataSourceId string = "dataSourceId-1015796374"
	var displayName string = "displayName1615086568"
	var description string = "description-1724546052"
	var clientId string = "clientId-1904089585"
	var supportsMultipleTransfers bool = true
	var updateDeadlineSeconds int32 = 991471694
	var defaultSchedule string = "defaultSchedule-800168235"
	var supportsCustomSchedule bool = true
	var helpUrl string = "helpUrl-789431439"
	var defaultDataRefreshWindowDays int32 = 1804935157
	var manualRunsDisabled bool = true
	var partnerLegalName string = "partnerLegalName-1307326424"
	var expectedResponse = &datatransferpb.DataSource{
		Name:                      name2,
		DataSourceId:              dataSourceId,
		DisplayName:               displayName,
		Description:               description,
		ClientId:                  clientId,
		SupportsMultipleTransfers: supportsMultipleTransfers,
		UpdateDeadlineSeconds:     updateDeadlineSeconds,
		DefaultSchedule:           defaultSchedule,
		SupportsCustomSchedule:    supportsCustomSchedule,
		HelpUrl:                   helpUrl,
		DefaultDataRefreshWindowDays: defaultDataRefreshWindowDays,
		ManualRunsDisabled:           manualRunsDisabled,
		PartnerLegalName:             partnerLegalName,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSources/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE]")
	var request = &datatransferpb.GetDataSourceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetDataSource(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceGetDataSourceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSources/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE]")
	var request = &datatransferpb.GetDataSourceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetDataSource(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceListDataSources(t *testing.T) {
	var nextPageToken string = ""
	var dataSourcesElement *datatransferpb.DataSource = &datatransferpb.DataSource{}
	var dataSources = []*datatransferpb.DataSource{dataSourcesElement}
	var expectedResponse = &datatransferpb.ListDataSourcesResponse{
		NextPageToken: nextPageToken,
		DataSources:   dataSources,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.ListDataSourcesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListDataSources(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.DataSources[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceListDataSourcesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.ListDataSourcesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListDataSources(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceCreateTransferConfig(t *testing.T) {
	var name string = "name3373707"
	var destinationDatasetId string = "destinationDatasetId1541564179"
	var displayName string = "displayName1615086568"
	var dataSourceId string = "dataSourceId-1015796374"
	var schedule string = "schedule-697920873"
	var dataRefreshWindowDays int32 = 327632845
	var disabled bool = true
	var userId int64 = 147132913
	var datasetRegion string = "datasetRegion959248539"
	var expectedResponse = &datatransferpb.TransferConfig{
		Name:                  name,
		DestinationDatasetId:  destinationDatasetId,
		DisplayName:           displayName,
		DataSourceId:          dataSourceId,
		Schedule:              schedule,
		DataRefreshWindowDays: dataRefreshWindowDays,
		Disabled:              disabled,
		UserId:                userId,
		DatasetRegion:         datasetRegion,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var transferConfig *datatransferpb.TransferConfig = &datatransferpb.TransferConfig{}
	var authorizationCode string = "authorizationCode1571154419"
	var request = &datatransferpb.CreateTransferConfigRequest{
		Parent:            formattedParent,
		TransferConfig:    transferConfig,
		AuthorizationCode: authorizationCode,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateTransferConfig(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceCreateTransferConfigError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var transferConfig *datatransferpb.TransferConfig = &datatransferpb.TransferConfig{}
	var authorizationCode string = "authorizationCode1571154419"
	var request = &datatransferpb.CreateTransferConfigRequest{
		Parent:            formattedParent,
		TransferConfig:    transferConfig,
		AuthorizationCode: authorizationCode,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateTransferConfig(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceUpdateTransferConfig(t *testing.T) {
	var name string = "name3373707"
	var destinationDatasetId string = "destinationDatasetId1541564179"
	var displayName string = "displayName1615086568"
	var dataSourceId string = "dataSourceId-1015796374"
	var schedule string = "schedule-697920873"
	var dataRefreshWindowDays int32 = 327632845
	var disabled bool = true
	var userId int64 = 147132913
	var datasetRegion string = "datasetRegion959248539"
	var expectedResponse = &datatransferpb.TransferConfig{
		Name:                  name,
		DestinationDatasetId:  destinationDatasetId,
		DisplayName:           displayName,
		DataSourceId:          dataSourceId,
		Schedule:              schedule,
		DataRefreshWindowDays: dataRefreshWindowDays,
		Disabled:              disabled,
		UserId:                userId,
		DatasetRegion:         datasetRegion,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var transferConfig *datatransferpb.TransferConfig = &datatransferpb.TransferConfig{}
	var authorizationCode string = "authorizationCode1571154419"
	var updateMask *field_maskpb.FieldMask = &field_maskpb.FieldMask{}
	var request = &datatransferpb.UpdateTransferConfigRequest{
		TransferConfig:    transferConfig,
		AuthorizationCode: authorizationCode,
		UpdateMask:        updateMask,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateTransferConfig(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceUpdateTransferConfigError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var transferConfig *datatransferpb.TransferConfig = &datatransferpb.TransferConfig{}
	var authorizationCode string = "authorizationCode1571154419"
	var updateMask *field_maskpb.FieldMask = &field_maskpb.FieldMask{}
	var request = &datatransferpb.UpdateTransferConfigRequest{
		TransferConfig:    transferConfig,
		AuthorizationCode: authorizationCode,
		UpdateMask:        updateMask,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateTransferConfig(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceDeleteTransferConfig(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var request = &datatransferpb.DeleteTransferConfigRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteTransferConfig(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestDataTransferServiceDeleteTransferConfigError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var request = &datatransferpb.DeleteTransferConfigRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteTransferConfig(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestDataTransferServiceGetTransferConfig(t *testing.T) {
	var name2 string = "name2-1052831874"
	var destinationDatasetId string = "destinationDatasetId1541564179"
	var displayName string = "displayName1615086568"
	var dataSourceId string = "dataSourceId-1015796374"
	var schedule string = "schedule-697920873"
	var dataRefreshWindowDays int32 = 327632845
	var disabled bool = true
	var userId int64 = 147132913
	var datasetRegion string = "datasetRegion959248539"
	var expectedResponse = &datatransferpb.TransferConfig{
		Name:                  name2,
		DestinationDatasetId:  destinationDatasetId,
		DisplayName:           displayName,
		DataSourceId:          dataSourceId,
		Schedule:              schedule,
		DataRefreshWindowDays: dataRefreshWindowDays,
		Disabled:              disabled,
		UserId:                userId,
		DatasetRegion:         datasetRegion,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var request = &datatransferpb.GetTransferConfigRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetTransferConfig(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceGetTransferConfigError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var request = &datatransferpb.GetTransferConfigRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetTransferConfig(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceListTransferConfigs(t *testing.T) {
	var nextPageToken string = ""
	var transferConfigsElement *datatransferpb.TransferConfig = &datatransferpb.TransferConfig{}
	var transferConfigs = []*datatransferpb.TransferConfig{transferConfigsElement}
	var expectedResponse = &datatransferpb.ListTransferConfigsResponse{
		NextPageToken:   nextPageToken,
		TransferConfigs: transferConfigs,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var dataSourceIds []string = nil
	var request = &datatransferpb.ListTransferConfigsRequest{
		Parent:        formattedParent,
		DataSourceIds: dataSourceIds,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListTransferConfigs(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.TransferConfigs[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceListTransferConfigsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var dataSourceIds []string = nil
	var request = &datatransferpb.ListTransferConfigsRequest{
		Parent:        formattedParent,
		DataSourceIds: dataSourceIds,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListTransferConfigs(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceScheduleTransferRuns(t *testing.T) {
	var expectedResponse *datatransferpb.ScheduleTransferRunsResponse = &datatransferpb.ScheduleTransferRunsResponse{}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var labels map[string]string = nil
	var startTime *timestamppb.Timestamp = &timestamppb.Timestamp{}
	var endTime *timestamppb.Timestamp = &timestamppb.Timestamp{}
	var request = &datatransferpb.ScheduleTransferRunsRequest{
		Parent:    formattedParent,
		Labels:    labels,
		StartTime: startTime,
		EndTime:   endTime,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ScheduleTransferRuns(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceScheduleTransferRunsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var labels map[string]string = nil
	var startTime *timestamppb.Timestamp = &timestamppb.Timestamp{}
	var endTime *timestamppb.Timestamp = &timestamppb.Timestamp{}
	var request = &datatransferpb.ScheduleTransferRunsRequest{
		Parent:    formattedParent,
		Labels:    labels,
		StartTime: startTime,
		EndTime:   endTime,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ScheduleTransferRuns(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceGetTransferRun(t *testing.T) {
	var name2 string = "name2-1052831874"
	var destinationDatasetId string = "destinationDatasetId1541564179"
	var dataSourceId string = "dataSourceId-1015796374"
	var userId int64 = 147132913
	var schedule string = "schedule-697920873"
	var expectedResponse = &datatransferpb.TransferRun{
		Name:                 name2,
		DestinationDatasetId: destinationDatasetId,
		DataSourceId:         dataSourceId,
		UserId:               userId,
		Schedule:             schedule,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var request = &datatransferpb.GetTransferRunRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetTransferRun(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceGetTransferRunError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var request = &datatransferpb.GetTransferRunRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetTransferRun(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceDeleteTransferRun(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var request = &datatransferpb.DeleteTransferRunRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteTransferRun(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestDataTransferServiceDeleteTransferRunError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var request = &datatransferpb.DeleteTransferRunRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteTransferRun(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestDataTransferServiceListTransferRuns(t *testing.T) {
	var nextPageToken string = ""
	var transferRunsElement *datatransferpb.TransferRun = &datatransferpb.TransferRun{}
	var transferRuns = []*datatransferpb.TransferRun{transferRunsElement}
	var expectedResponse = &datatransferpb.ListTransferRunsResponse{
		NextPageToken: nextPageToken,
		TransferRuns:  transferRuns,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var states []datatransferpb.TransferState = nil
	var runAttempt datatransferpb.ListTransferRunsRequest_RunAttempt = datatransferpb.ListTransferRunsRequest_RUN_ATTEMPT_UNSPECIFIED
	var request = &datatransferpb.ListTransferRunsRequest{
		Parent:     formattedParent,
		States:     states,
		RunAttempt: runAttempt,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListTransferRuns(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.TransferRuns[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceListTransferRunsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]")
	var states []datatransferpb.TransferState = nil
	var runAttempt datatransferpb.ListTransferRunsRequest_RunAttempt = datatransferpb.ListTransferRunsRequest_RUN_ATTEMPT_UNSPECIFIED
	var request = &datatransferpb.ListTransferRunsRequest{
		Parent:     formattedParent,
		States:     states,
		RunAttempt: runAttempt,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListTransferRuns(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceListTransferLogs(t *testing.T) {
	var nextPageToken string = ""
	var transferMessagesElement *datatransferpb.TransferMessage = &datatransferpb.TransferMessage{}
	var transferMessages = []*datatransferpb.TransferMessage{transferMessagesElement}
	var expectedResponse = &datatransferpb.ListTransferLogsResponse{
		NextPageToken:    nextPageToken,
		TransferMessages: transferMessages,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var messageTypes []datatransferpb.TransferMessage_MessageSeverity = nil
	var request = &datatransferpb.ListTransferLogsRequest{
		Parent:       formattedParent,
		MessageTypes: messageTypes,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListTransferLogs(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.TransferMessages[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceListTransferLogsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var messageTypes []datatransferpb.TransferMessage_MessageSeverity = nil
	var request = &datatransferpb.ListTransferLogsRequest{
		Parent:       formattedParent,
		MessageTypes: messageTypes,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListTransferLogs(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceCheckValidCreds(t *testing.T) {
	var hasValidCreds bool = false
	var expectedResponse = &datatransferpb.CheckValidCredsResponse{
		HasValidCreds: hasValidCreds,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSources/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE]")
	var request = &datatransferpb.CheckValidCredsRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CheckValidCreds(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceCheckValidCredsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSources/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE]")
	var request = &datatransferpb.CheckValidCredsRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CheckValidCreds(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataTransferServiceEnableDataTransferService(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.EnableDataTransferServiceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.EnableDataTransferService(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestDataTransferServiceEnableDataTransferServiceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.EnableDataTransferServiceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.EnableDataTransferService(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestDataTransferServiceIsDataTransferServiceEnabled(t *testing.T) {
	var enabled bool = false
	var reason string = "reason-934964668"
	var expectedResponse = &datatransferpb.IsDataTransferServiceEnabledResponse{
		Enabled: enabled,
		Reason:  reason,
	}

	mockDataTransfer.err = nil
	mockDataTransfer.reqs = nil

	mockDataTransfer.resps = append(mockDataTransfer.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.IsDataTransferServiceEnabledRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.IsDataTransferServiceEnabled(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataTransfer.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataTransferServiceIsDataTransferServiceEnabledError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataTransfer.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.IsDataTransferServiceEnabledRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.IsDataTransferServiceEnabled(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataSourceServiceUpdateTransferRun(t *testing.T) {
	var name string = "name3373707"
	var destinationDatasetId string = "destinationDatasetId1541564179"
	var dataSourceId string = "dataSourceId-1015796374"
	var userId int64 = 147132913
	var schedule string = "schedule-697920873"
	var expectedResponse = &datatransferpb.TransferRun{
		Name:                 name,
		DestinationDatasetId: destinationDatasetId,
		DataSourceId:         dataSourceId,
		UserId:               userId,
		Schedule:             schedule,
	}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var transferRun *datatransferpb.TransferRun = &datatransferpb.TransferRun{}
	var updateMask *field_maskpb.FieldMask = &field_maskpb.FieldMask{}
	var request = &datatransferpb.UpdateTransferRunRequest{
		TransferRun: transferRun,
		UpdateMask:  updateMask,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateTransferRun(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataSourceServiceUpdateTransferRunError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var transferRun *datatransferpb.TransferRun = &datatransferpb.TransferRun{}
	var updateMask *field_maskpb.FieldMask = &field_maskpb.FieldMask{}
	var request = &datatransferpb.UpdateTransferRunRequest{
		TransferRun: transferRun,
		UpdateMask:  updateMask,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateTransferRun(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataSourceServiceLogTransferRunMessages(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var transferMessages []*datatransferpb.TransferMessage = nil
	var request = &datatransferpb.LogTransferRunMessagesRequest{
		Name:             formattedName,
		TransferMessages: transferMessages,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.LogTransferRunMessages(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestDataSourceServiceLogTransferRunMessagesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var transferMessages []*datatransferpb.TransferMessage = nil
	var request = &datatransferpb.LogTransferRunMessagesRequest{
		Name:             formattedName,
		TransferMessages: transferMessages,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.LogTransferRunMessages(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestDataSourceServiceStartBigQueryJobs(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var importedData []*datatransferpb.ImportedDataInfo = nil
	var userCredentials []byte = []byte("-120")
	var request = &datatransferpb.StartBigQueryJobsRequest{
		Name:            formattedName,
		ImportedData:    importedData,
		UserCredentials: userCredentials,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.StartBigQueryJobs(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestDataSourceServiceStartBigQueryJobsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var importedData []*datatransferpb.ImportedDataInfo = nil
	var userCredentials []byte = []byte("-120")
	var request = &datatransferpb.StartBigQueryJobsRequest{
		Name:            formattedName,
		ImportedData:    importedData,
		UserCredentials: userCredentials,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.StartBigQueryJobs(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestDataSourceServiceGetCredentials(t *testing.T) {
	var name2 string = "name2-1052831874"
	var authToken string = "authToken-1956766558"
	var expectedResponse = &datatransferpb.Credentials{
		Name:      name2,
		AuthToken: authToken,
	}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSources/%s/credentials/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE]", "[CREDENTIAL]")
	var request = &datatransferpb.GetCredentialsRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetCredentials(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataSourceServiceGetCredentialsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSources/%s/credentials/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE]", "[CREDENTIAL]")
	var request = &datatransferpb.GetCredentialsRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetCredentials(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataSourceServiceFinishRun(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var request = &datatransferpb.FinishRunRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.FinishRun(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestDataSourceServiceFinishRunError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/transferConfigs/%s/runs/%s", "[PROJECT]", "[LOCATION]", "[TRANSFER_CONFIG]", "[RUN]")
	var request = &datatransferpb.FinishRunRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.FinishRun(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestDataSourceServiceCreateDataSourceDefinition(t *testing.T) {
	var name string = "name3373707"
	var transferRunPubsubTopic string = "transferRunPubsubTopic-1740562661"
	var supportEmail string = "supportEmail-648030420"
	var disabled bool = true
	var expectedResponse = &datatransferpb.DataSourceDefinition{
		Name: name,
		TransferRunPubsubTopic: transferRunPubsubTopic,
		SupportEmail:           supportEmail,
		Disabled:               disabled,
	}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var dataSourceDefinition *datatransferpb.DataSourceDefinition = &datatransferpb.DataSourceDefinition{}
	var request = &datatransferpb.CreateDataSourceDefinitionRequest{
		Parent:               formattedParent,
		DataSourceDefinition: dataSourceDefinition,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateDataSourceDefinition(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataSourceServiceCreateDataSourceDefinitionError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var dataSourceDefinition *datatransferpb.DataSourceDefinition = &datatransferpb.DataSourceDefinition{}
	var request = &datatransferpb.CreateDataSourceDefinitionRequest{
		Parent:               formattedParent,
		DataSourceDefinition: dataSourceDefinition,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateDataSourceDefinition(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataSourceServiceUpdateDataSourceDefinition(t *testing.T) {
	var name string = "name3373707"
	var transferRunPubsubTopic string = "transferRunPubsubTopic-1740562661"
	var supportEmail string = "supportEmail-648030420"
	var disabled bool = true
	var expectedResponse = &datatransferpb.DataSourceDefinition{
		Name: name,
		TransferRunPubsubTopic: transferRunPubsubTopic,
		SupportEmail:           supportEmail,
		Disabled:               disabled,
	}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var dataSourceDefinition *datatransferpb.DataSourceDefinition = &datatransferpb.DataSourceDefinition{}
	var updateMask *field_maskpb.FieldMask = &field_maskpb.FieldMask{}
	var request = &datatransferpb.UpdateDataSourceDefinitionRequest{
		DataSourceDefinition: dataSourceDefinition,
		UpdateMask:           updateMask,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateDataSourceDefinition(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataSourceServiceUpdateDataSourceDefinitionError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var dataSourceDefinition *datatransferpb.DataSourceDefinition = &datatransferpb.DataSourceDefinition{}
	var updateMask *field_maskpb.FieldMask = &field_maskpb.FieldMask{}
	var request = &datatransferpb.UpdateDataSourceDefinitionRequest{
		DataSourceDefinition: dataSourceDefinition,
		UpdateMask:           updateMask,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateDataSourceDefinition(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataSourceServiceDeleteDataSourceDefinition(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSourceDefinitions/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE_DEFINITION]")
	var request = &datatransferpb.DeleteDataSourceDefinitionRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteDataSourceDefinition(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestDataSourceServiceDeleteDataSourceDefinitionError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSourceDefinitions/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE_DEFINITION]")
	var request = &datatransferpb.DeleteDataSourceDefinitionRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteDataSourceDefinition(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestDataSourceServiceGetDataSourceDefinition(t *testing.T) {
	var name2 string = "name2-1052831874"
	var transferRunPubsubTopic string = "transferRunPubsubTopic-1740562661"
	var supportEmail string = "supportEmail-648030420"
	var disabled bool = true
	var expectedResponse = &datatransferpb.DataSourceDefinition{
		Name: name2,
		TransferRunPubsubTopic: transferRunPubsubTopic,
		SupportEmail:           supportEmail,
		Disabled:               disabled,
	}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSourceDefinitions/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE_DEFINITION]")
	var request = &datatransferpb.GetDataSourceDefinitionRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetDataSourceDefinition(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataSourceServiceGetDataSourceDefinitionError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/dataSourceDefinitions/%s", "[PROJECT]", "[LOCATION]", "[DATA_SOURCE_DEFINITION]")
	var request = &datatransferpb.GetDataSourceDefinitionRequest{
		Name: formattedName,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetDataSourceDefinition(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestDataSourceServiceListDataSourceDefinitions(t *testing.T) {
	var nextPageToken string = ""
	var dataSourceDefinitionsElement *datatransferpb.DataSourceDefinition = &datatransferpb.DataSourceDefinition{}
	var dataSourceDefinitions = []*datatransferpb.DataSourceDefinition{dataSourceDefinitionsElement}
	var expectedResponse = &datatransferpb.ListDataSourceDefinitionsResponse{
		NextPageToken:         nextPageToken,
		DataSourceDefinitions: dataSourceDefinitions,
	}

	mockDataSource.err = nil
	mockDataSource.reqs = nil

	mockDataSource.resps = append(mockDataSource.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.ListDataSourceDefinitionsRequest{
		Parent: formattedParent,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListDataSourceDefinitions(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockDataSource.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.DataSourceDefinitions[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestDataSourceServiceListDataSourceDefinitionsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockDataSource.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s", "[PROJECT]", "[LOCATION]")
	var request = &datatransferpb.ListDataSourceDefinitionsRequest{
		Parent: formattedParent,
	}

	c, err := NewDataSourceClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListDataSourceDefinitions(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
