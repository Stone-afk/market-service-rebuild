package handler

import "market-service/domain"

type SupportAssetRequest struct {
	ConsumerToken string `json:"consumer_token"`
	AssetName     string `json:"asset_name"`
}

type SupportAssetResponse struct {
	IsSupport bool `json:"is_support"`
}

type OfficialCoinRate struct {
	Name string `json:"name"`
	Rate string `json:"rate"`
}

type MarketPrice struct {
	AssetName   string `json:"asset_name"`
	AssetPrice  string `json:"asset_price"`
	AssetVolume string `json:"asset_volume"`
	AssetRate   string `json:"asset_rate"`
}

type GetMarketPriceRequest struct {
	ConsumerToken string `json:"consumer_token"`
	AssetName     string `json:"asset_name"`
}

type GetMarketPriceResponse struct {
	MarketPriceList      []MarketPrice      `json:"market_price_list"`
	OfficialCoinRateList []OfficialCoinRate `json:"official_coin_rate_list"`
}

func toMarketPriceVO(m domain.MarketPrice) MarketPrice {
	return MarketPrice{
		AssetName:   m.AssetName,
		AssetPrice:  m.AssetPrice,
		AssetVolume: m.AssetVolume,
		AssetRate:   m.AssetRate,
	}
}

func toOfficialCoinRateVO(o domain.OfficialCoinRate) OfficialCoinRate {
	return OfficialCoinRate{
		Name: o.Name,
		Rate: o.Rate,
	}
}
