package models

type GetNodeInfoResponse struct {
	BlockNumber              string   `json:"block_number"`
	NodeNumber               int      `json:"node_number"`
	TransactionNumber        string   `json:"transaction_number"`
	PendingTransactionNumber int      `json:"pending_transaction_number"`
	NodeInfoList             []string `json:"node_info_list"`
}
