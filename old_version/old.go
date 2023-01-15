// This file is an old version what only download an specific version and checksum.
package old

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const (
	URL      = "https://go.dev/dl/go1.19.4.linux-amd64.tar.gz"
	CHECKSUM = "c9c08f783325c4cf840a94333159cc937f05f75d36a8b307951d5bd959cf2ab8"
)

func downloadFile() error {
	tmp := os.TempDir()

	file, err := os.Create(tmp + "/go1.19.4.linux-amd64.tar.gz")
	if err != nil {
		return err
	}
	defer file.Close()

	res, err := http.Get(URL)
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

func matchChecksum() error {
	file, err := os.ReadFile("/tmp/go1.19.4.linux-amd64.tar.gz")
	if err != nil {
		return err
	}

	checksum := sha256.Sum256(file)

	if hex.EncodeToString(checksum[:]) != CHECKSUM {
		return errors.New("checksums not match")
	}

	return nil
}

func removeAndInstall() error {
	err := os.RemoveAll("/usr/local/go")
	if err != nil {
		return err
	}

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
		// /usr/local/
		switch header.Typeflag {
		case tar.TypeDir:
			if err = os.Mkdir("/usr/local/"+header.Name, 0755); err != nil {
				return errors.New("error at create go file")
			}
		case tar.TypeReg:
			file, err := os.Create("/usr/local/" + header.Name)
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

func Old() {
	// Download new version
	err := downloadFile()
	if err != nil {
		log.Fatal(err)
	}

	// Match checksums
	err = matchChecksum()
	if err != nil {
		log.Fatal(err)
	}

	// Remove previous installation and install new version
	err = removeAndInstall()
	if err != nil {
		log.Fatal(err)
	}

	exec.Command("ls", "-a")
	// check path
}
