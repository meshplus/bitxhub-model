// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Message_Type int32

const (
	Message_GET_BLOCK                Message_Type = 0
	Message_GET_BLOCK_ACK            Message_Type = 1
	Message_CONSENSUS                Message_Type = 2
	Message_FETCH_CERT               Message_Type = 3
	Message_FETCH_CERT_ACK           Message_Type = 4
	Message_FETCH_BLOCK_SIGN         Message_Type = 5
	Message_FETCH_BLOCK_SIGN_ACK     Message_Type = 6
	Message_FETCH_IBTP_REQUEST_SIGN  Message_Type = 7
	Message_FETCH_IBTP_RESPONSE_SIGN Message_Type = 8
	Message_FETCH_IBTP_SIGN_ACK      Message_Type = 9
	Message_GET_ADDRESS              Message_Type = 11
	Message_GET_ADDRESS_ACK          Message_Type = 12
	Message_GET_BLOCK_HEADERS        Message_Type = 13
	Message_GET_BLOCK_HEADERS_ACK    Message_Type = 14
	Message_GET_BLOCKS               Message_Type = 15
	Message_GET_BLOCKS_ACK           Message_Type = 16
	Message_CHECK_MASTER_PIER        Message_Type = 17
	Message_CHECK_MASTER_PIER_ACK    Message_Type = 18
	Message_FETCH_BURN_SIGN          Message_Type = 21
	Message_FETCH_BURN_SIGN_ACK      Message_Type = 22
	//pier message
	Message_APPCHAIN_REGISTER       Message_Type = 31
	Message_APPCHAIN_UPDATE         Message_Type = 32
	Message_APPCHAIN_GET            Message_Type = 33
	Message_RULE_DEPLOY             Message_Type = 34
	Message_INTERCHAIN_META_GET     Message_Type = 41
	Message_IBTP_GET                Message_Type = 42
	Message_IBTP_SEND               Message_Type = 43
	Message_IBTP_RECEIPT_GET        Message_Type = 44
	Message_IBTP_RECEIPT_SEND       Message_Type = 45
	Message_ROUTER_IBTP_SEND        Message_Type = 46
	Message_ROUTER_IBTP_GET         Message_Type = 47
	Message_ROUTER_IBTP_RECEIPT_GET Message_Type = 48
	Message_ROUTER_INTERCHAIN_GET   Message_Type = 49
	Message_ADDRESS_GET             Message_Type = 51
	Message_PUBKEY_GET              Message_Type = 52
	Message_PUBKEY_GET_ACK          Message_Type = 53
	Message_ACK                     Message_Type = 61
)

var Message_Type_name = map[int32]string{
	0:  "GET_BLOCK",
	1:  "GET_BLOCK_ACK",
	2:  "CONSENSUS",
	3:  "FETCH_CERT",
	4:  "FETCH_CERT_ACK",
	5:  "FETCH_BLOCK_SIGN",
	6:  "FETCH_BLOCK_SIGN_ACK",
	7:  "FETCH_IBTP_REQUEST_SIGN",
	8:  "FETCH_IBTP_RESPONSE_SIGN",
	9:  "FETCH_IBTP_SIGN_ACK",
	11: "GET_ADDRESS",
	12: "GET_ADDRESS_ACK",
	13: "GET_BLOCK_HEADERS",
	14: "GET_BLOCK_HEADERS_ACK",
	15: "GET_BLOCKS",
	16: "GET_BLOCKS_ACK",
	17: "CHECK_MASTER_PIER",
	18: "CHECK_MASTER_PIER_ACK",
	21: "FETCH_BURN_SIGN",
	22: "FETCH_BURN_SIGN_ACK",
	31: "APPCHAIN_REGISTER",
	32: "APPCHAIN_UPDATE",
	33: "APPCHAIN_GET",
	34: "RULE_DEPLOY",
	41: "INTERCHAIN_META_GET",
	42: "IBTP_GET",
	43: "IBTP_SEND",
	44: "IBTP_RECEIPT_GET",
	45: "IBTP_RECEIPT_SEND",
	46: "ROUTER_IBTP_SEND",
	47: "ROUTER_IBTP_GET",
	48: "ROUTER_IBTP_RECEIPT_GET",
	49: "ROUTER_INTERCHAIN_GET",
	51: "ADDRESS_GET",
	52: "PUBKEY_GET",
	53: "PUBKEY_GET_ACK",
	61: "ACK",
}

var Message_Type_value = map[string]int32{
	"GET_BLOCK":                0,
	"GET_BLOCK_ACK":            1,
	"CONSENSUS":                2,
	"FETCH_CERT":               3,
	"FETCH_CERT_ACK":           4,
	"FETCH_BLOCK_SIGN":         5,
	"FETCH_BLOCK_SIGN_ACK":     6,
	"FETCH_IBTP_REQUEST_SIGN":  7,
	"FETCH_IBTP_RESPONSE_SIGN": 8,
	"FETCH_IBTP_SIGN_ACK":      9,
	"GET_ADDRESS":              11,
	"GET_ADDRESS_ACK":          12,
	"GET_BLOCK_HEADERS":        13,
	"GET_BLOCK_HEADERS_ACK":    14,
	"GET_BLOCKS":               15,
	"GET_BLOCKS_ACK":           16,
	"CHECK_MASTER_PIER":        17,
	"CHECK_MASTER_PIER_ACK":    18,
	"FETCH_BURN_SIGN":          21,
	"FETCH_BURN_SIGN_ACK":      22,
	"APPCHAIN_REGISTER":        31,
	"APPCHAIN_UPDATE":          32,
	"APPCHAIN_GET":             33,
	"RULE_DEPLOY":              34,
	"INTERCHAIN_META_GET":      41,
	"IBTP_GET":                 42,
	"IBTP_SEND":                43,
	"IBTP_RECEIPT_GET":         44,
	"IBTP_RECEIPT_SEND":        45,
	"ROUTER_IBTP_SEND":         46,
	"ROUTER_IBTP_GET":          47,
	"ROUTER_IBTP_RECEIPT_GET":  48,
	"ROUTER_INTERCHAIN_GET":    49,
	"ADDRESS_GET":              51,
	"PUBKEY_GET":               52,
	"PUBKEY_GET_ACK":           53,
	"ACK":                      61,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}

func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0, 0}
}

type Message struct {
	Type    Message_Type `protobuf:"varint,1,opt,name=type,proto3,enum=pb.Message_Type" json:"type,omitempty"`
	Data    []byte       `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Version []byte       `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_GET_BLOCK
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.Message_Type", Message_Type_name, Message_Type_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x40, 0x21, 0x99, 0xf0, 0xb1, 0x4c, 0x42, 0xe3, 0xaa, 0x95, 0x4b, 0x51, 0x0f, 0xe9,
	0x17, 0xfd, 0x48, 0x7b, 0xec, 0xc1, 0xd8, 0x5b, 0xb0, 0x00, 0xe3, 0xee, 0xae, 0x2b, 0xe5, 0x64,
	0x81, 0x64, 0x55, 0x55, 0xa5, 0x80, 0x00, 0xa9, 0xca, 0x9f, 0xa8, 0xfa, 0xb3, 0x7a, 0xcc, 0xb1,
	0xc7, 0x0a, 0xee, 0xfd, 0x0d, 0xd5, 0xac, 0xc1, 0x26, 0xc9, 0xcd, 0xf3, 0xde, 0x9b, 0xc7, 0xbc,
	0x99, 0x05, 0xaa, 0x97, 0xf1, 0xea, 0xc7, 0x6c, 0xf1, 0xbd, 0x33, 0x5f, 0xcc, 0x56, 0x33, 0xcc,
	0xcf, 0xa7, 0xed, 0x7f, 0x25, 0x28, 0x8f, 0xe2, 0xe5, 0x72, 0xf2, 0x35, 0xc6, 0xa7, 0x50, 0x5c,
	0x5d, 0xcd, 0x63, 0xd3, 0x68, 0x19, 0x67, 0xb5, 0x77, 0xac, 0x33, 0x9f, 0x76, 0xb6, 0x54, 0x47,
	0x5d, 0xcd, 0x63, 0xa1, 0x59, 0x44, 0x28, 0xba, 0x93, 0xd5, 0xc4, 0xcc, 0xb7, 0x8c, 0xb3, 0x8a,
	0xd0, 0xdf, 0x68, 0x42, 0xf9, 0x4b, 0xbc, 0x58, 0x7e, 0x9b, 0x5d, 0x9a, 0x05, 0x0d, 0xef, 0xca,
	0xf6, 0xcf, 0x12, 0x14, 0xa9, 0x19, 0xab, 0x70, 0xd8, 0xe3, 0x2a, 0xea, 0x0e, 0xc7, 0xce, 0x80,
	0xe5, 0xb0, 0x01, 0xd5, 0xb4, 0x8c, 0x6c, 0x67, 0xc0, 0x0c, 0x52, 0x38, 0x63, 0x5f, 0x72, 0x5f,
	0x86, 0x92, 0xe5, 0xb1, 0x06, 0xf0, 0x89, 0x2b, 0xa7, 0x1f, 0x39, 0x5c, 0x28, 0x56, 0x40, 0x84,
	0x5a, 0x56, 0xeb, 0x96, 0x22, 0x9e, 0x00, 0x4b, 0xb0, 0xc4, 0x47, 0x7a, 0x3d, 0x9f, 0xdd, 0x43,
	0x13, 0x4e, 0x6e, 0xa3, 0x5a, 0x5f, 0xc2, 0x87, 0x70, 0x9a, 0x30, 0x5e, 0x57, 0x05, 0x91, 0xe0,
	0x9f, 0x43, 0x2e, 0x55, 0xd2, 0x56, 0xc6, 0x47, 0x60, 0xde, 0x20, 0x65, 0x40, 0xd3, 0x24, 0xec,
	0x01, 0x9e, 0xc2, 0xf1, 0x1e, 0x9b, 0x7a, 0x1e, 0x62, 0x1d, 0x8e, 0x28, 0x89, 0xed, 0xba, 0x82,
	0x4b, 0xc9, 0x8e, 0xf0, 0x18, 0xea, 0x7b, 0x80, 0x56, 0x55, 0xb0, 0x09, 0x8d, 0x2c, 0x6f, 0x9f,
	0xdb, 0x2e, 0x17, 0x92, 0x55, 0xf1, 0x01, 0x34, 0xef, 0xc0, 0xba, 0xa3, 0x46, 0xf9, 0x53, 0x4a,
	0xb2, 0x3a, 0xe5, 0xcf, 0x6a, 0xad, 0x61, 0xe4, 0xea, 0xf4, 0xb9, 0x33, 0x88, 0x46, 0xb6, 0x54,
	0x5c, 0x44, 0x81, 0xc7, 0x05, 0x6b, 0x90, 0xeb, 0x1d, 0x58, 0x77, 0x20, 0x0d, 0xb7, 0xdd, 0x4d,
	0x28, 0xfc, 0x24, 0x5b, 0x33, 0xcb, 0x96, 0x82, 0x5a, 0x7d, 0x9f, 0xfc, 0xed, 0x20, 0x70, 0xfa,
	0xb6, 0xe7, 0x47, 0x82, 0xf7, 0x3c, 0x72, 0x63, 0x8f, 0xc9, 0x24, 0x85, 0xc3, 0xc0, 0xb5, 0x15,
	0x67, 0x2d, 0x64, 0x50, 0x49, 0xc1, 0x1e, 0x57, 0xec, 0x09, 0x6d, 0x46, 0x84, 0x43, 0x1e, 0xb9,
	0x3c, 0x18, 0x8e, 0x2f, 0x58, 0x9b, 0x7e, 0xc7, 0xf3, 0x15, 0x17, 0x89, 0x68, 0xc4, 0x95, 0xad,
	0x95, 0xcf, 0xb0, 0x02, 0x07, 0x7a, 0xad, 0x54, 0x3d, 0xa7, 0x87, 0x90, 0x2c, 0x99, 0xfb, 0x2e,
	0x7b, 0x41, 0x47, 0xde, 0x5e, 0xc4, 0xe1, 0x5e, 0xa0, 0xb4, 0xe8, 0x25, 0x8d, 0x76, 0x03, 0xd5,
	0xe2, 0x57, 0x24, 0x16, 0xe3, 0x90, 0x42, 0x67, 0x16, 0x1d, 0x1a, 0x78, 0x1f, 0x25, 0x87, 0xd7,
	0xf4, 0x18, 0xf6, 0xc1, 0x7d, 0xfb, 0x37, 0xb4, 0xc2, 0x1d, 0x99, 0x4d, 0x4c, 0xd4, 0x5b, 0x8a,
	0xb5, 0xbb, 0x2d, 0x01, 0xe7, 0x74, 0xa9, 0x20, 0xec, 0x0e, 0xf8, 0x85, 0xae, 0xdf, 0xd3, 0xa5,
	0xb2, 0x5a, 0x6f, 0xf2, 0x03, 0x96, 0xa1, 0x40, 0x1f, 0x1f, 0xbb, 0xe6, 0xef, 0xb5, 0x65, 0x5c,
	0xaf, 0x2d, 0xe3, 0xef, 0xda, 0x32, 0x7e, 0x6d, 0xac, 0xdc, 0xf5, 0xc6, 0xca, 0xfd, 0xd9, 0x58,
	0xb9, 0x69, 0x49, 0xff, 0x2b, 0xcf, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x86, 0x87, 0xea, 0x8b,
	0xa6, 0x03, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintNetwork(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintNetwork(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintNetwork(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetwork(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetwork(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovNetwork(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovNetwork(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovNetwork(uint64(l))
	}
	return n
}

func sovNetwork(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetwork(x uint64) (n int) {
	return sovNetwork(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Message_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = append(m.Version[:0], dAtA[iNdEx:postIndex]...)
			if m.Version == nil {
				m.Version = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNetwork
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
func skipNetwork(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetwork
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
					return 0, ErrIntOverflowNetwork
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
					return 0, ErrIntOverflowNetwork
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
				return 0, ErrInvalidLengthNetwork
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetwork
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetwork
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetwork        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetwork          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetwork = fmt.Errorf("proto: unexpected end of group")
)
