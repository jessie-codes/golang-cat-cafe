package cat

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CatListTestSuite struct {
	suite.Suite
	TestCats *Cats
}

func (suite *CatListTestSuite) SetupTest() {
	suite.TestCats = &Cats{
		List: CafeCats,
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

// func (suite *CatListTestSuite) TestReserve() {
// 	t := suite.T()
// 	for i := 0; i < len(suite.TestCats.List); i++ {
// 		result := suite.TestCats.Reserve()
// 		assert.NotNil(t, result)
// 	}
// 	result := suite.TestCats.Reserve()
// 	assert.Nil(t, result)
// }

func TestCatListTestSuite(t *testing.T) {
	suite.Run(t, new(CatListTestSuite))
}
