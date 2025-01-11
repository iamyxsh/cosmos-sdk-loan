package keeper

import (
	"context"

	"loan/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, ok := k.GetLoan(ctx, msg.Id)
	if !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.State != types.REQUESTED {
		return nil, errorsmod.Wrapf(types.ErrLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)

	err := k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	loan.Lender = msg.Creator
	loan.State = types.IMBURSED
	k.SetLoan(ctx, loan)

	return &types.MsgApproveLoanResponse{}, nil
}
