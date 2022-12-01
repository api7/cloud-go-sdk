// Copyright 2022 API7.ai, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloud

import (
	"context"
	"encoding/json"
	"path"
)

// Consumer is an abstraction of Application/API caller.
type Consumer struct {
	// ID is the unique identify to mark an object.
	ID ID `json:"id"`
	// Name of the consumer, should be unique among all applications in the same control plane.
	Name string `json:"name" gorm:"column:name"`
	// Description for this consumer.
	Description string                 `json:"description"` // Certificates are used to authenticate the consumer.
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	// Plugins settings on Consumer level
	Plugins Plugins `json:"plugins,omitempty"`
	// Labels are used for resource classification and indexing
	Labels []string `json:"labels,omitempty"`
}

// ConsumerInterface is the interface for manipulating Consumers.
type ConsumerInterface interface {
	// CreateConsumer creates an API7 Cloud Consumer in the specified control plane.
	// The given `consumer` parameter should specify the desired Consumer specification.
	// Users need to specify the ControlPlane in the `opts`.
	// The returned Consumer will contain the same Consumer specification plus some
	// management fields and default values.
	CreateConsumer(ctx context.Context, consumer *Consumer, opts *ResourceCreateOptions) (*Consumer, error)
	// UpdateConsumer updates an existing API7 Cloud Consumer in the specified control plane.
	// The given `consumer` parameter should specify the desired Consumer specification.
	// Users need to specify the ControlPlane in the `opts`.
	// The returned Consumer will contain the same Consumer specification plus some
	// management fields and default values.
	UpdateConsumer(ctx context.Context, consumer *Consumer, opts *ResourceUpdateOptions) (*Consumer, error)
	// DeleteConsumer deletes an existing API7 Cloud Consumer in the specified control plane.
	// The given `consumerID` parameter should specify the Consumer that you want to delete.
	// Users need to specify the ControlPlane in the `opts`.
	DeleteConsumer(ctx context.Context, consumerID ID, opts *ResourceDeleteOptions) error
	// GetConsumer gets an existing API7 Cloud Consumer in the specified control plane.
	// The given `consumerID` parameter should specify the Consumer that you want to get.
	// Users need to specify the ControlPlane in the `opts`.
	GetConsumer(ctx context.Context, consumerID ID, opts *ResourceGetOptions) (*Consumer, error)
	// ListConsumers returns an iterator for listing Consumers in the specified control plane with the
	// given list conditions.
	// Users need to specify the ControlPlane, Paging conditions in the `opts`.
	ListConsumers(ctx context.Context, opts *ResourceListOptions) (ConsumerListIterator, error)
}

// ConsumerListIterator is an iterator for listing Consumers.
type ConsumerListIterator interface {
	// Next returns the next Consumer according to the filter conditions.
	Next() (*Consumer, error)
}

type consumerImpl struct {
	client httpClient
}

type consumerListIterator struct {
	iter listIterator
}

func (iter *consumerListIterator) Next() (*Consumer, error) {
	var consumer Consumer
	rawData, err := iter.iter.Next()
	if err != nil {
		return nil, err
	}
	if rawData == nil {
		return nil, nil
	}
	if err = json.Unmarshal(rawData, &consumer); err != nil {
		return nil, err
	}
	return &consumer, nil
}

func newConsumer(cli httpClient) ConsumerInterface {
	return &consumerImpl{
		client: cli,
	}
}

func (impl *consumerImpl) CreateConsumer(ctx context.Context, consumer *Consumer, opts *ResourceCreateOptions) (*Consumer, error) {
	var createdConsumer Consumer

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "consumers")
	err := impl.client.sendPostRequest(ctx, uri, "", consumer, jsonPayloadDecodeFactory(&createdConsumer))
	if err != nil {
		return nil, err
	}
	return &createdConsumer, nil
}

func (impl *consumerImpl) UpdateConsumer(ctx context.Context, consumer *Consumer, opts *ResourceUpdateOptions) (*Consumer, error) {
	var updatedConsumer Consumer

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "consumers", consumer.ID.String())
	err := impl.client.sendPutRequest(ctx, uri, "", consumer, jsonPayloadDecodeFactory(&updatedConsumer))
	if err != nil {
		return nil, err
	}
	return &updatedConsumer, nil
}

func (impl *consumerImpl) DeleteConsumer(ctx context.Context, consumerID ID, opts *ResourceDeleteOptions) error {
	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "consumers", consumerID.String())
	return impl.client.sendDeleteRequest(ctx, uri, "", nil)
}

func (impl *consumerImpl) GetConsumer(ctx context.Context, consumerID ID, opts *ResourceGetOptions) (*Consumer, error) {
	var consumer Consumer

	cpID := opts.ControlPlane.ID
	uri := path.Join(_apiPathPrefix, "controlplanes", cpID.String(), "consumers", consumerID.String())
	err := impl.client.sendGetRequest(ctx, uri, "", jsonPayloadDecodeFactory(&consumer))
	if err != nil {
		return nil, err
	}
	return &consumer, nil
}

func (impl *consumerImpl) ListConsumers(ctx context.Context, opts *ResourceListOptions) (ConsumerListIterator, error) {
	iter := listIterator{
		ctx:      ctx,
		resource: "consumer",
		client:   impl.client,
		path:     path.Join(_apiPathPrefix, "controlplanes", opts.ControlPlane.ID.String(), "consumers"),
		paging:   mergePagination(opts.Pagination),
	}

	return &consumerListIterator{iter: iter}, nil
}
