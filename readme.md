# Tinyport


Tinyport is a scaffolding tool for cosmos-sdk blockchains. Like [tinyseed](https://github.com/notional-labs/tinyseed) tinyport is a fork that improves usability through extreme minimalism. 

* [**Build a blockchain with Tinyport in a web-based IDE** (stable)](https://gitpod.io/#https://github.com/notional-labs/tinyport/tree/master) or use [nightly version](https://gitpod.io/#https://github.com/notional-labs/tinyport/)
* [Check out the latest features in v0.19](https://tinyport.com/blog/tinyport-v0-19)

## Quick start

Open Tinyport [in your browser](https://gitpod.io/#https://github.com/notional-labs/tinyport/tree/master), or [install it](https://docs.tinyport.network/guide/install.html). Create and start a blockchain:

```bash
tinyport scaffold chain github.com/cosmonaut/mars

cd mars

tinyport chain serve
```

## Documentation

To learn how to use Tinyport, check out the [Tinyport Documentation](https://docs.tinyport.com). To learn more about how to build blockchain apps with Tinyport, see the [Tinyport Developer Tutorials](https://docs.tinyport.com/guide/). 

To install Tinyport locally on GNU, Linux, or macOS, see [Install Tinyport](https://docs.tinyport.com/guide/install.html).

To learn more about building a JavaScript frontend for your Cosmos SDK blockchain, see [tendermint/vue](https://github.com/tendermint/vue).

## Questions

For questions and support, join the official [Tinyport Discord](https://discord.gg/ignt) server. The issue list in this repo is exclusively for bug reports and feature requests.

## Cosmos SDK Compatibility

Blockchains created with Tinyport use the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) framework. To ensure the best possible experience, use the version of Tinyport that corresponds to the version of Cosmos SDK that your blockchain is built with. Unless noted otherwise, a row refers to a minor version and all associated patch versions.

| Tinyport | Cosmos SDK | IBC                  | Notes                                            |
| -------- | ---------- | -------------------- | ------------------------------------------------ |
| v0.19.2  | v0.44.5    | v2.0.2               | |
| v0.19    | v0.44      | v1.2.2               | |
| v0.18    | v0.44      | v1.2.2               | `tinyport chain serve` works with v0.44.x chains |
| v0.17    | v0.42      | Same with Cosmos SDK | |

To upgrade your blockchain to the newer version of Cosmos SDK, see the [Migration guide](https://docs.tinyport.com/migration/).

## Contributing

We welcome contributions from everyone. The `develop` branch contains the development version of the code. You can create a branch from `develop` and create a pull request, or maintain your own fork and submit a cross-repository pull request. 

**Important** Before you start implementing a new Tinyport feature, the first step is to create an issue on Github that describes the proposed changes.

If you're not sure where to start, check out [contributing.md](contributing.md) for our guidelines and policies for how we develop Tinyport. Thank you to everyone who has contributed to Tinyport!

## Community

Tinyport is a free and open-source product maintained by [Notional](https://notional.ventures). Here's where you can find us. Stay in touch.


