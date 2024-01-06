package main

import (
	"errors"
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

const (
	ScreenshotsDir = "screenshots"
)

const (
	Single = iota
	Double
	Triple
	Custom
)

func githubUrlStringToFilename(str string) string {
	filetype := ".png"

	// Remove "https://" from the beginning
	str = strings.ReplaceAll(str, "://", "-")

	// Replace "/" with "-"
	str = strings.ReplaceAll(str, "/", "-")

	// Replace "?" with ""
	str = strings.ReplaceAll(str, "?", "")

	// Replace "&" with "-"
	str = strings.ReplaceAll(str, "&", "-")

	// Replace "=" with "-"
	str = strings.ReplaceAll(str, "=", "-")

	return str + filetype
}

func execGowitness(siteURL string) error {
	gowitness := "/go/bin/gowitness"
	// gowitness := "gowitness"

	cmd := exec.Command(gowitness, "single", "--fullpage", siteURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error executing command:", err)
		return err
	}

	log.Println("Executing gowitness: " + string(output))

	return nil
}

func getGithubUrl(username string) string {
	return "https://github.com/" + username + "?tab=overview&from=2024-01-01&to=2024-01-03"
}

type Screenshot struct {
	Username  string
	GithubUrl string
	Filename  string
	Directory string
}

func TakeScreenshot(username string, option int, userYOffset int) (string, error) {
	var inY int
	switch option {
	case Triple:
		inY = 1270
	case Double:
		inY = 970
	case Single:
		inY = 700
	case Custom:
		inY = userYOffset
	}

	u := getGithubUrl(username)
	f := githubUrlStringToFilename(u)

	s := Screenshot{
		Username:  username,
		GithubUrl: u,
		Filename:  f,
		Directory: ScreenshotsDir,
	}

	err := execGowitness(s.GithubUrl)
	if err != nil {
		log.Println(err)
		return "", errors.New("An internal error occurred attempting screenshot.")
	}

	// Open the resulting PNG image
	file, err := os.Open(filepath.Join(s.Directory, s.Filename))
	if err != nil {
		log.Println(err)
		return "", errors.New("Error opening file.")
	}
	defer file.Close()

	// Decode the PNG image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// maxHeight := img.Bounds().Dy()

	// Values for contribution graph
	// Create a new RGBA image with the calculated dimensions
	padding := 40
	contributionGraphWidth := 1800 + padding
	contributionGraphHeight := 425 + padding
	croppedImage := image.NewRGBA(image.Rect(0, 0, contributionGraphWidth, contributionGraphHeight))

	// Draw into the contribution graph image based on offset
	xOffset := 860 - padding/2
	yOffset := inY - padding/2
	draw.Draw(croppedImage, croppedImage.Bounds(), img, image.Point{xOffset, yOffset}, draw.Over)

	// Timestamp the filename
	timestamp := time.Now().Format("2006-01-02")

	// Save the cropped image to a new PNG file
	outFile := filepath.Join(s.Directory, username+"_"+timestamp+"_"+fmt.Sprintf("%dpx", yOffset)+".png")
	outputFile, err := os.Create(outFile)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, croppedImage)
	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println("Saved cropped file to: " + outFile)
	return outFile, nil
}
