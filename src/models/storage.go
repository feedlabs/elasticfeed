package models

import (
	"errors"
	"strconv"
	"time"
)

func init() {
	Storages = make(map[string]*Storage)
	Storages["1"] = &Storage{"1", "foo"}
	Storages["2"] = &Storage{"2", "bar"}
	Storages["3"] = &Storage{"3", "foobar"}
}

func AddStorage(storage Storage) (StorageId string) {
	storage.StorageId = strconv.FormatInt(time.Now().UnixNano(), 10)
	Storages[storage.StorageId] = &storage
	return storage.StorageId
}

func GetStorage(StorageId string) (storage *Storage, err error) {
	if v, ok := Storages[StorageId]; ok {
		return v, nil
	}
	return nil, errors.New("StorageId Not Exist")
}

func GetStorageList() map[string]*Storage {
	return Storages
}

func UpdateStorage(StorageId string, Name string) (err error) {
	if v, ok := Storages[StorageId]; ok {
		v.Name = Name
		return nil
	}
	return errors.New("StorageId Not Exist")
}

func DeleteStorage(StorageId string) {
	delete(Storages, StorageId)
}
