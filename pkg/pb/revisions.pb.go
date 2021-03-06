// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: google/cloud/run/v1/revisions.proto

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

type GetRevisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the revision to retrieve. For Cloud Run (fully managed), replace {namespace_id} with the project ID or number.
	// Authorization requires the following IAM permission on the specified resource name:
	// run.revisions.get
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRevisionRequest) Reset() {
	*x = GetRevisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRevisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRevisionRequest) ProtoMessage() {}

func (x *GetRevisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRevisionRequest.ProtoReflect.Descriptor instead.
func (*GetRevisionRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_revisions_proto_rawDescGZIP(), []int{0}
}

func (x *GetRevisionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The kind of this resource, in this case "Revision".
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Metadata associated with this Revision, including name, namespace, labels, and annotations.
	Metadata *ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec holds the desired state of the Revision (from the client).
	Spec *RevisionSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status communicates the observed state of the Revision (from the controller).
	Status *RevisionStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Revision) Reset() {
	*x = Revision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revision) ProtoMessage() {}

func (x *Revision) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revision.ProtoReflect.Descriptor instead.
func (*Revision) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_revisions_proto_rawDescGZIP(), []int{1}
}

func (x *Revision) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Revision) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Revision) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Revision) GetSpec() *RevisionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Revision) GetStatus() *RevisionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type RevisionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ObservedGeneration is the 'Generation' of the Revision that was last processed by the controller.
	// Clients polling for completed reconciliation should poll until observedGeneration = metadata.generation,
	// and the Ready condition's status is True or False.
	ObservedGeneration int32 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Conditions communicates information about ongoing/complete reconciliation processes that bring the "spec"
	// inline with the observed state of the world.
	// As a Revision is being prepared, it will incrementally update conditions.
	// Revision-specific conditions include:
	// * "ResourcesAvailable": True when underlying resources have been provisioned.
	// * "ContainerHealthy": True when the Revision readiness check completes.
	// * "Active": True when the Revision may receive traffic.
	Conditions []*Condition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Specifies the generated logging url for this particular revision based on the revision url template specified in the controller's config.
	LogUrl string `protobuf:"bytes,3,opt,name=log_url,json=logUrl,proto3" json:"log_url,omitempty"`
	// Not currently used by Cloud Run.
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// ImageDigest holds the resolved digest for the image specified within .Spec.Container.Image. The digest is resolved during the creation of Revision.
	// This field holds the digest value regardless of whether a tag or digest was originally specified in the Container object.
	ImageDigest string `protobuf:"bytes,5,opt,name=image_digest,json=imageDigest,proto3" json:"image_digest,omitempty"`
}

func (x *RevisionStatus) Reset() {
	*x = RevisionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevisionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevisionStatus) ProtoMessage() {}

func (x *RevisionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevisionStatus.ProtoReflect.Descriptor instead.
func (*RevisionStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_revisions_proto_rawDescGZIP(), []int{2}
}

func (x *RevisionStatus) GetObservedGeneration() int32 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *RevisionStatus) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *RevisionStatus) GetLogUrl() string {
	if x != nil {
		return x.LogUrl
	}
	return ""
}

func (x *RevisionStatus) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RevisionStatus) GetImageDigest() string {
	if x != nil {
		return x.ImageDigest
	}
	return ""
}

type ListRevisionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace from which the revisions should be listed. For Cloud Run (fully managed), replace {namespace_id} with the project ID or number.
	// Authorization requires the following IAM permission on the specified resource parent:
	// run.revisions.list
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

func (x *ListRevisionsRequest) Reset() {
	*x = ListRevisionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRevisionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRevisionsRequest) ProtoMessage() {}

func (x *ListRevisionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRevisionsRequest.ProtoReflect.Descriptor instead.
func (*ListRevisionsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_revisions_proto_rawDescGZIP(), []int{3}
}

func (x *ListRevisionsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListRevisionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRevisionsRequest) GetIncludeUninitialized() bool {
	if x != nil {
		return x.IncludeUninitialized
	}
	return false
}

func (x *ListRevisionsRequest) GetFieldSelector() string {
	if x != nil {
		return x.FieldSelector
	}
	return ""
}

func (x *ListRevisionsRequest) GetLabelSelector() string {
	if x != nil {
		return x.LabelSelector
	}
	return ""
}

func (x *ListRevisionsRequest) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *ListRevisionsRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

func (x *ListRevisionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListRevisionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The API version for this call such as "serving.knative.dev/v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The kind of this resource, in this case "RevisionList".
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Metadata associated with this revision list.
	Metadata *ListMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// List of Revisions.
	Items []*Revision `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,5,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListRevisionsResponse) Reset() {
	*x = ListRevisionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRevisionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRevisionsResponse) ProtoMessage() {}

func (x *ListRevisionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRevisionsResponse.ProtoReflect.Descriptor instead.
func (*ListRevisionsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_revisions_proto_rawDescGZIP(), []int{4}
}

func (x *ListRevisionsResponse) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ListRevisionsResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ListRevisionsResponse) GetMetadata() *ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ListRevisionsResponse) GetItems() []*Revision {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListRevisionsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

type DeleteRevisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the revision to delete. For Cloud Run (fully managed), replace {namespace_id} with the project ID or number.
	// Authorization requires the following IAM permission on the specified resource name:
	// run.revisions.delete
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the propagation policy of delete. Cloud Run currently ignores this setting, and deletes in the background.
	// Please see kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/ for more information.
	PropagationPolicy string `protobuf:"bytes,2,opt,name=propagation_policy,json=propagationPolicy,proto3" json:"propagation_policy,omitempty"`
	// Cloud Run currently ignores this parameter.
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	// Cloud Run currently ignores this parameter.
	ApiVersion string `protobuf:"bytes,4,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Indicates that the server should validate the request and populate default values without persisting the request.
	// Supported values: all
	DryRun string `protobuf:"bytes,5,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *DeleteRevisionRequest) Reset() {
	*x = DeleteRevisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRevisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRevisionRequest) ProtoMessage() {}

func (x *DeleteRevisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_revisions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRevisionRequest.ProtoReflect.Descriptor instead.
func (*DeleteRevisionRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_revisions_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRevisionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteRevisionRequest) GetPropagationPolicy() string {
	if x != nil {
		return x.PropagationPolicy
	}
	return ""
}

func (x *DeleteRevisionRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *DeleteRevisionRequest) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *DeleteRevisionRequest) GetDryRun() string {
	if x != nil {
		return x.DryRun
	}
	return ""
}

var File_google_cloud_run_v1_revisions_proto protoreflect.FileDescriptor

var file_google_cloud_run_v1_revisions_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xf0, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x35, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x55, 0x72,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0xae, 0x02, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x75, 0x6e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x55, 0x6e, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xde, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61,
	0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e,
	0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x70, 0x61,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x64,
	0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72,
	0x79, 0x52, 0x75, 0x6e, 0x32, 0xf3, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x57, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x1a, 0x46, 0xca, 0x41, 0x12, 0x72, 0x75, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v1_revisions_proto_rawDescOnce sync.Once
	file_google_cloud_run_v1_revisions_proto_rawDescData = file_google_cloud_run_v1_revisions_proto_rawDesc
)

func file_google_cloud_run_v1_revisions_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v1_revisions_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v1_revisions_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v1_revisions_proto_rawDescData)
	})
	return file_google_cloud_run_v1_revisions_proto_rawDescData
}

var file_google_cloud_run_v1_revisions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_run_v1_revisions_proto_goTypes = []interface{}{
	(*GetRevisionRequest)(nil),    // 0: google.cloud.run.v1.GetRevisionRequest
	(*Revision)(nil),              // 1: google.cloud.run.v1.Revision
	(*RevisionStatus)(nil),        // 2: google.cloud.run.v1.RevisionStatus
	(*ListRevisionsRequest)(nil),  // 3: google.cloud.run.v1.ListRevisionsRequest
	(*ListRevisionsResponse)(nil), // 4: google.cloud.run.v1.ListRevisionsResponse
	(*DeleteRevisionRequest)(nil), // 5: google.cloud.run.v1.DeleteRevisionRequest
	(*ObjectMeta)(nil),            // 6: google.cloud.run.v1.ObjectMeta
	(*RevisionSpec)(nil),          // 7: google.cloud.run.v1.RevisionSpec
	(*Condition)(nil),             // 8: google.cloud.run.v1.Condition
	(*ListMeta)(nil),              // 9: google.cloud.run.v1.ListMeta
	(*Status)(nil),                // 10: google.cloud.run.v1.Status
}
var file_google_cloud_run_v1_revisions_proto_depIdxs = []int32{
	6,  // 0: google.cloud.run.v1.Revision.metadata:type_name -> google.cloud.run.v1.ObjectMeta
	7,  // 1: google.cloud.run.v1.Revision.spec:type_name -> google.cloud.run.v1.RevisionSpec
	2,  // 2: google.cloud.run.v1.Revision.status:type_name -> google.cloud.run.v1.RevisionStatus
	8,  // 3: google.cloud.run.v1.RevisionStatus.conditions:type_name -> google.cloud.run.v1.Condition
	9,  // 4: google.cloud.run.v1.ListRevisionsResponse.metadata:type_name -> google.cloud.run.v1.ListMeta
	1,  // 5: google.cloud.run.v1.ListRevisionsResponse.items:type_name -> google.cloud.run.v1.Revision
	0,  // 6: google.cloud.run.v1.Revisions.GetRevision:input_type -> google.cloud.run.v1.GetRevisionRequest
	3,  // 7: google.cloud.run.v1.Revisions.ListRevisions:input_type -> google.cloud.run.v1.ListRevisionsRequest
	5,  // 8: google.cloud.run.v1.Revisions.DeleteRevision:input_type -> google.cloud.run.v1.DeleteRevisionRequest
	1,  // 9: google.cloud.run.v1.Revisions.GetRevision:output_type -> google.cloud.run.v1.Revision
	4,  // 10: google.cloud.run.v1.Revisions.ListRevisions:output_type -> google.cloud.run.v1.ListRevisionsResponse
	10, // 11: google.cloud.run.v1.Revisions.DeleteRevision:output_type -> google.cloud.run.v1.Status
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v1_revisions_proto_init() }
func file_google_cloud_run_v1_revisions_proto_init() {
	if File_google_cloud_run_v1_revisions_proto != nil {
		return
	}
	file_google_cloud_run_v1_common_proto_init()
	file_google_cloud_run_v1_metav1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v1_revisions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRevisionRequest); i {
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
		file_google_cloud_run_v1_revisions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revision); i {
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
		file_google_cloud_run_v1_revisions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevisionStatus); i {
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
		file_google_cloud_run_v1_revisions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRevisionsRequest); i {
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
		file_google_cloud_run_v1_revisions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRevisionsResponse); i {
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
		file_google_cloud_run_v1_revisions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRevisionRequest); i {
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
			RawDescriptor: file_google_cloud_run_v1_revisions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_run_v1_revisions_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v1_revisions_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v1_revisions_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v1_revisions_proto = out.File
	file_google_cloud_run_v1_revisions_proto_rawDesc = nil
	file_google_cloud_run_v1_revisions_proto_goTypes = nil
	file_google_cloud_run_v1_revisions_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RevisionsClient is the client API for Revisions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RevisionsClient interface {
	// Get information about a revision.
	GetRevision(ctx context.Context, in *GetRevisionRequest, opts ...grpc.CallOption) (*Revision, error)
	// List revisions.
	ListRevisions(ctx context.Context, in *ListRevisionsRequest, opts ...grpc.CallOption) (*ListRevisionsResponse, error)
	// Delete a revision.
	DeleteRevision(ctx context.Context, in *DeleteRevisionRequest, opts ...grpc.CallOption) (*Status, error)
}

type revisionsClient struct {
	cc grpc.ClientConnInterface
}

func NewRevisionsClient(cc grpc.ClientConnInterface) RevisionsClient {
	return &revisionsClient{cc}
}

func (c *revisionsClient) GetRevision(ctx context.Context, in *GetRevisionRequest, opts ...grpc.CallOption) (*Revision, error) {
	out := new(Revision)
	err := c.cc.Invoke(ctx, "/google.cloud.run.v1.Revisions/GetRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revisionsClient) ListRevisions(ctx context.Context, in *ListRevisionsRequest, opts ...grpc.CallOption) (*ListRevisionsResponse, error) {
	out := new(ListRevisionsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.run.v1.Revisions/ListRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revisionsClient) DeleteRevision(ctx context.Context, in *DeleteRevisionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/google.cloud.run.v1.Revisions/DeleteRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RevisionsServer is the server API for Revisions service.
type RevisionsServer interface {
	// Get information about a revision.
	GetRevision(context.Context, *GetRevisionRequest) (*Revision, error)
	// List revisions.
	ListRevisions(context.Context, *ListRevisionsRequest) (*ListRevisionsResponse, error)
	// Delete a revision.
	DeleteRevision(context.Context, *DeleteRevisionRequest) (*Status, error)
}

// UnimplementedRevisionsServer can be embedded to have forward compatible implementations.
type UnimplementedRevisionsServer struct {
}

func (*UnimplementedRevisionsServer) GetRevision(context.Context, *GetRevisionRequest) (*Revision, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRevision not implemented")
}
func (*UnimplementedRevisionsServer) ListRevisions(context.Context, *ListRevisionsRequest) (*ListRevisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRevisions not implemented")
}
func (*UnimplementedRevisionsServer) DeleteRevision(context.Context, *DeleteRevisionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRevision not implemented")
}

func RegisterRevisionsServer(s *grpc.Server, srv RevisionsServer) {
	s.RegisterService(&_Revisions_serviceDesc, srv)
}

func _Revisions_GetRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevisionsServer).GetRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.run.v1.Revisions/GetRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevisionsServer).GetRevision(ctx, req.(*GetRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Revisions_ListRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRevisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevisionsServer).ListRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.run.v1.Revisions/ListRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevisionsServer).ListRevisions(ctx, req.(*ListRevisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Revisions_DeleteRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevisionsServer).DeleteRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.run.v1.Revisions/DeleteRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevisionsServer).DeleteRevision(ctx, req.(*DeleteRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Revisions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.run.v1.Revisions",
	HandlerType: (*RevisionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRevision",
			Handler:    _Revisions_GetRevision_Handler,
		},
		{
			MethodName: "ListRevisions",
			Handler:    _Revisions_ListRevisions_Handler,
		},
		{
			MethodName: "DeleteRevision",
			Handler:    _Revisions_DeleteRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/run/v1/revisions.proto",
}
