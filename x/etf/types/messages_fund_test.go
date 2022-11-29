package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFund_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateFund
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateFund{
				Creator:   "invalid_address",
				BaseDenom: "uosmo",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: MsgCreateFund{
				Creator:   sample.AccAddress(),
				BaseDenom: "uosmo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
