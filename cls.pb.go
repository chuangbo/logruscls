// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cls.proto

/*
Package logruscls is a generated protocol buffer package.

It is generated from these files:
	cls.proto

It has these top-level messages:
	Log
	LogGroup
	LogGroupList
*/
package logruscls

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Log struct {
	Time             *int64         `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	Contents         []*Log_Content `protobuf:"bytes,2,rep,name=contents" json:"contents,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Log) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Log) GetContents() []*Log_Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Log_Content struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Log_Content) Reset()                    { *m = Log_Content{} }
func (m *Log_Content) String() string            { return proto.CompactTextString(m) }
func (*Log_Content) ProtoMessage()               {}
func (*Log_Content) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Log_Content) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Log_Content) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type LogGroup struct {
	Logs             []*Log  `protobuf:"bytes,1,rep,name=logs" json:"logs,omitempty"`
	ContextFlow      *string `protobuf:"bytes,2,opt,name=contextFlow" json:"contextFlow,omitempty"`
	Filename         *string `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	Source           *string `protobuf:"bytes,4,opt,name=source" json:"source,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogGroup) Reset()                    { *m = LogGroup{} }
func (m *LogGroup) String() string            { return proto.CompactTextString(m) }
func (*LogGroup) ProtoMessage()               {}
func (*LogGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogGroup) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *LogGroup) GetContextFlow() string {
	if m != nil && m.ContextFlow != nil {
		return *m.ContextFlow
	}
	return ""
}

func (m *LogGroup) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *LogGroup) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

type LogGroupList struct {
	LogGroupList     []*LogGroup `protobuf:"bytes,1,rep,name=logGroupList" json:"logGroupList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *LogGroupList) Reset()                    { *m = LogGroupList{} }
func (m *LogGroupList) String() string            { return proto.CompactTextString(m) }
func (*LogGroupList) ProtoMessage()               {}
func (*LogGroupList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogGroupList) GetLogGroupList() []*LogGroup {
	if m != nil {
		return m.LogGroupList
	}
	return nil
}

func init() {
	proto.RegisterType((*Log)(nil), "logruscls.Log")
	proto.RegisterType((*Log_Content)(nil), "logruscls.Log.Content")
	proto.RegisterType((*LogGroup)(nil), "logruscls.LogGroup")
	proto.RegisterType((*LogGroupList)(nil), "logruscls.LogGroupList")
}

func init() { proto.RegisterFile("cls.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0xe5, 0x38, 0x40, 0xb2, 0xad, 0x10, 0x5a, 0x50, 0x65, 0xf5, 0x14, 0xe5, 0x94, 0x53,
	0x24, 0x7a, 0xe1, 0x01, 0x90, 0xe8, 0x25, 0x27, 0xbf, 0x41, 0x15, 0x19, 0x2b, 0x62, 0x9b, 0xad,
	0x62, 0x87, 0x9f, 0x1b, 0x07, 0x1e, 0x1c, 0xb1, 0x0d, 0x51, 0x73, 0xdb, 0x99, 0x6f, 0xec, 0x19,
	0xc8, 0x5b, 0x0a, 0xf5, 0x69, 0xe0, 0xc8, 0x98, 0x13, 0xfb, 0x61, 0x0c, 0x2d, 0x85, 0xf2, 0x5b,
	0x81, 0x6e, 0xd8, 0x23, 0x42, 0x1a, 0xbb, 0xa3, 0x33, 0xaa, 0x48, 0x2a, 0x6d, 0xe5, 0xc6, 0x1d,
	0x64, 0x2d, 0xf7, 0xd1, 0xf5, 0x31, 0x98, 0xa4, 0xd0, 0xd5, 0x6a, 0xb7, 0xa9, 0xe7, 0x97, 0x75,
	0xc3, 0xbe, 0x7e, 0x3e, 0x63, 0x3b, 0xe7, 0xb6, 0x8f, 0x70, 0x33, 0x99, 0x78, 0x07, 0xfa, 0xcd,
	0x7d, 0xc9, 0x8f, 0xb9, 0xfd, 0x3b, 0xf1, 0x01, 0xae, 0xde, 0x0f, 0x34, 0x3a, 0x93, 0x88, 0x77,
	0x16, 0xe5, 0x8f, 0x82, 0xac, 0x61, 0xbf, 0x1f, 0x78, 0x3c, 0x61, 0x09, 0x29, 0xb1, 0x0f, 0x46,
	0x49, 0xdf, 0xed, 0xb2, 0xcf, 0x0a, 0xc3, 0x02, 0x56, 0xd2, 0xf7, 0x19, 0x5f, 0x88, 0x3f, 0x4c,
	0x52, 0xa8, 0x2a, 0xb7, 0x97, 0x16, 0x6e, 0x21, 0x7b, 0xed, 0xc8, 0xf5, 0x87, 0xa3, 0x33, 0x5a,
	0xf0, 0xac, 0x71, 0x03, 0xd7, 0x81, 0xc7, 0xa1, 0x75, 0x26, 0x15, 0x32, 0xa9, 0x72, 0x0f, 0xeb,
	0xff, 0x15, 0x4d, 0x17, 0x22, 0x3e, 0xc1, 0x9a, 0x2e, 0xf4, 0xb4, 0xe8, 0x7e, 0xb9, 0x48, 0xb0,
	0x5d, 0x04, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x32, 0xc8, 0x1f, 0x69, 0x01, 0x00, 0x00,
}