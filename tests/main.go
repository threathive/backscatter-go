/*
	Backscatter.io golang client tests
*/
package main

import (
	"github.com/urfave/cli"
	"os"
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

	var qtype,scope string
	app := cli.NewApp()
	app.Name = "Backscatter.io command line client."
	app.Usage = "Let's you lookup sensor data related to ip, domain, asn, country data."
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name: "qtype",
			Value: "ip",
			Usage: "query type",
			Destination: &qtype,
		},
		cli.StringFlag{
			Name: "scope",
			Value: "1d",
			Usage: "time scope",
			Destination: &scope,
		},
	}

	// we create our commands
	app.Commands = []cli.Command{

		{
			Name:  "ping",
			Usage: "Looks to see if we can make a basic authenticated request.",
			Flags: myFlags,
			// the action, or code that will be executed when
			// we execute our `ns` command
			Action: func(c *cli.Context) error {
				ctx := context.Background()
				h, _ := client.Hello(ctx)
				fmt.Println(prettyPrint(h))
				return nil
			},
		},
		{
			Name:  "observations",
			Usage: "Looks up observations for a ip, network , asn , port or country.",
			Flags: myFlags,
			// the action, or code that will be executed when
			// we execute our `ns` command
			Action: func(c *cli.Context) error {
				ctx := context.Background()
				client.Query = c.Args().Get(0)
				client.QueryType = qtype
				client.Scope = scope

				o, _ := client.SearchObservations(ctx)
				fmt.Println(prettyPrint(o))
				return nil

			},
		},
		{
			Name:  "trends",
			Usage: "Looks up trending data about an ip , network, asn , port or country",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ctx := context.Background()
				client.Query = ""
				client.QueryType = qtype
				client.Scope = scope

				h, _ := client.SearchTrends(ctx)
				fmt.Println(prettyPrint(h))
				return nil
			},
		},
		{
			Name:  "enrichments",
			Usage: "Looks up enrichment data about an ip , network, asn or port.",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ctx := context.Background()
				client.Query = c.Args().Get(0)
				client.QueryType = qtype
				client.Scope = scope
				h, _ := client.SearchEnrichments(ctx)
				fmt.Println(prettyPrint(h))
				return nil
			},

		},

	}

	// start our application
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}


}



