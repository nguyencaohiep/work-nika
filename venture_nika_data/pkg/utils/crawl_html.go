package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func GetHtmlDomJsRenderByUrl(url string) *goquery.Document {
	var dom *goquery.Document

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			dom = r.HTMLDoc
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()

	return dom
}

func ConvertClassesFormatFromBrowserToGoQuery(input string) string {
	classes := input
	classes = `.` + classes
	classes = strings.ReplaceAll(classes, ` `, `.`)
	return classes
}
