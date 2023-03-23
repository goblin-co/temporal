// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by cmd/tools/rpcwrappers. DO NOT EDIT.

package matching

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"google.golang.org/grpc"

	"go.temporal.io/server/common/backoff"
)

func (c *retryableClient) AddActivityTask(
	ctx context.Context,
	request *matchingservice.AddActivityTaskRequest,
	opts ...grpc.CallOption,
) (*matchingservice.AddActivityTaskResponse, error) {
	var resp *matchingservice.AddActivityTaskResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.AddActivityTask(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) AddWorkflowTask(
	ctx context.Context,
	request *matchingservice.AddWorkflowTaskRequest,
	opts ...grpc.CallOption,
) (*matchingservice.AddWorkflowTaskResponse, error) {
	var resp *matchingservice.AddWorkflowTaskResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.AddWorkflowTask(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) CancelOutstandingPoll(
	ctx context.Context,
	request *matchingservice.CancelOutstandingPollRequest,
	opts ...grpc.CallOption,
) (*matchingservice.CancelOutstandingPollResponse, error) {
	var resp *matchingservice.CancelOutstandingPollResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.CancelOutstandingPoll(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) DescribeTaskQueue(
	ctx context.Context,
	request *matchingservice.DescribeTaskQueueRequest,
	opts ...grpc.CallOption,
) (*matchingservice.DescribeTaskQueueResponse, error) {
	var resp *matchingservice.DescribeTaskQueueResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.DescribeTaskQueue(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) GetTaskQueueMetadata(
	ctx context.Context,
	request *matchingservice.GetTaskQueueMetadataRequest,
	opts ...grpc.CallOption,
) (*matchingservice.GetTaskQueueMetadataResponse, error) {
	var resp *matchingservice.GetTaskQueueMetadataResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.GetTaskQueueMetadata(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) GetWorkerBuildIdCompatibility(
	ctx context.Context,
	request *matchingservice.GetWorkerBuildIdCompatibilityRequest,
	opts ...grpc.CallOption,
) (*matchingservice.GetWorkerBuildIdCompatibilityResponse, error) {
	var resp *matchingservice.GetWorkerBuildIdCompatibilityResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.GetWorkerBuildIdCompatibility(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) InvalidateTaskQueueMetadata(
	ctx context.Context,
	request *matchingservice.InvalidateTaskQueueMetadataRequest,
	opts ...grpc.CallOption,
) (*matchingservice.InvalidateTaskQueueMetadataResponse, error) {
	var resp *matchingservice.InvalidateTaskQueueMetadataResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.InvalidateTaskQueueMetadata(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) ListTaskQueuePartitions(
	ctx context.Context,
	request *matchingservice.ListTaskQueuePartitionsRequest,
	opts ...grpc.CallOption,
) (*matchingservice.ListTaskQueuePartitionsResponse, error) {
	var resp *matchingservice.ListTaskQueuePartitionsResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.ListTaskQueuePartitions(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) PollActivityTaskQueue(
	ctx context.Context,
	request *matchingservice.PollActivityTaskQueueRequest,
	opts ...grpc.CallOption,
) (*matchingservice.PollActivityTaskQueueResponse, error) {
	var resp *matchingservice.PollActivityTaskQueueResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.PollActivityTaskQueue(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) PollWorkflowTaskQueue(
	ctx context.Context,
	request *matchingservice.PollWorkflowTaskQueueRequest,
	opts ...grpc.CallOption,
) (*matchingservice.PollWorkflowTaskQueueResponse, error) {
	var resp *matchingservice.PollWorkflowTaskQueueResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.PollWorkflowTaskQueue(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) QueryWorkflow(
	ctx context.Context,
	request *matchingservice.QueryWorkflowRequest,
	opts ...grpc.CallOption,
) (*matchingservice.QueryWorkflowResponse, error) {
	var resp *matchingservice.QueryWorkflowResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.QueryWorkflow(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) RespondQueryTaskCompleted(
	ctx context.Context,
	request *matchingservice.RespondQueryTaskCompletedRequest,
	opts ...grpc.CallOption,
) (*matchingservice.RespondQueryTaskCompletedResponse, error) {
	var resp *matchingservice.RespondQueryTaskCompletedResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.RespondQueryTaskCompleted(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}

func (c *retryableClient) UpdateWorkerBuildIdCompatibility(
	ctx context.Context,
	request *matchingservice.UpdateWorkerBuildIdCompatibilityRequest,
	opts ...grpc.CallOption,
) (*matchingservice.UpdateWorkerBuildIdCompatibilityResponse, error) {
	var resp *matchingservice.UpdateWorkerBuildIdCompatibilityResponse
	op := func(ctx context.Context) error {
		var err error
		resp, err = c.client.UpdateWorkerBuildIdCompatibility(ctx, request, opts...)
		return err
	}
	err := backoff.ThrottleRetryContext(ctx, op, c.policy, c.isRetryable)
	return resp, err
}
