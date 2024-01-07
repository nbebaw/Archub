package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	aurURL = "https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s"
)

type AURPackage struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Version     string `json:"Version"`
}

type AURResponse struct {
	Results []AURPackage `json:"results"`
}

func SearchAURPackage(packageName string) ([]AURPackage, error) {
	url := fmt.Sprintf(aurURL, packageName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aurResp AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, err
	}

	return aurResp.Results, nil
}

func MainSearch(pkg_search string) {

	counter := 1
	packageName := pkg_search
	results, err := SearchAURPackage(packageName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Search results for %s:\n", packageName)
	for _, pkg := range results {
		fmt.Printf("%v) %s - v%s\nDescription: %s\n\n", counter, pkg.Name, pkg.Version, pkg.Description)
		counter++
	}
}
