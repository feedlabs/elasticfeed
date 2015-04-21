package resource

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
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

		plugin := NewPlugin(id , node.Data["name"].(string), node.Data["group"].(string), node.Data["version"].(string), node.Data["path"].(string))
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

		return NewPlugin(id , node.Data["name"].(string), node.Data["group"].(string), node.Data["version"].(string), node.Data["path"].(string)), nil
	}

	return nil, errors.New("PluginId `"+id+"` not exist")
}

func AddPlugin(plugin *Plugin) (err error) {
	properties := graph.Props {
		"name": plugin.Name,
		"group": plugin.Group,
		"version": plugin.Version,
		"path": plugin.Path,
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
	return storage.SetPropertyNode(_id, "name", plugin.Name)
}

func DeletePlugin(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}

func NewPlugin(id string, name string, group string, version string, path string) *Plugin {
	return &Plugin{id, name, group, version, path}
}
