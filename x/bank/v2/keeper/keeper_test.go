package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/header"
	coretesting "cosmossdk.io/core/testing"
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	accountskeeper "cosmossdk.io/x/accounts"
	"cosmossdk.io/x/accounts/accountstd"
	asset "cosmossdk.io/x/accounts/defaults/asset"
	assetmock "cosmossdk.io/x/accounts/defaults/asset/mock"
	assetv1 "cosmossdk.io/x/accounts/defaults/asset/v1"
	"cosmossdk.io/x/bank/v2/keeper"
	banktestutil "cosmossdk.io/x/bank/v2/testutil"
	banktypes "cosmossdk.io/x/bank/v2/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	fooDenom   = "foo"   // use default asset account logic
	barDenom   = "bar"   // use custom asset account logic
	stakeDenom = "stake" // use sdk logic
)

var (
	burnerAcc = authtypes.NewEmptyModuleAccount(authtypes.Burner, authtypes.Burner, authtypes.Staking)
	mintAcc   = authtypes.NewEmptyModuleAccount(banktypes.MintModuleName, authtypes.Minter)

	accAddrs = []sdk.AccAddress{
		sdk.AccAddress([]byte("addr1_______________")),
		sdk.AccAddress([]byte("addr2_______________")),
		sdk.AccAddress([]byte("addr3_______________")),
		sdk.AccAddress([]byte("addr4_______________")),
		sdk.AccAddress([]byte("addr5_______________")),
	}
)

func newFooCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(fooDenom, amt)
}

func newBarCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(barDenom, amt)
}

func newStakeCoin(amt int64) sdk.Coin {
	return sdk.NewInt64Coin(stakeDenom, amt)
}

type KeeperTestSuite struct {
	suite.Suite

	ctx          context.Context
	bankKeeper   keeper.Keeper
	addressCodec address.Codec
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(banktypes.StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithHeaderInfo(header.Info{Time: time.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{})

	env := runtime.NewEnvironment(runtime.NewKVStoreService(key), coretesting.NewNopLogger())

	ac := codectestutil.CodecOptions{}.GetAddressCodec()
	suite.addressCodec = ac
	authority := authtypes.NewModuleAddress("gov")

	suite.ctx = ctx

	ir := codectypes.NewInterfaceRegistry()

	accountsKeeper, err := accountskeeper.NewKeeper(codec.NewProtoCodec(ir), env, suite.addressCodec, ir, nil, accountstd.AddAccount(asset.Type, asset.NewAssetAccount))
	suite.Require().NoError(err)
	suite.bankKeeper = *keeper.NewKeeper(
		authority,
		ac,
		env,
		encCfg.Codec,
		accountsKeeper,
	)

	// Init x/account for foo & bar coin

	msgInitFoo := assetv1.MsgInitAssetAccount{
		Owner: authority.String(),
		Denom: fooDenom,
		InitBalance: []assetv1.Balance{
			{
				Addr:   accAddrs[0],
				Amount: math.NewInt(100),
			},
		},
	}

	msgInitFooWraper := &assetv1.MsgInitAssetAccountWrapper{
		MsgInitAssetAccount: msgInitFoo,
		TransferFunc:        asset.DefaultTransfer,
		MintFunc:            asset.DefaultMint,
		BurnFunc:            asset.DefaultBurn,
	}
	_, fooAddr, err := accountsKeeper.Init(suite.ctx, asset.Type, authority, msgInitFooWraper, sdk.NewCoins())
	suite.Require().NoError(err)
	suite.Require().NoError(suite.bankKeeper.SetAssetAccount(suite.ctx, fooDenom, fooAddr))

	// Init bar with custom send, mint, burn logic
	// - send: receiver take 90%, owner take 10% (gov)
	// - mint: receiver take 90%, owner take 10%
	// - burn: just burn 90%
	msgInitBar := assetv1.MsgInitAssetAccount{
		Owner: authority.String(),
		Denom: barDenom,
		InitBalance: []assetv1.Balance{
			{
				Addr:   accAddrs[0],
				Amount: math.NewInt(100),
			},
		},
	}

	msgInitBarWraper := &assetv1.MsgInitAssetAccountWrapper{
		MsgInitAssetAccount: msgInitBar,
		TransferFunc:        assetmock.CustomTransfer,
		MintFunc:            assetmock.CustomMint,
		BurnFunc:            assetmock.CustomBurn,
	}
	_, barAddr, err := accountsKeeper.Init(suite.ctx, asset.Type, authority, msgInitBarWraper, sdk.NewCoins())
	suite.Require().NoError(err)
	suite.Require().NoError(suite.bankKeeper.SetAssetAccount(suite.ctx, barDenom, barAddr))
}

func (suite *KeeperTestSuite) TestSendCoins() {
	ctx := suite.ctx
	require := suite.Require()
	sendAmt := sdk.NewCoins(newFooCoin(10), newBarCoin(10), newStakeCoin(10))

	handler := keeper.NewHandlers(&suite.bankKeeper)

	// Set balances for acc0 and then try send to acc1
	require.NoError(banktestutil.FundAccount(ctx, suite.bankKeeper, accAddrs[0], newStakeCoin(100)))
	_, err := handler.MsgSend(ctx, &banktypes.MsgSend{
		FromAddress: accAddrs[0].String(),
		ToAddress:   accAddrs[1].String(),
		Amount:      sendAmt,
	})
	require.NoError(err)

	// Check acc0 balances
	acc0FooBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], fooDenom)
	require.Equal(acc0FooBalance.Amount, math.NewInt(90))
	acc0BarBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], barDenom)
	require.Equal(acc0BarBalance.Amount, math.NewInt(90))
	acc0StakeBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], stakeDenom)
	require.Equal(acc0StakeBalance.Amount, math.NewInt(90))

	// Check acc1 balances
	acc1FooBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], fooDenom)
	require.Equal(acc1FooBalance.Amount, math.NewInt(10))
	acc1BarBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], barDenom)
	require.Equal(acc1BarBalance.Amount, math.NewInt(9))
	acc1StakeBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], stakeDenom)
	require.Equal(acc1StakeBalance.Amount, math.NewInt(10))

	// Check gov balance
	govBarBalance := suite.bankKeeper.GetBalance(ctx, authtypes.NewModuleAddress("gov"), barDenom)
	require.Equal(govBarBalance.Amount, math.NewInt(1))
}

func (suite *KeeperTestSuite) TestMintCoins() {
	ctx := suite.ctx
	require := suite.Require()
	mintAmt := sdk.NewCoins(newFooCoin(10), newBarCoin(10), newStakeCoin(10))

	handler := keeper.NewHandlers(&suite.bankKeeper)

	_, err := handler.MsgMint(ctx, &banktypes.MsgMint{
		Authority: authtypes.NewModuleAddress("gov").String(),
		ToAddress: accAddrs[1].String(),
		Amount:    mintAmt,
	})
	require.NoError(err)

	// Check acc1 balances
	acc1FooBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], fooDenom)
	require.Equal(acc1FooBalance.Amount, math.NewInt(10))
	acc1BarBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], barDenom)
	require.Equal(acc1BarBalance.Amount, math.NewInt(9))
	acc1StakeBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], stakeDenom)
	require.Equal(acc1StakeBalance.Amount, math.NewInt(10))

	// Check gov balance
	govBarBalance := suite.bankKeeper.GetBalance(ctx, authtypes.NewModuleAddress("gov"), barDenom)
	require.Equal(govBarBalance.Amount, math.NewInt(1))

	fooSupply := suite.bankKeeper.GetSupply(ctx, fooDenom)
	require.Equal(fooSupply.Amount, math.NewInt(110))
	barSupply := suite.bankKeeper.GetSupply(ctx, barDenom)
	require.Equal(barSupply.Amount, math.NewInt(110))
	stakeSupply := suite.bankKeeper.GetSupply(ctx, stakeDenom)
	require.Equal(stakeSupply.Amount, math.NewInt(10))
}

func (suite *KeeperTestSuite) TestBurnCoins() {
	ctx := suite.ctx
	require := suite.Require()
	burnAmount := sdk.NewCoins(newFooCoin(10), newBarCoin(10), newStakeCoin(10))

	handler := keeper.NewHandlers(&suite.bankKeeper)

	// Set balances for acc0 and then try burn
	require.NoError(banktestutil.FundAccount(ctx, suite.bankKeeper, accAddrs[0], newStakeCoin(100)))
	_, err := handler.MsgBurn(ctx, &banktypes.MsgBurn{
		Authority:       authtypes.NewModuleAddress("gov").String(),
		BurnFromAddress: accAddrs[0].String(),
		Amount:          burnAmount,
	})
	require.NoError(err)

	// Check acc0 balances
	acc1FooBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], fooDenom)
	require.Equal(acc1FooBalance.Amount, math.NewInt(90))
	acc1BarBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], barDenom)
	require.Equal(acc1BarBalance.Amount, math.NewInt(91))
	acc1StakeBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], stakeDenom)
	require.Equal(acc1StakeBalance.Amount, math.NewInt(90))

	// Check denoms supply
	fooSupply := suite.bankKeeper.GetSupply(ctx, fooDenom)
	require.Equal(fooSupply.Amount, math.NewInt(90))
	barSupply := suite.bankKeeper.GetSupply(ctx, barDenom)
	require.Equal(barSupply.Amount, math.NewInt(91))
	stakeSupply := suite.bankKeeper.GetSupply(ctx, stakeDenom)
	require.Equal(stakeSupply.Amount, math.NewInt(90))
}

// func (suite *KeeperTestSuite) TestSendCoins_WithRestriction() {
// 	ctx := suite.ctx
// 	require := suite.Require()
// 	balances := sdk.NewCoins(newFooCoin(100), newBarCoin(50))
// 	sendAmt := sdk.NewCoins(newFooCoin(10), newBarCoin(10))

// 	require.NoError(banktestutil.FundAccount(ctx, suite.bankKeeper, accAddrs[0], balances))

// 	// Add first restriction
// 	addrRestrictFunc := func(ctx context.Context, from, to []byte, amount sdk.Coins) ([]byte, error) {
// 		if bytes.Equal(from, to) {
// 			return nil, fmt.Errorf("Can not send to same address")
// 		}
// 		return to, nil
// 	}
// 	suite.bankKeeper.AppendGlobalSendRestriction(addrRestrictFunc)

// 	err := suite.bankKeeper.SendCoins(ctx, accAddrs[0], accAddrs[0], sendAmt)
// 	require.Error(err)
// 	require.Contains(err.Error(), "Can not send to same address")

// 	// Add second restriction
// 	amtRestrictFunc := func(ctx context.Context, from, to []byte, amount sdk.Coins) ([]byte, error) {
// 		if len(amount) > 1 {
// 			return nil, fmt.Errorf("Allow only one denom per one send")
// 		}
// 		return to, nil
// 	}
// 	suite.bankKeeper.AppendGlobalSendRestriction(amtRestrictFunc)

// 	// Pass the 1st but failed at the 2nd
// 	err = suite.bankKeeper.SendCoins(ctx, accAddrs[0], accAddrs[1], sendAmt)
// 	require.Error(err)
// 	require.Contains(err.Error(), "Allow only one denom per one send")

// 	// Pass both 2 restrictions
// 	err = suite.bankKeeper.SendCoins(ctx, accAddrs[0], accAddrs[1], sdk.NewCoins(newFooCoin(10)))
// 	require.NoError(err)

// 	// Check balances
// 	acc0FooBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], fooDenom)
// 	require.Equal(acc0FooBalance.Amount, math.NewInt(90))
// 	acc0BarBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[0], barDenom)
// 	require.Equal(acc0BarBalance.Amount, math.NewInt(50))
// 	acc1FooBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], fooDenom)
// 	require.Equal(acc1FooBalance.Amount, math.NewInt(10))
// 	acc1BarBalance := suite.bankKeeper.GetBalance(ctx, accAddrs[1], barDenom)
// 	require.Equal(acc1BarBalance.Amount, math.ZeroInt())
// }
