package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"os"
)

type Item struct {
	Number int
	Color  string
}

const (
	tmplExt = ".tmpl"
	htmlExt = ".html"
)

func main() {
	outDir := os.Args[1]
	fmt.Println("outDir: ", outDir)
	generateFrom1to10(outDir)
	generateSquareOf1to10(outDir)
	generateListHtml(outDir)
	generateLinksHtml(outDir)
}

func generateHtml(tmplPath string, htmlPath string, data interface{}) error {
	tmpl := template.Must(template.ParseFiles(tmplPath))
	f, err := os.Create(htmlPath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func generateFrom1to10(outDir string) {
	filePrefix := "layout"
	tmplPath := filePrefix + tmplExt
	htmlPath := outDir + filePrefix + htmlExt
	items := make([]Item, 0, 0)

	for i := 1; i <= 10; i++ {
		items = append(items, Item{
			Number: i,
			Color:  fmt.Sprintf("#%06d", rand.Intn(999999)),
		})
	}
	err := generateHtml(tmplPath, htmlPath, struct {
		Items []Item
	}{items})
	if err != nil {
		log.Println("generateHtml: ", err)
		return
	}
}

type Layout2Item struct {
	Text string
}

func generateSquareOf1to10(outDir string) {
	filePrefix := "layout2"
	tmplPath := filePrefix + tmplExt
	htmlPath := outDir + filePrefix + htmlExt
	items := make([]Layout2Item, 0, 0)
	for i := 1; i <= 10; i++ {
		items = append(items, Layout2Item{
			fmt.Sprintf("%d2 = %d", i, i*i),
		})
	}
	err := generateHtml(tmplPath, htmlPath, struct{ Items []Layout2Item }{items})
	if err != nil {
		log.Println("generateHtml: ", err)
		return
	}
}

type Layout3Item struct {
	Text string
}

func generateListHtml(outDir string) {
	filePrefix := "layout3"
	tmplPath := filePrefix + tmplExt
	htmlPath := outDir + filePrefix + htmlExt
	ulItems := []Layout3Item{{"Tom Yum Kung"}, {"Khao Man Kai"}, {"Som Tum"}}
	olItems := []Layout3Item{{"Grape"}, {"Orange"}, {"Apple"}}

	err := generateHtml(tmplPath, htmlPath, struct {
		OlItems []Layout3Item
		UlItems []Layout3Item
	}{ulItems, olItems})
	if err != nil {
		log.Println("generateHtml: ", err)
		return
	}
}

type Layout6Item struct {
	Text string
	URL  string
}

func generateLinksHtml(outDir string) {
	filePrefix := "layout6"
	tmplPath := filePrefix + tmplExt
	htmlPath := outDir + filePrefix + htmlExt
	items := []Layout6Item{
		{"Google", "https://www.google.com/"},
		{"Facebook", "https://www.facebook.com/"},
		{"Github", "https://github.com/"},
		{"Golang", "https://golang.org/"},
	}

	tmpl := template.Must(template.ParseFiles(tmplPath))
	f, err := os.Create(htmlPath)
	if err != nil {
		log.Println("generateHtml:", err)
		return
	}
	err = tmpl.Execute(f, struct{ Links []Layout6Item }{Links: items})
	if err != nil {
		log.Println("generateHtml:", err)
		return
	}
}

//type Layout7Item {
//	ImagePath string
//	Text string
//}
