// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: process.proto

package modelpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ppid        uint32         `protobuf:"varint,1,opt,name=ppid,proto3" json:"ppid,omitempty"`
	Thread      *ProcessThread `protobuf:"bytes,2,opt,name=thread,proto3" json:"thread,omitempty"`
	Title       string         `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	CommandLine string         `protobuf:"bytes,4,opt,name=command_line,json=commandLine,proto3" json:"command_line,omitempty"`
	Executable  string         `protobuf:"bytes,5,opt,name=executable,proto3" json:"executable,omitempty"`
	Argv        []string       `protobuf:"bytes,6,rep,name=argv,proto3" json:"argv,omitempty"`
	Pid         uint32         `protobuf:"varint,7,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{0}
}

func (x *Process) GetPpid() uint32 {
	if x != nil {
		return x.Ppid
	}
	return 0
}

func (x *Process) GetThread() *ProcessThread {
	if x != nil {
		return x.Thread
	}
	return nil
}

func (x *Process) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Process) GetCommandLine() string {
	if x != nil {
		return x.CommandLine
	}
	return ""
}

func (x *Process) GetExecutable() string {
	if x != nil {
		return x.Executable
	}
	return ""
}

func (x *Process) GetArgv() []string {
	if x != nil {
		return x.Argv
	}
	return nil
}

func (x *Process) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type ProcessThread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProcessThread) Reset() {
	*x = ProcessThread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessThread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessThread) ProtoMessage() {}

func (x *ProcessThread) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessThread.ProtoReflect.Descriptor instead.
func (*ProcessThread) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessThread) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProcessThread) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_process_proto protoreflect.FileDescriptor

var file_process_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x22,
	0xd3, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x70, 0x69, 0x64, 0x12,
	0x35, 0x0a, 0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x06,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x76, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2f, 0x61, 0x70, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_process_proto_rawDescOnce sync.Once
	file_process_proto_rawDescData = file_process_proto_rawDesc
)

func file_process_proto_rawDescGZIP() []byte {
	file_process_proto_rawDescOnce.Do(func() {
		file_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_process_proto_rawDescData)
	})
	return file_process_proto_rawDescData
}

var file_process_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_process_proto_goTypes = []interface{}{
	(*Process)(nil),       // 0: elastic.apm.v1.Process
	(*ProcessThread)(nil), // 1: elastic.apm.v1.ProcessThread
}
var file_process_proto_depIdxs = []int32{
	1, // 0: elastic.apm.v1.Process.thread:type_name -> elastic.apm.v1.ProcessThread
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_process_proto_init() }
func file_process_proto_init() {
	if File_process_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
		file_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessThread); i {
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
			RawDescriptor: file_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_process_proto_goTypes,
		DependencyIndexes: file_process_proto_depIdxs,
		MessageInfos:      file_process_proto_msgTypes,
	}.Build()
	File_process_proto = out.File
	file_process_proto_rawDesc = nil
	file_process_proto_goTypes = nil
	file_process_proto_depIdxs = nil
}
