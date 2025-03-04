package starportcmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/notional-labs/tinyport/tinyport/pkg/cosmosaccount"
)

func NewAccountDelete() *cobra.Command {
	c := &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete an account by name",
		Args:  cobra.ExactArgs(1),
		RunE:  accountDeleteHandler,
	}

	c.Flags().AddFlagSet(flagSetKeyringBackend())

	return c
}

func accountDeleteHandler(cmd *cobra.Command, args []string) error {
	name := args[0]

	ca, err := cosmosaccount.New(
		cosmosaccount.WithKeyringBackend(getKeyringBackend(cmd)),
	)
	if err != nil {
		return err
	}

	if err := ca.DeleteByName(name); err != nil {
		return err
	}

	fmt.Printf("Account %s deleted.\n", name)
	return nil
}
