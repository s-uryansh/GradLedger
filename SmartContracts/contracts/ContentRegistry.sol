// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

//Smart Contract: ContentRegistry

contract ContentRegistry {
    struct Content {
        address uploader;
        string cid;      // IPFS hash
        string title;    // Title of content
        bool isPublic;
    }

    Content[] public contents;

    mapping(address => uint[]) public userUploads;

    event ContentUploaded(address indexed uploader, uint contentId, string title, bool isPublic);

    function uploadContent(string calldata _cid, string calldata _title, bool _isPublic) external {
        contents.push(Content({
            uploader: msg.sender,
            cid: _cid,
            title: _title,
            isPublic: _isPublic
        }));

        uint contentId = contents.length - 1;
        userUploads[msg.sender].push(contentId);

        emit ContentUploaded(msg.sender, contentId, _title, _isPublic);
    }

    function getContent(uint _id) external view returns (address uploader, string memory cid, string memory title, bool isPublic) {
        require(_id < contents.length, "Invalid ID");
        Content memory c = contents[_id];
        return (c.uploader, c.cid, c.title, c.isPublic);
    }

    function getUserUploads(address _user) external view returns (uint[] memory) {
        return userUploads[_user];
    }

    function getTotalContents() external view returns (uint) {
        return contents.length;
    }
}
