package ideas

import (
	"fmt"
	"strings"
)

var commands = struct {
	CreateChild string
	CreateNew   string
	MoveChild   string
}{
	CreateChild: "cc",
	CreateNew:   "cn",
	MoveChild:   "mc",
}

type IdeaId uint

type Idea struct {
	Id       IdeaId
	Text     string
	Children []IdeaId
}

type IdeaBank struct {
	NilIdea *Idea

	ideas      map[IdeaId]*Idea
	commandLog [][]byte
	nextId     IdeaId
}

func NewIdeaBank() *IdeaBank {
	nilIdea := &Idea{Id: 0, Text: "", Children: []IdeaId{}}

	return &IdeaBank{
		NilIdea: nilIdea,
		nextId:  1,
		ideas:   map[IdeaId]*Idea{0: nilIdea},
	}
}

func NewIdeaBankFromCommandLog(commandLog [][]byte) *IdeaBank {
	ideaBank := NewIdeaBank()
	for _, commandBytes := range commandLog {
		ideaBank.interpretCommand(commandBytes)
	}

	return ideaBank
}

func (ideaBank *IdeaBank) AllIdeas() []*Idea {
	ideas := make([]*Idea, 0, len(ideaBank.ideas))
	for _, idea := range ideaBank.ideas {
		ideas = append(ideas, idea)
	}
	return ideas
}

func (ideaBank *IdeaBank) CommandLog() [][]byte {
	return ideaBank.commandLog
}

func (ideaBank *IdeaBank) Count() int {
	return len(ideaBank.ideas)
}

func (ideaBank *IdeaBank) CreateChild(parentId IdeaId, text string) *Idea {
	parent := ideaBank.ideas[parentId]
	if parent == nil || parent.Id == 0 {
		return ideaBank.NilIdea
	}

	newIdea := &Idea{
		Id:       ideaBank.nextId,
		Text:     text,
		Children: []IdeaId{},
	}
	ideaBank.ideas[newIdea.Id] = newIdea
	parent.Children = append(parent.Children, newIdea.Id)

	ideaBank.nextId++
	ideaBank.commandLog = append(ideaBank.commandLog, []byte(fmt.Sprintf("%s %d %s", commands.CreateChild, parentId, text)))

	return newIdea
}

func (ideaBank *IdeaBank) CreateIdea(text string) *Idea {
	newIdea := &Idea{
		Id:       ideaBank.nextId,
		Text:     text,
		Children: []IdeaId{},
	}
	ideaBank.ideas[newIdea.Id] = newIdea

	ideaBank.nextId++
	ideaBank.commandLog = append(ideaBank.commandLog, []byte(fmt.Sprintf("%s %s", commands.CreateNew, text)))

	return newIdea
}

func (ideaBank *IdeaBank) interpretCommand(commandBytes []byte) {
	command := string(commandBytes)
	words := strings.SplitN(command, " ", 2)
	commandName := words[0]
	commandInput := ""
	if len(words) > 1 {
		commandInput = strings.Join(words[1:], " ")
	}

	switch commandName {
	case commands.CreateChild:
		var parentId IdeaId
		var text string
		fmt.Sscanf(commandInput, "%d %s", &parentId, &text)
		ideaBank.CreateChild(parentId, text)

	case commands.CreateNew:
		ideaBank.CreateIdea(commandInput)

	case commands.MoveChild:
		var parentId IdeaId
		var childIndex uint
		var offset int16
		fmt.Sscanf(commandInput, "%d %d %d", &parentId, &childIndex, &offset)
		ideaBank.MoveChild(parentId, childIndex, offset)

	default:
		fmt.Println("Unknown command:", commandName)
	}
}

func (ideaBank *IdeaBank) GetIdea(id IdeaId) *Idea {
	return ideaBank.ideas[id]
}

func (ideaBank *IdeaBank) MoveChild(parentId IdeaId, childIndex uint, offset int16) {
	// TODO: Change to using inline swapping to avoid extra memory copies
	// This is a super naive implementation since it rebuilds a fresh slice with the new order.
	// Perhaps we can count down from the offset to 0, swapping along the way. This should save
	// both time and memory.
	newIndex := int(childIndex) + int(offset)
	if newIndex < 0 {
		newIndex = 0
	}
	if newIndex >= len(ideaBank.ideas[parentId].Children) {
		newIndex = len(ideaBank.ideas[parentId].Children) - 1
	}

	parentIdea := ideaBank.ideas[parentId]
	newChildren := make([]IdeaId, 0, len(parentIdea.Children))
	for idx, childId := range parentIdea.Children {
		if idx == int(childIndex) {
			continue
		}
		if idx == newIndex {
			newChildren = append(newChildren, parentIdea.Children[childIndex])
		}
		newChildren = append(newChildren, childId)
	}
	parentIdea.Children = newChildren

	ideaBank.commandLog = append(ideaBank.commandLog, []byte(fmt.Sprintf("%s %d %d %d", commands.MoveChild, parentId, childIndex, offset)))
}
