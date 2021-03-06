// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: service/grpc/staff/protobuf/staff.proto

package protobuf

import (
	context "context"
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

// Request create company message
type CreateStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tel       string `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty"`
}

func (x *CreateStaffRequest) Reset() {
	*x = CreateStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffRequest) ProtoMessage() {}

func (x *CreateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffRequest.ProtoReflect.Descriptor instead.
func (*CreateStaffRequest) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStaffRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CreateStaffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateStaffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStaffRequest) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

// Response create company message
type CreateStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateStaffResponse) Reset() {
	*x = CreateStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaffResponse) ProtoMessage() {}

func (x *CreateStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaffResponse.ProtoReflect.Descriptor instead.
func (*CreateStaffResponse) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStaffResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateStaffResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request list company message
type ListStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    string   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage string   `protobuf:"bytes,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Filters []string `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	Sorts   []string `protobuf:"bytes,4,rep,name=sorts,proto3" json:"sorts,omitempty"`
}

func (x *ListStaffRequest) Reset() {
	*x = ListStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffRequest) ProtoMessage() {}

func (x *ListStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffRequest.ProtoReflect.Descriptor instead.
func (*ListStaffRequest) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{2}
}

func (x *ListStaffRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *ListStaffRequest) GetPerPage() string {
	if x != nil {
		return x.PerPage
	}
	return ""
}

func (x *ListStaffRequest) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListStaffRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

// Response list company message
type ListStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64                       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*ListStaffResponse_Output `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListStaffResponse) Reset() {
	*x = ListStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffResponse) ProtoMessage() {}

func (x *ListStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffResponse.ProtoReflect.Descriptor instead.
func (*ListStaffResponse) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{3}
}

func (x *ListStaffResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListStaffResponse) GetItems() []*ListStaffResponse_Output {
	if x != nil {
		return x.Items
	}
	return nil
}

// Request read company message
type ReadStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffId string `protobuf:"bytes,1,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
}

func (x *ReadStaffRequest) Reset() {
	*x = ReadStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadStaffRequest) ProtoMessage() {}

func (x *ReadStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadStaffRequest.ProtoReflect.Descriptor instead.
func (*ReadStaffRequest) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{4}
}

func (x *ReadStaffRequest) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

// Response read company message
type ReadStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReadStaffResponse) Reset() {
	*x = ReadStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadStaffResponse) ProtoMessage() {}

func (x *ReadStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadStaffResponse.ProtoReflect.Descriptor instead.
func (*ReadStaffResponse) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{5}
}

func (x *ReadStaffResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReadStaffResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request update company message
type UpdateStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tel       string `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty"`
}

func (x *UpdateStaffRequest) Reset() {
	*x = UpdateStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffRequest) ProtoMessage() {}

func (x *UpdateStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffRequest.ProtoReflect.Descriptor instead.
func (*UpdateStaffRequest) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateStaffRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *UpdateStaffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStaffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStaffRequest) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

// Response update company message
type UpdateStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStaffResponse) Reset() {
	*x = UpdateStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaffResponse) ProtoMessage() {}

func (x *UpdateStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaffResponse.ProtoReflect.Descriptor instead.
func (*UpdateStaffResponse) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{7}
}

// Request delete company message
type DeleteStaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStaffRequest) Reset() {
	*x = DeleteStaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaffRequest) ProtoMessage() {}

func (x *DeleteStaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaffRequest.ProtoReflect.Descriptor instead.
func (*DeleteStaffRequest) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteStaffRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response delete company message
type DeleteStaffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStaffResponse) Reset() {
	*x = DeleteStaffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaffResponse) ProtoMessage() {}

func (x *DeleteStaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaffResponse.ProtoReflect.Descriptor instead.
func (*DeleteStaffResponse) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{9}
}

type ListStaffResponse_Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListStaffResponse_Output) Reset() {
	*x = ListStaffResponse_Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStaffResponse_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaffResponse_Output) ProtoMessage() {}

func (x *ListStaffResponse_Output) ProtoReflect() protoreflect.Message {
	mi := &file_service_grpc_staff_protobuf_staff_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaffResponse_Output.ProtoReflect.Descriptor instead.
func (*ListStaffResponse_Output) Descriptor() ([]byte, []int) {
	return file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListStaffResponse_Output) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListStaffResponse_Output) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_service_grpc_staff_protobuf_staff_proto protoreflect.FileDescriptor

var file_service_grpc_staff_protobuf_staff_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x22, 0x69, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x22, 0x39,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x71, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x91, 0x01, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0x2c, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2d, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x22,
	0x37, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x65, 0x6c, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe9, 0x02, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_grpc_staff_protobuf_staff_proto_rawDescOnce sync.Once
	file_service_grpc_staff_protobuf_staff_proto_rawDescData = file_service_grpc_staff_protobuf_staff_proto_rawDesc
)

func file_service_grpc_staff_protobuf_staff_proto_rawDescGZIP() []byte {
	file_service_grpc_staff_protobuf_staff_proto_rawDescOnce.Do(func() {
		file_service_grpc_staff_protobuf_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_grpc_staff_protobuf_staff_proto_rawDescData)
	})
	return file_service_grpc_staff_protobuf_staff_proto_rawDescData
}

var file_service_grpc_staff_protobuf_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_grpc_staff_protobuf_staff_proto_goTypes = []interface{}{
	(*CreateStaffRequest)(nil),       // 0: protobuf.CreateStaffRequest
	(*CreateStaffResponse)(nil),      // 1: protobuf.CreateStaffResponse
	(*ListStaffRequest)(nil),         // 2: protobuf.ListStaffRequest
	(*ListStaffResponse)(nil),        // 3: protobuf.ListStaffResponse
	(*ReadStaffRequest)(nil),         // 4: protobuf.ReadStaffRequest
	(*ReadStaffResponse)(nil),        // 5: protobuf.ReadStaffResponse
	(*UpdateStaffRequest)(nil),       // 6: protobuf.UpdateStaffRequest
	(*UpdateStaffResponse)(nil),      // 7: protobuf.UpdateStaffResponse
	(*DeleteStaffRequest)(nil),       // 8: protobuf.DeleteStaffRequest
	(*DeleteStaffResponse)(nil),      // 9: protobuf.DeleteStaffResponse
	(*ListStaffResponse_Output)(nil), // 10: protobuf.ListStaffResponse.Output
}
var file_service_grpc_staff_protobuf_staff_proto_depIdxs = []int32{
	10, // 0: protobuf.ListStaffResponse.items:type_name -> protobuf.ListStaffResponse.Output
	2,  // 1: protobuf.StaffGrpcService.List:input_type -> protobuf.ListStaffRequest
	4,  // 2: protobuf.StaffGrpcService.Read:input_type -> protobuf.ReadStaffRequest
	0,  // 3: protobuf.StaffGrpcService.Create:input_type -> protobuf.CreateStaffRequest
	6,  // 4: protobuf.StaffGrpcService.Update:input_type -> protobuf.UpdateStaffRequest
	8,  // 5: protobuf.StaffGrpcService.Delete:input_type -> protobuf.DeleteStaffRequest
	3,  // 6: protobuf.StaffGrpcService.List:output_type -> protobuf.ListStaffResponse
	5,  // 7: protobuf.StaffGrpcService.Read:output_type -> protobuf.ReadStaffResponse
	1,  // 8: protobuf.StaffGrpcService.Create:output_type -> protobuf.CreateStaffResponse
	7,  // 9: protobuf.StaffGrpcService.Update:output_type -> protobuf.UpdateStaffResponse
	9,  // 10: protobuf.StaffGrpcService.Delete:output_type -> protobuf.DeleteStaffResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_service_grpc_staff_protobuf_staff_proto_init() }
func file_service_grpc_staff_protobuf_staff_proto_init() {
	if File_service_grpc_staff_protobuf_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffRequest); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaffResponse); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffRequest); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffResponse); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadStaffRequest); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadStaffResponse); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStaffRequest); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStaffResponse); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStaffRequest); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStaffResponse); i {
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
		file_service_grpc_staff_protobuf_staff_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStaffResponse_Output); i {
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
			RawDescriptor: file_service_grpc_staff_protobuf_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_grpc_staff_protobuf_staff_proto_goTypes,
		DependencyIndexes: file_service_grpc_staff_protobuf_staff_proto_depIdxs,
		MessageInfos:      file_service_grpc_staff_protobuf_staff_proto_msgTypes,
	}.Build()
	File_service_grpc_staff_protobuf_staff_proto = out.File
	file_service_grpc_staff_protobuf_staff_proto_rawDesc = nil
	file_service_grpc_staff_protobuf_staff_proto_goTypes = nil
	file_service_grpc_staff_protobuf_staff_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StaffGrpcServiceClient is the client API for StaffGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StaffGrpcServiceClient interface {
	List(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffResponse, error)
	Read(ctx context.Context, in *ReadStaffRequest, opts ...grpc.CallOption) (*ReadStaffResponse, error)
	Create(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error)
	Update(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*UpdateStaffResponse, error)
	Delete(ctx context.Context, in *DeleteStaffRequest, opts ...grpc.CallOption) (*DeleteStaffResponse, error)
}

type staffGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffGrpcServiceClient(cc grpc.ClientConnInterface) StaffGrpcServiceClient {
	return &staffGrpcServiceClient{cc}
}

func (c *staffGrpcServiceClient) List(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffResponse, error) {
	out := new(ListStaffResponse)
	err := c.cc.Invoke(ctx, "/protobuf.StaffGrpcService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffGrpcServiceClient) Read(ctx context.Context, in *ReadStaffRequest, opts ...grpc.CallOption) (*ReadStaffResponse, error) {
	out := new(ReadStaffResponse)
	err := c.cc.Invoke(ctx, "/protobuf.StaffGrpcService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffGrpcServiceClient) Create(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffResponse, error) {
	out := new(CreateStaffResponse)
	err := c.cc.Invoke(ctx, "/protobuf.StaffGrpcService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffGrpcServiceClient) Update(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*UpdateStaffResponse, error) {
	out := new(UpdateStaffResponse)
	err := c.cc.Invoke(ctx, "/protobuf.StaffGrpcService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffGrpcServiceClient) Delete(ctx context.Context, in *DeleteStaffRequest, opts ...grpc.CallOption) (*DeleteStaffResponse, error) {
	out := new(DeleteStaffResponse)
	err := c.cc.Invoke(ctx, "/protobuf.StaffGrpcService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffGrpcServiceServer is the server API for StaffGrpcService service.
type StaffGrpcServiceServer interface {
	List(context.Context, *ListStaffRequest) (*ListStaffResponse, error)
	Read(context.Context, *ReadStaffRequest) (*ReadStaffResponse, error)
	Create(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error)
	Update(context.Context, *UpdateStaffRequest) (*UpdateStaffResponse, error)
	Delete(context.Context, *DeleteStaffRequest) (*DeleteStaffResponse, error)
}

// UnimplementedStaffGrpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStaffGrpcServiceServer struct {
}

func (*UnimplementedStaffGrpcServiceServer) List(context.Context, *ListStaffRequest) (*ListStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStaffGrpcServiceServer) Read(context.Context, *ReadStaffRequest) (*ReadStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedStaffGrpcServiceServer) Create(context.Context, *CreateStaffRequest) (*CreateStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedStaffGrpcServiceServer) Update(context.Context, *UpdateStaffRequest) (*UpdateStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedStaffGrpcServiceServer) Delete(context.Context, *DeleteStaffRequest) (*DeleteStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterStaffGrpcServiceServer(s *grpc.Server, srv StaffGrpcServiceServer) {
	s.RegisterService(&_StaffGrpcService_serviceDesc, srv)
}

func _StaffGrpcService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffGrpcServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.StaffGrpcService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffGrpcServiceServer).List(ctx, req.(*ListStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffGrpcService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffGrpcServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.StaffGrpcService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffGrpcServiceServer).Read(ctx, req.(*ReadStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffGrpcService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffGrpcServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.StaffGrpcService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffGrpcServiceServer).Create(ctx, req.(*CreateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffGrpcService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffGrpcServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.StaffGrpcService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffGrpcServiceServer).Update(ctx, req.(*UpdateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffGrpcService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffGrpcServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.StaffGrpcService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffGrpcServiceServer).Delete(ctx, req.(*DeleteStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StaffGrpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.StaffGrpcService",
	HandlerType: (*StaffGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _StaffGrpcService_List_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _StaffGrpcService_Read_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _StaffGrpcService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffGrpcService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffGrpcService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/grpc/staff/protobuf/staff.proto",
}
