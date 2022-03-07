package docs

import "embed"

// Docs are Tinyport docs.
//go:embed *.md */*.md
var Docs embed.FS
