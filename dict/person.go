package dict

import (
	"code.google.com/p/go.text/encoding/japanese"
	"code.google.com/p/go.text/transform"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	// Thanks for Mt. Erech Ave.
	// http://www.geocities.jp/mt_erech_ave/data.html
	// *** Own your risk for scraping ***
	personDocUrls = []string{
		"http://www.geocities.jp/mt_erech_ave/english_boy.html",
		"http://www.geocities.jp/mt_erech_ave/english_girl.html",
		"http://www.geocities.jp/mt_erech_ave/english_family1.html",
		"http://www.geocities.jp/mt_erech_ave/english_family2.html",
		"http://www.geocities.jp/mt_erech_ave/france_boy.html",
		"http://www.geocities.jp/mt_erech_ave/france_girl.html",
		"http://www.geocities.jp/mt_erech_ave/france_family.html",
		"http://www.geocities.jp/mt_erech_ave/germany_boy.html",
		"http://www.geocities.jp/mt_erech_ave/germany_girl.html",
		"http://www.geocities.jp/mt_erech_ave/germany_family.html",
		"http://www.geocities.jp/mt_erech_ave/italiano_boy.html",
		"http://www.geocities.jp/mt_erech_ave/italiano_girl.html",
		"http://www.geocities.jp/mt_erech_ave/italiano_family.html",
		"http://www.geocities.jp/mt_erech_ave/spain_boy.html",
		"http://www.geocities.jp/mt_erech_ave/spain_girl.html",
		"http://www.geocities.jp/mt_erech_ave/spain_family.html",
		"http://www.geocities.jp/mt_erech_ave/greek_boy.html",
		"http://www.geocities.jp/mt_erech_ave/greek_girl.html",
		"http://www.geocities.jp/mt_erech_ave/finnish_boy.html",
		"http://www.geocities.jp/mt_erech_ave/finnish_girl.html",
		"http://www.geocities.jp/mt_erech_ave/finnish_family.html",
		"http://www.geocities.jp/mt_erech_ave/russian_boy.html",
		"http://www.geocities.jp/mt_erech_ave/russian_girl.html",
		"http://www.geocities.jp/mt_erech_ave/russian_family.html",
	}

	alphabetConvertMaps = func() map[string]string {
		maps := map[string]string{
			"－": "-",
			"　": " ",
		}

		i := 0
		for ch := 'Ａ'; ch <= 'Ｚ'; ch++ {
			maps[fmt.Sprintf("%c", ch)] = fmt.Sprintf("%c", ch-'Ａ'+'A')
			i++
		}

		i = 0
		for ch := 'ａ'; ch <= 'ｚ'; ch++ {
			maps[fmt.Sprintf("%c", ch)] = fmt.Sprintf("%c", ch-'ａ'+'a')
			i++
		}

		return maps
	}()
)

func PersonDict() (dict []string) {
	for _, docUrl := range personDocUrls {
		doc := utf8DocByEucjpUrl(docUrl)
		doc.Find("table tbody tr").Each(func(_ int, tr *goquery.Selection) {
			tr.Find("td").Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					word := s.Text()
					word = fullWidthToHalfWidth(word)
					if strings.Count(word, " ") > 0 || strings.Count(word, ">") > 0 {
						return
					}
					dict = append(dict, word)
				}
			})
		})

		// 1 sec request interval
		time.Sleep(time.Second)
	}
	return
}

func fullWidthToHalfWidth(fullWidth string) (result string) {
	result = fullWidth
	for o, n := range alphabetConvertMaps {
		result = strings.Replace(result, o, n, -1)
	}
	return
}

func utf8DocByEucjpUrl(eucjpUrl string) *goquery.Document {
	resp, err := http.Get(eucjpUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(transform.NewReader(resp.Body, japanese.EUCJP.NewDecoder()))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
