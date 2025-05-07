const Block = require('./Block');
const ProofOfWork = require('./ProofOfWork');

class Blockchain {
    constructor() {
        this.chain = [this.createGenesisBlock()];
        this.difficulty = 4; // Difficulty for mining
        this.pendingTransactions = [];
        this.miningReward = 100; // Reward for mining a block
    }

    createGenesisBlock() {
        return new Block(0, Date.now(), 'Genesis Block', '0');
    }

    getLatestBlock() {
        return this.chain[this.chain.length - 1];
    }

    addTransaction(transaction) {
        if (!transaction.sender || !transaction.receiver || !transaction.amount) {
            throw new Error('Invalid transaction');
        }
        this.pendingTransactions.push(transaction);
    }

    minePendingTransactions(miningRewardAddress) {
        const block = new Block(
            this.chain.length,
            Date.now(),
            this.pendingTransactions,
            this.getLatestBlock().hash
        );

        const proofOfWork = new ProofOfWork(this);
        proofOfWork.mineBlock(block);

        console.log('Mining complete.');
        this.chain.push(block);

        // Reward the miner
        this.pendingTransactions = [
            new Transaction(null, miningRewardAddress, this.miningReward, null),
        ];
    }

    getBalanceOfAddress(address) {
        let balance = 0;
        for (const block of this.chain) {
            for (const tx of block.data) {
                if (tx.receiver === address) {
                    balance += tx.amount;
                }
                if (tx.sender === address) {
                    balance -= tx.amount;
                }
            }
        }
        return balance;
    }

    isChainValid() {
        for (let i = 1; i < this.chain.length; i++) {
            const currentBlock = this.chain[i];
            const previousBlock = this.chain[i - 1];

            if (currentBlock.hash !== currentBlock.calculateHash()) {
                return false;
            }

            if (currentBlock.previousHash !== previousBlock.hash) {
                return false;
            }
        }
        return true;
    }
}

module.exports = Blockchain;
