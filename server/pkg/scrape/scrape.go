package scrape

import (
	"net/http"
	"strings"

	dynamodb "server/pkg/db"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func SaveScrapingResults(c *gin.Context) {
	var form struct {
		URL      string   `form:"url"`
		Keywords []string `form:"keywords"`
	}

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sentences, err := scrapeWebpage(form.URL, form.Keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(sentences) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No sentences found for the given keywords."})
		return
	}

	err = dynamodb.SaveScrapingResults(sentences, form.Keywords, form.URL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"success": sentences})
}

func scrapeWebpage(url string, keywords []string) ([]string, error) {
	var sentences []string

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		for _, keyword := range keywords {
			if strings.Contains(text, keyword) {
				sentences = append(sentences, text)
				break
			}
		}
	})

	return sentences, nil
}

func GetScrapingResults(c *gin.Context) {
	results, err := dynamodb.GetScrapingResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}