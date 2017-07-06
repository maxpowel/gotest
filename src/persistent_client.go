package main

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
	"encoding/gob"
	"runtime"
)

type persistentClient struct {
	client *http.Client

}

// Encode via Gob to file
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}


func NewClient() persistentClient {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	cookieJar, _ := cookiejar.New(&options)

	client := &http.Client{
		Jar: cookieJar,
	}

	return persistentClient{client: client}
}

func (p *persistentClient) get(url string) (*http.Response, error) {
	//"https://yosoymas.masmovil.es/validate/"
	 return p.client.Get(url)
}

func (p *persistentClient) cookies(url *url.URL) ([]*http.Cookie){
	//r.Request.URL
	return p.client.Jar.Cookies(url)
}


func (p *persistentClient) post(url string, values url.Values) (*http.Response, error) {
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
	return p.client.Do(rsp)
}

func cosa() {

	form := url.Values{}
	form.Add("action", "login")
	form.Add("url", "")
	form.Add("user", "alvaro_gg@hotmail.com")
	form.Add("password", "MBAR4B1")

	c := NewClient()


	/*r1, err := c.get("https://yosoymas.masmovil.es/validate/")
	time.Sleep(3 * time.Second)
	c.post("https://yosoymas.masmovil.es/validate/", form)
	//&url.URL{Host:"https://yosoymas.masmovil.es", Path:"/validate/"}
	Save("cookies.bin", c.cookies(r1.Request.URL))
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
	var cookies  []*http.Cookie
	Load("cookies.bin", &cookies)
	//r, err := http.NewRequest("GET", "https://yosoymas.masmovil.es/validate/", nil)
	//fmt.Println(r.URL.Host)
	fmt.Println(cookies)



	r2, err := http.NewRequest("GET", "https://yosoymas.masmovil.es", nil)
	r2.Header.Set("Referer", "https://yosoymas.masmovil.es/validate/")
	r2.Header.Set("Accept", "*/*")
	r2.Header.Set("User-Agent", "Mozilla")
	//res, err := c.get("https://yosoymas.masmovil.es")
	/*s, _ := json.Marshal(r2.URL)
	fmt.Println(string(s))
	return*/
	c.client.Jar.SetCookies(&url.URL{Host:"yosoymas.masmovil.es", Scheme:"https"}, cookies)
	res, err := c.client.Do(r2)
	if err != nil {
		panic(err)
	}

	//r, _ = client.Get("https://yosoymas.masmovil.es/")
	//data, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(data))

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".submenu-select").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		c := s.Find("option").Text()
		fmt.Printf("Contarto es %s", c)
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
