//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract RabbitCollectible is ERC721 {
    constructor() ERC721("RabbitCollectible", "RBO"){
    }

    function Mint(address to, uint256 tokenId, bytes memory data ) public {
        _safeMint(to, tokenId, data);
    }
}
