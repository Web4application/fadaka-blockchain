import React, { useState } from 'react';

function Wallet() {
  const [address, setAddress] = useState('');
  const [loading, setLoading] = useState(false);

  // Handler to call backend API for wallet generation
  const handleGenerate = async () => {
    setLoading(true);
    try {
      // You would likely call your backend here to generate the wallet
      const response = await fetch('http://localhost:5000/generate_wallet'); // Assuming backend is running on port 5000
      const data = await response.json();
      setAddress(data.walletAddress); // Set the address to state
    } catch (error) {
      console.error('Error generating wallet:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <h2>Generate a new wallet address</h2>
      <button onClick={handleGenerate} disabled={loading}>
        {loading ? 'Generating...' : 'Generate Wallet'}
      </button>
      {address && (
        <div>
          <p>Your new wallet address:</p>
          <pre>{address}</pre> {/* Display wallet address */}
        </div>
      )}
    </div>
  );
}

export default Wallet;
