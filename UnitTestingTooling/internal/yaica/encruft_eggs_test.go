package yaica

// TestEncruftEggs uses the dependencies established by the test suite to test the EncruftEggs method:
func (s *yaicaTestSuite) TestEncruftEggs() {

	// Remember how many times that Encruft had been called before we started:
	originalCallCount := s.fakeCrufter.EncruftCallCount()

	// Set some default eggs to start with:
	s.basket.SetEggs([]string{"this is original egg1", "this is original egg2"})

	// Configure our mock to return something:
	s.fakeCrufter.EncruftReturns("encrufted egg", nil)

	// Remember what the eggs were:
	existingEggs := s.basket.eggs

	// Reset the logBuffer so we know that it will be empty:
	s.logBuffer.Reset()

	// Replace the eggs:
	s.basket.EncruftEggs()

	// Make sure our fake crufter was called twice (once per egg):
	s.Equal(originalCallCount+2, s.fakeCrufter.EncruftCallCount())
	s.Contains(s.logBuffer.String(), "Encrufted an egg")
	s.Contains(s.logBuffer.String(), "this is original egg1")
	s.Contains(s.logBuffer.String(), "this is original egg2")
	s.Contains(s.logBuffer.String(), "encrufted egg")

	// Make sure the eggs were not modified:
	s.Equal(existingEggs, s.basket.eggs)
}
