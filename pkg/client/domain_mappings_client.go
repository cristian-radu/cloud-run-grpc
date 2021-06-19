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

var newDomainMappingsClientHook clientHook

// DomainMappingsCallOptions contains the retry settings for each method of DomainMappingsClient.
type DomainMappingsCallOptions struct {
	CreateDomainMapping []gax.CallOption
	DeleteDomainMapping []gax.CallOption
	GetDomainMapping    []gax.CallOption
	ListDomainMappings  []gax.CallOption
}

func defaultDomainMappingsGRPCClientOptions() []option.ClientOption {
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

func defaultDomainMappingsCallOptions() *DomainMappingsCallOptions {
	return &DomainMappingsCallOptions{
		CreateDomainMapping: []gax.CallOption{},
		DeleteDomainMapping: []gax.CallOption{},
		GetDomainMapping:    []gax.CallOption{},
		ListDomainMappings:  []gax.CallOption{},
	}
}

// internalDomainMappingsClient is an interface that defines the methods availaible from Cloud Run Admin API.
type internalDomainMappingsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateDomainMapping(context.Context, *pb.CreateDomainMappingRequest, ...gax.CallOption) (*pb.DomainMapping, error)
	DeleteDomainMapping(context.Context, *pb.DeleteDomainMappingRequest, ...gax.CallOption) (*pb.Status, error)
	GetDomainMapping(context.Context, *pb.GetDomainMappingRequest, ...gax.CallOption) (*pb.DomainMapping, error)
	ListDomainMappings(context.Context, *pb.ListDomainMappingsRequest, ...gax.CallOption) (*pb.ListDomainMappingsResponse, error)
}

// DomainMappingsClient is a client for interacting with Cloud Run Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type DomainMappingsClient struct {
	// The internal transport-dependent client.
	internalClient internalDomainMappingsClient

	// The call options for this service.
	CallOptions *DomainMappingsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DomainMappingsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DomainMappingsClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *DomainMappingsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateDomainMapping create a new domain mapping.
func (c *DomainMappingsClient) CreateDomainMapping(ctx context.Context, req *pb.CreateDomainMappingRequest, opts ...gax.CallOption) (*pb.DomainMapping, error) {
	return c.internalClient.CreateDomainMapping(ctx, req, opts...)
}

// DeleteDomainMapping delete a domain mapping.
func (c *DomainMappingsClient) DeleteDomainMapping(ctx context.Context, req *pb.DeleteDomainMappingRequest, opts ...gax.CallOption) (*pb.Status, error) {
	return c.internalClient.DeleteDomainMapping(ctx, req, opts...)
}

// GetDomainMapping get information about a domain mapping.
func (c *DomainMappingsClient) GetDomainMapping(ctx context.Context, req *pb.GetDomainMappingRequest, opts ...gax.CallOption) (*pb.DomainMapping, error) {
	return c.internalClient.GetDomainMapping(ctx, req, opts...)
}

// ListDomainMappings list domain mappings.
func (c *DomainMappingsClient) ListDomainMappings(ctx context.Context, req *pb.ListDomainMappingsRequest, opts ...gax.CallOption) (*pb.ListDomainMappingsResponse, error) {
	return c.internalClient.ListDomainMappings(ctx, req, opts...)
}

// domainMappingsGRPCClient is a client for interacting with Cloud Run Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type domainMappingsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing DomainMappingsClient
	CallOptions **DomainMappingsCallOptions

	// The gRPC API client.
	domainMappingsClient pb.DomainMappingsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDomainMappingsClient creates a new domain mappings client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewDomainMappingsClient(ctx context.Context, opts ...option.ClientOption) (*DomainMappingsClient, error) {
	clientOpts := defaultDomainMappingsGRPCClientOptions()
	if newDomainMappingsClientHook != nil {
		hookOpts, err := newDomainMappingsClientHook(ctx, clientHookParams{})
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
	client := DomainMappingsClient{CallOptions: defaultDomainMappingsCallOptions()}

	c := &domainMappingsGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		domainMappingsClient: pb.NewDomainMappingsClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *domainMappingsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *domainMappingsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *domainMappingsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *domainMappingsGRPCClient) CreateDomainMapping(ctx context.Context, req *pb.CreateDomainMappingRequest, opts ...gax.CallOption) (*pb.DomainMapping, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateDomainMapping[0:len((*c.CallOptions).CreateDomainMapping):len((*c.CallOptions).CreateDomainMapping)], opts...)
	var resp *pb.DomainMapping
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.domainMappingsClient.CreateDomainMapping(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *domainMappingsGRPCClient) DeleteDomainMapping(ctx context.Context, req *pb.DeleteDomainMappingRequest, opts ...gax.CallOption) (*pb.Status, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).DeleteDomainMapping[0:len((*c.CallOptions).DeleteDomainMapping):len((*c.CallOptions).DeleteDomainMapping)], opts...)
	var resp *pb.Status
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.domainMappingsClient.DeleteDomainMapping(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *domainMappingsGRPCClient) GetDomainMapping(ctx context.Context, req *pb.GetDomainMappingRequest, opts ...gax.CallOption) (*pb.DomainMapping, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetDomainMapping[0:len((*c.CallOptions).GetDomainMapping):len((*c.CallOptions).GetDomainMapping)], opts...)
	var resp *pb.DomainMapping
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.domainMappingsClient.GetDomainMapping(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *domainMappingsGRPCClient) ListDomainMappings(ctx context.Context, req *pb.ListDomainMappingsRequest, opts ...gax.CallOption) (*pb.ListDomainMappingsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListDomainMappings[0:len((*c.CallOptions).ListDomainMappings):len((*c.CallOptions).ListDomainMappings)], opts...)
	var resp *pb.ListDomainMappingsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.domainMappingsClient.ListDomainMappings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
