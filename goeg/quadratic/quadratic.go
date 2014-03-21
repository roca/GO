// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"math/cmplx"
	"net/http"
	//"sort"
	"strconv"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Quadratic</title>
<body><h3>Quadratic Equation Solver</h3>
<p>Solves equations of the for ax<sup>2</sup> + bx + c </p>`
	form = `<form action="/" method="POST">
	<p>
<input type="text" name="numberA" size="10"> x<sup>2</sup> + 
<input type="text" name="numberB" size="10"> x +
<input type="text" name="numberC" size="10"> -> 
<input type="submit" value="Calculate">
</p>

</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	sd      float64
	mode    []float64
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {

		if numbers, message, ok := processRequest(request); ok {

			x1, x2 := solve(numbers)
			fmt.Fprintf(writer, strconv.FormatFloat(real(x1), 'f', 4, 64)+", ")
			fmt.Fprintf(writer, strconv.FormatFloat(real(x2), 'f', 4, 64)+", ")
			//fmt.Fprint(writer, formatSolutions(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}

	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	var textFields []string

	textFields = append(textFields, request.FormValue("numberA"))
	textFields = append(textFields, request.FormValue("numberB"))
	textFields = append(textFields, request.FormValue("numberC"))

	for _, text := range textFields {
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Standard Deviation</td><td>%f</td></tr>
<tr><td>Mode</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.sd, stats.mode)
}

func solve(numbers []float64) (complex128, complex128) {

	var a, b, c complex128

	a = complex(numbers[0], 0)
	b = complex(numbers[1], 0)
	c = complex(numbers[2], 0)

	sqrt := (cmplx.Sqrt((b * b) - (4 * a * c)))

	x1 := (-b + sqrt) / (2 * a)
	x2 := (-b - sqrt) / (2 * a)

	return x1, x2

}

func formatQuestion(answer complex128) string {
	return ""
}

func formatSolutions(answer complex128) string {
	return ""
}
