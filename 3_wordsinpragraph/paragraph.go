package __wordsinpragraph

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindLongestWords() (int, string) {
	reader := bufio.NewReader(os.Stdin)
	strs, _, err := reader.ReadLine()
	if err != nil {
		return 0, ""
	}
	splited := strings.Split(string(strs), " ")

	maxLength := 0
	maxWords := ""
	for _, str := range splited {
		fmt.Println(str, len(str))
		if len(str) > maxLength {
			maxLength = len(str)
			maxWords = str
		}
	}
	return maxLength, maxWords
}
