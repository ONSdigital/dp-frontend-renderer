module github.com/ONSdigital/dp-frontend-renderer

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/ONSdigital/dp-frontend-models v1.10.2
	github.com/ONSdigital/dp-healthcheck v1.0.5
	github.com/ONSdigital/go-ns v0.0.0-20200902154605-290c8b5ba5eb
	github.com/ONSdigital/log.go v1.0.1
	github.com/c2h5oh/datasize v0.0.0-20200112174442-28bbd4740fee
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/pat v1.0.1
	github.com/gosimple/slug v1.9.0
	github.com/jteeuwen/go-bindata v3.0.7+incompatible // indirect
	github.com/justinas/alice v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nicksnyder/go-i18n/v2 v2.0.3
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/unrolled/render v1.0.2
	golang.org/x/text v0.3.2
	gopkg.in/russross/blackfriday.v2 v2.0.1
	gopkg.in/yaml.v2 v2.2.8
)

// This dependency fails if used directly. We need to do this replace. More info: https://github.com/russross/blackfriday/issues/500
replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
