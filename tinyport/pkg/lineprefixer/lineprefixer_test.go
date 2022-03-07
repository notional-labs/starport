package lineprefixer

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {
	logs := `hello,
this
is
Tinyport!`
	buf := bytes.Buffer{}
	w := NewWriter(&buf, func() string { return "[TENDERMINT] " })
	_, err := io.Copy(w, strings.NewReader(logs))
	require.NoError(t, err)
	require.Equal(t, `[TENDERMINT] hello,
[TENDERMINT] this
[TENDERMINT] is
[TENDERMINT] Tinyport!`,
		buf.String(),
	)
}
