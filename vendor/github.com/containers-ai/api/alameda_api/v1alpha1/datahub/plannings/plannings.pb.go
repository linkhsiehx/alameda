// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/plannings/plannings.proto

package plannings

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContainerPlanning struct {
	Name                    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitPlannings          []*common.MetricData `protobuf:"bytes,2,rep,name=limit_plannings,json=limitPlannings,proto3" json:"limit_plannings,omitempty"`
	RequestPlannings        []*common.MetricData `protobuf:"bytes,3,rep,name=request_plannings,json=requestPlannings,proto3" json:"request_plannings,omitempty"`
	InitialLimitPlannings   []*common.MetricData `protobuf:"bytes,4,rep,name=initial_limit_plannings,json=initialLimitPlannings,proto3" json:"initial_limit_plannings,omitempty"`
	InitialRequestPlannings []*common.MetricData `protobuf:"bytes,5,rep,name=initial_request_plannings,json=initialRequestPlannings,proto3" json:"initial_request_plannings,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *ContainerPlanning) Reset()         { *m = ContainerPlanning{} }
func (m *ContainerPlanning) String() string { return proto.CompactTextString(m) }
func (*ContainerPlanning) ProtoMessage()    {}
func (*ContainerPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{0}
}

func (m *ContainerPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPlanning.Unmarshal(m, b)
}
func (m *ContainerPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPlanning.Marshal(b, m, deterministic)
}
func (m *ContainerPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPlanning.Merge(m, src)
}
func (m *ContainerPlanning) XXX_Size() int {
	return xxx_messageInfo_ContainerPlanning.Size(m)
}
func (m *ContainerPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPlanning proto.InternalMessageInfo

func (m *ContainerPlanning) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerPlanning) GetLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.LimitPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.RequestPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetInitialLimitPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialLimitPlannings
	}
	return nil
}

func (m *ContainerPlanning) GetInitialRequestPlannings() []*common.MetricData {
	if m != nil {
		return m.InitialRequestPlannings
	}
	return nil
}

type PodPlanning struct {
	ObjectMeta           *resources.ObjectMeta      `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	PlanningType         PlanningType               `protobuf:"varint,2,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	ApplyPlanningNow     bool                       `protobuf:"varint,3,opt,name=apply_planning_now,json=applyPlanningNow,proto3" json:"apply_planning_now,omitempty"`
	AssignPodPolicy      *resources.AssignPodPolicy `protobuf:"bytes,4,opt,name=assign_pod_policy,json=assignPodPolicy,proto3" json:"assign_pod_policy,omitempty"`
	ContainerPlannings   []*ContainerPlanning       `protobuf:"bytes,5,rep,name=container_plannings,json=containerPlannings,proto3" json:"container_plannings,omitempty"`
	StartTime            *timestamp.Timestamp       `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp       `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	TopController        *resources.Controller      `protobuf:"bytes,8,opt,name=top_controller,json=topController,proto3" json:"top_controller,omitempty"`
	PlanningId           string                     `protobuf:"bytes,9,opt,name=planning_id,json=planningId,proto3" json:"planning_id,omitempty"`
	TotalCost            float64                    `protobuf:"fixed64,10,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PodPlanning) Reset()         { *m = PodPlanning{} }
func (m *PodPlanning) String() string { return proto.CompactTextString(m) }
func (*PodPlanning) ProtoMessage()    {}
func (*PodPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{1}
}

func (m *PodPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodPlanning.Unmarshal(m, b)
}
func (m *PodPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodPlanning.Marshal(b, m, deterministic)
}
func (m *PodPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodPlanning.Merge(m, src)
}
func (m *PodPlanning) XXX_Size() int {
	return xxx_messageInfo_PodPlanning.Size(m)
}
func (m *PodPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_PodPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_PodPlanning proto.InternalMessageInfo

func (m *PodPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *PodPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *PodPlanning) GetApplyPlanningNow() bool {
	if m != nil {
		return m.ApplyPlanningNow
	}
	return false
}

func (m *PodPlanning) GetAssignPodPolicy() *resources.AssignPodPolicy {
	if m != nil {
		return m.AssignPodPolicy
	}
	return nil
}

func (m *PodPlanning) GetContainerPlannings() []*ContainerPlanning {
	if m != nil {
		return m.ContainerPlannings
	}
	return nil
}

func (m *PodPlanning) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *PodPlanning) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *PodPlanning) GetTopController() *resources.Controller {
	if m != nil {
		return m.TopController
	}
	return nil
}

func (m *PodPlanning) GetPlanningId() string {
	if m != nil {
		return m.PlanningId
	}
	return ""
}

func (m *PodPlanning) GetTotalCost() float64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

type ControllerPlanning struct {
	ObjectMeta           *resources.ObjectMeta      `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind                 resources.Kind             `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	PlanningType         PlanningType               `protobuf:"varint,3,opt,name=planning_type,json=planningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.PlanningType" json:"planning_type,omitempty"`
	CtlPlanningType      ControllerPlanningType     `protobuf:"varint,4,opt,name=ctl_planning_type,json=ctlPlanningType,proto3,enum=containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanningType" json:"ctl_planning_type,omitempty"`
	CtlPlanningSpec      *ControllerPlanningSpec    `protobuf:"bytes,5,opt,name=ctl_planning_spec,json=ctlPlanningSpec,proto3" json:"ctl_planning_spec,omitempty"`
	CtlPlanningSpecK8S   *ControllerPlanningSpecK8S `protobuf:"bytes,6,opt,name=ctl_planning_spec_k8s,json=ctlPlanningSpecK8s,proto3" json:"ctl_planning_spec_k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ControllerPlanning) Reset()         { *m = ControllerPlanning{} }
func (m *ControllerPlanning) String() string { return proto.CompactTextString(m) }
func (*ControllerPlanning) ProtoMessage()    {}
func (*ControllerPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{2}
}

func (m *ControllerPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerPlanning.Unmarshal(m, b)
}
func (m *ControllerPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerPlanning.Marshal(b, m, deterministic)
}
func (m *ControllerPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerPlanning.Merge(m, src)
}
func (m *ControllerPlanning) XXX_Size() int {
	return xxx_messageInfo_ControllerPlanning.Size(m)
}
func (m *ControllerPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerPlanning proto.InternalMessageInfo

func (m *ControllerPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func (m *ControllerPlanning) GetKind() resources.Kind {
	if m != nil {
		return m.Kind
	}
	return resources.Kind_POD
}

func (m *ControllerPlanning) GetPlanningType() PlanningType {
	if m != nil {
		return m.PlanningType
	}
	return PlanningType_PT_UNDEFINED
}

func (m *ControllerPlanning) GetCtlPlanningType() ControllerPlanningType {
	if m != nil {
		return m.CtlPlanningType
	}
	return ControllerPlanningType_CPT_UNDEFINED
}

func (m *ControllerPlanning) GetCtlPlanningSpec() *ControllerPlanningSpec {
	if m != nil {
		return m.CtlPlanningSpec
	}
	return nil
}

func (m *ControllerPlanning) GetCtlPlanningSpecK8S() *ControllerPlanningSpecK8S {
	if m != nil {
		return m.CtlPlanningSpecK8S
	}
	return nil
}

type ApplicationPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ApplicationPlanning) Reset()         { *m = ApplicationPlanning{} }
func (m *ApplicationPlanning) String() string { return proto.CompactTextString(m) }
func (*ApplicationPlanning) ProtoMessage()    {}
func (*ApplicationPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{3}
}

func (m *ApplicationPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationPlanning.Unmarshal(m, b)
}
func (m *ApplicationPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationPlanning.Marshal(b, m, deterministic)
}
func (m *ApplicationPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationPlanning.Merge(m, src)
}
func (m *ApplicationPlanning) XXX_Size() int {
	return xxx_messageInfo_ApplicationPlanning.Size(m)
}
func (m *ApplicationPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationPlanning proto.InternalMessageInfo

func (m *ApplicationPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

type NamespacePlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NamespacePlanning) Reset()         { *m = NamespacePlanning{} }
func (m *NamespacePlanning) String() string { return proto.CompactTextString(m) }
func (*NamespacePlanning) ProtoMessage()    {}
func (*NamespacePlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{4}
}

func (m *NamespacePlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespacePlanning.Unmarshal(m, b)
}
func (m *NamespacePlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespacePlanning.Marshal(b, m, deterministic)
}
func (m *NamespacePlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacePlanning.Merge(m, src)
}
func (m *NamespacePlanning) XXX_Size() int {
	return xxx_messageInfo_NamespacePlanning.Size(m)
}
func (m *NamespacePlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacePlanning.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacePlanning proto.InternalMessageInfo

func (m *NamespacePlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

type NodePlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NodePlanning) Reset()         { *m = NodePlanning{} }
func (m *NodePlanning) String() string { return proto.CompactTextString(m) }
func (*NodePlanning) ProtoMessage()    {}
func (*NodePlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{5}
}

func (m *NodePlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodePlanning.Unmarshal(m, b)
}
func (m *NodePlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodePlanning.Marshal(b, m, deterministic)
}
func (m *NodePlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePlanning.Merge(m, src)
}
func (m *NodePlanning) XXX_Size() int {
	return xxx_messageInfo_NodePlanning.Size(m)
}
func (m *NodePlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePlanning.DiscardUnknown(m)
}

var xxx_messageInfo_NodePlanning proto.InternalMessageInfo

func (m *NodePlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

type ClusterPlanning struct {
	ObjectMeta           *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterPlanning) Reset()         { *m = ClusterPlanning{} }
func (m *ClusterPlanning) String() string { return proto.CompactTextString(m) }
func (*ClusterPlanning) ProtoMessage()    {}
func (*ClusterPlanning) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c6a4155f49dbb5, []int{6}
}

func (m *ClusterPlanning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterPlanning.Unmarshal(m, b)
}
func (m *ClusterPlanning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterPlanning.Marshal(b, m, deterministic)
}
func (m *ClusterPlanning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterPlanning.Merge(m, src)
}
func (m *ClusterPlanning) XXX_Size() int {
	return xxx_messageInfo_ClusterPlanning.Size(m)
}
func (m *ClusterPlanning) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterPlanning.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterPlanning proto.InternalMessageInfo

func (m *ClusterPlanning) GetObjectMeta() *resources.ObjectMeta {
	if m != nil {
		return m.ObjectMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ContainerPlanning")
	proto.RegisterType((*PodPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.PodPlanning")
	proto.RegisterType((*ControllerPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ControllerPlanning")
	proto.RegisterType((*ApplicationPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ApplicationPlanning")
	proto.RegisterType((*NamespacePlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.NamespacePlanning")
	proto.RegisterType((*NodePlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.NodePlanning")
	proto.RegisterType((*ClusterPlanning)(nil), "containersai.alameda.v1alpha1.datahub.plannings.ClusterPlanning")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/plannings/plannings.proto", fileDescriptor_36c6a4155f49dbb5)
}

var fileDescriptor_36c6a4155f49dbb5 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x5f, 0x4f, 0xdb, 0x3a,
	0x18, 0xc6, 0x55, 0xda, 0x02, 0x7d, 0x0b, 0x94, 0x1a, 0xa1, 0x93, 0x83, 0x74, 0x44, 0xd5, 0xab,
	0x5e, 0x9c, 0x93, 0x88, 0x1e, 0x81, 0x7a, 0x74, 0x34, 0x69, 0xd0, 0x4d, 0x13, 0x63, 0x30, 0x94,
	0x71, 0x35, 0x4d, 0x8a, 0x5c, 0xc7, 0x2b, 0x1e, 0x8e, 0x6d, 0x62, 0x77, 0xa8, 0x17, 0xfb, 0x02,
	0xfb, 0x86, 0xfb, 0x22, 0xbb, 0x9e, 0xe2, 0x26, 0x29, 0x6d, 0x25, 0xd6, 0x02, 0xea, 0x9d, 0xed,
	0xe6, 0x7d, 0x7e, 0x79, 0xff, 0x3c, 0xa9, 0xe1, 0x08, 0x73, 0x1c, 0xd1, 0x10, 0x07, 0x58, 0x31,
	0xef, 0xeb, 0x01, 0xe6, 0xea, 0x1a, 0x1f, 0x78, 0x21, 0x36, 0xf8, 0x7a, 0xd0, 0xf3, 0x14, 0xc7,
	0x42, 0x30, 0xd1, 0xd7, 0xe3, 0x95, 0xab, 0x62, 0x69, 0x24, 0xf2, 0x88, 0x14, 0x06, 0x33, 0x41,
	0x63, 0x8d, 0x99, 0x9b, 0x8a, 0xb8, 0x99, 0x80, 0x9b, 0x0a, 0xb8, 0x79, 0xd8, 0xde, 0xc1, 0x83,
	0x20, 0x22, 0xa3, 0x48, 0x0a, 0x2f, 0xa2, 0x26, 0x66, 0x24, 0x65, 0xec, 0xb5, 0xe7, 0x7c, 0x37,
	0x33, 0x54, 0x34, 0x8b, 0x39, 0x7c, 0x30, 0x26, 0xa6, 0x5a, 0x0e, 0x62, 0x42, 0x75, 0x42, 0xc2,
	0xc9, 0xe9, 0x82, 0x61, 0x4a, 0x72, 0x46, 0x58, 0x4e, 0x3b, 0x9a, 0x33, 0x2c, 0x5f, 0xa5, 0x71,
	0xfb, 0x7d, 0x29, 0xfb, 0x9c, 0x7a, 0x76, 0xd7, 0x1b, 0x7c, 0xf6, 0x0c, 0x8b, 0xa8, 0x36, 0x38,
	0x52, 0xa3, 0x07, 0x9a, 0x3f, 0x8a, 0x50, 0xef, 0x66, 0x15, 0xbe, 0x4c, 0x33, 0x45, 0x08, 0x4a,
	0x02, 0x47, 0xd4, 0x29, 0x34, 0x0a, 0xad, 0x8a, 0x6f, 0xd7, 0x08, 0x43, 0x8d, 0xb3, 0x88, 0x99,
	0x20, 0xaf, 0x87, 0xb3, 0xd2, 0x28, 0xb6, 0xaa, 0xed, 0x8e, 0x3b, 0x5f, 0x8b, 0x46, 0xa5, 0x77,
	0xcf, 0x6d, 0xe9, 0x5f, 0x61, 0x83, 0xfd, 0x2d, 0x2b, 0x98, 0x51, 0x35, 0xa2, 0x50, 0x8f, 0xe9,
	0xed, 0x80, 0xea, 0xfb, 0x90, 0xe2, 0x13, 0x21, 0xdb, 0xa9, 0xe4, 0x18, 0xa3, 0xe0, 0x0f, 0x26,
	0x98, 0x61, 0x98, 0x07, 0xd3, 0x19, 0x95, 0x9e, 0x08, 0xdb, 0x4d, 0x85, 0xdf, 0x4d, 0x26, 0x66,
	0xe0, 0xcf, 0x8c, 0x38, 0x9b, 0x60, 0xf9, 0x89, 0xcc, 0x2c, 0x19, 0x7f, 0x2a, 0xcf, 0xe6, 0xcf,
	0x32, 0x54, 0x2f, 0x65, 0x98, 0x77, 0xf5, 0x13, 0x54, 0x65, 0xef, 0x0b, 0x25, 0x26, 0x48, 0x86,
	0xd2, 0x36, 0xb7, 0xda, 0xfe, 0x7f, 0x4e, 0xee, 0x78, 0xb2, 0xde, 0x5b, 0x8d, 0x73, 0x6a, 0xb0,
	0x0f, 0x32, 0x5f, 0xa3, 0x1e, 0x6c, 0x66, 0x39, 0x05, 0x89, 0x51, 0x9c, 0x95, 0x46, 0xa1, 0xb5,
	0xd5, 0x7e, 0xe1, 0x2e, 0x68, 0x60, 0x37, 0x7b, 0xdf, 0xab, 0xa1, 0xa2, 0xfe, 0x86, 0xba, 0xb7,
	0x43, 0x7f, 0x03, 0xc2, 0x4a, 0xf1, 0x61, 0x5e, 0xbd, 0x40, 0xc8, 0x3b, 0xa7, 0xd8, 0x28, 0xb4,
	0xd6, 0xfd, 0x6d, 0xfb, 0x4b, 0x16, 0x7c, 0x21, 0xef, 0x10, 0x87, 0x3a, 0xd6, 0x9a, 0xf5, 0x45,
	0xa0, 0x64, 0x18, 0x58, 0x47, 0x0d, 0x9d, 0x92, 0xcd, 0xfa, 0xe5, 0xc2, 0x59, 0x1f, 0x5b, 0xa5,
	0xa4, 0x9c, 0x56, 0xc7, 0xaf, 0xe1, 0xc9, 0x03, 0xa4, 0x61, 0x27, 0xd7, 0x9c, 0xe9, 0xee, 0xc9,
	0xc2, 0x55, 0x98, 0x31, 0xa5, 0x8f, 0xc8, 0xf4, 0x91, 0x46, 0xff, 0x01, 0x68, 0x83, 0x63, 0x13,
	0x24, 0xbe, 0x76, 0x56, 0x6d, 0x6e, 0x7b, 0xee, 0xc8, 0xf4, 0x6e, 0x66, 0x7a, 0xf7, 0x2a, 0x33,
	0xbd, 0x5f, 0xb1, 0x4f, 0x27, 0x7b, 0x74, 0x08, 0xeb, 0x54, 0x84, 0xa3, 0xc0, 0xb5, 0xdf, 0x06,
	0xae, 0x51, 0x11, 0xda, 0xb0, 0x1e, 0x6c, 0x19, 0xa9, 0x82, 0xe4, 0x5d, 0x62, 0xc9, 0x39, 0x8d,
	0x9d, 0xf5, 0x47, 0xce, 0x51, 0x37, 0x97, 0xf0, 0x37, 0x8d, 0x54, 0xe3, 0x2d, 0xda, 0x87, 0x6a,
	0xde, 0x60, 0x16, 0x3a, 0x15, 0xfb, 0x15, 0x82, 0xec, 0xe8, 0x34, 0x44, 0x7f, 0x01, 0x18, 0x69,
	0x30, 0x0f, 0x88, 0xd4, 0xc6, 0x81, 0x46, 0xa1, 0x55, 0xf0, 0x2b, 0xf6, 0xa4, 0x2b, 0xb5, 0x69,
	0x7e, 0x2f, 0x03, 0x1a, 0xcb, 0x2d, 0x69, 0xfe, 0x4f, 0xa1, 0x74, 0xc3, 0x44, 0x98, 0x8e, 0xfd,
	0xe1, 0xc2, 0xb2, 0x67, 0x4c, 0x84, 0xbe, 0x95, 0x98, 0xb5, 0x52, 0xf1, 0xf9, 0xad, 0xa4, 0xa1,
	0x4e, 0x0c, 0x0f, 0x26, 0x39, 0x25, 0xcb, 0x79, 0xf3, 0xa8, 0x61, 0x9d, 0x2c, 0xb6, 0x25, 0xd6,
	0x88, 0xe1, 0x97, 0x0f, 0x41, 0xb5, 0xa2, 0xc4, 0x29, 0xdb, 0x3e, 0x3c, 0x07, 0xf4, 0x83, 0xa2,
	0x64, 0x02, 0x9a, 0x1c, 0xa0, 0x6f, 0xb0, 0x3b, 0x03, 0x0d, 0x6e, 0x3a, 0x3a, 0xb5, 0xcb, 0xdb,
	0x67, 0x02, 0x9f, 0x75, 0xb4, 0x8f, 0xa6, 0xd8, 0x67, 0x1d, 0xdd, 0xd4, 0xb0, 0x73, 0xac, 0x14,
	0x67, 0x04, 0x1b, 0x26, 0xc5, 0x72, 0x86, 0xb1, 0x79, 0x0b, 0xf5, 0x0b, 0x1c, 0x51, 0xad, 0x30,
	0xa1, 0x4b, 0x42, 0x72, 0xd8, 0xb8, 0x90, 0xe1, 0xb2, 0x68, 0x12, 0x6a, 0x5d, 0x3e, 0xd0, 0x66,
	0x59, 0xf6, 0x3e, 0x79, 0xfd, 0xb1, 0xdb, 0x67, 0x26, 0xfd, 0x07, 0xbe, 0x77, 0x29, 0xfd, 0x07,
	0x33, 0x2f, 0xb9, 0x94, 0xcd, 0x77, 0x85, 0xec, 0xad, 0xda, 0x6f, 0xeb, 0xbf, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xa1, 0xc9, 0x02, 0xfc, 0x0f, 0x0b, 0x00, 0x00,
}