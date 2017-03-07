package config

// DebugMode contains boolean of whether application is being run in development
var DebugMode bool

// SiteDomain is the current domain that the router is hosted on
var SiteDomain = "ons.gov.uk"

// PatternLibraryAssetsPath is the URL to the CSS and JS assets from the pattern library
var PatternLibraryAssetsPath = "//cdn.ons.gov.uk/sixteens/caf397f"

// DataDiscovery has the configuration values for the dataset JS application
var DataDiscovery = struct {
	ReactAppAssetsPath string
	ApiURL             string
	JobsApiURL         string
	BasePath           string
}{
	"https://cdn.ons.gov.uk/dp-dd-react-app/bbf32ff",
	"/dd/api",
	"/dd/api/jobs",
	"/dd",
}

// ZebedeeURL is the URL for Zebedee
var ZebedeeURL = "http://localhost:8082"
