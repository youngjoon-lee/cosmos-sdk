// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gov/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgSubmitProposal defines an sdk.Msg type that supports submitting arbitrary
// proposal Content.
type MsgSubmitProposal struct {
	Content        *types.Any                               `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	InitialDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_deposit,json=initialDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_deposit" yaml:"initial_deposit"`
	Proposer       string                                   `protobuf:"bytes,3,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (m *MsgSubmitProposal) Reset()      { *m = MsgSubmitProposal{} }
func (*MsgSubmitProposal) ProtoMessage() {}
func (*MsgSubmitProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c053992595e3dce, []int{0}
}
func (m *MsgSubmitProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitProposal.Merge(m, src)
}
func (m *MsgSubmitProposal) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitProposal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitProposal proto.InternalMessageInfo

// MsgSubmitProposalResponse defines the Msg/SubmitProposal response type.
type MsgSubmitProposalResponse struct {
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id" yaml:"proposal_id"`
}

func (m *MsgSubmitProposalResponse) Reset()         { *m = MsgSubmitProposalResponse{} }
func (m *MsgSubmitProposalResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitProposalResponse) ProtoMessage()    {}
func (*MsgSubmitProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c053992595e3dce, []int{1}
}
func (m *MsgSubmitProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitProposalResponse.Merge(m, src)
}
func (m *MsgSubmitProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitProposalResponse proto.InternalMessageInfo

func (m *MsgSubmitProposalResponse) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

// MsgVote defines a message to cast a vote.
type MsgVote struct {
	ProposalId uint64     `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id" yaml:"proposal_id"`
	Voter      string     `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	Option     VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"`
}

func (m *MsgVote) Reset()      { *m = MsgVote{} }
func (*MsgVote) ProtoMessage() {}
func (*MsgVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c053992595e3dce, []int{2}
}
func (m *MsgVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVote.Merge(m, src)
}
func (m *MsgVote) XXX_Size() int {
	return m.Size()
}
func (m *MsgVote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVote proto.InternalMessageInfo

// MsgVoteResponse defines the Msg/Vote response type.
type MsgVoteResponse struct {
}

func (m *MsgVoteResponse) Reset()         { *m = MsgVoteResponse{} }
func (m *MsgVoteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgVoteResponse) ProtoMessage()    {}
func (*MsgVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c053992595e3dce, []int{3}
}
func (m *MsgVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteResponse.Merge(m, src)
}
func (m *MsgVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteResponse proto.InternalMessageInfo

// MsgDeposit defines a message to submit a deposit to an existing proposal.
type MsgDeposit struct {
	ProposalId uint64                                   `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id" yaml:"proposal_id"`
	Depositor  string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *MsgDeposit) Reset()      { *m = MsgDeposit{} }
func (*MsgDeposit) ProtoMessage() {}
func (*MsgDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c053992595e3dce, []int{4}
}
func (m *MsgDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeposit.Merge(m, src)
}
func (m *MsgDeposit) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeposit proto.InternalMessageInfo

// MsgDepositResponse defines the Msg/Deposit response type.
type MsgDepositResponse struct {
}

func (m *MsgDepositResponse) Reset()         { *m = MsgDepositResponse{} }
func (m *MsgDepositResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDepositResponse) ProtoMessage()    {}
func (*MsgDepositResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c053992595e3dce, []int{5}
}
func (m *MsgDepositResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositResponse.Merge(m, src)
}
func (m *MsgDepositResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitProposal)(nil), "cosmos.gov.v1beta1.MsgSubmitProposal")
	proto.RegisterType((*MsgSubmitProposalResponse)(nil), "cosmos.gov.v1beta1.MsgSubmitProposalResponse")
	proto.RegisterType((*MsgVote)(nil), "cosmos.gov.v1beta1.MsgVote")
	proto.RegisterType((*MsgVoteResponse)(nil), "cosmos.gov.v1beta1.MsgVoteResponse")
	proto.RegisterType((*MsgDeposit)(nil), "cosmos.gov.v1beta1.MsgDeposit")
	proto.RegisterType((*MsgDepositResponse)(nil), "cosmos.gov.v1beta1.MsgDepositResponse")
}

func init() { proto.RegisterFile("cosmos/gov/v1beta1/tx.proto", fileDescriptor_3c053992595e3dce) }

var fileDescriptor_3c053992595e3dce = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x93, 0xd2, 0xd0, 0x8b, 0x94, 0xd2, 0x53, 0x84, 0x12, 0xb7, 0xb2, 0x23, 0xa3, 0xa2,
	0x2c, 0xb1, 0x69, 0x90, 0x18, 0xca, 0x44, 0x8a, 0x10, 0x20, 0x45, 0x80, 0x91, 0x18, 0x58, 0x22,
	0xdb, 0x71, 0x8d, 0x45, 0xe2, 0x67, 0xe5, 0x2e, 0x51, 0xb3, 0x31, 0x22, 0x06, 0x60, 0x64, 0xcc,
	0xcc, 0x86, 0xc4, 0x1f, 0x51, 0x31, 0x75, 0x64, 0x40, 0x01, 0x25, 0x0b, 0x30, 0xf6, 0x2f, 0x40,
	0xbe, 0x1f, 0x69, 0xd5, 0xa4, 0x01, 0xa4, 0x4e, 0xc9, 0x7b, 0xdf, 0xfb, 0x3e, 0xdd, 0xf7, 0xdd,
	0xf3, 0xa1, 0x4d, 0x1f, 0x48, 0x17, 0x88, 0x1d, 0xc2, 0xc0, 0x1e, 0xec, 0x78, 0x01, 0x75, 0x77,
	0x6c, 0x7a, 0x60, 0x25, 0x3d, 0xa0, 0x80, 0x31, 0x07, 0xad, 0x10, 0x06, 0x96, 0x00, 0x35, 0x5d,
	0x10, 0x3c, 0x97, 0x04, 0x33, 0x86, 0x0f, 0x51, 0xcc, 0x39, 0xda, 0xd6, 0x02, 0xc1, 0x94, 0xcf,
	0xd1, 0x32, 0x47, 0x5b, 0xac, 0xb2, 0x85, 0x3c, 0x87, 0x8a, 0x21, 0x84, 0xc0, 0xfb, 0xe9, 0x3f,
	0x49, 0x08, 0x01, 0xc2, 0x4e, 0x60, 0xb3, 0xca, 0xeb, 0xef, 0xdb, 0x6e, 0x3c, 0xe4, 0x90, 0xf9,
	0x2e, 0x83, 0x36, 0x9a, 0x24, 0x7c, 0xda, 0xf7, 0xba, 0x11, 0x7d, 0xdc, 0x83, 0x04, 0x88, 0xdb,
	0xc1, 0xb7, 0x51, 0xce, 0x87, 0x98, 0x06, 0x31, 0x2d, 0xa9, 0x15, 0xb5, 0x9a, 0xaf, 0x17, 0x2d,
	0x2e, 0x61, 0x49, 0x09, 0xeb, 0x4e, 0x3c, 0x6c, 0xe4, 0xbf, 0x7c, 0xae, 0xe5, 0xf6, 0xf8, 0xa0,
	0x23, 0x19, 0xf8, 0xad, 0x8a, 0xd6, 0xa3, 0x38, 0xa2, 0x91, 0xdb, 0x69, 0xb5, 0x83, 0x04, 0x48,
	0x44, 0x4b, 0x99, 0x4a, 0xb6, 0x9a, 0xaf, 0x97, 0x2d, 0x71, 0xd8, 0xd4, 0xb7, 0x0c, 0xc3, 0xda,
	0x83, 0x28, 0x6e, 0x3c, 0x3c, 0x1c, 0x1b, 0xca, 0xf1, 0xd8, 0xb8, 0x3a, 0x74, 0xbb, 0x9d, 0x5d,
	0xf3, 0x0c, 0xdf, 0xfc, 0xf8, 0xdd, 0xa8, 0x86, 0x11, 0x7d, 0xd1, 0xf7, 0x2c, 0x1f, 0xba, 0xc2,
	0xb3, 0xf8, 0xa9, 0x91, 0xf6, 0x4b, 0x9b, 0x0e, 0x93, 0x80, 0x30, 0x29, 0xe2, 0x14, 0x04, 0xfb,
	0x2e, 0x27, 0x63, 0x0d, 0x5d, 0x4e, 0x98, 0xb3, 0xa0, 0x57, 0xca, 0x56, 0xd4, 0xea, 0x9a, 0x33,
	0xab, 0x77, 0xaf, 0xbc, 0x1e, 0x19, 0xca, 0x87, 0x91, 0xa1, 0xfc, 0x1c, 0x19, 0xca, 0xab, 0x6f,
	0x15, 0xc5, 0xf4, 0x51, 0x79, 0x2e, 0x10, 0x27, 0x20, 0x09, 0xc4, 0x24, 0xc0, 0xf7, 0x50, 0x3e,
	0x11, 0xbd, 0x56, 0xd4, 0x66, 0xe1, 0xac, 0x34, 0xb6, 0x7f, 0x8f, 0x8d, 0xd3, 0xed, 0xe3, 0xb1,
	0x81, 0xb9, 0x8d, 0x53, 0x4d, 0xd3, 0x41, 0xb2, 0x7a, 0xd0, 0x36, 0x3f, 0xa9, 0x28, 0xd7, 0x24,
	0xe1, 0x33, 0xa0, 0x17, 0xa6, 0x89, 0x8b, 0xe8, 0xd2, 0x00, 0x68, 0xd0, 0x2b, 0x65, 0x98, 0x47,
	0x5e, 0xe0, 0x5b, 0x68, 0x15, 0x12, 0x1a, 0x41, 0xcc, 0xac, 0x17, 0xea, 0xba, 0x35, 0xbf, 0x8f,
	0x56, 0x7a, 0x8e, 0x47, 0x6c, 0xca, 0x11, 0xd3, 0x0b, 0x82, 0xd9, 0x40, 0xeb, 0xe2, 0xc8, 0x32,
	0x0e, 0xf3, 0x97, 0x8a, 0x50, 0x93, 0x84, 0x32, 0xe8, 0x8b, 0x72, 0xb2, 0x85, 0xd6, 0xc4, 0xc5,
	0x83, 0x74, 0x73, 0xd2, 0xc0, 0x3e, 0x5a, 0x75, 0xbb, 0xd0, 0x8f, 0x69, 0x29, 0xfb, 0xb7, 0xad,
	0xba, 0x91, 0x6e, 0xd5, 0x7f, 0xed, 0x8e, 0x90, 0x5e, 0x60, 0xbf, 0x88, 0xf0, 0x89, 0x55, 0x99,
	0x40, 0xfd, 0x4d, 0x06, 0x65, 0x9b, 0x24, 0xc4, 0xfb, 0xa8, 0x70, 0xe6, 0x1b, 0xda, 0x5e, 0x14,
	0xf4, 0xdc, 0x66, 0x69, 0xb5, 0x7f, 0x1a, 0x9b, 0x2d, 0xe0, 0x7d, 0xb4, 0xc2, 0x96, 0x66, 0xf3,
	0x1c, 0x5a, 0x0a, 0x6a, 0xd7, 0x96, 0x80, 0x33, 0xa5, 0x27, 0x28, 0x27, 0xef, 0x4d, 0x3f, 0x67,
	0x5e, 0xe0, 0xda, 0xf5, 0xe5, 0xb8, 0x94, 0x6c, 0x34, 0x0e, 0x27, 0xba, 0x7a, 0x34, 0xd1, 0xd5,
	0x1f, 0x13, 0x5d, 0x7d, 0x3f, 0xd5, 0x95, 0xa3, 0xa9, 0xae, 0x7c, 0x9d, 0xea, 0xca, 0xf3, 0xe5,
	0x17, 0x70, 0xc0, 0x1e, 0x3a, 0x76, 0x0d, 0xde, 0x2a, 0x7b, 0x61, 0x6e, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0x09, 0x3f, 0x57, 0x54, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error) {
	out := new(MsgSubmitProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/SubmitProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(context.Context, *MsgSubmitProposal) (*MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(context.Context, *MsgVote) (*MsgVoteResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitProposal(ctx context.Context, req *MsgSubmitProposal) (*MsgSubmitProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProposal not implemented")
}
func (*UnimplementedMsgServer) Vote(ctx context.Context, req *MsgVote) (*MsgVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (*UnimplementedMsgServer) Deposit(ctx context.Context, req *MsgDeposit) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/SubmitProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProposal(ctx, req.(*MsgSubmitProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(ctx, req.(*MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.gov.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitProposal",
			Handler:    _Msg_SubmitProposal_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/gov/v1beta1/tx.proto",
}

func (m *MsgSubmitProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InitialDeposit) > 0 {
		for iNdEx := len(m.InitialDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Option != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Option))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepositResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.InitialDeposit) > 0 {
		for _, e := range m.InitialDeposit {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovTx(uint64(m.ProposalId))
	}
	return n
}

func (m *MsgVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovTx(uint64(m.ProposalId))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Option != 0 {
		n += 1 + sovTx(uint64(m.Option))
	}
	return n
}

func (m *MsgVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovTx(uint64(m.ProposalId))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgDepositResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &types.Any{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialDeposit = append(m.InitialDeposit, types1.Coin{})
			if err := m.InitialDeposit[len(m.InitialDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitProposalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			m.Option = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Option |= VoteOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types1.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDepositResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDepositResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
