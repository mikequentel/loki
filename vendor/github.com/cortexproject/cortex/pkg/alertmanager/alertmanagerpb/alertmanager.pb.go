// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alertmanager.proto

package alertmanagerpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	clusterpb "github.com/prometheus/alertmanager/cluster/clusterpb"
	httpgrpc "github.com/weaveworks/common/httpgrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type UpdateStateStatus int32

const (
	OK             UpdateStateStatus = 0
	MERGE_ERROR    UpdateStateStatus = 2
	USER_NOT_FOUND UpdateStateStatus = 3
)

var UpdateStateStatus_name = map[int32]string{
	0: "OK",
	2: "MERGE_ERROR",
	3: "USER_NOT_FOUND",
}

var UpdateStateStatus_value = map[string]int32{
	"OK":             0,
	"MERGE_ERROR":    2,
	"USER_NOT_FOUND": 3,
}

func (UpdateStateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{0}
}

type UpdateStateResponse struct {
	Status UpdateStateStatus `protobuf:"varint,1,opt,name=status,proto3,enum=alertmanagerpb.UpdateStateStatus" json:"status,omitempty"`
	Error  string            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *UpdateStateResponse) Reset()      { *m = UpdateStateResponse{} }
func (*UpdateStateResponse) ProtoMessage() {}
func (*UpdateStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{0}
}
func (m *UpdateStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStateResponse.Merge(m, src)
}
func (m *UpdateStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdateStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStateResponse proto.InternalMessageInfo

func (m *UpdateStateResponse) GetStatus() UpdateStateStatus {
	if m != nil {
		return m.Status
	}
	return OK
}

func (m *UpdateStateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("alertmanagerpb.UpdateStateStatus", UpdateStateStatus_name, UpdateStateStatus_value)
	proto.RegisterType((*UpdateStateResponse)(nil), "alertmanagerpb.UpdateStateResponse")
}

func init() { proto.RegisterFile("alertmanager.proto", fileDescriptor_e60437b6e0c74c9a) }

var fileDescriptor_e60437b6e0c74c9a = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x6f, 0x1a, 0x31,
	0x18, 0xc6, 0x6d, 0xaa, 0x22, 0xd5, 0xb4, 0x40, 0xdd, 0x3f, 0x42, 0x37, 0x58, 0x94, 0x2e, 0xa8,
	0xc3, 0x9d, 0x44, 0xbb, 0x74, 0xa3, 0x88, 0x6b, 0x91, 0xaa, 0x72, 0xc8, 0xc0, 0xd2, 0x05, 0xf9,
	0x0e, 0x07, 0xa2, 0x70, 0x67, 0xc7, 0xf6, 0x85, 0x35, 0x1f, 0x21, 0x5b, 0xbe, 0x42, 0x3e, 0x4a,
	0x46, 0x46, 0xc6, 0x70, 0x2c, 0x19, 0xf9, 0x08, 0x91, 0x38, 0x02, 0x17, 0x22, 0x65, 0xf2, 0xeb,
	0x47, 0x7e, 0x7e, 0x7a, 0xde, 0xd7, 0x2f, 0xc2, 0x6c, 0xc6, 0x95, 0x09, 0x59, 0xc4, 0x26, 0x5c,
	0xd9, 0x52, 0x09, 0x23, 0x70, 0x31, 0xab, 0x49, 0xdf, 0xfa, 0x31, 0x39, 0x35, 0xd3, 0xd8, 0xb7,
	0x03, 0x11, 0x3a, 0x73, 0xce, 0x2e, 0xf8, 0x5c, 0xa8, 0x33, 0xed, 0x04, 0x22, 0x0c, 0x45, 0xe4,
	0x4c, 0x8d, 0x91, 0x13, 0x25, 0x83, 0x7d, 0x91, 0x52, 0xac, 0x56, 0xc6, 0x25, 0x95, 0x08, 0xb9,
	0x99, 0xf2, 0x58, 0x3b, 0x59, 0xb6, 0x13, 0xcc, 0x62, 0x6d, 0x0e, 0xa7, 0xf4, 0x1f, 0xab, 0x94,
	0x51, 0x3b, 0x41, 0x1f, 0x86, 0x72, 0xcc, 0x0c, 0xef, 0x1b, 0x66, 0x38, 0xe5, 0x5a, 0x8a, 0x48,
	0x73, 0xfc, 0x13, 0xe5, 0xb5, 0x61, 0x26, 0xd6, 0x15, 0x58, 0x85, 0xf5, 0x62, 0xe3, 0x8b, 0xfd,
	0x34, 0xb1, 0x9d, 0x31, 0xf5, 0xb7, 0x0f, 0xe9, 0xce, 0x80, 0x3f, 0xa2, 0xd7, 0x5c, 0x29, 0xa1,
	0x2a, 0xb9, 0x2a, 0xac, 0xbf, 0xa1, 0xe9, 0xe5, 0x5b, 0x13, 0xbd, 0x7f, 0x66, 0xc1, 0x79, 0x94,
	0xf3, 0xfe, 0x96, 0x01, 0x2e, 0xa1, 0xc2, 0x3f, 0x97, 0xfe, 0x71, 0x47, 0x2e, 0xa5, 0x1e, 0x2d,
	0xe7, 0x30, 0x46, 0xc5, 0x61, 0xdf, 0xa5, 0xa3, 0xae, 0x37, 0x18, 0xfd, 0xf6, 0x86, 0xdd, 0x76,
	0xf9, 0x55, 0xe3, 0x1a, 0xa2, 0xb7, 0xbf, 0x32, 0x21, 0x70, 0x13, 0xbd, 0xeb, 0xb0, 0x68, 0x3c,
	0xe3, 0x94, 0x9f, 0xc7, 0x5c, 0x1b, 0xfc, 0xc9, 0xde, 0x0f, 0xa8, 0x33, 0x18, 0xf4, 0x76, 0xb2,
	0xf5, 0xf9, 0x58, 0x4e, 0x7b, 0xac, 0x01, 0xec, 0xa2, 0x42, 0x26, 0x14, 0x2e, 0xd9, 0xfb, 0x29,
	0xd9, 0x3d, 0xa6, 0x8c, 0xf5, 0xf5, 0x85, 0xae, 0x0f, 0x98, 0x56, 0x7b, 0xb1, 0x22, 0x60, 0xb9,
	0x22, 0x60, 0xb3, 0x22, 0xf0, 0x32, 0x21, 0xf0, 0x26, 0x21, 0xf0, 0x36, 0x21, 0x70, 0x91, 0x10,
	0x78, 0x97, 0x10, 0x78, 0x9f, 0x10, 0xb0, 0x49, 0x08, 0xbc, 0x5a, 0x13, 0xb0, 0x58, 0x13, 0xb0,
	0x5c, 0x13, 0xf0, 0xff, 0x68, 0x07, 0xfc, 0xfc, 0xf6, 0x43, 0xbe, 0x3f, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0x96, 0xf1, 0xad, 0x30, 0x02, 0x00, 0x00,
}

func (x UpdateStateStatus) String() string {
	s, ok := UpdateStateStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *UpdateStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateStateResponse)
	if !ok {
		that2, ok := that.(UpdateStateResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *UpdateStateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&alertmanagerpb.UpdateStateResponse{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAlertmanager(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertmanagerClient is the client API for Alertmanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertmanagerClient interface {
	HandleRequest(ctx context.Context, in *httpgrpc.HTTPRequest, opts ...grpc.CallOption) (*httpgrpc.HTTPResponse, error)
	UpdateState(ctx context.Context, in *clusterpb.Part, opts ...grpc.CallOption) (*UpdateStateResponse, error)
}

type alertmanagerClient struct {
	cc *grpc.ClientConn
}

func NewAlertmanagerClient(cc *grpc.ClientConn) AlertmanagerClient {
	return &alertmanagerClient{cc}
}

func (c *alertmanagerClient) HandleRequest(ctx context.Context, in *httpgrpc.HTTPRequest, opts ...grpc.CallOption) (*httpgrpc.HTTPResponse, error) {
	out := new(httpgrpc.HTTPResponse)
	err := c.cc.Invoke(ctx, "/alertmanagerpb.Alertmanager/HandleRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertmanagerClient) UpdateState(ctx context.Context, in *clusterpb.Part, opts ...grpc.CallOption) (*UpdateStateResponse, error) {
	out := new(UpdateStateResponse)
	err := c.cc.Invoke(ctx, "/alertmanagerpb.Alertmanager/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertmanagerServer is the server API for Alertmanager service.
type AlertmanagerServer interface {
	HandleRequest(context.Context, *httpgrpc.HTTPRequest) (*httpgrpc.HTTPResponse, error)
	UpdateState(context.Context, *clusterpb.Part) (*UpdateStateResponse, error)
}

// UnimplementedAlertmanagerServer can be embedded to have forward compatible implementations.
type UnimplementedAlertmanagerServer struct {
}

func (*UnimplementedAlertmanagerServer) HandleRequest(ctx context.Context, req *httpgrpc.HTTPRequest) (*httpgrpc.HTTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRequest not implemented")
}
func (*UnimplementedAlertmanagerServer) UpdateState(ctx context.Context, req *clusterpb.Part) (*UpdateStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}

func RegisterAlertmanagerServer(s *grpc.Server, srv AlertmanagerServer) {
	s.RegisterService(&_Alertmanager_serviceDesc, srv)
}

func _Alertmanager_HandleRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpgrpc.HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertmanagerServer).HandleRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanagerpb.Alertmanager/HandleRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertmanagerServer).HandleRequest(ctx, req.(*httpgrpc.HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alertmanager_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clusterpb.Part)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertmanagerServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanagerpb.Alertmanager/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertmanagerServer).UpdateState(ctx, req.(*clusterpb.Part))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alertmanager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alertmanagerpb.Alertmanager",
	HandlerType: (*AlertmanagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleRequest",
			Handler:    _Alertmanager_HandleRequest_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _Alertmanager_UpdateState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alertmanager.proto",
}

func (m *UpdateStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintAlertmanager(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintAlertmanager(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAlertmanager(dAtA []byte, offset int, v uint64) int {
	offset -= sovAlertmanager(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovAlertmanager(uint64(m.Status))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	return n
}

func sovAlertmanager(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAlertmanager(x uint64) (n int) {
	return sovAlertmanager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UpdateStateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateStateResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAlertmanager(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UpdateStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlertmanager
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
			return fmt.Errorf("proto: UpdateStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= UpdateStateStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
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
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlertmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAlertmanager
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
func skipAlertmanager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlertmanager
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
					return 0, ErrIntOverflowAlertmanager
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
					return 0, ErrIntOverflowAlertmanager
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
				return 0, ErrInvalidLengthAlertmanager
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAlertmanager
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAlertmanager
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
				next, err := skipAlertmanager(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAlertmanager
				}
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
	ErrInvalidLengthAlertmanager = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlertmanager   = fmt.Errorf("proto: integer overflow")
)