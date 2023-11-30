# VerifyingL1RollupData_Part-1_Demo-nstallation

## command for CTC
```shell
solcjs --abi CanonicalTransactionChain.sol
```
CanonicalTransactionChain.abi created
```shell
solcjs --bin CanonicalTransactionChain.sol
```
CanonicalTransactionChain.bin created
```shell
abigen --bin=ICanonicalTransactionChain_sol_ICanonicalTransactionChain.bin --abi=ICanonicalTransactionChain_sol_ICanonicalTransactionChain.abi --pkg=contracts --out=CanonicalTransactionchain.go
```
CanonicalTransactionchain.go created



## command for SCC
```shell
solcjs --abi StateCommitmentChain.sol
```
StateCommitmentChain.abi created

```shell
solcjs --bin StateCommitmentChain.sol
```
StateCommitmentChain.bin created

```shell
abigen --bin=IStateCommitmentChain_sol_IStateCommitmentChain.bin --abi=IStateCommitmentChain_sol_IStateCommitmentChain.abi --pkg=contracts --out=StateCommitmentChain.go
```
StateCommitmentChain.go created



## command for running main package
```shell
go run event_read.go
```



## error
- Source "@openzeppelin/contracts/access/Ownable.sol" not found: File not found inside the base path or any of the include paths.
```shell
yarn add @openzeppelin/contracts
```

- can’t find path of specifit contract
```shell
solcjs --bin --abi --include-path node_modules/ --base-path . -o ethereum/contracts/build ethereum/contracts/TokenSwap.sol”```
```

