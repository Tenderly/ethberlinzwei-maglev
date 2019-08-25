pragma solidity ^0.5.1;

pragma experimental ABIEncoderV2;

import "./ERC725.sol";

contract Trebuchet {
    address public owner;

    constructor() public {
        owner = msg.sender;
    }


    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function execute(
        bytes[] memory _data
    ) public payable onlyOwner {
        for (uint i = 0; i < _data.length; i++) {
            bytes memory _datum = _data[i];

            assembly {
                pop(
                    call(
                        shr(232, mload(add(_datum, 32))),
                        shr(96, mload(add(_datum, 35))),
                        0,
                        add(_datum, 55),
                        sub(mload(_datum), 23),
                        0,
                        0
                    )
                )
            }
        }
    }
}
