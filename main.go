package main

import (
	"flag"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func main() {
	var domain, query string;
	var print_keys, queries bool;
	flag.StringVar(&domain, "domain", "", "Domain to perform whois lookup")
	flag.StringVar(&query, "query", "", "Comma-separated list of pathes to print. You can run script without this key to see the structure of response. Write path as 'Key/SubKey'. Note Pascal case")
	flag.BoolVar(&print_keys, "print-keys", true, "Whether to print keys before requested values")
	flag.BoolVar(&queries, "queries", false, "Whether to print queries before requested values, has priority over -print-keys")

	flag.Parse()

	if *flag.Bool("help", false, "Show defaults") {
		flag.PrintDefaults()
		return
	}

	if domain == "" {
		fmt.Println("No domain were given.")
		return
	}

	result_raw, err := whois.Whois(domain)
	if err != nil {
		return
	}
	result, err := whoisparser.Parse(result_raw)
	if err != nil {
		return
	}
}

