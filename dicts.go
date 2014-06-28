package main

import (
	"github.com/k0kubun/grimoire/dict"
)

type DictFuncMap struct {
	Name     string
	DictFunc func() []string
}

var dictFuncMaps = []DictFuncMap{
	{"english", dict.EnglishDict},
	{"greek", dict.GreekDict},
	{"norse", dict.NorseDict},
	{"person", dict.PersonDict},
}

func allDict() (dict []string) {
	name := "all"

	if isCached(name) {
		dict = loadDictByName(name)
	} else {
		all := []string{}
		for _, dictFuncMap := range dictFuncMaps {
			all = append(all, loadDictByName(dictFuncMap.Name)...)
		}
		dict = formatDict(all)
		cacheDict(name, dict)
	}
	return dict
}

func loadDictByName(name string) (dict []string) {
	if isCached(name) {
		dict = cachedDict(name)
	} else {
		dict = dictFuncByName(name)()
		dict = formatDict(dict)
		cacheDict(name, dict)
	}
	return
}

func dictFuncByName(name string) func() []string {
	for _, dictFuncMap := range dictFuncMaps {
		if dictFuncMap.Name == name {
			return dictFuncMap.DictFunc
		}
	}
	return func() []string { return []string{} }
}
