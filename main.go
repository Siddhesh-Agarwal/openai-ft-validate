package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func acceptableLine(line string, line_no int) bool {
	var rawMap map[string]interface{}
	if err := json.Unmarshal([]byte(line), &rawMap); err != nil {
		return false
	}

	// Check for unexpected top-level keys
	if len(rawMap) != 1 || rawMap["messages"] == nil {
		fmt.Printf("Line %d: Only \"messages\" key is allowed at top level\n", line_no)
		return false
	}

	messages, ok := rawMap["messages"].([]interface{})
	if !ok {
		fmt.Printf("Line %d: \"messages\" must be an array\n", line_no)
		return false
	}

	for _, msg := range messages {
		message, ok := msg.(map[string]interface{})
		if !ok {
			fmt.Printf("Line %d: Each message must be an object\n", line_no)
			return false
		}

		// Check for unexpected keys in message objects
		allowedKeys := map[string]bool{"role": true, "content": true, "weight": true}
		for key := range message {
			if !allowedKeys[key] {
				fmt.Printf("Line %d: Unexpected key \"%s\" in message object\n", line_no, key)
				return false
			}
		}

		role, roleOk := message["role"].(string)
		content, contentOk := message["content"].(string)
		weight, weightOk := message["weight"].(float64)

		if !roleOk {
			fmt.Printf("Line %d: Missing required field \"role\"\n", line_no)
			return false
		}
		if !contentOk {
			fmt.Printf("Line %d: Missing required field \"content\"\n", line_no)
			return false
		}
		if weightOk {
			if weight != 0 && weight != 1 {
				fmt.Printf("Line %d: Invalid \"weight\" %f found\n", line_no, weight)
				return false
			}
		}
		if role != "assistant" && role != "system" && role != "user" {
			fmt.Printf("Line %d: Invalid \"role\" %s found\n", line_no, role)
			return false
		}
		if content == "" {
			fmt.Printf("Line %d: \"content\" is empty\n", line_no)
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
