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
	"time"
)

type issuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*issue
}

type issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *user
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type user struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func searchIssues(terms []string) (*issuesSearchResult, error) {
	const issuesURL = "https://api.github.com/search/issues"
	var result issuesSearchResult

	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// a) streaming decoder
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	return nil, err
	// }

	// b) read whole responce. not sure how it works for large payload
	b, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
