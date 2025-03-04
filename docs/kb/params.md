---
order: 15
description: Scaffold module parameters to be accessible to the module.
---

# Module parameters

Sometimes you need to set default parameters for a module. The Cosmos SDK [params package](https://docs.cosmos.network/master/modules/params) provides a globally available parameter that is saved into the key-value store. 

Params are managed and centralized by the Cosmos SDK `params` module and are updated with a governance proposal.

You can use Tinyport to scaffold parameters to be accessible for the module. Parameters have default values that can be changed when the chain is live. Since the parameters are managed and centralized by the Cosmos SDK params module, they can be easily updated using a governance proposal.

To scaffold a module with params using the `--params` flag:

```shell
tinyport scaffold module launch --params minLaunch:uint,maxLaunch:int
```

After the parameters are scaffolded, change the `x/<module>/types/params.go` file to set the default values and validate the fields. 

The params module supports all [built-in Tinyport types](types.md).

## Params types

| Type   | Code type | Description             |
| ------ | --------- | ----------------------- |
| string | string    | Text type               |
| bool   | bool      | Boolean type            |
| int    | int32     | Integer number          |
| uint   | uint64    | Unsigned integer number |