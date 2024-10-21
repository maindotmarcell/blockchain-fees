package bitcoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BitcoinAPIClient struct {
	apiURL string
}

func NewBitcoinAPIClient(apiURL string) *BitcoinAPIClient {
	return &BitcoinAPIClient{apiURL: apiURL}
}

func (c *BitcoinAPIClient) FetchFee() (interface{}, error) {
	resp, err := http.Get(c.apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching Bitcoin fee: %v", err)
	}
	defer resp.Body.Close()

	var data BitcoinResponseData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return data, nil
}
