package config

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/palantir/stacktrace"
)

// ReadConfigPaths ...
func (c *ConfigFactory) ReadConfigPaths(paths []string, extension ConfigExtension) (map[string]Config, error) {

	result := make(map[string]Config)
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			err = stacktrace.Propagate(err, "could not open file at '%s'", path)
			return nil, err
		}
		fi, err := f.Stat()
		if err != nil {
			f.Close()
			err = stacktrace.Propagate(err, "could not stat file at '%s'", path)
			return nil, err
		}
		if !fi.IsDir() {
			if filepath.Ext(fi.Name()) == "."+JSON.String() {
				config, err := DecodeJSONConfig(f)
				if err != nil {
					f.Close()
					err = stacktrace.Propagate(err, "could not decode file at '%s'", path)
					return nil, err
				}
				result[path] = *config
			}
			if filepath.Ext(fi.Name()) == "."+CONF.String() {
				config, err := c.DecodeRawConfig(f)
				if err != nil {
					f.Close()
					err = stacktrace.Propagate(err, "could not decode file at '%s'", path)
					return nil, err
				}
				result[path] = *config
			}
			continue
		}
		contents, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			err = stacktrace.Propagate(err, "Error reading directory '%s'", path)
			return nil, err
		}
		sort.Sort(dirEnts(contents))
		for _, fi := range contents {
			// directory
			if fi.IsDir() {
				continue
			}
			subpath := filepath.Join(path, fi.Name())
			f, err := os.Open(subpath)
			if err != nil {
				err = stacktrace.Propagate(err, "could not read nested directory '%s'", subpath)
				return nil, err
			}
			if f == nil {
				err = stacktrace.NewError("could not get a file handle for %v", subpath)
				return nil, err
			}
			if filepath.Ext(fi.Name()) == "."+JSON.String() {
				config, err := DecodeJSONConfig(f)
				if err != nil {
					f.Close()
					err = stacktrace.Propagate(err, "could not decode file at '%s'", subpath)
					return nil, err
				}
				result[subpath] = *config
			}
			if filepath.Ext(fi.Name()) == "."+CONF.String() {
				config, err := c.DecodeRawConfig(f)
				if err != nil {
					f.Close()
					err = stacktrace.Propagate(err, "could not decode file at '%s'", subpath)
					return nil, err
				}
				result[subpath] = *config
			}
			f.Close()
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
