package main

import (
	"net/url"
	"time"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strconv"
)

type MasMovilFetcher struct {
	fetcher *WebFetcher
	credentials Credentials
}

func (f *MasMovilFetcher) login() {
	form := url.Values{}
	form.Add("action", "login")
	form.Add("url", "")
	form.Add("user", f.credentials.username)
	form.Add("password", f.credentials.password)

	f.fetcher.get("https://yosoymas.masmovil.es/validate/")
	time.Sleep(3 * time.Second)
	f.fetcher.post("https://yosoymas.masmovil.es/validate/", form)
	f.fetcher.SaveCookies("cookies.json")
	time.Sleep(1 * time.Second)
}

func (f *MasMovilFetcher) getInternetConsumption(phoneNumber string) (InternetConsumption, error){

	f.login()
	//f.fetcher.LoadCookies("cookies.json")
	res, err := f.fetcher.get("https://yosoymas.masmovil.es/consumo/?line="+phoneNumber)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	//ci := make(chan InternetConsumption)
	c := InternetConsumption{}
	doc.Find(".box-main-content").Find(".progress").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//c := s.Find("span").Text()
		re := regexp.MustCompile("([0-9]+)|(infinito)")
		r := re.FindAllString(s.Text(), -1)

		if i == 0 {

			//fmt.Printf("Megas gastados %v de %v\n", r[0], r[1])
			consumed, err := strconv.ParseInt(r[0],10, 64)
			if err == nil {
				c.consumed = consumed
			}

			total, err := strconv.ParseInt(r[1],10, 64)
			if err == nil {
				c.total = total
			}
		} else {
			//fmt.Printf("Minutos gastados %v de %v\n", r[0], r[1])
		}
	})
	//c := <- ci
	return c, nil
	//f.fetcher.LoadCookies("cookies.json")

}

func NewMasMovilFetcher (credentials Credentials) *MasMovilFetcher{
	return &MasMovilFetcher{
		fetcher: NewWebFetcher(&url.URL{Host:"yosoymas.masmovil.es", Scheme:"https"}),
		credentials: credentials,
	}
}