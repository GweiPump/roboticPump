const ethers = require("ethers") // npm i ethers@5.7.2 https://github.com/smartcontractkit/full-blockchain-solidity-course-js/discussions/5139#discussioncomment-5444517

const rpcURL =  process.env.baseSepoliaWSS //Use WSS to get live event data instead of polling constantly,

const provider = new ethers.providers.WebSocketProvider(rpcURL)

const contractAddress = "0x5c9E9e129dE558eCd87DaDa56CA8b058B562ee62";
const contractABI = [{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"EmptyArgs","type":"error"},{"inputs":[],"name":"EmptySource","type":"error"},{"inputs":[],"name":"NoInlineSecrets","type":"error"},{"inputs":[],"name":"OnlyRouterCanFulfill","type":"error"},{"inputs":[{"internalType":"bytes32","name":"requestId","type":"bytes32"}],"name":"UnexpectedRequestID","type":"error"},{"inputs":[],"name":"etherNotSent","type":"error"},{"inputs":[],"name":"msgValueTooSmall","type":"error"},{"inputs":[],"name":"oraclePriceFeedZero","type":"error"},{"inputs":[],"name":"pumpNotFilled","type":"error"},{"inputs":[],"name":"upKeepNotNeeded","type":"error"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"id","type":"bytes32"}],"name":"RequestFulfilled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"id","type":"bytes32"}],"name":"RequestSent","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"requestId","type":"bytes32"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"},{"indexed":false,"internalType":"bytes","name":"response","type":"bytes"},{"indexed":false,"internalType":"bytes","name":"err","type":"bytes"}],"name":"Response","type":"event"},{"anonymous":false,"inputs":[],"name":"oilBought","type":"event"},{"anonymous":false,"inputs":[],"name":"updateWti","type":"event"},{"inputs":[],"name":"MAX_BPS","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"SCALE_FEE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"buyOil40Milliliters","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"getLatestEthUsdPennies","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getLatestWtiEth","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getWti40Milliliters","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"requestId","type":"bytes32"},{"internalType":"bytes","name":"response","type":"bytes"},{"internalType":"bytes","name":"err","type":"bytes"}],"name":"handleOracleFulfillment","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"isPumpFilled","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"lastWtiPriceCheckUnixTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"status","type":"uint256"}],"name":"ownerPumpFilledStatus","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"s_lastError","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"s_lastRequestId","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"s_lastResponse","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint64","name":"subscriptionId","type":"uint64"},{"internalType":"string[]","name":"args","type":"string[]"}],"name":"sendRequest","outputs":[{"internalType":"bytes32","name":"requestId","type":"bytes32"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"wtiUsdPenniesPriceOracle","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}];
const contractDeployed = new ethers.Contract(contractAddress, contractABI, provider);

async function setupConnection() {

  // const chainIdConnected = await web3.eth.getChainId();
  // console.log("chainIdConnected: "+ chainIdConnected)

  await pumpValueStatus();

  console.log("Waiting for a GweiPump oilBought event...")

}

async function pumpValueStatus() {

  const connectedNetworkObject = await provider.getNetwork();
  const chainIdConnected = connectedNetworkObject.chainId;
  console.log("chainIdConnected: "+ chainIdConnected)

  const isPumpFilledValue = await contractDeployed.isPumpFilled()
  console.log("isPumpFilled: " + isPumpFilledValue)

}

setupConnection();

contractDeployed.on("oilBought", (eventDetected) => {

  console.log("EVENT DETECTED! PUMP OIL [40mL]!")
  console.log(eventDetected)
  pumpValueStatus();  //Call the get function to get the most accurate present state for the value.
  //Pump logic goes here.
  console.log("Waiting for a GweiPump oilBought event...")
  
});
