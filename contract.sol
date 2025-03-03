pragma solidity ^0.8.0;

contract Governance {
    mapping(address => uint256) public votes;
    address[] public candidates;

    event VoteCast(address indexed voter, address indexed candidate);

    function vote(address candidate) public {
        votes[candidate] += 1;
        emit VoteCast(msg.sender, candidate);
    }

    function getVotes(address candidate) public view returns (uint256) {
        return votes[candidate];
    }
}
