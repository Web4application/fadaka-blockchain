const API_URL = 'http://localhost:5000'; // Replace with your backend URL if different

// Function to generate a new wallet address
const generateWallet = async () => {
  try {
    const response = await fetch(`${API_URL}/generate_wallet`, {
      method: 'GET', // GET request to generate wallet
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // If the response is not okay, throw an error
    if (!response.ok) {
      throw new Error('Failed to generate wallet');
    }

    // Parse the JSON response from the backend
    const data = await response.json();
    return data.walletAddress; // Assuming the backend sends { walletAddress: '0x...' }
  } catch (error) {
    console.error('Error generating wallet:', error);
    return null; // Return null in case of an error
  }
};

// Function to get the transaction history (example, can be extended)
const getTransactionHistory = async () => {
  try {
    const response = await fetch(`${API_URL}/transactions`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch transaction history');
    }

    const data = await response.json();
    return data.transactions; // Assuming the backend sends a list of transactions
  } catch (error) {
    console.error('Error fetching transactions:', error);
    return [];
  }
};

export { generateWallet, getTransactionHistory };
