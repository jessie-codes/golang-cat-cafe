package cat

import (
	"sync"
	"time"
)

//Cat stores the attributes of a single cat.
type Cat struct {
	Name        string `json:"name" xml:"name"`
	Breed       string `json:"breed" xml:"breed"`
	Personality string `json:"personality" xml:"personality"`
	Available   bool   `json:"available" xml:"available"`
}

//Cats is a slice of Cat pointers
type Cats struct {
	List []*Cat `json:"cats" xml:"list"`
	mux  sync.Mutex
}

//List defines actions that can be taken on cats
type List interface {
	GetByPersonality(personality string) *Cats
	Reserve() *Cat
}

//Get returns a list of all cats
func Get() List {
	cats := Cats{
		List: CafeCats,
	}
	return &cats
}

//GetByPersonality returns a filtered list of cats by personality
func (c *Cats) GetByPersonality(personality string) *Cats {
	result := Cats{}
	for _, cat := range c.List {
		if cat.Personality == personality {
			result.List = append(result.List, cat)
		}
	}
	return &result
}

//Reserve returns the first available cat
func (c *Cats) Reserve() *Cat {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, cat := range c.List {
		if cat.Available {
			go cat.startAppointment()
			return cat
		}
	}
	return nil
}

func (c *Cat) startAppointment() {
	c.Available = false
	time.Sleep(5 * time.Second)
	c.Available = true
}
