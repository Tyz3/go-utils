package steamapis

type Item interface {
	GetItem(marketHashName string, appID string) (*MarketItem, int)
}

type MarketService struct {
	Item
}

func NewMarketService(apiKey string, globalHitsPerMin int64, hideLogs bool) *MarketService {
	return &MarketService{
		Item: NewItemService(apiKey, globalHitsPerMin, hideLogs),
	}
}
