package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type DictCache struct {
	DictDatas []DictData
}

type DictData struct {
	Name string
	Dict []string
}

var (
	configFilePath = func() string {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		return filepath.Join(usr.HomeDir, ".grimoire")
	}()
)

func cacheDict(name string, dict []string) {
	dictData := DictData{
		Name: name,
		Dict: dict,
	}

	dictCache := dictCache()
	dictCache.DictDatas = append(dictCache.DictDatas, dictData)
	dictCache.DictDatas = uniqDictDatas(dictCache.DictDatas)

	config, err := json.Marshal(dictCache)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(configFilePath, config, 0644)
}

func dictCache() *DictCache {
	if !fileExists(configFilePath) {
		return &DictCache{DictDatas: []DictData{}}
	}

	config := DictCache{}
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&config)

	return &config
}

func isCached(name string) bool {
	if !fileExists(configFilePath) {
		return false
	}

	for _, dictData := range dictCache().DictDatas {
		if dictData.Name == name {
			return true
		}
	}
	return false
}

func uniqDictDatas(dictDatas []DictData) []DictData {
	uniqed := []DictData{}
	for _, dictData := range dictDatas {
		exists := false

		for _, uniqedData := range uniqed {
			if uniqedData.Name == dictData.Name {
				exists = true
				break
			}
		}

		if !exists {
			uniqed = append(uniqed, dictData)
		}
	}
	return uniqed
}

func cachedDict(dictName string) []string {
	if !fileExists(configFilePath) {
		return []string{}
	}

	for _, dictData := range dictCache().DictDatas {
		if dictData.Name == dictName {
			return dictData.Dict
		}
	}
	return []string{}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
