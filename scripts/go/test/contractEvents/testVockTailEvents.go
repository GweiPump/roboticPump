//Step 1: Create abi file by running: solc --abi VockTails.sol > VockTails.abi
//Step 2: Create bin file by running: solc --bin VockTails.sol > VockTails.bin
//Step 3: Remove comments above the abi and bin files.
//Step 4: Generate Go contract interaction file by running:  abigen --bin=VockTails.bin --abi=VockTails.abi --pkg=vockTails --out=VockTails.go
//Step 5: Run: testVockTailEvents.go
package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "math/big"

    vockTails "gweiPumpProject/contracts/VockTails" //LOOK AT "go.mod" FOR YOUR RELATIVE PROJECT PATH TO FIND CONTRACT INTERFACE!

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"
)

func main() {

     // Use this endpoint when you are running your own node on a specific chain (no events)
     // client, chainID := clientSetup("http://localhost:8545")

     // Use this endpoint when you are running your own node on a specific chain (events allowed)
     // client, chainID := clientSetup("ws://localhost:8546")

     client, chainID := clientSetup(os.Getenv("mumbaiQuicknodeWSS"))

     fmt.Println("chainID: ", chainID)

     contractAddress := common.HexToAddress("0xC14708B1faf3737602EA69b68C893beb58baB5a7")
     contract := connectContractAddress(client,contractAddress)
     fmt.Println("contract type object: ")
     fmt.Printf("%T",contract)
     fmt.Println("")

     randomDrinkValue := getRandomDrinkValue(contract)
     fmt.Println("randomDrinkValue:", randomDrinkValue)

     fmt.Println("Listening for VockTail randomDrinkValue events...")
     SubscribeToEvents(client, contractAddress, contract)

}

func clientSetup(wssConnectionURL string) (client *ethclient.Client, chainID *big.Int) {

  client, err := ethclient.Dial(wssConnectionURL)
  if err != nil {
      log.Fatal(err)
  }

  chainID, err = client.NetworkID(context.Background())
  if err != nil {
     log.Fatal(err)
  }
  return
}

func connectContractAddress(client *ethclient.Client, contractAddress common.Address) (contract *vockTails.VockTails) {

  contract, err := vockTails.NewVockTails(contractAddress, client)
  if err != nil {
      log.Fatal(err)
  }
  return
}

func getRandomDrinkValue(contract *vockTails.VockTails) (isPumpFilled *big.Int) {

  isPumpFilled, err := contract.RandomDrinkValue(&bind.CallOpts{})
  if err != nil {
        log.Fatal(err)
  }
  return

}

func SubscribeToEvents(client *ethclient.Client, contractAddress common.Address, contract *vockTails.VockTails) {
  //Subscribe to events from smart contract address.
  query := ethereum.FilterQuery{
      Addresses: []common.Address{contractAddress},
  }

  logs := make(chan types.Log)
  sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
  if err != nil {
      log.Fatal(err)
  }

  for {
      select {
      case err := <-sub.Err():
          log.Fatal(err)
      case vLog := <-logs:
          fmt.Println("New Event Log:", vLog) // pointer to event log

          randomDrinkValue, err := contract.RandomDrinkValue(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
          }
          fmt.Println("randomDrinkValue:", randomDrinkValue)

          fmt.Println("Listening for VockTails randomDrinkValue events...")

      }
  }

  return
}
