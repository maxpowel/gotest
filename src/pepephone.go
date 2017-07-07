package main

import (
	"net/url"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"io/ioutil"
	"strings"

)

type PepephoneFetcher struct {
	fetcher *WebFetcher
	credentials Credentials
}

func (f *PepephoneFetcher) getInternetConsumption(phoneNumber string) (InternetConsumption, error){
	form := url.Values{}
	form.Add("request_uri", "/login")
	form.Add("email", f.credentials.username)
	form.Add("password", f.credentials.password)

	//res, _ := f.fetcher.post("https://www.pepephone.com/login", form)
	//f.fetcher.SaveCookies("cookies.json")
	//fmt.Println(f.fetcher.cookies())
	//data, _ := ioutil.ReadAll(res.Body)
	//data, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))
f.fetcher.LoadCookies("cookies.json")

	//time.Sleep(time.Second * 3)
	res, err := f.fetcher.get("https://www.pepephone.com/mipepephone")
	//fmt.Println(f.isLogged(res))
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(f.fetcher.cookies())
	//data, _ = ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))
	//ci := make(chan InternetConsumption)
	c := InternetConsumption{}
	doc.Find(".grpelem").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//c := s.Find("span").Text()
		c.consumed = 33
		fmt.Println("LEL")
		fmt.Println(s.Text())

	})
	fmt.Println("ESPERANDO CHANNEL")
	//c := <- ci
	return c, nil
	//f.fetcher.LoadCookies("cookies.json")

}

func (f *PepephoneFetcher) isLogged(res *http.Response) (bool){
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	return !strings.Contains(string(data), "Para acceder a Mi Pepephone")
}



func NewPepephoneFetcher (credentials Credentials) *PepephoneFetcher{
	return &PepephoneFetcher{
		fetcher: NewWebFetcher(&url.URL{Host:"www.pepephone.com", Scheme:"https"}),
		credentials: credentials,
	}
}