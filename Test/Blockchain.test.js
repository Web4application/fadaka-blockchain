const CryptoUtils = require('./utils/CryptoUtils');

// Generate a key pair
const { publicKey, privateKey } = CryptoUtils.generateKeyPair();
console.log("Public Key:", publicKey);
console.log("Private Key:", privateKey);

// Create sample data
const data = { sender: "0x123", receiver: "0x456", amount: 50 };

// Generate hash
const dataHash = CryptoUtils.generateHash(JSON.stringify(data));
console.log("Data Hash:", dataHash);

// Sign the data
const signature = CryptoUtils.signData(privateKey, data);
console.log("Signature:", signature);

// Verify the signature
const isValid = CryptoUtils.verifySignature(publicKey, data, signature);
console.log("Is Signature Valid?", isValid);

// Encrypt data
const encrypted = CryptoUtils.encryptData("Sensitive data");
console.log("Encrypted Data:", encrypted);

// Decrypt data
const decrypted = CryptoUtils.decryptData(encrypted);
console.log("Decrypted Data:", decrypted);
