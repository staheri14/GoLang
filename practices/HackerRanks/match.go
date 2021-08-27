package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	temp := make(map[string]int)
	for _, data := range strings{
		v, _:= temp[data]
		temp[data] = v + 1
	}

	var result = make([]int32, len(queries))
	for i, data := range queries{
		value, _ := temp[data]
		result[i] = int32(value)
	}

	return result

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	stringsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var strings []string

	for i := 0; i < int(stringsCount); i++ {
		stringsItem := readLine(reader)
		strings = append(strings, stringsItem)
	}

	 str, _:=  readLine(reader)
	var cleanStr = strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")
	queriesCount, err := strconv.ParseInt(TrimSpace(str), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(strings, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}


