package yaica

// TestReplaceEggs uses the dependencies established by the test suite to test the ReplaceEggs method:
func (s *yaicaTestSuite) TestReplaceEggs() {

	// Remember how many times that Encruft had been called before we started:
	originalCallCount := s.fakeCrufter.EncruftCallCount()

	// Configure our mock to return something:
	s.fakeCrufter.EncruftReturns("encrufted egg", nil)

	// Provide some new eggs:
	s.basket.SetEggs([]string{"this is egg1", "this is egg2", "this is egg3"})
	s.Len(s.basket.eggs, 3)

	// Replace the eggs:
	s.basket.ReplaceEggs()

	// Make sure our fake crufter was called three times (once per egg):
	s.Equal(originalCallCount+3, s.fakeCrufter.EncruftCallCount())

	// Make sure the eggs were corrected encrufted (according to our mock):
	s.Equal([]string{"encrufted egg", "encrufted egg", "encrufted egg"}, s.basket.eggs)
}
