const Web3 = require('web3') //HIDE KEYS WITH "Linux Environment Variables" https://www.youtube.com/watch?v=himEMfYQJ1w "vim .bashrc" > "i" > update > "esc" > ":w" enter
const rpcURL =  process.env.mumbaiQuicknodeWSS //Use WSS to get live event data instead of polling constantly,
const web3 = new Web3(rpcURL)

const gweiPumpAddress = '0xdf6110fE578B98DEF32d5066fE3538a646C9A48B'
const gweiPumpABI = [{"inputs":[],"name":"msgValueTooSmall","type":"error"},{"inputs":[],"name":"notOwner","type":"error"},{"inputs":[],"name":"oraclePriceFeedZero","type":"error"},{"inputs":[],"name":"pumpNotFilled","type":"error"},{"inputs":[],"name":"upKeepNotNeeded","type":"error"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"id","type":"bytes32"}],"name":"ChainlinkCancelled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"id","type":"bytes32"}],"name":"ChainlinkFulfilled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"id","type":"bytes32"}],"name":"ChainlinkRequested","type":"event"},{"anonymous":false,"inputs":[],"name":"oilBought","type":"event"},{"anonymous":false,"inputs":[],"name":"updateWti","type":"event"},{"inputs":[],"name":"buyOil40Milliliters","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"chainlinkNodeRequestWtiPrice","outputs":[{"internalType":"bytes32","name":"requestId","type":"bytes32"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"","type":"bytes"}],"name":"checkUpkeep","outputs":[{"internalType":"bool","name":"upkeepNeeded","type":"bool"},{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"_requestId","type":"bytes32"},{"internalType":"uint256","name":"memoryWtiPriceOracle","type":"uint256"}],"name":"fulfill","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"manualUpKeep","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"status","type":"uint256"}],"name":"ownerPumpFilledStatus","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"","type":"bytes"}],"name":"performUpkeep","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"erc20LINK","outputs":[{"internalType":"contractERC20TokenContract","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"feeThreeThousandthPercent","outputs":[{"internalType":"int256","name":"","type":"int256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getLatestMaticUsd","outputs":[{"internalType":"int256","name":"","type":"int256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getLatestWtiMatic","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"isPumpFilled","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"lastWtiPriceCheckUnixTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"Owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"Wti40Milliliters","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"WtiPriceOracle","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]
const deployedGweiPump = new web3.eth.Contract(gweiPumpABI, gweiPumpAddress)

async function setupConnection() {

  const chainIdConnected = await web3.eth.getChainId();
  console.log("chainIdConnected: "+ chainIdConnected)

  await pumpValueStatus();

  console.log("Waiting for a GweiPump oilBought event...")

}

async function pumpValueStatus() {

  const isPumpFilledValue = await deployedGweiPump.methods.isPumpFilled().call()
  console.log("isPumpFilled: " + isPumpFilledValue)

}

setupConnection();

deployedGweiPump.events.oilBought({ //Subscribe to event.
     fromBlock: 'latest'
 }, function(error, eventResult){})
 .on('data', function(eventResult){
   console.log("EVENT DETECTED! PUMP OIL [40mL]!")
   pumpValueStatus();  //Call the get function to get the most accurate present state for the value.
   //Pump logic goes here.
   console.log("Waiting for a GweiPump oilBought event...")
   })
 .on('changed', function(eventResult){
     // remove event from local database
 })
 .on('error', console.error);
