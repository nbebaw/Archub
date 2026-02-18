package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// AUR RPC API endpoint for searching packages
const (
	aurURL = "https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s"
)

// AURPackage represents a single package from the AUR search results
type AURPackage struct {
	Name        string `json:"Name"`        // Package name
	Description string `json:"Description"` // Package description
	Version     string `json:"Version"`     // Latest package version
}

// AURResponse represents the JSON response from the AUR RPC API
type AURResponse struct {
	Results []AURPackage `json:"results"` // Array of matching packages
}

// SearchAURPackage queries the AUR RPC API to search for packages.
// Uses the AUR RPC v5 API endpoint with type=search.
//
// Parameters:
//   - packageName: Search query string (partial or full package name)
//
// Returns:
//   - []AURPackage: Slice of matching packages
//   - error: nil on success, error if HTTP request or JSON parsing fails
func SearchAURPackage(packageName string) ([]AURPackage, error) {
	// Build the search URL with the package name
	url := fmt.Sprintf(aurURL, packageName)

	// Make HTTP GET request to AUR RPC API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse JSON response
	var aurResp AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, err
	}

	return aurResp.Results, nil
}

// MainSearch is the entry point for the search command.
// Searches for packages in AUR and displays formatted results.
//
// Parameters:
//   - pkg_search: Package name or search query
func MainSearch(pkg_search string) {
	counter := 1
	packageName := pkg_search

	// Search for packages in AUR
	results, err := SearchAURPackage(packageName)
	if err != nil {
		log.Fatal(err)
	}

	// Display search results
	fmt.Printf("Search results for %s:\n", packageName)
	for _, pkg := range results {
		// Format: 1) package-name - v1.2.3
		//         Description: Package description
		fmt.Printf("%s%v)%s %s%s%s - v%s\n%sDescription:%s %s\n\n",
			ColorYellow, counter, ColorNone,
			ColorLightRed, pkg.Name, ColorNone, pkg.Version,
			ColorGreen, ColorNone, pkg.Description)
		counter++
	}
}
