package yaica

// ReplaceEggs runs all of the eggs in the basket through the configured Crufter, and replaces their original values:
func (b *Basket) ReplaceEggs() {

	// Encruft all the eggs:
	for i, egg := range b.eggs {
		encruftedEgg, err := b.crufter.Encruft(egg)
		if err != nil {
			b.logger.WithError(err).Errorf("Unable to encruft an egg (%s)", egg)
			continue
		}

		b.eggs[i] = encruftedEgg
	}
}
