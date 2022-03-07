# Welcome to Tinyport ✨

Tinyport is an easy-to-use CLI tool for creating sovereign blockchains. Blockchains created with Tinyport use Cosmos SDK, the world's most widely used blockchain application framework.

In this browser-based development environment, the terminal window is in the lower part of the window. The `tinyport` binary is pre-installed and ready to use on the command line.

## Quick start

To create a blockchain and start a node in development:

```
tinyport scaffold chain github.com/cosmonaut/mars

cd mars

tinyport chain serve
```

## Next steps

📺 **[Introduction to Tinyport](https://www.youtube.com/watch?v=5RqAIE0b8Kw)**: Watch an introductory video to learn about Tinyport.

🧑‍🏫 **[Developer Guide](https://docs.tinyport.network/guide/)**: Learn by building a simple IBC-enabled module, nameservice, or a decentralized exchange (DEX).

📕 **[Tinyport Documentation](https://docs.tinyport.network)**: Explore the features of Tinyport.

📚 [Cosmos SDK Documentation](https://docs.cosmos.network): Learn about the framework for building application-specific blockchains.

⭐️ [Tinyport on Github](https://github.com/notional-labs/tinyport): Submit an issue or contribute to the source code.

## Tinyport features

* Scaffold modules, messages, types with CRUD operations, IBC packets, and more
* Start a blockchain node in development with live reloading
* Connect to other blockchains with a built-in IBC relayer
* Use automatically generated TypeScript/Vuex clients to interact with your blockchain
* Use the Vue.js web app template with a set of components and Vuex modules

## Install Tinyport locally

```
curl https://get.tinyport.network/tinyport! | bash
```

The latest `tinyport` binary is downloaded from the Github repo and installed in `/usr/local/bin`. Learn more about [installing Tinyport](https://docs.tinyport.network/guide/install.html).

## Stay in touch

Tinyport is a free and open source product maintained by [Tendermint](https://tendermint.com). Follow us on [Twitter](https://twitter.com/starportHQ) and [Medium](https://medium.com/tendermint) to get the latest updates!
