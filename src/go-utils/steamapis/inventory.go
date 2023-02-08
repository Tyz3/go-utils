package steamapis

import (
	"encoding/json"
	"fmt"
	"io"
	"kronos-utils/logger"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type InventoryService struct {
	httpClient *http.Client
	endpoint   string

	hideLogs   bool
	hitsPerMin int64
	lastHitAt  int64
	hitsMu     sync.Mutex
}

func NewInventoryService(apiKey string, hitsPerMin int64, hideLogs bool) *InventoryService {
	return &InventoryService{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		endpoint: strings.Replace(
			"https://api.steamapis.com/steam/inventory/{SteamID}/{AppID}/2?api_key={ApiKey}",
			"{ApiKey}",
			apiKey,
			1,
		),
		hitsPerMin: hitsPerMin,
		hideLogs:   hideLogs,
	}
}

func (s *InventoryService) GetInventory(steamID string, appID string) (*SteamInventory, int) {
	url := strings.NewReplacer("{SteamID}", steamID, "{AppID}", appID).Replace(s.endpoint)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetInventory(%s, %s): %s\n", steamID, appID, err.Error())
		panic(err)
	}

	// Замедляем работу по лимитам
	for {
		s.hitsMu.Lock()
		if s.lastHitAt+60*1000/s.hitsPerMin < time.Now().UnixMilli() {
			// Обновляем последнее время
			s.lastHitAt = time.Now().UnixMilli()
			s.hitsMu.Unlock()

			// Do Work
			return s.makeRequest(request)
		}
		s.hitsMu.Unlock()
	}
}

func (s *InventoryService) makeRequest(request *http.Request) (*SteamInventory, int) {
	memo := logger.NewMemo("SteamAPIs->GetInventory")
	if s.hideLogs {
		memo.HiddenMode()
	}
	memo.Info(request.URL.String())

	response, _ := s.httpClient.Do(request)
	if response == nil {
		memo.Error(fmt.Sprintf("Отсутствует ответ: %+v", response)).Timestamp().Failed()
		return nil, 0
	}

	data, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	memo.Info(fmt.Sprintf("status=%d", response.StatusCode))

	if response.StatusCode == 403 {
		memo.Warn("Профиль закрыт").Timestamp().Ok()
		return nil, response.StatusCode
	} else if response.StatusCode == 429 {
		memo.Warn(fmt.Sprintf("data=%s", string(data)))
	} else if response.StatusCode != 200 {
		memo.Error(fmt.Sprintf("data=%s", string(data))).Timestamp().Failed()
		return nil, response.StatusCode
	}

	var inv SteamInventory
	err := json.Unmarshal(data, &inv)
	if err != nil {
		memo.Error(
			fmt.Sprintf("Ошибка парсинга JSON: %s\n", err.Error()),
		).Timestamp().Failed()
		return nil, 0
	}

	if response.StatusCode == 429 {
		memo.Timestamp().Failed()
	} else {
		memo.Timestamp().Ok()
	}
	return &inv, response.StatusCode
}
