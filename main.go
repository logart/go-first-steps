package main

import (
	"net/http"
	"log"
	"golang.org/x/net/html"
	"fmt"
)

func main() {
	res, err := http.Get("http://price.ua/apple/apple_iphone_7_128gb/catc52t1m1548190.html?order=price_asc#prices")
	if err != nil {
		log.Fatal(err)
	}

	z := html.NewTokenizer(res.Body)
	defer res.Body.Close();
	printNext := false
	correctBlock := false

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isSpan := t.Data == "span"
			isDiv := t.Data == "div"

			if (isSpan && correctBlock) {
				for _, a := range t.Attr {
					if a.Key == "class" {
						if(a.Val == "price"){
							fmt.Println("We've found span tag")
							fmt.Println(t)
							fmt.Println("Found class:", a.Val)
							printNext = true
						}
						break
					}
				}
			}

			if (isDiv) {
				for _, a := range t.Attr {
					if a.Key == "class" {
						if(a.Val == "table-prices"){
							fmt.Println("We've found div tag")
							fmt.Println(t)
							fmt.Println("Found class:", a.Val)
							correctBlock = true
						}
						break
					}
				}
			}
		case tt == html.TextToken:
			t := z.Token()

			if(printNext){
				printNext = false
				fmt.Println("Span content is", t)
			}
		}
	}
}
