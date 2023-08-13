package scraper

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// initializing a data structure to keep the scraped data
type LazadaItem struct {
	url, image, name, price string
}

func LazadaScrapper(urlInput string, limitPage int) []LazadaItem {
	// c := colly.NewCollector()
	var lazadaItems []LazadaItem
	var urlPaginate = urlInput
	var i = 1
	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node

	maxPage := getMaxPage(urlInput)
	fmt.Printf("Max page is %d\n", maxPage)
	if maxPage < limitPage {
		fmt.Printf("Change limit to %d\n", maxPage)
		limitPage = maxPage
	}

	for i <= limitPage {
		fmt.Printf("Loading: %d - %s\n", i, urlPaginate)

		chromedp.Run(ctx,
			chromedp.Navigate(urlPaginate),
			chromedp.Nodes(".Bm3ON", &nodes, chromedp.ByQueryAll),
		)

		// scraping data from each node
		var url, image, name, price string
		for _, node := range nodes {
			chromedp.Run(ctx,
				chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text(".RfADt", &name, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text("span.ooOxS", &price, chromedp.ByQuery, chromedp.FromNode(node)),
			)
			var lazadaItem = LazadaItem{}

			lazadaItem.url = url
			lazadaItem.image = image
			lazadaItem.name = name
			lazadaItem.price = price
			// fmt.Printf("Name: %s \n", lazadaItem.name)
			lazadaItems = append(lazadaItems, lazadaItem)
		}
		i++
		paginateAnchor := fmt.Sprintf("li.ant-pagination-item-%d", i)

		chromedp.Run(ctx,
			chromedp.Nodes(paginateAnchor, &nodes, chromedp.ByQueryAll),
		)

		for _, node := range nodes {
			chromedp.Run(ctx,
				chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
			)
		}
		urlPaginate = "https://www.lazada.vn" + url

	}

	return lazadaItems
}

func getMaxPage(url string) int {
	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node

	//Check Maximum Page
	chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Nodes("li.ant-pagination-item", &nodes, chromedp.ByQueryAll),
	)

	var pages []string
	var page string
	for _, node := range nodes {
		chromedp.Run(ctx,
			chromedp.Text("a", &page, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		pages = append(pages, page)
	}

	size := len(pages)
	maxPage, _ := strconv.Atoi(pages[size-1])
	return maxPage
}

func Create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ExportLazadaCSV(lazadaItems []LazadaItem) {
	// opening the CSV file
	now := time.Now()
	timeString := now.Format("2006-01-02")
	nano := strconv.Itoa(now.Nanosecond())
	fileName := fmt.Sprintf("csv/lazada/%s-%s.csv", timeString, nano)
	file, err := Create(fileName)
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"name",
		"price",
		"url",
		"image",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, lazadaItem := range lazadaItems {
		// converting a PokemonProduct to an array of strings
		record := []string{
			lazadaItem.name,
			lazadaItem.price,
			lazadaItem.url,
			lazadaItem.image,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}
