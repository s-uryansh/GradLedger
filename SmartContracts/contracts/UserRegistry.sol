// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

//Smart Contract: UserRegistry

contract UserRegistry {
    address public owner;

    enum Role { None, Student, Alumni, Faculty }

    struct User {
        bool verified;
        Role role;
        string name;
        string rollNumber;
        string program;
        string major;
        string pictureHash; 
        // IPFS hash or URL for picture
    }

    mapping(address => User) public users;

    event UserVerified(address indexed user, Role role, string name, string rollNumber);

    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }
// deployer is owner
    constructor() {
        owner = msg.sender; 
    }

    function verifyUser(
        address _user,
        Role _role,
        string calldata _name,
        string calldata _rollNumber,
        string calldata _program,
        string calldata _major,
        string calldata _pictureHash
    ) external onlyOwner {
        require(_role != Role.None, "Invalid role");

        users[_user] = User({
            verified: true,
            role: _role,
            name: _name,
            rollNumber: _rollNumber,
            program: _program,
            major: _major,
            pictureHash: _pictureHash
        });

        emit UserVerified(_user, _role, _name, _rollNumber);
    }

    function isVerified(address _user) external view returns (bool) {
        return users[_user].verified;
    }

    function getUser(address _user) external view returns (
        bool verified,
        Role role,
        string memory name,
        string memory rollNumber,
        string memory program,
        string memory major,
        string memory pictureHash
    ) {
        User memory u = users[_user];
        return (u.verified, u.role, u.name, u.rollNumber, u.program, u.major, u.pictureHash);
    }
}
