// +build unit

package Testing

import (
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world" {
		t.Fatal("Contents do not match expected")
	}
}
