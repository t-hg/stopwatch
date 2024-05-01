package stopwatch

import (
	"strconv"
	"strings"
)

const F0 = `
  ___  
 / _ \ 
| | | |
| |_| |
 \___/ `

const F1 = `
 _ 
/ |
| |
| |
|_|`

const F2 = `
 ____  
|___ \ 
  __) |
 / __/ 
|_____|`

const F3 = `
 _____ 
|___ / 
  |_ \ 
 ___) |
|____/ `

const F4 = `
 _  _   
| || |  
| || |_ 
|__   _|
   |_|  `

const F5 = `
 ____  
| ___| 
|___ \ 
 ___) |
|____/ `

const F6 = `
  __   
 / /_  
| '_ \ 
| (_) |
 \___/ `

const F7 = `
 _____ 
|___  |
   / / 
  / /  
 /_/   `

const F8 = `
  ___  
 ( _ ) 
 / _ \ 
| (_) |
 \___/ `

const F9 = `
  ___  
 / _ \ 
| (_) |
 \__, |
   /_/ `

const FC = `
   
 _ 
(_)
 _ 
(_)`

const FS = `
 
 
 
 
 `

const FE = `




`

var FigletNums = []string{F0, F1, F2, F3, F4, F5, F6, F7, F8, F9}

func FigletAppend(figlets string, figlet string) string {
	figletsLines := strings.Split(figlets, "\n")
	figletLines := strings.Split(figlet, "\n")
	if len(figletsLines) != len(figletLines) {
		panic("FigletAppend: length mismatch")
	}
	for i := range figletsLines {
		figletsLines[i] += figletLines[i]
	}
	return strings.Join(figletsLines, "\n")
}

func Figletize(str string) string {
	figletized := FE
	for _, char := range str {
		switch char {
		case ' ':
			figletized = FigletAppend(figletized, FS)
		case ':':
			figletized = FigletAppend(figletized, FC)
		default:
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			} 
			figletized = FigletAppend(figletized, FigletNums[digit])
		}
	}
	return figletized
}
