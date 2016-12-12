package config

// DebugMode contains boolean of whether application is being run in development
var DebugMode bool

// SiteDomain is the current domain that the router is hosted on
var SiteDomain = "ons.gov.uk"

// PatternLibraryAssetsPath is the URL to the CSS and JS assets from the pattern library
var PatternLibraryAssetsPath = "//cdn.ons.gov.uk/sixteens/6cc1837"

// DataDiscoveryAssetsPath is the URL to the bundled React JS app build by webpack
var DataDiscoveryAssetsPath = "https://cdn.ons.gov.uk/dp-dd-react-app/f921145/bundle.js"
