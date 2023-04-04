package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

type GzipWithoutPoolModel struct{}

func NewGzipWithoutPool() *GzipWithoutPoolModel {
	return &GzipWithoutPoolModel{}
}

func (gzipWithoutPoolModel *GzipWithoutPoolModel) Zip(str string) error {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)

	_, error := writer.Write([]byte(str))
	if error != nil {
		fmt.Println(error)
	}

	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}
