package s3share

import "testing"

func TestUploadToS3(t *testing.T) {

	fileName := "testdata/L34-exercise-presign.pdf"
	bucketName := "testBucket"

	tc := &TestClient{}

	err := UploadToS3(tc, &fileName, &bucketName)
	if err != nil {
		t.Errorf("Error uploading to S3: %s", err)
	}

}
