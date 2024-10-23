package bsc

type jsonRpcResponseData struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

type BscResponseData = jsonRpcResponseData
