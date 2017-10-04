// datastore_test.go
package backend

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchrcom/testify/assert"
)

/*
func Test_(t *testing.T) {

}
*/

func Test_InvalidPathNew(t *testing.T) {
	invalid_path := "barf"
	file_store, err := New(invalid_path)

	assert.True(t, err != nil, fmt.Sprintf("Accepted bogus path: %s", err))
	assert.True(t, file_store == nil, fmt.Sprintf("No nil on error: %s", err))
}

func Test_CreateDeleteFile(t *testing.T) {
	test_dir := os.TempDir()
	file_store, err := New(test_dir)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("file store failed to create: %s", err))
	}
	full_test_path := ""
	// Create file
	test_file := `invalid/\:jargon`
	file, err := file_store.OpenFile(test_file, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	assert.True(t, err == nil, fmt.Sprintf("%s", err))
	if file != nil {
		t.Logf("File generated \"%s\" from \"%s\"", file.Name(), test_file)
		assert.Equal(t, test_file, file.Original_Name(),
			fmt.Sprintf("File mapping name mismatch, original name %s, original name on file %s", test_file, file.Original_Name()))
		full_test_path = file.Name()
		file.Close()
	}

	// Delete file
	err = file_store.DeleteFile(test_file)
	assert.True(t, err == nil, fmt.Sprintf("%s", err))
	_, exist_err := os.Stat(full_test_path)
	assert.True(t, exist_err != nil)

	// Cleanup in failure case
	if err != nil {
		os.Remove(full_test_path)
	}
}
