package sha1

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func Sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}

	sha := sha1.New()

	if _, err = io.Copy(sha, r); err != nil {
		return "", err
	}

	signedSum := sha.Sum(nil)

	return fmt.Sprintf("%x", signedSum), nil
}