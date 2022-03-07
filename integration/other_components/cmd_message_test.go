//go:build !relayer
// +build !relayer

package other_components_test

import (
	"path/filepath"
	"testing"

	envtest "github.com/notional-labs/tinyport/integration"
	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner/step"
)

func TestGenerateAnAppWithMessage(t *testing.T) {
	var (
		env  = envtest.New(t)
		path = env.Scaffold("blog")
	)

	env.Must(env.Exec("create a message",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"message",
				"do-foo",
				"text",
				"vote:int",
				"like:bool",
				"-r",
				"foo,bar:int,foobar:bool",
			),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a message with custom path",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"message",
				"app-path",
				"text",
				"vote:int",
				"like:bool",
				"-r",
				"foo,bar:int,foobar:bool",
				"--path",
				"blog",
				"--no-simulation",
			),
			step.Workdir(filepath.Dir(path)),
		)),
	))

	env.Must(env.Exec("should prevent creating an existing message",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "message", "do-foo", "bar"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("create a message with a custom signer name",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "message", "do-bar", "bar", "--signer", "bar-doer"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a custom field type",
		step.NewSteps(step.New(
			step.Exec("tinyport",
				"s",
				"type",
				"custom-type",
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
			),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a message with the custom field type",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "message", "foo-baz", "customField:CustomType"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a module",
		step.NewSteps(step.New(
			step.Exec("tinyport", "s", "module", "foo", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a message in a module",
		step.NewSteps(step.New(
			step.Exec(
				"tinyport",
				"s",
				"message",
				"do-foo",
				"text",
				"userIds:array.uint",
				"--module",
				"foo",
				"--desc",
				"foo bar foobar",
				"--response",
				"foo,bar:int,foobar:bool",
			),
			step.Workdir(path),
		)),
	))

	env.EnsureAppIsSteady(path)
}
