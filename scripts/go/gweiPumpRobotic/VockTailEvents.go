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

    "time"
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

     contractAddress := common.HexToAddress("0xC14708B1faf3737602EA69b68C893beb58baB5a7")
     contract := connectContractAddress(client,contractAddress)
     fmt.Println("contract type object: ")
     fmt.Printf("%T",contract)
     fmt.Println("")

     randomDrinkValue := getRandomDrinkValue(contract)
     fmt.Println("randomDrinkValue:", randomDrinkValue)

     robotJobSinglePumpLeft40mL  := setupRobotJobSinglePumpLeft40mL()
     robotJobSinglePumpRight40mL := setupRobotJobSinglePumpRight40mL()
     robotJobDoublePump40mL      := setupRobotJobDoublePump40mL()

     fmt.Println("Listening for VockTail randomDrinkValue events...")
     SubscribeToEventsWithRobot(client, contractAddress, contract, robotJobSinglePumpLeft40mL, robotJobSinglePumpRight40mL, robotJobDoublePump40mL)

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

func SubscribeToEventsWithRobot(client *ethclient.Client, contractAddress common.Address, contract *vockTails.VockTails,robotJobSinglePumpLeft40mL *gobot.Robot, robotJobSinglePumpRight40mL *gobot.Robot,robotJobDoublePump40mL *gobot.Robot) {
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

          if randomDrinkValue.Cmp(big.NewInt(1)) == 0 {
              fmt.Println("Gobot: use robotJobSinglePumpLeft40mL")
              go robotJobSinglePumpLeft40mL.Start() //Use a go routine to keep event listener on when starting gobot.
          }
          if randomDrinkValue.Cmp(big.NewInt(2)) == 0 {
             fmt.Println("Gobot: use robotJobSinglePumpRight40mL")
              go robotJobSinglePumpRight40mL.Start() //Use a go routine to keep event listener on when starting gobot.
          }
          if randomDrinkValue.Cmp(big.NewInt(3)) == 0 {
              fmt.Println("Gobot: use robotJobDoublePump40mL")
              go robotJobDoublePump40mL.Start() //Use a go routine to keep event listener on when starting gobot.
          }

          fmt.Println("Listening for VockTails randomDrinkValue events...")

      }
  }

  return
}

func setupRobotJobDoublePump40mL() (robot *gobot.Robot){

  r := raspi.NewAdaptor()
  pump1 := gpio.NewLedDriver(r, "40") //Physical pin 40, GPIO 21.
  pump2 := gpio.NewLedDriver(r, "16") //Physical pin 16, GPIO 23.

  work := func() {

    fmt.Println("ON!");

    pump1.On();
    pump2.On();

    time.Sleep(22 * time.Second)

    fmt.Println("OFF!");
    pump1.Off();
    pump2.Off();

  }

  robot = gobot.NewRobot("blinkBot",
          []gobot.Connection{r},
          //[]gobot.Device{led},
          []gobot.Device{pump1,pump2},
          work,
  )

  return
}

func setupRobotJobSinglePumpLeft40mL() (robot *gobot.Robot){

  r := raspi.NewAdaptor()
  pump1 := gpio.NewLedDriver(r, "40") //Physical pin 40, GPIO 21.

  work := func() {

    fmt.Println("ON!");

    pump1.On();

    time.Sleep(44 * time.Second)

    fmt.Println("OFF!");
    pump1.Off();

  }

  robot = gobot.NewRobot("blinkBot",
          []gobot.Connection{r},
          //[]gobot.Device{led},
          []gobot.Device{pump1},
          work,
  )

  return
}

func setupRobotJobSinglePumpRight40mL() (robot *gobot.Robot){

  r := raspi.NewAdaptor()
  pump2 := gpio.NewLedDriver(r, "16") //Physical pin 16, GPIO 23.

  work := func() {

    fmt.Println("ON!");

    pump2.On();

    time.Sleep(44 * time.Second)

    fmt.Println("OFF!");
    pump2.Off();

  }

  robot = gobot.NewRobot("blinkBot",
          []gobot.Connection{r},
          //[]gobot.Device{led},
          []gobot.Device{pump2},
          work,
  )

  return
}
