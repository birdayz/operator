// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/registry/registrypb/device_registry.proto

package registrypb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	context "golang.org/x/net/context"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

type Device struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Enabled              *wrappers.BoolValue `protobuf:"bytes,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Certificate          *Certificate        `protobuf:"bytes,4,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Tags                 []string            `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Namespace            string              `protobuf:"bytes,6,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *Device) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Device) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Device) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type CreateRequest struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *CreateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type Certificate struct {
	PemData              string   `protobuf:"bytes,1,opt,name=pem_data,json=pemData,proto3" json:"pem_data,omitempty"`
	Algorithm            string   `protobuf:"bytes,2,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Fingerprint          []byte   `protobuf:"bytes,3,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	FingerprintAlgorithm string   `protobuf:"bytes,4,opt,name=fingerprintAlgorithm,proto3" json:"fingerprintAlgorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{2}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetPemData() string {
	if m != nil {
		return m.PemData
	}
	return ""
}

func (m *Certificate) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

func (m *Certificate) GetFingerprint() []byte {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *Certificate) GetFingerprintAlgorithm() string {
	if m != nil {
		return m.FingerprintAlgorithm
	}
	return ""
}

type CreateResponse struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type GetByFingerprintRequest struct {
	Fingerprint          []byte   `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByFingerprintRequest) Reset()         { *m = GetByFingerprintRequest{} }
func (m *GetByFingerprintRequest) String() string { return proto.CompactTextString(m) }
func (*GetByFingerprintRequest) ProtoMessage()    {}
func (*GetByFingerprintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{4}
}

func (m *GetByFingerprintRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByFingerprintRequest.Unmarshal(m, b)
}
func (m *GetByFingerprintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByFingerprintRequest.Marshal(b, m, deterministic)
}
func (m *GetByFingerprintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByFingerprintRequest.Merge(m, src)
}
func (m *GetByFingerprintRequest) XXX_Size() int {
	return xxx_messageInfo_GetByFingerprintRequest.Size(m)
}
func (m *GetByFingerprintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByFingerprintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByFingerprintRequest proto.InternalMessageInfo

func (m *GetByFingerprintRequest) GetFingerprint() []byte {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

type DeviceForFingerprint struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceForFingerprint) Reset()         { *m = DeviceForFingerprint{} }
func (m *DeviceForFingerprint) String() string { return proto.CompactTextString(m) }
func (*DeviceForFingerprint) ProtoMessage()    {}
func (*DeviceForFingerprint) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{5}
}

func (m *DeviceForFingerprint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceForFingerprint.Unmarshal(m, b)
}
func (m *DeviceForFingerprint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceForFingerprint.Marshal(b, m, deterministic)
}
func (m *DeviceForFingerprint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceForFingerprint.Merge(m, src)
}
func (m *DeviceForFingerprint) XXX_Size() int {
	return xxx_messageInfo_DeviceForFingerprint.Size(m)
}
func (m *DeviceForFingerprint) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceForFingerprint.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceForFingerprint proto.InternalMessageInfo

func (m *DeviceForFingerprint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceForFingerprint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceForFingerprint) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type GetByFingerprintResponse struct {
	Devices              []*DeviceForFingerprint `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetByFingerprintResponse) Reset()         { *m = GetByFingerprintResponse{} }
func (m *GetByFingerprintResponse) String() string { return proto.CompactTextString(m) }
func (*GetByFingerprintResponse) ProtoMessage()    {}
func (*GetByFingerprintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{6}
}

func (m *GetByFingerprintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByFingerprintResponse.Unmarshal(m, b)
}
func (m *GetByFingerprintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByFingerprintResponse.Marshal(b, m, deterministic)
}
func (m *GetByFingerprintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByFingerprintResponse.Merge(m, src)
}
func (m *GetByFingerprintResponse) XXX_Size() int {
	return xxx_messageInfo_GetByFingerprintResponse.Size(m)
}
func (m *GetByFingerprintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByFingerprintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByFingerprintResponse proto.InternalMessageInfo

func (m *GetByFingerprintResponse) GetDevices() []*DeviceForFingerprint {
	if m != nil {
		return m.Devices
	}
	return nil
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{7}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{8}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type ListDevicesRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDevicesRequest) Reset()         { *m = ListDevicesRequest{} }
func (m *ListDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDevicesRequest) ProtoMessage()    {}
func (*ListDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{9}
}

func (m *ListDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDevicesRequest.Unmarshal(m, b)
}
func (m *ListDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDevicesRequest.Marshal(b, m, deterministic)
}
func (m *ListDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDevicesRequest.Merge(m, src)
}
func (m *ListDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDevicesRequest.Size(m)
}
func (m *ListDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDevicesRequest proto.InternalMessageInfo

func (m *ListDevicesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListDevicesRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type ListResponse struct {
	Devices              []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{10}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{11}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{12}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type UpdateRequest struct {
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,1,opt,name=fieldMask,proto3" json:"fieldMask,omitempty"`
	Device               *Device               `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{13}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *UpdateRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ff97fec4b2fb717, []int{14}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Device)(nil), "infinimesh.deviceregistry.Device")
	proto.RegisterType((*CreateRequest)(nil), "infinimesh.deviceregistry.CreateRequest")
	proto.RegisterType((*Certificate)(nil), "infinimesh.deviceregistry.Certificate")
	proto.RegisterType((*CreateResponse)(nil), "infinimesh.deviceregistry.CreateResponse")
	proto.RegisterType((*GetByFingerprintRequest)(nil), "infinimesh.deviceregistry.GetByFingerprintRequest")
	proto.RegisterType((*DeviceForFingerprint)(nil), "infinimesh.deviceregistry.DeviceForFingerprint")
	proto.RegisterType((*GetByFingerprintResponse)(nil), "infinimesh.deviceregistry.GetByFingerprintResponse")
	proto.RegisterType((*GetRequest)(nil), "infinimesh.deviceregistry.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "infinimesh.deviceregistry.GetResponse")
	proto.RegisterType((*ListDevicesRequest)(nil), "infinimesh.deviceregistry.ListDevicesRequest")
	proto.RegisterType((*ListResponse)(nil), "infinimesh.deviceregistry.ListResponse")
	proto.RegisterType((*DeleteRequest)(nil), "infinimesh.deviceregistry.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "infinimesh.deviceregistry.DeleteResponse")
	proto.RegisterType((*UpdateRequest)(nil), "infinimesh.deviceregistry.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "infinimesh.deviceregistry.UpdateResponse")
}

func init() {
	proto.RegisterFile("pkg/registry/registrypb/device_registry.proto", fileDescriptor_4ff97fec4b2fb717)
}

var fileDescriptor_4ff97fec4b2fb717 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xda, 0xae, 0x65, 0x37, 0x5b, 0x35, 0x59, 0x93, 0xc8, 0xaa, 0x09, 0x4a, 0x24, 0x46,
	0x79, 0x58, 0x2a, 0x75, 0x3c, 0x80, 0xf6, 0xb4, 0x0f, 0x75, 0x43, 0x1b, 0x2f, 0x11, 0x20, 0x84,
	0x34, 0x26, 0xb7, 0xb9, 0x4d, 0xad, 0xa5, 0x49, 0x88, 0x5d, 0xd0, 0xc4, 0x2b, 0x3f, 0x83, 0x7f,
	0xc3, 0xff, 0xe0, 0xb7, 0xa0, 0xc4, 0x71, 0x9a, 0xa4, 0x2c, 0xdb, 0x34, 0xf1, 0x96, 0x38, 0xe7,
	0x9e, 0x73, 0xee, 0xb1, 0xaf, 0x5b, 0xd8, 0x0d, 0xaf, 0xdc, 0x7e, 0x84, 0x2e, 0xe3, 0x22, 0xba,
	0xce, 0x1e, 0xc2, 0x51, 0xdf, 0xc1, 0x6f, 0x6c, 0x8c, 0x97, 0x6a, 0xc5, 0x0a, 0xa3, 0x40, 0x04,
	0x64, 0x8b, 0xf9, 0x13, 0xe6, 0xb3, 0x19, 0xf2, 0xa9, 0x25, 0x11, 0x0a, 0xd0, 0xe9, 0xba, 0x41,
	0xe0, 0x7a, 0xd8, 0x4f, 0x80, 0xa3, 0xf9, 0xa4, 0x3f, 0x61, 0xe8, 0x39, 0x97, 0x33, 0xca, 0xaf,
	0x64, 0x71, 0xe7, 0x49, 0x19, 0xf1, 0x3d, 0xa2, 0x61, 0x88, 0x11, 0x97, 0xdf, 0xcd, 0x3f, 0x1a,
	0x34, 0x8f, 0x13, 0x52, 0xd2, 0x86, 0x1a, 0x73, 0x0c, 0xad, 0xab, 0xf5, 0x56, 0xed, 0x1a, 0x73,
	0x08, 0x81, 0x86, 0x4f, 0x67, 0x68, 0xd4, 0x92, 0x95, 0xe4, 0x99, 0xbc, 0x82, 0x16, 0xfa, 0x74,
	0xe4, 0xa1, 0x63, 0xd4, 0xbb, 0x5a, 0x4f, 0x1f, 0x74, 0x2c, 0x29, 0x60, 0x29, 0x01, 0xeb, 0x30,
	0x08, 0xbc, 0x8f, 0xd4, 0x9b, 0xa3, 0xad, 0xa0, 0xe4, 0x14, 0xf4, 0x31, 0x46, 0x82, 0x4d, 0xd8,
	0x98, 0x0a, 0x34, 0x1a, 0x49, 0xe5, 0x8e, 0x75, 0x63, 0x5f, 0xd6, 0xd1, 0x02, 0x6d, 0xe7, 0x4b,
	0x63, 0x4f, 0x82, 0xba, 0xdc, 0x58, 0xe9, 0xd6, 0x63, 0x4f, 0xf1, 0x33, 0xd9, 0x86, 0xd5, 0xd8,
	0x1b, 0x0f, 0xe9, 0x18, 0x8d, 0x66, 0x62, 0x76, 0xb1, 0x60, 0x4e, 0x61, 0xfd, 0x28, 0xc2, 0x98,
	0x08, 0xbf, 0xce, 0x91, 0x0b, 0xf2, 0x06, 0x9a, 0x52, 0x2d, 0x69, 0x55, 0x1f, 0x3c, 0xab, 0xf0,
	0x21, 0x93, 0xb1, 0xd3, 0x82, 0xa2, 0x52, 0xad, 0xac, 0xf4, 0x4b, 0x03, 0x3d, 0x67, 0x9c, 0x6c,
	0xc1, 0xa3, 0x10, 0x67, 0x97, 0x0e, 0x15, 0x34, 0x4d, 0xb5, 0x15, 0xe2, 0xec, 0x98, 0x0a, 0x1a,
	0x13, 0x51, 0xcf, 0x0d, 0x22, 0x26, 0xa6, 0x33, 0x45, 0x94, 0x2d, 0x90, 0x2e, 0xe8, 0x13, 0xe6,
	0xbb, 0x18, 0x85, 0x11, 0xf3, 0x45, 0x12, 0xf4, 0x9a, 0x9d, 0x5f, 0x22, 0x03, 0xd8, 0xcc, 0xbd,
	0x1e, 0x64, 0x54, 0x8d, 0x84, 0xea, 0x9f, 0xdf, 0xcc, 0x33, 0x68, 0xab, 0x20, 0x78, 0x18, 0xf8,
	0x1c, 0x1f, 0x90, 0x84, 0xb9, 0x0f, 0x8f, 0x4f, 0x50, 0x1c, 0x5e, 0x0f, 0x17, 0x4a, 0x2a, 0xdf,
	0x92, 0x7b, 0x6d, 0xc9, 0xbd, 0xf9, 0x09, 0x36, 0x25, 0xdd, 0x30, 0x88, 0x72, 0x04, 0x77, 0x3a,
	0x80, 0x85, 0x2d, 0xa8, 0x97, 0xb7, 0x00, 0xc1, 0x58, 0xb6, 0x95, 0x76, 0xfb, 0x16, 0x5a, 0xd2,
	0x3c, 0x37, 0xb4, 0x6e, 0xbd, 0xa7, 0x0f, 0xfa, 0xb7, 0xb6, 0x5b, 0xf4, 0x67, 0xab, 0x7a, 0x73,
	0x1b, 0xe0, 0x04, 0xb3, 0x86, 0x4b, 0xb6, 0xcd, 0x53, 0xd0, 0x93, 0xaf, 0x0f, 0x4f, 0xf9, 0x1c,
	0xc8, 0x39, 0xe3, 0x42, 0xae, 0x72, 0xa5, 0x57, 0x88, 0x40, 0x2b, 0x45, 0x40, 0x0c, 0x68, 0xd1,
	0xf1, 0x38, 0x98, 0xfb, 0x22, 0xcd, 0x4d, 0xbd, 0x9a, 0x67, 0xb0, 0x16, 0xb3, 0x65, 0xc6, 0xf6,
	0xcb, 0x81, 0xdc, 0xc1, 0x59, 0x16, 0xc1, 0x53, 0x58, 0x3f, 0x46, 0x0f, 0x17, 0x63, 0x55, 0x4e,
	0x61, 0x03, 0xda, 0x0a, 0x20, 0xf5, 0xcc, 0x9f, 0x1a, 0xac, 0x7f, 0x08, 0x9d, 0xdc, 0x28, 0xbe,
	0x86, 0xd5, 0xe4, 0xc2, 0x7a, 0x47, 0xf9, 0x55, 0x9a, 0xce, 0xf2, 0x7d, 0x32, 0x54, 0x08, 0x7b,
	0x01, 0xce, 0x85, 0x5a, 0xbb, 0x6f, 0xa8, 0x1b, 0xd0, 0x56, 0x2e, 0xa4, 0xb1, 0xc1, 0xef, 0x15,
	0x68, 0xa5, 0x19, 0x93, 0x0b, 0x68, 0xca, 0x29, 0x21, 0xbd, 0xaa, 0xfb, 0x29, 0x7f, 0xa3, 0x74,
	0x5e, 0xde, 0x01, 0x99, 0x66, 0x7e, 0x01, 0x4d, 0x29, 0x5e, 0x49, 0x5f, 0x48, 0xa9, 0x92, 0xbe,
	0xd8, 0x09, 0xf9, 0x01, 0x1b, 0xe5, 0xf3, 0x4f, 0x06, 0x15, 0xe5, 0x37, 0xcc, 0x70, 0x67, 0xef,
	0x5e, 0x35, 0xa9, 0xf8, 0x7b, 0xa8, 0x9f, 0xa0, 0x20, 0xcf, 0xab, 0x6b, 0x95, 0xc4, 0xce, 0x6d,
	0xb0, 0x94, 0xf5, 0x0b, 0x34, 0xe2, 0x53, 0x4b, 0x76, 0x2b, 0xf0, 0xcb, 0x43, 0xd2, 0x79, 0x71,
	0x0b, 0x3c, 0xe3, 0x9f, 0x42, 0x3b, 0x7e, 0x1f, 0x06, 0xd1, 0x81, 0x9c, 0x93, 0xff, 0xa6, 0x74,
	0x11, 0xff, 0xd2, 0xc6, 0x13, 0x51, 0xb9, 0xf7, 0x85, 0xa9, 0xaa, 0xdc, 0xfb, 0xe2, 0x78, 0x1d,
	0xae, 0x7d, 0x86, 0xc5, 0x5f, 0x89, 0x51, 0x33, 0x99, 0x9f, 0xbd, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0x36, 0x07, 0x3a, 0x6c, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevicesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetByFingerprint(ctx context.Context, in *GetByFingerprintRequest, opts ...grpc.CallOption) (*GetByFingerprintResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListResponse, error)
	ListForAccount(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type devicesClient struct {
	cc *grpc.ClientConn
}

func NewDevicesClient(cc *grpc.ClientConn) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetByFingerprint(ctx context.Context, in *GetByFingerprintRequest, opts ...grpc.CallOption) (*GetByFingerprintResponse, error) {
	out := new(GetByFingerprintResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/GetByFingerprint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) ListForAccount(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/ListForAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.deviceregistry.Devices/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
type DevicesServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetByFingerprint(context.Context, *GetByFingerprintRequest) (*GetByFingerprintResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListDevicesRequest) (*ListResponse, error)
	ListForAccount(context.Context, *ListDevicesRequest) (*ListResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterDevicesServer(s *grpc.Server, srv DevicesServer) {
	s.RegisterService(&_Devices_serviceDesc, srv)
}

func _Devices_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetByFingerprint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByFingerprintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetByFingerprint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/GetByFingerprint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetByFingerprint(ctx, req.(*GetByFingerprintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).List(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_ListForAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).ListForAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/ListForAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).ListForAccount(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.deviceregistry.Devices/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Devices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.deviceregistry.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Devices_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Devices_Update_Handler,
		},
		{
			MethodName: "GetByFingerprint",
			Handler:    _Devices_GetByFingerprint_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Devices_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Devices_List_Handler,
		},
		{
			MethodName: "ListForAccount",
			Handler:    _Devices_ListForAccount_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Devices_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/registry/registrypb/device_registry.proto",
}
