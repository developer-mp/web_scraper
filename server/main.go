package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/scrape", handleScrape)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleScrape(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	url := r.Form.Get("url")
	keywords := strings.Split(r.Form.Get("keywords"), ",")

	article, err := scrapePage(url, keywords)
	if err != nil {
		http.Error(w, "Failed to scrape page", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"title": "%s", "content": "%s"}`, article.Title, article.Content)
}

func scrapePage(url string, keywords []string) (*Article, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()

	var content strings.Builder
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		content.WriteString(s.Text())
		content.WriteString("\n")
	})

	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(title), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(content.String()), strings.ToLower(keyword)) {
			return &Article{Title: title, Content: content.String()}, nil
		}
	}

	return nil, fmt.Errorf("keywords not found")
}
