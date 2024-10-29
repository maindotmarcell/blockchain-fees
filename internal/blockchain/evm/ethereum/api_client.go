package ethereum

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type EthApiClient struct {
	apiURL string
}

func NewEthApiClient() *EthApiClient {
	return &EthApiClient{apiURL: API_URL}
}

func (c *EthApiClient) FetchFee() (interface{}, error) {
	reqObj := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_maxPriorityFeePerGas",
		"params":  []interface{}{},
		"id":      1,
	}
	jsonData, err := json.Marshal(reqObj)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(c.apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data EthResponseData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
