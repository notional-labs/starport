package scaffolder

import (
	"context"
	"errors"

	"github.com/gobuffalo/genny"

	"github.com/notional-labs/tinyport/tinyport/pkg/multiformatname"
	"github.com/notional-labs/tinyport/tinyport/pkg/placeholder"
	"github.com/notional-labs/tinyport/tinyport/pkg/xgenny"
	"github.com/notional-labs/tinyport/tinyport/templates/field"
	"github.com/notional-labs/tinyport/tinyport/templates/query"
)

// AddQuery adds a new query to scaffolded app
func (s Scaffolder) AddQuery(
	ctx context.Context,
	tracer *placeholder.Tracer,
	moduleName,
	queryName,
	description string,
	reqFields,
	resFields []string,
	paginated bool,
) (sm xgenny.SourceModification, err error) {
	// If no module is provided, we add the type to the app's module
	if moduleName == "" {
		moduleName = s.modpath.Package
	}
	mfName, err := multiformatname.NewName(moduleName, multiformatname.NoNumber)
	if err != nil {
		return sm, err
	}
	moduleName = mfName.LowerCase

	name, err := multiformatname.NewName(queryName)
	if err != nil {
		return sm, err
	}

	if err := checkComponentValidity(s.path, moduleName, name, true); err != nil {
		return sm, err
	}

	// Check and parse provided request fields
	if ok := containCustomTypes(reqFields); ok {
		return sm, errors.New("query request params can't contain custom type")
	}
	parsedReqFields, err := field.ParseFields(reqFields, checkGoReservedWord)
	if err != nil {
		return sm, err
	}

	// Check and parse provided response fields
	if err := checkCustomTypes(ctx, s.path, moduleName, resFields); err != nil {
		return sm, err
	}
	parsedResFields, err := field.ParseFields(resFields, checkGoReservedWord)
	if err != nil {
		return sm, err
	}

	var (
		g    *genny.Generator
		opts = &query.Options{
			AppName:     s.modpath.Package,
			AppPath:     s.path,
			ModulePath:  s.modpath.RawPath,
			ModuleName:  moduleName,
			OwnerName:   owner(s.modpath.RawPath),
			QueryName:   name,
			ReqFields:   parsedReqFields,
			ResFields:   parsedResFields,
			Description: description,
			Paginated:   paginated,
		}
	)

	// Scaffold
	g, err = query.NewStargate(tracer, opts)
	if err != nil {
		return sm, err
	}
	sm, err = xgenny.RunWithValidation(tracer, g)
	if err != nil {
		return sm, err
	}
	return sm, finish(opts.AppPath, s.modpath.RawPath)
}
