// File name: ...\s15\01_call_by_ref\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	// Uncomment one at a time
	testCase1()
	// testCase2()
	// testCase3()
	// testCase4()

}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
//'int' is a non-reference type
func testCase1() {
	x1 := 10
	x2 := 20

	fmt.Printf("x1=%d x2=%d\n", x1, x2)
	f1(x1, &x2)
	fmt.Printf("x1=%d x2=%d\n", x1, x2)
}

func f1(y1 int, y2 *int) {
	y1++
	*y2++
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
//'string' is a non-reference type
func testCase2() {
	s1 := "Hello "
	s2 := "Hi "

	fmt.Printf("s1=%s s2=%s\n", s1, s2)
	f2(s1, &s2)
	fmt.Printf("s1=%s s2=%s\n", s1, s2)
}

func f2(q1 string, q2 *string) {
	q1 += "there!"
	*q2 += "there!"
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
//'slice' is a reference type; 'array' is a non-ref type
const aLen = 7

func testCase3() {
	langs := [aLen]string{1: "Go", 2: "C", 3: "Java", 4: "Python", 5: "JS", 6: "SQL"}
	fmt.Printf("%v %v %T \n", langs, len(langs), langs)

	processA1(langs)
	fmt.Println(langs, len(langs))

	processA2(&langs)
	fmt.Println(langs, len(langs))

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	fmt.Println()
	favLangs := langs[1:4]
	fmt.Println(favLangs, len(favLangs))

	processS1(favLangs)
	fmt.Println(favLangs, len(favLangs))

	processA3(&langs)
	fmt.Println(langs, len(langs))
}

func processA1(lang [aLen]string) {
	lang[1] = "GoLang"
}

func processA2(lang *[aLen]string) {
	lang[2] = "C++"
}

func processA3(lang *[aLen]string) {
	*lang = [aLen]string{}
}

func processS1(lang []string) {
	lang[0] = "GoLang"
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
//'map' is a reference type
func testCase4() {
	sal := map[string]float64{
		"Blake":  60000.00,
		"Parker": 70000.00,
		"Dakota": 80000.00,
	}

	fmt.Printf("sal=%v\n", sal)
	f4(sal)
	fmt.Printf("sal=%v\n", sal)
}

func f4(sal map[string]float64) {

	fmt.Println("")
	for name, sal := range sal {
		fmt.Printf("%-10s %.2f\n", name, sal)
	}

	sal["Parker"] += 9999
}
