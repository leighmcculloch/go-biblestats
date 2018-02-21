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
	m := map[string][]int{}

	books, err := getBookIDs()
	if err != nil {
		log.Fatalf("getting books: %v", err)
	}

	for _, b := range books {
		m[b.Name] = []int{}

		chapters, err := getChapterIDs(b.ID)
		if err != nil {
			log.Fatalf("getting chapters: %v", err)
		}

		for ci, c := range chapters {
			m[b.Name] = append(m[b.Name], 0)

			verses, err := getVerseIDs(c)
			if err != nil {
				log.Fatalf("getting verses: %v", err)
			}

			m[b.Name][ci] = len(verses)
		}
	}

	fmt.Printf("package biblestats\n\nvar verses = %#v\n", m)
}

type book struct {
	ID   string
	Name string
}

func getBookIDs() ([]book, error) {
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

func getChapterIDs(bookID string) ([]string, error) {
	path := fmt.Sprintf("/v2/books/%s/chapters.js", bookID)

	res, err := get(path)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var chaptersRes struct {
		Response struct {
			Chapters []struct {
				ID      string
				Chapter string
			}
		}
	}

	err = json.Unmarshal(body, &chaptersRes)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, body)
	}

	chapters := []string{}
	for _, c := range chaptersRes.Response.Chapters {
		if c.Chapter == "int" {
			continue
		}
		chapters = append(chapters, c.ID)
	}

	return chapters, nil
}

func getVerseIDs(chapterID string) ([]string, error) {
	log.Println(chapterID)
	path := fmt.Sprintf("/v2/chapters/%s/verses.js", chapterID)

	res, err := get(path)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var versesRes struct {
		Response struct {
			Verses []struct {
				ID string
			}
		}
	}

	err = json.Unmarshal(body, &versesRes)
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, body)
	}

	verses := make([]string, len(versesRes.Response.Verses))
	for i, v := range versesRes.Response.Verses {
		verses[i] = v.ID
	}

	return verses, nil
}
