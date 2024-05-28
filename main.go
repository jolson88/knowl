package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jolson88/knowl/ideas"
)

func main() {
	activeId := uint(0)
	nextId := uint(0)
	ideaMap := make(map[uint]ideas.Idea)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("knowl> ")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				break
			}

			words := strings.SplitN(input, " ", 2)
			command := words[0]
			commandInput := ""
			if len(words) > 1 {
				commandInput = strings.Join(words[1:], " ")
			}

			if command == "new" {
				ideaMap[nextId] = ideas.Idea{
					Id:       nextId,
					Text:     strings.Trim(commandInput, " "),
					Children: []uint{},
				}
				activeId = nextId
				nextId++
			} else if command == "list" {
				for k, idea := range ideaMap {
					if k == activeId {
						fmt.Print("* ")
					}
					fmt.Printf("[%d] %s\n", k, idea.Text)
				}
			} else {
				fmt.Println("Unknown command:", command)
			}
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
