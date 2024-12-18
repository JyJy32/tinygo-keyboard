package midi

type Board struct {
	buttons []*MidiControlButton
}

func (b *Board) AddButton(button *MidiControlButton) *Board {
	b.buttons = append(b.buttons, button)
	return b
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
