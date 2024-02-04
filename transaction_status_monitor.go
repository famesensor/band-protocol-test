package bandprotocoltest

import (
	"encoding/json"
	"net/http"
)

type TransactionStatusMonitorRes struct {
	TxStatus string `json:"tx_status"`
}

func TransactionStatusMonitor(txHash string) (TransactionStatusMonitorRes, error) {
	client := http.DefaultClient

	url := "https://mock-node-wgqbnxruha-as.a.run.app/check/" + txHash

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return TransactionStatusMonitorRes{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return TransactionStatusMonitorRes{}, err
	}
	defer resp.Body.Close()

	res := new(TransactionStatusMonitorRes)
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return TransactionStatusMonitorRes{}, err
	}

	return *res, nil
}
