/*
	Backscatter.io golang client tests
*/
package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
        "crypto/tls"

	"github.com/threathive/backscatter-go"
	"github.com/threathive/backscatter-go/internal/config"
	"golang.org/x/net/context"
        "encoding/json"

)

var timeout time.Duration

/*
	Helper method to ppretty print a struct.
*/
func prettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}


func main() {

        configuration, err := config.New()
        if err != nil {
                log.Panicln("configuration error", err)
        }

        timeout = time.Duration(30 * time.Second) // sets 5 second timeout

        apikey := configuration.Constants.BackScatter.ApiKey
        apiserver := configuration.Constants.BackScatter.ApiServer

	client := backscatter.Client{
	        ApiUrl: apiserver,
        	ApiKey: apikey,
	}

        client.HTTPClient = http.Client{
                Timeout: timeout,
                Transport: &http.Transport{
                        TLSClientConfig: &tls.Config{InsecureSkipVerify: false}, // sets config to ignore $
                },
        }

	ctx := context.Background()
	h, _ := client.Hello(ctx)
	fmt.Println(h)

	client.Query = "60.28.178.124"
	client.QueryType = "ip"

	o, _ := client.SearchObservations(ctx)
	fmt.Println(prettyPrint(o))

	client.Query = "74.96.0.0/16"
	client.QueryType = "network"

	o, _ = client.SearchObservations(ctx)
	fmt.Println(prettyPrint(o))

	client.Query = "701"
	client.QueryType = "asn"

	o, _ = client.SearchObservations(ctx)
	fmt.Println(prettyPrint(o))

	client.Query = "6666"
	client.QueryType = "port"

	o, _ = client.SearchObservations(ctx)
	fmt.Println(prettyPrint(o))

	client.Query = "US"
	client.QueryType = "country"

	o, _ = client.SearchObservations(ctx)
	fmt.Println(prettyPrint(o))

	client.Query = ""
	client.QueryType = "ip"

	t, _ := client.SearchTrends(ctx)
	fmt.Println(prettyPrint(t))


	client.QueryType = "network"

	t, _ = client.SearchTrends(ctx)
	fmt.Println(prettyPrint(t))

	client.QueryType = "asn"

	t, _ = client.SearchTrends(ctx)
	fmt.Println(prettyPrint(t))

	client.QueryType = "port"

	t, _ = client.SearchTrends(ctx)
	fmt.Println(prettyPrint(t))

	client.QueryType = "country"

	t, _ = client.SearchTrends(ctx)
	fmt.Println(prettyPrint(t))

	client.Query = "74.96.192.82"
	client.QueryType = "ip"

	e, _ := client.SearchEnrichments(ctx)
	fmt.Println(prettyPrint(e))

	client.Query = "74.96.0.0/32"
	client.QueryType = "network"

	e, _ = client.SearchEnrichments(ctx)
	fmt.Println(prettyPrint(e))

	client.Query = "701"
	client.QueryType = "asn"

	e, _ = client.SearchEnrichments(ctx)
	fmt.Println(prettyPrint(e))

	client.Query = "443"
	client.QueryType = "port"

	e, _ = client.SearchEnrichments(ctx)
	fmt.Println(prettyPrint(e))

}



