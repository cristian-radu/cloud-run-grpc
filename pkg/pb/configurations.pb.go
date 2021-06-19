// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: google/cloud/run/v1/configurations.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the configuration to retrieve. For Cloud Run (fully managed), replace {namespace_id} with the project ID or number.
	// Authorization requires the following IAM permission on the specified resource name:
	// run.configurations.get
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetConfigurationRequest) Reset() {
	*x = GetConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationRequest) ProtoMessage() {}

func (x *GetConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationRequest.ProtoReflect.Descriptor instead.
func (*GetConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_configurations_proto_rawDescGZIP(), []int{0}
}

func (x *GetConfigurationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The kind of resource, in this case always "Configuration"
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Metadata associated with this Configuration, including name, namespace, labels, and annotations.
	Metadata *ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec holds the desired state of the Configuration (from the client).
	Spec *ConfigurationSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status communicates the observed state of the Configuration (from the controller).
	Status *ConfigurationStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_configurations_proto_rawDescGZIP(), []int{1}
}

func (x *Configuration) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Configuration) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Configuration) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Configuration) GetSpec() *ConfigurationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Configuration) GetStatus() *ConfigurationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ConfigurationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Template holds the latest specification for the Revision to be stamped out.
	Template *RevisionTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *ConfigurationSpec) Reset() {
	*x = ConfigurationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationSpec) ProtoMessage() {}

func (x *ConfigurationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurationSpec.ProtoReflect.Descriptor instead.
func (*ConfigurationSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_configurations_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigurationSpec) GetTemplate() *RevisionTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

type ConfigurationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ObservedGeneration is the 'Generation' of the Configuration that was last processed by the controller.
	// The observed generation is updated even if the controller failed to process the spec and create the Revision.
	// Clients polling for completed reconciliation should poll until observedGeneration = metadata.generation,
	// and the Ready condition's status is True or False.
	ObservedGeneration int32 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// LatestCreatedRevisionName is the last revision that was created from this Configuration.
	// It might not be ready yet, for that use LatestReadyRevisionName.
	LatestCreatedRevisionName string `protobuf:"bytes,2,opt,name=latest_created_revision_name,json=latestCreatedRevisionName,proto3" json:"latest_created_revision_name,omitempty"`
	// LatestReadyRevisionName holds the name of the latest
	// Revision stamped out from this Configuration that has had its "Ready" condition become "True".
	LatestReadyRevisionName string `protobuf:"bytes,3,opt,name=latest_ready_revision_name,json=latestReadyRevisionName,proto3" json:"latest_ready_revision_name,omitempty"`
	// Conditions communicates information about ongoing/complete reconciliation processes that bring the "spec"
	// inline with the observed state of the world.
	Conditions []*Condition `protobuf:"bytes,4,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *ConfigurationStatus) Reset() {
	*x = ConfigurationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationStatus) ProtoMessage() {}

func (x *ConfigurationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurationStatus.ProtoReflect.Descriptor instead.
func (*ConfigurationStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_configurations_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigurationStatus) GetObservedGeneration() int32 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *ConfigurationStatus) GetLatestCreatedRevisionName() string {
	if x != nil {
		return x.LatestCreatedRevisionName
	}
	return ""
}

func (x *ConfigurationStatus) GetLatestReadyRevisionName() string {
	if x != nil {
		return x.LatestReadyRevisionName
	}
	return ""
}

func (x *ConfigurationStatus) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type ListConfigurationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace from which the configurations should be listed. For Cloud Run (fully managed), replace {namespace_id} with the project ID or number.
	// Authorization requires the following IAM permission on the specified resource parent:
	// run.configurations.list
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of records that should be returned.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Not currently used by Cloud Run
	IncludeUninitialized bool `protobuf:"varint,3,opt,name=include_uninitialized,json=includeUninitialized,proto3" json:"include_uninitialized,omitempty"`
	// Allows to filter resources based on a specific value for a field name. Send this in a query string format. i.e. 'metadata.name%3Dlorem'.
	// Not currently used by Cloud Run.
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

func (x *ListConfigurationsRequest) Reset() {
	*x = ListConfigurationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigurationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigurationsRequest) ProtoMessage() {}

func (x *ListConfigurationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigurationsRequest.ProtoReflect.Descriptor instead.
func (*ListConfigurationsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_configurations_proto_rawDescGZIP(), []int{4}
}

func (x *ListConfigurationsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListConfigurationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListConfigurationsRequest) GetIncludeUninitialized() bool {
	if x != nil {
		return x.IncludeUninitialized
	}
	return false
}

func (x *ListConfigurationsRequest) GetFieldSelector() string {
	if x != nil {
		return x.FieldSelector
	}
	return ""
}

func (x *ListConfigurationsRequest) GetLabelSelector() string {
	if x != nil {
		return x.LabelSelector
	}
	return ""
}

func (x *ListConfigurationsRequest) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *ListConfigurationsRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

func (x *ListConfigurationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListConfigurationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The kind of this resource, in this case "ConfigurationList".
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Metadata associated with this configuration list.
	Metadata *ListMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// List of Configurations.
	Items []*Configuration `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,5,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListConfigurationsResponse) Reset() {
	*x = ListConfigurationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigurationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigurationsResponse) ProtoMessage() {}

func (x *ListConfigurationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_configurations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigurationsResponse.ProtoReflect.Descriptor instead.
func (*ListConfigurationsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_configurations_proto_rawDescGZIP(), []int{5}
}

func (x *ListConfigurationsResponse) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ListConfigurationsResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ListConfigurationsResponse) GetMetadata() *ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ListConfigurationsResponse) GetItems() []*Configuration {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListConfigurationsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

var File_google_cloud_run_v1_configurations_proto protoreflect.FileDescriptor

var file_google_cloud_run_v1_configurations_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x3a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x40, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x56, 0x0a,
	0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x41, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x84, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a,
	0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f,
	0x0a, 0x1c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3b, 0x0a, 0x1a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb3, 0x02, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x33, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x55, 0x6e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xe8, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75,
	0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x32, 0xb9, 0x02,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x66, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x1a, 0x46, 0xca, 0x41, 0x12, 0x72, 0x75, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v1_configurations_proto_rawDescOnce sync.Once
	file_google_cloud_run_v1_configurations_proto_rawDescData = file_google_cloud_run_v1_configurations_proto_rawDesc
)

func file_google_cloud_run_v1_configurations_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v1_configurations_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v1_configurations_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v1_configurations_proto_rawDescData)
	})
	return file_google_cloud_run_v1_configurations_proto_rawDescData
}

var file_google_cloud_run_v1_configurations_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_run_v1_configurations_proto_goTypes = []interface{}{
	(*GetConfigurationRequest)(nil),    // 0: google.cloud.run.v1.GetConfigurationRequest
	(*Configuration)(nil),              // 1: google.cloud.run.v1.Configuration
	(*ConfigurationSpec)(nil),          // 2: google.cloud.run.v1.ConfigurationSpec
	(*ConfigurationStatus)(nil),        // 3: google.cloud.run.v1.ConfigurationStatus
	(*ListConfigurationsRequest)(nil),  // 4: google.cloud.run.v1.ListConfigurationsRequest
	(*ListConfigurationsResponse)(nil), // 5: google.cloud.run.v1.ListConfigurationsResponse
	(*ObjectMeta)(nil),                 // 6: google.cloud.run.v1.ObjectMeta
	(*RevisionTemplate)(nil),           // 7: google.cloud.run.v1.RevisionTemplate
	(*Condition)(nil),                  // 8: google.cloud.run.v1.Condition
	(*ListMeta)(nil),                   // 9: google.cloud.run.v1.ListMeta
}
var file_google_cloud_run_v1_configurations_proto_depIdxs = []int32{
	6, // 0: google.cloud.run.v1.Configuration.metadata:type_name -> google.cloud.run.v1.ObjectMeta
	2, // 1: google.cloud.run.v1.Configuration.spec:type_name -> google.cloud.run.v1.ConfigurationSpec
	3, // 2: google.cloud.run.v1.Configuration.status:type_name -> google.cloud.run.v1.ConfigurationStatus
	7, // 3: google.cloud.run.v1.ConfigurationSpec.template:type_name -> google.cloud.run.v1.RevisionTemplate
	8, // 4: google.cloud.run.v1.ConfigurationStatus.conditions:type_name -> google.cloud.run.v1.Condition
	9, // 5: google.cloud.run.v1.ListConfigurationsResponse.metadata:type_name -> google.cloud.run.v1.ListMeta
	1, // 6: google.cloud.run.v1.ListConfigurationsResponse.items:type_name -> google.cloud.run.v1.Configuration
	0, // 7: google.cloud.run.v1.Configurations.GetConfiguration:input_type -> google.cloud.run.v1.GetConfigurationRequest
	4, // 8: google.cloud.run.v1.Configurations.ListConfigurations:input_type -> google.cloud.run.v1.ListConfigurationsRequest
	1, // 9: google.cloud.run.v1.Configurations.GetConfiguration:output_type -> google.cloud.run.v1.Configuration
	5, // 10: google.cloud.run.v1.Configurations.ListConfigurations:output_type -> google.cloud.run.v1.ListConfigurationsResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v1_configurations_proto_init() }
func file_google_cloud_run_v1_configurations_proto_init() {
	if File_google_cloud_run_v1_configurations_proto != nil {
		return
	}
	file_google_cloud_run_v1_common_proto_init()
	file_google_cloud_run_v1_metav1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v1_configurations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigurationRequest); i {
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
		file_google_cloud_run_v1_configurations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_google_cloud_run_v1_configurations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurationSpec); i {
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
		file_google_cloud_run_v1_configurations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurationStatus); i {
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
		file_google_cloud_run_v1_configurations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigurationsRequest); i {
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
		file_google_cloud_run_v1_configurations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigurationsResponse); i {
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
			RawDescriptor: file_google_cloud_run_v1_configurations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_run_v1_configurations_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v1_configurations_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v1_configurations_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v1_configurations_proto = out.File
	file_google_cloud_run_v1_configurations_proto_rawDesc = nil
	file_google_cloud_run_v1_configurations_proto_goTypes = nil
	file_google_cloud_run_v1_configurations_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConfigurationsClient is the client API for Configurations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationsClient interface {
	// Get information about a configuration.
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*Configuration, error)
	// List configurations.
	ListConfigurations(ctx context.Context, in *ListConfigurationsRequest, opts ...grpc.CallOption) (*ListConfigurationsResponse, error)
}

type configurationsClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationsClient(cc grpc.ClientConnInterface) ConfigurationsClient {
	return &configurationsClient{cc}
}

func (c *configurationsClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/google.cloud.run.v1.Configurations/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationsClient) ListConfigurations(ctx context.Context, in *ListConfigurationsRequest, opts ...grpc.CallOption) (*ListConfigurationsResponse, error) {
	out := new(ListConfigurationsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.run.v1.Configurations/ListConfigurations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationsServer is the server API for Configurations service.
type ConfigurationsServer interface {
	// Get information about a configuration.
	GetConfiguration(context.Context, *GetConfigurationRequest) (*Configuration, error)
	// List configurations.
	ListConfigurations(context.Context, *ListConfigurationsRequest) (*ListConfigurationsResponse, error)
}

// UnimplementedConfigurationsServer can be embedded to have forward compatible implementations.
type UnimplementedConfigurationsServer struct {
}

func (*UnimplementedConfigurationsServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*Configuration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (*UnimplementedConfigurationsServer) ListConfigurations(context.Context, *ListConfigurationsRequest) (*ListConfigurationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfigurations not implemented")
}

func RegisterConfigurationsServer(s *grpc.Server, srv ConfigurationsServer) {
	s.RegisterService(&_Configurations_serviceDesc, srv)
}

func _Configurations_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationsServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.run.v1.Configurations/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationsServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configurations_ListConfigurations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigurationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationsServer).ListConfigurations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.run.v1.Configurations/ListConfigurations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationsServer).ListConfigurations(ctx, req.(*ListConfigurationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configurations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.run.v1.Configurations",
	HandlerType: (*ConfigurationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfiguration",
			Handler:    _Configurations_GetConfiguration_Handler,
		},
		{
			MethodName: "ListConfigurations",
			Handler:    _Configurations_ListConfigurations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/run/v1/configurations.proto",
}
