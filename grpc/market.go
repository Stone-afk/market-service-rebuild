package grpc

import (
	"context"
	"github.com/ecodeclub/ekit/slice"
	"google.golang.org/grpc"
	"market-service/domain"
	mtv1 "market-service/proto/market"
	"market-service/service"
)

type MarketServiceServer struct {
	mtv1.UnimplementedMarketServicesServer
	marketPricesService     service.MarketPricesService
	officialCoinRateService service.OfficialCoinRateService
}

func (s *MarketServiceServer) GetSupportAsset(ctx context.Context, req *mtv1.SupportAssetRequest) (*mtv1.SupportAssetResponse, error) {
	return &mtv1.SupportAssetResponse{
		ReturnCode: 100,
		Message:    "support this asset",
		IsSupport:  true,
	}, nil
}

func (s *MarketServiceServer) GetMarketPrice(ctx context.Context, req *mtv1.MarketPriceRequest) (*mtv1.MarketPriceResponse, error) {
	assetPriceList, err := s.marketPricesService.ListMarketPricesByAsset(ctx, req.AssetName)
	if err != nil {
		return nil, err
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

	ocrList, err := s.officialCoinRateService.ListOfficialCoinRatesByAsset(ctx, req.AssetName)
	if err != nil {
		return nil, err
	}
	var officialCoinRateList []domain.OfficialCoinRate
	for _, ocrItem := range ocrList {
		officialCoinRateItem := domain.OfficialCoinRate{
			Name: ocrItem.AssetName,
			Rate: ocrItem.Price,
		}
		officialCoinRateList = append(officialCoinRateList, officialCoinRateItem)
	}
	return &mtv1.MarketPriceResponse{
		MarketPrice: slice.Map(marketPriceList, func(idx int, src domain.MarketPrice) *mtv1.MarketPrice {
			return toMarketPriceVO(src)
		}),
		OfficialCoinRate: slice.Map(officialCoinRateList, func(idx int, src domain.OfficialCoinRate) *mtv1.OfficialCoinRate {
			return toOfficialCoinRateVO(src)
		}),
	}, nil
}

func (s *MarketServiceServer) Register(server *grpc.Server) {
	mtv1.RegisterMarketServicesServer(server, s)
}

func NewMarketServiceServer(marketPricesService service.MarketPricesService,
	officialCoinRateService service.OfficialCoinRateService) *MarketServiceServer {
	return &MarketServiceServer{
		marketPricesService:     marketPricesService,
		officialCoinRateService: officialCoinRateService,
	}
}

func toMarketPriceVO(m domain.MarketPrice) *mtv1.MarketPrice {
	return &mtv1.MarketPrice{
		AssetName:   m.AssetName,
		AssetPrice:  m.AssetPrice,
		AssetVolume: m.AssetVolume,
		AssetRate:   m.AssetRate,
	}
}

func toOfficialCoinRateVO(o domain.OfficialCoinRate) *mtv1.OfficialCoinRate {
	return &mtv1.OfficialCoinRate{
		Name: o.Name,
		Rate: o.Rate,
	}
}
