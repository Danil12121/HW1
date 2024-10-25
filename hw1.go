package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

//Напиши программу на языке Go, которая:

//1. Откроет указанный текстовый файл.
//2. Найдёт все строки, встречающиеся ровно один раз.
//3. Преобразует каждую уникальную строку в верхний регистр.
//4. Посчитает количество байт, которые занимает каждая уникальная строка.
//5. Запишет результат в новый файл в формате, каждая строка в файле-выводе должна быть отсортирована по алфавиту:
//УНИКАЛЬНАЯ СТРОКА В ВЕРХНЕМ РЕГИСТРЕ - X байт
//Правила:
//Каждая строка начинается с новой строки.
//Возможна подача на вход пустого и несуществующего файлов

func findUnical(lines []string) ([]string, error) {
	mapLines := make(map[string]int, len(lines))
	for _, line := range lines {
		mapLines[line] += 1
	}
	var res []string
	//for _, key := range mapLines.keys()
	for key, value := range mapLines {
		if value == 1 {
			res = append(res, key)
		}
	}
	sort.Strings(res)
	return res, nil
}

func writeResult(resData []string) error {
	fileRes, e := os.Create("results.txt")

	if e != nil {
		return fmt.Errorf("Error with create file: %w", e)
	}

	for _, line := range resData {
		data := fmt.Sprintf("%s - %d байт\n", line, len(line))
		_, err := fileRes.WriteString(data)
		if err != nil {
			return fmt.Errorf("Error with write to file: %w", err)
		}
	}

	fileRes.Close()
	return nil
}

func main() {

	if len(os.Args) < 2 {
		panic(errors.New("Wrong name of file"))
	}
	fileName := os.Args[1]

	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(errors.New("Error with read file"))
	}

	fileDataStr := string(fileData)
	fileDataStr = strings.Replace(fileDataStr, "\r", "", -1)
	fileDataStr = strings.ToUpper(fileDataStr)

	resData, _ := findUnical(strings.Split(fileDataStr, "\n"))

	e := writeResult(resData)
	if e != nil {
		panic(e)
	}

	fmt.Println("Результат в файле 'results.txt'")
}
