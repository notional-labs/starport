package giturl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	parsed, err := Parse("http://github.com/notional-labs/tinyport/a/b")
	require.NoError(t, err)
	require.Equal(t, "github.com", parsed.Host)
	require.Equal(t, "tendermint", parsed.User)
	require.Equal(t, "tinyport", parsed.Repo)
	require.Equal(t, "notional-labs/tinyport", parsed.UserAndRepo())
}
