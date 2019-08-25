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

            bytes32 r;
            bytes32 s;
            uint8 v;
            address identity;
        
            assembly {
                mstore(_datum, sub(mload(_datum), 65))
                identity := shr(96, mload(add(_datum, 35)))

                r := mload(add(_datum, mload(_datum)))
                s := mload(add(_datum, add(mload(_datum), 32)))
                v := byte(0, mload(add(_datum, add(mload(_datum), 64))))
                
            }
            
            require(ecrecover(keccak256(_datum), v, r, s) == identity);
             
            assembly {
                pop(
                    call(
                        shr(232, mload(add(_datum, 32))),
                        identity,
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
