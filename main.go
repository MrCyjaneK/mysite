package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	// Updating readme's
	downloadReadme("https://git.mrcyjanek.net/mrcyjanek/btnet/raw/branch/master/README.md", "content/projects/btnet/index.md", `[ "btnet", "projects", "FOSS" ]`, "BTnet - Using BitTorrent to serve websites.")
	downloadReadme("https://git.mrcyjanek.net/mrcyjanek/jwapi/raw/branch/master/README.md", "content/projects/jwapi/index.md", `[ "jwapi", "FOSS" ]`, "JWAPI - Golang library, and JW Library FOSS replacement.")
	downloadReadme("https://git.mrcyjanek.net/mrcyjanek/simple-tor-file-server/raw/branch/master/README.md", "content/projects/simple-tor-file-server/index.md", `[ "TOR", "FOSS" ]`, "Simple Tor File Server - Minimal OnionShare replacement")
	downloadReadme("https://git.mrcyjanek.net/mrcyjanek/userbot.php/raw/branch/master/README.md", "content/projects/userbot_php/index.md", `[ "PHP", "Telegram", "FOSS" ]`, "userbot.php - Badly written Telegram Userbot using Madeline Proto")
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	if *port == 0 {
		new(builder).Build("content", "build", "cache")
	} else {
		devserver.DevServe(new(builder), *port, "content", "build", "cache")
	}
}

func downloadReadme(url string, target string, tags string, title string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(`+++
Area = "portfolio"
CrumbParent = "portfolio"
Layout = "page"
Tags = ` + tags + `
Title = "` + title + `"
+++

`)
	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
