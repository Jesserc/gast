// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract CurrentYear {
    uint256 year;

    constructor() {
        year = 2024;
    }

    function getYear() external view returns (uint256) {
        return year;
    }

    function updateYear(uint256 _year) external {
        year = _year;
    }
}
