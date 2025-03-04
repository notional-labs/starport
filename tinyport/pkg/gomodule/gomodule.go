package gomodule

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"

	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner"
	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner/step"
)

// ErrGoModNotFound returned when go.mod file cannot be found for an app.
var ErrGoModNotFound = errors.New("go.mod not found")

// ParseAt finds and parses go.mod at app's path.
func ParseAt(path string) (*modfile.File, error) {
	gomod, err := os.ReadFile(filepath.Join(path, "go.mod"))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, ErrGoModNotFound
		}
		return nil, err
	}
	return modfile.Parse("", gomod, nil)
}

// FilterVersions filters dependencies under require section by their paths.
func FilterVersions(dependencies []module.Version, paths ...string) []module.Version {
	var filtered []module.Version

	for _, dep := range dependencies {
		for _, path := range paths {
			if dep.Path == path {
				filtered = append(filtered, dep)
				break
			}
		}
	}

	return filtered
}

func ResolveDependencies(f *modfile.File) ([]module.Version, error) {
	var versions []module.Version

	isReplacementAdded := func(rv module.Version) bool {
		for _, rep := range f.Replace {
			if rv.Path == rep.Old.Path {
				versions = append(versions, rep.New)

				return true
			}
		}

		return false
	}

	for _, req := range f.Require {
		if req.Indirect {
			continue
		}
		if !isReplacementAdded(req.Mod) {
			versions = append(versions, req.Mod)
		}
	}

	return versions, nil
}

// LocatePath locates pkg's absolute path managed by 'go mod' on the local filesystem.
func LocatePath(ctx context.Context, src string, pkg module.Version) (path string, err error) {
	// can be a local package.
	if pkg.Version == "" { // indicates that this is a local package.
		if filepath.IsAbs(pkg.Path) {
			return pkg.Path, nil
		}
		return filepath.Join(src, pkg.Path), nil
	}

	// otherwise, it is hosted.
	out := &bytes.Buffer{}

	if err := cmdrunner.
		New().
		Run(ctx, step.New(
			step.Exec("go", "mod", "download", "-json"),
			step.Workdir(src),
			step.Stdout(out),
		)); err != nil {
		return "", err
	}

	d := json.NewDecoder(out)

	for {
		var module struct {
			Path, Version, Dir string
		}
		if err := d.Decode(&module); err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		if module.Path == pkg.Path && module.Version == pkg.Version {
			return module.Dir, nil
		}
	}

	return "", fmt.Errorf("module %q not found", pkg.Path)
}
