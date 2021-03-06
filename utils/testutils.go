package testutils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "bazel-remote")
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func CreateRandomFile(dir string, size int64) (string, error) {
	data, hash := RandomDataAndHash(size)
	os.MkdirAll(dir, os.FileMode(0744))
	filepath := dir + "/" + hash

	return hash, ioutil.WriteFile(filepath, data, 0744)
}

func CreateCacheFile(dir string, size int64) (string, error) {
	data, hash := RandomDataAndHash(size)
	subdir := dir + "/" + hash[0:2]
	os.MkdirAll(subdir, os.FileMode(0744))
	filepath := subdir + "/" + hash

	return hash, ioutil.WriteFile(filepath, data, 0744)
}

func RandomDataAndHash(size int64) ([]byte, string) {
	data := make([]byte, size)
	rand.Read(data)
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])
	return data, hashStr
}

func CreateTmpCacheDirs(t *testing.T) string {
	path, err := ioutil.TempDir("", "bazel-remote-test")
	if err != nil {
		t.Error("Couldn't create tmp dir", err)
	}
	os.MkdirAll(path, os.FileMode(0744))
	return path
}

// NewSilentLogger returns a cheap logger that doesn't print anything, useful
// for tests.
func NewSilentLogger() *log.Logger {
	return log.New(ioutil.Discard, "", 0)
}
