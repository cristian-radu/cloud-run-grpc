// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: google/cloud/run/v1/common.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type TrafficTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ConfigurationName of a configuration to whose latest revision we will send this portion of traffic.
	// When the "status.latestReadyRevisionName" of the referenced configuration changes, we will automatically migrate traffic
	// from the prior "latest ready" revision to the new one. This field is never set in Route's status, only its spec.
	// This is mutually exclusive with RevisionName.
	// Cloud Run currently supports a single ConfigurationName.
	ConfigurationName string `protobuf:"bytes,1,opt,name=configuration_name,json=configurationName,proto3" json:"configuration_name,omitempty"`
	// RevisionName of a specific revision to which to send this portion of traffic. This is mutually exclusive with ConfigurationName.
	// Providing RevisionName in spec is not currently supported by Cloud Run.
	RevisionName string `protobuf:"bytes,2,opt,name=revision_name,json=revisionName,proto3" json:"revision_name,omitempty"`
	// Percent specifies percent of the traffic to this Revision or Configuration. This defaults to zero if unspecified.
	// Cloud Run currently requires 100 percent for a single ConfigurationName TrafficTarget entry.
	Percent int32 `protobuf:"varint,3,opt,name=percent,proto3" json:"percent,omitempty"`
	// Tag is optionally used to expose a dedicated url for referencing this target exclusively.
	Tag string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	// LatestRevision may be optionally provided to indicate that the latest ready Revision of the Configuration should be used for this traffic target.
	// When provided LatestRevision must be true if RevisionName is empty; it must be false when RevisionName is non-empty.
	LatestRevision bool `protobuf:"varint,5,opt,name=latest_revision,json=latestRevision,proto3" json:"latest_revision,omitempty"`
	// Output only. URL displays the URL for accessing tagged traffic targets. URL is displayed in status, and is disallowed on spec.
	// URL must contain a scheme (e.g. http://) and a hostname, but may not contain anything else (e.g. basic auth, url path, etc.)
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TrafficTarget) Reset() {
	*x = TrafficTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficTarget) ProtoMessage() {}

func (x *TrafficTarget) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficTarget.ProtoReflect.Descriptor instead.
func (*TrafficTarget) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *TrafficTarget) GetConfigurationName() string {
	if x != nil {
		return x.ConfigurationName
	}
	return ""
}

func (x *TrafficTarget) GetRevisionName() string {
	if x != nil {
		return x.RevisionName
	}
	return ""
}

func (x *TrafficTarget) GetPercent() int32 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *TrafficTarget) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *TrafficTarget) GetLatestRevision() bool {
	if x != nil {
		return x.LatestRevision
	}
	return false
}

func (x *TrafficTarget) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is used to communicate the status of the reconciliation process.
	// See also: https://github.com/knative/serving/blob/master/docs/spec/errors.md#error-conditions-and-reporting
	// Types common to all resources include: * "Ready": True when the Resource is ready.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// One-word CamelCase reason for the condition's last transition.
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	// Human readable message indicating details about the current status.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_transition_time,json=lastTransitionTime,proto3" json:"last_transition_time,omitempty"`
	// How to interpret failures of this condition, one of Error, Warning, Info
	Severity string `protobuf:"bytes,6,opt,name=severity,proto3" json:"severity,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Condition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Condition) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Condition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Condition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Condition) GetLastTransitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *Condition) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

type Addressable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information for connecting over HTTP(s).
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Addressable) Reset() {
	*x = Addressable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addressable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addressable) ProtoMessage() {}

func (x *Addressable) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addressable.ProtoReflect.Descriptor instead.
func (*Addressable) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *Addressable) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type RevisionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ContainerConcurrency specifies the maximum allowed in-flight (concurrent) requests per container instance of the Revision.
	// Cloud Run fully managed: supported, defaults to 80
	// Cloud Run for Anthos: supported, defaults to 0, which means concurrency to the application is not limited,
	// and the system decides the target concurrency for the autoscaler.
	ContainerConcurrency int32 `protobuf:"varint,1,opt,name=container_concurrency,json=containerConcurrency,proto3" json:"container_concurrency,omitempty"`
	// TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
	// Cloud Run fully managed: defaults to 300 seconds (5 minutes). Maximum allowed value is 900 seconds (15 minutes).
	// Cloud Run for Anthos: defaults to 300 seconds (5 minutes). Maximum allowed value is configurable by the cluster operator.
	TimeoutSeconds int32 `protobuf:"varint,2,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	// Email address of the IAM service account associated with the revision of the service.
	// The service account represents the identity of the running revision, and determines what permissions the revision has.
	// If not provided, the revision will use the project's default service account.
	ServiceAccountName string `protobuf:"bytes,3,opt,name=service_account_name,json=serviceAccountName,proto3" json:"service_account_name,omitempty"`
	// Containers holds the single container that defines the unit of execution for this Revision.
	// In the context of a Revision, we disallow a number of fields on this Container, including: name and lifecycle.
	// In Cloud Run, only a single container may be provided.
	// The runtime contract is documented here: https://github.com/knative/serving/blob/master/docs/runtime-contract.md
	Containers []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	Volumes    []*Volume    `protobuf:"bytes,5,rep,name=volumes,proto3" json:"volumes,omitempty"`
}

func (x *RevisionSpec) Reset() {
	*x = RevisionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevisionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevisionSpec) ProtoMessage() {}

func (x *RevisionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevisionSpec.ProtoReflect.Descriptor instead.
func (*RevisionSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *RevisionSpec) GetContainerConcurrency() int32 {
	if x != nil {
		return x.ContainerConcurrency
	}
	return 0
}

func (x *RevisionSpec) GetTimeoutSeconds() int32 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

func (x *RevisionSpec) GetServiceAccountName() string {
	if x != nil {
		return x.ServiceAccountName
	}
	return ""
}

func (x *RevisionSpec) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *RevisionSpec) GetVolumes() []*Volume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

type RevisionTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional metadata for this Revision, including labels and annotations. Name will be generated by the Configuration.
	// The following annotation keys set properties of the created revision:
	// autoscaling.knative.dev/minScale sets the minimum number of instances.
	// autoscaling.knative.dev/maxScale sets the maximum number of instances.
	// run.googleapis.com/cloudsql-instances sets Cloud SQL connections. Multiple values should be comma separated.
	// run.googleapis.com/vpc-access-connector sets a Serverless VPC Access connector.
	// run.googleapis.com/vpc-access-egress sets VPC egress. Supported values are all-traffic, all (deprecated), and private-ranges-only.
	// all-traffic and all provide the same functionality. all is deprecated but will continue to be supported. Prefer all-traffic.
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// RevisionSpec holds the desired state of the Revision (from the client).
	Spec *RevisionSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *RevisionTemplate) Reset() {
	*x = RevisionTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_run_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevisionTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevisionTemplate) ProtoMessage() {}

func (x *RevisionTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_run_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevisionTemplate.ProtoReflect.Descriptor instead.
func (*RevisionTemplate) Descriptor() ([]byte, []int) {
	return file_google_cloud_run_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *RevisionTemplate) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RevisionTemplate) GetSpec() *RevisionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

var File_google_cloud_run_v1_common_proto protoreflect.FileDescriptor

var file_google_cloud_run_v1_common_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a,
	0x0d, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2d,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x27,
	0x0a, 0x0f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xd3, 0x01, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x22,
	0x1f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x95, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52,
	0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_run_v1_common_proto_rawDescOnce sync.Once
	file_google_cloud_run_v1_common_proto_rawDescData = file_google_cloud_run_v1_common_proto_rawDesc
)

func file_google_cloud_run_v1_common_proto_rawDescGZIP() []byte {
	file_google_cloud_run_v1_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_run_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_run_v1_common_proto_rawDescData)
	})
	return file_google_cloud_run_v1_common_proto_rawDescData
}

var file_google_cloud_run_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_run_v1_common_proto_goTypes = []interface{}{
	(*TrafficTarget)(nil),         // 0: google.cloud.run.v1.TrafficTarget
	(*Condition)(nil),             // 1: google.cloud.run.v1.Condition
	(*Addressable)(nil),           // 2: google.cloud.run.v1.Addressable
	(*RevisionSpec)(nil),          // 3: google.cloud.run.v1.RevisionSpec
	(*RevisionTemplate)(nil),      // 4: google.cloud.run.v1.RevisionTemplate
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Container)(nil),             // 6: google.cloud.run.v1.Container
	(*Volume)(nil),                // 7: google.cloud.run.v1.Volume
	(*ObjectMeta)(nil),            // 8: google.cloud.run.v1.ObjectMeta
}
var file_google_cloud_run_v1_common_proto_depIdxs = []int32{
	5, // 0: google.cloud.run.v1.Condition.last_transition_time:type_name -> google.protobuf.Timestamp
	6, // 1: google.cloud.run.v1.RevisionSpec.containers:type_name -> google.cloud.run.v1.Container
	7, // 2: google.cloud.run.v1.RevisionSpec.volumes:type_name -> google.cloud.run.v1.Volume
	8, // 3: google.cloud.run.v1.RevisionTemplate.metadata:type_name -> google.cloud.run.v1.ObjectMeta
	3, // 4: google.cloud.run.v1.RevisionTemplate.spec:type_name -> google.cloud.run.v1.RevisionSpec
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_run_v1_common_proto_init() }
func file_google_cloud_run_v1_common_proto_init() {
	if File_google_cloud_run_v1_common_proto != nil {
		return
	}
	file_google_cloud_run_v1_corev1_proto_init()
	file_google_cloud_run_v1_metav1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_run_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficTarget); i {
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
		file_google_cloud_run_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_google_cloud_run_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addressable); i {
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
		file_google_cloud_run_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevisionSpec); i {
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
		file_google_cloud_run_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevisionTemplate); i {
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
			RawDescriptor: file_google_cloud_run_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_run_v1_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_run_v1_common_proto_depIdxs,
		MessageInfos:      file_google_cloud_run_v1_common_proto_msgTypes,
	}.Build()
	File_google_cloud_run_v1_common_proto = out.File
	file_google_cloud_run_v1_common_proto_rawDesc = nil
	file_google_cloud_run_v1_common_proto_goTypes = nil
	file_google_cloud_run_v1_common_proto_depIdxs = nil
}
