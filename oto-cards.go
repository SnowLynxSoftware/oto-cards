package otocards

import "math/rand/v2"

type CardData struct {
	Front string   `json:"front"`
	Back  string   `json:"back"`
	Tags  []string `json:"tags"`
}

type ICard interface {
	Front() string
	Back() string
	Tags() []string
}

type Card struct {
	front string
	back  string
	tags  []string
}

func NewCard(front, back string, tags []string) ICard {
	return &Card{
		front: front,
		back:  back,
		tags:  tags,
	}
}

func (c *Card) Front() string {
	return c.front
}

func (c *Card) Back() string {
	return c.back
}

func (c *Card) Tags() []string {
	return c.tags
}

type ICardDeck interface {
	Cards() []ICard
	Shuffle()
	Size() int
	Tags() []string
}

type CardDeck struct {
	cards []ICard
	tags  []string
}

func NewCardDeck(cards []ICard, tags []string) ICardDeck {
	return &CardDeck{
		cards: cards,
		tags:  tags,
	}
}

func (d *CardDeck) Cards() []ICard {
	return d.cards
}

func (d *CardDeck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *CardDeck) Size() int {
	return len(d.cards)
}

func (d *CardDeck) Tags() []string {
	return d.tags
}

type ICardDeckFactory interface {
	MakeDeck() ICardDeck
}

type CardDeckFactory struct {
}

func (f *CardDeckFactory) MakeDeckFromCardData(data []CardData) ICardDeck {
	var cards []ICard
	var allTags []string
	for _, d := range data {
		cards = append(cards, NewCard(d.Front, d.Back, d.Tags))
		allTags = append(allTags, d.Tags...)
	}

	return NewCardDeck(cards, allTags)
}
