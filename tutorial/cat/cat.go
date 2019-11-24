package cat

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
	return &result
}
