package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
)

type MetaStruct struct {
	Channel    string
	URL        string
	Models     []string
	Categories []string
	Tags       []string
	ImageCount int
}

// var url string
var metadata MetaStruct

func getImages(url string) (MetaStruct, bytes.Buffer) {

	//url = "https://example.com/gallery"
	//url = "https://www.pornpics.com/galleries/pretty-amateur-babe-soledad-lomas-flashes-her-hot-boobs-at-the-store-20119808/"

	//url = "https://www.pornpics.com/galleries/ebony-babe-desiree-jacobsen-loses-her-lingerie-plays-with-her-hard-nipples-38146132/"

	// Create a new collector
	c := colly.NewCollector()
	var fullurl string
	// Check if the URL starts with the specified prefix
	if strings.HasPrefix(url, "https://www.pornpics.com/galleries/") {
		fullurl = url
	} else {
		fullurl = "https://www.pornpics.com/galleries/" + url
	}

	var imageLinks []string
	//metadata := new(MetaStruct)
	metadata.URL = fullurl

	// Find and visit all image links
	c.OnHTML("#tiles .rel-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "http") {
			link = e.Request.AbsoluteURL(link)
		}
		imageLinks = append(imageLinks, link)
	})

	c.OnHTML(".gallery-info__item", func(e *colly.HTMLElement) {

		if e.ChildText(".gallery-info__title") == "Models:" {
			actors := e.ChildAttrs("a", "title")

			for _, s := range actors {
				metadata.Models = append(metadata.Models, extractText(s))
			}
		} else if e.ChildText(".gallery-info__title") == "Channel:" {
			channels := e.ChildAttrs("a", "title")
			for _, s := range channels {
				metadata.Channel = extractText(s)
			}
		} else if e.ChildText(".gallery-info__title") == "Categories:" {
			categories := e.ChildAttrs("a", "title")
			for _, s := range categories {
				metadata.Categories = append(metadata.Categories, extractText(s))
			}
		} else if e.ChildText(".gallery-info__title") == "Tags List:" {
			tags := e.ChildAttrs("a", "title")
			for _, s := range tags {
				metadata.Tags = append(metadata.Tags, extractText(s))

			}
		}

	})

	// Start the collector
	err := c.Visit(fullurl) // Change to your target URL
	if err != nil {
		fmt.Println("Error visiting page:", err)
	}

	metadata.ImageCount = len(imageLinks)

	// Download and zip the images
	archiveName := "images.zip"
	data, err := downloadAndZipImages(imageLinks, archiveName)
	if err != nil {
		fmt.Println("Error downloading images:", err)
	}
	/*
		data, err := json.Marshal(metadata)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))*/

	return metadata, data
}

func extractText(s string) string {
	m := len(s) - 4
	return strings.TrimSpace(s[:m])
}

func downloadAndZipImages(links []string, archiveName string) (bytes.Buffer, error) {
	// Create a buffer to write our archive to

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	//Write metadata file to zip

	zipFile, err := zipWriter.Create("metadata.json")
	if err != nil {
		return *buf, fmt.Errorf("could not create zip entry for metadata.json: %v", err)
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(zipFile, bytes.NewReader(data))
	if err != nil {
		return *buf, fmt.Errorf("could not copy data for metadata.json: %v", err)
	}

	for _, link := range links {
		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("Failed to download %s: %v\n", link, err)
			continue
		}
		defer resp.Body.Close()

		fileName := filepath.Base(link)
		zipFile, err := zipWriter.Create(fileName)
		if err != nil {
			return *buf, fmt.Errorf("could not create zip entry for %s: %v", fileName, err)
		}

		_, err = io.Copy(zipFile, resp.Body)
		if err != nil {
			return *buf, fmt.Errorf("could not copy data for %s: %v", fileName, err)
		}
	}

	err = zipWriter.Close()
	if err != nil {
		return *buf, fmt.Errorf("could not close zip writer: %v", err)
	}

	// Write the buffer to a file
	err = os.WriteFile(archiveName, buf.Bytes(), 0644)
	if err != nil {
		return *buf, fmt.Errorf("could not write zip file to disk: %v", err)
	}

	return *buf, nil
}
