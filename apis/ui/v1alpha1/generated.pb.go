/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go.openviz.dev/grafana-tools/apis/ui/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	v1 "kmodules.xyz/client-go/api/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *DashboardRef) Reset()      { *m = DashboardRef{} }
func (*DashboardRef) ProtoMessage() {}
func (*DashboardRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4e2fa1ec64fe07, []int{0}
}
func (m *DashboardRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DashboardRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DashboardRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardRef.Merge(m, src)
}
func (m *DashboardRef) XXX_Size() int {
	return m.Size()
}
func (m *DashboardRef) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardRef.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardRef proto.InternalMessageInfo

func (m *EmbeddedDashboard) Reset()      { *m = EmbeddedDashboard{} }
func (*EmbeddedDashboard) ProtoMessage() {}
func (*EmbeddedDashboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4e2fa1ec64fe07, []int{1}
}
func (m *EmbeddedDashboard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmbeddedDashboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EmbeddedDashboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbeddedDashboard.Merge(m, src)
}
func (m *EmbeddedDashboard) XXX_Size() int {
	return m.Size()
}
func (m *EmbeddedDashboard) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbeddedDashboard.DiscardUnknown(m)
}

var xxx_messageInfo_EmbeddedDashboard proto.InternalMessageInfo

func (m *EmbeddedDashboardRequest) Reset()      { *m = EmbeddedDashboardRequest{} }
func (*EmbeddedDashboardRequest) ProtoMessage() {}
func (*EmbeddedDashboardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4e2fa1ec64fe07, []int{2}
}
func (m *EmbeddedDashboardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmbeddedDashboardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EmbeddedDashboardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbeddedDashboardRequest.Merge(m, src)
}
func (m *EmbeddedDashboardRequest) XXX_Size() int {
	return m.Size()
}
func (m *EmbeddedDashboardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbeddedDashboardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmbeddedDashboardRequest proto.InternalMessageInfo

func (m *EmbeddedDashboardResponse) Reset()      { *m = EmbeddedDashboardResponse{} }
func (*EmbeddedDashboardResponse) ProtoMessage() {}
func (*EmbeddedDashboardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4e2fa1ec64fe07, []int{3}
}
func (m *EmbeddedDashboardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmbeddedDashboardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EmbeddedDashboardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbeddedDashboardResponse.Merge(m, src)
}
func (m *EmbeddedDashboardResponse) XXX_Size() int {
	return m.Size()
}
func (m *EmbeddedDashboardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbeddedDashboardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmbeddedDashboardResponse proto.InternalMessageInfo

func (m *PanelURL) Reset()      { *m = PanelURL{} }
func (*PanelURL) ProtoMessage() {}
func (*PanelURL) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4e2fa1ec64fe07, []int{4}
}
func (m *PanelURL) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PanelURL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PanelURL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PanelURL.Merge(m, src)
}
func (m *PanelURL) XXX_Size() int {
	return m.Size()
}
func (m *PanelURL) XXX_DiscardUnknown() {
	xxx_messageInfo_PanelURL.DiscardUnknown(m)
}

var xxx_messageInfo_PanelURL proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DashboardRef)(nil), "go.openviz.dev.grafana_tools.apis.ui.v1alpha1.DashboardRef")
	proto.RegisterType((*EmbeddedDashboard)(nil), "go.openviz.dev.grafana_tools.apis.ui.v1alpha1.EmbeddedDashboard")
	proto.RegisterType((*EmbeddedDashboardRequest)(nil), "go.openviz.dev.grafana_tools.apis.ui.v1alpha1.EmbeddedDashboardRequest")
	proto.RegisterType((*EmbeddedDashboardResponse)(nil), "go.openviz.dev.grafana_tools.apis.ui.v1alpha1.EmbeddedDashboardResponse")
	proto.RegisterType((*PanelURL)(nil), "go.openviz.dev.grafana_tools.apis.ui.v1alpha1.PanelURL")
}

func init() {
	proto.RegisterFile("go.openviz.dev/grafana-tools/apis/ui/v1alpha1/generated.proto", fileDescriptor_5e4e2fa1ec64fe07)
}

var fileDescriptor_5e4e2fa1ec64fe07 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x33, 0xf4, 0x2f, 0x19, 0x17, 0xaa, 0xba, 0x9b, 0xd0, 0x85, 0x13, 0x99, 0x05, 0xd9,
	0x64, 0xac, 0x56, 0xfc, 0x49, 0x88, 0x8d, 0xd5, 0x0a, 0x90, 0x2a, 0x15, 0x8d, 0xda, 0x05, 0x6c,
	0xaa, 0x89, 0x7d, 0xe3, 0x98, 0xda, 0x1e, 0x33, 0x33, 0xb6, 0x68, 0x57, 0x3c, 0x02, 0xcf, 0xc0,
	0xd3, 0x74, 0xd9, 0x65, 0x57, 0x11, 0x35, 0x0f, 0xc1, 0x0a, 0x09, 0xc5, 0x1e, 0x37, 0x6d, 0x43,
	0x04, 0x15, 0x3b, 0xdf, 0x7b, 0x7d, 0xce, 0xf9, 0x66, 0xe6, 0xe2, 0x57, 0x01, 0x27, 0x3c, 0x85,
	0x24, 0x0f, 0x4f, 0x89, 0x0f, 0xb9, 0x13, 0x08, 0x36, 0x64, 0x09, 0xeb, 0x2b, 0xce, 0x23, 0xe9,
	0xb0, 0x34, 0x94, 0x4e, 0x16, 0x3a, 0xf9, 0x16, 0x8b, 0xd2, 0x11, 0xdb, 0x72, 0x02, 0x48, 0x40,
	0x30, 0x05, 0x3e, 0x49, 0x05, 0x57, 0xdc, 0xec, 0xdf, 0x94, 0x13, 0x2d, 0x3f, 0x2a, 0xe5, 0x64,
	0x22, 0x27, 0x59, 0x48, 0x6a, 0xf9, 0x66, 0x3f, 0x08, 0xd5, 0x28, 0x1b, 0x10, 0x8f, 0xc7, 0x4e,
	0xc0, 0x03, 0xee, 0x94, 0x2e, 0x83, 0x6c, 0x58, 0x56, 0x65, 0x51, 0x7e, 0x55, 0xee, 0x9b, 0x4f,
	0x8e, 0x5f, 0x48, 0x12, 0xf2, 0x09, 0x46, 0xcc, 0xbc, 0x51, 0x98, 0x80, 0x38, 0x71, 0xd2, 0xe3,
	0xa0, 0xe2, 0x8a, 0x41, 0x31, 0x27, 0x9f, 0x61, 0xda, 0x74, 0xe6, 0xa9, 0x44, 0x96, 0xa8, 0x30,
	0x86, 0x19, 0xc1, 0xb3, 0xbf, 0x09, 0xa4, 0x37, 0x82, 0x98, 0xcd, 0xe8, 0xfa, 0xc7, 0x31, 0xf7,
	0xb3, 0x08, 0x24, 0xf9, 0x7c, 0x72, 0xea, 0x78, 0x51, 0x08, 0x89, 0xea, 0x07, 0xa5, 0xcf, 0x1f,
	0xb8, 0xec, 0x6f, 0x08, 0xaf, 0xee, 0x30, 0x39, 0x1a, 0x70, 0x26, 0x7c, 0x0a, 0x43, 0x33, 0xc6,
	0x6b, 0x7c, 0xf0, 0x11, 0x3c, 0x45, 0x61, 0x08, 0x02, 0x12, 0x0f, 0xda, 0xa8, 0x8b, 0x7a, 0xc6,
	0x36, 0x21, 0xd7, 0x9d, 0x49, 0xe5, 0x7c, 0x14, 0xf0, 0xc9, 0x85, 0x92, 0x7c, 0x8b, 0xec, 0xdf,
	0x54, 0xb9, 0x1b, 0xc5, 0xb8, 0xb3, 0x76, 0xab, 0x49, 0x6f, 0x7b, 0x9b, 0x8f, 0xf0, 0x92, 0x0a,
	0x55, 0x04, 0xed, 0x7b, 0x5d, 0xd4, 0x6b, 0xb9, 0xf7, 0xcf, 0xc6, 0x9d, 0x46, 0x31, 0xee, 0x2c,
	0x1d, 0x4c, 0x9a, 0xb4, 0x9a, 0xd9, 0x3f, 0x11, 0x5e, 0xdf, 0x8d, 0x07, 0xe0, 0xfb, 0xe0, 0x5f,
	0xc1, 0x9a, 0x09, 0x5e, 0x11, 0xf0, 0x29, 0x03, 0xa9, 0x34, 0xe1, 0x6b, 0x72, 0xa7, 0x87, 0x27,
	0x33, 0x96, 0xb4, 0xb2, 0x73, 0x8d, 0x62, 0xdc, 0x59, 0xd1, 0x05, 0xad, 0x43, 0x4c, 0x81, 0x9b,
	0x02, 0x64, 0xca, 0x13, 0x59, 0xd1, 0x1a, 0xdb, 0x6f, 0xfe, 0x3f, 0xb0, 0xf2, 0x73, 0x57, 0x8b,
	0x71, 0xa7, 0x59, 0x57, 0xf4, 0x2a, 0xc7, 0xfe, 0x85, 0x70, 0x7b, 0x1e, 0xa6, 0xb9, 0x8f, 0x97,
	0x15, 0x13, 0x01, 0xd4, 0xe7, 0x7f, 0xfc, 0x4f, 0x2f, 0xf4, 0x76, 0xc7, 0x7d, 0xa0, 0x6f, 0x79,
	0xf9, 0xa0, 0x94, 0x53, 0x6d, 0x63, 0x46, 0xb8, 0xe5, 0xd7, 0x21, 0xfa, 0x88, 0x2f, 0xef, 0x78,
	0xc4, 0xeb, 0xbb, 0xe4, 0xae, 0xeb, 0x9c, 0xd6, 0xb4, 0x3b, 0x0d, 0x30, 0x6d, 0xbc, 0x9c, 0xb2,
	0x04, 0x22, 0xd9, 0x5e, 0xe8, 0x2e, 0xf4, 0x5a, 0x2e, 0x9e, 0x10, 0xbd, 0x2b, 0x3b, 0x54, 0x4f,
	0xec, 0x1c, 0x3f, 0x9c, 0x7b, 0x69, 0xe6, 0x7b, 0xbc, 0x98, 0x89, 0x48, 0xb6, 0x51, 0x77, 0xa1,
	0x67, 0x6c, 0x3f, 0xbf, 0x23, 0x69, 0x99, 0x74, 0x48, 0xf7, 0xdc, 0x55, 0x4d, 0xb9, 0x78, 0x48,
	0xf7, 0x24, 0x2d, 0x2d, 0xed, 0x21, 0x6e, 0xd6, 0xf3, 0xe9, 0x8a, 0xa2, 0xf9, 0x2b, 0x6a, 0x3e,
	0xc5, 0x06, 0x68, 0xd0, 0x43, 0xba, 0xa7, 0xb7, 0x79, 0x43, 0xff, 0x6a, 0xec, 0x4e, 0x47, 0xf4,
	0xfa, 0x7f, 0x2e, 0x39, 0xbb, 0xb4, 0x1a, 0xe7, 0x97, 0x56, 0xe3, 0xe2, 0xd2, 0x6a, 0x7c, 0x29,
	0x2c, 0x74, 0x56, 0x58, 0xe8, 0xbc, 0xb0, 0xd0, 0x45, 0x61, 0xa1, 0xef, 0x85, 0x85, 0xbe, 0xfe,
	0xb0, 0x1a, 0x1f, 0x9a, 0x35, 0xf4, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x13, 0x49, 0x4d,
	0x1a, 0x05, 0x00, 0x00,
}

func (m *DashboardRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DashboardRef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DashboardRef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Title)
	copy(dAtA[i:], m.Title)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Title)))
	i--
	dAtA[i] = 0x12
	if m.ObjectReference != nil {
		{
			size, err := m.ObjectReference.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EmbeddedDashboard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmbeddedDashboard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmbeddedDashboard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Response != nil {
		{
			size, err := m.Response.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EmbeddedDashboardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmbeddedDashboardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmbeddedDashboardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Panels) > 0 {
		for iNdEx := len(m.Panels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Panels[iNdEx])
			copy(dAtA[i:], m.Panels[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Panels[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Dashboard.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Target.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EmbeddedDashboardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmbeddedDashboardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmbeddedDashboardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.URLs) > 0 {
		for iNdEx := len(m.URLs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.URLs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PanelURL) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PanelURL) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PanelURL) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.EmbeddedURL)
	copy(dAtA[i:], m.EmbeddedURL)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.EmbeddedURL)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Title)
	copy(dAtA[i:], m.Title)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Title)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DashboardRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ObjectReference != nil {
		l = m.ObjectReference.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Title)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *EmbeddedDashboard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *EmbeddedDashboardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Target.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Dashboard.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Panels) > 0 {
		for _, s := range m.Panels {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *EmbeddedDashboardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.URLs) > 0 {
		for _, e := range m.URLs {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *PanelURL) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.EmbeddedURL)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DashboardRef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DashboardRef{`,
		`ObjectReference:` + strings.Replace(fmt.Sprintf("%v", this.ObjectReference), "ObjectReference", "v1.ObjectReference", 1) + `,`,
		`Title:` + fmt.Sprintf("%v", this.Title) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EmbeddedDashboard) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EmbeddedDashboard{`,
		`Request:` + strings.Replace(this.Request.String(), "EmbeddedDashboardRequest", "EmbeddedDashboardRequest", 1) + `,`,
		`Response:` + strings.Replace(this.Response.String(), "EmbeddedDashboardResponse", "EmbeddedDashboardResponse", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EmbeddedDashboardRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EmbeddedDashboardRequest{`,
		`Target:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Target), "ObjectID", "v1.ObjectID", 1), `&`, ``, 1) + `,`,
		`Dashboard:` + strings.Replace(strings.Replace(this.Dashboard.String(), "DashboardRef", "DashboardRef", 1), `&`, ``, 1) + `,`,
		`Panels:` + fmt.Sprintf("%v", this.Panels) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EmbeddedDashboardResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForURLs := "[]PanelURL{"
	for _, f := range this.URLs {
		repeatedStringForURLs += strings.Replace(strings.Replace(f.String(), "PanelURL", "PanelURL", 1), `&`, ``, 1) + ","
	}
	repeatedStringForURLs += "}"
	s := strings.Join([]string{`&EmbeddedDashboardResponse{`,
		`URLs:` + repeatedStringForURLs + `,`,
		`}`,
	}, "")
	return s
}
func (this *PanelURL) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PanelURL{`,
		`Title:` + fmt.Sprintf("%v", this.Title) + `,`,
		`EmbeddedURL:` + fmt.Sprintf("%v", this.EmbeddedURL) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DashboardRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DashboardRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DashboardRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectReference", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ObjectReference == nil {
				m.ObjectReference = &v1.ObjectReference{}
			}
			if err := m.ObjectReference.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmbeddedDashboard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmbeddedDashboard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmbeddedDashboard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &EmbeddedDashboardRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &EmbeddedDashboardResponse{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmbeddedDashboardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmbeddedDashboardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmbeddedDashboardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Target.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dashboard", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Dashboard.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Panels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Panels = append(m.Panels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmbeddedDashboardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmbeddedDashboardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmbeddedDashboardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URLs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URLs = append(m.URLs, PanelURL{})
			if err := m.URLs[len(m.URLs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PanelURL) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PanelURL: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PanelURL: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmbeddedURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmbeddedURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
