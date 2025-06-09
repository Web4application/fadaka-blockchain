npm install --save-dev hardhat @nomicfoundation/hardhat-toolbox dotenv

npx hardhat compile

# In a separate terminal, start local blockchain node
npx hardhat node

# Deploy contract to local node
npx hardhat run scripts/deploy.js --network localhost
