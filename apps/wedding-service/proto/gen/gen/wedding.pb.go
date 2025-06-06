// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: wedding.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
	Gender_OTHER  Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
		2: "OTHER",
	}
	Gender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
		"OTHER":  2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_wedding_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_wedding_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_wedding_proto_rawDescGZIP(), []int{0}
}

type Person struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FirstName     string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender        Gender                 `protobuf:"varint,3,opt,name=gender,proto3,enum=wedding.Gender" json:"gender,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_wedding_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_wedding_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_wedding_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_MALE
}

type Couple struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Person_1      *Person                `protobuf:"bytes,1,opt,name=person_1,json=person1,proto3" json:"person_1,omitempty"`
	Person_2      *Person                `protobuf:"bytes,2,opt,name=person_2,json=person2,proto3" json:"person_2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Couple) Reset() {
	*x = Couple{}
	mi := &file_wedding_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Couple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Couple) ProtoMessage() {}

func (x *Couple) ProtoReflect() protoreflect.Message {
	mi := &file_wedding_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Couple.ProtoReflect.Descriptor instead.
func (*Couple) Descriptor() ([]byte, []int) {
	return file_wedding_proto_rawDescGZIP(), []int{1}
}

func (x *Couple) GetPerson_1() *Person {
	if x != nil {
		return x.Person_1
	}
	return nil
}

func (x *Couple) GetPerson_2() *Person {
	if x != nil {
		return x.Person_2
	}
	return nil
}

type Wedding struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Couple        *Couple                `protobuf:"bytes,2,opt,name=couple,proto3" json:"couple,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Location      string                 `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	GuestCount    int32                  `protobuf:"varint,5,opt,name=guest_count,json=guestCount,proto3" json:"guest_count,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Wedding) Reset() {
	*x = Wedding{}
	mi := &file_wedding_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Wedding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wedding) ProtoMessage() {}

func (x *Wedding) ProtoReflect() protoreflect.Message {
	mi := &file_wedding_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wedding.ProtoReflect.Descriptor instead.
func (*Wedding) Descriptor() ([]byte, []int) {
	return file_wedding_proto_rawDescGZIP(), []int{2}
}

func (x *Wedding) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Wedding) GetCouple() *Couple {
	if x != nil {
		return x.Couple
	}
	return nil
}

func (x *Wedding) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Wedding) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Wedding) GetGuestCount() int32 {
	if x != nil {
		return x.GuestCount
	}
	return 0
}

func (x *Wedding) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Wedding) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateWeddingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wedding       *Wedding               `protobuf:"bytes,1,opt,name=wedding,proto3" json:"wedding,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWeddingRequest) Reset() {
	*x = CreateWeddingRequest{}
	mi := &file_wedding_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWeddingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWeddingRequest) ProtoMessage() {}

func (x *CreateWeddingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wedding_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWeddingRequest.ProtoReflect.Descriptor instead.
func (*CreateWeddingRequest) Descriptor() ([]byte, []int) {
	return file_wedding_proto_rawDescGZIP(), []int{3}
}

func (x *CreateWeddingRequest) GetWedding() *Wedding {
	if x != nil {
		return x.Wedding
	}
	return nil
}

type CreateWeddingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wedding       *Wedding               `protobuf:"bytes,1,opt,name=wedding,proto3" json:"wedding,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWeddingResponse) Reset() {
	*x = CreateWeddingResponse{}
	mi := &file_wedding_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWeddingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWeddingResponse) ProtoMessage() {}

func (x *CreateWeddingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wedding_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWeddingResponse.ProtoReflect.Descriptor instead.
func (*CreateWeddingResponse) Descriptor() ([]byte, []int) {
	return file_wedding_proto_rawDescGZIP(), []int{4}
}

func (x *CreateWeddingResponse) GetWedding() *Wedding {
	if x != nil {
		return x.Wedding
	}
	return nil
}

var File_wedding_proto protoreflect.FileDescriptor

var file_wedding_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x06, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x70,
	0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x31, 0x12, 0x2a,
	0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x32, 0x22, 0xa5, 0x02, 0x0a, 0x07, 0x57,
	0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6c, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x67, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x77, 0x65,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x65,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x77,
	0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x43, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x07, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x65, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x07, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2a, 0x29, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f,
	0x54, 0x48, 0x45, 0x52, 0x10, 0x02, 0x32, 0x62, 0x0a, 0x0e, 0x57, 0x65, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_wedding_proto_rawDescOnce sync.Once
	file_wedding_proto_rawDescData []byte
)

func file_wedding_proto_rawDescGZIP() []byte {
	file_wedding_proto_rawDescOnce.Do(func() {
		file_wedding_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_wedding_proto_rawDesc), len(file_wedding_proto_rawDesc)))
	})
	return file_wedding_proto_rawDescData
}

var file_wedding_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wedding_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wedding_proto_goTypes = []any{
	(Gender)(0),                   // 0: wedding.Gender
	(*Person)(nil),                // 1: wedding.Person
	(*Couple)(nil),                // 2: wedding.Couple
	(*Wedding)(nil),               // 3: wedding.Wedding
	(*CreateWeddingRequest)(nil),  // 4: wedding.CreateWeddingRequest
	(*CreateWeddingResponse)(nil), // 5: wedding.CreateWeddingResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_wedding_proto_depIdxs = []int32{
	0,  // 0: wedding.Person.gender:type_name -> wedding.Gender
	1,  // 1: wedding.Couple.person_1:type_name -> wedding.Person
	1,  // 2: wedding.Couple.person_2:type_name -> wedding.Person
	2,  // 3: wedding.Wedding.couple:type_name -> wedding.Couple
	6,  // 4: wedding.Wedding.date:type_name -> google.protobuf.Timestamp
	6,  // 5: wedding.Wedding.created_at:type_name -> google.protobuf.Timestamp
	6,  // 6: wedding.Wedding.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 7: wedding.CreateWeddingRequest.wedding:type_name -> wedding.Wedding
	3,  // 8: wedding.CreateWeddingResponse.wedding:type_name -> wedding.Wedding
	4,  // 9: wedding.WeddingService.CreateWedding:input_type -> wedding.CreateWeddingRequest
	5,  // 10: wedding.WeddingService.CreateWedding:output_type -> wedding.CreateWeddingResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_wedding_proto_init() }
func file_wedding_proto_init() {
	if File_wedding_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wedding_proto_rawDesc), len(file_wedding_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wedding_proto_goTypes,
		DependencyIndexes: file_wedding_proto_depIdxs,
		EnumInfos:         file_wedding_proto_enumTypes,
		MessageInfos:      file_wedding_proto_msgTypes,
	}.Build()
	File_wedding_proto = out.File
	file_wedding_proto_goTypes = nil
	file_wedding_proto_depIdxs = nil
}
