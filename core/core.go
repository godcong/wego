package core

import (
	"github.com/godcong/wego/log"
	"golang.org/x/text/transform"
	"os"
)

// SaveTo ...
func SaveTo(response Responder, path string) error {
	var err error
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Debug("Responder|ToFile", err)
		return err
	}
	defer func() {
		err = file.Close()
	}()
	_, err = file.Write(response.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// SaveEncodingTo ...
func SaveEncodingTo(response Responder, path string, t transform.Transformer) error {
	var err error
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Debug("Responder|ToFile", err)
		return err
	}
	defer func() {
		err = file.Close()
	}()
	writer := transform.NewWriter(file, t)
	_, err = writer.Write(response.Bytes())
	if err != nil {
		return err
	}
	defer func() {
		err = writer.Close()
	}()
	return nil
}
