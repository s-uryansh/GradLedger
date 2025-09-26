// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IUserRegistry {
    // Role: 0 None, 1 Student, 2 Alumni, 3 Faculty
    function isVerified(address _user) external view returns (bool);
    function getUserRole(address _user) external view returns (uint8);
}

interface IReputation {
    function add(address mentor, uint256 amount) external;
    function sub(address mentor, uint256 amount) external;
}

// Smart Contract: Mentorship
contract Mentorship {
    enum Status { Requested, Accepted, Completed, Canceled }

    struct Session {
        address student;
        address mentor;
        Status status;
        bool feedbackGiven;
    }

    address public owner;
    IUserRegistry public userRegistry;
    IReputation public reputation;

    Session[] public sessions;

    event Requested(uint indexed id, address indexed student, address indexed mentor);
    event Accepted(uint indexed id, address indexed mentor);
    event Completed(uint indexed id, address indexed mentor);
    event Canceled(uint indexed id, address indexed by);
    event Feedback(uint indexed id, address indexed student, address indexed mentor, bool upvote);

    modifier onlyOwner() { require(msg.sender == owner, "Not owner"); _; }

    constructor(address _userRegistry, address _reputation) {
        owner = msg.sender;
        userRegistry = IUserRegistry(_userRegistry);
        reputation   = IReputation(_reputation);
    }

    function request(address mentor) external returns (uint id) {
        require(userRegistry.isVerified(msg.sender), "Student not verified");
        require(userRegistry.isVerified(mentor), "Mentor not verified");

        uint8 mRole = userRegistry.getUserRole(mentor);
        require(mRole == 2 || mRole == 3, "Mentor must be Alumni/Faculty");

        sessions.push(Session({
            student: msg.sender,
            mentor: mentor,
            status: Status.Requested,
            feedbackGiven: false
        }));
        id = sessions.length - 1;
        emit Requested(id, msg.sender, mentor);
    }

    function accept(uint id) external {
        Session storage s = sessions[id];
        require(msg.sender == s.mentor, "Only mentor");
        require(s.status == Status.Requested, "Bad state");
        s.status = Status.Accepted;
        emit Accepted(id, s.mentor);
    }

    function complete(uint id) external {
        Session storage s = sessions[id];
        require(msg.sender == s.mentor, "Only mentor");
        require(s.status == Status.Accepted, "Bad state");
        s.status = Status.Completed;
        emit Completed(id, s.mentor);
    }

    function cancel(uint id) external {
        Session storage s = sessions[id];
        require(msg.sender == s.student || msg.sender == s.mentor || msg.sender == owner, "Not allowed");
        require(s.status == Status.Requested || s.status == Status.Accepted, "Bad state");
        s.status = Status.Canceled;
        emit Canceled(id, msg.sender);
    }

    function giveFeedback(uint id, bool upvote) external {
        Session storage s = sessions[id];
        require(msg.sender == s.student, "Only student");
        require(s.status == Status.Completed, "Not completed");
        require(!s.feedbackGiven, "Already given");

        s.feedbackGiven = true;

        if (upvote) {
            reputation.add(s.mentor, 1);
        } else {
            reputation.sub(s.mentor, 1);
        }
        emit Feedback(id, s.student, s.mentor, upvote);
    }

    function getSession(uint id) external view returns (
        address student, address mentor, Status status, bool feedbackGiven
    ) {
        Session memory s = sessions[id];
        return (s.student, s.mentor, s.status, s.feedbackGiven);
    }

    function totalSessions() external view returns (uint) {
        return sessions.length;
    }
}
