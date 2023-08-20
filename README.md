
# Web Scraper E-commerce

Data Scraping from E-commerce Websites: Chotot, and Lazada.

Demo Video: [Youtube](https://youtu.be/pY0gu1pWHsg)


## Tech Stack
- Go v1.20.6
- [chromedp](https://github.com/chromedp/chromedp)


## Run

Clone the project

```bash
  git clone https://github.com/TinTran96/go23-ecom-scraper.git
```

Go to the project directory

```bash
  cd go23-ecom-scraper
```

Run

For **Lazada** Link
```bash
  go run main.go -lazada https://www.lazada.vn/dien-thoai-di-dong/?spm=a2o4n.home.cate_1.1.d57c3bdcARpMip 1
```

For **ChoToT** Link
```bash
  go run main.go -chotot https://www.chotot.com/mua-ban-laptop-tp-ho-chi-minh 1
```


## Instruction

Currently, our tool support Lazada and Chotot site by search URL or category URL
For example:

https://www.lazada.vn/catalog/?q=laptop&_keyori=ss&from=input&spm=..search.go.

https://www.lazada.vn/trang-phuc-nam/?spm=a2o4n.home.cate_9.1.20183bdcbcdT9p

https://xe.chotot.com/mua-ban-xe-may-honda-tp-ho-chi-minh-sdmb1

https://www.chotot.com/mua-ban-dien-thoai-tp-ho-chi-minh

...

Our command is 

```bash
  go run main.go -chotot {url} {limitPage}
```

The solution to scraping data from Chotot & Lazada is **chromedp**, which is a library that provides browser capabilities and allows you to load a web page in a special browser with no GUI. You can then instruct the headless browser to mimic user interactions.

CSV File format: `YYYY-MM-DD-Nanosecond.csv`

To-do (future):
- Unable to access the link with "?page=2". I can only gather information from the original page 1.
- Not effectively optimizing the time it takes to scrape data.
- It is important to have a validation and error handling system in place for web pages in case errors occur.
- Lazada is capable of detecting the maximum page size, but Chotot does not possess this capability.


## Authors

- [@TinTran96](https://github.com/TinTran96)
