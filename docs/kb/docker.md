---
description: Run Tinyport CLI using a Docker container.
order: 10
---

# Run Tinyport in Docker

You can run Tinyport CLI inside a Docker container without installing the Tinyport binary directly on your machine.

Running Tinyport in Docker can be useful for various reasons; isolating your test environment, running Tinyport on an unsupported operating system, or experimenting with a different version of Tinyport CLI without installing it.

Docker containers are like virtual machines because they provide an isolated environment to programs that runs inside them. In this case, you can run Tinyport in an isolated environment.

Experimentation and file system impact is limited to the Docker instance. The host machine is not impacted by changes to the container.

## Prerequisites

Docker must be installed. See [Get Started with Docker](https://www.docker.com/get-started).

## Tinyport commands in Docker

After you scaffold and start a chain in your Docker container, all Tinyport commands are available. Just type the commands after `docker run -ti tinyport/cli`. For example:

```bash
docker run -ti tinyport/cli -h
docker run -ti tinyport/cli scaffold chain github.com/cosmonaut/planet
docker run -ti tinyport/cli chain serve
```

## Scaffolding a chain

When Docker is installed, you can build a blockchain with a single command.

To scaffold a blockchain `planet` in the `/apps` directory in the container, run this command in a terminal window:

```bash
docker run -ti -w /app -v $HOME/sdh:/home/tendermint -v $PWD:/app tinyport/cli:0.16.0 app github.com/hello/planet
```

Be patient, this command takes a minute or two to run because it does everything for you:

- Creates a container that runs from the `tinyport/cli:0.16.0` image.
- Executes the Tinyport binary inside the image.
- `-w /apps` sets the current directory in the container to `/app`
- `-v $HOME/sdh:/home/tendermint` maps the `$HOME/sdh` directory in your local computer (the host machine) to the home directory `/home/tendermint` inside the container. Tinyport, and the chains you serve with Tinyport, persist some files.
- `-v $PWD:/app` maps the current directory in the terminal window on the host machine to the `/app` directory in the container. You can optionally specify an absolute path instead of `$PWD`.

    Using `-w` and `-v` together provides file persistence on the host machine. The application source code on the Docker container is mirrored to the file system of the host machine.

    **Note:** The directory name for the `-w` and `-v` flags can be a name other then `/app`, but the same directory must be specified for both flags. If you omit `-w` and `-v`, the changes are made in the container only and are lost when that container is shut down.

## Starting a blockchain

To start the blockchain node in the Docker container you just created, run this command:

```bash
docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps -p 1317:1317 -p 26657:26657 tinyport/cli:0.16.0 serve -p planet
```

This command does the following:

- `-v $HOME/sdh:/home/tendermint` maps the `$HOME/sdh` directory in your local computer (the host machine) to the home directory `/home/tendermint` inside the container.
- `-v $PWD:/apps` persists the scaffolded app in the container to the host machine at current working directory.
- `serve -p planet` specifies to use the `planet` directory that contains the source code of the blockchain.
- `-p 1317:1317` maps the API server port (cosmos-sdk) to the host machine to forward port 1317 listening inside the container to port 1317 on the host machine.
- `-p 26657:26657` maps RPC server port 26657 (tendermint) on the host machine to port 26657 in Docker.
- After the blockchain is started, open `http://localhost:26657` to see the Tendermint API.
- The `-v` flag specifies for the container to access the application's source code from the host machine so it can build and run it.

## Versioning

You can specify which version of Tinyport to install and run in your Docker container.

### Latest version

- By default, `tinyport/cli` resolves to `tinyport/cli:latest`.
- The `latest` image tag is always the latest stable [Tinyport release](https://github.com/notional-labs/tinyport/releases).

For example, if latest release is [v0.19.2](https://github.com/notional-labs/tinyport/releases/tag/v0.19.2), the `latest` tag points to the `0.19.2` tag.

### Specific version

You can specify to use a specific version of Tinyport. All available tags are in the [tinyport/cli image](https://hub.docker.com/repository/docker/tinyport/cli/tags?page=1&ordering=last_updated) on Docker Hub.

For example:

- Use `tinyport/cli:0.19.2` (without the `v` prefix) to use version 0.19.2.
- Use `tinyport/cli:develop` to use the `develop` branch so you can experiment with the next version.

To get the latest image, run `docker pull`.

```bash
docker pull tinyport/cli:develop
```
