// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/*
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

interface IGoTest {
    function mint(address to) external;
}

contract GoTest is ERC721, IGoTest {
    uint256 private id;

    constructor() ERC721("GoTest", "GO") {}

    function mint(address to) public virtual override {
        _safeMint(to, id);
        id++;
    }
}
*/

interface IGoTest {
    function mint(address to) external;
}

