package dict

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

var (
	greekDoc = func() *goquery.Document {
		sourceUrl := "http://en.wikipedia.org/wiki/List_of_Greek_mythological_figures"
		doc, err := goquery.NewDocument(sourceUrl)
		if err != nil {
			log.Fatal(err)
		}
		return doc
	}()

	greekHeaderIds = []string{
		"#Gigantes_and_other_\\.22giants\\.22",
		"#Personified_concepts",
		"#Chthonic_deities",
		"#Sea_deities",
		"#Sky_deities",
		"#Rustic_deities",
		"#Agricultural_deities",
		"#Health_deities",
		"#Other_deities",
		"#Deified_mortals",
		"#Heroes",
		"#Notable_women",
		"#Kings",
		"#Seers",
		"#Amazons",
		"#Inmates_of_Tartarus",
	}
)

func GreekDict() (dict []string) {
	for _, headerId := range greekHeaderIds {
		dict = append(dict, wordsForGreekHeaderId(headerId)...)
	}
	return
}

func wordsForGreekHeaderId(headerId string) (words []string) {
	div := greekDoc.Find(headerId).Closest("h3").Next()
	div.Find("ul li a").Each(func(_ int, s *goquery.Selection) {
		word := s.Text()
		if strings.Count(word, " ") > 0 || strings.Count(word, "[") > 0 {
			return
		}

		words = append(words, word)
	})
	return
}
