// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.6
// source: greeting.proto

package greeting

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The reply message containing the greetings
type GreetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GreetReply) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GreetReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ShouldGreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShouldGreetRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

type ShouldGreetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShouldGreet bool `protobuf:"varint,1,opt,name=should_greet,json=shouldGreet,proto3" json:"should_greet,omitempty"`
}

func (x *ShouldGreetReply) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *ShouldGreetReply) GetShouldGreet() bool {
	if x != nil {
		return x.ShouldGreet
	}
	return false
}

// The greeting service definition.
// go:plugin type=plugin version=1
type Greeter interface {
	// Sends a greeting
	Greet(context.Context, GreetRequest) (GreetReply, error)
}

// go:plugin type=host
type HostFunctions interface {
	ShouldGreet(context.Context, ShouldGreetRequest) (ShouldGreetReply, error)
}
