// Code generated by protoc-gen-gogo.
// source: github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto
// DO NOT EDIT!

/*
	Package networkserver is a generated protocol buffer package.

	It is generated from these files:
		github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto

	It has these top-level messages:
		DevicesRequest
		DevicesResponse
		StatusRequest
		Status
*/
package networkserver

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/TheThingsNetwork/ttn/api"
import lorawan "github.com/TheThingsNetwork/ttn/api/protocol/lorawan"
import broker "github.com/TheThingsNetwork/ttn/api/broker"
import handler "github.com/TheThingsNetwork/ttn/api/handler"

import github_com_TheThingsNetwork_ttn_core_types "github.com/TheThingsNetwork/ttn/core/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DevicesRequest struct {
	// Device address from the uplink message
	DevAddr *github_com_TheThingsNetwork_ttn_core_types.DevAddr `protobuf:"bytes,1,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevAddr" json:"dev_addr,omitempty"`
	// Frame counter from the uplink message
	FCnt uint32 `protobuf:"varint,2,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
}

func (m *DevicesRequest) Reset()                    { *m = DevicesRequest{} }
func (m *DevicesRequest) String() string            { return proto.CompactTextString(m) }
func (*DevicesRequest) ProtoMessage()               {}
func (*DevicesRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{0} }

func (m *DevicesRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

type DevicesResponse struct {
	Results []*lorawan.Device `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *DevicesResponse) Reset()                    { *m = DevicesResponse{} }
func (m *DevicesResponse) String() string            { return proto.CompactTextString(m) }
func (*DevicesResponse) ProtoMessage()               {}
func (*DevicesResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{1} }

func (m *DevicesResponse) GetResults() []*lorawan.Device {
	if m != nil {
		return m.Results
	}
	return nil
}

// message StatusRequest is used to request the status of this NetworkServer
type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{2} }

// message Status is the response to the StatusRequest
type Status struct {
	System            *api.SystemStats    `protobuf:"bytes,1,opt,name=system" json:"system,omitempty"`
	Component         *api.ComponentStats `protobuf:"bytes,2,opt,name=component" json:"component,omitempty"`
	Uplink            *api.Rates          `protobuf:"bytes,11,opt,name=uplink" json:"uplink,omitempty"`
	Downlink          *api.Rates          `protobuf:"bytes,12,opt,name=downlink" json:"downlink,omitempty"`
	Activations       *api.Rates          `protobuf:"bytes,13,opt,name=activations" json:"activations,omitempty"`
	DevicesPerAddress *api.Percentiles    `protobuf:"bytes,21,opt,name=devices_per_address,json=devicesPerAddress" json:"devices_per_address,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptorNetworkserver, []int{3} }

func (m *Status) GetSystem() *api.SystemStats {
	if m != nil {
		return m.System
	}
	return nil
}

func (m *Status) GetComponent() *api.ComponentStats {
	if m != nil {
		return m.Component
	}
	return nil
}

func (m *Status) GetUplink() *api.Rates {
	if m != nil {
		return m.Uplink
	}
	return nil
}

func (m *Status) GetDownlink() *api.Rates {
	if m != nil {
		return m.Downlink
	}
	return nil
}

func (m *Status) GetActivations() *api.Rates {
	if m != nil {
		return m.Activations
	}
	return nil
}

func (m *Status) GetDevicesPerAddress() *api.Percentiles {
	if m != nil {
		return m.DevicesPerAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*DevicesRequest)(nil), "networkserver.DevicesRequest")
	proto.RegisterType((*DevicesResponse)(nil), "networkserver.DevicesResponse")
	proto.RegisterType((*StatusRequest)(nil), "networkserver.StatusRequest")
	proto.RegisterType((*Status)(nil), "networkserver.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkServer service

type NetworkServerClient interface {
	// Broker requests devices with DevAddr and matching FCnt (or disabled FCnt check)
	GetDevices(ctx context.Context, in *DevicesRequest, opts ...grpc.CallOption) (*DevicesResponse, error)
	// Broker requests device activation "template" from Network Server
	PrepareActivation(ctx context.Context, in *broker.DeduplicatedDeviceActivationRequest, opts ...grpc.CallOption) (*broker.DeduplicatedDeviceActivationRequest, error)
	// Broker confirms device activation (after response from Handler)
	Activate(ctx context.Context, in *handler.DeviceActivationResponse, opts ...grpc.CallOption) (*handler.DeviceActivationResponse, error)
	// Broker informs Network Server about Uplink
	Uplink(ctx context.Context, in *broker.DeduplicatedUplinkMessage, opts ...grpc.CallOption) (*broker.DeduplicatedUplinkMessage, error)
	// Broker informs Network Server about Downlink, NetworkServer may add MAC commands and re-set MIC
	Downlink(ctx context.Context, in *broker.DownlinkMessage, opts ...grpc.CallOption) (*broker.DownlinkMessage, error)
}

type networkServerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerClient(cc *grpc.ClientConn) NetworkServerClient {
	return &networkServerClient{cc}
}

func (c *networkServerClient) GetDevices(ctx context.Context, in *DevicesRequest, opts ...grpc.CallOption) (*DevicesResponse, error) {
	out := new(DevicesResponse)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/GetDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) PrepareActivation(ctx context.Context, in *broker.DeduplicatedDeviceActivationRequest, opts ...grpc.CallOption) (*broker.DeduplicatedDeviceActivationRequest, error) {
	out := new(broker.DeduplicatedDeviceActivationRequest)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/PrepareActivation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Activate(ctx context.Context, in *handler.DeviceActivationResponse, opts ...grpc.CallOption) (*handler.DeviceActivationResponse, error) {
	out := new(handler.DeviceActivationResponse)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/Activate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Uplink(ctx context.Context, in *broker.DeduplicatedUplinkMessage, opts ...grpc.CallOption) (*broker.DeduplicatedUplinkMessage, error) {
	out := new(broker.DeduplicatedUplinkMessage)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/Uplink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Downlink(ctx context.Context, in *broker.DownlinkMessage, opts ...grpc.CallOption) (*broker.DownlinkMessage, error) {
	out := new(broker.DownlinkMessage)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServer/Downlink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServer service

type NetworkServerServer interface {
	// Broker requests devices with DevAddr and matching FCnt (or disabled FCnt check)
	GetDevices(context.Context, *DevicesRequest) (*DevicesResponse, error)
	// Broker requests device activation "template" from Network Server
	PrepareActivation(context.Context, *broker.DeduplicatedDeviceActivationRequest) (*broker.DeduplicatedDeviceActivationRequest, error)
	// Broker confirms device activation (after response from Handler)
	Activate(context.Context, *handler.DeviceActivationResponse) (*handler.DeviceActivationResponse, error)
	// Broker informs Network Server about Uplink
	Uplink(context.Context, *broker.DeduplicatedUplinkMessage) (*broker.DeduplicatedUplinkMessage, error)
	// Broker informs Network Server about Downlink, NetworkServer may add MAC commands and re-set MIC
	Downlink(context.Context, *broker.DownlinkMessage) (*broker.DownlinkMessage, error)
}

func RegisterNetworkServerServer(s *grpc.Server, srv NetworkServerServer) {
	s.RegisterService(&_NetworkServer_serviceDesc, srv)
}

func _NetworkServer_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/GetDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).GetDevices(ctx, req.(*DevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_PrepareActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(broker.DeduplicatedDeviceActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).PrepareActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/PrepareActivation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).PrepareActivation(ctx, req.(*broker.DeduplicatedDeviceActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(handler.DeviceActivationResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Activate(ctx, req.(*handler.DeviceActivationResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Uplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(broker.DeduplicatedUplinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Uplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/Uplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Uplink(ctx, req.(*broker.DeduplicatedUplinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Downlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(broker.DownlinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Downlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServer/Downlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Downlink(ctx, req.(*broker.DownlinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "networkserver.NetworkServer",
	HandlerType: (*NetworkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDevices",
			Handler:    _NetworkServer_GetDevices_Handler,
		},
		{
			MethodName: "PrepareActivation",
			Handler:    _NetworkServer_PrepareActivation_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _NetworkServer_Activate_Handler,
		},
		{
			MethodName: "Uplink",
			Handler:    _NetworkServer_Uplink_Handler,
		},
		{
			MethodName: "Downlink",
			Handler:    _NetworkServer_Downlink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto",
}

// Client API for NetworkServerManager service

type NetworkServerManagerClient interface {
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Status, error)
}

type networkServerManagerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerManagerClient(cc *grpc.ClientConn) NetworkServerManagerClient {
	return &networkServerManagerClient{cc}
}

func (c *networkServerManagerClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/networkserver.NetworkServerManager/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServerManager service

type NetworkServerManagerServer interface {
	GetStatus(context.Context, *StatusRequest) (*Status, error)
}

func RegisterNetworkServerManagerServer(s *grpc.Server, srv NetworkServerManagerServer) {
	s.RegisterService(&_NetworkServerManager_serviceDesc, srv)
}

func _NetworkServerManager_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerManagerServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkserver.NetworkServerManager/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerManagerServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "networkserver.NetworkServerManager",
	HandlerType: (*NetworkServerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _NetworkServerManager_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto",
}

func (m *DevicesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DevAddr != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.DevAddr.Size()))
		n1, err := m.DevAddr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FCnt != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.FCnt))
	}
	return i, nil
}

func (m *DevicesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			dAtA[i] = 0xa
			i++
			i = encodeVarintNetworkserver(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.System != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.System.Size()))
		n2, err := m.System.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Component != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.Component.Size()))
		n3, err := m.Component.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Uplink != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.Uplink.Size()))
		n4, err := m.Uplink.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Downlink != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.Downlink.Size()))
		n5, err := m.Downlink.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Activations != nil {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.Activations.Size()))
		n6, err := m.Activations.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.DevicesPerAddress != nil {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintNetworkserver(dAtA, i, uint64(m.DevicesPerAddress.Size()))
		n7, err := m.DevicesPerAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeFixed64Networkserver(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Networkserver(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintNetworkserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DevicesRequest) Size() (n int) {
	var l int
	_ = l
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.FCnt != 0 {
		n += 1 + sovNetworkserver(uint64(m.FCnt))
	}
	return n
}

func (m *DevicesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovNetworkserver(uint64(l))
		}
	}
	return n
}

func (m *StatusRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Status) Size() (n int) {
	var l int
	_ = l
	if m.System != nil {
		l = m.System.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Component != nil {
		l = m.Component.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Uplink != nil {
		l = m.Uplink.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Downlink != nil {
		l = m.Downlink.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.Activations != nil {
		l = m.Activations.Size()
		n += 1 + l + sovNetworkserver(uint64(l))
	}
	if m.DevicesPerAddress != nil {
		l = m.DevicesPerAddress.Size()
		n += 2 + l + sovNetworkserver(uint64(l))
	}
	return n
}

func sovNetworkserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetworkserver(x uint64) (n int) {
	return sovNetworkserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DevicesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCnt", wireType)
			}
			m.FCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FCnt |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func (m *DevicesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &lorawan.Device{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func (m *StatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field System", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.System == nil {
				m.System = &api.SystemStats{}
			}
			if err := m.System.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Component == nil {
				m.Component = &api.ComponentStats{}
			}
			if err := m.Component.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uplink", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Uplink == nil {
				m.Uplink = &api.Rates{}
			}
			if err := m.Uplink.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downlink", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Downlink == nil {
				m.Downlink = &api.Rates{}
			}
			if err := m.Downlink.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Activations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Activations == nil {
				m.Activations = &api.Rates{}
			}
			if err := m.Activations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevicesPerAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DevicesPerAddress == nil {
				m.DevicesPerAddress = &api.Percentiles{}
			}
			if err := m.DevicesPerAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkserver
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
func skipNetworkserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkserver
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
					return 0, ErrIntOverflowNetworkserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkserver
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNetworkserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetworkserver
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNetworkserver(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNetworkserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/networkserver/networkserver.proto", fileDescriptorNetworkserver)
}

var fileDescriptorNetworkserver = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0x5d, 0x4f, 0x13, 0x4f,
	0x14, 0xc6, 0xb3, 0xf0, 0xff, 0x97, 0x72, 0x4a, 0x45, 0x06, 0x89, 0x4d, 0x95, 0x0a, 0x4d, 0x34,
	0x35, 0xea, 0x6e, 0xa8, 0x89, 0x57, 0x24, 0xf2, 0x52, 0xc3, 0x85, 0x81, 0xd4, 0x05, 0x13, 0xe3,
	0x0d, 0x99, 0xee, 0x1e, 0xda, 0x0d, 0xdb, 0x99, 0x75, 0x66, 0xb6, 0xc8, 0xc7, 0xf1, 0xda, 0x2f,
	0xe2, 0xa5, 0xd7, 0x5e, 0x18, 0xc3, 0x27, 0x31, 0x9d, 0x97, 0xc2, 0x0a, 0xa4, 0xe1, 0x6a, 0xf7,
	0x3c, 0xcf, 0x6f, 0x67, 0xce, 0xce, 0x73, 0x76, 0xe1, 0x5d, 0x3f, 0x51, 0x83, 0xbc, 0xe7, 0x47,
	0x7c, 0x18, 0x1c, 0x0d, 0xf0, 0x68, 0x90, 0xb0, 0xbe, 0x3c, 0x40, 0x75, 0xc6, 0xc5, 0x69, 0xa0,
	0x14, 0x0b, 0x68, 0x96, 0x04, 0xcc, 0xd4, 0x12, 0xc5, 0x08, 0x45, 0xb1, 0xf2, 0x33, 0xc1, 0x15,
	0x27, 0xd5, 0x82, 0x58, 0x7f, 0x75, 0x65, 0xd5, 0x3e, 0xef, 0xf3, 0x40, 0x53, 0xbd, 0xfc, 0x44,
	0x57, 0xba, 0xd0, 0x77, 0xe6, 0xe9, 0xfa, 0x92, 0xdb, 0x88, 0x66, 0x89, 0x95, 0x9e, 0x3a, 0x49,
	0x97, 0x11, 0x4f, 0x83, 0x94, 0x0b, 0x7a, 0x46, 0x59, 0x10, 0xe3, 0x28, 0x89, 0xd0, 0x62, 0x8f,
	0x1c, 0xd6, 0x13, 0xfc, 0x14, 0x85, 0xbd, 0x58, 0x73, 0xd5, 0x99, 0x03, 0xca, 0xe2, 0x14, 0x85,
	0xbb, 0x1a, 0xbb, 0xf9, 0x15, 0xee, 0x75, 0xf4, 0x5a, 0x32, 0xc4, 0x2f, 0x39, 0x4a, 0x45, 0x3e,
	0x40, 0x39, 0xc6, 0xd1, 0x31, 0x8d, 0x63, 0x51, 0xf3, 0xd6, 0xbc, 0xd6, 0xc2, 0xce, 0x9b, 0x5f,
	0xbf, 0x9f, 0xb4, 0xa7, 0x1d, 0x51, 0xc4, 0x05, 0x06, 0xea, 0x3c, 0x43, 0xe9, 0x77, 0x70, 0xb4,
	0x1d, 0xc7, 0x22, 0x9c, 0x8b, 0xcd, 0x0d, 0x59, 0x86, 0xff, 0x4f, 0x8e, 0x23, 0xa6, 0x6a, 0x33,
	0x6b, 0x5e, 0xab, 0x1a, 0xfe, 0x77, 0xb2, 0xcb, 0x54, 0x73, 0x13, 0x16, 0x27, 0x3b, 0xcb, 0x8c,
	0x33, 0x89, 0xe4, 0x39, 0xcc, 0x09, 0x94, 0x79, 0xaa, 0x64, 0xcd, 0x5b, 0x9b, 0x6d, 0x55, 0xda,
	0x8b, 0xbe, 0x7d, 0x61, 0xdf, 0xa0, 0xa1, 0xf3, 0x9b, 0x8b, 0x50, 0x3d, 0x54, 0x54, 0xe5, 0xae,
	0xed, 0xe6, 0xb7, 0x19, 0x28, 0x19, 0x85, 0xb4, 0xa0, 0x24, 0xcf, 0xa5, 0xc2, 0xa1, 0xee, 0xbf,
	0xd2, 0xbe, 0xef, 0x8f, 0x8f, 0xf4, 0x50, 0x4b, 0x63, 0x44, 0x86, 0xd6, 0x27, 0x1b, 0x30, 0x1f,
	0xf1, 0x61, 0xc6, 0x19, 0xda, 0xe6, 0x2a, 0xed, 0x65, 0x0d, 0xef, 0x3a, 0xd5, 0xf0, 0x97, 0x14,
	0x69, 0x42, 0x29, 0xcf, 0xd2, 0x84, 0x9d, 0xd6, 0x2a, 0x9a, 0x07, 0xcd, 0x87, 0x54, 0xa1, 0x0c,
	0xad, 0x43, 0x9e, 0x41, 0x39, 0xe6, 0x67, 0x4c, 0x53, 0x0b, 0xd7, 0xa8, 0x89, 0x47, 0x5e, 0x42,
	0x85, 0x46, 0x2a, 0x19, 0x51, 0x95, 0x70, 0x26, 0x6b, 0xd5, 0x6b, 0xe8, 0x55, 0x9b, 0x6c, 0xc1,
	0xb2, 0x89, 0x5d, 0x1e, 0x67, 0x28, 0x74, 0x40, 0x28, 0x65, 0x6d, 0xe5, 0xca, 0x3b, 0x76, 0x51,
	0x44, 0xc8, 0x54, 0x92, 0xa2, 0x0c, 0x97, 0x2c, 0xdc, 0x45, 0xb1, 0x6d, 0xd0, 0xf6, 0xf7, 0x59,
	0xa8, 0xda, 0xcc, 0x0e, 0xf5, 0x8c, 0x92, 0xf7, 0x00, 0x7b, 0xa8, 0x6c, 0x0e, 0x64, 0xd5, 0x2f,
	0x8e, 0x75, 0x71, 0x32, 0xea, 0x8d, 0xdb, 0x6c, 0x1b, 0xdf, 0x10, 0x96, 0xba, 0x02, 0x33, 0x2a,
	0x70, 0x7b, 0xd2, 0x36, 0x79, 0xe1, 0xdb, 0x71, 0xec, 0x60, 0x3c, 0x3e, 0x9e, 0x88, 0x2a, 0x8c,
	0xcd, 0x93, 0x97, 0x94, 0xdb, 0xe1, 0x2e, 0x30, 0xe9, 0x42, 0xd9, 0x8a, 0x48, 0xd6, 0x7d, 0x37,
	0xd6, 0xd7, 0x69, 0xd3, 0x5d, 0x7d, 0x3a, 0x42, 0x0e, 0xa0, 0xf4, 0xd1, 0x24, 0xb8, 0x7e, 0x53,
	0x23, 0xc6, 0xdb, 0x47, 0x29, 0x69, 0x7f, 0xbc, 0xde, 0x34, 0x84, 0x6c, 0x42, 0xb9, 0xe3, 0xb2,
	0x7e, 0x38, 0xc1, 0xad, 0xe2, 0xd6, 0xb9, 0xcd, 0x68, 0x7f, 0x82, 0x07, 0x85, 0xb0, 0xf6, 0x29,
	0xa3, 0x7d, 0x14, 0x64, 0x0b, 0xe6, 0xf7, 0x50, 0xd9, 0x59, 0x7f, 0xfc, 0x4f, 0x26, 0x85, 0x8f,
	0xa2, 0xbe, 0x72, 0xa3, 0xbb, 0xf3, 0xf6, 0xc7, 0x45, 0xc3, 0xfb, 0x79, 0xd1, 0xf0, 0xfe, 0x5c,
	0x34, 0xbc, 0xcf, 0x1b, 0x77, 0xfe, 0xfb, 0xf5, 0x4a, 0xfa, 0xe7, 0xf1, 0xfa, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0x28, 0xd8, 0x5e, 0x39, 0x05, 0x00, 0x00,
}