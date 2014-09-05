package stream

import (
	"errors"
	"github.com/barakmich/glog"
)

type StreamAdapterStore interface {
	Name() string
	Connect()
	Subscribe([]string, func(bool, string, string))
	Publish(string)
}

type Options map[string]interface{}

func (d Options) IntKey(key string) (int, bool) {
	if val, ok := d[key]; ok {
		switch vv := val.(type) {
		case float64:
			return int(vv), true
		default:
			glog.Fatalln("Invalid", key, "parameter type from config.")
		}
	}
	return 0, false
}

func (d Options) StringKey(key string) (string, bool) {
	if val, ok := d[key]; ok {
		switch vv := val.(type) {
		case string:
			return vv, true
		default:
			glog.Fatalln("Invalid", key, "parameter type from config.")
		}
	}
	return "", false
}

func (d Options) BoolKey(key string) (bool, bool) {
	if val, ok := d[key]; ok {
		switch vv := val.(type) {
		case bool:
			return vv, true
		default:
			glog.Fatalln("Invalid", key, "parameter type from config.")
		}
	}
	return false, false
}

type NewStoreFunc func(Options) (StreamAdapterStore, error)
type InitStoreFunc func(Options) error

type register struct {
	newFunc  NewStoreFunc
	initFunc InitStoreFunc
}

var storeRegistry = make(map[string]register)

func RegisterAdapterStore(name string, newFunc NewStoreFunc, initFunc InitStoreFunc) {
	if _, found := storeRegistry[name]; found {
		panic("already registered stream adapter " + name)
	}
	storeRegistry[name] = register{
		newFunc:  newFunc,
		initFunc: initFunc,
	}
}

func NewAdapterStore(name string, opts Options) (StreamAdapterStore, error) {
	r, registered := storeRegistry[name]
	if !registered {
		return nil, errors.New("adapterstore: name '" + string(name) + "' is not registered")
	}
	return r.newFunc(opts)
}

func InitAdapterStore(name string, opts Options) error {
	r, registered := storeRegistry[name]
	if registered {
		return r.initFunc(opts)
	}
	return errors.New("adapterstore: name '" + string(name) + "' is not registered")
}

func AdapterStores() []string {
	t := make([]string, 0, len(storeRegistry))
	for n := range storeRegistry {
		t = append(t, n)
	}
	return t
}
