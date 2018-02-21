package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var apiKey = func() string {
	apiKey := os.Getenv("API_KEY_BIBLESORG")
	if apiKey == "" {
		log.Fatalf("environment variable API_KEY_BIBLESORG must be set")
	}
	return apiKey
}()

func get(path string) (*http.Response, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "bibles.org",
		Path:   path,
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.SetBasicAuth(apiKey, "X")

	return client.Do(req)
}

func main() {
	m := map[string]string{}

	books, err := getBooks()
	if err != nil {
		log.Fatalf("getting books: %v", err)
	}

	for _, b := range books {
		m[b.Name] = b.Abbr
	}

	fmt.Printf("package biblestats\n\nvar abbr = %#v\n", m)
}

type book struct {
	ID   string
	Name string
	Abbr string
}

func getBooks() ([]book, error) {
	path := "/v2/versions/eng-CEV/books.js"

	res, err := get(path)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var booksRes struct {
		Response struct {
			Books []book
		}
	}

	err = json.Unmarshal(body, &booksRes)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, body)
	}

	for i, b := range booksRes.Response.Books {
		b.Name = strings.Title(b.Name)
		booksRes.Response.Books[i] = b
	}

	return booksRes.Response.Books, nil
}
