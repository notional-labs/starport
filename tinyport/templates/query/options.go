package query

import (
	"github.com/notional-labs/tinyport/tinyport/pkg/multiformatname"
	"github.com/notional-labs/tinyport/tinyport/templates/field"
)

// Options ...
type Options struct {
	AppName     string
	AppPath     string
	ModuleName  string
	ModulePath  string
	OwnerName   string
	QueryName   multiformatname.Name
	Description string
	ResFields   field.Fields
	ReqFields   field.Fields
	Paginated   bool
}
