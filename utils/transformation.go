package utils

import (
	"compress/gzip"
	"io"
	"os"

	"bytes"
)

// Error type to determine if gzip succeeded or not
type gzipError struct {
	error_msg   string
	gzip_failed bool
}

func newGzipError(msg string, gzip_fail bool) (err *gzipError) {
	return &gzipError{error_msg: msg, gzip_failed: gzip_fail}
}

func (err *gzipError) Error() (msg string) {
	return err.error_msg
}

// Whether gzip was the cause of failure or not
func (err *gzipError) GzipFailed() (failed bool) {
	return err.gzip_failed
}

// Transforms a reader that is possibly gzipped into another reader,
// if it is not a gzip stream then the data is merely copied, returning the error from gzip decompression,
// if it is a gzip stream then it is uncompressed and copied.
func UnGzipTransform(src_reader io.Reader, read_n int64) (transformed_stream *bytes.Reader, gzip_err *gzipError) {
	if src_reader == nil {
		return nil, newGzipError("util: src_reader cannot be nil", false)
	}

	// copy of src reader in buffer
	src_mem := []byte{}
	src_buf := bytes.NewBuffer(src_mem)

	_, err := io.CopyN(src_buf, src_reader, read_n)
	if err != nil && err != io.EOF {
		return nil, newGzipError(err.Error(), false)
	}

	src_seek := bytes.NewReader(src_buf.Bytes())

	// dest buffer to write to from gzip reader
	dst_mem := []byte{}
	dst_buf := bytes.NewBuffer(dst_mem)

	// this automatically reads for header check, so we need to seek to start again
	gzip_reader, err := gzip.NewReader(src_seek)
	if gzip_reader != nil {
		defer gzip_reader.Close()
	}
	// gzip failed somehow, just return copied reader after seeking to start
	if err != nil {
		_, seek_err := src_seek.Seek(0, os.SEEK_SET)
		if seek_err != nil {
			return nil, newGzipError(seek_err.Error(), false)
		}
		return src_seek, newGzipError(err.Error(), true)
	}
	_, err = io.CopyN(dst_buf, gzip_reader, read_n)
	if err != nil && err != io.EOF {
		return nil, newGzipError(err.Error(), false)
	}

	// src_reader was gzip compressed and uncompressed properly
	return bytes.NewReader(dst_buf.Bytes()), nil
}
