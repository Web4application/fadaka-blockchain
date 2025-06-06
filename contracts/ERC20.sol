// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract fadakaToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("fadakaToken", "FDK") {
        _mint(msg.sender, initialSupply);
    }
}
