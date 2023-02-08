package steamapis

type Inventory interface {
	GetInventory(steamID string, appID string) (*SteamInventory, int)
}

type SteamService struct {
	Inventory
}

func NewSteamService(apiKey string, inventoryHitsPerMin int64, hideLogs bool) *SteamService {
	return &SteamService{
		Inventory: NewInventoryService(apiKey, inventoryHitsPerMin, hideLogs),
	}
}
