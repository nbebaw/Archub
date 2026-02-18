package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// AUR RPC API endpoint for package info
const (
	aurInfoURL = "https://aur.archlinux.org/rpc/?v=5&type=info&arg[]=%s"
)

// AURPackageInfo represents detailed package information from AUR RPC info endpoint
type AURPackageInfo struct {
	Name        string   `json:"Name"`
	Version     string   `json:"Version"`
	Description string   `json:"Description"`
	URL         string   `json:"URL"`
	License     []string `json:"License"`
	Maintainer  string   `json:"Maintainer"`
	Submitter   string   `json:"Submitter"`
	NumVotes    int      `json:"NumVotes"`
	Popularity  float64  `json:"Popularity"`

	// Unix timestamps
	FirstSubmitted int64  `json:"FirstSubmitted"`
	LastModified   int64  `json:"LastModified"`
	OutOfDate      *int64 `json:"OutOfDate"` // Pointer because it can be null

	// Dependencies
	Depends      []string `json:"Depends"`
	MakeDepends  []string `json:"MakeDepends"`
	OptDepends   []string `json:"OptDepends"`
	CheckDepends []string `json:"CheckDepends"`

	// Conflicts and Provides
	Conflicts []string `json:"Conflicts"`
	Provides  []string `json:"Provides"`
	Replaces  []string `json:"Replaces"`

	// Additional metadata
	PackageBase string   `json:"PackageBase"`
	Keywords    []string `json:"Keywords"`
	URLPath     string   `json:"URLPath"`
}

// AURInfoResponse represents the JSON response from AUR RPC info endpoint
type AURInfoResponse struct {
	ResultCount int                `json:"resultcount"`
	Results     []AURPackageInfo   `json:"results"`
}

// FetchPackageInfo queries the AUR RPC API for detailed package information.
// Uses AUR RPC v5 API endpoint with type=info.
//
// Parameters:
//   - packageName: Exact package name to query
//
// Returns:
//   - *AURPackageInfo: Package information if found, nil if not found
//   - error: nil on success, error if HTTP request or JSON parsing fails
func FetchPackageInfo(packageName string) (*AURPackageInfo, error) {
	// Build the info URL with the package name
	url := fmt.Sprintf(aurInfoURL, packageName)

	// Make HTTP GET request to AUR RPC API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse JSON response
	var aurResp AURInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, err
	}

	// Check if package was found
	if aurResp.ResultCount == 0 {
		return nil, nil
	}

	// Return first result
	return &aurResp.Results[0], nil
}

// CheckIfInstalled checks if a package is currently installed using pacman.
//
// Parameters:
//   - packageName: Name of package to check
//
// Returns:
//   - bool: true if installed, false otherwise
//   - string: Installed version if found, empty string otherwise
func CheckIfInstalled(packageName string) (bool, string) {
	// Execute pacman -Q to check if package is installed
	cmd := exec.Command("pacman", "-Q", packageName)
	output, err := cmd.Output()

	// If command fails, package is not installed
	if err != nil {
		return false, ""
	}

	// Parse output: "packagename version"
	parts := strings.Fields(string(output))
	if len(parts) >= 2 {
		return true, parts[1]
	}

	return false, ""
}

// FormatTimestamp converts Unix timestamp to human-readable date.
//
// Parameters:
//   - timestamp: Unix timestamp in seconds
//
// Returns:
//   - string: Formatted date string (e.g., "Mon Jan 02 2006")
func FormatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("Mon Jan 02 2006")
}

// FormatStringArray formats a string array for display.
// Handles None case and proper spacing.
//
// Parameters:
//   - arr: Array of strings to format
//   - indent: String to use for indentation (for multi-line)
//
// Returns:
//   - string: Formatted string representation
func FormatStringArray(arr []string, indent string) string {
	if len(arr) == 0 {
		return "None"
	}

	// For single item, return it directly
	if len(arr) == 1 {
		return arr[0]
	}

	// For multiple items, join with spaces
	return strings.Join(arr, "  ")
}

// FormatOptDepends formats optional dependencies with proper indentation.
//
// Parameters:
//   - arr: Array of optional dependency strings
//
// Returns:
//   - string: Formatted string with newlines and indentation
func FormatOptDepends(arr []string) string {
	if len(arr) == 0 {
		return "None"
	}

	if len(arr) == 1 {
		return arr[0]
	}

	// Join with newlines and indentation for multiple items
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result += "\n                " + arr[i]
	}
	return result
}

// DisplayPackageInfo formats and displays package information.
// Output format matches pacman -Si for user familiarity.
//
// Parameters:
//   - pkg: Package information to display
func DisplayPackageInfo(pkg *AURPackageInfo) {
	// Repository
	fmt.Printf("%sRepository      :%s aur\n", ColorGreen, ColorNone)

	// Name
	fmt.Printf("%sName            :%s %s%s%s\n", ColorGreen, ColorNone, ColorLightRed, pkg.Name, ColorNone)

	// Version
	fmt.Printf("%sVersion         :%s %s\n", ColorGreen, ColorNone, pkg.Version)

	// Description
	fmt.Printf("%sDescription     :%s %s\n", ColorGreen, ColorNone, pkg.Description)

	// URL
	url := pkg.URL
	if url == "" {
		url = "None"
	}
	fmt.Printf("%sURL             :%s %s\n", ColorGreen, ColorNone, url)

	// Licenses
	licenses := FormatStringArray(pkg.License, "")
	fmt.Printf("%sLicenses        :%s %s\n", ColorGreen, ColorNone, licenses)

	// Groups (always None for AUR)
	fmt.Printf("%sGroups          :%s None\n", ColorGreen, ColorNone)

	// Provides
	provides := FormatStringArray(pkg.Provides, "")
	fmt.Printf("%sProvides        :%s %s\n", ColorGreen, ColorNone, provides)

	// Depends On
	depends := FormatStringArray(pkg.Depends, "")
	fmt.Printf("%sDepends On      :%s %s\n", ColorGreen, ColorNone, depends)

	// Make Deps
	makeDeps := FormatStringArray(pkg.MakeDepends, "")
	fmt.Printf("%sMake Deps       :%s %s\n", ColorGreen, ColorNone, makeDeps)

	// Check Deps
	checkDeps := FormatStringArray(pkg.CheckDepends, "")
	fmt.Printf("%sCheck Deps      :%s %s\n", ColorGreen, ColorNone, checkDeps)

	// Optional Deps (with special formatting)
	optDeps := FormatOptDepends(pkg.OptDepends)
	fmt.Printf("%sOptional Deps   :%s %s\n", ColorGreen, ColorNone, optDeps)

	// Conflicts With
	conflicts := FormatStringArray(pkg.Conflicts, "")
	fmt.Printf("%sConflicts With  :%s %s\n", ColorGreen, ColorNone, conflicts)

	// Replaces
	replaces := FormatStringArray(pkg.Replaces, "")
	fmt.Printf("%sReplaces        :%s %s\n", ColorGreen, ColorNone, replaces)

	// Maintainer
	maintainer := pkg.Maintainer
	if maintainer == "" {
		maintainer = "None"
	}
	fmt.Printf("%sMaintainer      :%s %s\n", ColorGreen, ColorNone, maintainer)

	// Votes
	fmt.Printf("%sVotes           :%s %d\n", ColorGreen, ColorNone, pkg.NumVotes)

	// Popularity
	fmt.Printf("%sPopularity      :%s %.2f\n", ColorGreen, ColorNone, pkg.Popularity)

	// First Submitted
	firstSubmitted := FormatTimestamp(pkg.FirstSubmitted)
	fmt.Printf("%sFirst Submitted :%s %s\n", ColorGreen, ColorNone, firstSubmitted)

	// Last Modified
	lastModified := FormatTimestamp(pkg.LastModified)
	fmt.Printf("%sLast Modified   :%s %s\n", ColorGreen, ColorNone, lastModified)

	// Out-of-Date
	if pkg.OutOfDate != nil {
		outOfDate := FormatTimestamp(*pkg.OutOfDate)
		fmt.Printf("%sOut-of-Date     :%s %s%s%s\n", ColorGreen, ColorNone, ColorRed, outOfDate, ColorNone)
	} else {
		fmt.Printf("%sOut-of-Date     :%s No\n", ColorGreen, ColorNone)
	}

	// Installed status
	installed, version := CheckIfInstalled(pkg.Name)
	if installed {
		fmt.Printf("%sInstalled       :%s %sYes (%s)%s\n", ColorGreen, ColorNone, ColorGreen, version, ColorNone)
	} else {
		fmt.Printf("%sInstalled       :%s %sNo%s\n", ColorGreen, ColorNone, ColorRed, ColorNone)
	}
}

// MainPkgInfo is the entry point for the package info command.
// Fetches and displays detailed AUR package information.
//
// Parameters:
//   - packageName: Name of package to query
func MainPkgInfo(packageName string) {
	// Fetch package information from AUR
	pkg, err := FetchPackageInfo(packageName)
	if err != nil {
		fmt.Printf("%sError fetching package information:%s %v\n", ColorRed, ColorNone, err)
		return
	}

	// Check if package was found
	if pkg == nil {
		fmt.Printf("%sPackage '%s' not found in AUR%s\n", ColorRed, packageName, ColorNone)
		return
	}

	// Display package information
	DisplayPackageInfo(pkg)
}
