package otocards

import (
	"reflect"
	"testing"
)

func TestNewCard(t *testing.T) {
	front := "What is the capital of France?"
	back := "Paris"
	tags := []string{"Geography", "Europe"}

	card := NewCard(front, back, tags)

	if card.Front() != front {
		t.Errorf("Expected Front() to return %s, got %s", front, card.Front())
	}

	if card.Back() != back {
		t.Errorf("Expected Back() to return %s, got %s", back, card.Back())
	}

	if !reflect.DeepEqual(card.Tags(), tags) {
		t.Errorf("Expected Tags() to return %v, got %v", tags, card.Tags())
	}
}

func TestCardDeck(t *testing.T) {
	card1 := NewCard("Q1", "A1", []string{"Tag1"})
	card2 := NewCard("Q2", "A2", []string{"Tag2"})
	cards := []ICard{card1, card2}
	tags := []string{"Tag1", "Tag2"}

	deck := NewCardDeck(cards, tags)

	if deck.Size() != 2 {
		t.Errorf("Expected deck size to be 2, got %d", deck.Size())
	}

	if !reflect.DeepEqual(deck.Tags(), tags) {
		t.Errorf("Expected deck tags to be %v, got %v", tags, deck.Tags())
	}
}

func TestCardDeckFactory(t *testing.T) {
	factory := &CardDeckFactory{}
	data := []CardData{
		{Front: "Q1", Back: "A1", Tags: []string{"Tag1"}},
		{Front: "Q2", Back: "A2", Tags: []string{"Tag2"}},
	}

	deck := factory.MakeDeckFromCardData(data)

	if deck.Size() != 2 {
		t.Errorf("Expected deck size to be 2, got %d", deck.Size())
	}

	// Check if all tags are present in the deck
	expected := []string{"Tag1", "Tag2"}
	allTagsPresent := true
	for _, tag := range expected {
		found := false
		for _, deckTag := range deck.Tags() {
			if tag == deckTag {
				found = true
				break
			}
		}
		if !found {
			allTagsPresent = false
			break
		}
	}

	if !allTagsPresent {
		t.Errorf("Not all expected tags were found in deck tags: %v", deck.Tags())
	}
}

func TestCardDeck_Shuffle(t *testing.T) {
	card1 := NewCard("Q1", "A1", []string{"Tag1"})
	card2 := NewCard("Q2", "A2", []string{"Tag2"})
	card3 := NewCard("Q3", "A3", []string{"Tag3"})
	card4 := NewCard("Q4", "A4", []string{"Tag4"})
	card5 := NewCard("Q5", "A5", []string{"Tag5"})
	cards := []ICard{card1, card2, card3, card4, card5}
	tags := []string{"Tag1", "Tag2", "Tag3", "Tag4", "Tag5"}

	deck := NewCardDeck(cards, tags)

	originalOrder := make([]ICard, deck.Size())
	copy(originalOrder, deck.Cards())

	deck.Shuffle()

	if reflect.DeepEqual(originalOrder, deck.Cards()) {
		t.Error("Deck was not shuffled; original order and shuffled order are the same")
	}
}
