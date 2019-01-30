package backscatter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/threathive/backscatter-go/internal/config"
	"crypto/tls"
	"encoding/json"
)

type Query map[string]interface{}
type Results map[string]interface{}


type Response struct {
	Message string `json:"message,omitempty"`
	Success bool `json:"success"`
	Query `json:"query,omitempty"`
	Results `json:"results,omitempty"`
}

func ApiTests(){
	configuration, err := config.New()
	if err != nil {
		log.Panicln("configuration error", err)
	}

	api_key := configuration.Constants.BackScatter.ApiKey
	api_server := configuration.Constants.BackScatter.ApiServer

	timeout := time.Duration(30 * time.Second) // sets 5 second timeout
	req, err := http.NewRequest("GET", api_server+"hello", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false}, // sets config to ignore ssl issues
		},
		CheckRedirect: func(req2 *http.Request, via []*http.Request) error {
			req = req2
			return nil
		},
	}

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}

	var r = new(Response)
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)

	req, err = http.NewRequest("GET", api_server+"observations/ip", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q := req.URL.Query()
	q.Add("query", "8.8.8.8")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))

	//var r = new(Response)
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)


	req, err = http.NewRequest("GET", api_server+"observations/network", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("query", "74.96.0.0/16")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)

	req, err = http.NewRequest("GET", api_server+"observations/asn", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("query", "701")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)

	req, err = http.NewRequest("GET", api_server+"observations/port", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("query", "6666")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)


	req, err = http.NewRequest("GET", api_server+"observations/country", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("query", "CN")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))

	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)




	req, err = http.NewRequest("GET", api_server+"trends/popular/ip", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("scope", "1d")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)


	req, err = http.NewRequest("GET", api_server+"trends/popular/network", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("scope", "1d")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)

	req, err = http.NewRequest("GET", api_server+"trends/popular/asn", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("scope", "1d")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)


	req, err = http.NewRequest("GET", api_server+"trends/popular/port", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("scope", "1d")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)


	req, err = http.NewRequest("GET", api_server+"trends/popular/country", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("scope", "1d")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)



	req, err = http.NewRequest("GET", api_server+"enrichment/ip", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("query" , "74.96.192.82")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)





	req, err = http.NewRequest("GET", api_server+"enrichment/network", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("query" , "74.96.0.0/32")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)


	req, err = http.NewRequest("GET", api_server+"enrichment/asn", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", api_key)

	q = req.URL.Query()
	q.Add("query" , "701")
	req.URL.RawQuery = q.Encode()

	//fmt.Println(req.URL.String())

	resp, err = client.Do(req) //makes the request
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &r)
	if err != nil{
		fmt.Println("error parsing hello" , err)
	}
	fmt.Println(r)

}
