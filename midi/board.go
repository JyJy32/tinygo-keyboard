package midi

type Board struct {
	buttons []MidiButton
}

func (b *Board) OnTick() error {
	for _, button := range b.buttons {
		err := button.OnTick()
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Board) AddButton(button MidiButton) {
	b.buttons = append(b.buttons, button)
}
