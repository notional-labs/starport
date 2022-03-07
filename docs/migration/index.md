---
order: 1
title: v0.19.2
parent:
  title: Migration
  order: 3
description: For chains that were scaffolded with Tinyport versions lower than v0.19.2, changes are required to use Tinyport v0.19.2. 
---

# Upgrading a blockchain to use Tinyport v0.19.2

Tinyport v0.19.2 comes with IBC v2.0.2.

With Tinyport v0.19.2, the contents of the deprecated Tinyport Modules `tendermint/spm` repo are moved to the official Tinyport repo which introduces breaking changes.

To migrate your chain that was scaffolded with Tinyport versions lower than v0.19.2: 

1. IBC upgrade: Use the [IBC migration documents](https://github.com/cosmos/ibc-go/blob/main/docs/migrations/v1-to-v2.md)
   
2. In your chain's `go.mod` file, remove `tendermint/spm` and add the v0.19.2 version of `notional-labs/tinyport`. If your chain uses these packages, change the import paths as shown: 

    - `github.com/tendermint/spm/ibckeeper` moved to `github.com/notional-labs/tinyport/tinyport/pkg/cosmosibckeeper`
    - `github.com/tendermint/spm/cosmoscmd` moved to `github.com/notional-labs/tinyport/tinyport/pkg/cosmoscmd` 
    - `github.com/tendermint/spm/openapiconsole` moved to `github.com/notional-labs/tinyport/tinyport/pkg/openapiconsole`
    - `github.com/tendermint/spm/testutil/sample` moved to `github.com/notional-labs/tinyport/tinyport/pkg/cosmostestutil/sample`