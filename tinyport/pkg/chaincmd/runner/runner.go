// Package chaincmdrunner provides a high level access to a blockchain's commands.
package chaincmdrunner

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	"github.com/notional-labs/tinyport/tinyport/pkg/chaincmd"
	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner"
	"github.com/notional-labs/tinyport/tinyport/pkg/cmdrunner/step"
	"github.com/notional-labs/tinyport/tinyport/pkg/lineprefixer"
	"github.com/notional-labs/tinyport/tinyport/pkg/truncatedbuffer"
)

// Runner provides a high level access to a blockchain's commands.
type Runner struct {
	chainCmd                      chaincmd.ChainCmd
	stdout, stderr                io.Writer
	daemonLogPrefix, cliLogPrefix string
}

// Option configures Runner.
type Option func(r *Runner)

// Stdout sets stdout for executed commands.
func Stdout(w io.Writer) Option {
	return func(runner *Runner) {
		runner.stdout = w
	}
}

// DaemonLogPrefix is a prefix added to app's daemon logs.
func DaemonLogPrefix(prefix string) Option {
	return func(runner *Runner) {
		runner.daemonLogPrefix = prefix
	}
}

// CLILogPrefix is a prefix added to app's cli logs.
func CLILogPrefix(prefix string) Option {
	return func(runner *Runner) {
		runner.cliLogPrefix = prefix
	}
}

// Stderr sets stderr for executed commands.
func Stderr(w io.Writer) Option {
	return func(runner *Runner) {
		runner.stderr = w
	}
}

// New creates a new Runner with cc and options.
func New(ctx context.Context, chainCmd chaincmd.ChainCmd, options ...Option) (Runner, error) {
	runner := Runner{
		chainCmd: chainCmd,
		stdout:   io.Discard,
		stderr:   io.Discard,
	}

	applyOptions(&runner, options)

	// auto detect the chain id and get it applied to chaincmd if auto
	// detection is enabled.
	if chainCmd.IsAutoChainIDDetectionEnabled() {
		status, err := runner.Status(ctx)
		if err != nil {
			return Runner{}, err
		}

		runner.chainCmd = runner.chainCmd.Copy(chaincmd.WithChainID(status.ChainID))
	}

	return runner, nil
}

func applyOptions(r *Runner, options []Option) {
	for _, apply := range options {
		apply(r)
	}
}

// Copy makes a copy of runner by overwriting its options with given options.
func (r Runner) Copy(options ...Option) Runner {
	applyOptions(&r, options)

	return r
}

// Cmd returns underlying chain cmd.
func (r Runner) Cmd() chaincmd.ChainCmd {
	return r.chainCmd
}

type runOptions struct {
	// wrappedStdErrMaxLen determines the maximum length of the wrapped error logs
	// this option is used for long running command to prevent the buffer containing stderr getting too big
	// 0 can be used for no maximum length
	wrappedStdErrMaxLen int

	// stdout and stderr used to collect a copy of command's outputs.
	stdout, stderr io.Writer

	// stdin defines input for the command
	stdin io.Reader
}

// run executes a command.
func (r Runner) run(ctx context.Context, runOptions runOptions, stepOptions ...step.Option) error {
	var (
		// we use a truncated buffer to prevent memory leak
		// this is because Stargate app currently send logs to StdErr
		// therefore if the app successfully starts, the written logs can become extensive
		errb = truncatedbuffer.NewTruncatedBuffer(runOptions.wrappedStdErrMaxLen)

		// add optional prefixes to output streams.
		stdout io.Writer = lineprefixer.NewWriter(r.stdout,
			func() string { return r.daemonLogPrefix },
		)
		stderr io.Writer = lineprefixer.NewWriter(r.stderr,
			func() string { return r.cliLogPrefix },
		)
	)

	// Set standard outputs
	if runOptions.stdout != nil {
		stdout = io.MultiWriter(stdout, runOptions.stdout)
	}
	if runOptions.stderr != nil {
		stderr = io.MultiWriter(stderr, runOptions.stderr)
	}

	stderr = io.MultiWriter(stderr, errb)

	runnerOptions := []cmdrunner.Option{
		cmdrunner.DefaultStdout(stdout),
		cmdrunner.DefaultStderr(stderr),
	}

	// Set standard input if defined
	if runOptions.stdin != nil {
		runnerOptions = append(runnerOptions, cmdrunner.DefaultStdin(runOptions.stdin))
	}

	err := cmdrunner.
		New(runnerOptions...).
		Run(ctx, step.New(stepOptions...))

	return errors.Wrap(err, errb.GetBuffer().String())
}

func newBuffer() *buffer {
	return &buffer{
		Buffer: new(bytes.Buffer),
	}
}

// buffer is a bytes.Buffer with additional features.
type buffer struct {
	*bytes.Buffer
}

// JSONEnsuredBytes ensures that encoding format for returned bytes is always
// JSON even if the written data is originally encoded in YAML.
func (b *buffer) JSONEnsuredBytes() ([]byte, error) {
	bytes := b.Buffer.Bytes()

	var out interface{}

	if err := yaml.Unmarshal(bytes, &out); err == nil {
		return yaml.YAMLToJSON(bytes)
	}

	return bytes, nil
}

type txResult struct {
	Code   int    `json:"code"`
	RawLog string `json:"raw_log"`
	TxHash string `json:"txhash"`
}

func decodeTxResult(b *buffer) (txResult, error) {
	var r txResult

	data, err := b.JSONEnsuredBytes()
	if err != nil {
		return r, err
	}

	return r, json.Unmarshal(data, &r)
}
