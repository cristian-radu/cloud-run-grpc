// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: google/cloud/run/v1/routes.proto

package run

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the route to retrieve. For Cloud Run (fully managed), replace {namespace_id} with the project ID or number.
	// Authorization requires the following IAM permission on the specified resource name:
	// run.routes.get
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRouteRequest) Reset() {
	*x = GetRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_routes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteRequest) ProtoMessage() {}

func (x *GetRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_routes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteRequest.ProtoReflect.Descriptor instead.
func (*GetRouteRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_routes_proto_rawDescGZIP(), []int{0}
}

func (x *GetRouteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The kind of this resource, in this case always "Route".
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Metadata associated with this Route, including name, namespace, labels, and annotations.
	Metadata *v1.ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec holds the desired state of the Route (from the client).
	Spec *RouteSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status communicates the observed state of the Route (from the controller).
	Status *RouteStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_routes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_routes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_routes_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Route) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Route) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Route) GetSpec() *RouteSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Route) GetStatus() *RouteStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type RouteSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations.
	// Cloud Run currently supports a single configurationName.
	Target []*TrafficTarget `protobuf:"bytes,1,rep,name=target,proto3" json:"target,omitempty"`
}

func (x *RouteSpec) Reset() {
	*x = RouteSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_routes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteSpec) ProtoMessage() {}

func (x *RouteSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_routes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteSpec.ProtoReflect.Descriptor instead.
func (*RouteSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_routes_proto_rawDescGZIP(), []int{2}
}

func (x *RouteSpec) GetTarget() []*TrafficTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

type RouteStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ObservedGeneration is the 'Generation' of the Route that was last processed by the controller.
	// Clients polling for completed reconciliation should poll until observedGeneration = metadata.generation and the Ready condition's status is True or False.
	// Note that providing a trafficTarget that only has a configurationName will result in a Route that does not increment either its metadata.generation
	// or its observedGeneration, as new "latest ready" revisions from the Configuration are processed without an update to the Route's spec.
	ObservedGeneration int32 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Conditions communicates information about ongoing/complete reconciliation processes that bring the "spec" inline with the observed state of the world.
	Conditions []*Condition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Traffic holds the configured traffic distribution. These entries will always contain RevisionName references.
	// When ConfigurationName appears in the spec, this will hold the LatestReadyRevisionName that we last observed.
	Traffic []*TrafficTarget `protobuf:"bytes,3,rep,name=traffic,proto3" json:"traffic,omitempty"`
	// URL holds the url that will distribute traffic over the provided traffic targets.
	// It generally has the form: https://{route-hash}-{project-hash}-{cluster-level-suffix}.a.run.app
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	// Similar to url, information on where the service is available on HTTP.
	Address *Addressable `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RouteStatus) Reset() {
	*x = RouteStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_routes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteStatus) ProtoMessage() {}

func (x *RouteStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_routes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteStatus.ProtoReflect.Descriptor instead.
func (*RouteStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_routes_proto_rawDescGZIP(), []int{3}
}

func (x *RouteStatus) GetObservedGeneration() int32 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *RouteStatus) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *RouteStatus) GetTraffic() []*TrafficTarget {
	if x != nil {
		return x.Traffic
	}
	return nil
}

func (x *RouteStatus) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RouteStatus) GetAddress() *Addressable {
	if x != nil {
		return x.Address
	}
	return nil
}

type ListRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace from which the routes should be listed. For Cloud Run (fully managed), replace {namespace_id} with the project ID or number.
	// Authorization requires the following IAM permission on the specified resource parent:
	// run.routes.list
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of records that should be returned.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Not currently used by Cloud Run.
	IncludeUninitialized bool `protobuf:"varint,3,opt,name=include_uninitialized,json=includeUninitialized,proto3" json:"include_uninitialized,omitempty"`
	// Allows to filter resources based on a specific value for a field name.
	// Send this in a query string format. i.e. 'metadata.name%3Dlorem'. Not currently used by Cloud Run.
	FieldSelector string `protobuf:"bytes,4,opt,name=field_selector,json=fieldSelector,proto3" json:"field_selector,omitempty"`
	// Allows to filter resources based on a label. Supported operations are =, !=, exists, in, and notIn.
	LabelSelector string `protobuf:"bytes,5,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	// The baseline resource version from which the list or watch operation should start. Not currently used by Cloud Run.
	ResourceVersion string `protobuf:"bytes,6,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	// Flag that indicates that the client expects to watch this resource as well. Not currently used by Cloud Run.
	Watch bool `protobuf:"varint,7,opt,name=watch,proto3" json:"watch,omitempty"`
	// Optional encoded string to continue paging.
	PageToken string `protobuf:"bytes,8,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListRoutesRequest) Reset() {
	*x = ListRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_routes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesRequest) ProtoMessage() {}

func (x *ListRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_routes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesRequest.ProtoReflect.Descriptor instead.
func (*ListRoutesRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_routes_proto_rawDescGZIP(), []int{4}
}

func (x *ListRoutesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListRoutesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRoutesRequest) GetIncludeUninitialized() bool {
	if x != nil {
		return x.IncludeUninitialized
	}
	return false
}

func (x *ListRoutesRequest) GetFieldSelector() string {
	if x != nil {
		return x.FieldSelector
	}
	return ""
}

func (x *ListRoutesRequest) GetLabelSelector() string {
	if x != nil {
		return x.LabelSelector
	}
	return ""
}

func (x *ListRoutesRequest) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *ListRoutesRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

func (x *ListRoutesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The kind of this resource, in this case always "RouteList".
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	//  Metadata associated with this Route list.
	Metadata *v1.ListMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// List of Routes.
	Items []*Route `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,5,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListRoutesResponse) Reset() {
	*x = ListRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_routes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesResponse) ProtoMessage() {}

func (x *ListRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_routes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesResponse.ProtoReflect.Descriptor instead.
func (*ListRoutesResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_routes_proto_rawDescGZIP(), []int{5}
}

func (x *ListRoutesResponse) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ListRoutesResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ListRoutesResponse) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ListRoutesResponse) GetItems() []*Route {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListRoutesResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

var File_google_cloud_run_v1_routes_proto protoreflect.FileDescriptor

var file_google_cloud_run_v1_routes_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xf8, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4c,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x47, 0x0a, 0x09, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3a, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0x8a, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0xab, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x55, 0x6e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xe9,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65,
	0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x32, 0x81, 0x02, 0x0a, 0x06, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x1a, 0x46, 0xca, 0x41, 0x12, 0x72, 0x75, 0x6e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v1_routes_proto_rawDescOnce sync.Once
	file_google_cloud_run_v1_routes_proto_rawDescData = file_google_cloud_run_v1_routes_proto_rawDesc
)

func file_google_cloud_run_v1_routes_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v1_routes_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v1_routes_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v1_routes_proto_rawDescData)
	})
	return file_google_cloud_run_v1_routes_proto_rawDescData
}

var file_google_cloud_run_v1_routes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_run_v1_routes_proto_goTypes = []interface{}{
	(*GetRouteRequest)(nil),    // 0: google.cloud.run.v1.GetRouteRequest
	(*Route)(nil),              // 1: google.cloud.run.v1.Route
	(*RouteSpec)(nil),          // 2: google.cloud.run.v1.RouteSpec
	(*RouteStatus)(nil),        // 3: google.cloud.run.v1.RouteStatus
	(*ListRoutesRequest)(nil),  // 4: google.cloud.run.v1.ListRoutesRequest
	(*ListRoutesResponse)(nil), // 5: google.cloud.run.v1.ListRoutesResponse
	(*v1.ObjectMeta)(nil),      // 6: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*TrafficTarget)(nil),      // 7: google.cloud.run.v1.TrafficTarget
	(*Condition)(nil),          // 8: google.cloud.run.v1.Condition
	(*Addressable)(nil),        // 9: google.cloud.run.v1.Addressable
	(*v1.ListMeta)(nil),        // 10: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
}
var file_google_cloud_run_v1_routes_proto_depIdxs = []int32{
	6,  // 0: google.cloud.run.v1.Route.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	2,  // 1: google.cloud.run.v1.Route.spec:type_name -> google.cloud.run.v1.RouteSpec
	3,  // 2: google.cloud.run.v1.Route.status:type_name -> google.cloud.run.v1.RouteStatus
	7,  // 3: google.cloud.run.v1.RouteSpec.target:type_name -> google.cloud.run.v1.TrafficTarget
	8,  // 4: google.cloud.run.v1.RouteStatus.conditions:type_name -> google.cloud.run.v1.Condition
	7,  // 5: google.cloud.run.v1.RouteStatus.traffic:type_name -> google.cloud.run.v1.TrafficTarget
	9,  // 6: google.cloud.run.v1.RouteStatus.address:type_name -> google.cloud.run.v1.Addressable
	10, // 7: google.cloud.run.v1.ListRoutesResponse.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	1,  // 8: google.cloud.run.v1.ListRoutesResponse.items:type_name -> google.cloud.run.v1.Route
	0,  // 9: google.cloud.run.v1.Routes.GetRoute:input_type -> google.cloud.run.v1.GetRouteRequest
	4,  // 10: google.cloud.run.v1.Routes.ListRoutes:input_type -> google.cloud.run.v1.ListRoutesRequest
	1,  // 11: google.cloud.run.v1.Routes.GetRoute:output_type -> google.cloud.run.v1.Route
	5,  // 12: google.cloud.run.v1.Routes.ListRoutes:output_type -> google.cloud.run.v1.ListRoutesResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v1_routes_proto_init() }
func file_google_cloud_run_v1_routes_proto_init() {
	if File_google_cloud_run_v1_routes_proto != nil {
		return
	}
	file_google_cloud_run_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v1_routes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_run_v1_routes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_run_v1_routes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_run_v1_routes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_run_v1_routes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_run_v1_routes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_run_v1_routes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_run_v1_routes_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v1_routes_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v1_routes_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v1_routes_proto = out.File
	file_google_cloud_run_v1_routes_proto_rawDesc = nil
	file_google_cloud_run_v1_routes_proto_goTypes = nil
	file_google_cloud_run_v1_routes_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoutesClient is the client API for Routes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutesClient interface {
	// Get information about a route.
	GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*Route, error)
	// List routes.
	ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error)
}

type routesClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutesClient(cc grpc.ClientConnInterface) RoutesClient {
	return &routesClient{cc}
}

func (c *routesClient) GetRoute(ctx context.Context, in *GetRouteRequest, opts ...grpc.CallOption) (*Route, error) {
	out := new(Route)
	err := c.cc.Invoke(ctx, "/google.cloud.run.v1.Routes/GetRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routesClient) ListRoutes(ctx context.Context, in *ListRoutesRequest, opts ...grpc.CallOption) (*ListRoutesResponse, error) {
	out := new(ListRoutesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.run.v1.Routes/ListRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutesServer is the server API for Routes service.
type RoutesServer interface {
	// Get information about a route.
	GetRoute(context.Context, *GetRouteRequest) (*Route, error)
	// List routes.
	ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error)
}

// UnimplementedRoutesServer can be embedded to have forward compatible implementations.
type UnimplementedRoutesServer struct {
}

func (*UnimplementedRoutesServer) GetRoute(context.Context, *GetRouteRequest) (*Route, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoute not implemented")
}
func (*UnimplementedRoutesServer) ListRoutes(context.Context, *ListRoutesRequest) (*ListRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoutes not implemented")
}

func RegisterRoutesServer(s *grpc.Server, srv RoutesServer) {
	s.RegisterService(&_Routes_serviceDesc, srv)
}

func _Routes_GetRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesServer).GetRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.run.v1.Routes/GetRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesServer).GetRoute(ctx, req.(*GetRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Routes_ListRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesServer).ListRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.run.v1.Routes/ListRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesServer).ListRoutes(ctx, req.(*ListRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Routes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.run.v1.Routes",
	HandlerType: (*RoutesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoute",
			Handler:    _Routes_GetRoute_Handler,
		},
		{
			MethodName: "ListRoutes",
			Handler:    _Routes_ListRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/run/v1/routes.proto",
}