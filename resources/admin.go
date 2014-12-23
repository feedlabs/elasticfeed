package resources

import (
	"errors"
	"strconv"

	"github.com/feedlabs/feedify/graph"
)

func init() {
	Admins = make(map[string]*Admin)
}

func GetAdminList() []*Admin {
	nodes, err := storage.FindNodesByLabel(RESOURCE_ADMIN_LABEL)
	if err != nil {
		nodes = nil
	}

	var admins []*Admin

	for _, node := range nodes {
		data := node.Data["data"].(string)
		id := strconv.Itoa(node.Id)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")

		admin := &Admin{id , data, len(rels)}
		admins = append(admins, admin)
	}

	return admins
}

func GetAdmin(id string) (admin *Admin, err error) {
	_id, err := strconv.Atoi(id)
	node, err := storage.Node(_id)

	if err != nil {
		return nil, err
	}

	if node != nil && contains(node.Labels, RESOURCE_ADMIN_LABEL) {
		data := node.Data["data"].(string)
		rels, _ := storage.RelationshipsNode(node.Id, "contains")
		return &Admin{strconv.Itoa(node.Id), data, len(rels)}, nil
	}

	return nil, errors.New("AdminId not exist")
}

func AddAdmin(admin Admin) (id string, err error) {
	properties := graph.Props{"data": admin.Data}
	_admin, err := storage.NewNode(properties, RESOURCE_ADMIN_LABEL)

	if err != nil {
		return "0", err
	}

	admin.Id = strconv.Itoa(_admin.Id)

	return admin.Id, nil
}

func UpdateAdmin(id string, data string) (err error) {
	_id, _ := strconv.Atoi(id)
	return storage.SetPropertyNode(_id, "data", data)
}

func DeleteAdmin(id string) (error) {
	_id, _ := strconv.Atoi(id)
	return storage.DeleteNode(_id)
}
