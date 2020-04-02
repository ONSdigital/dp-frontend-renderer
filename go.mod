module github.com/ONSdigital/dp-frontend-renderer

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/ONSdigital/dp-frontend-models v1.5.0
	github.com/ONSdigital/dp-healthcheck v1.0.2
	github.com/ONSdigital/go-ns v0.0.0-20200205115900-a11716f93bad
	github.com/ONSdigital/log.go v1.0.0
	github.com/c2h5oh/datasize v0.0.0-20170519143321-54516c931ae9
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/pat v0.0.0-20160413041632-cf955c3d1f2c
	github.com/gosimple/slug v1.1.1
	github.com/justinas/alice v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nicksnyder/go-i18n/v2 v2.0.3
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/unrolled/render v1.0.2
	golang.org/x/text v0.3.2
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.2
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
