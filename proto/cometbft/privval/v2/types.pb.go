// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/privval/v2/types.proto

package v2

import (
	fmt "fmt"
	v1 "github.com/cometbft/cometbft/proto/cometbft/privval/v1"
	v3 "github.com/cometbft/cometbft/proto/cometbft/types/v3"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// SignVoteRequest is a request to sign a vote
type SignVoteRequest struct {
	Vote    *v3.Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
	ChainId string   `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *SignVoteRequest) Reset()         { *m = SignVoteRequest{} }
func (m *SignVoteRequest) String() string { return proto.CompactTextString(m) }
func (*SignVoteRequest) ProtoMessage()    {}
func (*SignVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4682e79d99bb618f, []int{0}
}
func (m *SignVoteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignVoteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignVoteRequest.Merge(m, src)
}
func (m *SignVoteRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignVoteRequest proto.InternalMessageInfo

func (m *SignVoteRequest) GetVote() *v3.Vote {
	if m != nil {
		return m.Vote
	}
	return nil
}

func (m *SignVoteRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

// SignedVoteResponse is a response containing a signed vote or an error
type SignedVoteResponse struct {
	Vote  v3.Vote               `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote"`
	Error *v1.RemoteSignerError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *SignedVoteResponse) Reset()         { *m = SignedVoteResponse{} }
func (m *SignedVoteResponse) String() string { return proto.CompactTextString(m) }
func (*SignedVoteResponse) ProtoMessage()    {}
func (*SignedVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4682e79d99bb618f, []int{1}
}
func (m *SignedVoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignedVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignedVoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignedVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedVoteResponse.Merge(m, src)
}
func (m *SignedVoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *SignedVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignedVoteResponse proto.InternalMessageInfo

func (m *SignedVoteResponse) GetVote() v3.Vote {
	if m != nil {
		return m.Vote
	}
	return v3.Vote{}
}

func (m *SignedVoteResponse) GetError() *v1.RemoteSignerError {
	if m != nil {
		return m.Error
	}
	return nil
}

type Message struct {
	// Types that are valid to be assigned to Sum:
	//	*Message_PubKeyRequest
	//	*Message_PubKeyResponse
	//	*Message_SignVoteRequest
	//	*Message_SignedVoteResponse
	//	*Message_SignProposalRequest
	//	*Message_SignedProposalResponse
	//	*Message_PingRequest
	//	*Message_PingResponse
	Sum isMessage_Sum `protobuf_oneof:"sum"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4682e79d99bb618f, []int{2}
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

type isMessage_Sum interface {
	isMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Message_PubKeyRequest struct {
	PubKeyRequest *v1.PubKeyRequest `protobuf:"bytes,1,opt,name=pub_key_request,json=pubKeyRequest,proto3,oneof" json:"pub_key_request,omitempty"`
}
type Message_PubKeyResponse struct {
	PubKeyResponse *v1.PubKeyResponse `protobuf:"bytes,2,opt,name=pub_key_response,json=pubKeyResponse,proto3,oneof" json:"pub_key_response,omitempty"`
}
type Message_SignVoteRequest struct {
	SignVoteRequest *SignVoteRequest `protobuf:"bytes,3,opt,name=sign_vote_request,json=signVoteRequest,proto3,oneof" json:"sign_vote_request,omitempty"`
}
type Message_SignedVoteResponse struct {
	SignedVoteResponse *SignedVoteResponse `protobuf:"bytes,4,opt,name=signed_vote_response,json=signedVoteResponse,proto3,oneof" json:"signed_vote_response,omitempty"`
}
type Message_SignProposalRequest struct {
	SignProposalRequest *v1.SignProposalRequest `protobuf:"bytes,5,opt,name=sign_proposal_request,json=signProposalRequest,proto3,oneof" json:"sign_proposal_request,omitempty"`
}
type Message_SignedProposalResponse struct {
	SignedProposalResponse *v1.SignedProposalResponse `protobuf:"bytes,6,opt,name=signed_proposal_response,json=signedProposalResponse,proto3,oneof" json:"signed_proposal_response,omitempty"`
}
type Message_PingRequest struct {
	PingRequest *v1.PingRequest `protobuf:"bytes,7,opt,name=ping_request,json=pingRequest,proto3,oneof" json:"ping_request,omitempty"`
}
type Message_PingResponse struct {
	PingResponse *v1.PingResponse `protobuf:"bytes,8,opt,name=ping_response,json=pingResponse,proto3,oneof" json:"ping_response,omitempty"`
}

func (*Message_PubKeyRequest) isMessage_Sum()          {}
func (*Message_PubKeyResponse) isMessage_Sum()         {}
func (*Message_SignVoteRequest) isMessage_Sum()        {}
func (*Message_SignedVoteResponse) isMessage_Sum()     {}
func (*Message_SignProposalRequest) isMessage_Sum()    {}
func (*Message_SignedProposalResponse) isMessage_Sum() {}
func (*Message_PingRequest) isMessage_Sum()            {}
func (*Message_PingResponse) isMessage_Sum()           {}

func (m *Message) GetSum() isMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Message) GetPubKeyRequest() *v1.PubKeyRequest {
	if x, ok := m.GetSum().(*Message_PubKeyRequest); ok {
		return x.PubKeyRequest
	}
	return nil
}

func (m *Message) GetPubKeyResponse() *v1.PubKeyResponse {
	if x, ok := m.GetSum().(*Message_PubKeyResponse); ok {
		return x.PubKeyResponse
	}
	return nil
}

func (m *Message) GetSignVoteRequest() *SignVoteRequest {
	if x, ok := m.GetSum().(*Message_SignVoteRequest); ok {
		return x.SignVoteRequest
	}
	return nil
}

func (m *Message) GetSignedVoteResponse() *SignedVoteResponse {
	if x, ok := m.GetSum().(*Message_SignedVoteResponse); ok {
		return x.SignedVoteResponse
	}
	return nil
}

func (m *Message) GetSignProposalRequest() *v1.SignProposalRequest {
	if x, ok := m.GetSum().(*Message_SignProposalRequest); ok {
		return x.SignProposalRequest
	}
	return nil
}

func (m *Message) GetSignedProposalResponse() *v1.SignedProposalResponse {
	if x, ok := m.GetSum().(*Message_SignedProposalResponse); ok {
		return x.SignedProposalResponse
	}
	return nil
}

func (m *Message) GetPingRequest() *v1.PingRequest {
	if x, ok := m.GetSum().(*Message_PingRequest); ok {
		return x.PingRequest
	}
	return nil
}

func (m *Message) GetPingResponse() *v1.PingResponse {
	if x, ok := m.GetSum().(*Message_PingResponse); ok {
		return x.PingResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_PubKeyRequest)(nil),
		(*Message_PubKeyResponse)(nil),
		(*Message_SignVoteRequest)(nil),
		(*Message_SignedVoteResponse)(nil),
		(*Message_SignProposalRequest)(nil),
		(*Message_SignedProposalResponse)(nil),
		(*Message_PingRequest)(nil),
		(*Message_PingResponse)(nil),
	}
}

func init() {
	proto.RegisterType((*SignVoteRequest)(nil), "cometbft.privval.v2.SignVoteRequest")
	proto.RegisterType((*SignedVoteResponse)(nil), "cometbft.privval.v2.SignedVoteResponse")
	proto.RegisterType((*Message)(nil), "cometbft.privval.v2.Message")
}

func init() { proto.RegisterFile("cometbft/privval/v2/types.proto", fileDescriptor_4682e79d99bb618f) }

var fileDescriptor_4682e79d99bb618f = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x13, 0xd6, 0xae, 0xc3, 0xdd, 0x28, 0x78, 0x03, 0xca, 0x24, 0xb2, 0x52, 0x10, 0x54,
	0x9a, 0x94, 0xa8, 0xa9, 0xc4, 0x89, 0x53, 0xa5, 0x49, 0x41, 0x80, 0xa8, 0x8c, 0x84, 0x04, 0x48,
	0x54, 0xcd, 0x6a, 0xbc, 0x88, 0x35, 0x36, 0xb1, 0x13, 0xa9, 0x1f, 0x80, 0x3b, 0x9f, 0x87, 0x4f,
	0xb0, 0xe3, 0x8e, 0x9c, 0x10, 0x6a, 0xbf, 0x08, 0xca, 0x8b, 0x97, 0xa4, 0x6b, 0x36, 0xed, 0x66,
	0x3f, 0xff, 0xdf, 0xef, 0xfd, 0x9f, 0x9f, 0x6c, 0x74, 0x70, 0xcc, 0x67, 0x54, 0xf9, 0xdf, 0x94,
	0x23, 0xa2, 0x20, 0x49, 0x26, 0xa7, 0x4e, 0xe2, 0x3a, 0x6a, 0x2e, 0xa8, 0xb4, 0x45, 0xc4, 0x15,
	0xc7, 0xbb, 0x17, 0x02, 0x5b, 0x0b, 0xec, 0xc4, 0xdd, 0xaf, 0xc8, 0xea, 0x97, 0xb3, 0xf6, 0x1f,
	0xe7, 0x02, 0x88, 0x3a, 0xc9, 0x60, 0xe5, 0x78, 0x8f, 0x71, 0xc6, 0x61, 0xe9, 0xa4, 0xab, 0x2c,
	0xda, 0xfd, 0x84, 0x5a, 0x1f, 0x02, 0x16, 0x7e, 0xe4, 0x8a, 0x12, 0xfa, 0x23, 0xa6, 0x52, 0xe1,
	0x43, 0x54, 0x4b, 0xb8, 0xa2, 0x6d, 0xb3, 0x63, 0xf6, 0x9a, 0xee, 0x43, 0x3b, 0x37, 0x93, 0xd1,
	0x92, 0x81, 0x0d, 0x6a, 0x10, 0xe1, 0x47, 0x68, 0xeb, 0xf8, 0x64, 0x12, 0x84, 0xe3, 0x60, 0xda,
	0xbe, 0xd5, 0x31, 0x7b, 0xb7, 0x49, 0x03, 0xf6, 0xaf, 0xa7, 0xdd, 0x9f, 0x26, 0xc2, 0x29, 0x9b,
	0x4e, 0x33, 0xba, 0x14, 0x3c, 0x94, 0x14, 0xf7, 0x6f, 0x84, 0x1f, 0xd6, 0xce, 0xfe, 0x1e, 0x18,
	0xba, 0xc8, 0x2b, 0x54, 0xa7, 0x51, 0xc4, 0x23, 0xa8, 0xd0, 0x74, 0x9f, 0xdb, 0xeb, 0xf7, 0xd3,
	0xb7, 0x09, 0x9d, 0x71, 0x45, 0xa1, 0x60, 0x74, 0x94, 0xaa, 0x49, 0x96, 0xd4, 0xfd, 0x5d, 0x47,
	0x8d, 0x77, 0x54, 0xca, 0x09, 0xa3, 0xf8, 0x2d, 0x6a, 0x89, 0xd8, 0x1f, 0x7f, 0xa7, 0xf3, 0x71,
	0x94, 0xb5, 0xab, 0x7d, 0x74, 0x2b, 0x99, 0xa3, 0xd8, 0x7f, 0x43, 0xe7, 0xfa, 0x62, 0x3c, 0x83,
	0xec, 0x88, 0x72, 0x00, 0xbf, 0x47, 0x77, 0x0b, 0x5a, 0xd6, 0x9e, 0xb6, 0xf8, 0xf4, 0x5a, 0x5c,
	0x26, 0xf5, 0x0c, 0x72, 0x47, 0xac, 0x44, 0x30, 0x41, 0xf7, 0x64, 0xc0, 0xc2, 0x71, 0xda, 0x75,
	0x6e, 0x70, 0x03, 0x88, 0xcf, 0x2a, 0x88, 0xae, 0x7d, 0x69, 0x76, 0x9e, 0x41, 0x5a, 0xf2, 0xd2,
	0x38, 0xbf, 0xa0, 0x3d, 0x09, 0x53, 0xb8, 0xa0, 0x6a, 0xa3, 0x35, 0xc0, 0xbe, 0xb8, 0x12, 0xbb,
	0x3a, 0x36, 0xcf, 0x20, 0x58, 0xae, 0x0f, 0xf3, 0x2b, 0xba, 0x0f, 0x86, 0x45, 0xc4, 0x05, 0x97,
	0x93, 0xd3, 0xdc, 0x74, 0x1d, 0xe8, 0xbd, 0xca, 0x6b, 0x48, 0xe9, 0x23, 0x9d, 0x50, 0x18, 0xdf,
	0x95, 0xeb, 0x61, 0xcc, 0x50, 0x5b, 0x9b, 0x2f, 0x55, 0xd0, 0x0d, 0x6c, 0x42, 0x89, 0xc3, 0x2b,
	0x4b, 0xd0, 0x69, 0x41, 0xcb, 0x9b, 0x78, 0x20, 0x2b, 0x4f, 0xf0, 0x11, 0xda, 0x16, 0x41, 0xc8,
	0x72, 0xff, 0x0d, 0x80, 0x77, 0xaa, 0xc7, 0x18, 0x84, 0xac, 0xf0, 0xdd, 0x14, 0xc5, 0x16, 0x7b,
	0x68, 0x47, 0x63, 0xb4, 0xc9, 0x2d, 0xe0, 0x3c, 0xb9, 0x86, 0x93, 0x5b, 0xdb, 0x16, 0xa5, 0xfd,
	0xb0, 0x8e, 0x36, 0x64, 0x3c, 0x1b, 0x8e, 0xce, 0x16, 0x96, 0x79, 0xbe, 0xb0, 0xcc, 0x7f, 0x0b,
	0xcb, 0xfc, 0xb5, 0xb4, 0x8c, 0xf3, 0xa5, 0x65, 0xfc, 0x59, 0x5a, 0xc6, 0xe7, 0x97, 0x2c, 0x50,
	0x27, 0xb1, 0x9f, 0x92, 0x9d, 0xfc, 0xe5, 0x97, 0xfe, 0x88, 0xf4, 0xa1, 0x57, 0x7c, 0x34, 0xfe,
	0x26, 0x1c, 0x0d, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x67, 0x83, 0x2e, 0x86, 0x04, 0x00,
	0x00,
}

func (m *SignVoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignVoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignVoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Vote != nil {
		{
			size, err := m.Vote.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignedVoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignedVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignedVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		{
			size, err := m.Error.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Vote.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message_PubKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PubKeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PubKeyRequest != nil {
		{
			size, err := m.PubKeyRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Message_PubKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PubKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PubKeyResponse != nil {
		{
			size, err := m.PubKeyResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Message_SignVoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_SignVoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignVoteRequest != nil {
		{
			size, err := m.SignVoteRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Message_SignedVoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_SignedVoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignedVoteResponse != nil {
		{
			size, err := m.SignedVoteResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Message_SignProposalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_SignProposalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignProposalRequest != nil {
		{
			size, err := m.SignProposalRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Message_SignedProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_SignedProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignedProposalResponse != nil {
		{
			size, err := m.SignedProposalResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Message_PingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PingRequest != nil {
		{
			size, err := m.PingRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Message_PingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PingResponse != nil {
		{
			size, err := m.PingResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignVoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vote != nil {
		l = m.Vote.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *SignedVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Vote.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Message_PubKeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKeyRequest != nil {
		l = m.PubKeyRequest.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_PubKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKeyResponse != nil {
		l = m.PubKeyResponse.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_SignVoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignVoteRequest != nil {
		l = m.SignVoteRequest.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_SignedVoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedVoteResponse != nil {
		l = m.SignedVoteResponse.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_SignProposalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignProposalRequest != nil {
		l = m.SignProposalRequest.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_SignedProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedProposalResponse != nil {
		l = m.SignedProposalResponse.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_PingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingRequest != nil {
		l = m.PingRequest.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_PingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingResponse != nil {
		l = m.PingResponse.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignVoteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SignVoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignVoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Vote == nil {
				m.Vote = &v3.Vote{}
			}
			if err := m.Vote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SignedVoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SignedVoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignedVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Vote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &v1.RemoteSignerError{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeyRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.PubKeyRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PubKeyRequest{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeyResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.PubKeyResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PubKeyResponse{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignVoteRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignVoteRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_SignVoteRequest{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedVoteResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignedVoteResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_SignedVoteResponse{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignProposalRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.SignProposalRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_SignProposalRequest{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedProposalResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.SignedProposalResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_SignedProposalResponse{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.PingRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PingRequest{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.PingResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PingResponse{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
