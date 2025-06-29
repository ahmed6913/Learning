// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    // State variable to store a number
    uint256 private storedNumber;

    // Function to set the number
    function set(uint256 _num) public {
        storedNumber = _num;
    }

    // Function to get the number
    function get() public view returns (uint256) {
        return storedNumber;
    }
} 