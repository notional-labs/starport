package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	starportcmd "github.com/notional-labs/tinyport/tinyport/cmd"
	"github.com/notional-labs/tinyport/tinyport/pkg/clictx"
	"github.com/notional-labs/tinyport/tinyport/pkg/validation"
)

func main() {
	ctx := clictx.From(context.Background())

	err := starportcmd.New(ctx).ExecuteContext(ctx)

	if ctx.Err() == context.Canceled || err == context.Canceled {
		fmt.Println("aborted")
		return
	}

	if err != nil {
		var validationErr validation.Error

		if errors.As(err, &validationErr) {
			fmt.Println(validationErr.ValidationInfo())
		} else {
			fmt.Println(err)
		}

		os.Exit(1)
	}
}
