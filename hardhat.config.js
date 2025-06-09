require("dotenv").config();
require("@nomiclabs/hardhat-ethers");
require("@nomiclabs/hardhat-etherscan"); // Optional if Fadaka supports verification
require("hardhat-contract-sizer");

module.exports = {
  solidity: "0.8.21",
  networks: {
    fadaka: {
      url: process.env.FADAKA_RPC_URL,
      accounts: [process.env.FADAKA_PRIVATE_KEY]
    }
  },
  etherscan: {
    apiKey: process.env.FADAKA_EXPLORER_API_KEY || "" // Optional
  },
  contractSizer: {
    runOnCompile: true,
    strict: false,
    only: []
  }
};
