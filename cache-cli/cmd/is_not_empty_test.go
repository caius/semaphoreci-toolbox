package cmd

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/semaphoreci/toolbox/cache-cli/pkg/storage"
	assert "github.com/stretchr/testify/assert"
)

func Test__IsNotEmpty(t *testing.T) {
	runTestForAllBackends(t, func(backend string, storage storage.Storage) {
		t.Run(fmt.Sprintf("%s cache is empty", backend), func(*testing.T) {
			storage.Clear()
			assert.False(t, RunIsNotEmpty(isNotEmptyCmd, []string{}))
		})

		t.Run(fmt.Sprintf("%s cache is not empty", backend), func(*testing.T) {
			storage.Clear()
			tempFile, _ := ioutil.TempFile("/tmp", "*")
			storage.Store("abc001", tempFile.Name())

			assert.True(t, RunIsNotEmpty(isNotEmptyCmd, []string{}))
		})
	})
}
