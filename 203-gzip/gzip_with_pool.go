package main

import (
	"bytes"
	"compress/gzip"
	"sync"
)

type GzipWithPoolModel struct {
	pool sync.Pool
}

func NewGzipWithPool() *GzipWithPoolModel {
	return &GzipWithPoolModel{
		pool: sync.Pool{
			New: func() interface{} {
				var buffer bytes.Buffer
				return gzip.NewWriter(&buffer)
			},
		},
	}
}

func (gzipWithPoolModel *GzipWithPoolModel) GetGzipWriter() *gzip.Writer {
	return gzipWithPoolModel.pool.Get().(*gzip.Writer)
}

func (gzipWithPoolModel *GzipWithPoolModel) PutGzipWriter(buf *gzip.Writer) {
	buf.Flush()
	gzipWithPoolModel.pool.Put(buf)
}
