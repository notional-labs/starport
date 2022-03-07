---
order: 1
description: Steps to install Tinyport on your local computer.
---

# Install Tinyport

You can run [Tinyport](https://github.com/notional-labs/tinyport) in a web-based Gitpod IDE or you can install Tinyport on your local computer.

## Prerequisites

Be sure you have met the prerequisites before you install and use Tinyport.

### Operating systems

Tinyport is supported for the following operating systems:

- GNU/Linux
- macOS
- Windows Subsystem for Linux (WSL)

### Go

Tinyport is written in the Go programming language. To use Tinyport on a local system:

- Install [Go](https://golang.org/doc/install) (**version 1.16** or higher)
- Ensure the Go environment variables are [set properly](https://golang.org/doc/gopath_code#GOPATH) on your system

## Verify your Tinyport version

To verify the version of Tinyport you have installed, run the following command:

```sh
tinyport version
```

## Installing Tinyport

To install the latest version of the `tinyport` binary use the following command.

```bash
curl https://get.tinyport.network/tinyport! | bash
```

This command invokes `curl` to download the install script and pipes the output to `bash` to perform the installation. The `tinyport` binary is installed in `/usr/local/bin`.

To learn more or customize the installation process, see the [installer docs](https://github.com/allinbits/tinyport-installer) on GitHub.

### Write permission

Tinyport installation requires write permission to the `/usr/local/bin/` directory. If the installation fails because you do not have write permission to `/usr/local/bin/`, run the following command:

```bash
curl https://get.tinyport.network/tinyport | bash
```

Then run this command to move the `tinyport` executable to `/usr/local/bin/`:

```bash
sudo mv tinyport /usr/local/bin/
```

On some machines, a permissions error occurs:

```bash
mv: rename ./tinyport to /usr/local/bin/tinyport: Permission denied
============
Error: mv failed
```

In this case, use sudo before `curl` and before `bash`:

```bash
sudo curl https://get.tinyport.network/tinyport! | sudo bash
```

## Upgrading your Tinyport installation

Before you install a new version of Tinyport, remove all existing Tinyport installations.

To remove the current Tinyport installation:

1. On your terminal window, press `Ctrl+C` to stop the chain that you started with `tinyport chain serve`.
1. Remove the Tinyport binary with `rm $(which tinyport)`.
   Depending on your user permissions, run the command with or without `sudo`.
1. Repeat this step until all `tinyport` installations are removed from your system.

After all existing Tinyport installations are removed, follow the  [Installing Tinyport](#installing-tinyport) instructions.

For details on version features and changes, see the [changelog.md](https://github.com/notional-labs/tinyport/blob/develop/changelog.md) in the repo.

## Build from source

To experiment with the source code, you can build from source:

```bash
git clone https://github.com/notional-labs/tinyport --depth=1
cd tinyport && make install
```

## Summary

- Verify the prerequisites.
- To setup a local development environment, install Tinyport locally on your computer.
- Install Tinyport by fetching the binary using cURL or by building from source.
- The latest version is installed by default. You can install previous versions of the precompiled `tinyport` binary.
- Stop the chain and remove existing versions before installing a new version.
