package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/structs"

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

	if query == "" {
		result_s, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(result_s))
		return
	}

	result_m := structs.New(result)
	for _, path := range strings.Split(query, ",") {
		res, err := followPath(result_m.Fields(), strings.Split(path, "/"))
		if err {
			fmt.Println("Path " + path + " not found!")
		}

		if queries {
			fmt.Print(path + ": ")
		} else if print_keys {
			keys := strings.Split(path, "/")
			fmt.Print(keys[len(keys) - 1] + ": ")
		} 
		fmt.Println(res)
	}
}

func followPath(fields []*structs.Field, path []string) (string, bool) {
	defer func() {
        if err := recover(); err != nil {
            fmt.Println("Path isn't full, panic")
        }
    }()

	for _, field := range fields {
		if field.Name() == path[0] {
			if len(path) == 1 {
				return field.Value().(string), false
			}
			return followPath(field.Fields(), path[1:])
		}
	}
	return "Not found or Incorrect path", true
}