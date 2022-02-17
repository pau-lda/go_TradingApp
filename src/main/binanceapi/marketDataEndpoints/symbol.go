package marketDataEndpoints

type Symbol struct {
	Symbol               string        `json:"symbol"`
	Status               string        `json:"status"`
	BaseAsset            string        `json:"baseAsset"`
	BaseAssetPrecision   int64         `json:"baseAssetPrecision"`
	QuoteAsset           string        `json:"quoteAsset"`
	QuotePrecision       int64         `json:"quotePrecision"`
	QuoteAssetPrecision  int64         `json:"quoteAssetPrecision"`
	OrderTypes           []string      `json:"orderTypes"`
	IcebergAllowed       bool          `json:"icebergAllowed"`
	OcoAllowed           bool          `json:"ocoAllowed"`
	SpotTradingAllowed   bool          `json:"isSpotTradingAllowed"`
	MarginTradingAllowed bool          `json:"isMarginTradingAllowed"`
	Filters              []interface{} `json:"filters"`
	Permissions          []string      `json:"permissions"`
}

func (s *Symbol) GetSymbol() string {
	return s.Symbol
}
func (s *Symbol) GetStatus() string {
	return s.Status
}
func (s *Symbol) GetBaseAsset() string {
	return s.BaseAsset
}
func (s *Symbol) GetBaseAssetPrecision() int64 {
	return s.BaseAssetPrecision
}
func (s *Symbol) GetQuoteAsset() string {
	return s.QuoteAsset
}
func (s *Symbol) GetQuotePrecision() int64 {
	return s.QuotePrecision
}
func (s *Symbol) GetQuoteAssetPrecision() int64 {
	return s.QuoteAssetPrecision
}
func (s *Symbol) GetOrderTypes() []string {
	return s.OrderTypes
}
func (s *Symbol) IsIcebergAllowed() bool {
	return s.IcebergAllowed
}
func (s *Symbol) IsOcoAllowed() bool {
	return s.OcoAllowed
}
func (s *Symbol) IsSpotTradingAllowed() bool {
	return s.SpotTradingAllowed
}
func (s *Symbol) IsMarginTradingAllowed() bool {
	return s.MarginTradingAllowed
}
func (s *Symbol) GetFilters() []interface{} {
	return s.Filters
}
func (s *Symbol) getPermissions() []string {
	return s.Permissions
}
