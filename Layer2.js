// Import necessary libraries
const ethers = require('ethers');

// Initialize provider and signer
const provider = new ethers.providers.JsonRpcProvider('YOUR_INFURA_API_URL');
const signer = provider.getSigner();

// Define state channel contract
const stateChannelContract = new ethers.Contract(
    'CONTRACT_ADDRESS',
    ['function deposit() public payable', 'function close(uint256 value, bytes memory signature) public'],
    signer
);

// Function to deposit into the state channel
async function deposit(amount) {
    const tx = await stateChannelContract.deposit({ value: ethers.utils.parseEther(amount) });
    await tx.wait();
    console.log('Deposited:', amount);
}

// Example usage
deposit('0.1');
