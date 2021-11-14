package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LineCounter(str string) (int, error) {
	var count int
	rr := strings.Split(str, "\n")
	var rrr []string
	cc := 0
	comstart := false

	for _, s := range rr {
		//kiểm tra xem chuỗi s có bắt đầu bằng "/*" hay không
		if strings.HasPrefix(s, "/*") {
			comstart = true
			continue
		}


		if strings.Contains(s, "*/") {
			comstart = false
			cc += 1
		}


		if (str != "" && !strings.HasPrefix(str, "//") && !comstart) {
			rrr = append(rrr, str)
		}
	}

	count = len(rrr) - cc
	return count, nil
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func main() {
	arg := os.Args[1]
	all := 0

	err := filepath.Walk(arg, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if (!info.IsDir()  && !strings.Contains(info.Name(), "test") ){
			ext := filepath.Ext(path)


			if ext == ".go" {
				jj, _ := os.ReadFile(path)
				g, _ := LineCounter(string(jj))
				all += g
				var u string

				if g < 200 {
					z := strconv.Itoa(g)
					u = Info(z)
				}

				if g >= 200 && g < 1000 {
					z := strconv.Itoa(g)
					u = Warn(z)
				}

				if g >= 1000 {
					z := strconv.Itoa(g)
					u = Fata(z)
				}

				fmt.Printf("File %s have %s lines of code \n",path, u)
			}
		}

		return nil
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("All lines", all)

}