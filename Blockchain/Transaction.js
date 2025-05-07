class Transaction {
    constructor(sender, receiver, amount, signature) {
        this.sender = sender;
        this.receiver = receiver;
        this.amount = amount;
        this.signature = signature; // Used to verify the transaction's authenticity
    }
}

module.exports = Transaction;
