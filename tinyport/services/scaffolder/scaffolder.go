// Package scaffolder initializes Tinyport apps and modifies existing ones
// to add more features in a later time.
package scaffolder

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/notional-labs/tinyport/tinyport/chainconfig"
	sperrors "github.com/notional-labs/tinyport/tinyport/errors"
	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner"
	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner/step"
	"github.com/notional-labs/tinyport/tinyport/pkg/cosmosanalysis"
	"github.com/notional-labs/tinyport/tinyport/pkg/cosmosanalysis/module"
	"github.com/notional-labs/tinyport/tinyport/pkg/cosmosgen"
	"github.com/notional-labs/tinyport/tinyport/pkg/cosmosver"
	"github.com/notional-labs/tinyport/tinyport/pkg/giturl"
	"github.com/notional-labs/tinyport/tinyport/pkg/gocmd"
	"github.com/notional-labs/tinyport/tinyport/pkg/gomodule"
	"github.com/notional-labs/tinyport/tinyport/pkg/gomodulepath"
)

// Scaffolder is Tinyport app scaffolder.
type Scaffolder struct {
	// path of the app.
	path string

	// modpath represents the go module path of the app.
	modpath gomodulepath.Path

	// Version of the chain
	Version cosmosver.Version
}

// App creates a new scaffolder for an existent app.
func App(path string) (Scaffolder, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return Scaffolder{}, err
	}

	modpath, path, err := gomodulepath.Find(path)
	if err != nil {
		return Scaffolder{}, err
	}
	modfile, err := gomodule.ParseAt(path)
	if err != nil {
		return Scaffolder{}, err
	}
	if err := cosmosanalysis.ValidateGoMod(modfile); err != nil {
		return Scaffolder{}, err
	}

	version, err := cosmosver.Detect(path)
	if err != nil {
		return Scaffolder{}, err
	}

	if !version.IsFamily(cosmosver.Stargate) {
		return Scaffolder{}, sperrors.ErrOnlyStargateSupported
	}

	s := Scaffolder{
		path:    path,
		modpath: modpath,
		Version: version,
	}

	return s, nil
}

func owner(modulePath string) string {
	return strings.Split(modulePath, "/")[1]
}

func finish(path, gomodPath string) error {
	if err := protoc(path, gomodPath); err != nil {
		return err
	}
	if err := tidy(path); err != nil {
		return err
	}
	return fmtProject(path)
}

func protoc(projectPath, gomodPath string) error {
	if err := cosmosgen.InstallDependencies(context.Background(), projectPath); err != nil {
		return err
	}

	confpath, err := chainconfig.LocateDefault(projectPath)
	if err != nil {
		return err
	}
	conf, err := chainconfig.ParseFile(confpath)
	if err != nil {
		return err
	}

	options := []cosmosgen.Option{
		cosmosgen.WithGoGeneration(gomodPath),
		cosmosgen.IncludeDirs(conf.Build.Proto.ThirdPartyPaths),
	}

	// generate Vuex code as well if it is enabled.
	if conf.Client.Vuex.Path != "" {
		storeRootPath := filepath.Join(projectPath, conf.Client.Vuex.Path, "generated")

		options = append(options,
			cosmosgen.WithVuexGeneration(
				false,
				func(m module.Module) string {
					parsedGitURL, _ := giturl.Parse(m.Pkg.GoImportName)
					return filepath.Join(storeRootPath, parsedGitURL.UserAndRepo(), m.Pkg.Name, "module")
				},
				storeRootPath,
			),
		)
	}
	if conf.Client.OpenAPI.Path != "" {
		options = append(options, cosmosgen.WithOpenAPIGeneration(conf.Client.OpenAPI.Path))
	}

	return cosmosgen.Generate(context.Background(), projectPath, conf.Build.Proto.Path, options...)
}

func tidy(path string) error {
	return cmdrunner.
		New(
			cmdrunner.DefaultStderr(os.Stderr),
			cmdrunner.DefaultWorkdir(path),
		).
		Run(context.Background(),
			step.New(
				step.Exec(gocmd.Name(), "mod", "tidy"),
			),
		)
}

func fmtProject(path string) error {
	return cmdrunner.
		New(
			cmdrunner.DefaultStderr(os.Stderr),
			cmdrunner.DefaultWorkdir(path),
		).
		Run(context.Background(),
			step.New(
				step.Exec(
					gocmd.Name(),
					"fmt",
					"./...",
				),
			),
		)
}
