package resource

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

const (
	PLUGIN_INDEXER  = 1
	PLUGIN_CRAWLER  = 2
	PLUGIN_SENSOR   = 3
	PLUGIN_SCENARIO = 4
	PLUGIN_PIPELINE = 5
	PLUGIN_HELPER   = 6
)

func GetPluginList() (pluginList []*Plugin, err error) {
	nodes, err := storage.FindNodesByLabel(RESOURCE_PLUGIN_LABEL)
	if err != nil {
		nodes = nil
	}

	var plugins []*Plugin

	for _, node := range nodes {
		id := strconv.Itoa(node.Id)

		if node.Data["name"] == nil {
			node.Data["name"] = ""
		}

		if node.Data["group"] == nil {
			node.Data["group"] = ""
		}

		if node.Data["version"] == nil {
			node.Data["version"] = ""
		}

		if node.Data["path"] == nil {
			node.Data["path"] = ""
		}

		if node.Data["license"] == nil {
			node.Data["license"] = ""
		}

		plugin := NewPlugin(id , node.Data["name"].(string), node.Data["group"].(string), node.Data["version"].(string), node.Data["path"].(string), node.Data["license"].(string))
		plugins = append(plugins, plugin)
	}

	if plugins == nil {
		plugins = make([]*Plugin, 0)
	}

	return plugins, nil
}

func GetPlugin(id string) (plugin *Plugin, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && Contains(node.Labels, RESOURCE_PLUGIN_LABEL) {

		if node.Data["name"] == nil {
			node.Data["name"] = ""
		}

		if node.Data["group"] == nil {
			node.Data["group"] = ""
		}

		if node.Data["version"] == nil {
			node.Data["version"] = ""
		}

		if node.Data["path"] == nil {
			node.Data["path"] = ""
		}

		if node.Data["license"] == nil {
			node.Data["license"] = ""
		}

		return NewPlugin(id , node.Data["name"].(string), node.Data["group"].(string), node.Data["version"].(string), node.Data["path"].(string), node.Data["license"].(string)), nil
	}

	return nil, errors.New("PluginId `"+id+"` not exist")
}

func AddPlugin(plugin *Plugin) (err error) {
	properties := graph.Props {
		"name": plugin.Name,
		"group": plugin.Group,
		"version": plugin.Version,
		"path": plugin.Path,
		"license": plugin.License,
	}

	_plugin, err := storage.NewNode(properties, RESOURCE_PLUGIN_LABEL)

	if err != nil {
		return err
	}

	plugin.Id = strconv.Itoa(_plugin.Id)

	return nil
}

func UpdatePlugin(plugin *Plugin) (err error) {
	_id, _ := strconv.Atoi(plugin.Id)

	err = storage.SetPropertyNode(_id, "name", plugin.Name)
	err = storage.SetPropertyNode(_id, "group", plugin.Group)
	err = storage.SetPropertyNode(_id, "version", plugin.Version)
	err = storage.SetPropertyNode(_id, "path", plugin.Path)
	err = storage.SetPropertyNode(_id, "license", plugin.License)

	return err
}

func DeletePlugin(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}

func NewPlugin(id string, name string, group string, version string, path string, license string) *Plugin {
	return &Plugin{id, name, group, version, path, license}
}
