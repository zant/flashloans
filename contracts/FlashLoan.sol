// SPDX-License-Identifier: agpl-3.0
pragma solidity ^0.8.4;

import {FlashLoanReceiverBase} from "./FlashLoanReceiverBase.sol";
import {ILendingPool} from "./interfaces/ILendingPool.sol";
import {IERC20} from "./interfaces/IERC20.sol";
import {SafeMath} from "./Libraries.sol";
import "./Ownable.sol";

contract FlashLoan is FlashLoanReceiverBase, Ownable {
    using SafeMath for uint256;
    bool public check = false;

    constructor(address _addressProvider)
        FlashLoanReceiverBase(_addressProvider)
    {
        check = true;
    }

    function executeOperation(
        address[] calldata assets,
        uint256[] calldata amounts,
        uint256[] calldata premiums,
        address initiator,
        bytes calldata params
    ) external override returns (bool) {
        check = false;
        uint256 amountOwing = amounts[0].add(premiums[0]);
        IERC20(assets[0]).approve(address(LENDING_POOL), amountOwing);

        return true;
    }

    function flashLoan() public {
        address receiverAddress = address(this);
        address onBehalfOf = address(this);

        uint256 amount = 10000000000000000;

        address[] memory assets = new address[](1);
        assets[0] = address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = amount;

        uint256[] memory modes = new uint256[](1);
        modes[0] = 0;
        bytes memory params = "";
        uint16 referralCode = 0;


        LENDING_POOL.flashLoan(
            receiverAddress,
            assets,
            amounts,
            modes,
            onBehalfOf,
            params,
            referralCode
        );
    }
}
