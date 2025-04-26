// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Import ERC-20 standard from OpenZeppelin
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// FadakaToken contract inheriting from ERC20
contract FadakaToken is ERC20 {
    // The contract owner (creator)
    address public owner;

    // Constructor to set the initial supply
    constructor(uint256 initialSupply) ERC20("Fadaka", "FDK") {
        _mint(msg.sender, initialSupply); // Mint the initial supply to the deployer's address
        owner = msg.sender; // Set the contract deployer as the owner
    }

    // Function to mint new tokens
    function mint(address to, uint256 amount) external {
        require(msg.sender == owner, "Only the owner can mint tokens");
        _mint(to, amount); // Mint new tokens to the specified address
    }

    // Function to burn tokens
    function burn(uint256 amount) external {
        _burn(msg.sender, amount); // Burn tokens from the caller's balance
    }

    // Function to check the balance of an address
    function getBalance(address account) external view returns (uint256) {
        return balanceOf(account);
    }

    // A function to pause or stop token transfers (example)
    bool public paused = false;

    modifier whenNotPaused() {
        require(!paused, "Token transfers are paused");
        _;
    }

    // Overriding the _transfer function to include the pause feature
    function _transfer(address from, address to, uint256 amount) internal override whenNotPaused {
        super._transfer(from, to, amount);
    }

    // Function to pause or unpause the contract
    function togglePause() external {
        require(msg.sender == owner, "Only the owner can pause/unpause");
        paused = !paused;
    }
}
