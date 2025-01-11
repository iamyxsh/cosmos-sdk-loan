# Cosmos-Loan

Welcome to **Cosmos-Loan**, a blockchain application built using the Cosmos SDK and Ignite CLI. This repository contains the source code, modules, and configurations for our custom blockchain designed to provide seamless decentralized loan services.

## Features

- **Custom Modules**: Includes purpose-built modules to handle decentralized lending and borrowing.
- **Cosmos SDK Integration**: Leverages the Cosmos SDK for modular and secure blockchain infrastructure.
- **Interoperability**: Fully supports the Inter-Blockchain Communication (IBC) protocol for cross-chain interaction.
- **Customizable Parameters**: Easily adjust chain parameters such as gas limits, fees, and staking mechanisms.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Running the Blockchain Locally](#running-the-blockchain-locally)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

To build and run this project, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (v1.19 or later)
- [Ignite CLI](https://ignite.com/) (v1.8.0 or later)
- [Node.js](https://nodejs.org/) (for frontend interactions, if applicable)

## Getting Started

Clone the repository:

```bash
git clone git@github.com:iamyxsh/cosmos-sdk-loan.git
cd cosmos-sdk-loan
```

Initialize and configure the blockchain:

```bash
ignite chain serve
```

This command will:

- Build the project.
- Start a local blockchain node.
- Open REST and gRPC endpoints for interaction.

## Running the Blockchain Locally

To start a local node for development:

1. Build and run the chain:

   ```bash
   ignite chain serve
   ```

2. Access endpoints for testing:

   - REST: `http://localhost:1317`
   - gRPC: `localhost:9090`
   - gRPC-Web: `http://localhost:9091`

3. Interact with the blockchain via the command-line interface (CLI):

   ```bash
   ./cosmos-loand query [module] [command]
   ```

## Project Structure

```plaintext
.
├── app/                 # Cosmos SDK application setup
├── cmd/                 # Main CLI commands
├── x/
│   ├── [module-name]/   # Custom modules
│   ├── ...
├── proto/               # Protocol Buffers definitions
├── config/              # Chain configurations
├── docs/                # Documentation
└── ...
```

## Usage

### Querying the Blockchain

To query the blockchain, use the CLI:

```bash
./cosmos-loand query [module] [query-command]
```

For example, to check the chain status:

```bash
./cosmos-loand status
```

### Submitting Transactions

To submit transactions, use:

```bash
./cosmos-loand tx [module] [tx-command] --from [key-name]
```

For example, to transfer tokens:

```bash
./cosmos-loand tx bank send [from-address] [to-address] [amount] --from [key-name]
```

### Interacting via REST or gRPC

- REST: Use `curl` or Postman to interact with the REST API.
- gRPC: Use a gRPC client for advanced interactions.

## Contributing

We welcome contributions! To get started:

1. Fork the repository.
2. Create a new branch for your feature/bugfix.
3. Submit a pull request with detailed descriptions of your changes.

For detailed guidelines, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute this project as per the license terms.

---

For any issues, please [open an issue](https://github.com/iamyxsh/cosmos-sdk-loan/issues) or contact us at [iamyxsh@gmail.com].
