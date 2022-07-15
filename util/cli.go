package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadStringFromCli(name string, output *string) error {
	fmt.Printf("Enter your %s:", name)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return err
	}
	input = strings.TrimSuffix(input,"\n")
	*output = input
	return nil
}