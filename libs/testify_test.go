package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t) // Membuat objek assertion

	result := 2 + 3
	assert.Equal(5, result, "Hasil harus 5") // Tidak perlu menyertakan `t` lagi
	assert.NotEqual(6, result, "Hasil tidak boleh 6")
}
