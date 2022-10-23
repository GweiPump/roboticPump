//Step 1: Create abi file by running: solc --abi GweiPump.sol > GweiPump.abi
//Step 2: Create bin file by running: solc --bin GweiPump.sol > GweiPump.bin
//Step 3: Remove comments above the abi and bin files.
//Step 4: Generate Go contract interaction file by running:  abigen --bin=GweiPump.bin --abi=GweiPump.abi --pkg=gweiPump --out=GweiPump.go
//Step 5: Run: testGweiPumpEvents.go
package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "math/big"

    gweiPump "gweiPumpProject/contracts/GweiPump" //LOOK AT "go.mod" FOR YOUR RELATIVE PROJECT PATH TO FIND CONTRACT INTERFACE!

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/core/types"


        "time"
	//"fmt"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {

     // Use this endpoint when you are running your own node on a specific chain (no events)
     // client, chainID := clientSetup("http://localhost:8545")

     // Use this endpoint when you are running your own node on a specific chain (events allowed)
     // client, chainID := clientSetup("ws://localhost:8546")

     client, chainID := clientSetup(os.Getenv("mumbaiQuicknodeWSS"))

     fmt.Println("chainID: ", chainID)

     contractAddress := common.HexToAddress("0xd27759C36967E299ef16df8FAac24D4adb21665c")
     contract := connectContractAddress(client,contractAddress)
     fmt.Println("contract type object: ")
     fmt.Printf("%T",contract)
     fmt.Println("")

     isPumpFilled := getIsPumpFilled(contract)
     fmt.Println("isPumpFilled:", isPumpFilled)

     fmt.Println("Listening for GweiPump oilBought events...")
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

func connectContractAddress(client *ethclient.Client, contractAddress common.Address) (contract *gweiPump.GweiPump) {

  contract, err := gweiPump.NewGweiPump(contractAddress, client)
  if err != nil {
      log.Fatal(err)
  }
  return
}

func getIsPumpFilled(contract *gweiPump.GweiPump) (isPumpFilled *big.Int) {

  isPumpFilled, err := contract.IsPumpFilled(&bind.CallOpts{})
  if err != nil {
        log.Fatal(err)
  }
  return

}

func SubscribeToEvents(client *ethclient.Client, contractAddress common.Address, contract *gweiPump.GweiPump) {
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

          isPumpFilled, err := contract.IsPumpFilled(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
          }
          fmt.Println("isPumpFilled:", isPumpFilled)



 	r := raspi.NewAdaptor()
        //led := gpio.NewLedDriver(r, "38") //Physical pin 38, GPIO 20.
        pump1 := gpio.NewLedDriver(r, "40") //Physical pin 40, GPIO 21.
        pump2 := gpio.NewLedDriver(r, "16") //Physical pin 16, GPIO 23.        

        work := func() {
           //     gobot.Every(2*time.Second, func() {
	//		fmt.Println(time.Now().Unix());
	//		fmt.Println(time.Now().Unix()%4);
			
	//		if (time.Now().Unix()%4) == 2 { 
				fmt.Println("ON!");
				//led.On();
				pump1.On();
				pump2.On();
	
					
				time.Sleep(2 * time.Second)
	//		} else {
				fmt.Println("OFF!");
				//led.Off();
				pump1.Off();
				pump2.Off();
	//		}

          //      })
        }

        robot := gobot.NewRobot("blinkBot",
                []gobot.Connection{r},
                //[]gobot.Device{led},
                []gobot.Device{pump1,pump2},
                work,
        )

        robot.Start()
	  

          fmt.Println("Listening for GweiPump oilBought events...")

      }
  }

  return
}
