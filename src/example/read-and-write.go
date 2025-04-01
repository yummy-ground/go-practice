package example

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLineFrom(origin string, lines *[]string) error {
	scanner := bufio.NewScanner(strings.NewReader(origin)) // NewReader -> io.Reader 계열
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func readFileFrom(path string, lines *[]string) error {
	fileReader, err := os.Open(path)
	if err != nil {
		return err
	}

	defer fileReader.Close() // 수행 순서는 역순 -> 아래 코드가 끝난 다음에 실행됨
	var line string
	if _, err := fmt.Fscanf(fileReader, "%s\n", &line); err == nil {
		*lines = append(*lines, line)
	}
	return nil
}

func writeLineFrom(lines []string) error {
	writer := os.Stdout
	for _, line := range lines {
		if _, err := fmt.Fprintln(writer, line); err != nil {
			return err
		}

	}
	return nil
}

func writeFile(path string, lines []string) error {
	fileWriter, err := os.Create(path)
	if err != nil {
		return err
	}

	defer fileWriter.Close() // 수행 순서는 역순 -> 아래 코드가 끝난 다음에 실행됨
	for _, line := range lines {
		if _, err := fmt.Fprintf(fileWriter, "%s\n", line); err == nil {
			return err
		}
	}
	return nil

}
