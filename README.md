# Gast - Ethereum Toolkit

## (STILL IN DEVELOPMENT)

Gast is a command-line toolkit designed to streamline the management of Ethereum transactions and gas prices. It provides a suite of commands for fetching current gas prices, managing transactions (including creation, signing, and tracing), and much more directly from your terminal.

## Installation

TODO

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
gast tx --help
```

### Transaction Management

Manage Ethereum transactions with ease. The `tx` command supports a variety of subcommands:
```shell
gast tx [sub-command] [flags]
```

- `create-contract`: Deploy Solidity contract (solc must be installed).
- `create-raw`: Generates a raw, unsigned EIP-1559 transaction.
- `estimate-gas`: Provides an estimate of the gas required to execute a given transaction.
- `get-nonce`: Get transaction count of an account. 
- `send-raw`: Submits a raw, signed transaction to the Ethereum network.
- `sign-message`: Signs a given message with the private key.
- `trace`: Retrieves and displays the execution trace (path) of a given transaction hash using `ots_traceTransaction`.
- `verify-sig`: Verifies the signature of a signed message.

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

Gast uses a configuration file located by default at `$HOME/.gast.yaml`. This file allows you to set default values for various flags and commands.

(TODO - not implemented yet)

## Contributing

TODO

## License
TODO