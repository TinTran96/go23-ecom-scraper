package scraper

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

// initializing a data structure to keep the scraped data
type ChoTotItem struct {
	url, image, name, price, state string
}

func ChoTotScrapper(urlInput string, limitPage int) []ChoTotItem {
	// c := colly.NewCollector()
	var choTotItems []ChoTotItem
	var urlPaginate = urlInput
	var i = 1
	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var nodes []*cdp.Node

	for i <= limitPage {
		fmt.Printf("Loading: %d - %s\n", i, urlPaginate)

		chromedp.Run(ctx,
			chromedp.Navigate(urlPaginate),
			chromedp.KeyEvent(kb.PageDown),
			chromedp.Sleep(2*time.Second),
			chromedp.KeyEvent(kb.PageDown),
			chromedp.Sleep(2*time.Second),
			chromedp.KeyEvent(kb.PageDown),
			chromedp.Sleep(2*time.Second),
			chromedp.KeyEvent(kb.PageDown),
			chromedp.Sleep(2*time.Second),
			chromedp.KeyEvent(kb.PageDown),
			chromedp.Sleep(2*time.Second),
			chromedp.Nodes("li.AdItem_wrapperAdItem__S6qPH", &nodes, chromedp.ByQueryAll),
		)

		// scraping data from each node
		// var url, image, name, price, state string
		var name, url, image, price, state string
		for _, node := range nodes {
			chromedp.Run(ctx,
				chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text("h3", &name, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text("p.AdBody_adPriceNormal___OYFU", &price, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text("span", &state, chromedp.ByQuery, chromedp.FromNode(node)),
			)

			var chototItem ChoTotItem

			chototItem.url = url
			chototItem.price = price
			chototItem.state = state
			chototItem.image = image
			chototItem.name = name

			choTotItems = append(choTotItems, chototItem)
		}
		i++

		urlPaginate = fmt.Sprintf("%s&page=%d", urlInput, i)
	}

	return choTotItems
}

func ExportChototCSV(choTotCars []ChoTotItem) {
	// opening the CSV file
	now := time.Now()
	timeString := now.Format("2006-01-02")
	nano := strconv.Itoa(now.Nanosecond())
	fileName := fmt.Sprintf("csv/chotot/%s-%s.csv", timeString, nano)
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
		"state",
		"url",
		"image",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, choTotCar := range choTotCars {
		// converting a PokemonProduct to an array of strings
		record := []string{
			choTotCar.name,
			choTotCar.price,
			choTotCar.state,
			choTotCar.url,
			choTotCar.image,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}
