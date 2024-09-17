This is a fork of https://github.com/Layr-Labs/eigenpod-proofs-generation that wraps
checkpoint related transactions to be use `etherfiNodesManager.forwardEigenpodCall()`
instead of calling the methods directly on the eigenpod. The reason for this is
so that we don't need to set `eigenpod.proofSubmitter` for all 50k eigenpods we control

You should be able to use this fork exactly as you would the normal CLI. The
`--sender` key you use must have `operatingAdmin` permissions on `EtherfiNodesManager` contract

Only supports mainnet at this time

# Introduction

PEPE Changes how we prove balances to EigenLayer. For more information, check out some of the links below.

WARNING: You should only use this software with Holesky Testnet. When PePe goes to production, we will update this tool to support further networks. Thanks!

## Links

- [More about PEPE](https://hackmd.io/U36dE9lnQha3tbf7D0GtKw?view)
- [Contract Documentation](https://github.com/Layr-Labs/eigenlayer-contracts/blob/feat/partial-withdrawal-batching/docs/core/EigenPod.md)

# Usage

- If you want to produce and submit proofs onchain -- either immediately, or by writing to a file to submit later -- check out our [CLI](./cli/README.md). The CLI can produce both credential and checkpoint proofs, and submit them onchain if given a private key.

- If you want to produce proofs from within Golang, please use `cli/core:GenerateValidatorProof` or `cli/core:GenerateCheckpointProof` for our high-level APIs. These will handle downloading beacon state, interfacing with an eth node, and generating the relevant proofs. Lower level APIs are available in `prove_validator.go`.

## Questions

For any questions, feel free to;

- Open a Github Issue
- Ask in [Discord](https://discord.com/invite/eigenlayer)
