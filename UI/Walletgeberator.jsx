import React, { useState } from 'react';

function Wallet() {
    const [address, setAddress] = useState('');

    const handleGenerate = () => {
        // Logic to generate wallet address
        const newAddress = '0x...' // Dummy address
        setAddress(newAddress);
    };

    return (
        <div>
            <h1>Wallet Generator</h1>
            <button onClick={handleGenerate}>Generate Wallet</button>
            {address && <p>Your new address: {address}</p>}
        </div>
    );
}

export default Wallet;
