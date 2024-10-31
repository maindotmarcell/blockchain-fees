package solana

type SolanaResponseData struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			APIVersion string `json:"apiVersion"`
			Slot       int64  `json:"slot"`
		} `json:"context"`
		Value struct {
			Blockhash     string `json:"blockhash"`
			FeeCalculator struct {
				LamportsPerSignature uint64 `json:"lamportsPerSignature"`
			} `json:"feeCalculator"`
		} `json:"value"`
	} `json:"result"`
	ID int `json:"id"`
}
