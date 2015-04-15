package plugin

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/osext"
	"github.com/mitchellh/elasticfeed/elasticfeed"
	"github.com/mitchellh/elasticfeed/elasticfeed/plugin"
)

// EnvConfig is the global EnvironmentConfig we use to initialize the CLI.
var EnvConfig elasticfeed.EnvironmentConfig

type config struct {
	PluginMinPort              uint
	PluginMaxPort              uint

	Pipelines       map[string]string
	Indexers       map[string]string
	Crawlers       map[string]string
	Scenarios       map[string]string
}

// ConfigFile returns the default path to the configuration file. On
// Unix-like systems this is the ".elasticfeedconfig" file in the home directory.
// On Windows, this is the "elasticfeed.config" file in the application data
// directory.
func ConfigFile() (string, error) {
	return configFile()
}

// ConfigDir returns the configuration directory for Elasticfeed.
func ConfigDir() (string, error) {
	return configDir()
}

// Decodes configuration in JSON format from the given io.Reader into
// the config object pointed to.
func decodeConfig(r io.Reader, c *config) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(c)
}

// Discover discovers plugins.
//
// This looks in the directory of the executable and the CWD, in that
// order for priority.
func (c *config) Discover() error {
	// Next, look in the same directory as the executable. Any conflicts
	// will overwrite those found in our current directory.
	exePath, err := osext.Executable()
	if err != nil {
		log.Printf("[ERR] Error loading exe directory: %s", err)
	} else {
		if err := c.discover(filepath.Dir(exePath)); err != nil {
			return err
		}
	}

	// Look in the plugins directory
	dir, err := ConfigDir()
	if err != nil {
		log.Printf("[ERR] Error loading config directory: %s", err)
	} else {
		if err := c.discover(filepath.Join(dir, "plugins")); err != nil {
			return err
		}
	}

	// Look in the cwd.
	if err := c.discover("."); err != nil {
		return err
	}

	return nil
}

func (c *config) LoadPiplines(name string) (elasticfeed.Pipeline, error) {
	log.Printf("Loading provisioner: %s\n", name)
	bin, ok := c.Pipelines[name]
	if !ok {
		log.Printf("Pipelines not found: %s\n", name)
		return nil, nil
	}

	return c.pluginClient(bin).Pipeline()
}

func (c *config) LoadIndexers(name string) (elasticfeed.Indexer, error) {
	log.Printf("Loading provisioner: %s\n", name)
	bin, ok := c.Indexers[name]
	if !ok {
		log.Printf("Indexers not found: %s\n", name)
		return nil, nil
	}

	return c.pluginClient(bin).Indexer()
}

func (c *config) LoadCrawlers(name string) (elasticfeed.Crawler, error) {
	log.Printf("Loading provisioner: %s\n", name)
	bin, ok := c.Crawlers[name]
	if !ok {
		log.Printf("Crawlers not found: %s\n", name)
		return nil, nil
	}

	return c.pluginClient(bin).Crawler()
}

func (c *config) LoadScenarios(name string) (elasticfeed.Scenario, error) {
	log.Printf("Loading provisioner: %s\n", name)
	bin, ok := c.Scenarios[name]
	if !ok {
		log.Printf("Scenarios not found: %s\n", name)
		return nil, nil
	}

	return c.pluginClient(bin).Scenario()
}

func (c *config) discover(path string) error {
	var err error

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return err
		}
	}

	err = c.discoverSingle(
		filepath.Join(path, "elasticfeed-pipeline-*"), &c.Pipelines)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "elasticfeed-indexer-*"), &c.Indexers)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "elasticfeed-crawler-*"), &c.Crawlers)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "elasticfeed-scenario-*"), &c.Scenarios)
	if err != nil {
		return err
	}

	return nil
}

func (c *config) discoverSingle(glob string, m *map[string]string) error {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return err
	}

	if *m == nil {
		*m = make(map[string]string)
	}

	prefix := filepath.Base(glob)
	prefix = prefix[:strings.Index(prefix, "*")]
	for _, match := range matches {
		file := filepath.Base(match)

		// If the filename has a ".", trim up to there
		if idx := strings.Index(file, "."); idx >= 0 {
			file = file[:idx]
		}

		// Look for foo-bar-baz. The plugin name is "baz"
		plugin := file[len(prefix):]
		log.Printf("[DEBUG] Discoverd plugin: %s = %s", plugin, match)
		(*m)[plugin] = match
	}

	return nil
}

func (c *config) pluginClient(path string) *plugin.Client {
	originalPath := path

	// First attempt to find the executable by consulting the PATH.
	path, err := exec.LookPath(path)
	if err != nil {
		// If that doesn't work, look for it in the same directory
		// as the `packer` executable (us).
		log.Printf("Plugin could not be found. Checking same directory as executable.")
		exePath, err := osext.Executable()
		if err != nil {
			log.Printf("Couldn't get current exe path: %s", err)
		} else {
			log.Printf("Current exe path: %s", exePath)
			path = filepath.Join(filepath.Dir(exePath), filepath.Base(originalPath))
		}
	}

	// If everything failed, just use the original path and let the error
	// bubble through.
	if path == "" {
		path = originalPath
	}

	log.Printf("Creating plugin client for path: %s", path)
	var config plugin.ClientConfig
	config.Cmd = exec.Command(path)
	config.Managed = true
	config.MinPort = c.PluginMinPort
	config.MaxPort = c.PluginMaxPort
	return plugin.NewClient(&config)
}
