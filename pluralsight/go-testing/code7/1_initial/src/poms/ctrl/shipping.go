package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/GOCODE/pluralsight/go-testing/code7/1_initial/src/poms/model"
)

type ShippingController struct{}

func (sc *ShippingController) GetReceivers(w http.ResponseWriter, r *http.Request) {
	receivers := model.GetReceivers()

	w.Header().Add("Content-Type", "application/json")

	data, _ := json.Marshal(receivers)

	w.Write(data)
}

func (sc *ShippingController) GetVendors(w http.ResponseWriter, r *http.Request) {
	vendors, err := model.GetVendors()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	data, _ := json.Marshal(vendors)

	w.Write(data)
}
