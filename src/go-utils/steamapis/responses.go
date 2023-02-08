package steamapis

type SteamInventory struct {
	Type   string `json:"type"`
	Assets []struct {
		AppID      int    `json:"appid"`
		ContextID  string `json:"contextid"`
		AssetID    string `json:"assetid"`
		ClassID    string `json:"classid"`
		InstanceID string `json:"instanceid"`
		Amount     string `json:"amount"`
	} `json:"assets"`
	Descriptions []struct {
		AppID           int     `json:"appid"`
		ClassID         string  `json:"classid"`
		InstanceID      string  `json:"instanceid"`
		Currency        float64 `json:"currency"`
		BackgroundColor string  `json:"background_color"`
		IconUrl         string  `json:"icon_url"`
		IconUrlLarge    string  `json:"icon_url_large"`
		Descriptions    []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"descriptions"`
		Tradable int `json:"tradable"`
		Actions  []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"actions"`
		Name           string `json:"name"`
		NameColor      string `json:"name_color"`
		Type           string `json:"type"`
		MarketName     string `json:"market_name"`
		MarketHashName string `json:"market_hash_name"`
		MarketActions  []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"market_actions"`
		Commodity                 int `json:"commodity"`
		MarketTradableRestriction int `json:"market_tradable_restriction"`
		Marketable                int `json:"marketable"`
		Tags                      []struct {
			Category              string `json:"category"`
			InternalName          string `json:"internal_name"`
			LocalizedCategoryName string `json:"localized_category_name"`
			LocalizedTagName      string `json:"localized_tag_name"`
		} `json:"tags"`
	} `json:"descriptions"`
	TotalInventoryCount int `json:"total_inventory_count"`
	Success             int `json:"success"`
	Rwgrsn              int `json:"rwgrsn"`
}

type MarketItem struct {
	Type                  string          `json:"type"`
	NameID                int             `json:"nameID"`
	AppID                 int             `json:"appID"`
	MarketName            string          `json:"market_name"`
	MarketHashName        string          `json:"market_hash_name"`
	Description           interface{}     `json:"description"`
	Url                   string          `json:"url"`
	Image                 string          `json:"image"`
	BorderColor           string          `json:"border_color"`
	MedianAvgPrices15Days [][]interface{} `json:"median_avg_prices_15days"`
	Histogram             struct {
		SellOrderArray []struct {
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
		} `json:"sell_order_array"`
	} `json:"histogram"`
	SellOrderSummary struct {
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	} `json:"sell_order_summary"`
	BuyOrderArray   []interface{} `json:"buy_order_array"`
	BuyOrderSummary struct {
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	} `json:"buy_order_summary"`
	HighestBuyOrder interface{}     `json:"highest_buy_order"`
	LowestSellOrder float64         `json:"lowest_sell_order"`
	BuyOrderGraph   []interface{}   `json:"buy_order_graph"`
	SellOrderGraph  [][]interface{} `json:"sell_order_graph"`
	GraphMaxY       int             `json:"graph_max_y"`
	GraphMinX       float64         `json:"graph_min_x"`
	GraphMaxX       float64         `json:"graph_max_x"`
	PricePrefix     string          `json:"price_prefix"`
	PriceSuffix     string          `json:"price_suffix"`
	Assets          struct {
		Descriptions []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"descriptions"`
		Actions []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"actions"`
		Type string `json:"type"`
	} `json:"assets"`
	AssetInfo struct {
		NameColor       string `json:"name_color"`
		BackgroundColor string `json:"background_color"`
		Type            string `json:"type"`
		Tradable        bool   `json:"tradable"`
		Marketable      bool   `json:"marketable"`
		Commodity       bool   `json:"commodity"`
		Descriptions    []struct {
			Type    string      `json:"type"`
			Value   string      `json:"value"`
			AppData interface{} `json:"app_data"`
		} `json:"descriptions"`
		Actions []struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"actions"`
		MarketActions []struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"market_actions"`
		OwnerActions []interface{} `json:"owner_actions"`
		Tags         []struct {
			InternalName string `json:"internal_name"`
			Name         string `json:"name"`
			Category     string `json:"category"`
			CategoryName string `json:"category_name"`
		} `json:"tags"`
	} `json:"asset_info"`
	UpdateAt int64 `json:"update_at"`
}
