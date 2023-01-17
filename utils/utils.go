package utils

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	goUrlDomain      = "go.dev"
	installationPath = "/usr/local/"
	linuxExtFile     = "tar.gz"
	windowsExtFile   = "zip"
)

func SystemDist() string {
	return fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
}

func ExtFile() string {
	var ext string

	switch runtime.GOOS {
	case "linux":
		ext = linuxExtFile
	case "windows":
		ext = windowsExtFile
	default:
		ext = ""
	}

	return ext
}

func DownloadUrl(fileName string) error {
	url := fmt.Sprintf("https://%s/dl/%s", goUrlDomain, fileName)
	tmpPath := filepath.Join(os.TempDir(), fileName)

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

func MatchChecksum(fileName, checksum string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	fileChecksum := sha256.Sum256(file)
	if hex.EncodeToString(fileChecksum[:]) != checksum {
		return errors.New("checksums not match")
	}

	return nil
}

func Remove() error {
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

// Update tmp file, gzipreader names.
func Install() error {
	file, err := os.Open("/tmp/go1.19.4.linux-amd64.tar.gz")
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
			err := os.Mkdir(installationPath+header.Name, 0755)
			if err != nil {
				return errors.New("error at create go file")
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
