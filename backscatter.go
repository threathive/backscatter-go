package backscatter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"crypto/tls"
	"encoding/json"

	"github.com/threathive/backscatter-go/internal/config"
)

type query map[string]interface{}
type results map[string]interface{}

type response struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
	query   `json:"query,omitempty"`
	results `json:"results,omitempty"`
}

/*
	Basic function to test all the various api endpoints.
*/
func APITests() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("configuration error", err)
	}

	apikey := configuration.Constants.BackScatter.ApiKey
	apiserver := configuration.Constants.BackScatter.ApiServer

	timeout := time.Duration(30 * time.Second) // sets 5 second timeout

	OnlineCheck(apikey, apiserver)
	ObservationsIP(apikey, apiserver)
	ObservationsASN(apikey, apiserver)
	ObservationsPort(apikey, apiserver)
	ObservationsCountry(apikey, apiserver)
	ObservationsNetwork(apikey, apiserver)
}

/*
	OnlineCheck(apikey string, apiurl string)
	Checks if the api is accessible.
*/
func OnlineCheck(apikey string, apiserver string) {
	timeout := time.Duration(30 * time.Second) // sets 5 second timeout
	req, err := http.NewRequest("GET", apiserver+"hello", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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

	var r = new(response)
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)

}

/*
 */
func ObservationsIP(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"observations/ip", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func ObservationsNetwork(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"observations/network", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func ObservationsASN(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"observations/asn", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func ObservationsPort(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"observations/port", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func ObservationsCountry(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"observations/country", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func TrendsIP(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"trends/popular/ip", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func TrendsNetwork(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"trends/popular/network", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func TrendsASN(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"trends/popular/asn", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func TrendsPort(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"trends/popular/port", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func TrendsCountry(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"trends/popular/country", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func EnrichmentIP(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"enrichment/ip", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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

	q = req.URL.Query()
	q.Add("query", "74.96.192.82")
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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func EnrichmentNetwork(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"enrichment/network", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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

	q = req.URL.Query()
	q.Add("query", "74.96.0.0/32")
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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)
}

/*
 */
func EnrichmentASN(apikey string, apiserver string) {
	req, err := http.NewRequest("GET", apiserver+"enrichment/asn", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0") // sets the useragent
	req.Header.Set("X-API-KEY", apikey)

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
	if err != nil {
		fmt.Println("error parsing hello", err)
	}
	fmt.Println(r)

}
