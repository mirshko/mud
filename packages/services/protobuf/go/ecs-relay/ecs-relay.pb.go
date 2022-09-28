// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/ecs-relay.proto

package ecs_relay

import (
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

// Identifies a client connecting to Relay Service.
type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{0}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id        string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity     *Identity     `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Subscription *Subscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{2}
}

func (x *SubscriptionRequest) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *SubscriptionRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{3}
}

func (x *Subscription) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *Identity  `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Label    string     `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Messages []*Message `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{4}
}

func (x *PushRequest) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *PushRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PushRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{5}
}

type CountIdentitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountIdentitiesRequest) Reset() {
	*x = CountIdentitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountIdentitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountIdentitiesRequest) ProtoMessage() {}

func (x *CountIdentitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountIdentitiesRequest.ProtoReflect.Descriptor instead.
func (*CountIdentitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{6}
}

type CountIdentitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountIdentitiesResponse) Reset() {
	*x = CountIdentitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_relay_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountIdentitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountIdentitiesResponse) ProtoMessage() {}

func (x *CountIdentitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_relay_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountIdentitiesResponse.ProtoReflect.Descriptor instead.
func (*CountIdentitiesResponse) Descriptor() ([]byte, []int) {
	return file_proto_ecs_relay_proto_rawDescGZIP(), []int{7}
}

func (x *CountIdentitiesResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_ecs_relay_proto protoreflect.FileDescriptor

var file_proto_ecs_relay_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x73, 0x2d, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x22, 0x1e, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x65, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x0c,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x63, 0x73,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x2f, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xe7, 0x04, 0x0a, 0x0f, 0x45, 0x43, 0x53, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x65, 0x63, 0x73,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x12, 0x2e, 0x65, 0x63, 0x73,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x12,
	0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x65,
	0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x1a, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x65,
	0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x63, 0x73, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x1d, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x4f, 0x70,
	0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x65,
	0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x15, 0x2e, 0x65, 0x63,
	0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x63, 0x73, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x63, 0x73, 0x2d,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ecs_relay_proto_rawDescOnce sync.Once
	file_proto_ecs_relay_proto_rawDescData = file_proto_ecs_relay_proto_rawDesc
)

func file_proto_ecs_relay_proto_rawDescGZIP() []byte {
	file_proto_ecs_relay_proto_rawDescOnce.Do(func() {
		file_proto_ecs_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ecs_relay_proto_rawDescData)
	})
	return file_proto_ecs_relay_proto_rawDescData
}

var file_proto_ecs_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_ecs_relay_proto_goTypes = []interface{}{
	(*Identity)(nil),                // 0: ecsrelay.Identity
	(*Message)(nil),                 // 1: ecsrelay.Message
	(*SubscriptionRequest)(nil),     // 2: ecsrelay.SubscriptionRequest
	(*Subscription)(nil),            // 3: ecsrelay.Subscription
	(*PushRequest)(nil),             // 4: ecsrelay.PushRequest
	(*PushResponse)(nil),            // 5: ecsrelay.PushResponse
	(*CountIdentitiesRequest)(nil),  // 6: ecsrelay.CountIdentitiesRequest
	(*CountIdentitiesResponse)(nil), // 7: ecsrelay.CountIdentitiesResponse
}
var file_proto_ecs_relay_proto_depIdxs = []int32{
	0,  // 0: ecsrelay.SubscriptionRequest.identity:type_name -> ecsrelay.Identity
	3,  // 1: ecsrelay.SubscriptionRequest.subscription:type_name -> ecsrelay.Subscription
	0,  // 2: ecsrelay.PushRequest.identity:type_name -> ecsrelay.Identity
	1,  // 3: ecsrelay.PushRequest.messages:type_name -> ecsrelay.Message
	0,  // 4: ecsrelay.ECSRelayService.Authenticate:input_type -> ecsrelay.Identity
	0,  // 5: ecsrelay.ECSRelayService.Revoke:input_type -> ecsrelay.Identity
	0,  // 6: ecsrelay.ECSRelayService.Ping:input_type -> ecsrelay.Identity
	6,  // 7: ecsrelay.ECSRelayService.CountAuthenticated:input_type -> ecsrelay.CountIdentitiesRequest
	6,  // 8: ecsrelay.ECSRelayService.CountConnected:input_type -> ecsrelay.CountIdentitiesRequest
	2,  // 9: ecsrelay.ECSRelayService.Subscribe:input_type -> ecsrelay.SubscriptionRequest
	2,  // 10: ecsrelay.ECSRelayService.Unsubscribe:input_type -> ecsrelay.SubscriptionRequest
	0,  // 11: ecsrelay.ECSRelayService.OpenStream:input_type -> ecsrelay.Identity
	4,  // 12: ecsrelay.ECSRelayService.Push:input_type -> ecsrelay.PushRequest
	0,  // 13: ecsrelay.ECSRelayService.Authenticate:output_type -> ecsrelay.Identity
	0,  // 14: ecsrelay.ECSRelayService.Revoke:output_type -> ecsrelay.Identity
	0,  // 15: ecsrelay.ECSRelayService.Ping:output_type -> ecsrelay.Identity
	7,  // 16: ecsrelay.ECSRelayService.CountAuthenticated:output_type -> ecsrelay.CountIdentitiesResponse
	7,  // 17: ecsrelay.ECSRelayService.CountConnected:output_type -> ecsrelay.CountIdentitiesResponse
	3,  // 18: ecsrelay.ECSRelayService.Subscribe:output_type -> ecsrelay.Subscription
	3,  // 19: ecsrelay.ECSRelayService.Unsubscribe:output_type -> ecsrelay.Subscription
	1,  // 20: ecsrelay.ECSRelayService.OpenStream:output_type -> ecsrelay.Message
	5,  // 21: ecsrelay.ECSRelayService.Push:output_type -> ecsrelay.PushResponse
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ecs_relay_proto_init() }
func file_proto_ecs_relay_proto_init() {
	if File_proto_ecs_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ecs_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_proto_ecs_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_ecs_relay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRequest); i {
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
		file_proto_ecs_relay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_proto_ecs_relay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_proto_ecs_relay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_proto_ecs_relay_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountIdentitiesRequest); i {
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
		file_proto_ecs_relay_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountIdentitiesResponse); i {
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
			RawDescriptor: file_proto_ecs_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ecs_relay_proto_goTypes,
		DependencyIndexes: file_proto_ecs_relay_proto_depIdxs,
		MessageInfos:      file_proto_ecs_relay_proto_msgTypes,
	}.Build()
	File_proto_ecs_relay_proto = out.File
	file_proto_ecs_relay_proto_rawDesc = nil
	file_proto_ecs_relay_proto_goTypes = nil
	file_proto_ecs_relay_proto_depIdxs = nil
}
