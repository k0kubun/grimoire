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

func dictFuncByName(name string) func() []string {
	for _, dictFuncMap := range dictFuncMaps {
		if dictFuncMap.Name == name {
			return dictFuncMap.DictFunc
		}
	}
	return func() []string { return []string{} }
}
