package bandprotocoltest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type BroadcastTransactionReq struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type BroadcastTransactionRes struct {
	TxHash string `json:"tx_hash"`
}

func BroadcastTransaction(symbol string, price, timestamp uint64) (BroadcastTransactionRes, error) {
	client := http.DefaultClient

	url := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"

	body := BroadcastTransactionReq{
		Symbol:    symbol,
		Price:     price,
		Timestamp: timestamp,
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return BroadcastTransactionRes{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return BroadcastTransactionRes{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return BroadcastTransactionRes{}, err
	}
	defer resp.Body.Close()

	res := new(BroadcastTransactionRes)
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return BroadcastTransactionRes{}, err
	}

	return *res, nil
}
