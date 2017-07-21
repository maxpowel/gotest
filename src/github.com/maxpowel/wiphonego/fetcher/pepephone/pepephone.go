package pepephone

import (
	"net/url"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/maxpowel/wiphonego"

)

type Fetcher struct {
	Fetcher *wiphonego.WebFetcher
	Credentials *wiphonego.Credentials
}

func (f *Fetcher) login() {
	form := url.Values{}
	form.Add("request_uri", "/login")
	form.Add("email", f.Credentials.Username)
	form.Add("password", f.Credentials.Password)

	res, _ := f.Fetcher.Post("https://www.pepephone.com/login", form)
	f.isLogged(res)
	f.Fetcher.SaveCookies("cookies.json")
	//fmt.Println(f.fetcher.cookies())
	//data, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))
}

func (f *Fetcher) getInternetConsumption(phoneNumber string) (wiphonego.InternetConsumption, error){

	f.login()
f.Fetcher.LoadCookies("cookies.json")

	//time.Sleep(time.Second * 3)
	res, err := f.Fetcher.Get("https://www.pepephone.com/mipepephone")
	//fmt.Println(f.isLogged(res))
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(f.fetcher.cookies())
	//data, _ = ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))
	//ci := make(chan InternetConsumption)
	c := wiphonego.InternetConsumption{}
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//c := s.Find("span").Text()
		c.Consumed = 33
		fmt.Println("LEL")
		fmt.Println(s.Text())

	})
	fmt.Println("ESPERANDO CHANNEL")
	//c := <- ci
	return c, nil
	//f.fetcher.LoadCookies("cookies.json")

}

func (f *Fetcher) isLogged(res *http.Response) (bool){
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	return !strings.Contains(string(data), "Para acceder a Mi Pepephone")
}



func NewFetcher (credentials *wiphonego.Credentials) *Fetcher{
	return &Fetcher{
		Fetcher: wiphonego.NewWebFetcher(&url.URL{Host:"www.pepephone.com", Scheme:"https"}),
		Credentials: credentials,
	}
}