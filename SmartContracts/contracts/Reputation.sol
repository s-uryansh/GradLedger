// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

//Smart Contract: Reputation

contract Reputation {
    address public owner;
    address public operator; // Mentorship contract

    mapping(address => uint256) private _score;

    event OperatorUpdated(address indexed operator);
    event ReputationChanged(address indexed mentor, int256 delta, uint256 newScore);

    modifier onlyOwner() { require(msg.sender == owner, "Not owner"); _; }
    modifier onlyOperator() { require(msg.sender == operator, "Not operator"); _; }

    constructor() { owner = msg.sender; }

    function setOperator(address _op) external onlyOwner {
        operator = _op;
        emit OperatorUpdated(_op);
    }

    function add(address mentor, uint256 amount) external onlyOperator {
        uint256 s = _score[mentor] + amount;
        _score[mentor] = s;
        emit ReputationChanged(mentor, int256(amount), s);
    }

    function sub(address mentor, uint256 amount) external onlyOperator {
        uint256 s = _score[mentor];
        uint256 ns = amount >= s ? 0 : s - amount;
        _score[mentor] = ns;
        emit ReputationChanged(mentor, -int256(amount), ns);
    }

    function scoreOf(address user) external view returns (uint256) {
        return _score[user];
    }
}
