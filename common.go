package main

import (
	"encoding/xml"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var commandCommon = cli.Command{
	Name:  "common",
	Usage: "common English word list",
	Description: `
`,
	Action: commonDict,
}

type DicItemResult struct {
	TitleList     TitleList `xml:"TitleList"`
	TotalHitCount int       `xml:"TotalHitCount"`
	ItemCount     int       `xml:"ItemCount"`
}

type TitleList struct {
	DicItemTitles []DicItemTitle `xml:"DicItemTitle"`
}

type DicItemTitle struct {
	Title  Title `xml:"Title"`
	ItemId int   `xml:"ItemID"`
}

type Title struct {
	NetDicTitle string `xml:"span"`
}

func commonDict(c *cli.Context) {
	for _, w := range commonWords() {
		println(w)
	}
}

func commonWords() []string {
	words := []string{}
	for ch := 'a'; ch <= 'z'; ch++ {
		ws := wordsStartWith(ch)
		words = append(words, ws...)
	}
	return words
}

func wordsStartWith(ch rune) []string {
	totalCount := requestDejizoAPI(ch, 1).TotalHitCount
	result := requestDejizoAPI(ch, totalCount)

	words := []string{}
	for _, dicItemTitle := range result.TitleList.DicItemTitles {
		word := dicItemTitle.Title.NetDicTitle
		words = append(words, strings.Split(word, ",")...)
	}

	return words
}

// Thanks for dejizo RESTful API
// http://dejizo.jp/dev/rest.html
func requestDejizoAPI(ch rune, count int) *DicItemResult {
	apiEndPoint := "http://public.dejizo.jp/NetDicV09.asmx/SearchDicItemLite"
	params := url.Values{
		"Dic":       {"EJdict"},
		"Word":      {fmt.Sprintf("%c", ch)},
		"Scope":     {"HEADWORD"},
		"Match":     {"STARTWITH"},
		"Merge":     {"OR"},
		"Prof":      {"XHTML"},
		"PageSize":  {fmt.Sprintf("%d", count)},
		"PageIndex": {"0"},
	}
	requestUrl := strings.Join([]string{apiEndPoint, "?", params.Encode()}, "")

	resp, err := http.Get(requestUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := DicItemResult{}
	xml.Unmarshal(buf, &result)

	return &result
}
