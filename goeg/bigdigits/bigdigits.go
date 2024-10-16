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
	"flag"
	"fmt"
	"log"
	"os"
	//"path/filepath"
	"strings"
)

func main() {
	var bar = flag.Bool("bar", false, "bar flag")
	var b = flag.Bool("b", false, "bar flag")

	var help = flag.Bool("help", false, "help flag")
	var h = flag.Bool("h", false, "help flag")

	flag.Parse()

	if (len(os.Args) == 1) || *help || *h {
		fmt.Println("usage: bigdigits [-b|--bar] <whole-number>")
		fmt.Println("-b --bar draw an underbar and an overbar")
		os.Exit(1)
	}

	stringOfDigits := os.Args[len(os.Args)-1]
	lines := map[int]string{}
	maxRowLength := 0
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		lines[row] = line
		if len(line) > maxRowLength {
			maxRowLength = len(line)
		}
	}

	if *bar || *b {
		fmt.Println(strings.Repeat("*", maxRowLength))
	}
	for _, value := range lines {
		fmt.Println(value)
	}

	if *bar || *b {
		fmt.Println(strings.Repeat("*", maxRowLength))
	}

}

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
