package starportcmd

import (
	"github.com/spf13/cobra"

	"github.com/notional-labs/tinyport/docs"
	"github.com/notional-labs/tinyport/tinyport/pkg/localfs"
	"github.com/notional-labs/tinyport/tinyport/pkg/markdownviewer"
)

func NewDocs() *cobra.Command {
	c := &cobra.Command{
		Use:   "docs",
		Short: "Show Tinyport docs",
		Args:  cobra.NoArgs,
		RunE:  docsHandler,
	}
	return c
}

func docsHandler(cmd *cobra.Command, args []string) error {
	path, cleanup, err := localfs.SaveTemp(docs.Docs)
	if err != nil {
		return err
	}
	defer cleanup()

	return markdownviewer.View(path)
}
