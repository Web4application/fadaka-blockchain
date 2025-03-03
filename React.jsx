import React, { useState } from 'react';
import { View, Text, Button } from 'react-native';

const Wallet = () => {
    const [address, setAddress] = useState('');

    const handleGenerate = () => {
        const newAddress = '0x...'; // Dummy address
        setAddress(newAddress);
    };

    return (
        <View>
            <Text>Wallet Generator</Text>
            <Button title="Generate Wallet" onPress={handleGenerate} />
            {address && <Text>Your new address: {address}</Text>}
        </View>
    );
};

export default Wallet;
