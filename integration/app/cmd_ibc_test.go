//go:build !relayer
// +build !relayer

package app_test

import (
	"path/filepath"
	"testing"

	envtest "github.com/notional-labs/tinyport/integration"
	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner/step"
)

func TestCreateModuleWithIBC(t *testing.T) {

	var (
		env  = envtest.New(t)
		path = env.Scaffold("blogibc")
	)

	env.Must(env.Exec("create an IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "module", "foo", "--ibc", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create an IBC module with custom path",
		step.NewSteps(step.New(
			step.Exec("tinyport",
				"s",
				"module",
				"appPath",
				"--ibc",
				"--require-registration",
				"--path",
				"./blogibc",
			),
			step.Workdir(filepath.Dir(path)),
		)),
	))

	env.Must(env.Exec("create a type in an IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "list", "user", "email", "--module", "foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create an IBC module with an ordered channel",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"module",
				"orderedfoo",
				"--ibc",
				"--ordering",
				"ordered",
				"--require-registration",
			),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create an IBC module with an unordered channel",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"module",
				"unorderedfoo",
				"--ibc",
				"--ordering",
				"unordered",
				"--require-registration",
			),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a non IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "module", "non_ibc", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create an IBC module with dependencies",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"module",
				"example_with_dep",
				"--ibc",
				"--dep",
				"account,bank,staking,slashing",
				"--require-registration",
			),
			step.Workdir(path),
		)),
	))

	env.EnsureAppIsSteady(path)
}

func TestCreateIBCOracle(t *testing.T) {

	var (
		env  = envtest.New(t)
		path = env.Scaffold("ibcoracle")
	)

	env.Must(env.Exec("create an IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "module", "foo", "--ibc", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create an IBC module with params",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"module",
				"paramsFoo",
				"--ibc",
				"--params",
				"defaultName,isLaunched:bool,minLaunch:uint,maxLaunch:int",
				"--require-registration",
			),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create the first BandChain oracle integration",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "band", "oracleone", "--module", "foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create the second BandChain oracle integration",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "band", "oracletwo", "--module", "foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should prevent creating a BandChain oracle with no module specified",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "band", "invalidOracle"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating a BandChain oracle in a non existent module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "band", "invalidOracle", "--module", "nomodule"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("create a non-IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "module", "bar", "--params", "name,minLaunch:uint,maxLaunch:int", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should prevent creating a BandChain oracle in a non IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "band", "invalidOracle", "--module", "bar"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.EnsureAppIsSteady(path)
}

func TestCreateIBCPacket(t *testing.T) {

	var (
		env  = envtest.New(t)
		path = env.Scaffold("blogibc2")
	)

	env.Must(env.Exec("create an IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "module", "foo", "--ibc", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a packet",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"packet",
				"bar",
				"text",
				"texts:strings",
				"--module",
				"foo",
				"--ack",
				"foo:string,bar:int,baz:bool",
			),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should prevent creating a packet with no module specified",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "packet", "bar", "text"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating a packet in a non existent module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "packet", "bar", "text", "--module", "nomodule"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating an existing packet",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "packet", "bar", "post", "--module", "foo"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("create a packet with custom type fields",
		step.NewSteps(step.New(
			step.Exec("tinyport",
				"s",
				"packet",
				"ticket",
				"numInt:int",
				"numsInt:array.int",
				"numsIntAlias:ints",
				"numUint:uint",
				"numsUint:array.uint",
				"numsUintAlias:uints",
				"textString:string",
				"textStrings:array.string",
				"textStringsAlias:strings",
				"textCoin:coin",
				"textCoins:array.coin",
				"textCoinsAlias:coins",
				"victory:bool",
				"--module",
				"foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a custom field type",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "type", "custom-type", "customField:uint", "--module", "foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a packet with a custom field type",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "packet", "foo-baz", "customField:CustomType", "--module", "foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a packet with no send message",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "packet", "nomessage", "foo", "--no-message", "--module", "foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a packet with no field",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "packet", "empty", "--module", "foo"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a non-IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "module", "bar", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should prevent creating a packet in a non IBC module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "packet", "foo", "text", "--module", "bar"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.EnsureAppIsSteady(path)
}
