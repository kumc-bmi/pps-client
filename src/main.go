package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	TOKEN := os.Getenv("LOCKBOX")
	URL := os.Getenv("LOCKBOX_URL")

	// parse args for '-update'
	boolPtr := flag.Bool("update", false, "a bool")
	wordPtr := flag.String("pass", "foo", "a string")
	flag.Parse()

	// read item id from stdin, trim newline if needed
	reader := bufio.NewReader(os.Stdin)
	var ITEM string
	ITEM, _ = reader.ReadString('\n')
	ITEM = strings.TrimSuffix(ITEM, "\n")

	if TOKEN == "" {
		log.Println("missing LOCKBOX environment variable")
		log.Println("usage (update): pps-client -update -pass=foo <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		log.Println("usage (retrieve): pps-client <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		os.Exit(1)
	}

	if URL == "" {
		log.Println("missing LOCKBOX_URL environment variable")
		log.Println("usage (update): pps-client -update -pass=foo <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		log.Println("usage (retrieve): pps-client <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		os.Exit(1)
	}

	if ITEM == "" {
		log.Println("missing ITEM")
		log.Println("usage (update): pps-client -update -pass=foo <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		log.Println("usage (retrieve): pps-client <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		os.Exit(1)
	}

	bearer := "Bearer " + TOKEN
	URL = URL + "/api/v5/rest/Entries/"

	// if -update is defined, update the item; else, get the item
	if *boolPtr {

		URL = URL + ITEM

		// parse args for '-pass'
		//fmt.Println("pass:", *wordPtr)

		// prepare patch request
		var jsonStr = []byte(`{"Password": "` + *wordPtr + `"}`)
		req, err := http.NewRequest("PATCH", URL, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", bearer)

		// send request
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string([]byte(body)))

	} else {

		URL = URL + ITEM + "/password"

		// prepare get request
		req, err := http.NewRequest("GET", URL, nil)
		req.Header.Add("Authorization", bearer)

		// send request
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string([]byte(body)))
	}

}
