// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Product struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                int32    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Colors               []string `protobuf:"bytes,5,rep,name=colors,proto3" json:"colors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetColors() []string {
	if m != nil {
		return m.Colors
	}
	return nil
}

type ProductList struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProductList) Reset()         { *m = ProductList{} }
func (m *ProductList) String() string { return proto.CompactTextString(m) }
func (*ProductList) ProtoMessage()    {}
func (*ProductList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *ProductList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductList.Unmarshal(m, b)
}
func (m *ProductList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductList.Marshal(b, m, deterministic)
}
func (m *ProductList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductList.Merge(m, src)
}
func (m *ProductList) XXX_Size() int {
	return xxx_messageInfo_ProductList.Size(m)
}
func (m *ProductList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductList.DiscardUnknown(m)
}

var xxx_messageInfo_ProductList proto.InternalMessageInfo

func (m *ProductList) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "proto.Product")
	proto.RegisterType((*ProductList)(nil), "proto.ProductList")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x0e, 0xc2, 0x20,
	0x10, 0x86, 0x43, 0x29, 0xd5, 0x5e, 0x63, 0x87, 0x8b, 0x31, 0x8c, 0xa4, 0x13, 0x71, 0xe8, 0xa0,
	0x93, 0xef, 0xe0, 0x60, 0x78, 0x03, 0x85, 0x0e, 0x24, 0x5a, 0x08, 0xe0, 0xe8, 0xbb, 0x9b, 0x00,
	0x31, 0x4e, 0x77, 0xff, 0xff, 0x5d, 0x72, 0x1f, 0xec, 0x7c, 0x70, 0xe6, 0xad, 0xd3, 0xec, 0x83,
	0x4b, 0x0e, 0x59, 0x1e, 0xd3, 0x07, 0x36, 0xb7, 0xd2, 0xe3, 0x08, 0x8d, 0x35, 0x9c, 0x08, 0x22,
	0xa9, 0x6a, 0xac, 0x41, 0x84, 0x76, 0xbd, 0xbf, 0x16, 0xde, 0x08, 0x22, 0x7b, 0x95, 0x77, 0x14,
	0x30, 0x98, 0x25, 0xea, 0x60, 0x7d, 0xb2, 0x6e, 0xe5, 0x34, 0xa3, 0xff, 0x0a, 0xf7, 0xc0, 0x7c,
	0xb0, 0x7a, 0xe1, 0xad, 0x20, 0x92, 0xa9, 0x12, 0xf0, 0x00, 0x9d, 0x76, 0x4f, 0x17, 0x22, 0x67,
	0x82, 0xca, 0x5e, 0xd5, 0x34, 0x5d, 0x60, 0xa8, 0xef, 0xaf, 0x36, 0x26, 0x3c, 0xc2, 0xb6, 0x5a,
	0x46, 0x4e, 0x04, 0x95, 0xc3, 0x69, 0x2c, 0xba, 0x73, 0xbd, 0x52, 0x3f, 0xfe, 0xe8, 0x32, 0x38,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x04, 0x34, 0xe4, 0x69, 0xd8, 0x00, 0x00, 0x00,
}
