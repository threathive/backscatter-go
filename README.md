Golang Backscatter
===================

This is a golang library built to interface with the Backscatter service. It allows for integrating of the service into any golang based project.
This project also includes a cli client for easily quering the api.


Backscatter: https://backscatter.io/
Dev documentation: https://backscatter.io/developers

Quick Start
-----------
**Install the library**:

```
	go get github.com/threathive/backscatter-go
```

**Build client**

To build the client move to the cli directory and issue the following command.

```
	go build cli.go
```

This will build a cli binary for which you can utilize all the client methods below.


**Search observations**:
```
./cli observations --qtype ip 60.28.178.124 --scope 7d
./cli observations --qtype network 60.28.0.0/15 --scope 2d
./cli observations --qtype asn 701 --scope 1d
./cli observations --qtype port 22
./cli observations --qtype country ru
```

**Get Trends**:
```
./cli trends --qtype ip
./cli trends --qtype network
./cli trends --qtype asn
./cli trends --qtype port
./cli trends --qtype country
```

**Enrich values**:
```
./cli enrichments --qtype ip 74.96.192.82
./cli enrichments --qtype network 74.96.0.0/16
./cli enrichments --qtype asn 701
./cli enrichments --qtype port 52
```
Features
--------
* Run observation searches for ip, network, asn, country and ports
* Run enrichment searches for ip, network, asn, country and ports
* Get trend data for all data types

Changelog
---------

02-17-19
~~~~~~~
* Initial launch of the library
