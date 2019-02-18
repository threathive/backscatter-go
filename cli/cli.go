/*
	Backscatter.io golang client tests
*/
package main

import (
	"strings"
	"github.com/urfave/cli"
	"os"
	"fmt"
	"net/http"
	"time"
	"log"
        "crypto/tls"

	"github.com/threathive/backscatter-go"
        "github.com/spf13/viper"
	"golang.org/x/net/context"
        "encoding/json"

)

var Timeout = time.Duration(30 * time.Second) // sets 30 second timeout for all requests.

/*
	client config settings.
*/
type config struct {
        BackScatter struct {
            ApiServer string
            ApiKey string
        }

}

/*
	Helper method to ppretty print a struct.
*/
func prettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

/*
	Attempts to load a config file by reading a file api.config.
*/
func loadConfig() (config, error){
        var Config config

        viper.SetConfigName("api.config") // Configuration fileName without the .TOML or .YAML extension
        viper.AddConfigPath(".")           // Search the root directory for the configuration file
        err := viper.ReadInConfig()        // Find and read the config file
        if err != nil {                    // Handle errors reading the config file
		return Config, err
        }

        err = viper.Unmarshal(&Config)
	if err != nil {
		return Config, err
	}

	return Config, nil
}


func main() {
	Config, err := loadConfig()
	if err != nil {
		log.Fatal("Can not load config")
	}

	client := backscatter.Client{
	        ApiKey: Config.BackScatter.ApiKey,
        	ApiUrl: Config.BackScatter.ApiServer,
	}

        client.HTTPClient = http.Client{
                Timeout: Timeout,
                Transport: &http.Transport{
                        TLSClientConfig: &tls.Config{InsecureSkipVerify: false}, // sets config to ignore self signed certs.
                },
        }

	var qtype,scope string //stores the query type and the time frame for searches.

	app := cli.NewApp()
	app.Version = "0.0.1"
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
			Action: func(c *cli.Context) error {
				ctx := context.Background()
				client.Query = strings.ToUpper(c.Args().Get(0))
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



