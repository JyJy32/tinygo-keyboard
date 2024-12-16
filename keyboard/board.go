package keyboard

type Board struct {
	buttons []*Button
}

func (b *Board) AddButton(btn *Button) *Board {
	b.buttons = append(b.buttons, btn)
	return b
}

func (b *Board) OnTick() error {
	for _, btn := range b.buttons {
		err := btn.OnTick()
		if err != nil {
			return err
		}
	}
	return nil
}
