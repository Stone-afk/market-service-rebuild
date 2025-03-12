package domain

type MarketPrice struct {
	GUID        string `json:"guid"`
	AssetName   string `json:"asset_name"`
	AssetPrice  string `json:"asset_price"`
	AssetVolume string `json:"asset_volume"`
	AssetRate   string `json:"asset_rate"`
	Timestamp   int64  `json:"timestamp"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
