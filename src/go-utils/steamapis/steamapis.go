package steamapis

type Steam interface {
	Inventory
}

type Market interface {
	Item
}

type SteamApis struct {
	Steam
	Market
}

func NewSteamApis(apiKey string, globalHitsPerMin int64, inventoryHitsPerMin int64, hideLogs bool) *SteamApis {
	return &SteamApis{
		Steam:  NewSteamService(apiKey, inventoryHitsPerMin, hideLogs),
		Market: NewMarketService(apiKey, globalHitsPerMin, hideLogs),
	}
}
