package plugin

import (
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/feedlabs/elasticfeed/plugin/model"
	"github.com/feedlabs/elasticfeed/common/config"

	"github.com/mitchellh/osext"

	emodel "github.com/feedlabs/elasticfeed/elasticfeed/model"
)

type PluginManager struct {
	engine 			emodel.Elasticfeed

	Indexers        map[string]string
	Crawlers        map[string]string
	Sensors         map[string]string
	Pipelines       map[string]string
	Scenarios       map[string]string
	Helpers         map[string]string

	PluginMinPort              uint
	PluginMaxPort              uint
}

func (this *PluginManager) FindPlugin(name string, profiler *model.Profiler) *interface{} {
	return nil
}

func (this *PluginManager) RunPlugin(p Plugin) (err error) {
	err = p.Run()

	if err != nil {
		return err
	}

	return nil
}

// Discover discovers plugins.
//
// This looks in the directory of the executable and the CWD, in that
// order for priority.
func (c *PluginManager) Discover() error {
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
	dir, err := config.ConfigDir()
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

func (c *PluginManager) LoadIndexer(name string) (model.Indexer, error) {
	log.Printf("Loading indexer: %s\n", name)
	bin, ok := c.Indexers[name]
	if !ok {
		log.Printf("Indexer not found: %s\n", name)
		return nil, nil
	}

	return c.pluginClient(bin).Indexer()
}

func (c *PluginManager) LoadPipeline(name string) (model.Pipeline, error) {
	log.Printf("Loading pipeline: %s\n", name)
	bin, ok := c.Pipelines[name]
	if !ok {
		log.Printf("Pipeline not found: %s\n", name)
		return nil, nil
	}

	return c.pluginClient(bin).Pipeline()
}

func (c *PluginManager) discover(path string) error {
	var err error

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return err
		}
	}

	err = c.discoverSingle(
		filepath.Join(path, "pipeline-*"), &c.Pipelines)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "indexer-*"), &c.Indexers)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "crawler-*"), &c.Crawlers)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "scenario-*"), &c.Scenarios)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "helper-*"), &c.Helpers)
	if err != nil {
		return err
	}

	err = c.discoverSingle(
		filepath.Join(path, "sensor-*"), &c.Sensors)
	if err != nil {
		return err
	}

	return nil
}

func (c *PluginManager) discoverSingle(glob string, m *map[string]string) error {
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

func (c *PluginManager) pluginClient(path string) *Client {
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
	var config ClientConfig
	config.Cmd = exec.Command(path)
	config.Managed = true
	config.MinPort = c.PluginMinPort
	config.MaxPort = c.PluginMaxPort
	return NewClient(&config)
}

func NewPluginManager(engine emodel.Elasticfeed) emodel.PluginManager {

	pm := &PluginManager{engine, nil, nil, nil, nil, nil, nil, config.GetPluginPortMin(), config.GetPluginPortMax()}

	pm.discover(filepath.Join(config.GetPluginStoragePath()))

	return pm
}
