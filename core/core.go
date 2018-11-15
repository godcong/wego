package core

import (
	"github.com/godcong/wego/log"
	"golang.org/x/text/transform"
	"os"
)

// SaveTo ...
func SaveTo(response Responder, path string) error {
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if e != nil {
		log.Debug("Responder|ToFile", e)
		return e
	}
	defer file.Close()
	_, e = file.Write(response.Bytes())
	if e != nil {
		return e
	}
	return nil
}

// SaveEncodingTo ...
func SaveEncodingTo(response Responder, path string, t transform.Transformer) error {
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if e != nil {
		log.Debug("Responder|ToFile", e)
		return e
	}
	defer file.Close()
	writer := transform.NewWriter(file, t)
	_, e = writer.Write(response.Bytes())
	if e != nil {
		return e
	}
	defer writer.Close()
	return nil
}
