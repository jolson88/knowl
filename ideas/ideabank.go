package ideas

import (
	"fmt"
	"strings"
)

var commands = struct {
	AppendChild  string
	CreateNew    string
	SetActive    string
	SwapChildren string
}{
	AppendChild:  "ac",
	CreateNew:    "cn",
	SetActive:    "sa",
	SwapChildren: "sc",
}

type Idea struct {
	Id       uint
	Text     string
	Children []uint
}

type IdeaBank struct {
	ideas      []*Idea
	commandLog [][]byte
	nextId     uint

	ActiveIdea *Idea
	NilIdea    *Idea
}

func NewIdeaBank() *IdeaBank {
	var nilIdea = Idea{Id: 0, Text: "", Children: []uint{}}

	return &IdeaBank{
		NilIdea:    &nilIdea,
		ActiveIdea: &nilIdea,
		nextId:     1,
		ideas:      []*Idea{&nilIdea},
	}
}

func NewIdeaBankFromCommandLog(commandLog [][]byte) *IdeaBank {
	var ideaBank = NewIdeaBank()
	for _, commandBytes := range commandLog {
		ideaBank.InterpretCommand(commandBytes)
	}

	return ideaBank
}

func (ideaBank *IdeaBank) AppendChild(text string) {
	ideaBank.ideas = append(ideaBank.ideas, &Idea{
		Id:       ideaBank.nextId,
		Text:     text,
		Children: []uint{},
	})
	ideaBank.ActiveIdea.Children = append(ideaBank.ActiveIdea.Children, ideaBank.nextId)
	ideaBank.nextId++
	ideaBank.commandLog = append(ideaBank.commandLog, []byte(fmt.Sprintf("%s %s", commands.AppendChild, text)))
}

func (ideaBank *IdeaBank) Count() int {
	return len(ideaBank.ideas)
}

func (ideaBank *IdeaBank) CreateIdea(text string) *Idea {
	var newIdea = &Idea{
		Id:       ideaBank.nextId,
		Text:     text,
		Children: []uint{},
	}
	ideaBank.ActiveIdea = newIdea
	ideaBank.ideas = append(ideaBank.ideas, newIdea)
	ideaBank.nextId++

	ideaBank.commandLog = append(ideaBank.commandLog, []byte(fmt.Sprintf("%s %s", commands.CreateNew, text)))
	return newIdea
}

func (ideaBank *IdeaBank) InterpretCommand(commandBytes []byte) {
	var command = string(commandBytes)
	var words = strings.SplitN(command, " ", 2)
	var commandName = words[0]
	var commandInput = ""
	if len(words) > 1 {
		commandInput = strings.Join(words[1:], " ")
	}

	switch commandName {
	case commands.AppendChild:
		ideaBank.AppendChild(commandInput)
	case commands.CreateNew:
		ideaBank.CreateIdea(commandInput)
	case commands.SetActive:
		var id uint
		fmt.Sscanf(commandInput, "%d", &id)
		ideaBank.SetActiveIdea(id)
	case commands.SwapChildren:
		var firstChildIndex, secondChildIndex uint
		fmt.Sscanf(commandInput, "%d %d", &firstChildIndex, &secondChildIndex)
		ideaBank.SwapChildren(firstChildIndex, secondChildIndex)
	default:
		fmt.Println("Unknown command:", commandName)
	}
}

func (ideaBank *IdeaBank) GetAllIdeas() []*Idea {
	return ideaBank.ideas[0:]
}

func (ideaBank *IdeaBank) GetIdea(id uint) *Idea {
	if id >= uint(len(ideaBank.ideas)) {
		return ideaBank.NilIdea
	}
	return ideaBank.ideas[id]
}

func (ideaBank *IdeaBank) CommandLog() [][]byte {
	return ideaBank.commandLog
}

func (ideaBank *IdeaBank) SetActiveIdea(id uint) {
	if id >= uint(len(ideaBank.ideas)) {
		ideaBank.ActiveIdea = ideaBank.NilIdea
		return
	}
	ideaBank.ActiveIdea = ideaBank.ideas[id]
	ideaBank.commandLog = append(ideaBank.commandLog, []byte(fmt.Sprintf("%s %d", commands.SetActive, id)))
}

func (ideaBank *IdeaBank) SwapChildren(firstChildIndex uint, secondChildIndex uint) {
	var childrenCount = uint(len(ideaBank.ActiveIdea.Children))
	if firstChildIndex >= childrenCount || secondChildIndex >= childrenCount {
		return
	}

	var firstChildId = ideaBank.ActiveIdea.Children[firstChildIndex]
	var secondChildId = ideaBank.ActiveIdea.Children[secondChildIndex]
	ideaBank.ActiveIdea.Children[firstChildIndex] = secondChildId
	ideaBank.ActiveIdea.Children[secondChildIndex] = firstChildId

	ideaBank.commandLog = append(ideaBank.commandLog, []byte(fmt.Sprintf("%s %d %d", commands.SwapChildren, firstChildIndex, secondChildIndex)))
}
