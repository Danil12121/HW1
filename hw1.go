package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func findUnical(lines []string) []string {
	mapLines := make(map[string]int, len(lines))
	for _, line := range lines {
		mapLines[line] += 1
	}

	var res []string
	for key, value := range mapLines {
		if value == 1 {
			res = append(res, key)
		}
	}
	sort.Strings(res)
	return res
}

func writeResult(resData []string) {
	fileRes, e := os.Create("results.txt")

	if e != nil {
		fmt.Println("Error with create file: ", e)
		os.Exit(1)
	}

	for _, line := range resData {
		data := fmt.Sprintf("%s - %d байт\n", line, len(line))
		_, err := fileRes.WriteString(data)
		if err != nil {
			fmt.Println("Error with write to file: ", err)
			os.Exit(1)
		}
	}

	fileRes.Close()
}

func main() {
	var fileName string
	fmt.Println("Введите имя файла:")
	fmt.Fscan(os.Stdin, &fileName)
	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error with file: ", err)
		os.Exit(1)
	}

	fileDataStr := string(fileData)
	fileDataStr = strings.Replace(fileDataStr, "\r", "", -1)
	fileDataStr = strings.ToUpper(fileDataStr)

	resData := findUnical(strings.Split(fileDataStr, "\n"))

	writeResult(resData)

	fmt.Println("Результат в файле 'results.txt'")
}
