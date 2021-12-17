# BSC-Relay

This project contains the smart contracts and the relayer software for a simple blockchain relay that enables the verification of BSC blocks on Ethereum. 
This is exclusively a research prototype and only the basic functionality is implemented.

## Installation

### Prerequisites

* [Hardhat](https://hardhat.org/)
* [Golang](https://go.dev/)
* [Solidity](https://docs.soliditylang.org/en/v0.8.10/)

### Contract

1. Change into the contract directory: `cd bsc_contracts/`
2. Install all dependencies: `npm install`
3. Change the Ropsten network configuration in `hardhat.config.js`
4. Update the provider for BSC in `./scripts/deploy.js`
5. Run `hardhat run --network ropsten ./scripts/deploy.js` to deploy the example on Ropsten

### Relayer

1. Change into the contract directory: `cd bsc_relayer/`
2. Build the example caller with `go build -o bsc_relayer ./cmd/main.go`
3. Adapt the configuration in `./configs/config.json`
4. Run the example caller with `./bsc_relayer`

## Contributing

This is a research prototype. We welcome anyone to contribute. File a bug report or submit feature requests through the issue tracker. If you want to contribute feel free to submit a pull request.

## Acknowledgement

The financial support by the Austrian Federal Ministry for Digital and Economic Affairs, the National Foundation for Research, Technology and Development as well as the Christian Doppler Research Association is gratefully acknowledged.

## License

This project is licensed under the [MIT License](LICENSE).
