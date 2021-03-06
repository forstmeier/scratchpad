# Quehook

> Monitor open source activity with webhooks :loudspeaker:

[![Build Status](https://travis-ci.org/forstmeier/quehook.svg?branch=master)](https://travis-ci.org/forstmeier/quehook) [![Coverage Status](https://coveralls.io/repos/github/forstmeier/quehook/badge.svg?branch=master)](https://coveralls.io/github/forstmeier/quehook?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/forstmeier/quehook)](https://goreportcard.com/report/github.com/forstmeier/quehook)

<a href="https://www.buymeacoffee.com/forstmeier" target="_blank"><img src="https://bmc-cdn.nyc3.digitaloceanspaces.com/BMC-button-images/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## :beers: Introduction

**Quehook** provides subscribable webhooks to queries on **GitHub** data. New **[BigQuery](https://www.gharchive.org/#bigquery)** queries in the [standard dialect](https://cloud.google.com/bigquery/docs/reference/standard-sql/query-syntax) can be submitted by anyone via the application API endpoints and anyone can subscribe to webhook updates for any submitted queries. These queries are run against the public **[GH Archive](https://www.gharchive.org/)** datasets on an hourly basis every time the archive is updated. This allows for recurring questions regarding the open source community to be regularly answered on the newest available data.

## :octocat: Usage

For instructions on how to submit new queries or to subscribe to hourly updates, checkout the **Usage** section of the currently unpublished public website [here](https://forstmeier.github.io/quehook/). You can use [curl](https://curl.haxx.se/), [Postwoman](https://liyasthomas.github.io/postwoman/), or whatever other tool you want, to make the required requests to the appropriate endpoint. The included `cft.json` file can be used to create the necessary AWS resources for the application.

## :green_book: FAQ

Head over to the [FAQ page](https://forstmeier.github.io/quehook/faq) to get more information on the data being used and how it's being used. A changelog will be maintained on the [releases tab](https://github.com/forstmeier/quehook/releases) and planned features/fixes are listed in the [issues tab](https://github.com/forstmeier/quehook/issues).
