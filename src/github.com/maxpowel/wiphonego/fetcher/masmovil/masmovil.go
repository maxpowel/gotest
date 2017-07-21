package masmovil

import (
	"net/url"
	"time"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strconv"
	"github.com/maxpowel/wiphonego"
)

type Fetcher struct {
	Fetcher *wiphonego.WebFetcher
	Credentials *wiphonego.Credentials
}

func (f *Fetcher) login() (error) {
	form := url.Values{}
	form.Add("action", "login")
	form.Add("url", "")
	form.Add("user", f.Credentials.Username)
	form.Add("password", f.Credentials.Password)

	f.Fetcher.Get("https://yosoymas.masmovil.es/validate/")
	time.Sleep(3 * time.Second)
	f.Fetcher.Post("https://yosoymas.masmovil.es/validate/", form)
	f.Fetcher.SaveCookies("cookies.json")
	time.Sleep(1 * time.Second)

	return nil
}

func (f *Fetcher) GetInternetConsumption(phoneNumber string) (wiphonego.InternetConsumption, error){

	f.login()
	//f.fetcher.LoadCookies("cookies.json")
	res, err := f.Fetcher.Get("https://yosoymas.masmovil.es/consumo/?line="+phoneNumber)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	//ci := make(chan InternetConsumption)
	c := wiphonego.InternetConsumption{}
	doc.Find(".box-main-content").Find(".progress").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//c := s.Find("span").Text()
		re := regexp.MustCompile("([0-9]+)|(infinito)")
		r := re.FindAllString(s.Text(), -1)

		if i == 0 {

			//fmt.Printf("Megas gastados %v de %v\n", r[0], r[1])
			consumed, err := strconv.ParseInt(r[0],10, 64)
			if err == nil {
				c.Consumed = consumed
			}

			total, err := strconv.ParseInt(r[1],10, 64)
			if err == nil {
				c.Total = total
			}
		} else {
			//fmt.Printf("Minutos gastados %v de %v\n", r[0], r[1])
		}
	})
	//c := <- ci
	return c, nil
	//f.fetcher.LoadCookies("cookies.json")

}

func NewFetcher (credentials *wiphonego.Credentials) *Fetcher{
	return &Fetcher{
		Fetcher: wiphonego.NewWebFetcher(&url.URL{Host:"yosoymas.masmovil.es", Scheme:"https"}),
		Credentials: credentials,
	}
}