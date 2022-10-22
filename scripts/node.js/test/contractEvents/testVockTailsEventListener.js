const Web3 = require('web3') //HIDE KEYS WITH "Linux Environment Variables" https://www.youtube.com/watch?v=himEMfYQJ1w "vim .bashrc" > "i" > update > "esc" > ":w" enter
const rpcURL =  process.env.mumbaiQuicknodeWSS //Use WSS to get live event data instead of polling constantly,
const web3 = new Web3(rpcURL)

const vockTailsAddress = "0xC14708B1faf3737602EA69b68C893beb58baB5a7"
const vockTailsAbi =[{"inputs":[{"internalType":"address","name":"have","type":"address"},{"internalType":"address","name":"want","type":"address"}],"name":"OnlyCoordinatorCanFulfill","type":"error"},{"anonymous":false,"inputs":[],"name":"drinkVRF","type":"event"},{"inputs":[{"internalType":"uint256","name":"requestId","type":"uint256"},{"internalType":"uint256[]","name":"randomWords","type":"uint256[]"}],"name":"rawFulfillRandomWords","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"requestRandomWords","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"COORDINATOR","outputs":[{"internalType":"contractVRFCoordinatorV2Interface","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"LINKTOKEN","outputs":[{"internalType":"contractLinkTokenInterface","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"randomDrinkValue","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"requestId","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]
const deployedVockTails = new web3.eth.Contract(vockTailsAbi, vockTailsAddress)

async function setupConnection() {

  const chainIdConnected = await web3.eth.getChainId();
  console.log("chainIdConnected: "+ chainIdConnected)

  await drinkValueStatus();

  console.log("Waiting for a VockTails drinkVRF event...")

}

async function drinkValueStatus() {

  const randomDrinkValueContract = await deployedVockTails.methods.randomDrinkValue().call()
  console.log("randomDrinkValue: " + randomDrinkValueContract)

}

setupConnection();

deployedVockTails.events.drinkVRF({ //Subscribe to event.
     fromBlock: 'latest'
 }, function(error, eventResult){})
 .on('data', function(eventResult){
   console.log("EVENT DETECTED! PUMP VOCKTAIL [40mL]!")
   drinkValueStatus();  //Call the get function to get the most accurate present state for the value.
   //Pump logic goes here.
   console.log("Waiting for a Vocktails drinkVRF event...")
   })
 .on('changed', function(eventResult){
     // remove event from local database
 })
 .on('error', console.error);
