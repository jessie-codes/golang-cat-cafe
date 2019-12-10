package cat

import (
	"encoding/json"
	"time"
)

//Cat stores the attributes of a single cat.
type Cat struct {
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Personality string `json:"personality"`
	Available   bool   `json:"available"`
}

//Cats is a slice of Cat pointers
type Cats struct {
	List []*Cat `json:"cats"`
}

//List defines actions that can be taken on cats
type List interface {
	GetByPersonality(personality string) *Cats
	Reserve() *Cat
}

//Get returns a list of all cats
func Get() List {
	cats := Cats{}
	err := json.Unmarshal(cafeCats, &cats.List)
	if err != nil {
		panic("Unable to load cats")
	}
	return &cats
}

//GetByPersonality returns a filtered list of cats by personality
func (c *Cats) GetByPersonality(personality string) *Cats {
	result := Cats{}
	return &result
}

//Reserve returns the first available cat and sets them to unavailable
func (c *Cats) Reserve() *Cat {
	for _, cat := range c.List {
		if cat.Available {
			cat.Available = false
			time.Sleep(5 * time.Second)
			cat.Available = true
			return cat
		}
	}
	return nil
}
