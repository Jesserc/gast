# Gast - Ethereum  Transaction Toolkit

Gast is an open-source command-line transaction toolkit designed to streamline the management of Ethereum (including testnets and EVM compatible Layer 2s) transactions and gas prices. It provides some sets of commands for managing transactions (including creation, signing, and tracing).

## Installation

Firstly, ensure you have Go language installed. To verify, run:
```shell
 go version
```
If not installed, you can here: https://go.dev/dl/ 

After installing Go, run this command to install Gast:
```shell
go install github.com/Jesserc/gast@latest
```
Check if Gast is installed:
```shell
gast help
```
If the command is not recognized, you would have to add the Go Bin directory to your path.
```shell
export PATH=$PATH:$(go env GOPATH)/bin
```
Or manually paste it in your `.bashrc`, `.zshrc` file.
```shell
nano ~/.zshrc # add => export PATH=$PATH:$(go env GOPATH)/bin
# or
nano ~/.bashrc # export => PATH=$PATH:$(go env GOPATH)/bin
```

## Usage

To use Gast, run:

```shell
gast [command]
```

### Available Commands

- `completion`: Generate the autocompletion script for the specified shell.
- `gas-price`: Fetch the current gas price from specified Ethereum networks.
- `help`: Help about any command.
- `tx`: Manages Ethereum transactions, including creation, signing, and tracing.

### Flags

- `-h, --help`: Help for Gast.

For more information about a command, use:

```shell
gast [command] --help
```

### Transaction Management

Manage Ethereum transactions with ease. The `tx` command supports a variety of subcommands:
```shell
gast tx [sub-command] [flags]
```
* `create-contract`: Deploy Solidity contract
* `create-raw`: Generate a raw, signed EIP-1559 transaction
* `send-raw`: Submit a raw, signed transaction
* `send-blob`: Create and send an EIP-4844 blob transaction
* `send`: Send EIP-1559 transaction
* `trace`: Retrieve and display the execution trace (path) of a given transaction hash
* `sign-message`: Sign a given message with a private key
* `verify-sig`: Verify the signature of a signed message (can be created with the sign-message command)
* `estimate-gas`: Estimate the gas required to execute a given transaction
* `get-nonce`: Get the transaction count of an account

[//]: # (- `create-contract`: Deploy Solidity contract &#40;solc must be installed&#41;.)

[//]: # (- `create-raw`: Generates a raw, unsigned EIP-1559 transaction.)

[//]: # (- `send-raw`: Submits a raw, signed transaction to the Ethereum network.)

[//]: # (- `send`: Submits a constructed transaction.)

[//]: # (- `send-blob`: Submits a constructed EIP-4844 blob transaction.)

[//]: # (- `estimate-gas`: Provides an estimate of the gas required to execute a given transaction.)

[//]: # (- `get-nonce`: Get transaction count of an account.)

[//]: # (- `sign-message`: Signs a given message with the private key.)

[//]: # (- `trace`: Retrieves and displays the execution trace &#40;path&#41; of a given transaction hash using `ots_traceTransaction`.)

[//]: # (- `verify-sig`: Verifies the signature of a signed message.)

### Fetching Gas Prices

To fetch the current gas prices from specific Ethereum networks, use:

```shell
gast gas-price [flags]
```

Supported flags include:

- `--eth`: Use the default Ethereum RPC URL.
- `--op`: Use the default Optimism RPC URL.
- `--arb`: Use the default Arbitrum RPC URL.
- `--base`: Use the default Base RPC URL.
- `--linea`: Use the default Linea RPC URL.
- `--zksync`: Use the default zkSync RPC URL.
- `-u, --rpc-url string`: Specify a custom RPC URL for fetching the gas price.

## Generating Completions
To generate completions for zsh shell, run:
```shell
gast completion zsh # or use your shell
```
Save completion script to a completion file:
```shell
mkdir -p ~/.zsh/completion
nano ~/.zsh/completion/_gast # paste the completion script here
```
Open .zshrc to Include the Completion Directory:
```shell
nano ~/.zshrc
```
Add the following lines to set completions:
```shell
fpath=(~/.zsh/completion $fpath)
autoload -Uz compinit
compinit
```
Apply changes:
```shell
source ~/.zshrc
```

## Command Examples

Here are some practical examples to help you get started with Gast. These commands demonstrate how to use Gast for managing Ethereum transactions and gas prices.

### Fetching Gas Prices

```shell
# Fetch current gas price for Ethereum network
gast gas-price --eth 

# Fetch gas price from a custom RPC URL
gast gas-price --rpc-url https://forno.celo.org
```

### Contract Deployment

```shell
# Deploy a Solidity contract
gast tx create-contract --rpc-url https://sepolia.drpc.org --private-key "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662" --gas-limit 1599000 -d contracts/CurrentYear.sol
```

### Creating and Sending Transactions

```shell
# Create a raw, signed EIP-1559 transaction
gast tx create-raw --rpc-url "https://eth-sepolia.g.alchemy.com/v2/Of6ow3pvkFafGPn8Y2uk9vz4bSveZQxa" --to "0x4924Fb92285Cb10BC440E6fb4A53c2B94f2930c5" --private-key "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662" --gas-limit 21000 --wei 10000000000000

# Submit a raw, signed transaction to the Ethereum network
gast tx send-raw --raw-tx b87402f87183aa36a781d7843b9aca0084f702d4c28256ce944924fb92285cb10bc440e6fb4a53c2b94f2930c58398968080c080a081725247a454fc36e3ecd411ef6e7ddb89e668745fb2a5169ea08bfc4f5b617ba013cce55e74f620f15904e30a1c0f3e5dad22919e782468afe372d3bc6f5222b0 --rpc-url "https://eth-sepolia.g.alchemy.com/v2/Of6ow3pvkFafGPn8Y2uk9vz4bSveZQxa"

# Estimate gas for a transaction
gast tx estimate-gas --from 0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045 --to 0xbe0eb53f46cd790cd13851d5eff43d12404d33e8 --rpc-url https://rpc.mevblocker.io --data 'Hello Ethereum!'

# Get the transaction count of an account
gast tx get-nonce --address 0x8741Fb04b7d8f5A01e0ec1D454602Bc08BDB0c8c --rpc-url https://sepolia.drpc.org
```

### EIP-4844 Blob Transactions

```shell
# Create and send an EIP-4844 blob transaction
gast tx send-blob --to 0x571B102323C3b8B8Afb30619Ac1d36d85359fb84 --rpc-url "https://rpc2.sepolia.org" --private-key "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662" --blob-data 'Hello Blobs!' --dir gast/blob-tx # dir to save blob tx result
```

### Message Signing and Verification

```shell
# Sign a message with a private key
gast tx sign-message -m Jesserc -p "2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662"

# Verify the signature of a signed message
gast tx verify-sig --sig 0x5e9faa36429804f79bd8ca495e21095f29f1038ec2b3f10788437a16d52f79682aca534e2b4ff0f426d6444555d807e6bc1c7c8a6b21aaaa4676d4f5e8d45b541b --address 0x571B102323C3b8B8Afb30619Ac1d36d85359fb84 --msg Jesserc
```

### Transaction Tracing

```shell
# Retrieve the execution trace of a transaction
gast tx trace --hash 0xee92800f24e23971c0ab031b30d60d6414e2255a308993d902604f4cfc1e4e7f -u https://rpc.builder0x69.io/
```


## Tests
To run unit test:
```shell
go test -v ./... -short
```

To run integration test (connects to network):
```shell
 go test -v ./...  -run Integration
```

## Contribution
Good PRs or suggestions are welcomed.