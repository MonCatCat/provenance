// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/reward/v1/reward.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/provenance-io/provenance/x/epoch/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// claim
type Action int32

const (
	ActionAddLiquidity  Action = 0
	ActionSwap          Action = 1
	ActionVote          Action = 2
	ActionDelegateStake Action = 3
)

var Action_name = map[int32]string{
	0: "ActionAddLiquidity",
	1: "ActionSwap",
	2: "ActionVote",
	3: "ActionDelegateStake",
}

var Action_value = map[string]int32{
	"ActionAddLiquidity":  0,
	"ActionSwap":          1,
	"ActionVote":          2,
	"ActionDelegateStake": 3,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c3894741a216575, []int{0}
}

// Params holds parameters for the reward module
type Params struct {
	// distribution epoch identifier
	DistrEpochIdentifier string `protobuf:"bytes,1,opt,name=distr_epoch_identifier,json=distrEpochIdentifier,proto3" json:"distr_epoch_identifier,omitempty" yaml:"distr_epoch_identifier"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c3894741a216575, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDistrEpochIdentifier() string {
	if m != nil {
		return m.DistrEpochIdentifier
	}
	return ""
}

type Criteria struct {
	Id               uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributeTo     string                                   `protobuf:"bytes,2,opt,name=distribute_to,json=distributeTo,proto3" json:"distribute_to,omitempty"`
	Coins            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	StartHeight      int64                                    `protobuf:"varint,4,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight        int64                                    `protobuf:"varint,5,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	DistributedCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=distributed_coins,json=distributedCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"distributed_coins"`
}

func (m *Criteria) Reset()         { *m = Criteria{} }
func (m *Criteria) String() string { return proto.CompactTextString(m) }
func (*Criteria) ProtoMessage()    {}
func (*Criteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c3894741a216575, []int{1}
}
func (m *Criteria) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Criteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Criteria.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Criteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Criteria.Merge(m, src)
}
func (m *Criteria) XXX_Size() int {
	return m.Size()
}
func (m *Criteria) XXX_DiscardUnknown() {
	xxx_messageInfo_Criteria.DiscardUnknown(m)
}

var xxx_messageInfo_Criteria proto.InternalMessageInfo

func (m *Criteria) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Criteria) GetDistributeTo() string {
	if m != nil {
		return m.DistributeTo
	}
	return ""
}

func (m *Criteria) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Criteria) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *Criteria) GetEndHeight() int64 {
	if m != nil {
		return m.EndHeight
	}
	return 0
}

func (m *Criteria) GetDistributedCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DistributedCoins
	}
	return nil
}

// A Reward is the metadata of reward data per address
type Reward struct {
	// address of user reward
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial reward amount for the user
	InitialRewardAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_reward_amount,json=initialRewardAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_reward_amount" yaml:"initial_reward_amount"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,3,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *Reward) Reset()         { *m = Reward{} }
func (m *Reward) String() string { return proto.CompactTextString(m) }
func (*Reward) ProtoMessage()    {}
func (*Reward) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c3894741a216575, []int{2}
}
func (m *Reward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reward.Merge(m, src)
}
func (m *Reward) XXX_Size() int {
	return m.Size()
}
func (m *Reward) XXX_DiscardUnknown() {
	xxx_messageInfo_Reward.DiscardUnknown(m)
}

var xxx_messageInfo_Reward proto.InternalMessageInfo

func (m *Reward) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Reward) GetInitialRewardAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialRewardAmount
	}
	return nil
}

func (m *Reward) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("provenance.reward.v1.Action", Action_name, Action_value)
	proto.RegisterType((*Params)(nil), "provenance.reward.v1.Params")
	proto.RegisterType((*Criteria)(nil), "provenance.reward.v1.Criteria")
	proto.RegisterType((*Reward)(nil), "provenance.reward.v1.Reward")
}

func init() { proto.RegisterFile("provenance/reward/v1/reward.proto", fileDescriptor_0c3894741a216575) }

var fileDescriptor_0c3894741a216575 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0x9d, 0x36, 0x5f, 0x3b, 0xed, 0x57, 0xc2, 0xb4, 0xb4, 0xa1, 0x50, 0xbb, 0x35, 0x9b,
	0x08, 0x51, 0x9b, 0x94, 0x1d, 0xbb, 0xa6, 0x80, 0x40, 0x62, 0x51, 0xb9, 0x08, 0x24, 0x36, 0xd6,
	0xc4, 0x73, 0xeb, 0x8c, 0x1a, 0x7b, 0xcc, 0xcc, 0x24, 0x6d, 0xdf, 0x80, 0x25, 0x6f, 0x00, 0x12,
	0x3b, 0x9e, 0xa4, 0xcb, 0x2e, 0x59, 0x05, 0xd4, 0xbc, 0x41, 0x9e, 0x00, 0x79, 0xc6, 0x51, 0xa2,
	0x52, 0x09, 0x21, 0xb1, 0xf2, 0xdc, 0x73, 0x8f, 0xcf, 0x9c, 0x33, 0x7f, 0x68, 0x27, 0x17, 0x7c,
	0x00, 0x19, 0xc9, 0x62, 0x08, 0x04, 0x9c, 0x12, 0x41, 0x83, 0x41, 0xab, 0x1c, 0xf9, 0xb9, 0xe0,
	0x8a, 0xe3, 0xb5, 0x29, 0xc5, 0x2f, 0x1b, 0x83, 0xd6, 0xe6, 0x5a, 0xc2, 0x13, 0xae, 0x09, 0x41,
	0x31, 0x32, 0xdc, 0x4d, 0x27, 0xe1, 0x3c, 0xe9, 0x41, 0xa0, 0xab, 0x4e, 0xff, 0x38, 0xa0, 0x7d,
	0x41, 0x14, 0xe3, 0x59, 0xd9, 0x77, 0xaf, 0xf7, 0x15, 0x4b, 0x41, 0x2a, 0x92, 0xe6, 0x13, 0x81,
	0x98, 0xcb, 0x94, 0xcb, 0xa0, 0x43, 0x24, 0x04, 0x83, 0x56, 0x07, 0x14, 0x69, 0x05, 0x31, 0x67,
	0x13, 0x81, 0x59, 0xbf, 0x90, 0xf3, 0xb8, 0x5b, 0xd8, 0x4d, 0x20, 0x03, 0xc9, 0xa4, 0xa1, 0x78,
	0x04, 0xd5, 0x0e, 0x89, 0x20, 0xa9, 0xc4, 0xef, 0xd0, 0x3a, 0x65, 0x52, 0x89, 0x48, 0x33, 0x23,
	0x46, 0x21, 0x53, 0xec, 0x98, 0x81, 0x68, 0x58, 0xdb, 0x56, 0x73, 0xb1, 0xbd, 0x33, 0x1e, 0xba,
	0x5b, 0xe7, 0x24, 0xed, 0x3d, 0xf5, 0x6e, 0xe6, 0x79, 0xe1, 0x9a, 0x6e, 0x3c, 0x2f, 0xf0, 0x57,
	0x53, 0x78, 0x68, 0xa3, 0x85, 0x03, 0xc1, 0x14, 0x08, 0x46, 0xf0, 0x0a, 0xb2, 0x19, 0xd5, 0x8a,
	0x73, 0xa1, 0xcd, 0x28, 0x7e, 0x80, 0xfe, 0xd7, 0x3f, 0xb1, 0x4e, 0x5f, 0x41, 0xa4, 0x78, 0xc3,
	0x2e, 0x26, 0x0b, 0x97, 0xa7, 0xe0, 0x1b, 0x8e, 0x09, 0x9a, 0x2f, 0x52, 0xc9, 0x46, 0x75, 0xbb,
	0xda, 0x5c, 0xda, 0xbb, 0xeb, 0x9b, 0xdc, 0x7e, 0x91, 0xdb, 0x2f, 0x73, 0xfb, 0x07, 0x9c, 0x65,
	0xed, 0xc7, 0x17, 0x43, 0xb7, 0xf2, 0xed, 0x87, 0xdb, 0x4c, 0x98, 0xea, 0xf6, 0x3b, 0x7e, 0xcc,
	0xd3, 0xa0, 0x5c, 0x24, 0xf3, 0xd9, 0x95, 0xf4, 0x24, 0x50, 0xe7, 0x39, 0x48, 0xfd, 0x83, 0x0c,
	0x8d, 0x32, 0xde, 0x41, 0xcb, 0x52, 0x11, 0xa1, 0xa2, 0x2e, 0xb0, 0xa4, 0xab, 0x1a, 0x73, 0xdb,
	0x56, 0xb3, 0x1a, 0x2e, 0x69, 0xec, 0xa5, 0x86, 0xf0, 0x16, 0x42, 0x90, 0xd1, 0x09, 0x61, 0x5e,
	0x13, 0x16, 0x21, 0xa3, 0x65, 0xfb, 0x0c, 0xdd, 0x9e, 0x9a, 0xa6, 0x91, 0x31, 0x5c, 0xfb, 0xf7,
	0x86, 0xeb, 0x33, 0xb3, 0x68, 0xc4, 0xfb, 0x62, 0xa3, 0x5a, 0xa8, 0xcf, 0x1a, 0x7e, 0x84, 0xfe,
	0x23, 0x94, 0x0a, 0x90, 0xb2, 0xdc, 0x35, 0x3c, 0x1e, 0xba, 0x2b, 0x66, 0xd7, 0xca, 0x86, 0x17,
	0x4e, 0x28, 0xf8, 0xb3, 0x85, 0xee, 0xb0, 0x8c, 0x29, 0x46, 0x7a, 0x91, 0x39, 0xac, 0x11, 0x49,
	0x79, 0x3f, 0x53, 0x0d, 0xfb, 0x4f, 0xbe, 0x0f, 0x0b, 0xdf, 0xe3, 0xa1, 0x7b, 0xdf, 0x68, 0xdf,
	0xa8, 0xe2, 0xfd, 0x55, 0xae, 0xd5, 0x52, 0xc3, 0x24, 0xd9, 0xd7, 0x0a, 0xf8, 0x05, 0xaa, 0x93,
	0xb8, 0xb8, 0x12, 0x51, 0xcc, 0xd3, 0xbc, 0x07, 0x0a, 0xa8, 0x3e, 0x04, 0x0b, 0xed, 0x7b, 0xe3,
	0xa1, 0xbb, 0x51, 0x06, 0xbb, 0xc6, 0xf0, 0xc2, 0x5b, 0x06, 0x3a, 0x98, 0x20, 0x0f, 0x23, 0x54,
	0xdb, 0xd7, 0x10, 0x5e, 0x47, 0xd8, 0x8c, 0xf6, 0x29, 0x7d, 0xcd, 0x3e, 0xf4, 0x19, 0x65, 0xea,
	0xbc, 0x5e, 0xc1, 0x2b, 0x08, 0x19, 0xfc, 0xe8, 0x94, 0xe4, 0x75, 0x6b, 0x5a, 0xbf, 0xe5, 0x0a,
	0xea, 0x36, 0xde, 0x40, 0xab, 0xa6, 0x7e, 0x06, 0x3d, 0x48, 0x88, 0x82, 0x23, 0x45, 0x4e, 0xa0,
	0x5e, 0xdd, 0x9c, 0xfb, 0xf8, 0xd5, 0xa9, 0xb4, 0x93, 0x8b, 0x2b, 0xc7, 0xba, 0xbc, 0x72, 0xac,
	0x9f, 0x57, 0x8e, 0xf5, 0x69, 0xe4, 0x54, 0x2e, 0x47, 0x4e, 0xe5, 0xfb, 0xc8, 0xa9, 0xa0, 0x0d,
	0xa6, 0xef, 0xfb, 0x6f, 0x8f, 0xc2, 0xa1, 0xf5, 0x7e, 0x6f, 0x66, 0x71, 0xa6, 0x94, 0x5d, 0xc6,
	0x67, 0xaa, 0xe0, 0x6c, 0xf2, 0xd4, 0xe8, 0xc5, 0xea, 0xd4, 0xf4, 0xbd, 0x7d, 0xf2, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xd0, 0x1a, 0x3e, 0xa3, 0x8c, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DistrEpochIdentifier) > 0 {
		i -= len(m.DistrEpochIdentifier)
		copy(dAtA[i:], m.DistrEpochIdentifier)
		i = encodeVarintReward(dAtA, i, uint64(len(m.DistrEpochIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Criteria) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Criteria) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Criteria) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DistributedCoins) > 0 {
		for iNdEx := len(m.DistributedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DistributedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.EndHeight != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.EndHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.StartHeight != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DistributeTo) > 0 {
		i -= len(m.DistributeTo)
		copy(dAtA[i:], m.DistributeTo)
		i = encodeVarintReward(dAtA, i, uint64(len(m.DistributeTo)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Reward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintReward(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InitialRewardAmount) > 0 {
		for iNdEx := len(m.InitialRewardAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialRewardAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintReward(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReward(dAtA []byte, offset int, v uint64) int {
	offset -= sovReward(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DistrEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovReward(uint64(l))
	}
	return n
}

func (m *Criteria) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovReward(uint64(m.Id))
	}
	l = len(m.DistributeTo)
	if l > 0 {
		n += 1 + l + sovReward(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	if m.StartHeight != 0 {
		n += 1 + sovReward(uint64(m.StartHeight))
	}
	if m.EndHeight != 0 {
		n += 1 + sovReward(uint64(m.EndHeight))
	}
	if len(m.DistributedCoins) > 0 {
		for _, e := range m.DistributedCoins {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	return n
}

func (m *Reward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovReward(uint64(l))
	}
	if len(m.InitialRewardAmount) > 0 {
		for _, e := range m.InitialRewardAmount {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovReward(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovReward(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReward(x uint64) (n int) {
	return sovReward(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistrEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReward
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
func (m *Criteria) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
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
			return fmt.Errorf("proto: Criteria: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Criteria: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeTo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributeTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndHeight", wireType)
			}
			m.EndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributedCoins = append(m.DistributedCoins, types.Coin{})
			if err := m.DistributedCoins[len(m.DistributedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReward
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
func (m *Reward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
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
			return fmt.Errorf("proto: Reward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialRewardAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
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
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialRewardAmount = append(m.InitialRewardAmount, types.Coin{})
			if err := m.InitialRewardAmount[len(m.InitialRewardAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowReward
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
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowReward
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthReward
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthReward
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowReward
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
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReward
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
func skipReward(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReward
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
					return 0, ErrIntOverflowReward
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
					return 0, ErrIntOverflowReward
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
				return 0, ErrInvalidLengthReward
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReward
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReward
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReward        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReward          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReward = fmt.Errorf("proto: unexpected end of group")
)
