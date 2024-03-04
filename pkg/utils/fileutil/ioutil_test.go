package fileutil

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFileWithDataFunc(t *testing.T) {
	content := "hello\nworld\n"
	tests := []struct {
		filename string
		flag     int
		mode     os.FileMode
		dataFunc func(w io.Writer) error
	}{
		{
			filename: "./test.txt",
			flag:     os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
			mode:     0644,
			dataFunc: func(w io.Writer) error {
				_, err := w.Write([]byte(content))
				return err
			},
		},
	}
	for _, tt := range tests {
		err := WriteFileWithDataFunc(tt.filename, tt.flag, tt.mode, tt.dataFunc, false)
		assert.NoError(t, err)
		data, err := os.ReadFile(tt.filename)
		assert.NoError(t, err)
		assert.Equal(t, content, string(data))
		os.Remove(tt.filename)
	}
}
