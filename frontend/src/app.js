import React, { useState } from 'react';
import './App.css';  // Import custom styles
import Wallet from './components/wallet';  // Import Wallet component

function App() {
  const [walletAddress, setWalletAddress] = useState(null); // State for storing wallet address
  const [loading, setLoading] = useState(false); // State to track loading

  // Function to handle wallet generation
  const handleGenerateWallet = async () => {
    setLoading(true);
    try {
      const response = await fetch('http://localhost:5000/generate_wallet'); // Backend URL for generating wallet
      const data = await response.json();
      setWalletAddress(data.walletAddress); // Update wallet address state with response
    } catch (error) {
      console.error('Error generating wallet:', error);
      alert('Failed to generate wallet.');
    } finally {
      setLoading(false); // Reset loading state after the request is completed
    }
  };

  return (
    <div className="App">
      <h1>Welcome to Fadaka Blockchain</h1>
      <button onClick={handleGenerateWallet} disabled={loading}>
        {loading ? 'Generating Wallet...' : 'Generate Wallet'}
      </button>

      {walletAddress && (
        <div className="wallet-info">
          <h2>Your Generated Wallet</h2>
          <p>Address: {walletAddress}</p>
        </div>
      )}
    </div>
  );
}

export default App;
