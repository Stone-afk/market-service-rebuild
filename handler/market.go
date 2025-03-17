package handler

import (
	"github.com/ecodeclub/ekit/slice"
	"github.com/gin-gonic/gin"
	"market-service/domain"
	validator "market-service/pkg/xrest/validator"
	"market-service/service"
)

type MarketHandler struct {
	v                   *validator.Validator
	marketPricesSvc     service.MarketPricesService
	officialCoinRateSvc service.OfficialCoinRateService
}

func (h *MarketHandler) GetSupportAsset(ctx *gin.Context, req SupportAssetRequest) (SupportAssetResponse, error) {
	return SupportAssetResponse{
		IsSupport: true,
	}, nil
}

func (h *MarketHandler) GetMarketPrice(ctx *gin.Context, req GetMarketPriceRequest) (GetMarketPriceResponse, error) {
	recordCtx := ctx.Request.Context()
	assetPriceList, err := h.marketPricesSvc.ListMarketPricesByAsset(recordCtx, req.AssetName)
	if err != nil {
		return GetMarketPriceResponse{}, err
	}
	var marketPriceList []domain.MarketPrice
	for _, assetPrice := range assetPriceList {
		mpItem := domain.MarketPrice{
			AssetName:   assetPrice.AssetName,
			AssetPrice:  assetPrice.AssetPrice,
			AssetVolume: assetPrice.AssetVolume,
			AssetRate:   assetPrice.AssetRate,
		}
		marketPriceList = append(marketPriceList, mpItem)
	}

	ocrList, err := h.officialCoinRateSvc.ListOfficialCoinRatesByAsset(recordCtx, req.AssetName)
	if err != nil {
		return GetMarketPriceResponse{}, err
	}
	var officialCoinRateList []domain.OfficialCoinRate
	for _, ocrItem := range ocrList {
		officialCoinRateItem := domain.OfficialCoinRate{
			Name: ocrItem.AssetName,
			Rate: ocrItem.Price,
		}
		officialCoinRateList = append(officialCoinRateList, officialCoinRateItem)
	}
	return GetMarketPriceResponse{
		MarketPriceList: slice.Map(marketPriceList, func(idx int, src domain.MarketPrice) MarketPrice {
			return toMarketPriceVO(src)
		}),
		OfficialCoinRateList: slice.Map(officialCoinRateList, func(idx int, src domain.OfficialCoinRate) OfficialCoinRate {
			return toOfficialCoinRateVO(src)
		}),
	}, nil
}

func NewMarketHandler(v *validator.Validator, marketPricesSvc service.MarketPricesService, officialCoinRateSvc service.OfficialCoinRateService) *MarketHandler {
	return &MarketHandler{
		v:                   v,
		marketPricesSvc:     marketPricesSvc,
		officialCoinRateSvc: officialCoinRateSvc,
	}

}
