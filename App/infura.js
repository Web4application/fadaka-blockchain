import Web3 from 'web3';

const web3 = new Web3(new Web3.providers.HttpProvider('https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID'));

async function getBlockNumber() {
  const blockNumber = await web3.eth.getBlockNumber();
  console.log('Current block number:', blockNumber);
}

export default getBlockNumber;
