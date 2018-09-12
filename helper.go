// Helpers package

package helpers

import (
	//"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var FlagDbg bool

func ToStr(values ...interface{}) string {
	var retsrt string
	for _, value := range values {
		if value != nil {
			switch value.(type) {
			case string:
				retsrt += value.(string)
			case bool:
				retsrt += strconv.FormatBool(value.(bool))
			case int:
				retsrt += strconv.Itoa(value.(int))
			case uint8:
				retsrt += strconv.Itoa(int(value.(uint8)))
			case int16:
				retsrt += strconv.Itoa(int(value.(int16)))
			case []byte:
				retsrt += string(value.([]byte))
			case error:
				retsrt += value.(error).Error()
			case []string:
				newsrt := ""
				for _, v := range value.([]string) {
					newsrt += v
				}
				retsrt += newsrt
			case []int:
				newsrt := ""
				for _, v := range value.([]int) {
					newsrt += strconv.Itoa(v)
				}
				retsrt += newsrt
			default:
				fmt.Println("Error! ToStr")
			}
		}
	}
	return retsrt
}

func ToInt(values ...interface{}) int {
	var retsrt int
	for _, value := range values {
		if value != nil {
			switch value.(type) {
			case string:
				tmp, _ := strconv.Atoi(value.(string))
				retsrt += tmp
			case bool:
				if value.(bool) {
					retsrt += 1
				} else {
					retsrt += 0
				}
			case int:
				retsrt += value.(int)
			case uint8:
				retsrt += int(value.(uint8))
			case int16:
				retsrt += int(value.(int16))
			case []byte:
				tmp, _ := strconv.Atoi(string(value.([]byte)))
				retsrt += tmp
			case []string:
				newsrt := 0
				for _, v := range value.([]string) {
					tmp, _ := strconv.Atoi(v)
					newsrt += tmp
				}
				retsrt += newsrt
			case []int:
				newsrt := 0
				for _, v := range value.([]int) {
					newsrt += v
				}
				retsrt += newsrt
			default:
				fmt.Println("Error! ToInt")
			}
		}
	}
	return retsrt
}

func IndexArray(index int, arr ...interface{}) interface{} {
	if index < 0 {
		return nil
	}
	var items []interface{}
	for _, text := range arr {
		var item interface{}
		if text != nil {
			bInt, fInt := text.([]int)
			bErr, fErr := text.([]error)
			bStr, fStr := text.([]string)
			bByt, fByt := text.([]byte)
			bRun, fRun := text.([]rune)
			if fInt {
				if len(bInt)-1 < index {
					return nil
				} else {
					return bInt[index]
				}
			}
			if fErr {
				if len(bErr)-1 < index {
					return nil
				} else {
					return bErr[index]
				}
			}
			if fStr {
				if len(bStr)-1 < index {
					return nil
				} else {
					return bStr[index]
				}
			}
			if fByt {
				if len(bByt)-1 < index {
					return nil
				} else {
					return bByt[index]
				}
			}
			if fRun {
				if len(bRun)-1 < index {
					return nil
				} else {
					return bRun[index]
				}
			}
		}
		items = append(items, item)
	}
	return nil
}

func ToInt2(value string) int {
	var i int
	if _, err := fmt.Sscanf(value, "%d", &i); err == nil {
		Dbg(err)
	}
	return i
}

func Dbg(texts ...interface{}) {
	if FlagDbg {
		for _, text := range texts {
			if text != nil {
				fmt.Printf("%s \n", ToStr(text))
			}
		}
	}
}

func ToTime(text string) time.Duration {
	format, _ := time.ParseDuration(text)
	return format

}

func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:4001/").Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}
