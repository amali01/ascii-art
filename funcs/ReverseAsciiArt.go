package funcs

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReverseAsciiArt() {
	// the argument must be BannerLineed from --reverse= , checks the validity of the arguments
	if len(os.Args) != 2 || len(os.Args[1]) < 9 || os.Args[1][0:10] != "--reverse=" {
		fmt.Println(ErrorsM("revUsage"))
		os.Exit(0)
	}

	//read art file for reverse
	textFile := os.Args[1][10:] // 10 is the len of --reverse=
	file, err := os.Open(textFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if ReadInput(file) {
		fmt.Println(textFile)

		textContent, err := os.ReadFile(textFile)
		if err != nil {
			fmt.Println(ErrorsM("textCont"))
			os.Exit(0)
		}
		textData := string(textContent)
		textDatatest := strings.TrimSpace(textData)
		asciiArt := strings.Split(textData, "\n")

		BannerName := ""
		for _, supStr := range textDatatest {
			if strings.Contains(string(supStr), "o") {
				BannerName = "thinkertoy.txt"
				break
			} else if string(supStr) == "V" || string(supStr) == ")" || string(supStr) == "(" || string(supStr) == "/" ||
				string(supStr) == "\\" || string(supStr) == "<" || string(supStr) == ">" || string(supStr) == "'" {
				BannerName = "standard.txt"
				break
			} else {
				BannerName = "shadow.txt"
			}
		}

		// if we get from checkArgs true, work BannerLines =)
		//read Banner art
		fmt.Println("BannerName:", BannerName)
		fContent, err := os.ReadFile(BannerName)
		if err != nil {
			fmt.Println(ErrorsM("fCont"))
			os.Exit(0)
		}
		fontData := string(fContent)
		Banner := strings.Split(fontData, "\n")
		c := 0

		//If the file contains more than one object ascii-art
		if len(asciiArt) > 9 {
			for i := 0; i < len(asciiArt)-1; {
				if len(asciiArt[i]) > 0 {
					c = i + 8
					if len(Banner)-c < 8 && len(Banner)-c != 0 {
						fmt.Println("111Error with ascii art !!!!@@@@")
						os.Exit(0)
					}
					reverse(Banner, asciiArt[i:c], 0, 0, 1)
					fmt.Println()
					i = i + 8
				} else {
					fmt.Println()
					i = i + 1
				}
			}
		} else {
			reverse(Banner, asciiArt, 0, 0, 1)
		}
		fmt.Println()
	}

	os.Exit(0)
}

var x = 0

// LattersFound - symbol in ascii-art file(input.txt). ArtLine - count of lines in ascii-art file(input.txt). BannerLine - number of line in Banner(standard.txt)
func reverse(Banner []string, asciiArt []string, LattersFound int, ArtLine int, BannerLine int) {

	if LattersFound != len(asciiArt[ArtLine]) { //if we are not finish reverse.

		bannerWidth := len(Banner[BannerLine]) // len of candite for research

		if LattersFound+bannerWidth <= len(asciiArt[ArtLine]) {

			if ArtLine < 7 { // befor the last line
				if Banner[BannerLine+ArtLine] == asciiArt[ArtLine][LattersFound:LattersFound+bannerWidth] { // if the Banner-line equal the sup line of ascii-art
					reverse(Banner, asciiArt, LattersFound, ArtLine+1, BannerLine) // compare the next lines "next rows of each"
				} else {
					x = BannerLine + 9
					if len(Banner)-1-x < 9 && len(Banner)-1-x != 0 {
						fmt.Println("222Error with ascii art or not a standard banner!!!!@@@@")
						os.Exit(0)
					}
					reverse(Banner, asciiArt, LattersFound, 0, x) // if not equal we try the next symbol in Banner-file and cheack from the biganning of the rows "line of ascii-art = 0 && BannerLine+9"
				}
			} else { // at last line "ArtLine=8" add it than print and redo with for the next latter
				r := ((BannerLine - 1) / 9) + 32 // find and print the letter
				fmt.Printf("%c", r)
				reverse(Banner, asciiArt, LattersFound+bannerWidth, 0, 1) // we redo with the next sup line of the ascii-art
			}

		} else {
			x = BannerLine + 9
			if len(Banner)-1-x < 9 && len(Banner)-1-x != 0 {
				fmt.Println("Error with ascii art or not any of the (standard,thinkertoy,shadow) banners !!!!@@@@")
				os.Exit(0)
			}
			reverse(Banner, asciiArt, LattersFound, 0, x) // if not equal we try the next symbol in Banner-file and cheack from the biganning of the rows "line of ascii-art = 0 && BannerLine+9"
		}

	}

}

func ReadInput(file io.Reader) bool {
	// Initialize variables to store tetromino data and parsing state
	scanner := bufio.NewScanner(file)
	lineCount, wordNum, latterWidth := 0, 0, 0
	counter := 0
	// Read lines from the input file
	for scanner.Scan() {
		counter++
		line := scanner.Text()
		if lineCount == 0 {
			latterWidth = len(line)
		}

		// Check for an empty line indicating newline
		if line == "" {
			continue
		}

		if len(line) != latterWidth {
			// Check for invalid line length
			fmt.Printf("ERROR3: at line %d\n", counter)
			// fmt.Println("latterWidth", latterWidth)
			// fmt.Println("len(line)", len(line))
			// fmt.Println("lineCount", lineCount)
			// fmt.Println("wordNum", wordNum)
			// fmt.Println("cc", counter)

			os.Exit(0)
		}

		lineCount++
		if lineCount == 8 {
			lineCount = 0
			wordNum++

		}
	}

	// Check for scanner errors and handle them
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return true
}
