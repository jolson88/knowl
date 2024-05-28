package ideas

type Idea struct {
	Id       uint
	Text     string
	Children []uint
}

type IdeaBank struct {
	NilIdea    *Idea
	ActiveIdea *Idea
	nextId     uint
	ideas      []Idea
}

func NewIdeaBank() IdeaBank {
	var nilIdea = Idea{Id: 0, Text: "", Children: []uint{}}

	return IdeaBank{
		NilIdea:    &nilIdea,
		ActiveIdea: &nilIdea,
		nextId:     1,
		ideas:      []Idea{nilIdea},
	}
}

func (ideaBank *IdeaBank) Count() int {
	return len(ideaBank.ideas)
}

func (ideaBank *IdeaBank) CreateNewIdea(text string) *Idea {
	var newIdea = Idea{
		Id:       ideaBank.nextId,
		Text:     text,
		Children: []uint{},
	}
	ideaBank.ideas = append(ideaBank.ideas, newIdea)
	ideaBank.ActiveIdea = &ideaBank.ideas[len(ideaBank.ideas)-1]
	ideaBank.nextId++
	return &newIdea
}

func (ideaBank *IdeaBank) GetIdea(id uint) *Idea {
	if id >= uint(len(ideaBank.ideas)) {
		return ideaBank.NilIdea
	}
	return &ideaBank.ideas[id]
}

func (ideaBank *IdeaBank) AppendChild(text string) {
	ideaBank.ideas = append(ideaBank.ideas, Idea{
		Id:       ideaBank.nextId,
		Text:     text,
		Children: []uint{},
	})
	ideaBank.ActiveIdea.Children = append(ideaBank.ActiveIdea.Children, ideaBank.nextId)
	ideaBank.nextId++
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
}
