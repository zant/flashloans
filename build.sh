solc --abi --bin --overwrite contracts/FlashLoan.sol -o build && \
abigen --bin=./build/FlashLoan.bin --abi=./build/FlashLoan.abi --pkg=flashLoan --type=FlashLoan --out=abigen/FlashLoan/FlashLoan.go
