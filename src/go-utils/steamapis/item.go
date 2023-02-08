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

type ItemService struct {
	httpClient *http.Client
	endpoint   string

	hideLogs   bool
	hitsPerMin int64
	lastHitAt  int64
	hitsMu     sync.Mutex
}

func NewItemService(apiKey string, hitsPerMin int64, hideLogs bool) *ItemService {
	return &ItemService{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		endpoint: strings.Replace(
			"https://api.steamapis.com/market/item/{AppID}/{MarketHashName}?api_key={ApiKey}",
			"{ApiKey}",
			apiKey,
			1,
		),
		hitsPerMin: hitsPerMin,
		hideLogs:   hideLogs,
	}
}

func (s *ItemService) GetItem(marketHashName string, appID string) (*MarketItem, int) {
	url := strings.NewReplacer("{MarketHashName}", marketHashName, "{AppID}", appID).Replace(s.endpoint)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetItem(%s, %s): %s\n", marketHashName, appID, err.Error())
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

func (s *ItemService) makeRequest(request *http.Request) (*MarketItem, int) {
	memo := logger.NewMemo("SteamAPIs->GetItem")
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

	if response.StatusCode == 400 {
		memo.Warn("Предмет не найден").Timestamp().Ok()
		return nil, response.StatusCode
	} else if response.StatusCode == 429 {
		memo.Warn(fmt.Sprintf("data=%s", string(data)))
	} else if response.StatusCode != 200 {
		memo.Error(
			fmt.Sprintf("data=%s", string(data)),
		).Timestamp().Failed()
		return nil, response.StatusCode
	}

	var inv MarketItem
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
