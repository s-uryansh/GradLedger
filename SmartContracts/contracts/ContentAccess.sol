// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IContentRegistry {
    function getContent(uint _id) external view returns (
        address uploader,
        string memory cid,
        string memory title,
        bool isPublic
    );
}

//Smart Contract: ContentAccess
contract ContentAccess {
    IContentRegistry public contentRegistry;

    // contentId => viewer => allowed
    mapping(uint => mapping(address => bool)) public isWhitelisted;

    event Granted(uint indexed contentId, address indexed viewer);
    event Revoked(uint indexed contentId, address indexed viewer);

    constructor(address _contentRegistry) {
        contentRegistry = IContentRegistry(_contentRegistry);
    }

    function grant(uint contentId, address viewer) external {
        (address uploader,, ,) = contentRegistry.getContent(contentId);
        require(msg.sender == uploader, "Only uploader");
        isWhitelisted[contentId][viewer] = true;
        emit Granted(contentId, viewer);
    }

    function revoke(uint contentId, address viewer) external {
        (address uploader,, ,) = contentRegistry.getContent(contentId);
        require(msg.sender == uploader, "Only uploader");
        isWhitelisted[contentId][viewer] = false;
        emit Revoked(contentId, viewer);
    }

    function canAccess(uint contentId, address viewer) external view returns (bool) {
        (address uploader,, ,bool pub_) = contentRegistry.getContent(contentId);
        if (pub_) return true;
        if (viewer == uploader) return true;
        return isWhitelisted[contentId][viewer];
    }
}
