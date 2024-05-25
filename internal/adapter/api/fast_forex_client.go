package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/shopspring/decimal"
)

type FastForexClientImpl struct {
	apiKey string
	client *resty.Client
}

func NewFastForexClient(apiKey, mainURL string) *FastForexClientImpl {
	client := resty.New()

	client.SetQueryParam("api_key", apiKey)
	client.BaseURL = mainURL

	return &FastForexClientImpl{
		apiKey: apiKey,
		client: client,
	}
}

type FetchOneRequest struct {
	From string `url:"from"`
	To   string `url:"to"`
}

type FetchOneResponse struct {
	Base    string                     `json:"base"`
	Result  map[string]decimal.Decimal `json:"result"`
	Updated string                     `json:"updated"`
	Ms      int                        `json:"ms"`
}

func (ffc *FastForexClientImpl) GetRates(ctx context.Context,
	baseCurrency string, symbol string,
) (*FetchOneResponse, error) {
	const fetchOne = "/fetch-one"

	resp, err := ffc.client.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"from": baseCurrency,
			"to":   symbol,
		}).
		Get(fetchOne)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rates: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch rate for %s: %v", symbol, resp.String())
	}

	var response *FetchOneResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return response, nil
}
