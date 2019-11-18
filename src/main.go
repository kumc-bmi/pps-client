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

	token := os.Getenv("LOCKBOX")
	url := os.Getenv("LOCKBOX_URL")

	// parse args for '-update'
	boolPtr := flag.Bool("update", false, "a bool")
	wordPtr := flag.String("pass", "foo", "a string")
	flag.Parse()

	// read item id from stdin, trim newline if needed
	reader := bufio.NewReader(os.Stdin)
	var item string
	item, _ = reader.ReadString('\n')
	item = strings.TrimSuffix(item, "\n")

	if token == "" {
		log.Println("missing LOCKBOX environment variable")
		log.Println("usage (update): pps-client -update -pass=foo <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		log.Println("usage (retrieve): pps-client <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		os.Exit(1)
	}

	if url == "" {
		log.Println("missing LOCKBOX_URL environment variable")
		log.Println("usage (update): pps-client -update -pass=foo <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		log.Println("usage (retrieve): pps-client <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		os.Exit(1)
	}

	if item == "" {
		log.Println("missing item")
		log.Println("usage (update): pps-client -update -pass=foo <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		log.Println("usage (retrieve): pps-client <<< f03d2746-c28e-11e9-aba2-df487c779baf")
		os.Exit(1)
	}

	bearer := "Bearer " + token
	url = url + "/api/v5/rest/Entries/"

	// if -update is defined, update the item; else, get the item
	if *boolPtr {

		url = url + item

		// parse args for '-pass'
		//fmt.Println("pass:", *wordPtr)

		// prepare patch request
		var jsonStr = []byte(`{"Password": "` + *wordPtr + `"}`)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Println("Error initializing request\n", err)
			os.Exit(1)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", bearer)

		// send request
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error on response.\n", err)
			os.Exit(1)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string([]byte(body)))

	} else {

		url = url + item + "/password"

		// prepare get request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error initializing request\n", err)
			os.Exit(1)
		}

		req.Header.Add("Authorization", bearer)
		req.Header.Add("X-Pleasant-Client-Identifier", "00000000-0000-0000-0000-000000000000")

		// send request
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Println("Error on response.\n", err)
			os.Exit(1)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string([]byte(body)))
	}

}
