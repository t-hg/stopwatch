package style

import (
	"strconv"
	"strings"
)

const figlet0 = `
  ___  
 / _ \ 
| | | |
| |_| |
 \___/ `

const figlet1 = `
 _ 
/ |
| |
| |
|_|`

const figlet2 = `
 ____  
|___ \ 
  __) |
 / __/ 
|_____|`

const figlet3 = `
 _____ 
|___ / 
  |_ \ 
 ___) |
|____/ `

const figlet4 = `
 _  _   
| || |  
| || |_ 
|__   _|
   |_|  `

const figlet5 = `
 ____  
| ___| 
|___ \ 
 ___) |
|____/ `

const figlet6 = `
  __   
 / /_  
| '_ \ 
| (_) |
 \___/ `

const figlet7 = `
 _____ 
|___  |
   / / 
  / /  
 /_/   `

const figlet8 = `
  ___  
 ( _ ) 
 / _ \ 
| (_) |
 \___/ `

const figlet9 = `
  ___  
 / _ \ 
| (_) |
 \__, |
   /_/ `

const figletColon = `
   
 _ 
(_)
 _ 
(_)`

const figletSpace = `
 
 
 
 
 `

const figletEmtpy = `




`

var FigletNums = []string{
	figlet0,
	figlet1,
	figlet2,
	figlet3,
	figlet4,
	figlet5,
	figlet6,
	figlet7,
	figlet8,
	figlet9,
}

func figletAppend(figlets string, figlet string) string {
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
	figletized := figletEmtpy
	for _, char := range str {
		switch char {
		case ' ':
			figletized = figletAppend(figletized, figletSpace)
		case ':':
			figletized = figletAppend(figletized, figletColon)
		default:
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			figletized = figletAppend(figletized, FigletNums[digit])
		}
	}
	return figletized
}
