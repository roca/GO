package ctrl

import "testing"
import "net/http"

func TestCurrencyController(t *testing.T) {
	// arrange

	//act

	//assert

	// sets content-type Header

	//writes correct data to client
}

type mockResponseWriter struct {
	header       http.Header
	capturedData []byte
}

func (mrw mockResponseWriter) Header() http.Header {
	return mrw.header
}

func (mrw mockResponseWriter) Write(data []byte) (int, error) {
	mrw.capturedData = data
	return len(data), nil
}

func (mrw mockResponseWriter) WriteHeader(code int) {

}
