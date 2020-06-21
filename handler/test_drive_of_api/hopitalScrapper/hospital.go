package hospital

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly/v2"
)

func hop() {
	count := 0
	sliceHo := []string{}
	hosiptals := map[string][]string{}

	for x := 1; x <= 27; x++ {
		c := colly.NewCollector()
		count++
		c.OnHTML("div .company > h4 >  a[href]", func(e *colly.HTMLElement) {
			sliceHo = append(sliceHo, e.Attr("title"))
		})
		c.Visit(fmt.Sprintf("https://www.businesslist.com.ng/category/Hospitals/%d", x))
	}
	hosiptals["name"] = sliceHo
	filebyte, err := json.MarshalIndent(&hosiptals, " ", " ")
	if err != nil {
		log.Panic(err)
	}
	filebyte = bytes.Replace(filebyte, []byte("\\u003c"), []byte("<"), -1)
	filebyte = bytes.Replace(filebyte, []byte("\\u003e"), []byte(">"), -1)
	filebyte = bytes.Replace(filebyte, []byte("\\u0026"), []byte("&"), -1)
	ioutil.WriteFile("hospitalList.json", filebyte, 0766)
}
