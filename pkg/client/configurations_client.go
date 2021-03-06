// Copyright 2021 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package client

import (
	"context"
	"math"
	"time"

	pb "github.com/cristian-radu/cloud-run-grpc/pkg/pb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newConfigurationsClientHook clientHook

// ConfigurationsCallOptions contains the retry settings for each method of ConfigurationsClient.
type ConfigurationsCallOptions struct {
	GetConfiguration   []gax.CallOption
	ListConfigurations []gax.CallOption
}

func defaultConfigurationsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("run.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("run.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://run.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConfigurationsCallOptions() *ConfigurationsCallOptions {
	return &ConfigurationsCallOptions{
		GetConfiguration:   []gax.CallOption{},
		ListConfigurations: []gax.CallOption{},
	}
}

// internalConfigurationsClient is an interface that defines the methods availaible from Cloud Run Admin API.
type internalConfigurationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetConfiguration(context.Context, *pb.GetConfigurationRequest, ...gax.CallOption) (*pb.Configuration, error)
	ListConfigurations(context.Context, *pb.ListConfigurationsRequest, ...gax.CallOption) (*pb.ListConfigurationsResponse, error)
}

// ConfigurationsClient is a client for interacting with Cloud Run Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ConfigurationsClient struct {
	// The internal transport-dependent client.
	internalClient internalConfigurationsClient

	// The call options for this service.
	CallOptions *ConfigurationsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConfigurationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConfigurationsClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ConfigurationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetConfiguration get information about a configuration.
func (c *ConfigurationsClient) GetConfiguration(ctx context.Context, req *pb.GetConfigurationRequest, opts ...gax.CallOption) (*pb.Configuration, error) {
	return c.internalClient.GetConfiguration(ctx, req, opts...)
}

// ListConfigurations list configurations.
func (c *ConfigurationsClient) ListConfigurations(ctx context.Context, req *pb.ListConfigurationsRequest, opts ...gax.CallOption) (*pb.ListConfigurationsResponse, error) {
	return c.internalClient.ListConfigurations(ctx, req, opts...)
}

// configurationsGRPCClient is a client for interacting with Cloud Run Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type configurationsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ConfigurationsClient
	CallOptions **ConfigurationsCallOptions

	// The gRPC API client.
	configurationsClient pb.ConfigurationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewConfigurationsClient creates a new configurations client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewConfigurationsClient(ctx context.Context, opts ...option.ClientOption) (*ConfigurationsClient, error) {
	clientOpts := defaultConfigurationsGRPCClientOptions()
	if newConfigurationsClientHook != nil {
		hookOpts, err := newConfigurationsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ConfigurationsClient{CallOptions: defaultConfigurationsCallOptions()}

	c := &configurationsGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		configurationsClient: pb.NewConfigurationsClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *configurationsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *configurationsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *configurationsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *configurationsGRPCClient) GetConfiguration(ctx context.Context, req *pb.GetConfigurationRequest, opts ...gax.CallOption) (*pb.Configuration, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetConfiguration[0:len((*c.CallOptions).GetConfiguration):len((*c.CallOptions).GetConfiguration)], opts...)
	var resp *pb.Configuration
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.configurationsClient.GetConfiguration(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *configurationsGRPCClient) ListConfigurations(ctx context.Context, req *pb.ListConfigurationsRequest, opts ...gax.CallOption) (*pb.ListConfigurationsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListConfigurations[0:len((*c.CallOptions).ListConfigurations):len((*c.CallOptions).ListConfigurations)], opts...)
	var resp *pb.ListConfigurationsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.configurationsClient.ListConfigurations(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
