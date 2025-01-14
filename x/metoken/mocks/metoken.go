package mocks

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v6/util/coin"

	ltypes "github.com/umee-network/umee/v6/x/leverage/types"
	"github.com/umee-network/umee/v6/x/metoken"
	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

const (
	USDTBaseDenom    = "ibc/223420B0E8CF9CC47BCAB816AB3A20AE162EED27C1177F4B2BC270C83E11AD8D"
	USDTSymbolDenom  = "USDT"
	USDCBaseDenom    = "ibc/49788C29CD84E08D25CA7BE960BC1F61E88FEFC6333F58557D236D693398466A"
	USDCSymbolDenom  = "USDC"
	ISTBaseDenom     = "ibc/BA460328D9ABA27E643A924071FDB3836E4CE8084C6D2380F25EFAB85CF8EB11"
	ISTSymbolDenom   = "IST"
	WBTCBaseDenom    = "ibc/153B97FE395140EAAA2D7CAC537AF1804AEC5F0595CBC5F1603094018D158C0C"
	WBTCSymbolDenom  = "WBTC"
	ETHBaseDenom     = "ibc/04CE51E6E02243E565AE676DD60336E48D455F8AAD0611FA0299A22FDAC448D6"
	ETHSymbolDenom   = "ETH"
	CMSTBaseDenom    = "ibc/31FA0BA043524F2EFBC9AB0539C43708B8FC549E4800E02D103DDCECAC5FF40C"
	CMSTSymbolDenom  = "CMST"
	MeUSDDenom       = "me/USD"
	MeNonStableDenom = "me/NonStable"
	TestDenom1       = "testDenom1"
	BondDenom        = "uumee"
	MeBondDenom      = "me/" + BondDenom
)

var (
	USDTPrice = sdk.MustNewDecFromStr("0.998")
	USDCPrice = sdk.MustNewDecFromStr("1.0")
	ISTPrice  = sdk.MustNewDecFromStr("1.02")
	CMSTPrice = sdk.MustNewDecFromStr("0.998")
	WBTCPrice = sdk.MustNewDecFromStr("27268.938478585498709550")
	ETHPrice  = sdk.MustNewDecFromStr("1851.789229542837161069")
)

func StableIndex(denom string) metoken.Index {
	return metoken.NewIndex(
		denom,
		sdkmath.NewInt(1_000_000_000_000),
		6,
		ValidFee(),
		[]metoken.AcceptedAsset{
			acceptedAsset(USDTBaseDenom, "0.33"),
			acceptedAsset(USDCBaseDenom, "0.34"),
			acceptedAsset(ISTBaseDenom, "0.33"),
		},
	)
}

func NonStableIndex(denom string) metoken.Index {
	return metoken.NewIndex(
		denom,
		sdkmath.NewInt(1_000_000_000_000),
		8,
		ValidFee(),
		[]metoken.AcceptedAsset{
			acceptedAsset(CMSTBaseDenom, "0.33"),
			acceptedAsset(WBTCBaseDenom, "0.34"),
			acceptedAsset(ETHBaseDenom, "0.33"),
		},
	)
}

func BondIndex() metoken.Index {
	return metoken.Index{
		Denom:     MeBondDenom,
		MaxSupply: sdk.NewInt(1000000_00000),
		Exponent:  6,
		Fee:       ValidFee(),
		AcceptedAssets: []metoken.AcceptedAsset{
			metoken.NewAcceptedAsset(
				BondDenom, sdk.MustNewDecFromStr("0.2"),
				sdk.MustNewDecFromStr("1.0"),
			),
		},
	}
}

func BondBalance() metoken.IndexBalances {
	return metoken.IndexBalances{
		MetokenSupply: coin.Zero(MeBondDenom),
		AssetBalances: []metoken.AssetBalance{
			{
				Denom:     BondDenom,
				Leveraged: sdkmath.ZeroInt(),
				Reserved:  sdkmath.ZeroInt(),
				Fees:      sdkmath.ZeroInt(),
				Interest:  sdkmath.ZeroInt(),
			},
		},
	}
}

func acceptedAsset(denom, targetAllocation string) metoken.AcceptedAsset {
	return metoken.NewAcceptedAsset(denom, sdk.MustNewDecFromStr("0.2"), sdk.MustNewDecFromStr(targetAllocation))
}

func ValidFee() metoken.Fee {
	return metoken.NewFee(
		sdk.MustNewDecFromStr("0.01"),
		sdk.MustNewDecFromStr("0.2"),
		sdk.MustNewDecFromStr("0.5"),
	)
}

func EmptyUSDIndexBalances(denom string) metoken.IndexBalances {
	return metoken.NewIndexBalances(
		sdk.NewCoin(denom, sdkmath.ZeroInt()),
		[]metoken.AssetBalance{
			metoken.NewZeroAssetBalance(USDTBaseDenom),
			metoken.NewZeroAssetBalance(USDCBaseDenom),
			metoken.NewZeroAssetBalance(ISTBaseDenom),
		},
	)
}

func EmptyNonStableIndexBalances(denom string) metoken.IndexBalances {
	return metoken.NewIndexBalances(
		sdk.NewCoin(denom, sdkmath.ZeroInt()),
		[]metoken.AssetBalance{
			metoken.NewZeroAssetBalance(USDTBaseDenom),
			metoken.NewZeroAssetBalance(WBTCBaseDenom),
			metoken.NewZeroAssetBalance(ETHBaseDenom),
		},
	)
}

func ValidUSDIndexBalances(denom string) metoken.IndexBalances {
	return metoken.NewIndexBalances(
		sdk.NewCoin(denom, sdkmath.NewInt(4960_000000)),
		[]metoken.AssetBalance{
			metoken.NewAssetBalance(
				USDTBaseDenom,
				sdkmath.NewInt(960_000000),
				sdkmath.NewInt(240_000000),
				sdkmath.NewInt(34_000000),
				sdkmath.ZeroInt(),
			),
			metoken.NewAssetBalance(
				USDCBaseDenom,
				sdkmath.NewInt(608_000000),
				sdkmath.NewInt(152_000000),
				sdkmath.NewInt(28_000000),
				sdkmath.ZeroInt(),
			),
			metoken.NewAssetBalance(
				ISTBaseDenom,
				sdkmath.NewInt(2400_000000),
				sdkmath.NewInt(600_000000),
				sdkmath.NewInt(76_000000),
				sdkmath.ZeroInt(),
			),
		},
	)
}

// ValidPrices return 24 medians, each one with different prices
func ValidPrices() otypes.Prices {
	prices := otypes.Prices{}
	usdtPrice := USDTPrice.Sub(sdk.MustNewDecFromStr("0.24"))
	usdcPrice := USDCPrice.Sub(sdk.MustNewDecFromStr("0.24"))
	istPrice := ISTPrice.Sub(sdk.MustNewDecFromStr("0.24"))
	cmstPrice := CMSTPrice.Sub(sdk.MustNewDecFromStr("0.24"))
	wbtcPrice := WBTCPrice.Sub(sdk.MustNewDecFromStr("0.24"))
	ethPrice := ETHPrice.Sub(sdk.MustNewDecFromStr("0.24"))
	for i := 1; i <= 24; i++ {
		median := otypes.Price{
			ExchangeRateTuple: otypes.NewExchangeRateTuple(
				USDTSymbolDenom,
				usdtPrice.Add(sdk.MustNewDecFromStr("0.01").MulInt(sdkmath.NewInt(int64(i)))),
			),
			BlockNum: uint64(i),
		}
		prices = append(prices, median)
		median = otypes.Price{
			ExchangeRateTuple: otypes.NewExchangeRateTuple(
				USDCSymbolDenom,
				usdcPrice.Add(sdk.MustNewDecFromStr("0.01").MulInt(sdkmath.NewInt(int64(i)))),
			),
			BlockNum: uint64(i),
		}
		prices = append(prices, median)
		median = otypes.Price{
			ExchangeRateTuple: otypes.NewExchangeRateTuple(
				ISTSymbolDenom,
				istPrice.Add(sdk.MustNewDecFromStr("0.01").MulInt(sdkmath.NewInt(int64(i)))),
			),
			BlockNum: uint64(i),
		}
		prices = append(prices, median)
		median = otypes.Price{
			ExchangeRateTuple: otypes.NewExchangeRateTuple(
				CMSTSymbolDenom,
				cmstPrice.Add(sdk.MustNewDecFromStr("0.01").MulInt(sdkmath.NewInt(int64(i)))),
			),
			BlockNum: uint64(i),
		}
		prices = append(prices, median)
		median = otypes.Price{
			ExchangeRateTuple: otypes.NewExchangeRateTuple(
				WBTCSymbolDenom,
				wbtcPrice.Add(sdk.MustNewDecFromStr("0.01").MulInt(sdkmath.NewInt(int64(i)))),
			),
			BlockNum: uint64(i),
		}
		prices = append(prices, median)
		median = otypes.Price{
			ExchangeRateTuple: otypes.NewExchangeRateTuple(
				ETHSymbolDenom,
				ethPrice.Add(sdk.MustNewDecFromStr("0.01").MulInt(sdkmath.NewInt(int64(i)))),
			),
			BlockNum: uint64(i),
		}
		prices = append(prices, median)
	}

	return prices
}

// ValidPricesFunc return mock func for x/oracle
func ValidPricesFunc() func(ctx sdk.Context) otypes.Prices {
	return func(ctx sdk.Context) otypes.Prices {
		return ValidPrices()
	}
}

func ValidToken(baseDenom, symbolDenom string, exponent uint32) ltypes.Token {
	maxSupply := sdk.NewInt(1000000_00000000)
	if baseDenom == ETHBaseDenom {
		maxSupply = sdk.ZeroInt()
	}
	return ltypes.Token{
		BaseDenom:              baseDenom,
		SymbolDenom:            symbolDenom,
		Exponent:               exponent,
		ReserveFactor:          sdk.MustNewDecFromStr("0.25"),
		CollateralWeight:       sdk.MustNewDecFromStr("0.5"),
		LiquidationThreshold:   sdk.MustNewDecFromStr("0.51"),
		BaseBorrowRate:         sdk.MustNewDecFromStr("0.01"),
		KinkBorrowRate:         sdk.MustNewDecFromStr("0.05"),
		MaxBorrowRate:          sdk.MustNewDecFromStr("1"),
		KinkUtilization:        sdk.MustNewDecFromStr("0.75"),
		LiquidationIncentive:   sdk.MustNewDecFromStr("0.05"),
		EnableMsgSupply:        true,
		EnableMsgBorrow:        true,
		Blacklist:              false,
		MaxCollateralShare:     sdk.MustNewDecFromStr("1"),
		MaxSupplyUtilization:   sdk.MustNewDecFromStr("1"),
		MinCollateralLiquidity: sdk.MustNewDecFromStr("0.05"),
		MaxSupply:              maxSupply,
		HistoricMedians:        24,
	}
}
