package tonapi

type AddressInformationResponse struct {
	Ok     bool   `json:"ok"`
	Error  string `json:"error"`
	Code   int    `json:"code"`
	Result struct {
		Type              string `json:"@type"`
		Balance           int    `json:"balance"`
		Code              string `json:"code"`
		Data              string `json:"data"`
		LastTransactionId struct {
			Type string `json:"@type"`
			Lt   string `json:"lt"`
			Hash string `json:"hash"`
		} `json:"last_transaction_id"`
		BlockId struct {
			Type      string `json:"@type"`
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
			RootHash  string `json:"root_hash"`
			FileHash  string `json:"file_hash"`
		} `json:"block_id"`
		FrozenHash string `json:"frozen_hash"`
		SyncUtime  int    `json:"sync_utime"`
		Extra      string `json:"@extra"`
		State      string `json:"state"`
	} `json:"result"`
}
