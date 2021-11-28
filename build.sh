mkdir -p abigen
mkdir -p abigen/FlashLoan
mkdir -p abigen/ERC20
solc --abi --bin --overwrite contracts/FlashLoan.sol -o build && \
abigen --bin=./build/FlashLoan.bin --abi=./build/FlashLoan.abi --pkg=flashLoan --type=FlashLoan --out=abigen/FlashLoan/FlashLoan.go
abigen --bin=./build/IERC20.bin --abi=./build/IERC20.abi --pkg=erc20 --type=ERC20 --out=abigen/ERC20/ERC20.go
