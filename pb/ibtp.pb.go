// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibtp.proto

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

type IBTP_Type int32

const (
	IBTP_INTERCHAIN       IBTP_Type = 0
	IBTP_RECEIPT_SUCCESS  IBTP_Type = 1
	IBTP_RECEIPT_FAILURE  IBTP_Type = 2
	IBTP_ROLLBACK         IBTP_Type = 3
	IBTP_RECEIPT_ROLLBACK IBTP_Type = 4
)

var IBTP_Type_name = map[int32]string{
	0: "INTERCHAIN",
	1: "RECEIPT_SUCCESS",
	2: "RECEIPT_FAILURE",
	3: "ROLLBACK",
	4: "RECEIPT_ROLLBACK",
}

var IBTP_Type_value = map[string]int32{
	"INTERCHAIN":       0,
	"RECEIPT_SUCCESS":  1,
	"RECEIPT_FAILURE":  2,
	"ROLLBACK":         3,
	"RECEIPT_ROLLBACK": 4,
}

func (x IBTP_Type) String() string {
	return proto.EnumName(IBTP_Type_name, int32(x))
}

func (IBTP_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7576a0a5bf0190a3, []int{0, 0}
}

type IBTP_Category int32

const (
	IBTP_REQUEST  IBTP_Category = 0
	IBTP_RESPONSE IBTP_Category = 1
	IBTP_UNKNOWN  IBTP_Category = 2
)

var IBTP_Category_name = map[int32]string{
	0: "REQUEST",
	1: "RESPONSE",
	2: "UNKNOWN",
}

var IBTP_Category_value = map[string]int32{
	"REQUEST":  0,
	"RESPONSE": 1,
	"UNKNOWN":  2,
}

func (x IBTP_Category) String() string {
	return proto.EnumName(IBTP_Category_name, int32(x))
}

func (IBTP_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7576a0a5bf0190a3, []int{0, 1}
}

// Inter-blockchain Transfer Protocol
type IBTP struct {
	// ID of sending chain and sending service
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// ID of receiving chain and receiving service
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// Index of inter-chain transaction
	Index uint64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// inter-chain transaction type
	Type IBTP_Type `protobuf:"varint,4,opt,name=type,proto3,enum=pb.IBTP_Type" json:"type,omitempty"`
	// timeout height of inter-chain transaction on BitXHub
	TimeoutHeight int64 `protobuf:"varint,5,opt,name=timeoutHeight,proto3" json:"timeoutHeight,omitempty"`
	// Proof of inter-chain transactions
	Proof []byte `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"`
	// Encoded content used by inter-chain
	Payload []byte `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
	// info about other txs in the same group
	Group *StringUint64Map `protobuf:"bytes,8,opt,name=group,proto3" json:"group,omitempty"`
	// Message version
	Version string `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	// Self-defined fields used by app-chain
	Extra []byte `protobuf:"bytes,10,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (m *IBTP) Reset()         { *m = IBTP{} }
func (m *IBTP) String() string { return proto.CompactTextString(m) }
func (*IBTP) ProtoMessage()    {}
func (*IBTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_7576a0a5bf0190a3, []int{0}
}
func (m *IBTP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBTP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBTP.Merge(m, src)
}
func (m *IBTP) XXX_Size() int {
	return m.Size()
}
func (m *IBTP) XXX_DiscardUnknown() {
	xxx_messageInfo_IBTP.DiscardUnknown(m)
}

var xxx_messageInfo_IBTP proto.InternalMessageInfo

func (m *IBTP) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *IBTP) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *IBTP) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *IBTP) GetType() IBTP_Type {
	if m != nil {
		return m.Type
	}
	return IBTP_INTERCHAIN
}

func (m *IBTP) GetTimeoutHeight() int64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *IBTP) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *IBTP) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *IBTP) GetGroup() *StringUint64Map {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *IBTP) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *IBTP) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

type Payload struct {
	Encrypted bool   `protobuf:"varint,1,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	Content   []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Hash      []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_7576a0a5bf0190a3, []int{1}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return m.Size()
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func (m *Payload) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Payload) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Content struct {
	Func string   `protobuf:"bytes,1,opt,name=func,proto3" json:"func,omitempty"`
	Args [][]byte `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_7576a0a5bf0190a3, []int{2}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Content.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return m.Size()
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *Content) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

type Result struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_7576a0a5bf0190a3, []int{3}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return m.Size()
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type IBTPs struct {
	Ibtps []*IBTP `protobuf:"bytes,1,rep,name=ibtps,proto3" json:"ibtps,omitempty"`
}

func (m *IBTPs) Reset()         { *m = IBTPs{} }
func (m *IBTPs) String() string { return proto.CompactTextString(m) }
func (*IBTPs) ProtoMessage()    {}
func (*IBTPs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7576a0a5bf0190a3, []int{4}
}
func (m *IBTPs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBTPs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBTPs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBTPs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBTPs.Merge(m, src)
}
func (m *IBTPs) XXX_Size() int {
	return m.Size()
}
func (m *IBTPs) XXX_DiscardUnknown() {
	xxx_messageInfo_IBTPs.DiscardUnknown(m)
}

var xxx_messageInfo_IBTPs proto.InternalMessageInfo

func (m *IBTPs) GetIbtps() []*IBTP {
	if m != nil {
		return m.Ibtps
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.IBTP_Type", IBTP_Type_name, IBTP_Type_value)
	proto.RegisterEnum("pb.IBTP_Category", IBTP_Category_name, IBTP_Category_value)
	proto.RegisterType((*IBTP)(nil), "pb.IBTP")
	proto.RegisterType((*Payload)(nil), "pb.payload")
	proto.RegisterType((*Content)(nil), "pb.content")
	proto.RegisterType((*Result)(nil), "pb.result")
	proto.RegisterType((*IBTPs)(nil), "pb.IBTPs")
}

func init() { proto.RegisterFile("ibtp.proto", fileDescriptor_7576a0a5bf0190a3) }

var fileDescriptor_7576a0a5bf0190a3 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x86, 0x71, 0x08, 0x0b, 0x0c, 0x2c, 0x8d, 0xbc, 0x7b, 0xb0, 0xaa, 0x55, 0x94, 0x46, 0x95,
	0x9a, 0x5e, 0x90, 0x4a, 0xab, 0xde, 0x21, 0x4a, 0xb5, 0x68, 0x69, 0xa0, 0x0e, 0x51, 0x8f, 0x55,
	0x00, 0x03, 0x91, 0x76, 0x63, 0xcb, 0x31, 0xd5, 0xf2, 0x16, 0x7d, 0x92, 0x3e, 0x47, 0x8f, 0x7b,
	0xec, 0xb1, 0x82, 0x17, 0xa9, 0xec, 0x08, 0xba, 0xbd, 0xcd, 0xff, 0xfd, 0xf6, 0x4c, 0x9c, 0xf9,
	0x01, 0xf2, 0x85, 0x12, 0x7d, 0x21, 0xb9, 0xe2, 0xd8, 0x12, 0x8b, 0x97, 0x9d, 0x45, 0x56, 0xe6,
	0xcb, 0x0a, 0xf8, 0x3f, 0xeb, 0x60, 0x8f, 0x47, 0xf3, 0x19, 0xc6, 0x60, 0xaf, 0x25, 0x7f, 0x20,
	0xc8, 0x43, 0x41, 0x9b, 0x9a, 0x1a, 0xf7, 0xc0, 0x52, 0x9c, 0x58, 0x86, 0x58, 0x8a, 0xe3, 0x6b,
	0x68, 0xe4, 0xc5, 0x8a, 0x3d, 0x92, 0xba, 0x87, 0x02, 0x9b, 0x56, 0x02, 0xbf, 0x02, 0x5b, 0xed,
	0x05, 0x23, 0xb6, 0x87, 0x82, 0xde, 0xe0, 0xb2, 0x2f, 0x16, 0x7d, 0xdd, 0xb1, 0x3f, 0xdf, 0x0b,
	0x46, 0x8d, 0x85, 0x5f, 0xc3, 0xa5, 0xca, 0x1f, 0x18, 0xdf, 0xa9, 0x5b, 0x96, 0x6f, 0xb6, 0x8a,
	0x34, 0x3c, 0x14, 0xd4, 0xe9, 0xff, 0x50, 0xb7, 0x17, 0x92, 0xf3, 0x35, 0xb9, 0xf0, 0x50, 0xd0,
	0xa5, 0x95, 0xc0, 0x04, 0x9a, 0x22, 0xdb, 0xdf, 0xf3, 0x6c, 0x45, 0x9a, 0x86, 0x9f, 0x24, 0x7e,
	0x0b, 0x8d, 0x8d, 0xe4, 0x3b, 0x41, 0x5a, 0x1e, 0x0a, 0x3a, 0x83, 0x2b, 0x3d, 0x39, 0x51, 0x32,
	0x2f, 0x36, 0x69, 0x5e, 0xa8, 0x8f, 0x1f, 0x3e, 0x67, 0x82, 0x56, 0x27, 0x74, 0x93, 0xef, 0x4c,
	0x96, 0x39, 0x2f, 0x48, 0xdb, 0x3c, 0xe7, 0x24, 0xf5, 0x50, 0xf6, 0xa8, 0x64, 0x46, 0xa0, 0x1a,
	0x6a, 0x84, 0xbf, 0x02, 0x5b, 0x7f, 0x3e, 0xee, 0x01, 0x8c, 0xe3, 0x79, 0x44, 0xc3, 0xdb, 0xe1,
	0x38, 0x76, 0x6a, 0xf8, 0x0a, 0x5e, 0xd0, 0x28, 0x8c, 0xc6, 0xb3, 0xf9, 0xb7, 0x24, 0x0d, 0xc3,
	0x28, 0x49, 0x1c, 0xf4, 0x1c, 0x7e, 0x1a, 0x8e, 0x27, 0x29, 0x8d, 0x1c, 0x0b, 0x77, 0xa1, 0x45,
	0xa7, 0x93, 0xc9, 0x68, 0x18, 0xde, 0x39, 0x75, 0x7c, 0x0d, 0xce, 0xe9, 0xc8, 0x99, 0xda, 0xfe,
	0x00, 0x5a, 0x61, 0xa6, 0xd8, 0x86, 0xcb, 0x3d, 0xee, 0x40, 0x93, 0x46, 0x5f, 0xd2, 0x28, 0x99,
	0x3b, 0x35, 0x73, 0x39, 0x4a, 0x66, 0xd3, 0x38, 0x89, 0x1c, 0xa4, 0xad, 0x34, 0xbe, 0x8b, 0xa7,
	0x5f, 0x63, 0xc7, 0xf2, 0xd3, 0xf3, 0xef, 0xc0, 0x37, 0xd0, 0x66, 0xc5, 0x52, 0xee, 0x85, 0x62,
	0x2b, 0xb3, 0xb7, 0x16, 0xfd, 0x07, 0xf4, 0x93, 0x97, 0xbc, 0x50, 0xac, 0x50, 0x66, 0x83, 0x5d,
	0x7a, 0x92, 0x7a, 0xd5, 0xdb, 0xac, 0xdc, 0x9a, 0x2d, 0x76, 0xa9, 0xa9, 0xfd, 0x77, 0xf0, 0xdc,
	0x5e, 0xef, 0x8a, 0xe5, 0x39, 0x09, 0xbb, 0x62, 0xa9, 0x59, 0x26, 0x37, 0x25, 0xb1, 0xbc, 0xba,
	0xbe, 0xa2, 0x6b, 0xff, 0x06, 0x2e, 0x24, 0x2b, 0x77, 0xf7, 0xe6, 0xc6, 0x2a, 0x53, 0x19, 0x41,
	0x95, 0xab, 0x6b, 0xff, 0x0d, 0x34, 0x74, 0x0a, 0x4a, 0xec, 0x42, 0x43, 0x07, 0xb0, 0x34, 0x6e,
	0x67, 0xd0, 0x3a, 0xe5, 0x83, 0x56, 0x78, 0x44, 0x7e, 0x1d, 0x5c, 0xf4, 0x74, 0x70, 0xd1, 0x9f,
	0x83, 0x8b, 0x7e, 0x1c, 0xdd, 0xda, 0xd3, 0xd1, 0xad, 0xfd, 0x3e, 0xba, 0xb5, 0xc5, 0x85, 0x89,
	0xe8, 0xfb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0xfe, 0x08, 0x86, 0xc1, 0x02, 0x00, 0x00,
}

func (m *IBTP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBTP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBTP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Group != nil {
		{
			size, err := m.Group.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIbtp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintIbtp(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.Type != 0 {
		i = encodeVarintIbtp(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if m.Index != 0 {
		i = encodeVarintIbtp(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if m.Encrypted {
		i--
		if m.Encrypted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Content) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Content) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Content) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Args) > 0 {
		for iNdEx := len(m.Args) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Args[iNdEx])
			copy(dAtA[i:], m.Args[iNdEx])
			i = encodeVarintIbtp(dAtA, i, uint64(len(m.Args[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Func) > 0 {
		i -= len(m.Func)
		copy(dAtA[i:], m.Func)
		i = encodeVarintIbtp(dAtA, i, uint64(len(m.Func)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Data[iNdEx])
			copy(dAtA[i:], m.Data[iNdEx])
			i = encodeVarintIbtp(dAtA, i, uint64(len(m.Data[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IBTPs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBTPs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBTPs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ibtps) > 0 {
		for iNdEx := len(m.Ibtps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Ibtps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIbtp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintIbtp(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbtp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IBTP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovIbtp(uint64(m.Index))
	}
	if m.Type != 0 {
		n += 1 + sovIbtp(uint64(m.Type))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovIbtp(uint64(m.TimeoutHeight))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	if m.Group != nil {
		l = m.Group.Size()
		n += 1 + l + sovIbtp(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	l = len(m.Extra)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	return n
}

func (m *Payload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Encrypted {
		n += 2
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	return n
}

func (m *Content) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Func)
	if l > 0 {
		n += 1 + l + sovIbtp(uint64(l))
	}
	if len(m.Args) > 0 {
		for _, b := range m.Args {
			l = len(b)
			n += 1 + l + sovIbtp(uint64(l))
		}
	}
	return n
}

func (m *Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, b := range m.Data {
			l = len(b)
			n += 1 + l + sovIbtp(uint64(l))
		}
	}
	return n
}

func (m *IBTPs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ibtps) > 0 {
		for _, e := range m.Ibtps {
			l = e.Size()
			n += 1 + l + sovIbtp(uint64(l))
		}
	}
	return n
}

func sovIbtp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbtp(x uint64) (n int) {
	return sovIbtp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IBTP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbtp
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
			return fmt.Errorf("proto: IBTP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBTP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= IBTP_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Group == nil {
				m.Group = &StringUint64Map{}
			}
			if err := m.Group.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbtp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbtp
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
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbtp
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
			return fmt.Errorf("proto: payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encrypted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Encrypted = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbtp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbtp
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
func (m *Content) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbtp
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
			return fmt.Errorf("proto: content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Func", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Func = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, make([]byte, postIndex-iNdEx))
			copy(m.Args[len(m.Args)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbtp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbtp
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
func (m *Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbtp
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
			return fmt.Errorf("proto: result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, make([]byte, postIndex-iNdEx))
			copy(m.Data[len(m.Data)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbtp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbtp
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
func (m *IBTPs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbtp
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
			return fmt.Errorf("proto: IBTPs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBTPs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ibtps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbtp
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
				return ErrInvalidLengthIbtp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbtp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ibtps = append(m.Ibtps, &IBTP{})
			if err := m.Ibtps[len(m.Ibtps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbtp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbtp
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
func skipIbtp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbtp
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
					return 0, ErrIntOverflowIbtp
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
					return 0, ErrIntOverflowIbtp
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
				return 0, ErrInvalidLengthIbtp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbtp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbtp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbtp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbtp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbtp = fmt.Errorf("proto: unexpected end of group")
)
