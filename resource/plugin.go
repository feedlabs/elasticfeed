package resource

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func GetPluginList() (pluginList []*Plugin, err error) {
	nodes, err := storage.FindNodesByLabel(RESOURCE_ORG_LABEL)
	if err != nil {
		nodes = nil
	}

	var plugins []*Plugin

	for _, node := range nodes {
		id := strconv.Itoa(node.Id)

		if node.Data["name"] == nil {
			node.Data["name"] = ""
		}

		plugin := NewPlugin(id , node.Data["name"].(string))
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

	if node != nil && Contains(node.Labels, RESOURCE_ORG_LABEL) {

		if node.Data["name"] == nil {
			node.Data["name"] = ""
		}

		return NewPlugin(id , node.Data["name"].(string)), nil
	}

	return nil, errors.New("PluginId `"+id+"` not exist")
}

func AddPlugin(plugin *Plugin) (err error) {
	properties := graph.Props {
		"name": plugin.Name,
		"version": plugin.Version,
		"path": plugin.Path,
	}

	_plugin, err := storage.NewNode(properties, RESOURCE_ORG_LABEL)

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

func NewPlugin(id string, name string) *Plugin {
	_type := "sensor"
	_version := "1"
	_path := "/tmp"

	return &Plugin{id, name, _type, _version, _path}
}
