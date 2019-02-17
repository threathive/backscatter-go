/*
	Backscatter.io api client.
*/
package backscatter

import (
	"time"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"golang.org/x/net/context"

)

// Default api url
const DefaultBaseUrl = "https://api.backscatter.io/v0/"

// Version of this library
const Version = "0.1"

// Generic interface we will use for various Enrichment types
type GenericResponse interface {

}

/*
{
  "message": "bonjour",
  "success": true
}
*/
type HelloResponse struct {
        Message string `json:"message"`
        Success bool   `json:"success"`
}

/*
⊖{
    "query": ⊖{
        "after_time": "2019-01-05",
        "focus": "60.28.178.124",
        "scope": "now-7d",
        "type": "ip_query"
    },
    "results": ⊖{
        "observations": ⊖[
           ⊕{...}
        ],
        "summary": ⊖{
            "autonomous_system_count": 1,
            "has_observations": true,
            "ip_address_count": 1,
            "network_count": 1,
            "observations": false,
            "observations_count": 1,
            "port_count": 1,
            "protocol_count": 1
        },
        "unique": ⊖{
            "autonomous_systems": ⊕[ ... ],
            "countries": ⊕[ ... ],
            "ip_addresses": ⊕[ ... ],
            "networks": ⊕[ ... ],
            "ports": ⊕[ ... ],
            "protocols": ⊕[ ... ]
        }
    },
    "success": true
}
*/
type Observations struct {
	Query struct {
		AfterTime string `json:"after_time"`
		Focus     string `json:"focus"`
		Scope     string `json:"scope"`
		Type      string `json:"type"`
	} `json:"query"`
	Results struct {
		Observations []struct {
			DstPort       string    `json:"dst_port"`
			FragmentFlags string    `json:"fragment_flags"`
			ID            string    `json:"id"`
			Length        string    `json:"length"`
			Observed      time.Time `json:"observed"`
			Precedence    string    `json:"precedence"`
			Protocol      string    `json:"protocol"`
			Res           string    `json:"res"`
			SrcIP         string    `json:"src_ip"`
			SrcPort       string    `json:"src_port"`
			TCPFlags      string    `json:"tcp_flags"`
			Tos           string    `json:"tos"`
			TTL           string    `json:"ttl"`
			Window        string    `json:"window"`
		} `json:"observations"`
		Summary struct {
			AutonomousSystemCount int  `json:"autonomous_system_count"`
			HasObservations       bool `json:"has_observations"`
			IPAddressCount        int  `json:"ip_address_count"`
			NetworkCount          int  `json:"network_count"`
			Observations          bool `json:"observations"`
			ObservationsCount     int  `json:"observations_count"`
			PortCount             int  `json:"port_count"`
			ProtocolCount         int  `json:"protocol_count"`
		} `json:"summary"`
		Unique struct {
			AutonomousSystems []string `json:"autonomous_systems"`
			Countries         []string `json:"countries"`
			IPAddresses       []string `json:"ip_addresses"`
			Networks          []string `json:"networks"`
			Ports             []string `json:"ports"`
			Protocols         []string `json:"protocols"`
		} `json:"unique"`
	} `json:"results"`
	Success bool `json:"success"`
}

/*
⊖{
    "query": ⊖{
        "after_time": "2019-01-05",
        "focus": "ip",
        "scope": "now-7d",
        "type": "top_x_query"
    },
    "results": ⊖{
        "122.114.63.133": 610,
        "138.197.12.178": 1184,
        "138.197.216.210": 707,
        "139.59.69.217": 1322,
        "158.69.53.42": 336,
        "159.65.246.2": 3126,
        "159.65.246.252": 3003,
        "174.138.63.166": 1599,
        "176.119.4.18": 317,
        "176.119.4.77": 412,
        "18.224.214.192": 1321,
        "185.156.177.144": 565,
        "185.246.128.25": 366,
        "185.247.62.106": 256,
        "193.106.30.154": 245,
        "193.164.132.183": 1040,
        "193.238.46.82": 449,
        "194.28.115.243": 479,
        "194.28.115.245": 474,
        "198.108.67.48": 757,
        "37.26.130.146": 305,
        "5.8.18.70": 324,
        "51.75.62.169": 14397,
        "64.71.165.10": 332,
        "87.251.81.86": 456
    },
    "success": true
}
*/
type Trends struct {
	Query struct {
		AfterTime string `json:"after_time"`
		Focus     string `json:"focus"`
		Scope     string `json:"scope"`
		Type      string `json:"type"`
	} `json:"query"`
	Results map[string]int `json:"results"`
	Success bool `json:"success"`
}

/*
⊖{
    "results": ⊖{
        "as_name": "UUNET - MCI Communications Services, Inc. d/b/a Verizon Business, US",
        "as_num": 701,
        "city": "Vienna",
        "country_iso": "US",
        "country_name": "United States",
        "ip": "74.96.192.82",
        "ip_hex": "0x4a60c052",
        "ip_version": 4,
        "latitude": 38.8977,
        "longitude": -77.288,
        "network": "74.96.0.0/16",
        "network_broadcast": "74.96.255.255",
        "network_hostmask": "0.0.255.255",
        "network_netmask": "255.255.0.0",
        "network_size": 65536,
        "postal_code": "22181",
        "region_iso": "VA",
        "region_name": "Virginia"
    },
    "success": true
}
*/
type IPEncrichment struct {
	Results struct {
		AsName           string  `json:"as_name"`
		AsNum            int     `json:"as_num"`
		City             string  `json:"city"`
		CountryIso       string  `json:"country_iso"`
		CountryName      string  `json:"country_name"`
		IP               string  `json:"ip"`
		IPHex            string  `json:"ip_hex"`
		IPVersion        int     `json:"ip_version"`
		Latitude         float64 `json:"latitude"`
		Longitude        float64 `json:"longitude"`
		Network          string  `json:"network"`
		NetworkBroadcast string  `json:"network_broadcast"`
		NetworkHostmask  string  `json:"network_hostmask"`
		NetworkNetmask   string  `json:"network_netmask"`
		NetworkSize      int     `json:"network_size"`
		PostalCode       string  `json:"postal_code"`
		RegionIso        string  `json:"region_iso"`
		RegionName       string  `json:"region_name"`
	} `json:"results"`
	Success bool `json:"success"`
}

/*
⊖{
    "results": ⊖{
        "cidr": "74.96.0.0/32",
        "network_addresses": ⊖[
            "74.96.0.0"
        ],
        "network_size": 1
    },
    "success": true
}
*/
type NetworkEnrichment struct {
	Results struct {
		Cidr             string   `json:"cidr"`
		NetworkAddresses []string `json:"network_addresses"`
		NetworkSize      int      `json:"network_size"`
	} `json:"results"`
	Success bool `json:"success"`
}

/*
⊖{
    "results": ⊖{
        "as_num": 701,
        "as_name": "UUNET - MCI Communications Services, Inc. d/b/a Verizon Business, US",
        "prefix_count": 1104,
        "prefixes": ⊖[
            "71.115.128.0/17",
            "96.240.160.0/19",
            "108.8.192.0/18",
            "141.152.224.0/19",
            "152.208.0.0/13",
            "104.99.144.0/20",
            "206.64.0.0/14",
            "141.156.22.0/24",
            "71.187.0.0/16",
            "205.134.4.0/24",
            "104.100.160.0/22",
            "96.249.224.0/19",
            "104.95.192.0/20"
        ]
    },
    "success": true
}
*/
type ASNEnchriment struct {
	Results struct {
		AsNum       int      `json:"as_num"`
		AsName      string   `json:"as_name"`
		PrefixCount int      `json:"prefix_count"`
		Prefixes    []string `json:"prefixes"`
	} `json:"results"`
	Success bool `json:"success"`
}

/*
⊖{
    "results": ⊖[
       ⊖{
            "description": "http protocol over TLS/SSL",
            "port": "443",
            "protocol": "tcp",
            "service": "https"
        },
       ⊖{
            "description": "http protocol over TLS/SSL",
            "port": "443",
            "protocol": "udp",
            "service": "https"
        },
       ⊖{
            "description": "HTTPS",
            "port": "443",
            "protocol": "sctp",
            "service": "https"
        }
    ],
    "success": true
}
*/
type PortEnrichment struct {
	Results []struct {
		Description string `json:"description"`
		Port        string `json:"port"`
		Protocol    string `json:"protocol"`
		Service     string `json:"service"`
	} `json:"results"`
	Success bool `json:"success"`
}

/*
	Client struct representing a client connection to the backscatter.io api.
	@ApiUrl The api url.
	@ApiKey The api key.
	@HTTPClient Http Client
	@BaseURL The base api url
	@QueryType The type of query we are making
	@Query The query value we are searching for.
	@Scope The time frame we are searching against.
*/
type Client struct {
	ApiUrl string
	ApiKey string
	HTTPClient http.Client
	BaseURL string
	QueryType string
	Query string
	Scope string
}

/*
	Helper wrapper around the http request.
*/
func withCancel(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	canCancel, ok := client.Transport.(canceler)
	if !ok {
		return client.Do(req)
	}

	type doResponse struct {
		resp *http.Response
		err  error
	}

	c := make(chan doResponse, 1)
	go func() {
		resp, err := client.Do(req)
		c <- doResponse{
			resp: resp,
			err:  err,
		}
	}()
	select {
	case <-ctx.Done():
		canCancel.CancelRequest(req)
		<-c // Wait for f to return.
		return nil, ctx.Err()
	case r := <-c:
		return r.resp, r.err
	}
}

/*
	Basic Authentication request test.
*/
func (c *Client) Hello(ctx context.Context) (HelloResponse, error) {
	var m HelloResponse
	if err:= c.doReqURL(ctx, c.urlBase("hello"), &m); err != nil {
		return m, err
	}
	return m, nil
}

/*
	Search for observations on an ip, asn, network, port, country.
*/
func (c *Client) SearchObservations(ctx context.Context) (Observations, error) {
	var m Observations
	if err:= c.doReqURL(ctx, c.urlBase("observations/" + c.QueryType), &m); err != nil {
		return m, err
	}
	return m, nil
}

/*
	Search for trends on an ip, network, asn, port, country.
*/
func (c *Client) SearchTrends(ctx context.Context) (Trends, error) {
	var m Trends
	if err:= c.doReqURL(ctx, c.urlBase("trends/popular/" + c.QueryType), &m); err != nil {
		return m, err
	}
	return m, nil
}

/*
	Search for enrichment data on an ip, network, asn, port.
*/
func (c *Client) SearchEnrichments(ctx context.Context) ( GenericResponse, error) {
	var m GenericResponse

	if c.QueryType == "ip" {
		m = new(IPEncrichment)

	}else if c.QueryType == "network" {
		m = new(NetworkEnrichment)

	}else if c.QueryType == "asn" {
		m = new(ASNEnchriment)

	}else if c.QueryType == "port" {
		m = new(PortEnrichment)
	}

	if err:= c.doReqURL(ctx, c.urlBase("enrichment/" + c.QueryType), &m); err != nil {
		return m, err
	}
	return m, nil

}

/*
	Builds the request and dumps the response data into a passed struct.
*/
func (c *Client) doReqURL(ctx context.Context, u string, jsonInto interface{}) error {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/48.0")
	req.Header.Set("X-API-KEY", c.ApiKey)

	if c.Query != ""{
		q := req.URL.Query()
		q.Add("query", c.Query)

		if c.Scope != ""{
			q.Add("scope", c.Scope)
		}

		req.URL.RawQuery = q.Encode()
	}

	resp, err := withCancel(ctx, &c.HTTPClient, req)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error...")
	}

	if err := json.NewDecoder(&b).Decode(jsonInto); err != nil {
		fmt.Println("can't decode json")
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil

}

/*
	Builds and updates the current base url.
*/
func (c *Client) urlBase(endpoint string) string {
	base := c.BaseURL
	if c.BaseURL == "" {
		base = DefaultBaseUrl
	}
	return fmt.Sprintf("%s%s", base, endpoint)
}
