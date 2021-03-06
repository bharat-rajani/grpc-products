// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: products.proto

package products_v1

import (
	proto "github.com/golang/protobuf/proto"
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

type ClientRequestType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vendor string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
}

func (x *ClientRequestType) Reset() {
	*x = ClientRequestType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequestType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequestType) ProtoMessage() {}

func (x *ClientRequestType) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequestType.ProtoReflect.Descriptor instead.
func (*ClientRequestType) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRequestType) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

type ClientResponseType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductType string `protobuf:"bytes,1,opt,name=productType,proto3" json:"productType,omitempty"`
}

func (x *ClientResponseType) Reset() {
	*x = ClientResponseType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponseType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponseType) ProtoMessage() {}

func (x *ClientResponseType) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponseType.ProtoReflect.Descriptor instead.
func (*ClientResponseType) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{1}
}

func (x *ClientResponseType) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

type ClientRequestProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vendor      string `protobuf:"bytes,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	ProductType string `protobuf:"bytes,2,opt,name=productType,proto3" json:"productType,omitempty"`
}

func (x *ClientRequestProducts) Reset() {
	*x = ClientRequestProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequestProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequestProducts) ProtoMessage() {}

func (x *ClientRequestProducts) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequestProducts.ProtoReflect.Descriptor instead.
func (*ClientRequestProducts) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{2}
}

func (x *ClientRequestProducts) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *ClientRequestProducts) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

type ClientResponseProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *ProdsPrep `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *ClientResponseProducts) Reset() {
	*x = ClientResponseProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponseProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponseProducts) ProtoMessage() {}

func (x *ClientResponseProducts) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponseProducts.ProtoReflect.Descriptor instead.
func (*ClientResponseProducts) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{3}
}

func (x *ClientResponseProducts) GetProduct() *ProdsPrep {
	if x != nil {
		return x.Product
	}
	return nil
}

type ProdsPrep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	ShortUrl string `protobuf:"bytes,3,opt,name=shortUrl,proto3" json:"shortUrl,omitempty"`
}

func (x *ProdsPrep) Reset() {
	*x = ProdsPrep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdsPrep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdsPrep) ProtoMessage() {}

func (x *ProdsPrep) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdsPrep.ProtoReflect.Descriptor instead.
func (*ProdsPrep) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{4}
}

func (x *ProdsPrep) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProdsPrep) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ProdsPrep) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type AdminClientRequestProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product     *ProdsPrep `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Vendor      string     `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	ProductType string     `protobuf:"bytes,3,opt,name=productType,proto3" json:"productType,omitempty"`
}

func (x *AdminClientRequestProducts) Reset() {
	*x = AdminClientRequestProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminClientRequestProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminClientRequestProducts) ProtoMessage() {}

func (x *AdminClientRequestProducts) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminClientRequestProducts.ProtoReflect.Descriptor instead.
func (*AdminClientRequestProducts) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{5}
}

func (x *AdminClientRequestProducts) GetProduct() *ProdsPrep {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *AdminClientRequestProducts) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *AdminClientRequestProducts) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

type ProductCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ProductCount) Reset() {
	*x = ProductCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCount) ProtoMessage() {}

func (x *ProductCount) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCount.ProtoReflect.Descriptor instead.
func (*ProductCount) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{6}
}

func (x *ProductCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageContent string `protobuf:"bytes,1,opt,name=messageContent,proto3" json:"messageContent,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_products_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_products_proto_rawDescGZIP(), []int{7}
}

func (x *ChatMessage) GetMessageContent() string {
	if x != nil {
		return x.MessageContent
	}
	return ""
}

var File_products_proto protoreflect.FileDescriptor

var file_products_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x2b, 0x0a,
	0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x22, 0x36, 0x0a, 0x12, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x30, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x65, 0x70, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x22, 0x4f, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x65, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x88, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x65, 0x70, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x24, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xf0, 0x02, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x28, 0x01, 0x12, 0x49, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x74, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_products_proto_rawDescOnce sync.Once
	file_products_proto_rawDescData = file_products_proto_rawDesc
)

func file_products_proto_rawDescGZIP() []byte {
	file_products_proto_rawDescOnce.Do(func() {
		file_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_products_proto_rawDescData)
	})
	return file_products_proto_rawDescData
}

var file_products_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_products_proto_goTypes = []interface{}{
	(*ClientRequestType)(nil),          // 0: products.v1.ClientRequestType
	(*ClientResponseType)(nil),         // 1: products.v1.ClientResponseType
	(*ClientRequestProducts)(nil),      // 2: products.v1.ClientRequestProducts
	(*ClientResponseProducts)(nil),     // 3: products.v1.ClientResponseProducts
	(*ProdsPrep)(nil),                  // 4: products.v1.ProdsPrep
	(*AdminClientRequestProducts)(nil), // 5: products.v1.AdminClientRequestProducts
	(*ProductCount)(nil),               // 6: products.v1.ProductCount
	(*ChatMessage)(nil),                // 7: products.v1.ChatMessage
}
var file_products_proto_depIdxs = []int32{
	4, // 0: products.v1.ClientResponseProducts.product:type_name -> products.v1.ProdsPrep
	4, // 1: products.v1.AdminClientRequestProducts.product:type_name -> products.v1.ProdsPrep
	0, // 2: products.v1.ProductService.GetVendorProductTypes:input_type -> products.v1.ClientRequestType
	2, // 3: products.v1.ProductService.GetVendorProducts:input_type -> products.v1.ClientRequestProducts
	5, // 4: products.v1.ProductService.SetVendorProducts:input_type -> products.v1.AdminClientRequestProducts
	7, // 5: products.v1.ProductService.ChatVendorSales:input_type -> products.v1.ChatMessage
	1, // 6: products.v1.ProductService.GetVendorProductTypes:output_type -> products.v1.ClientResponseType
	3, // 7: products.v1.ProductService.GetVendorProducts:output_type -> products.v1.ClientResponseProducts
	6, // 8: products.v1.ProductService.SetVendorProducts:output_type -> products.v1.ProductCount
	7, // 9: products.v1.ProductService.ChatVendorSales:output_type -> products.v1.ChatMessage
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_products_proto_init() }
func file_products_proto_init() {
	if File_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequestType); i {
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
		file_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponseType); i {
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
		file_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequestProducts); i {
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
		file_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponseProducts); i {
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
		file_products_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdsPrep); i {
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
		file_products_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminClientRequestProducts); i {
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
		file_products_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCount); i {
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
		file_products_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_products_proto_goTypes,
		DependencyIndexes: file_products_proto_depIdxs,
		MessageInfos:      file_products_proto_msgTypes,
	}.Build()
	File_products_proto = out.File
	file_products_proto_rawDesc = nil
	file_products_proto_goTypes = nil
	file_products_proto_depIdxs = nil
}
