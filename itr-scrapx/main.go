package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	var imageLinks []string

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
				m := len(s) - 4
				actor := s[:m]
				fmt.Println("actor:", actor)

			}

		}

	})

	// Start the collector
	err := c.Visit("https://example.com/gallery") // Change to your target URL
	if err != nil {
		fmt.Println("Error visiting page:", err)
		return
	}

	fmt.Printf("Found %d images\n", len(imageLinks))

	// Download and zip the images
	archiveName := "images.zip"
	err = downloadAndZipImages(imageLinks, archiveName)
	if err != nil {
		fmt.Println("Error downloading images:", err)
	}
	fmt.Printf("Images downloaded and saved to %s\n", archiveName)
}

func downloadAndZipImages(links []string, archiveName string) error {
	// Create a buffer to write our archive to
	/*
		buf := new(bytes.Buffer)
		zipWriter := zip.NewWriter(buf)

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
				return fmt.Errorf("could not create zip entry for %s: %v", fileName, err)
			}

			_, err = io.Copy(zipFile, resp.Body)
			if err != nil {
				return fmt.Errorf("could not copy data for %s: %v", fileName, err)
			}
		}

		err := zipWriter.Close()
		if err != nil {
			return fmt.Errorf("could not close zip writer: %v", err)
		}

		// Write the buffer to a file
		err = os.WriteFile(archiveName, buf.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("could not write zip file to disk: %v", err)
		}
	*/
	return nil
}
