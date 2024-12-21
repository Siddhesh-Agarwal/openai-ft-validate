package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func acceptableLine(line string, line_no int) bool {
	var jsonMap map[string][]map[string]interface{}
	err := json.Unmarshal([]byte(line), &jsonMap)
	if err != nil {
		return false
	}

	messages, ok := jsonMap["messages"]
	if !ok {
		return false
	}

	for _, message := range messages {
		role, roleOk := message["role"].(string)          // required
		content, contentOk := message["content"].(string) // required
		weight, weightOk := message["weight"].(uint8)     // optional
		if !roleOk {
			fmt.Printf("Line %d: Missing required field \"role\"", line_no)
			return false
		}
		if !contentOk {
			fmt.Printf("Line %d: Missing required field \"content\"", line_no)
			return false
		}
		if weightOk {
			if weight != 0 && weight != 1 {
				fmt.Printf("Line %d: Invalid \"weight\" %d found\n", line_no, weight)
				return false
			}
		}
		if role != "assistant" && role != "system" && role != "user" {
			fmt.Printf("Line %d: Invalid \"role\" %s found\n", line_no, role)
			return false
		}
		if content == "" {
			fmt.Printf("Line %d: \"content\" is empty", line_no)
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the JSONL file.")
		return
	}

	jsonlFilePath := os.Args[1]

	file, err := os.Open(jsonlFilePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line_no := 1
	flag := true
	for scanner.Scan() {
		line := scanner.Text()
		if !acceptableLine(line, line_no) {
			flag = false
		}
		line_no++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	} else if flag {
		fmt.Println("The JSONL is valid 🎉")
	} else {
		fmt.Println("The JSONL is invalid ❌")
	}
}
