package cat

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CatListTestSuite struct {
	suite.Suite
	TestCats *Cats
}

func (suite *CatListTestSuite) SetupTest() {
	suite.TestCats = &Cats{}
	err := json.Unmarshal(cafeCats, &suite.TestCats.List)
	if err != nil {
		panic("Unable to load cats")
	}
}

func (suite *CatListTestSuite) TestGetByPersonality() {
	t := suite.T()
	result := suite.TestCats.GetByPersonality("diva")
	assert.Equal(t, 2, len(result.List))
	result = suite.TestCats.GetByPersonality("shy")
	assert.Equal(t, 1, len(result.List))
	result = suite.TestCats.GetByPersonality("mean")
	assert.Equal(t, 0, len(result.List))
}

func TestCatListTestSuite(t *testing.T) {
	suite.Run(t, new(CatListTestSuite))
}
