package dict

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

var (
	norseDoc = func() *goquery.Document {
		sourceUrl := "http://en.wikipedia.org/wiki/List_of_people,_items_and_places_in_Norse_mythology"
		doc, err := goquery.NewDocument(sourceUrl)
		if err != nil {
			log.Fatal(err)
		}
		return doc
	}()

	norseHeaderIds = []string{
		"#Places",
		"#Events",
		"#Artifacts",
		"#People",
		"#Other_assorted_beings",
	}
)

func NorseDict() (dict []string) {
	for _, headerId := range norseHeaderIds {
		dict = append(dict, wordsForNorseHeaderId(headerId)...)
	}
	return
}

func wordsForNorseHeaderId(headerId string) (words []string) {
	div := norseDoc.Find(headerId).Closest("h2").Next()
	div.Find("ul li a").Each(func(_ int, s *goquery.Selection) {
		word := s.Text()
		if strings.Count(word, " ") > 0 || strings.Count(word, "[") > 0 {
			return
		}

		words = append(words, word)
	})
	return
}
