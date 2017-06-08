package poms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/GOCODE/pluralsight/go-testing/code7/1_initial/src/poms/ctrl"
	"github.com/GOCODE/pluralsight/go-testing/code7/1_initial/src/poms/model"
)

func TestMain(m *testing.M) {
	ctrl.Setup()

	go http.ListenAndServe(":3000", new(GZipServer))

	m.Run()

	os.Exit(0)
}

func TestGetVendors(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/api/vendors")

	if err == nil {
		var vendors []model.Vendor
		data, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			resp.Body.Read(data)

			err = json.Unmarshal(data, &vendors)
		}

	}

	if err != nil {
		t.Error("Failed to retrieve vendors")
	}
}
