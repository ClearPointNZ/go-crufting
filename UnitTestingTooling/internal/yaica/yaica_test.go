package yaica

import (
	"bytes"
	"testing"

	"github.com/ClearPointNZ/go-crufting/UnitTestingTooling/internal/cruft/mock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

// yaicaTestSuite provides everything we need to test an egg basket:
type yaicaTestSuite struct {
	suite.Suite
	basket      *Basket
	fakeCrufter *mock.FakeCrufter
	logBuffer   *bytes.Buffer
}

// SetupTest puts a basket together using a mock Crufter:
// - This is automatically called when we do "suite.Run()"
func (s *yaicaTestSuite) SetupTest() {

	// Use a FakeStore mock:
	fakeCrufter := new(mock.FakeCrufter)
	s.fakeCrufter = fakeCrufter

	// Get a new logger:
	logger := logrus.New()
	s.logBuffer = bytes.NewBufferString("")
	logger.SetOutput(s.logBuffer)
	logger.SetLevel(logrus.TraceLevel)

	// Put these into a new Basket:
	basket, err := New(logger, fakeCrufter)
	s.NoError(err)

	// Now make this configured basket available to the test suite:
	s.basket = basket
}

// TestYaica initiates the entire test-suite:
func TestYaica(t *testing.T) {
	suite.Run(t, new(yaicaTestSuite))
}
