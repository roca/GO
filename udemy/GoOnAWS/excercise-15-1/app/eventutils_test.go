// begin package
package dslapp_test

//end package

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"gotest.tools/assert"

	"dslapp"

	"github.com/aws/aws-lambda-go/events"
)

// begin testfunction
func TestAppExtractObject(t *testing.T) {
	//end testfunction

	//begin mockevent
	var s3event events.S3Event

	const testfile = "testdata/put.json"
	jsonFile, err := os.Open(testfile)
	if err != nil {
		t.Log(err)
	}
	t.Log("Successfully Opened ", testfile)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal([]byte(byteValue), &s3event)
	assert.Equal(t, nil, err)
	//end mockevent

	//begin assertion
	expectedKey := "my2etestkey.txt"
	realKey := dslapp.ExtractKey(s3event)
	assert.Equal(t, expectedKey, realKey)
	//end assertion

}
