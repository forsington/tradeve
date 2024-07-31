package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/forsington/tradeve/item"
	"github.com/gocarina/gocsv"
	"github.com/hashicorp/go-hclog"
)

const (
	SDEFilename = "invTypes.csv"
	SDEFileURL  = "https://www.fuzzwork.co.uk/dump/latest/invTypes.csv"
)

func ensureSDEFileExists(force bool, logger hclog.Logger) error {
	shouldDownload := false

	if _, err := os.Stat(SDEFilename); os.IsNotExist(err) {
		logger.Info("SDE file not found, downloading from fuzzwork", "filename", SDEFilename)
		shouldDownload = true
	} else if force {
		logger.Info("SDE file exists but force flag is set, downloading from fuzzwork", "filename", SDEFilename)

		err := os.Rename(SDEFilename, fmt.Sprintf("%s.bak-%s", SDEFilename, time.Now().Format("20060102-150405")))
		if err != nil {
			return err
		}

		shouldDownload = true
	}

	if shouldDownload {
		err := downloadCSV(SDEFileURL, SDEFilename)
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadCSV(url, path string) error {
	if url == "" {
		return fmt.Errorf("no URL provided")
	}

	if path == "" {
		return fmt.Errorf("no path provided")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	itemCsv, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer itemCsv.Body.Close()

	itemsFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer itemsFile.Close()

	_, err = io.Copy(itemsFile, itemCsv.Body)
	if err != nil {
		return err
	}

	return nil
}

func loadCSV(filename string) (item.Types, error) {
	typesFile, err := os.OpenFile(filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer typesFile.Close()

	types := item.Types{}

	if err := gocsv.UnmarshalFile(typesFile, &types); err != nil {
		return nil, err
	}

	if len(types) == 0 {
		return nil, fmt.Errorf("no types found in CSV file")
	}

	return types, nil
}
