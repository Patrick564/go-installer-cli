package cli

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	downloadUrl      = "https://go.dev/dl"
	installationPath = "/usr/local/"
	linuxExtFile     = "tar.gz"
)

// Remove the previous installation using GOROOT path.
func remove() error {
	goRoot, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		return err
	}

	err = os.RemoveAll(string(goRoot))
	if err != nil {
		return err
	}

	return nil
}

// Create the download name using the given version, system os and arch.
func downloadFilename(version string) string {
	dist := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)

	var ext string
	switch runtime.GOOS {
	case "linux":
		ext = linuxExtFile
	default:
		ext = ""
	}

	return fmt.Sprintf("%s.%s.%s", version, dist, ext)
}

// Download the selected Go version in tmp path.
func download(filename string) error {
	url := fmt.Sprintf("%s/%s", downloadUrl, filename)
	tmpPath := filepath.Join(os.TempDir(), filename)

	file, err := os.Create(tmpPath)
	if err != nil {
		return err
	}
	defer file.Close()

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}

// Install the given file in the system path show in go.dev installation instructions.
func install(filename string) error {
	file, err := os.Open(filepath.Join(os.TempDir(), filename))
	if err != nil {
		return err
	}

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			err := os.Mkdir(installationPath+header.Name, 0777)
			if err != nil {
				return errors.New("error at create folder")
			}
		case tar.TypeReg:
			file, err := os.Create(installationPath + header.Name)
			if err != nil {
				return err
			}

			_, err = io.Copy(file, tarReader)
			if err != nil {
				return errors.New("error at copy")
			}
			file.Close()
		default:
			return errors.New("unknow file type")
		}
	}

	return nil
}

// Remove the previous version, download and install the given version.
func Install(version string) {
	err := remove()
	if err != nil {
		log.Fatalf("Error at remove previous installation: %s", err.Error())
	}

	filename := downloadFilename(version)
	err = download(filename)
	if err != nil {
		log.Fatalf("Error at download Go version in tmp path: %s", err.Error())
	}

	err = install(filename)
	if err != nil {
		log.Fatalf("Error at install the Go version: %s", err.Error())
	}

	fmt.Println("Installed correctly.")
}
