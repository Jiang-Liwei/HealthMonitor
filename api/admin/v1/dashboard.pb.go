// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.0
// source: api/admin/v1/dashboard.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DashboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DashboardRequest) Reset() {
	*x = DashboardRequest{}
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardRequest) ProtoMessage() {}

func (x *DashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardRequest.ProtoReflect.Descriptor instead.
func (*DashboardRequest) Descriptor() ([]byte, []int) {
	return file_api_admin_v1_dashboard_proto_rawDescGZIP(), []int{0}
}

type DashboardReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardList         []*CardList         `protobuf:"bytes,1,rep,name=cardList,proto3" json:"cardList,omitempty"`
	LastSevenRecords []*LastSevenRecords `protobuf:"bytes,2,rep,name=lastSevenRecords,proto3" json:"lastSevenRecords,omitempty"`
}

func (x *DashboardReply) Reset() {
	*x = DashboardReply{}
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DashboardReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardReply) ProtoMessage() {}

func (x *DashboardReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardReply.ProtoReflect.Descriptor instead.
func (*DashboardReply) Descriptor() ([]byte, []int) {
	return file_api_admin_v1_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *DashboardReply) GetCardList() []*CardList {
	if x != nil {
		return x.CardList
	}
	return nil
}

func (x *DashboardReply) GetLastSevenRecords() []*LastSevenRecords {
	if x != nil {
		return x.LastSevenRecords
	}
	return nil
}

type CardList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num          int64           `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	ChangeRatio  float32         `protobuf:"fixed32,2,opt,name=changeRatio,proto3" json:"changeRatio,omitempty"`
	Category     string          `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	MonthDetails []*MonthDetails `protobuf:"bytes,4,rep,name=monthDetails,proto3" json:"monthDetails,omitempty"`
}

func (x *CardList) Reset() {
	*x = CardList{}
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CardList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardList) ProtoMessage() {}

func (x *CardList) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardList.ProtoReflect.Descriptor instead.
func (*CardList) Descriptor() ([]byte, []int) {
	return file_api_admin_v1_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *CardList) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *CardList) GetChangeRatio() float32 {
	if x != nil {
		return x.ChangeRatio
	}
	return 0
}

func (x *CardList) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CardList) GetMonthDetails() []*MonthDetails {
	if x != nil {
		return x.MonthDetails
	}
	return nil
}

type MonthDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num   int64  `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Month string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *MonthDetails) Reset() {
	*x = MonthDetails{}
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonthDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonthDetails) ProtoMessage() {}

func (x *MonthDetails) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonthDetails.ProtoReflect.Descriptor instead.
func (*MonthDetails) Descriptor() ([]byte, []int) {
	return file_api_admin_v1_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *MonthDetails) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *MonthDetails) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

type LastSevenRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordDate        int64  `protobuf:"varint,1,opt,name=recordDate,proto3" json:"recordDate,omitempty"`
	StatusSummary     string `protobuf:"bytes,2,opt,name=statusSummary,proto3" json:"statusSummary,omitempty"`
	SystolicPressure  uint64 `protobuf:"varint,3,opt,name=systolicPressure,proto3" json:"systolicPressure,omitempty"`
	DiastolicPressure uint64 `protobuf:"varint,4,opt,name=diastolicPressure,proto3" json:"diastolicPressure,omitempty"`
	Pulse             uint64 `protobuf:"varint,5,opt,name=pulse,proto3" json:"pulse,omitempty"`
	Mood              string `protobuf:"bytes,6,opt,name=mood,proto3" json:"mood,omitempty"`
	Id                string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LastSevenRecords) Reset() {
	*x = LastSevenRecords{}
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LastSevenRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastSevenRecords) ProtoMessage() {}

func (x *LastSevenRecords) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_v1_dashboard_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastSevenRecords.ProtoReflect.Descriptor instead.
func (*LastSevenRecords) Descriptor() ([]byte, []int) {
	return file_api_admin_v1_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *LastSevenRecords) GetRecordDate() int64 {
	if x != nil {
		return x.RecordDate
	}
	return 0
}

func (x *LastSevenRecords) GetStatusSummary() string {
	if x != nil {
		return x.StatusSummary
	}
	return ""
}

func (x *LastSevenRecords) GetSystolicPressure() uint64 {
	if x != nil {
		return x.SystolicPressure
	}
	return 0
}

func (x *LastSevenRecords) GetDiastolicPressure() uint64 {
	if x != nil {
		return x.DiastolicPressure
	}
	return 0
}

func (x *LastSevenRecords) GetPulse() uint64 {
	if x != nil {
		return x.Pulse
	}
	return 0
}

func (x *LastSevenRecords) GetMood() string {
	if x != nil {
		return x.Mood
	}
	return ""
}

func (x *LastSevenRecords) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_admin_v1_dashboard_proto protoreflect.FileDescriptor

var file_api_admin_v1_dashboard_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x90,
	0x01, 0x0a, 0x0e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x63, 0x61, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x76,
	0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x61, 0x73, 0x74, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x9a, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x3e,
	0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x36,
	0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0xec, 0x01, 0x0a, 0x10, 0x4c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x79, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x63, 0x50, 0x72, 0x65,
	0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x79, 0x73,
	0x74, 0x6f, 0x6c, 0x69, 0x63, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x64, 0x69, 0x61, 0x73, 0x74, 0x6f, 0x6c, 0x69, 0x63, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x64, 0x69, 0x61, 0x73, 0x74, 0x6f,
	0x6c, 0x69, 0x63, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x75, 0x6c, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x75, 0x6c, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x73, 0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x66, 0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x2f, 0x0a, 0x0c, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_admin_v1_dashboard_proto_rawDescOnce sync.Once
	file_api_admin_v1_dashboard_proto_rawDescData = file_api_admin_v1_dashboard_proto_rawDesc
)

func file_api_admin_v1_dashboard_proto_rawDescGZIP() []byte {
	file_api_admin_v1_dashboard_proto_rawDescOnce.Do(func() {
		file_api_admin_v1_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_v1_dashboard_proto_rawDescData)
	})
	return file_api_admin_v1_dashboard_proto_rawDescData
}

var file_api_admin_v1_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_admin_v1_dashboard_proto_goTypes = []any{
	(*DashboardRequest)(nil), // 0: api.admin.v1.DashboardRequest
	(*DashboardReply)(nil),   // 1: api.admin.v1.DashboardReply
	(*CardList)(nil),         // 2: api.admin.v1.CardList
	(*MonthDetails)(nil),     // 3: api.admin.v1.MonthDetails
	(*LastSevenRecords)(nil), // 4: api.admin.v1.LastSevenRecords
}
var file_api_admin_v1_dashboard_proto_depIdxs = []int32{
	2, // 0: api.admin.v1.DashboardReply.cardList:type_name -> api.admin.v1.CardList
	4, // 1: api.admin.v1.DashboardReply.lastSevenRecords:type_name -> api.admin.v1.LastSevenRecords
	3, // 2: api.admin.v1.CardList.monthDetails:type_name -> api.admin.v1.MonthDetails
	0, // 3: api.admin.v1.Dashboard.Dashboard:input_type -> api.admin.v1.DashboardRequest
	1, // 4: api.admin.v1.Dashboard.Dashboard:output_type -> api.admin.v1.DashboardReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_admin_v1_dashboard_proto_init() }
func file_api_admin_v1_dashboard_proto_init() {
	if File_api_admin_v1_dashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_admin_v1_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_admin_v1_dashboard_proto_goTypes,
		DependencyIndexes: file_api_admin_v1_dashboard_proto_depIdxs,
		MessageInfos:      file_api_admin_v1_dashboard_proto_msgTypes,
	}.Build()
	File_api_admin_v1_dashboard_proto = out.File
	file_api_admin_v1_dashboard_proto_rawDesc = nil
	file_api_admin_v1_dashboard_proto_goTypes = nil
	file_api_admin_v1_dashboard_proto_depIdxs = nil
}