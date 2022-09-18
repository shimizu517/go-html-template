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
	generateImageHtml(outDir)
	generateCompaniesTable(outDir)
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

type Layout7Item struct {
	ImagePath string
	Text      string
}

func generateImageHtml(outDir string) {
	filePrefix := "layout7"
	tmplPath := filePrefix + tmplExt
	htmlPath := outDir + filePrefix + htmlExt
	items := []Layout7Item{
		{"https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png", "Google"},
		{"https://www.facebook.com/images/fb_icon_325x325.png", "Facebook"},
		{"https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png", "Github"},
	}
	tmpl := template.Must(template.ParseFiles(tmplPath))
	f, err := os.Create(htmlPath)
	if err != nil {
		log.Println("generateHtml:", err)
		return
	}
	err = tmpl.Execute(f, struct{ Items []Layout7Item }{Items: items})
	if err != nil {
		log.Println("generateHtml:", err)
		return
	}
}

type CompanyTableCell struct {
	Text  string
	Style template.CSS
}

type CompaniesTable struct {
	TableTitle string
	Heads      []CompanyTableCell
	Rows       [][]CompanyTableCell
}

func generateCompaniesTable(outDir string) {
	filePrefix := "layout_companies_table"
	tmplPath := filePrefix + tmplExt
	htmlPath := outDir + filePrefix + htmlExt
	Heads := []CompanyTableCell{
		{"NAME", ""},
		{"SYMBOL", "text-align: center;"},
		{"CURRENT", "text-align: center;"},
		{"52WK HI", "text-align: center;"},
		{"52WK LO", "text-align: center;"},
		{"P/E RATIO", "text-align: center;"},
	}
	Rows := [][]CompanyTableCell{
		{{"Intel", ""}, {"INTC", "text-align: center;"}, {"28.84", "text-align: center;"}, {"56.28", "text-align: center;"}, {"28.42", "text-align: center;"}, {"56.09", "text-align: center;"}},
		{{"Apple", ""}, {"AAPL", "text-align: center;"}, {"152.37", "text-align: center;"}, {"182.94", "text-align: center;"}, {"129.04", "text-align: center;"}, {"24.90", "text-align: center;"}},
		{{"Verizon", ""}, {"VZ", "text-align: center;"}, {"41.03", "text-align: center;"}, {"55.51", "text-align: center;"}, {"40.66", "text-align: center;"}, {"8.28", "text-align: center;"}},
		{{"META", ""}, {"META", "text-align: center;"}, {"149.55", "text-align: center;"}, {"373.56", "text-align: center;"}, {"144.29", "text-align: center;"}, {"12.12", "text-align: center;"}},
		{{"Google", ""}, {"GOOG", "text-align: center;"}, {"103.90", "text-align: center;"}, {"152.10", "text-align: center;"}, {"101.86", "text-align: center;"}, {"19.63", "text-align: center;"}},
	}
	tmpl := template.Must(template.ParseFiles(tmplPath))
	f, err := os.Create(htmlPath)
	if err != nil {
		log.Println("generateHtml:", err)
		return
	}
	err = tmpl.Execute(f, CompaniesTable{
		TableTitle: "My tech stock picks",
		Heads:      Heads,
		Rows:       Rows,
	})
	if err != nil {
		log.Println("generateHtml:", err)
		return
	}
}
