package main

import (
	"bufio"
	"fmt"
	"go/token"
	"go/types"
	"os"
	"strings"
)

const addition = "abcd"
const subtraction = "bcde"
const multiplication = "dede"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the String: ")
	str, err := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)
	if err != nil {
		panic(err)
	}
	str = setOperation(str)
	strList := strings.Split(str, " ")
	strList = calculate(strList)
	strList = strings.Split(strings.Join(strList, ""), " ")
	strList = splitCharacter(strList)
	fs := token.NewFileSet()
	finalStr:= isExpTrust(strings.Join(strList, ""))
	tv, err := types.Eval(fs, nil, token.NoPos, finalStr)
	if err != nil {
		panic(err)
	}

	println("\nExperction: ", finalStr)
	println("\nExpected result: ", tv.Value.String())
}

func setOperation(str string) string {
	str = strings.ReplaceAll(str, addition, "+")
	str = strings.ReplaceAll(str, subtraction, "/")
	str = strings.ReplaceAll(str, multiplication, "*")
	return str
}

func isExpTrust(str string) string {
	str = strings.ReplaceAll(str,"+)",")")
	str = strings.ReplaceAll(str,"*)",")")
	str = strings.ReplaceAll(str,"/)",")")
	str = strings.ReplaceAll(str,"+*)",")")
	str = strings.ReplaceAll(str,"+/)",")")
	str = strings.ReplaceAll(str,"++","")
	str = strings.ReplaceAll(str,"**","")
	str = strings.ReplaceAll(str,"//","")
	return str
}

func splitCharacter(strItems []string) []string {
	var list []string
	var op = "+"
	for i := 0; i < len(strItems); i++ {
		var strNew string
		if string(strItems[i][0]) != "*" && string(strItems[i][0]) != "+" && string(strItems[i][0]) != "/" {
			strNew = "("
		}
		for index, val := range strItems[i] {
			switch string(val) {
			case "a":
				if index+1 == len(strItems[i]) {
					strNew = strNew + "1)%5^2+"
					break
				} else if len(strItems) == i+1 && index+1 == len(strItems[i]) {
					strNew = strNew + "1)%5^2"
					break
				}
				strNew = strNew + "1" + op
				break
			case "b":
				if index+1 == len(strItems[i]) {
					strNew = strNew + "2)%5^2+"
					break
				} else if len(strItems)-1 == i && index+1 == len(strItems[i]) {
					strNew = strNew + "2)%5^2"
					break
				}
				strNew = strNew + "2" + op
				break
			case "c":
				if index+1 == len(strItems[i]) && len(strItems)-1 != i {
					strNew = strNew + "3)%5^2+"
					break
				} else if len(strItems)-1 == i && index+1 == len(strItems[i]) {
					strNew = strNew + "3)%5^2"
					break
				}
				strNew = strNew + "3" + op
				break
			case "e":
				if index+1 == len(strItems[i]) && len(strItems)-1 != i {
					strNew = strNew + "4)%5^2+"
					break
				} else if len(strItems)-1 == i && index+1 == len(strItems[i]) {
					strNew = strNew + "4)%5^2"
					break
				}
				strNew = strNew + "4" + op
				break
			case "d":
				if index+1 == len(strItems[i]) && len(strItems)-1 != i {
					strNew = strNew + "5)%5^2+"
					break
				} else if len(strItems)-1 == i && index+1 == len(strItems[i]) {
					strNew = strNew + "5)%5^2"
					break
				}
				strNew = strNew + "5" + op
				break
			case ")":
				if index+1 == len(strItems[i]) && index+1 < len(strItems[i]) {
					strNew = strNew + ")%5^2+"
					break
				} else if len(strItems)-1 == i && index+1 == len(strItems[i]) {
					strNew = strNew + ")%5^2)"
					break
				}
				if index+1 == len(strItems[i]) && string(strItems[i][index]) == ")"{
					strNew = strNew + ")%5^2)+"
					break
				}
				strNew = strNew + ")" + op
				break
			default:
				strNew = strNew + string(val)
				break
			}
		}
		list = append(list, strNew)
		switch strItems[i] {
		case "+":
			op = "+"
			break
		case "*":
			op = "*"
			break
		case "/":
			op = "/"
			break
		default:
			op = "+"
			break
		}
	}
	return list
}

func calculate(str []string) []string {
	var s = ""
	var mm []string
	for i, v := range str {
		var valueCh string
		for _, vv := range v {
			if string(vv) != "*" && string(vv) != "+" && string(vv) != "/" && string(vv) != "(" && string(vv) != ")" {
				switch string(vv) {
				case "a":
					if string(vv) == s {
						valueCh = valueCh + "a"
					} else if string(vv) != s && s != "" {
						valueCh = valueCh + ")(a"
					} else {
						valueCh = valueCh + "(a"
					}
					break
				case "b":
					if string(vv) == s {
						valueCh = valueCh + "b"
					} else if string(vv) != s && s != "" {
						valueCh = valueCh + ") (b"
					} else {
						valueCh = valueCh + "(b"
					}
					break

				case "c":
					if string(vv) == s {
						valueCh = valueCh + "c"
					} else if string(vv) != s && s != "" {
						valueCh = valueCh + ") (c"
					} else {
						valueCh = valueCh + "(c"
					}
					break

				case "d":
					if string(vv) == s {
						valueCh = valueCh + "d"
					} else if string(vv) != s && s != "" {
						valueCh = valueCh + ") (d"
					} else {
						valueCh = valueCh + "(d"
					}
					break

				case "e":
					if string(vv) == s {
						valueCh = valueCh + "e"
					} else if string(vv) != s && s != "" {
						valueCh = valueCh + ") (e"
					} else {
						valueCh = valueCh + "(e"
					}
					break
				default:
					break
				}
				s = string(vv)
			} else {
				valueCh = string(vv)
			}
		}
		mm = append(mm, valueCh)
		if len(str)-1 == i {
			mm[i] = valueCh + ")"
		}

	}
	return mm
}
