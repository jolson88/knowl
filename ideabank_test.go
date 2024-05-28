package main

import (
	"testing"

	"github.com/jolson88/knowl/ideas"
)

func TestCreatesAndInteractsWithNewIdea(t *testing.T) {
	ideaBank := ideas.NewIdeaBank()
	var countBefore = ideaBank.Count()

	//
	// Create new ideas
	//
	var firstIdea = ideaBank.CreateIdea("first idea")
	var secondIdea = ideaBank.CreateIdea("second idea")

	if firstIdea.Id == secondIdea.Id {
		t.Fatalf("Expected different ids, but got %d and %d", firstIdea.Id, secondIdea.Id)
	}
	if ideaBank.Count() != countBefore+2 {
		t.Fatalf("Expected %d ideas, got %d", countBefore+2, ideaBank.Count())
	}
	if ideaBank.ActiveIdea == nil {
		t.Fatalf("Expected ActiveIdea to not be nil")
	}
	if ideaBank.ActiveIdea.Text != "second idea" {
		t.Fatalf("Expected ActiveIdea to have text 'second idea', got '%s'", ideaBank.ActiveIdea.Text)
	}

	//
	// Append Children
	//
	countBefore = ideaBank.Count()

	const firstChildText = "first child"
	const secondChildText = "second child"
	ideaBank.AppendChild(firstChildText)
	ideaBank.AppendChild(secondChildText)

	var allIdeas = ideaBank.GetAllIdeas()
	if len(allIdeas) != countBefore+2 {
		t.Fatalf("Expected children to be added for %d total ideas, got %d", countBefore+2, len(allIdeas))
	}
	if ideaBank.ActiveIdea.Text != "second idea" {
		t.Fatalf("Expected ActiveIdea to not change when children are appended, got '%s'", ideaBank.ActiveIdea.Text)
	}
	if len(ideaBank.ActiveIdea.Children) != 2 {
		t.Fatalf("Expected ActiveIdea to have 2 children, got %d", len(ideaBank.ActiveIdea.Children))
	}

	var firstChild = ideaBank.GetIdea(ideaBank.ActiveIdea.Children[0])
	var secondChild = ideaBank.GetIdea(ideaBank.ActiveIdea.Children[1])
	if firstChild.Text != firstChildText {
		t.Fatalf("Expected first child to have text '%s', got '%s'", firstChildText, firstChild.Text)
	}
	if secondChild.Text != secondChildText {
		t.Fatalf("Expected second child to have text '%s', got '%s'", secondChildText, secondChild.Text)
	}

	//
	// Re-ordering
	//
	ideaBank.SwapChildren(0, 1)

	var newFirstChild = ideaBank.GetIdea(ideaBank.ActiveIdea.Children[0])
	var newSecondChild = ideaBank.GetIdea(ideaBank.ActiveIdea.Children[1])
	if newFirstChild.Text != secondChildText {
		t.Fatalf("Expected first child to have text '%s', got '%s'", firstChildText, firstChild.Text)
	}
	if newSecondChild.Text != firstChildText {
		t.Fatalf("Expected second child to have text '%s', got '%s'", secondChildText, secondChild.Text)
	}

	//
	// Idea Activation
	//
	ideaBank.SetActiveIdea(firstIdea.Id)

	if ideaBank.ActiveIdea.Text != "first idea" {
		t.Fatalf("Expected ActiveIdea to be 'first idea', got '%s'", ideaBank.ActiveIdea.Text)
	}
}
