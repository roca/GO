package share

import "testing"

func TestUploadToS3(t *testing.T) {

	filenName := "excercise-10-2/testdata/L34-exercise-presign.pdf"
	bucktName := "testBucket"

	tc := &testClient{}

	err := UploadToS3(tc, &fileName, &bucketName)

}
