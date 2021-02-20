package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/FooSoft/goldsmith"
	"github.com/FooSoft/goldsmith-components/devserver"
	"github.com/FooSoft/goldsmith-components/plugins/absolute"
	"github.com/FooSoft/goldsmith-components/plugins/breadcrumbs"
	"github.com/FooSoft/goldsmith-components/plugins/collection"
	"github.com/FooSoft/goldsmith-components/plugins/document"
	"github.com/FooSoft/goldsmith-components/plugins/frontmatter"
	"github.com/FooSoft/goldsmith-components/plugins/index"
	"github.com/FooSoft/goldsmith-components/plugins/layout"
	"github.com/FooSoft/goldsmith-components/plugins/markdown"
	"github.com/FooSoft/goldsmith-components/plugins/summary"
	"github.com/FooSoft/goldsmith-components/plugins/syntax"
	"github.com/FooSoft/goldsmith-components/plugins/tags"
	"github.com/FooSoft/goldsmith-components/plugins/thumbnail"
	"github.com/PuerkitoBio/goquery"
)

func fixup(doc *goquery.Document) error {
	doc.Find("table").AddClass("table").Find("thead").AddClass("thead-light")
	doc.Find("blockquote").AddClass("blockquote")
	doc.Find("img[src*='thumb']").Each(func(i int, s *goquery.Selection) {
		thumbLink := s.ParentFiltered("a")
		thumbLink.AddClass("img-thumbnail", "img-thumbnail-inline")
		thumbLink.SetAttr("data-title", s.AttrOr("alt", ""))
		thumbLink.SetAttr("data-toggle", "lightbox")
		thumbLink.SetAttr("data-gallery", "gallery")
	})

	return nil
}

type builder struct{}

func (b *builder) Build(contentDir, buildDir, cacheDir string) {
	tagMeta := map[string]interface{}{
		"Area":        "tags",
		"CrumbParent": "tags",
		"Layout":      "tag",
	}

	indexMeta := map[string]interface{}{
		"Layout": "index",
	}

	errs := goldsmith.Begin(contentDir).
		Cache(cacheDir).
		Chain(frontmatter.New()).
		Chain(markdown.New()).
		Chain(summary.New()).
		Chain(collection.New()).
		Chain(index.New(indexMeta)).
		Chain(tags.New().IndexMeta(tagMeta)).
		Chain(breadcrumbs.New()).
		Chain(layout.New()).
		Chain(syntax.New().Placement(syntax.PlaceInline)).
		Chain(document.New(fixup)).
		Chain(thumbnail.New()).
		Chain(absolute.New()).
		End(buildDir)

	for _, err := range errs {
		log.Print(err)
	}
}

func main() {
	fmt.Println("http://127.0.0.1:8080/")
	port := flag.Int("port", 8080, "server port")
	flag.Parse()

	devserver.DevServe(new(builder), *port, "content", "build", "cache")
}
