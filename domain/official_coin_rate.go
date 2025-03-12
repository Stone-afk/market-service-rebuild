package domain

type OfficialCoinRate struct {
	GUID      string `json:"guid"`
	Name      string `json:"name"`
	Rate      string `json:"rate"`
	AssetName string `json:"asset_name"`
	BaseAsset string `json:"base_asset"`
	Price     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
