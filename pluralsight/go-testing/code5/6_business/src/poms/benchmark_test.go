package poms

import (
	"net/http"
	"os"
	"testing"

	"github.com/GOCODE/pluralsight/go-testing/code2/src/poms/ctrl"
)

func TestMain(m *testing.M) {
	ctrl.Setup()

	go http.ListenAndServe(":3000", new(GZipServer))

	m.Run()

	os.Exit(0)
}

func BenchmarkGetPurchaseOrder(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:3000/api/purchaserOrders/1")
	}
}
