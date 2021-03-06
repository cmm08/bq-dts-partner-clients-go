// Copyright 2019 Google LLC
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

// Code generated by gapic-generator. DO NOT EDIT.

package datatransfer_test

import (
	"context"

	datatransfer "cloud.google.com/go/bigquery/datatransfer/apiv1"
	"google.golang.org/api/iterator"
	datatransferpb "google.golang.org/genproto/googleapis/cloud/bigquery/datatransfer/v1"
)

func ExampleNewDataSourceClient() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleDataSourceClient_UpdateTransferRun() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.UpdateTransferRunRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateTransferRun(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataSourceClient_LogTransferRunMessages() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.LogTransferRunMessagesRequest{
		// TODO: Fill request struct fields.
	}
	err = c.LogTransferRunMessages(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDataSourceClient_StartBigQueryJobs() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.StartBigQueryJobsRequest{
		// TODO: Fill request struct fields.
	}
	err = c.StartBigQueryJobs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDataSourceClient_FinishRun() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.FinishRunRequest{
		// TODO: Fill request struct fields.
	}
	err = c.FinishRun(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDataSourceClient_CreateDataSourceDefinition() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.CreateDataSourceDefinitionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateDataSourceDefinition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataSourceClient_UpdateDataSourceDefinition() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.UpdateDataSourceDefinitionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateDataSourceDefinition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataSourceClient_DeleteDataSourceDefinition() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.DeleteDataSourceDefinitionRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteDataSourceDefinition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDataSourceClient_GetDataSourceDefinition() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.GetDataSourceDefinitionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDataSourceDefinition(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDataSourceClient_ListDataSourceDefinitions() {
	ctx := context.Background()
	c, err := datatransfer.NewDataSourceClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datatransferpb.ListDataSourceDefinitionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListDataSourceDefinitions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
