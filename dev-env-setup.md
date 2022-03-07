# Set up your Tinyport development environment 

To ensure you have a successful experience developing with Tinyport, make sure that your local system meets these technical requirements.

Tinyport is supported for the following operating systems:

- GNU/Linux
- macOS
- Windows Subsystem for Linux (WSL)

## Install Go

This installation method removes existing Go installations, installs Go, and sets the environment variables.

1. Install the latest version of [Go](https://golang.org/doc/install).

2. Download the release that is suitable for your system.

3. Follow the installation instructions.

**Note:** We recommend not using brew to install Go.

## Add the Go bin directory to your PATH 

Ensure the Go environment variables are [set properly](https://golang.org/doc/gopath_code#GOPATH) on your system. Many of the initial problems are related to incorrect environment variables. 

1. Edit your `~/.bashrc` file and add `export PATH=$PATH:$(go env GOPATH)/bin`. 
2. To apply the changes, run `source ~/.bashrc`.

## Remove existing Tinyport installations 

Before you install a new version of Tinyport, remove all existing Tinyport installations. 

1. Remove the Tinyport binary with `rm $(which tinyport)`

   Depending on your user permissions, run the command with or without `sudo`.

2. Repeat this step until all `tinyport` installations are removed from your system.

`curl https://get.tinyport.network/tinyport! | bash`

See [Install Tinyport](docs/install.md).

## Clone the Tinyport repo

`git clone --depth=1 git@github.com:notional-labs/tinyport.git`

## Run make install 

1. After you clone the `tinyport` repo, change into the root directory `cd tinyport`.

2. Run `make install`.

## Verify Your Tinyport version 

To verify the version of Tinyport you have installed, run `tinyport version`. 

The latest version is `Tinyport version: development`. 
