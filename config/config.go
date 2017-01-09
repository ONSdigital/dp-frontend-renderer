package config

// DebugMode contains boolean of whether application is being run in development
var DebugMode bool

// SiteDomain is the current domain that the router is hosted on
var SiteDomain = "ons.gov.uk"

// PatternLibraryAssetsPath is the URL to the CSS and JS assets from the pattern library
var PatternLibraryAssetsPath = "//cdn.ons.gov.uk/sixteens/6cc1837"

// DataDiscovery has configuration values for the dataset JS application
var DataDiscovery = struct {
	AssetsPath  string
	API_URL     string
	JOB_API_URL string
	BASE_PATH   string
}{
	"https://cdn.ons.gov.uk/dp-dd-react-app/78210cf",
	"http://localhost:20099",
	"http://localhost:20010",
	"/dd",
}
