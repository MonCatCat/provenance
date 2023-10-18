// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/exchange/v1/orders.proto

package exchange

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Order associates an order id with one of the order types.
type Order struct {
	// order_id is the numerical identifier for this order.
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// order is the specifics of this order.
	//
	// Types that are valid to be assigned to Order:
	//	*Order_AskOrder
	//	*Order_BidOrder
	Order isOrder_Order `protobuf_oneof:"order"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab7cbe63f582471, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

type isOrder_Order interface {
	isOrder_Order()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Order_AskOrder struct {
	AskOrder *AskOrder `protobuf:"bytes,2,opt,name=ask_order,json=askOrder,proto3,oneof" json:"ask_order,omitempty"`
}
type Order_BidOrder struct {
	BidOrder *BidOrder `protobuf:"bytes,3,opt,name=bid_order,json=bidOrder,proto3,oneof" json:"bid_order,omitempty"`
}

func (*Order_AskOrder) isOrder_Order() {}
func (*Order_BidOrder) isOrder_Order() {}

func (m *Order) GetOrder() isOrder_Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Order) GetAskOrder() *AskOrder {
	if x, ok := m.GetOrder().(*Order_AskOrder); ok {
		return x.AskOrder
	}
	return nil
}

func (m *Order) GetBidOrder() *BidOrder {
	if x, ok := m.GetOrder().(*Order_BidOrder); ok {
		return x.BidOrder
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Order) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Order_AskOrder)(nil),
		(*Order_BidOrder)(nil),
	}
}

// AskOrder represents someone's desire to sell something at a minimum price.
type AskOrder struct {
	// market_id identifies the market that this order belongs to.
	MarketId uint32 `protobuf:"varint,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// seller is the address of the account that owns this order and has the assets to sell.
	Seller string `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	// assets are the things that the seller wishes to sell.
	// A hold is placed on this until the order is filled or cancelled.
	Assets types.Coin `protobuf:"bytes,3,opt,name=assets,proto3" json:"assets"`
	// price is the minimum amount that the seller is willing to accept for the assets. The seller's settlement
	// proportional fee (and possibly the settlement flat fee) is taken out of the amount the seller receives,
	// so it's possible that the seller will still receive less than this price.
	Price types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	// seller_settlement_flat_fee is the flat fee for sellers that will be charged during settlement. If this denom is the
	// same denom as the price, it will come out of the actual price received. If this denom is different, the amount must
	// be in the seller's account and a hold is placed on it until the order is filled or cancelled.
	SellerSettlementFlatFee *types.Coin `protobuf:"bytes,5,opt,name=seller_settlement_flat_fee,json=sellerSettlementFlatFee,proto3" json:"seller_settlement_flat_fee,omitempty"`
	// allow_partial should be true if partial fulfillment of this order should be allowed, and should be false if the
	// order must be either filled in full or not filled at all.
	AllowPartial bool `protobuf:"varint,6,opt,name=allow_partial,json=allowPartial,proto3" json:"allow_partial,omitempty"`
	// external_id is an optional string used to externally identify this order. Max length is 100 characters.
	// If an order in this market with this external id already exists, this order will be rejected.
	ExternalId string `protobuf:"bytes,7,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (m *AskOrder) Reset()         { *m = AskOrder{} }
func (m *AskOrder) String() string { return proto.CompactTextString(m) }
func (*AskOrder) ProtoMessage()    {}
func (*AskOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab7cbe63f582471, []int{1}
}
func (m *AskOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AskOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AskOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AskOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskOrder.Merge(m, src)
}
func (m *AskOrder) XXX_Size() int {
	return m.Size()
}
func (m *AskOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_AskOrder.DiscardUnknown(m)
}

var xxx_messageInfo_AskOrder proto.InternalMessageInfo

// BidOrder represents someone's desire to buy something at a specific price.
type BidOrder struct {
	// market_id identifies the market that this order belongs to.
	MarketId uint32 `protobuf:"varint,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// buyer is the address of the account that owns this order and has the price to spend.
	Buyer string `protobuf:"bytes,2,opt,name=buyer,proto3" json:"buyer,omitempty"`
	// assets are the things that the buyer wishes to buy.
	Assets types.Coin `protobuf:"bytes,3,opt,name=assets,proto3" json:"assets"`
	// price is the amount that the buyer will pay for the assets.
	// A hold is placed on this until the order is filled or cancelled.
	Price types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	// buyer_settlement_fees are the fees (both flat and proportional) that the buyer will pay (in addition to the price)
	// when the order is settled. A hold is placed on this until the order is filled or cancelled.
	BuyerSettlementFees github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=buyer_settlement_fees,json=buyerSettlementFees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"buyer_settlement_fees"`
	// allow_partial should be true if partial fulfillment of this order should be allowed, and should be false if the
	// order must be either filled in full or not filled at all.
	AllowPartial bool `protobuf:"varint,6,opt,name=allow_partial,json=allowPartial,proto3" json:"allow_partial,omitempty"`
	// external_id is an optional string used to externally identify this order. Max length is 100 characters.
	// If an order in this market with this external id already exists, this order will be rejected.
	ExternalId string `protobuf:"bytes,7,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (m *BidOrder) Reset()         { *m = BidOrder{} }
func (m *BidOrder) String() string { return proto.CompactTextString(m) }
func (*BidOrder) ProtoMessage()    {}
func (*BidOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab7cbe63f582471, []int{2}
}
func (m *BidOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BidOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidOrder.Merge(m, src)
}
func (m *BidOrder) XXX_Size() int {
	return m.Size()
}
func (m *BidOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_BidOrder.DiscardUnknown(m)
}

var xxx_messageInfo_BidOrder proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Order)(nil), "provenance.exchange.v1.Order")
	proto.RegisterType((*AskOrder)(nil), "provenance.exchange.v1.AskOrder")
	proto.RegisterType((*BidOrder)(nil), "provenance.exchange.v1.BidOrder")
}

func init() {
	proto.RegisterFile("provenance/exchange/v1/orders.proto", fileDescriptor_dab7cbe63f582471)
}

var fileDescriptor_dab7cbe63f582471 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0xb5, 0x9b, 0x38, 0x71, 0xae, 0x74, 0x31, 0x05, 0x9c, 0x20, 0x39, 0x51, 0xba, 0x64, 0xc9,
	0xb9, 0x01, 0x21, 0x24, 0x16, 0xd4, 0x20, 0x55, 0x64, 0xa2, 0x72, 0x25, 0x06, 0x16, 0xeb, 0x6c,
	0x7f, 0x4d, 0x4f, 0x71, 0x7c, 0x91, 0xef, 0x1a, 0xd2, 0x89, 0x95, 0x91, 0x81, 0x1f, 0xc0, 0xcc,
	0x0a, 0x3f, 0xa2, 0x63, 0xc5, 0xc4, 0x04, 0x28, 0x99, 0xf9, 0x0f, 0xc8, 0x77, 0x97, 0x34, 0x48,
	0x10, 0x3a, 0xa0, 0x4e, 0xbe, 0xef, 0xdd, 0x7b, 0xef, 0xfb, 0xfc, 0x4e, 0x77, 0x68, 0x6f, 0x92,
	0xb3, 0x29, 0x64, 0x24, 0x8b, 0xc1, 0x87, 0x59, 0x7c, 0x4a, 0xb2, 0x21, 0xf8, 0xd3, 0x9e, 0xcf,
	0xf2, 0x04, 0x72, 0x8e, 0x27, 0x39, 0x13, 0xcc, 0xb9, 0x7b, 0x45, 0xc2, 0x4b, 0x12, 0x9e, 0xf6,
	0x1a, 0x5e, 0xcc, 0xf8, 0x98, 0x71, 0x3f, 0x22, 0xbc, 0x10, 0x45, 0x20, 0x48, 0xcf, 0x8f, 0x19,
	0xcd, 0x94, 0xae, 0x51, 0x57, 0xfb, 0xa1, 0xac, 0x7c, 0x55, 0xe8, 0xad, 0xdd, 0x21, 0x1b, 0x32,
	0x85, 0x17, 0x2b, 0x85, 0xb6, 0x3f, 0x99, 0xc8, 0x7a, 0x51, 0x74, 0x76, 0xea, 0xc8, 0x96, 0x23,
	0x84, 0x34, 0x71, 0xcd, 0x96, 0xd9, 0x29, 0x07, 0x55, 0x59, 0x0f, 0x12, 0xe7, 0x29, 0xaa, 0x11,
	0x3e, 0x0a, 0x65, 0xe9, 0x6e, 0xb5, 0xcc, 0xce, 0xf6, 0x83, 0x16, 0xfe, 0xf3, 0x84, 0xf8, 0x80,
	0x8f, 0xa4, 0xdf, 0x73, 0x23, 0xb0, 0x89, 0x5e, 0x17, 0x06, 0x11, 0x4d, 0xb4, 0x41, 0x69, 0xb3,
	0x41, 0x9f, 0x26, 0x2b, 0x83, 0x48, 0xaf, 0x9f, 0x94, 0xdf, 0x7e, 0x68, 0x1a, 0xfd, 0x2a, 0xb2,
	0xa4, 0x45, 0xfb, 0xe7, 0x16, 0xb2, 0x97, 0x8d, 0x9c, 0xfb, 0xa8, 0x36, 0x26, 0xf9, 0x08, 0xc4,
	0x72, 0xf2, 0x9d, 0xc0, 0x56, 0xc0, 0x20, 0x71, 0xf6, 0x51, 0x85, 0x43, 0x9a, 0xea, 0xb9, 0x6b,
	0x7d, 0xf7, 0xcb, 0xe7, 0xee, 0xae, 0xce, 0xe5, 0x20, 0x49, 0x72, 0xe0, 0xfc, 0x58, 0xe4, 0x34,
	0x1b, 0x06, 0x9a, 0xe7, 0x3c, 0x46, 0x15, 0xc2, 0x39, 0x08, 0xae, 0x07, 0xad, 0x63, 0x4d, 0x2f,
	0x32, 0xc7, 0x3a, 0x73, 0xfc, 0x8c, 0xd1, 0xac, 0x5f, 0xbe, 0xf8, 0xd6, 0x34, 0x02, 0x4d, 0x77,
	0x1e, 0x21, 0x6b, 0x92, 0xd3, 0x18, 0xdc, 0xf2, 0xf5, 0x74, 0x8a, 0xed, 0xbc, 0x44, 0x0d, 0xd5,
	0x39, 0xe4, 0x20, 0x44, 0x0a, 0x63, 0xc8, 0x44, 0x78, 0x92, 0x12, 0x11, 0x9e, 0x00, 0xb8, 0xd6,
	0x3f, 0xbc, 0x82, 0x7b, 0x4a, 0x7c, 0xbc, 0xd2, 0x1e, 0xa6, 0x44, 0x1c, 0x02, 0x38, 0x7b, 0x68,
	0x87, 0xa4, 0x29, 0x7b, 0x1d, 0x4e, 0x48, 0x2e, 0x28, 0x49, 0xdd, 0x4a, 0xcb, 0xec, 0xd8, 0xc1,
	0x2d, 0x09, 0x1e, 0x29, 0xcc, 0x69, 0xa2, 0x6d, 0x98, 0x09, 0xc8, 0x33, 0x92, 0x16, 0xe9, 0x55,
	0x8b, 0x8c, 0x02, 0xb4, 0x84, 0x06, 0x89, 0x0a, 0xbe, 0xfd, 0xbe, 0x84, 0xec, 0xe5, 0xb9, 0x6c,
	0xce, 0x1b, 0x23, 0x2b, 0x3a, 0x3b, 0xbf, 0x46, 0xdc, 0x8a, 0x76, 0xe3, 0x69, 0xbf, 0x41, 0x77,
	0x64, 0xe3, 0xdf, 0xc2, 0x06, 0xe0, 0xae, 0xd5, 0x2a, 0x6d, 0xb6, 0xd9, 0x2f, 0x6c, 0x3e, 0x7e,
	0x6f, 0x76, 0x86, 0x54, 0x9c, 0x9e, 0x45, 0x38, 0x66, 0x63, 0x7d, 0xc1, 0xf4, 0xa7, 0xcb, 0x93,
	0x91, 0x2f, 0xce, 0x27, 0xc0, 0xa5, 0x80, 0x07, 0xb7, 0x65, 0xa7, 0xb5, 0x93, 0x01, 0xe0, 0xff,
	0xf3, 0x58, 0xfa, 0x70, 0x31, 0xf7, 0xcc, 0xcb, 0xb9, 0x67, 0xfe, 0x98, 0x7b, 0xe6, 0xbb, 0x85,
	0x67, 0x5c, 0x2e, 0x3c, 0xe3, 0xeb, 0xc2, 0x33, 0x50, 0x9d, 0xb2, 0xbf, 0xdc, 0xaf, 0x23, 0xf3,
	0x15, 0x5e, 0xfb, 0x83, 0x2b, 0x52, 0x97, 0xb2, 0xb5, 0xca, 0x9f, 0xad, 0x1e, 0xa7, 0xa8, 0x22,
	0x9f, 0x8a, 0x87, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x1d, 0x88, 0x31, 0xba, 0x04, 0x00,
	0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Order != nil {
		{
			size := m.Order.Size()
			i -= size
			if _, err := m.Order.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.OrderId != 0 {
		i = encodeVarintOrders(dAtA, i, uint64(m.OrderId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Order_AskOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order_AskOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.AskOrder != nil {
		{
			size, err := m.AskOrder.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrders(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Order_BidOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order_BidOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BidOrder != nil {
		{
			size, err := m.BidOrder.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrders(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *AskOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AskOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AskOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExternalId) > 0 {
		i -= len(m.ExternalId)
		copy(dAtA[i:], m.ExternalId)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.ExternalId)))
		i--
		dAtA[i] = 0x3a
	}
	if m.AllowPartial {
		i--
		if m.AllowPartial {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.SellerSettlementFlatFee != nil {
		{
			size, err := m.SellerSettlementFlatFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOrders(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrders(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Assets.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrders(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x12
	}
	if m.MarketId != 0 {
		i = encodeVarintOrders(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BidOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BidOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExternalId) > 0 {
		i -= len(m.ExternalId)
		copy(dAtA[i:], m.ExternalId)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.ExternalId)))
		i--
		dAtA[i] = 0x3a
	}
	if m.AllowPartial {
		i--
		if m.AllowPartial {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.BuyerSettlementFees) > 0 {
		for iNdEx := len(m.BuyerSettlementFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BuyerSettlementFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOrders(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrders(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Assets.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrders(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x12
	}
	if m.MarketId != 0 {
		i = encodeVarintOrders(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrders(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrders(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderId != 0 {
		n += 1 + sovOrders(uint64(m.OrderId))
	}
	if m.Order != nil {
		n += m.Order.Size()
	}
	return n
}

func (m *Order_AskOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AskOrder != nil {
		l = m.AskOrder.Size()
		n += 1 + l + sovOrders(uint64(l))
	}
	return n
}
func (m *Order_BidOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BidOrder != nil {
		l = m.BidOrder.Size()
		n += 1 + l + sovOrders(uint64(l))
	}
	return n
}
func (m *AskOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MarketId != 0 {
		n += 1 + sovOrders(uint64(m.MarketId))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = m.Assets.Size()
	n += 1 + l + sovOrders(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovOrders(uint64(l))
	if m.SellerSettlementFlatFee != nil {
		l = m.SellerSettlementFlatFee.Size()
		n += 1 + l + sovOrders(uint64(l))
	}
	if m.AllowPartial {
		n += 2
	}
	l = len(m.ExternalId)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	return n
}

func (m *BidOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MarketId != 0 {
		n += 1 + sovOrders(uint64(m.MarketId))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = m.Assets.Size()
	n += 1 + l + sovOrders(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovOrders(uint64(l))
	if len(m.BuyerSettlementFees) > 0 {
		for _, e := range m.BuyerSettlementFees {
			l = e.Size()
			n += 1 + l + sovOrders(uint64(l))
		}
	}
	if m.AllowPartial {
		n += 2
	}
	l = len(m.ExternalId)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	return n
}

func sovOrders(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrders(x uint64) (n int) {
	return sovOrders(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrders
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			m.OrderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskOrder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AskOrder{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Order = &Order_AskOrder{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidOrder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BidOrder{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Order = &Order_BidOrder{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrders
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
func (m *AskOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrders
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
			return fmt.Errorf("proto: AskOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AskOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Assets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerSettlementFlatFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SellerSettlementFlatFee == nil {
				m.SellerSettlementFlatFee = &types.Coin{}
			}
			if err := m.SellerSettlementFlatFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowPartial", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
			m.AllowPartial = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrders
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
func (m *BidOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrders
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
			return fmt.Errorf("proto: BidOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BidOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Assets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyerSettlementFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyerSettlementFees = append(m.BuyerSettlementFees, types.Coin{})
			if err := m.BuyerSettlementFees[len(m.BuyerSettlementFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowPartial", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
			m.AllowPartial = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrders
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
func skipOrders(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrders
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
					return 0, ErrIntOverflowOrders
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
					return 0, ErrIntOverflowOrders
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
				return 0, ErrInvalidLengthOrders
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrders
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrders
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrders        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrders          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrders = fmt.Errorf("proto: unexpected end of group")
)
