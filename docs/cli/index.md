---
order: 1
description: Tinyport CLI docs. 
parent:
  order: 8
  title: CLI Reference
---

# CLI Reference

Documentation for Tinyport CLI.

## tinyport

Tinyport offers everything you need to scaffold, test, build, and launch your blockchain

**Synopsis**

Tinyport is a tool for creating sovereign blockchains built with Cosmos SDK, the world’s
most popular modular blockchain framework. Tinyport offers everything you need to scaffold,
test, build, and launch your blockchain.

To get started, create a blockchain:

tinyport scaffold chain github.com/cosmonaut/mars

**Options**

```
  -h, --help   help for tinyport
```

**SEE ALSO**

* [tinyport account](#tinyport-account)	 - Commands for managing accounts
* [tinyport chain](#tinyport-chain)	 - Build, initialize and start a blockchain node or perform other actions on the blockchain
* [tinyport docs](#tinyport-docs)	 - Show Tinyport docs
* [tinyport generate](#tinyport-generate)	 - Generate clients, API docs from source code
* [tinyport relayer](#tinyport-relayer)	 - Connect blockchains by using IBC protocol
* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more
* [tinyport tools](#tinyport-tools)	 - Tools for advanced users
* [tinyport version](#tinyport-version)	 - Print the current build information


## tinyport account

Commands for managing accounts

**Synopsis**

Commands for managing accounts. An account is a pair of a private key and a public key.
Tinyport uses accounts to interact with the Tinyport Network blockchain, use an IBC relayer, and more.

**Options**

```
  -h, --help   help for account
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain
* [tinyport account create](#tinyport-account-create)	 - Create a new account
* [tinyport account delete](#tinyport-account-delete)	 - Delete an account by name
* [tinyport account export](#tinyport-account-export)	 - Export an account as a private key
* [tinyport account import](#tinyport-account-import)	 - Import an account by using a mnemonic or a private key
* [tinyport account list](#tinyport-account-list)	 - Show a list of all accounts
* [tinyport account show](#tinyport-account-show)	 - Show detailed information about a particular account


## tinyport account create

Create a new account

```
tinyport account create [name] [flags]
```

**Options**

```
  -h, --help                     help for create
      --keyring-backend string   Keyring backend to store your account keys (default "test")
```

**SEE ALSO**

* [tinyport account](#tinyport-account)	 - Commands for managing accounts


## tinyport account delete

Delete an account by name

```
tinyport account delete [name] [flags]
```

**Options**

```
  -h, --help                     help for delete
      --keyring-backend string   Keyring backend to store your account keys (default "test")
```

**SEE ALSO**

* [tinyport account](#tinyport-account)	 - Commands for managing accounts


## tinyport account export

Export an account as a private key

```
tinyport account export [name] [flags]
```

**Options**

```
  -h, --help                     help for export
      --keyring-backend string   Keyring backend to store your account keys (default "test")
      --non-interactive          Do not enter into interactive mode
      --passphrase string        Account passphrase
      --path string              path to export private key. default: ./key_[name]
```

**SEE ALSO**

* [tinyport account](#tinyport-account)	 - Commands for managing accounts


## tinyport account import

Import an account by using a mnemonic or a private key

```
tinyport account import [name] [flags]
```

**Options**

```
  -h, --help                     help for import
      --keyring-backend string   Keyring backend to store your account keys (default "test")
      --non-interactive          Do not enter into interactive mode
      --passphrase string        Account passphrase
      --secret string            Your mnemonic or path to your private key (use interactive mode instead to securely pass your mnemonic)
```

**SEE ALSO**

* [tinyport account](#tinyport-account)	 - Commands for managing accounts


## tinyport account list

Show a list of all accounts

```
tinyport account list [flags]
```

**Options**

```
      --address-prefix string    Account address prefix (default "cosmos")
  -h, --help                     help for list
      --keyring-backend string   Keyring backend to store your account keys (default "test")
```

**SEE ALSO**

* [tinyport account](#tinyport-account)	 - Commands for managing accounts


## tinyport account show

Show detailed information about a particular account

```
tinyport account show [name] [flags]
```

**Options**

```
      --address-prefix string    Account address prefix (default "cosmos")
  -h, --help                     help for show
      --keyring-backend string   Keyring backend to store your account keys (default "test")
```

**SEE ALSO**

* [tinyport account](#tinyport-account)	 - Commands for managing accounts


## tinyport chain

Build, initialize and start a blockchain node or perform other actions on the blockchain

**Synopsis**

Build, initialize and start a blockchain node or perform other actions on the blockchain.

**Options**

```
  -h, --help          help for chain
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain
* [tinyport chain build](#tinyport-chain-build)	 - Build a node binary
* [tinyport chain faucet](#tinyport-chain-faucet)	 - Send coins to an account
* [tinyport chain init](#tinyport-chain-init)	 - Initialize your chain
* [tinyport chain serve](#tinyport-chain-serve)	 - Start a blockchain node in development


## tinyport chain build

Build a node binary

**Synopsis**

By default, build your node binaries
and add the binaries to your $(go env GOPATH)/bin path.

To build binaries for a release, use the --release flag. The app binaries
for one or more specified release targets are built in a release/ dir under the app's
source. Specify the release targets with GOOS:GOARCH build tags.
If the optional --release.targets is not specified, a binary is created for your current environment.

Sample usages:
	- tinyport chain build
	- tinyport chain build --release -t linux:amd64 -t darwin:amd64 -t darwin:arm64

```
tinyport chain build [flags]
```

**Options**

```
  -h, --help                      help for build
      --home string               Home directory used for blockchains
  -o, --output string             binary output path
      --proto-all-modules         Enables proto code generation for 3rd party modules used in your chain. Available only without the --release flag
      --release                   build for a release
      --release.prefix string     tarball prefix for each release target. Available only with --release flag
  -t, --release.targets strings   release targets. Available only with --release flag
  -v, --verbose                   Verbose output
```

**Options inherited from parent commands**

```
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport chain](#tinyport-chain)	 - Build, initialize and start a blockchain node or perform other actions on the blockchain


## tinyport chain faucet

Send coins to an account

```
tinyport chain faucet [address] [coin<,...>] [flags]
```

**Options**

```
  -h, --help          help for faucet
      --home string   Home directory used for blockchains
  -v, --verbose       Verbose output
```

**Options inherited from parent commands**

```
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport chain](#tinyport-chain)	 - Build, initialize and start a blockchain node or perform other actions on the blockchain


## tinyport chain init

Initialize your chain

```
tinyport chain init [flags]
```

**Options**

```
  -h, --help          help for init
      --home string   Home directory used for blockchains
```

**Options inherited from parent commands**

```
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport chain](#tinyport-chain)	 - Build, initialize and start a blockchain node or perform other actions on the blockchain


## tinyport chain serve

Start a blockchain node in development

**Synopsis**

Start a blockchain node with automatic reloading

```
tinyport chain serve [flags]
```

**Options**

```
  -c, --config string       Tinyport config file (default: ./config.yml)
  -f, --force-reset         Force reset of the app state on start and every source change
  -h, --help                help for serve
      --home string         Home directory used for blockchains
      --proto-all-modules   Enables proto code generation for 3rd party modules used in your chain
  -r, --reset-once          Reset of the app state on first start
  -v, --verbose             Verbose output
```

**Options inherited from parent commands**

```
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport chain](#tinyport-chain)	 - Build, initialize and start a blockchain node or perform other actions on the blockchain


## tinyport docs

Show Tinyport docs

```
tinyport docs [flags]
```

**Options**

```
  -h, --help   help for docs
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain


## tinyport generate

Generate clients, API docs from source code

**Synopsis**

Generate clients, API docs from source code.

Such as compiling protocol buffer files into Go or implement particular functionality, for example, generating an OpenAPI spec.

Produced source code can be regenerated by running a command again and is not meant to be edited by hand.

**Options**

```
  -h, --help          help for generate
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain
* [tinyport generate openapi](#tinyport-generate-openapi)	 - Generate generates an OpenAPI spec for your chain from your config.yml
* [tinyport generate proto-go](#tinyport-generate-proto-go)	 - Generate proto based Go code needed for the app's source code
* [tinyport generate vuex](#tinyport-generate-vuex)	 - Generate Vuex store for you chain's frontend from your config.yml


## tinyport generate openapi

Generate generates an OpenAPI spec for your chain from your config.yml

```
tinyport generate openapi [flags]
```

**Options**

```
  -h, --help   help for openapi
```

**Options inherited from parent commands**

```
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport generate](#tinyport-generate)	 - Generate clients, API docs from source code


## tinyport generate proto-go

Generate proto based Go code needed for the app's source code

```
tinyport generate proto-go [flags]
```

**Options**

```
  -h, --help   help for proto-go
```

**Options inherited from parent commands**

```
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport generate](#tinyport-generate)	 - Generate clients, API docs from source code


## tinyport generate vuex

Generate Vuex store for you chain's frontend from your config.yml

```
tinyport generate vuex [flags]
```

**Options**

```
  -h, --help                help for vuex
      --proto-all-modules   Enables proto code generation for 3rd party modules used in your chain
```

**Options inherited from parent commands**

```
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [tinyport generate](#tinyport-generate)	 - Generate clients, API docs from source code


## tinyport relayer

Connect blockchains by using IBC protocol

**Options**

```
  -h, --help   help for relayer
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain
* [tinyport relayer configure](#tinyport-relayer-configure)	 - Configure source and target chains for relaying
* [tinyport relayer connect](#tinyport-relayer-connect)	 - Link chains associated with paths and start relaying tx packets in between


## tinyport relayer configure

Configure source and target chains for relaying

```
tinyport relayer configure [flags]
```

**Options**

```
  -a, --advanced                 Advanced configuration options for custom IBC modules
  -h, --help                     help for configure
      --keyring-backend string   Keyring backend to store your account keys (default "test")
      --ordered                  Set the channel as ordered
      --source-account string    Source Account
      --source-faucet string     Faucet address of the source chain
      --source-gaslimit int      Gas limit used for transactions on source chain
      --source-gasprice string   Gas price used for transactions on source chain
      --source-port string       IBC port ID on the source chain
      --source-prefix string     Address prefix of the source chain
      --source-rpc string        RPC address of the source chain
      --source-version string    Module version on the source chain
      --target-account string    Target Account
      --target-faucet string     Faucet address of the target chain
      --target-gaslimit int      Gas limit used for transactions on target chain
      --target-gasprice string   Gas price used for transactions on target chain
      --target-port string       IBC port ID on the target chain
      --target-prefix string     Address prefix of the target chain
      --target-rpc string        RPC address of the target chain
      --target-version string    Module version on the target chain
```

**SEE ALSO**

* [tinyport relayer](#tinyport-relayer)	 - Connect blockchains by using IBC protocol


## tinyport relayer connect

Link chains associated with paths and start relaying tx packets in between

```
tinyport relayer connect [<path>,...] [flags]
```

**Options**

```
  -h, --help                     help for connect
      --keyring-backend string   Keyring backend to store your account keys (default "test")
```

**SEE ALSO**

* [tinyport relayer](#tinyport-relayer)	 - Connect blockchains by using IBC protocol


## tinyport scaffold

Scaffold a new blockchain, module, message, query, and more

**Synopsis**

Scaffold commands create and modify the source code files to add functionality.

CRUD stands for "create, read, update, delete".

**Options**

```
  -h, --help   help for scaffold
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain
* [tinyport scaffold band](#tinyport-scaffold-band)	 - Scaffold an IBC BandChain query oracle to request real-time data
* [tinyport scaffold chain](#tinyport-scaffold-chain)	 - Fully-featured Cosmos SDK blockchain
* [tinyport scaffold list](#tinyport-scaffold-list)	 - CRUD for data stored as an array
* [tinyport scaffold map](#tinyport-scaffold-map)	 - CRUD for data stored as key-value pairs
* [tinyport scaffold message](#tinyport-scaffold-message)	 - Message to perform state transition on the blockchain
* [tinyport scaffold module](#tinyport-scaffold-module)	 - Scaffold a Cosmos SDK module
* [tinyport scaffold packet](#tinyport-scaffold-packet)	 - Message for sending an IBC packet
* [tinyport scaffold query](#tinyport-scaffold-query)	 - Query to get data from the blockchain
* [tinyport scaffold single](#tinyport-scaffold-single)	 - CRUD for data stored in a single location
* [tinyport scaffold type](#tinyport-scaffold-type)	 - Scaffold only a type definition
* [tinyport scaffold vue](#tinyport-scaffold-vue)	 - Vue 3 web app template


## tinyport scaffold band

Scaffold an IBC BandChain query oracle to request real-time data

**Synopsis**

Scaffold an IBC BandChain query oracle to request real-time data from BandChain scripts in a specific IBC-enabled Cosmos SDK module

```
tinyport scaffold band [queryName] --module [moduleName] [flags]
```

**Options**

```
  -h, --help            help for band
      --module string   IBC Module to add the packet into
  -p, --path string     path of the app (default ".")
      --signer string   Label for the message signer (default: creator)
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold chain

Fully-featured Cosmos SDK blockchain

**Synopsis**

Scaffold a new Cosmos SDK blockchain with a default directory structure

```
tinyport scaffold chain [github.com/org/repo] [flags]
```

**Options**

```
      --address-prefix string   Address prefix (default "cosmos")
  -h, --help                    help for chain
      --no-module               Prevent scaffolding a default module in the app
  -p, --path string             path to scaffold the chain (default ".")
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold list

CRUD for data stored as an array

```
tinyport scaffold list NAME [field]... [flags]
```

**Options**

```
  -h, --help            help for list
      --module string   Module to add into. Default is app's main module
      --no-message      Disable CRUD interaction messages scaffolding
  -p, --path string     path of the app (default ".")
      --signer string   Label for the message signer (default: creator)
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold map

CRUD for data stored as key-value pairs

```
tinyport scaffold map NAME [field]... [flags]
```

**Options**

```
  -h, --help            help for map
      --index strings   fields that index the value (default [index])
      --module string   Module to add into. Default is app's main module
      --no-message      Disable CRUD interaction messages scaffolding
  -p, --path string     path of the app (default ".")
      --signer string   Label for the message signer (default: creator)
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold message

Message to perform state transition on the blockchain

```
tinyport scaffold message [name] [field1] [field2] ... [flags]
```

**Options**

```
  -d, --desc string        Description of the command
  -h, --help               help for message
      --module string      Module to add the message into. Default: app's main module
  -p, --path string        path of the app (default ".")
  -r, --response strings   Response fields
      --signer string      Label for the message signer (default: creator)
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold module

Scaffold a Cosmos SDK module

**Synopsis**

Scaffold a new Cosmos SDK module in the `x` directory

```
tinyport scaffold module [name] [flags]
```

**Options**

```
      --dep strings            module dependencies (e.g. --dep account,bank)
  -h, --help                   help for module
      --ibc                    scaffold an IBC module
      --ordering string        channel ordering of the IBC module [none|ordered|unordered] (default "none")
  -p, --path string            path of the app (default ".")
      --require-registration   if true command will fail if module can't be registered
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold packet

Message for sending an IBC packet

**Synopsis**

Scaffold an IBC packet in a specific IBC-enabled Cosmos SDK module

```
tinyport scaffold packet [packetName] [field1] [field2] ... --module [moduleName] [flags]
```

**Options**

```
      --ack strings     Custom acknowledgment type (field1,field2,...)
  -h, --help            help for packet
      --module string   IBC Module to add the packet into
      --no-message      Disable send message scaffolding
  -p, --path string     path of the app (default ".")
      --signer string   Label for the message signer (default: creator)
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold query

Query to get data from the blockchain

```
tinyport scaffold query [name] [request_field1] [request_field2] ... [flags]
```

**Options**

```
  -d, --desc string        Description of the command
  -h, --help               help for query
      --module string      Module to add the query into. Default: app's main module
      --paginated          Define if the request can be paginated
  -p, --path string        path of the app (default ".")
  -r, --response strings   Response fields
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold single

CRUD for data stored in a single location

```
tinyport scaffold single NAME [field]... [flags]
```

**Options**

```
  -h, --help            help for single
      --module string   Module to add into. Default is app's main module
      --no-message      Disable CRUD interaction messages scaffolding
  -p, --path string     path of the app (default ".")
      --signer string   Label for the message signer (default: creator)
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold type

Scaffold only a type definition

```
tinyport scaffold type NAME [field]... [flags]
```

**Options**

```
  -h, --help            help for type
      --module string   Module to add into. Default is app's main module
      --no-message      Disable CRUD interaction messages scaffolding
  -p, --path string     path of the app (default ".")
      --signer string   Label for the message signer (default: creator)
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport scaffold vue

Vue 3 web app template

```
tinyport scaffold vue [flags]
```

**Options**

```
  -h, --help          help for vue
  -p, --path string   path to scaffold content of the Vue.js app (default "./vue")
```

**SEE ALSO**

* [tinyport scaffold](#tinyport-scaffold)	 - Scaffold a new blockchain, module, message, query, and more


## tinyport tools

Tools for advanced users

**Options**

```
  -h, --help   help for tools
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain
* [tinyport tools completions](#tinyport-tools-completions)	 - Generate completions script
* [tinyport tools ibc-relayer](#tinyport-tools-ibc-relayer)	 - Typescript implementation of an IBC relayer
* [tinyport tools ibc-setup](#tinyport-tools-ibc-setup)	 - Collection of commands to quickly setup a relayer
* [tinyport tools protoc](#tinyport-tools-protoc)	 - Execute the protoc command


## tinyport tools completions

Generate completions script

**Synopsis**

 The completions command outputs a completion script you can use in your shell. The output script requires 
				that [bash-completion](https://github.com/scop/bash-completion)	is installed and enabled in your 
				system. Since most Unix-like operating systems come with bash-completion by default, bash-completion 
				is probably already installed and operational.

Bash:

  $ source <(tinyport  tools completions bash)

  To load completions for every new session, run:

  ** Linux **
  $ tinyport  tools completions bash > /etc/bash_completion.d/tinyport

  ** macOS **
  $ tinyport  tools completions bash > /usr/local/etc/bash_completion.d/tinyport

Zsh:

  If shell completions is not already enabled in your environment, you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  To load completions for each session, execute once:
  
  $ tinyport  tools completions zsh > "${fpath[1]}/_starport"

  You will need to start a new shell for this setup to take effect.

fish:

  $ tinyport  tools completions fish | source

  To load completions for each session, execute once:
  
  $ tinyport  tools completions fish > ~/.config/fish/completionss/tinyport.fish

PowerShell:

  PS> tinyport  tools completions powershell | Out-String | Invoke-Expression

  To load completions for every new session, run:
  
  PS> tinyport  tools completions powershell > tinyport.ps1
  
  and source this file from your PowerShell profile.


```
tinyport tools completions
```

**Options**

```
  -h, --help   help for completions
```

**SEE ALSO**

* [tinyport tools](#tinyport-tools)	 - Tools for advanced users


## tinyport tools ibc-relayer

Typescript implementation of an IBC relayer

```
tinyport tools ibc-relayer [--] [...] [flags]
```

**Examples**

```
tinyport tools ibc-relayer -- -h
```

**Options**

```
  -h, --help   help for ibc-relayer
```

**SEE ALSO**

* [tinyport tools](#tinyport-tools)	 - Tools for advanced users


## tinyport tools ibc-setup

Collection of commands to quickly setup a relayer

```
tinyport tools ibc-setup [--] [...] [flags]
```

**Examples**

```
tinyport tools ibc-setup -- -h
tinyport tools ibc-setup -- init --src relayer_test_1 --dest relayer_test_2
```

**Options**

```
  -h, --help   help for ibc-setup
```

**SEE ALSO**

* [tinyport tools](#tinyport-tools)	 - Tools for advanced users


## tinyport tools protoc

Execute the protoc command

**Synopsis**

The protoc command. You don't need to setup the global protoc include folder with -I, it's automatically handled

```
tinyport tools protoc [--] [...] [flags]
```

**Examples**

```
tinyport tools protoc -- --version
```

**Options**

```
  -h, --help   help for protoc
```

**SEE ALSO**

* [tinyport tools](#tinyport-tools)	 - Tools for advanced users


## tinyport version

Print the current build information

```
tinyport version [flags]
```

**Options**

```
  -h, --help   help for version
```

**SEE ALSO**

* [tinyport](#tinyport)	 - Tinyport offers everything you need to scaffold, test, build, and launch your blockchain

