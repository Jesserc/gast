# Gast - Ethereum Toolkit

### (IN DEVELOPMENT)

Gast is a command-line toolkit designed to streamline the management of Ethereum transactions and gas prices. It provides some sets of commands for managing transactions (including creation, signing, and tracing).

## Installation

Firstly, ensure you have Go language installed. To verify, run:
```shell
 go version
```
If not installed, you can see installation steps here: 

After installing Go, run this command to install Gast:
```shell
go install github.com/jesserc/gast
```
Check if Gast is installed:
```shell
Gast help
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

- `--config string`: Config file (default is `$HOME/.gast.yaml`).
- `-h, --help`: Help for Gast.
- `-t, --toggle`: Help message for toggle.

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
* `create-raw     `: Generate a raw, signed EIP-1559 transaction
* `send-raw       `: Submit a raw, signed transaction
* `send-blob      `: Create and send an EIP-4844 blob transaction
* `send           `: Send EIP-1559 transaction
* `trace          `: Retrieve and display the execution trace (path) of a given transaction hash
* `sign-message   `: Sign a given message with a private key
* `verify-sig     `: Verify the signature of a signed message (can be created with the sign-message command)
* `estimate-gas   `: Estimate the gas required to execute a given transaction
* `get-nonce      `: Get the transaction count of an account

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
- `-u, --url string`: Specify a custom RPC URL for fetching the gas price.

## Configuration

[//]: # (Gast uses a configuration file located by default at `$HOME/.gast.yaml`. This file allows you to set default values for various flags and commands.)

(TODO)

## Tests
To run unit test:
```shell
 go test -v ./...  -skip Integration
```

To run integration test (connects to network):
```shell
 go test -v ./...  -run Integration
```
