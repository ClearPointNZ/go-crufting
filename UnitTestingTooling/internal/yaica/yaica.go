package yaica

import (
	"fmt"

	"github.com/ClearPointNZ/go-crufting/UnitTestingTooling/internal/cruft"
	"github.com/sirupsen/logrus"
)

// Basket is a thing with some dependencies:
type Basket struct {
	eggs    []string
	crufter cruft.Crufter
	logger  *logrus.Logger
}

// New returns a configured egg basket:
func New(logger *logrus.Logger, crufter cruft.Crufter) (*Basket, error) {
	if crufter == nil || logger == nil {
		return nil, fmt.Errorf("Please don't give me nil inputs")
	}

	return &Basket{
		crufter: crufter,
		logger:  logger,
	}, nil
}

// SetEggs configures new eggs:
func (b *Basket) SetEggs(newEggs []string) {
	b.eggs = newEggs
}
