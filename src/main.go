package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/alexflint/go-arg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// returns a pointer to an http.Request with the appropriate headers
func prepareRequest(url string, bearer string, method string, jsonStr []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("X-Pleasant-Client-Identifier", "00000000-0000-0000-0000-000000000000")
	req.Header.Set("Content-Type", "application/json")

	return req, err
}

// performs a request.  returns the response json string or an error
func performRequest(req *http.Request) (string, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if req.Method == "PATCH" && resp.StatusCode != 204 {
		return "", errors.New("endpoint returned code " + strconv.Itoa(resp.StatusCode))
	} else if req.Method == "GET" && resp.StatusCode != 200 {
		return "", errors.New("endpoint returned code " + strconv.Itoa(resp.StatusCode))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return (string([]byte(body))), err
}

// update an item password.  returns empty string or error
func updateItem(url string, bearer string, item string, pass string) (string, error) {
	url = url + item

	// prepare patch request
	var jsonStr = []byte(`{"Password": "` + pass + `"}`)
	req, err := prepareRequest(url, bearer, "PATCH", jsonStr)
	if err != nil {
		return "", err
	}

	resp, err := performRequest(req)
	if err != nil {
		return "", err
	}

	return resp, err
}

// fetch an item password or attachment.  does not manipulate response from api.  returns json string or error
func fetchItem(url string, bearer string, item string, attachment string) (string, error) {
	if attachment == "" {
		url = url + item + "/password"
	} else {
		url = url + item + "/Attachments/" + attachment
	}

	// prepare get request
	req, err := prepareRequest(url, bearer, "GET", nil)
	if err != nil {
		return "", err
	}

	resp, err := performRequest(req)
	if err != nil {
		return "", err
	}

	return resp, err
}

func main() {

	token := os.Getenv("LOCKBOX")
	url := os.Getenv("LOCKBOX_URL")

	var args struct {
		Item       string
		Attachment string
		Update     string
	}

	parser := arg.MustParse(&args)

	bearer := "Bearer " + token
	url = url + "/api/v5/rest/Entries/"

	switch {
	case args.Attachment != "":
		res, err := fetchItem(url, bearer, args.Item, args.Attachment)
		if err != nil {
			log.Println("error fetching attachment\n", err)
			os.Exit(1)
		}
		fmt.Println(res)
	case args.Update != "":
		res, err := updateItem(url, bearer, args.Item, args.Update)
		if err != nil {
			log.Println("error updating item\n", err)
			os.Exit(1)
		}
		fmt.Println(res)
	case args.Item != "":
		res, err := fetchItem(url, bearer, args.Item, "")
		if err != nil {
			log.Println("error fetching item\n", err)
			os.Exit(1)
		}
		fmt.Println(res)
	default:
		parser.Fail("incorrect usage")
	}

}
