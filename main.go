package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	Scraper "github.com/TinTran96/go23-ecom-scraper/scraper"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	lazadaFlags := flag.String("lazada", "foo", "lazada Flag")
	choTotFlags := flag.String("chotot", "foo", "cho tot Flag")

	flag.Parse()
	var url = ""
	var lazadaItems []Scraper.LazadaItem
	var chototItems []Scraper.ChoTotItem
	limitPage, err := strconv.Atoi((flag.Args())[0])

	if err != nil {
		fmt.Println("Error: Wrong Page Format")
	} else {
		if *lazadaFlags != "foo" {
			url = *lazadaFlags
			lazadaItems = Scraper.LazadaScrapper(url, limitPage)
			Scraper.ExportLazadaCSV(lazadaItems)
		} else {
			if *choTotFlags != "foo" {
				url = *choTotFlags
				chototItems = Scraper.ChoTotScrapper(url, limitPage)
				Scraper.ExportChototCSV(chototItems)
			} else {
				fmt.Println("Unavailable Flag")
			}
		}
	}

	// url := "https://tiki.vn/dien-thoai-may-tinh-bang/c1789"
	// lazadaItems := tikiScrapper(url)
	// exportLazadaCSV(lazadaItems)

	// url := "https://www.lazada.vn/dien-thoai-di-dong/?spm=a2o4n.home.cate_1.1.d57c3bdcARpMip"
	// lazadaItems := lazadaScrapper(url)
	// exportLazadaCSV(lazadaItems)

	// url := "https://www.nhatot.com/mua-ban-can-ho-chung-cu"
	// choTotCars := choTotScrapper(url)
	// exportChototCSV(choTotCars)
}
