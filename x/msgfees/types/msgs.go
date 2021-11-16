package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

const (
	TypeCreateMsgBasedFeeRequest  = "createmsgbasedfee"
	TypeCalculateFeePerMsgRequest = "calculatefeepermsg"
)

// Compile time interface checks.
var (
	_ sdk.Msg            = &CreateMsgBasedFeeRequest{}
	_ sdk.Msg            = &CalculateFeePerMsgRequest{}
	_ legacytx.LegacyMsg = &CreateMsgBasedFeeRequest{}  // For amino support.
	_ legacytx.LegacyMsg = &CalculateFeePerMsgRequest{} // For amino support.
)

func NewMsgBasedFee(msgTypeURL string, additionalFee sdk.Coin) MsgBasedFee {
	return MsgBasedFee{
		MsgTypeUrl: msgTypeURL, AdditionalFee: additionalFee,
	}
}

func (msg *CreateMsgBasedFeeRequest) ValidateBasic() error {
	if msg.MsgBasedFee == nil {
		return ErrEmptyMsgType
	}

	if msg.MsgBasedFee.AdditionalFee.IsZero() {
		return ErrInvalidFee
	}
	if err := msg.MsgBasedFee.AdditionalFee.Validate(); err == nil {
		return err
	}

	return nil
}

func (msg *CreateMsgBasedFeeRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes encodes the message for signing
func (msg *CreateMsgBasedFeeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(legacy.Cdc.MustMarshalJSON(&msg))
}

func (msg *CreateMsgBasedFeeRequest) Type() string {
	return TypeCreateMsgBasedFeeRequest
}

// Route implements Msg
func (msg *CreateMsgBasedFeeRequest) Route() string { return ModuleName }

func NewCalculateFeePerMsgRequest(tx []byte, fromAddress string) CalculateFeePerMsgRequest {
	return CalculateFeePerMsgRequest{
		Tx:          tx,
		FromAddress: fromAddress,
	}
}

func (msg *CalculateFeePerMsgRequest) ValidateBasic() error {
	if len(msg.Tx) == 0 {
		return fmt.Errorf("tx must be defined")
	}
	return nil
}

func (msg *CalculateFeePerMsgRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg *CalculateFeePerMsgRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(legacy.Cdc.MustMarshalJSON(&msg))
}

func (msg *CalculateFeePerMsgRequest) Route() string {
	return ModuleName
}

func (msg *CalculateFeePerMsgRequest) Type() string {
	return TypeCalculateFeePerMsgRequest
}
