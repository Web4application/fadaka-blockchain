const jwt = require('jsonwebtoken');

function generateToken(user) {
    return jwt.sign({ id: user.id, username: user.username }, 'your_jwt_secret', { expiresIn: '1h' });
}

// Example usage
const user = { id: 1, username: 'user1' };
const token = generateToken(user);
console.log('JWT Token:', token);
