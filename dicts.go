package main

import (
	"github.com/k0kubun/grimoire/dict"
)

type DictFuncMap struct {
	Name     string
	DictFunc func() []string
}

var dictFuncMaps = []DictFuncMap{
	{"common", dict.CommonDict},
	{"greek", dict.GreekDict},
	{"norse", dict.NorseDict},
	{"person", dict.PersonDict},
}

func allDict() []string {
	all := []string{}
	for _, dictFuncMap := range dictFuncMaps {
		all = append(all, loadDictByName(dictFuncMap.Name)...)
	}
	return all
}

func loadDictByName(name string) (dict []string) {
	if isCached(name) {
		dict = cachedDict(name)
	} else {
		dict = dictFuncByName(name)()
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
