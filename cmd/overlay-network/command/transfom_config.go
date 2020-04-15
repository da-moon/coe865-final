package command

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	daemon "github.com/da-moon/coe865-final/cmd/overlay-network/command/daemon"
	flags "github.com/da-moon/coe865-final/cmd/overlay-network/flags"
	cli "github.com/mitchellh/cli"
	stacktrace "github.com/palantir/stacktrace"
	"strings"
)

// TransformConfigCommand is a Command implementation that generates an encryption
// key.
type TransformConfigCommand struct {
	Ui   cli.Ui
	args []string
}

var _ cli.Command = &TransformConfigCommand{}

const entrypoint = "transform-config"

// Run ...
func (c *TransformConfigCommand) Run(_ []string) int {
	cmdFlags := flag.NewFlagSet(entrypoint, flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	var configFiles []string

	cmdFlags.Var((*flags.AppendSliceValue)(&configFiles), "config-file",
		"raw file to read config from")
	cmdFlags.Var((*flags.AppendSliceValue)(&configFiles), "config-dir",
		"directory of raw config files to read")
	if err := cmdFlags.Parse(c.args); err != nil {
		c.Ui.Error(fmt.Sprintf("[ERROR] could not parse arguments : %v", err))
		return -1
	}
	if len(configFiles) > 0 {
		err := ReadConfigPaths(configFiles)
		if err != nil {
			c.Ui.Error(fmt.Sprintf("[ERROR]: %s", err.Error()))
			return nil
		}
	}

	const length = 32
	key := make([]byte, length)
	n, err := rand.Reader.Read(key)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("[ERROR] could not read random data: %s", err))
		return 1
	}
	if n != length {
		c.Ui.Error(fmt.Sprintf("[ERROR] could not read enough entropy. Generate more entropy!"))
		return 1
	}
	c.Ui.Output(hex.EncodeToString(key))
	return 0
}

// Synopsis ...
func (c *TransformConfigCommand) Synopsis() string {
	return "transform a given config file to sane format"
}

// Help ...
func (c *TransformConfigCommand) Help() string {
	helpText := `
Usage: overlay-network transform-config
  reads a config file as defined in project specification and
  converts is into a normal mashalling format such as JSON.
  it stores the converted file with the same name.

    -config-file=foo       Path to a JSON file to read configuration from.
                           This can be specified multiple times.
    -config-dir=foo       Path to a JSON file to read configuration from.
                           This can be specified multiple times.

  sample config file (before transform) :

  1 100 10.2.2.1	; RCID ASN IP Address (local rc info)
  2	                ; No. of RC connected
  2 200 10.1.1.2	; RCID ASN IP Address
  3 300 11.1.1.2	; RCID ASN IP Address
  4	                ; No. of ASN connected
  10 2 5 	        ; ASN Mbps(link capacity) cost
  20 5 5	        ; ASN Mbps(link capacity) cost
  200 10 5          ; ASN Mbps(link capacity) cost
  300 10 5          ; ASN Mbps(link capacity) cost
`
	return strings.TrimSpace(helpText)
}

// ReadConfigPaths reads the paths in the given order to load configurations.
func ReadConfigPaths(paths []string) error {
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return stacktrace.NewError("Error reading '%s': %s", path, err)
		}

		fi, err := f.Stat()
		if err != nil {
			f.Close()
			return stacktrace.NewError("Error reading '%s': %s", path, err)
		}

		if !fi.IsDir() {
			config, err := DecodeConfig(f)
			f.Close()

			if err != nil {
				return nil, stacktrace.NewError("Error decoding '%s': %s", path, err)
			}

			result = MergeConfig(result, config)
			continue
		}

		contents, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, stacktrace.NewError("Error reading '%s': %s", path, err)
		}

		sort.Sort(dirEnts(contents))

		for _, fi := range contents {
			if fi.IsDir() {
				continue
			}

			if !strings.HasSuffix(fi.Name(), ".json") {
				continue
			}

			subpath := filepath.Join(path, fi.Name())
			f, err := os.Open(subpath)
			if err != nil {
				return nil, stacktrace.NewError("Error reading '%s': %s", subpath, err)
			}

			config, err := DecodeConfig(f)
			f.Close()

			if err != nil {
				return nil, stacktrace.NewError("Error decoding '%s': %s", subpath, err)
			}

			result = MergeConfig(result, config)
		}
	}

	return result, nil
}

type dirEnts []os.FileInfo

// Len ...
func (d dirEnts) Len() int {
	return len(d)
}

// Less ...
func (d dirEnts) Less(i, j int) bool {
	return d[i].Name() < d[j].Name()
}

// Swap ...
func (d dirEnts) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// DecodeConfig ...
func DecodeConfig(r io.Reader) (*daemon.Config, error) {

	var result daemon.Config
	return &result, nil
}
