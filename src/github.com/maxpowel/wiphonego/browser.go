package wiphonego

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http/cookiejar"
	"net/http"

	"net/url"
	"bytes"
)

import (
	"golang.org/x/net/publicsuffix"

	"os"
	"regexp"

	"encoding/json"
	"io/ioutil"
)


type WebFetcher struct {
	Client *http.Client
	BaseUrl *url.URL

}

func (wb *WebFetcher) SaveCookies(path string) error {
	b, err := json.Marshal(wb.Cookies())
	ioutil.WriteFile(path, b, os.FileMode(0777))
	return err
}

func (wb *WebFetcher) LoadCookies(path string) error {
	a, err := ioutil.ReadFile("cookies.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	var cookies  []*http.Cookie
	json.Unmarshal(a, &cookies)
	wb.Client.Jar.SetCookies(wb.BaseUrl, cookies)
	return err
}


func NewWebFetcher(url *url.URL) *WebFetcher {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	cookieJar, _ := cookiejar.New(&options)

	client := &http.Client{
		Jar: cookieJar,
	}

	return &WebFetcher{Client: client, BaseUrl: url}
}

func (wb *WebFetcher) Get(url string) (*http.Response, error) {
	//"https://yosoymas.masmovil.es/validate/"
	 return wb.Client.Get(url)
}

func (wb *WebFetcher) Cookies() ([]*http.Cookie){
	//r.Request.URL
	return wb.Client.Jar.Cookies(wb.BaseUrl)
}


func (wb *WebFetcher) Post(url string, values url.Values) (*http.Response, error) {
	body := bytes.NewBufferString(values.Encode())
//"https://yosoymas.masmovil.es/validate/"
	rsp, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//rsp, err := http.NewRequest("POST", "https://httpbin.org/post", body)
	rsp.Header.Set("User-Agent", "Mozilla")
	rsp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//rsp.Header.Set("Cookie", "sid=141295f502e07f041d21801eff4b9384; visid_incap_967703=KFIHTnZTTECZ+q3CZ0Sf7fgeXlkAAAAAQUIPAAAAAAAdLZPA/2B9PFQqHXoL6hnS; incap_ses_504_967703=0A50WubNsgTYi+7YupH+BvgeXlkAAAAA4tmZh7UuOlx/uPF0v2Hz2w==")
	rsp.Header.Set("Accept", "*/*")
	return wb.Client.Do(rsp)
}

func cosa() {

	form := url.Values{}
	form.Add("action", "login")
	form.Add("url", "")
	form.Add("user", "alvaro_gg@hotmail.com")
	form.Add("password", "MBAR4B1")

	c := NewWebFetcher(&url.URL{Host:"yosoymas.masmovil.es", Scheme:"https"})

/*
	c.get("https://yosoymas.masmovil.es/validate/")
	time.Sleep(3 * time.Second)
	c.post("https://yosoymas.masmovil.es/validate/", form)
	//&url.URL{Host:"https://yosoymas.masmovil.es", Path:"/validate/"}
	//Save("cookies.bin", c.cookies(r1.Request.URL))
	c.SaveCookies("cookies.json")
	//fmt.Println(r1.Request.URL)
	//b, _ := json.Marshal(c.cookies(r1.Request.URL))
	//ioutil.WriteFile("cookie.json", b, os.FileMode(0777))

*/

	//a :=[]byte("[{\"Name\":\"sid\",\"Value\":\"54f400872f9ca7270089295d17ff8345\",\"Path\":\"\",\"Domain\":\"\",\"Expires\":\"0001-01-01T00:00:00Z\",\"RawExpires\":\"\",\"MaxAge\":0,\"Secure\":false,\"HttpOnly\":false,\"Raw\":\"\",\"Unparsed\":null},{\"Name\":\"visid_incap_967703\",\"Value\":\"t3tj71p/RACwbK4GnJLEt+EoXlkAAAAAQUIPAAAAAACtuvA6zQA7XlAmU9IlqO+9\",\"Path\":\"\",\"Domain\":\"\",\"Expires\":\"0001-01-01T00:00:00Z\",\"RawExpires\":\"\",\"MaxAge\":0,\"Secure\":false,\"HttpOnly\":false,\"Raw\":\"\",\"Unparsed\":null},{\"Name\":\"incap_ses_504_967703\",\"Value\":\"zUOiCbZAGg3wu+/YupH+BuEoXlkAAAAAIcOVr+X2/gQMaKJk9ZqEYw==\",\"Path\":\"\",\"Domain\":\"\",\"Expires\":\"0001-01-01T00:00:00Z\",\"RawExpires\":\"\",\"MaxAge\":0,\"Secure\":false,\"HttpOnly\":false,\"Raw\":\"\",\"Unparsed\":null}]")
	/*a, err := ioutil.ReadFile("cookie.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	var cookies  []*http.Cookie
	json.Unmarshal(a, &cookies)*/
	c.LoadCookies("cookies.json")

	//r, err := http.NewRequest("GET", "https://yosoymas.masmovil.es/validate/", nil)
	//fmt.Println(r.URL.Host)
	fmt.Println(c.Cookies())

	//res, err := c.get("https://yosoymas.masmovil.es")
	/*s, _ := json.Marshal(r2.URL)
	fmt.Println(string(s))
	return*/


	res, err := c.Get("https://yosoymas.masmovil.es/consumo/?line=677077536")
	//r, _ = client.Get("https://yosoymas.masmovil.es/")
	//data, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(data))

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".box-main-content").Find(".progress").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//c := s.Find("span").Text()
		re := regexp.MustCompile("([0-9]+)|(infinito)")
		r := re.FindAllString(s.Text(), -1)
		if i == 0 {

			fmt.Printf("Megas gastados %v de %v\n", r[0], r[1])
		} else {
			fmt.Printf("Minutos gastados %v de %v\n", r[0], r[1])
		}
	})

	//gCurCookies = cookieJar.Cookies(r.Request.URL)
	//fmt.Println(gCurCookies)
	fmt.Println("Estado ", res.StatusCode);
	return

	//r, _ = client.Get("https://yosoymas.masmovil.es/")
	//data, _ = ioutil.ReadAll(r.Body)

	//fmt.Println(string(data))

	/*requestDump, err := httputil.DumpRequest(r.Request, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println((requestDump))*/

	return
}
