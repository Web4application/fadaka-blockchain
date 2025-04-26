import React, { useState } from 'react';

function App() {
    const [wallet, setWallet] = useState(null);
    const [transactions, setTransactions] = useState([]);

    const createWallet = async () => {
        // Call backend API to create a new wallet
        const response = await fetch('/api/create-wallet');
        const data = await response.json();
        setWallet(data);
    };

    const sendTransaction = async (receiver, amount) => {
        // Call backend API to send a transaction
        const response = await fetch('/api/send-transaction', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ receiver, amount }),
        });
        const data = await response.json();
        setTransactions([...transactions, data]);
    };

    return (
        <div>
            <h1>fadaka Wallet</h1>
            <button onClick={createWallet}>Create Wallet</button>
            {wallet && <div>Public Key: {wallet.publicKey}</div>}
            <h2>Transactions</h2>
            <ul>
                {transactions.map((tx, index) => (
                    <li key={index}>{tx.sender} sent {tx.amount} to {tx.receiver}</li>
                ))}
            </ul>
        </div>
    );
}

export default App;
