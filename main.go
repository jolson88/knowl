package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jolson88/knowl/ideas"
)

func main() {
	ideaBank := ideas.NewIdeaBank()
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

			switch command {
			case "log":
				for _, command := range ideaBank.CommandLog() {
					fmt.Println(string(command))
				}
			case "ls":
				var visited = make(map[uint]bool)
				for _, idea := range ideaBank.GetAllIdeas() {
					if idea.Id == ideaBank.NilIdea.Id {
						continue
					}
					if visited[idea.Id] {
						continue
					}

					if idea.Id == ideaBank.ActiveIdea.Id {
						fmt.Printf("*[%d] %s\n", idea.Id, idea.Text)
					} else {
						fmt.Printf("[%d] %s\n", idea.Id, idea.Text)
					}
					visited[idea.Id] = true

					for _, childId := range idea.Children {
						fmt.Printf("    - [%d] %s\n", childId, ideaBank.GetIdea(childId).Text)
						visited[childId] = true
					}
				}
			default:
				ideaBank.InterpretCommand([]byte(input))
			}
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
