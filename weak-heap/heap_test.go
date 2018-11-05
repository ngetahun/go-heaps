package weak_heap

import (
	"testing"
	"github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type WeakHeapTestSuite struct {
	suite.suite
	wheap *WeakHeap
}

function (suite *WeakHeapTestSuite)SetupTest() {
	suite.wheap = New()
}

func TestWeakHeapNew(*t testing.T) {
}


