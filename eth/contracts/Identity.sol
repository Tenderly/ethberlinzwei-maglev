pragma solidity ^0.5.1;

import "./ERC725.sol";

contract ProxyAccount is ERC725 {

    uint256 constant OPERATION_CALL = 0;
    uint256 constant OPERATION_CREATE = 1;
    // bytes32 constant KEY_OWNER = 0x0000000000000000000000000000000000000000000000000000000000000000;

    mapping(bytes32 => bytes) store;

    address public owner;

    constructor(address _owner) public {
        owner = _owner;
    }


    modifier onlyOwner() {
        require(msg.sender == owner, "only-owner-allowed");
        _;
    }

    function () external payable {}

    function changeOwner(address _owner)
    external
    onlyOwner
    {
        owner = _owner;
        emit OwnerChanged(owner);
    }

    function getData(bytes32 _key)
    external
    view
    returns (bytes memory _value)
    {
        return store[_key];
    }
    
    function getBalance()
    external
    view
    returns (uint256 _value)
    {
        return address(this).balance;
    }

    function setData(bytes32 _key, bytes calldata _value)
    external
    onlyOwner
    {
        store[_key] = _value;
        emit DataChanged(_key, _value);
    }

    function execute(bytes memory _data)
    public
    onlyOwner
    {
        assembly {
            if eq(shr(248, mload(add(_data, 32))), 0) {
                pop(
                    call(
                        gas,
                        shr(96, mload(add(_data, 33))),
                        shr(160, mload(add(_data, 53))),
                        add(_data, 65),
                        sub(mload(_data), 33),
                        0,
                        0
                    )
                )

                return(0, 0)
            }
            
            pop(
                create(
                    0,
                    add(_data, 0x21),
                    mload(_data)
                )
            )
            
        }
    }
}