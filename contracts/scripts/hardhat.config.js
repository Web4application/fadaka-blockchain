require("@nomiclabs/hardhat-ethers");

module.exports = {
  solidity: "0.8.21",
  networks: {
    fadaka: {
      url: "https://rpc.fadaka.chain",  // Replace with your Fadaka RPC
      accounts: ["0xYOUR_PRIVATE_KEY"] // Never commit this key!
    }
  }
};
