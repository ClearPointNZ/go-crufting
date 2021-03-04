package yaica

// EncruftEggs runs all of the eggs in the basket through the configured Crufter:
func (b *Basket) EncruftEggs() {

	// Encruft all the eggs:
	for _, egg := range b.eggs {
		encruftedEgg, err := b.crufter.Encruft(egg)
		if err != nil {
			b.logger.WithError(err).Errorf("Unable to encruft an egg (%s)", egg)
		}
		b.logger.WithField("egg", egg).WithField("encrufted", encruftedEgg).Infof("Encrufted an egg")
	}
}
