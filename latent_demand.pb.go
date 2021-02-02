// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: latent_demand.proto

package proto_latent_demand

import (
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

// @type: LatentDemand
type LatentDemand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAtContext  string                 `protobuf:"bytes,1,opt,name=_at_context,json=AtContext,proto3" json:"@context,omitempty"`
	XAtType     string                 `protobuf:"bytes,2,opt,name=_at_type,json=AtType,proto3" json:"@type,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Location    *Place                 `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Domain      *LatentDemandDomain    `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty"`
	Description string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	DemandFrom  string                 `protobuf:"bytes,8,opt,name=demand_from,json=demandFrom,proto3" json:"demand_from,omitempty"`
	Evidence    string                 `protobuf:"bytes,9,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Option      []byte                 `protobuf:"bytes,10,opt,name=option,proto3" json:"option,omitempty"` // option
}

func (x *LatentDemand) Reset() {
	*x = LatentDemand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latent_demand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatentDemand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatentDemand) ProtoMessage() {}

func (x *LatentDemand) ProtoReflect() protoreflect.Message {
	mi := &file_latent_demand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatentDemand.ProtoReflect.Descriptor instead.
func (*LatentDemand) Descriptor() ([]byte, []int) {
	return file_latent_demand_proto_rawDescGZIP(), []int{0}
}

func (x *LatentDemand) GetXAtContext() string {
	if x != nil {
		return x.XAtContext
	}
	return ""
}

func (x *LatentDemand) GetXAtType() string {
	if x != nil {
		return x.XAtType
	}
	return ""
}

func (x *LatentDemand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LatentDemand) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *LatentDemand) GetLocation() *Place {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *LatentDemand) GetDomain() *LatentDemandDomain {
	if x != nil {
		return x.Domain
	}
	return nil
}

func (x *LatentDemand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LatentDemand) GetDemandFrom() string {
	if x != nil {
		return x.DemandFrom
	}
	return ""
}

func (x *LatentDemand) GetEvidence() string {
	if x != nil {
		return x.Evidence
	}
	return ""
}

func (x *LatentDemand) GetOption() []byte {
	if x != nil {
		return x.Option
	}
	return nil
}

// @type: Place
type Place struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAtType   string         `protobuf:"bytes,1,opt,name=_at_type,json=AtType,proto3" json:"@type,omitempty"`
	Name      string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Latitude  float64        `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64        `protobuf:"fixed64,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Address   *PostalAddress `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Place) Reset() {
	*x = Place{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latent_demand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Place) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Place) ProtoMessage() {}

func (x *Place) ProtoReflect() protoreflect.Message {
	mi := &file_latent_demand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Place.ProtoReflect.Descriptor instead.
func (*Place) Descriptor() ([]byte, []int) {
	return file_latent_demand_proto_rawDescGZIP(), []int{1}
}

func (x *Place) GetXAtType() string {
	if x != nil {
		return x.XAtType
	}
	return ""
}

func (x *Place) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Place) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Place) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Place) GetAddress() *PostalAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

type PostalAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAtType         string `protobuf:"bytes,1,opt,name=_at_type,json=AtType,proto3" json:"@type,omitempty"`
	AddressLocality string `protobuf:"bytes,2,opt,name=address_locality,json=addressLocality,proto3" json:"address_locality,omitempty"`
	AddressRegion   string `protobuf:"bytes,3,opt,name=address_region,json=addressRegion,proto3" json:"address_region,omitempty"`
	StreetAddress   string `protobuf:"bytes,4,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
}

func (x *PostalAddress) Reset() {
	*x = PostalAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latent_demand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostalAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostalAddress) ProtoMessage() {}

func (x *PostalAddress) ProtoReflect() protoreflect.Message {
	mi := &file_latent_demand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostalAddress.ProtoReflect.Descriptor instead.
func (*PostalAddress) Descriptor() ([]byte, []int) {
	return file_latent_demand_proto_rawDescGZIP(), []int{2}
}

func (x *PostalAddress) GetXAtType() string {
	if x != nil {
		return x.XAtType
	}
	return ""
}

func (x *PostalAddress) GetAddressLocality() string {
	if x != nil {
		return x.AddressLocality
	}
	return ""
}

func (x *PostalAddress) GetAddressRegion() string {
	if x != nil {
		return x.AddressRegion
	}
	return ""
}

func (x *PostalAddress) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

type LatentDemandDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAtType  string `protobuf:"bytes,1,opt,name=_at_type,json=AtType,proto3" json:"@type,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"` // ほんとに string ?
}

func (x *LatentDemandDomain) Reset() {
	*x = LatentDemandDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latent_demand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatentDemandDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatentDemandDomain) ProtoMessage() {}

func (x *LatentDemandDomain) ProtoReflect() protoreflect.Message {
	mi := &file_latent_demand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatentDemandDomain.ProtoReflect.Descriptor instead.
func (*LatentDemandDomain) Descriptor() ([]byte, []int) {
	return file_latent_demand_proto_rawDescGZIP(), []int{3}
}

func (x *LatentDemandDomain) GetXAtType() string {
	if x != nil {
		return x.XAtType
	}
	return ""
}

func (x *LatentDemandDomain) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type LatentDemandSupply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAtContext  string `protobuf:"bytes,1,opt,name=_at_context,json=AtContext,proto3" json:"@context,omitempty"`
	XAtType     string `protobuf:"bytes,2,opt,name=_at_type,json=AtType,proto3" json:"@type,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Option      []byte `protobuf:"bytes,6,opt,name=option,proto3" json:"option,omitempty"` // option
}

func (x *LatentDemandSupply) Reset() {
	*x = LatentDemandSupply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latent_demand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatentDemandSupply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatentDemandSupply) ProtoMessage() {}

func (x *LatentDemandSupply) ProtoReflect() protoreflect.Message {
	mi := &file_latent_demand_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatentDemandSupply.ProtoReflect.Descriptor instead.
func (*LatentDemandSupply) Descriptor() ([]byte, []int) {
	return file_latent_demand_proto_rawDescGZIP(), []int{4}
}

func (x *LatentDemandSupply) GetXAtContext() string {
	if x != nil {
		return x.XAtContext
	}
	return ""
}

func (x *LatentDemandSupply) GetXAtType() string {
	if x != nil {
		return x.XAtType
	}
	return ""
}

func (x *LatentDemandSupply) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LatentDemandSupply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LatentDemandSupply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LatentDemandSupply) GetOption() []byte {
	if x != nil {
		return x.Option
	}
	return nil
}

type LatentDemandDisplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAtContext string                 `protobuf:"bytes,1,opt,name=_at_context,json=AtContext,proto3" json:"@context,omitempty"`
	XAtType    string                 `protobuf:"bytes,2,opt,name=_at_type,json=AtType,proto3" json:"@type,omitempty"`
	DemandName string                 `protobuf:"bytes,3,opt,name=demand_name,json=demandName,proto3" json:"demand_name,omitempty"`
	SupplyName string                 `protobuf:"bytes,4,opt,name=supply_name,json=supplyName,proto3" json:"supply_name,omitempty"`
	Date       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	ExpDate    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=exp_date,json=expDate,proto3" json:"exp_date,omitempty"`
	Url        string                 `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Option     []byte                 `protobuf:"bytes,8,opt,name=option,proto3" json:"option,omitempty"` // option
	Id         uint64                 `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LatentDemandDisplay) Reset() {
	*x = LatentDemandDisplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latent_demand_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatentDemandDisplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatentDemandDisplay) ProtoMessage() {}

func (x *LatentDemandDisplay) ProtoReflect() protoreflect.Message {
	mi := &file_latent_demand_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatentDemandDisplay.ProtoReflect.Descriptor instead.
func (*LatentDemandDisplay) Descriptor() ([]byte, []int) {
	return file_latent_demand_proto_rawDescGZIP(), []int{5}
}

func (x *LatentDemandDisplay) GetXAtContext() string {
	if x != nil {
		return x.XAtContext
	}
	return ""
}

func (x *LatentDemandDisplay) GetXAtType() string {
	if x != nil {
		return x.XAtType
	}
	return ""
}

func (x *LatentDemandDisplay) GetDemandName() string {
	if x != nil {
		return x.DemandName
	}
	return ""
}

func (x *LatentDemandDisplay) GetSupplyName() string {
	if x != nil {
		return x.SupplyName
	}
	return ""
}

func (x *LatentDemandDisplay) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *LatentDemandDisplay) GetExpDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpDate
	}
	return nil
}

func (x *LatentDemandDisplay) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LatentDemandDisplay) GetOption() []byte {
	if x != nil {
		return x.Option
	}
	return nil
}

func (x *LatentDemandDisplay) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LatentDemandSupplyURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAtContext string `protobuf:"bytes,1,opt,name=_at_context,json=AtContext,proto3" json:"@context,omitempty"`
	XAtType    string `protobuf:"bytes,2,opt,name=_at_type,json=AtType,proto3" json:"@type,omitempty"`
	Url        string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Option     []byte `protobuf:"bytes,4,opt,name=option,proto3" json:"option,omitempty"` // option
}

func (x *LatentDemandSupplyURL) Reset() {
	*x = LatentDemandSupplyURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_latent_demand_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatentDemandSupplyURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatentDemandSupplyURL) ProtoMessage() {}

func (x *LatentDemandSupplyURL) ProtoReflect() protoreflect.Message {
	mi := &file_latent_demand_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatentDemandSupplyURL.ProtoReflect.Descriptor instead.
func (*LatentDemandSupplyURL) Descriptor() ([]byte, []int) {
	return file_latent_demand_proto_rawDescGZIP(), []int{6}
}

func (x *LatentDemandSupplyURL) GetXAtContext() string {
	if x != nil {
		return x.XAtContext
	}
	return ""
}

func (x *LatentDemandSupplyURL) GetXAtType() string {
	if x != nil {
		return x.XAtType
	}
	return ""
}

func (x *LatentDemandSupplyURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LatentDemandSupplyURL) GetOption() []byte {
	if x != nil {
		return x.Option
	}
	return nil
}

var File_latent_demand_proto protoreflect.FileDescriptor

var file_latent_demand_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x02, 0x0a, 0x0c, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x5f, 0x61, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x08, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x08, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x61,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x08, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4b, 0x0a, 0x12, 0x4c, 0x61, 0x74, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a,
	0x08, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x12, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x5f,
	0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x08, 0x5f,
	0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb2, 0x02, 0x0a, 0x13, 0x4c, 0x61, 0x74, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x1e, 0x0a, 0x0b, 0x5f, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x18, 0x0a, 0x08, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65,
	0x78, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x15,
	0x4c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x1e, 0x0a, 0x0b, 0x5f, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x08, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x65, 0x78, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_latent_demand_proto_rawDescOnce sync.Once
	file_latent_demand_proto_rawDescData = file_latent_demand_proto_rawDesc
)

func file_latent_demand_proto_rawDescGZIP() []byte {
	file_latent_demand_proto_rawDescOnce.Do(func() {
		file_latent_demand_proto_rawDescData = protoimpl.X.CompressGZIP(file_latent_demand_proto_rawDescData)
	})
	return file_latent_demand_proto_rawDescData
}

var file_latent_demand_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_latent_demand_proto_goTypes = []interface{}{
	(*LatentDemand)(nil),          // 0: latent_demand.LatentDemand
	(*Place)(nil),                 // 1: latent_demand.Place
	(*PostalAddress)(nil),         // 2: latent_demand.PostalAddress
	(*LatentDemandDomain)(nil),    // 3: latent_demand.LatentDemandDomain
	(*LatentDemandSupply)(nil),    // 4: latent_demand.LatentDemandSupply
	(*LatentDemandDisplay)(nil),   // 5: latent_demand.LatentDemandDisplay
	(*LatentDemandSupplyURL)(nil), // 6: latent_demand.LatentDemandSupplyURL
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_latent_demand_proto_depIdxs = []int32{
	7, // 0: latent_demand.LatentDemand.date:type_name -> google.protobuf.Timestamp
	1, // 1: latent_demand.LatentDemand.location:type_name -> latent_demand.Place
	3, // 2: latent_demand.LatentDemand.domain:type_name -> latent_demand.LatentDemandDomain
	2, // 3: latent_demand.Place.address:type_name -> latent_demand.PostalAddress
	7, // 4: latent_demand.LatentDemandDisplay.date:type_name -> google.protobuf.Timestamp
	7, // 5: latent_demand.LatentDemandDisplay.exp_date:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_latent_demand_proto_init() }
func file_latent_demand_proto_init() {
	if File_latent_demand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_latent_demand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatentDemand); i {
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
		file_latent_demand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Place); i {
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
		file_latent_demand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostalAddress); i {
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
		file_latent_demand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatentDemandDomain); i {
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
		file_latent_demand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatentDemandSupply); i {
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
		file_latent_demand_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatentDemandDisplay); i {
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
		file_latent_demand_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatentDemandSupplyURL); i {
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
			RawDescriptor: file_latent_demand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_latent_demand_proto_goTypes,
		DependencyIndexes: file_latent_demand_proto_depIdxs,
		MessageInfos:      file_latent_demand_proto_msgTypes,
	}.Build()
	File_latent_demand_proto = out.File
	file_latent_demand_proto_rawDesc = nil
	file_latent_demand_proto_goTypes = nil
	file_latent_demand_proto_depIdxs = nil
}
