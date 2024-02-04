## Transaction Broadcasting and Monitoring Client Module

#### How to

how to use and integrate modules

- Install

  ```
  go get github.com/famesensor/band-protocol-test
  ```

- Broadcast Transaction
  - `BroadcastTransaction(symbol, price, timestamp)`
    - Arguments
      - symbol: Transaction symbol, e.g., "BTC"
      - price: Symbol price, e.g., 100000
      - timestamp: Timestamp of price retrieval
    - Return
      > BroadcastTransactionRes, Error
- Transaction Status Monitoring
  - `TransactionStatusMonitor(tx_hash)`
    - Arguments
      - tx_hash: response value from broadcast transaction
    - Return
      > TransactionStatusMonitorRes, Error

#### Example

The example show how to use and integrate modules in your script

```Go
package main

import (
 "log"

 bandprotocoltest "github.com/famesensor/band-protocol-test"
)

func main() {
 // call broadcast transaction
 // I try call api in code and postman
 // api is not work
 // response 404 Not Found, The requested resource could not be found, error from postman
 resBroadcast, err := bandprotocoltest.BroadcastTransaction("ETH", 4500, 1678912345)
 if err != nil {
  log.Fatal("error.BroadcastTransaction: ", err)
 }

 log.Println("tx_hash: ", resBroadcast.TxHash)

 // call transaction status monitor
 // api is work!!
 resTransaction, err := bandprotocoltest.TransactionStatusMonitor(resBroadcast.TxHash)
 if err != nil {
  log.Fatal("error.TransactionStatusMonitor: ", err)
 }

 log.Println("tx_status: ", resTransaction.TxStatus)
}
```
