package logger

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("should be able to create a logger", func(t *testing.T) {
		logger := NewLogger()
		assert.NotNil(t, logger)
	})

	t.Run("should log info", func(t *testing.T) {
		logger := NewLogger()
		output := captureOutput(func() {
			logger.Info("a", "b")
		})

		assert.Equal(t, "[INFO] - ab\n", output)
	})

	t.Run("should log error", func(t *testing.T) {
		logger := NewLogger()
		output := captureOutput(func() {
			logger.Error("a", "b")
		})

		assert.Equal(t, "[ERROR] - ab\n", output)
	})
}

func captureOutput(f func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}
