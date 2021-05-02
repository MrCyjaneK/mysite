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
	port := flag.Int("port", 0, "server port")
	fetch := flag.Bool("fetch", false, "Should I update data?")
	flag.Parse()
	if *fetch {
		// Updating readme's
		log.Println("Downloading external data...")
		downloadReadme("https://git.mrcyjanek.net/mrcyjanek/btnet/raw/branch/master/README.md", "content/projects/btnet/index.md", `[ "btnet", "projects", "FOSS" ]`, "BTnet - Using BitTorrent to serve websites.")
		downloadReadme("https://git.mrcyjanek.net/mrcyjanek/jwapi/raw/branch/master/README.md", "content/projects/jwapi/index.md", `[ "jwapi", "FOSS" ]`, "JWAPI - Golang library, and JW Library FOSS replacement.")
		downloadReadme("https://git.mrcyjanek.net/mrcyjanek/simple-tor-file-server/raw/branch/master/README.md", "content/projects/simple-tor-file-server/index.md", `[ "TOR", "FOSS" ]`, "Simple Tor File Server - Minimal OnionShare replacement")
		downloadReadme("https://git.mrcyjanek.net/mrcyjanek/userbot.php/raw/branch/master/README.md", "content/projects/userbot_php/index.md", `[ "PHP", "Telegram", "FOSS" ]`, "userbot.php - Badly written Telegram Userbot using Madeline Proto")
		// testportal
		downloadReadme("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/README.md", "content/projects/testportal-multitool/index.md", `[ "TestPortal", "Hack" ]`, "TestPortal MultiTool - Hack your exams like a boss")
		downloadReadme("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/README.pl.md", "content/projects/testportal-multitool/README.pl.md", `[ "TestPortal", "Hack" ]`, "[PL] TestPortal MultiTool - Hack your exams like a boss")
		downloadReadme("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/README.en.md", "content/projects/testportal-multitool/README.en.md", `[ "TestPortal", "Hack" ]`, "[EN] TestPortal MultiTool - Hack your exams like a boss")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/flags/pl.svg", "content/projects/testportal-multitool/static/flags/pl.svg")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/flags/us.svg", "content/projects/testportal-multitool/static/flags/us.svg")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/6-plus-9-x.png", "content/projects/testportal-multitool/static/screenshots/6-plus-9-x.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/100leg.jpeg", "content/projects/testportal-multitool/static/screenshots/100leg.jpeg")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/bruh-you-cant-screenshot-time.png", "content/projects/testportal-multitool/static/screenshots/bruh-you-cant-screenshot-time.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/hihi-we-are-qt.png", "content/projects/testportal-multitool/static/screenshots/hihi-we-are-qt.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/im-in-baby.png", "content/projects/testportal-multitool/static/screenshots/im-in-baby.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/knoppers.png", "content/projects/testportal-multitool/static/screenshots/knoppers.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/oh-john-you-are-my-hero.png", "content/projects/testportal-multitool/static/screenshots/oh-john-you-are-my-hero.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/pada-pada-i-padl.png", "content/projects/testportal-multitool/static/screenshots/pada-pada-i-padl.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/polikarp-have-you-talk-to-the-other-death.png", "content/projects/testportal-multitool/static/screenshots/polikarp-have-you-talk-to-the-other-death.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/screenshots/senpai-im-honest.png", "content/projects/testportal-multitool/static/screenshots/senpai-im-honest.png")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/firefox.svg", "content/projects/testportal-multitool/static/firefox.svg")
		downloadFile("https://git.mrcyjanek.net/mrcyjanek/testportal-multitool/raw/branch/main/static/javascript.svg", "content/projects/testportal-multitool/static/javascript.svg")
	}
	log.Println("Starting...")

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

func downloadFile(url string, target string) {
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
	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
