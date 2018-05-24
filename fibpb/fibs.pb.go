// Code generated by protoc-gen-gogo.
// source: fibpb/fibs.proto
// DO NOT EDIT!

/*
Package fibpb is a generated protocol buffer package.

It is generated from these files:
	fibpb/fibs.proto

It has these top-level messages:
	GetFibNRecurReq
	GetFibNRecurRes
	GetFibNIterReq
	GetFibNIterRes
*/
package fibpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type GetFibNRecurReq struct {
	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (m *GetFibNRecurReq) Reset()                    { *m = GetFibNRecurReq{} }
func (m *GetFibNRecurReq) String() string            { return proto.CompactTextString(m) }
func (*GetFibNRecurReq) ProtoMessage()               {}
func (*GetFibNRecurReq) Descriptor() ([]byte, []int) { return fileDescriptorFibs, []int{0} }

func (m *GetFibNRecurReq) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type GetFibNRecurRes struct {
	FibNums []int32 `protobuf:"varint,1,rep,packed,name=fib_nums,json=fibNums" json:"fib_nums,omitempty"`
}

func (m *GetFibNRecurRes) Reset()                    { *m = GetFibNRecurRes{} }
func (m *GetFibNRecurRes) String() string            { return proto.CompactTextString(m) }
func (*GetFibNRecurRes) ProtoMessage()               {}
func (*GetFibNRecurRes) Descriptor() ([]byte, []int) { return fileDescriptorFibs, []int{1} }

func (m *GetFibNRecurRes) GetFibNums() []int32 {
	if m != nil {
		return m.FibNums
	}
	return nil
}

type GetFibNIterReq struct {
	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (m *GetFibNIterReq) Reset()                    { *m = GetFibNIterReq{} }
func (m *GetFibNIterReq) String() string            { return proto.CompactTextString(m) }
func (*GetFibNIterReq) ProtoMessage()               {}
func (*GetFibNIterReq) Descriptor() ([]byte, []int) { return fileDescriptorFibs, []int{2} }

func (m *GetFibNIterReq) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type GetFibNIterRes struct {
	FibNums []int32 `protobuf:"varint,1,rep,packed,name=fib_nums,json=fibNums" json:"fib_nums,omitempty"`
}

func (m *GetFibNIterRes) Reset()                    { *m = GetFibNIterRes{} }
func (m *GetFibNIterRes) String() string            { return proto.CompactTextString(m) }
func (*GetFibNIterRes) ProtoMessage()               {}
func (*GetFibNIterRes) Descriptor() ([]byte, []int) { return fileDescriptorFibs, []int{3} }

func (m *GetFibNIterRes) GetFibNums() []int32 {
	if m != nil {
		return m.FibNums
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFibNRecurReq)(nil), "fibpb.GetFibNRecurReq")
	proto.RegisterType((*GetFibNRecurRes)(nil), "fibpb.GetFibNRecurRes")
	proto.RegisterType((*GetFibNIterReq)(nil), "fibpb.GetFibNIterReq")
	proto.RegisterType((*GetFibNIterRes)(nil), "fibpb.GetFibNIterRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Fibg service

type FibgClient interface {
	// GetFibNRecur 递归获取第n个斐波那契数
	GetFibNRecur(ctx context.Context, in *GetFibNRecurReq, opts ...grpc.CallOption) (*GetFibNRecurRes, error)
	// GetFibNIter 迭代获取第n个斐波那契数
	GetFibNIter(ctx context.Context, in *GetFibNIterReq, opts ...grpc.CallOption) (*GetFibNIterRes, error)
}

type fibgClient struct {
	cc *grpc.ClientConn
}

func NewFibgClient(cc *grpc.ClientConn) FibgClient {
	return &fibgClient{cc}
}

func (c *fibgClient) GetFibNRecur(ctx context.Context, in *GetFibNRecurReq, opts ...grpc.CallOption) (*GetFibNRecurRes, error) {
	out := new(GetFibNRecurRes)
	err := grpc.Invoke(ctx, "/fibpb.fibg/GetFibNRecur", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fibgClient) GetFibNIter(ctx context.Context, in *GetFibNIterReq, opts ...grpc.CallOption) (*GetFibNIterRes, error) {
	out := new(GetFibNIterRes)
	err := grpc.Invoke(ctx, "/fibpb.fibg/GetFibNIter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fibg service

type FibgServer interface {
	// GetFibNRecur 递归获取第n个斐波那契数
	GetFibNRecur(context.Context, *GetFibNRecurReq) (*GetFibNRecurRes, error)
	// GetFibNIter 迭代获取第n个斐波那契数
	GetFibNIter(context.Context, *GetFibNIterReq) (*GetFibNIterRes, error)
}

func RegisterFibgServer(s *grpc.Server, srv FibgServer) {
	s.RegisterService(&_Fibg_serviceDesc, srv)
}

func _Fibg_GetFibNRecur_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFibNRecurReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibgServer).GetFibNRecur(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fibpb.fibg/GetFibNRecur",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibgServer).GetFibNRecur(ctx, req.(*GetFibNRecurReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fibg_GetFibNIter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFibNIterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibgServer).GetFibNIter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fibpb.fibg/GetFibNIter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibgServer).GetFibNIter(ctx, req.(*GetFibNIterReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fibg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fibpb.fibg",
	HandlerType: (*FibgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFibNRecur",
			Handler:    _Fibg_GetFibNRecur_Handler,
		},
		{
			MethodName: "GetFibNIter",
			Handler:    _Fibg_GetFibNIter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fibpb/fibs.proto",
}

func (m *GetFibNRecurReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFibNRecurReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Num != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFibs(dAtA, i, uint64(m.Num))
	}
	return i, nil
}

func (m *GetFibNRecurRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFibNRecurRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FibNums) > 0 {
		dAtA2 := make([]byte, len(m.FibNums)*10)
		var j1 int
		for _, num1 := range m.FibNums {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintFibs(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func (m *GetFibNIterReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFibNIterReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Num != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFibs(dAtA, i, uint64(m.Num))
	}
	return i, nil
}

func (m *GetFibNIterRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFibNIterRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FibNums) > 0 {
		dAtA4 := make([]byte, len(m.FibNums)*10)
		var j3 int
		for _, num1 := range m.FibNums {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintFibs(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	return i, nil
}

func encodeFixed64Fibs(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Fibs(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFibs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetFibNRecurReq) Size() (n int) {
	var l int
	_ = l
	if m.Num != 0 {
		n += 1 + sovFibs(uint64(m.Num))
	}
	return n
}

func (m *GetFibNRecurRes) Size() (n int) {
	var l int
	_ = l
	if len(m.FibNums) > 0 {
		l = 0
		for _, e := range m.FibNums {
			l += sovFibs(uint64(e))
		}
		n += 1 + sovFibs(uint64(l)) + l
	}
	return n
}

func (m *GetFibNIterReq) Size() (n int) {
	var l int
	_ = l
	if m.Num != 0 {
		n += 1 + sovFibs(uint64(m.Num))
	}
	return n
}

func (m *GetFibNIterRes) Size() (n int) {
	var l int
	_ = l
	if len(m.FibNums) > 0 {
		l = 0
		for _, e := range m.FibNums {
			l += sovFibs(uint64(e))
		}
		n += 1 + sovFibs(uint64(l)) + l
	}
	return n
}

func sovFibs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFibs(x uint64) (n int) {
	return sovFibs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetFibNRecurReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFibs
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
			return fmt.Errorf("proto: GetFibNRecurReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFibNRecurReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFibs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFibs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFibs
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
func (m *GetFibNRecurRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFibs
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
			return fmt.Errorf("proto: GetFibNRecurRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFibNRecurRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFibs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FibNums = append(m.FibNums, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFibs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthFibs
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFibs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FibNums = append(m.FibNums, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FibNums", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFibs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFibs
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
func (m *GetFibNIterReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFibs
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
			return fmt.Errorf("proto: GetFibNIterReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFibNIterReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFibs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFibs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFibs
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
func (m *GetFibNIterRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFibs
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
			return fmt.Errorf("proto: GetFibNIterRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFibNIterRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFibs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FibNums = append(m.FibNums, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFibs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthFibs
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFibs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FibNums = append(m.FibNums, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FibNums", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFibs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFibs
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
func skipFibs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFibs
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
					return 0, ErrIntOverflowFibs
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
					return 0, ErrIntOverflowFibs
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
				return 0, ErrInvalidLengthFibs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFibs
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
				next, err := skipFibs(dAtA[start:])
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
	ErrInvalidLengthFibs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFibs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("fibpb/fibs.proto", fileDescriptorFibs) }

var fileDescriptorFibs = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcb, 0x4c, 0x2a,
	0x48, 0xd2, 0x4f, 0xcb, 0x4c, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x8b,
	0x28, 0x29, 0x73, 0xf1, 0xbb, 0xa7, 0x96, 0xb8, 0x65, 0x26, 0xf9, 0x05, 0xa5, 0x26, 0x97, 0x16,
	0x05, 0xa5, 0x16, 0x0a, 0x09, 0x70, 0x31, 0xe7, 0x95, 0xe6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0,
	0x06, 0x81, 0x98, 0x4a, 0x3a, 0xe8, 0x8a, 0x8a, 0x85, 0x24, 0xb9, 0x38, 0xd2, 0x32, 0x93, 0xe2,
	0xf3, 0x4a, 0x73, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0x58, 0x83, 0xd8, 0xd3, 0x32, 0x93, 0xfc,
	0x4a, 0x73, 0x8b, 0x95, 0x94, 0xb8, 0xf8, 0xa0, 0xaa, 0x3d, 0x4b, 0x52, 0x71, 0x98, 0xa8, 0x8d,
	0xa6, 0x06, 0x9f, 0x81, 0x46, 0xed, 0x8c, 0x5c, 0x2c, 0x69, 0x99, 0x49, 0xe9, 0x42, 0x0e, 0x5c,
	0x3c, 0xc8, 0xee, 0x10, 0x12, 0xd3, 0x03, 0x7b, 0x42, 0x0f, 0xcd, 0x07, 0x52, 0xd8, 0xc5, 0x8b,
	0x95, 0x18, 0x84, 0x6c, 0xb9, 0xb8, 0x91, 0xec, 0x15, 0x12, 0x45, 0x55, 0x08, 0x75, 0xaf, 0x14,
	0x56, 0xe1, 0x62, 0x25, 0x06, 0x27, 0x9e, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x31, 0x89, 0x0d, 0x1c, 0x92, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75,
	0x56, 0x70, 0xd1, 0x5d, 0x01, 0x00, 0x00,
}
