package poms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"poms/ctrl"
	"poms/model"
	"testing"
)

func TestMain(m *testing.M) {
	ms := setupMockService()
	defer ms.Close()

	model.VendorServiceURL = ms.URL

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

	if err != nil || resp.StatusCode == 500 {
		t.Error("Failed to retrieve vendors")
	}
}

func setupMockService() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var vendors []*Vendor
		vendors = GetVendors()
		data, _ := json.Marshal(vendors)

		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}))

}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Vendor struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Contact *Contact `json:"contact"`
}

func GetVendors() []*Vendor {
	result := []*Vendor{
		{
			Id:   1,
			Name: "Smith and Jones Mfg",
			Contact: &Contact{
				Name:  "Martha Jones",
				Phone: "+1 (888)555-9639",
			},
		}, {
			Id:   2,
			Name: "Oswald Unlimited",
			Contact: &Contact{
				Name:  "Clara Smith",
				Phone: "(926)555-1798",
			},
		}, {
			Id:   3,
			Name: "Noble Products",
			Contact: &Contact{
				Name:  "Brian Jeffries",
				Phone: "(425)555-1998",
			},
		},
	}
	return result
}
