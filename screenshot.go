package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func stripSpecialChars(url string) string {
	// Remove "https://" from the beginning
	url = strings.ReplaceAll(url, "://", "-")

	// Replace "/" with "-"
	url = strings.ReplaceAll(url, "/", "-")

	// Replace "?" with ""
	url = strings.ReplaceAll(url, "?", "")

	// Replace "&" with "-"
	url = strings.ReplaceAll(url, "&", "-")

	// Replace "=" with "-"
	url = strings.ReplaceAll(url, "=", "-")

	return url
}

func takeScreenshot(username string) string {
	originalURL := "https://github.com/" + username + "?tab=overview&from=2024-01-01&to=2024-01-03"
	// fmt.Println("taking screenshot of: " + originalURL)

	screenshotsDir := "screenshots"

	filename := stripSpecialChars(originalURL)
	filetype := "png"

	cmd := exec.Command("gowitness", "single", originalURL)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	screenshotFilename := filepath.Join(screenshotsDir, fmt.Sprintf("%s.%s", filename, filetype))

	file, err := os.Open(screenshotFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the PNG image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	padding := 40

	contributionGraphWidth := 1800 + padding
	contributionGraphHeight := 425 + padding

	// Create a new RGBA image with the calculated dimensions
	croppedImage := image.NewRGBA(image.Rect(0, 0, contributionGraphWidth, contributionGraphHeight))

	xOffset := 860 - padding/2
	yOffset := 700 - padding/2
	draw.Draw(croppedImage, croppedImage.Bounds(), img, image.Point{xOffset, yOffset}, draw.Over)

	// Save the cropped image to a new PNG file
	// outfile := filepath r
	tempName := fmt.Sprintf("%s-%s-out.%s", time.Now().Format("2006-01-02"), filename, filetype)
	croppedFilename := filepath.Join(screenshotsDir, tempName)
	outputFile, err := os.Create(croppedFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, croppedImage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(croppedFilename)
	return croppedFilename
}
