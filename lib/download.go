package lib

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Download(url string, saveDir string, fileName string, unzip bool) error {
	log.Printf("[Download] URL: %s", url)
	_, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(saveDir, 0755)
	} else {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	downloadFilePath := filepath.Join(saveDir, fileName)
	out, err := os.Create(downloadFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	if unzip {
		log.Printf("[Decompression] Start - File: %s", downloadFilePath)
		r, err := zip.OpenReader(downloadFilePath)
		if err != nil {
			return err
		}
		defer r.Close()

		DirectoryName := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
		createDirectory := filepath.Join(saveDir, DirectoryName)
		log.Println()

		err = os.MkdirAll(createDirectory, 0755)
		if err != nil {
			return err
		}

		for _, f := range r.File {
			log.Printf("[Decompression] Processing - File: %s", f.Name)
			if f.Mode().IsDir() {
				// ディレクトリは無視して構わない
				continue
			}
			if err := saveUnZipFile(createDirectory, *f); err != nil {
				return err
			}
		}

		log.Printf("[Decompression] Delete - File: %s", downloadFilePath)
		err = os.Remove(downloadFilePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveUnZipFile(destDir string, f zip.File) error {
	destPath := filepath.Join(destDir, f.Name)
	if err := os.MkdirAll(filepath.Dir(destPath), f.Mode()); err != nil {
		return err
	}
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()
	if _, err := io.Copy(destFile, rc); err != nil {
		return err
	}

	return nil
}
