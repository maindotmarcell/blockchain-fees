package solana

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type APIClient struct {
	apiURL string
}

func NewSolanaAPIClient() *APIClient {
	return &APIClient{apiURL: API_URL}
}

func (c *APIClient) FetchFee() (interface{}, error) {
	reqObj := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "getFees", // Changed method
		"params": []interface{}{
			map[string]string{
				"commitment": COMMITMENT,
			},
		},
		"id": 1,
	}

	jsonData, err := json.Marshal(reqObj)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(c.apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	// For debugging, print the actual response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data SolanaResponseData
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}
