package treedisplay

type DisplayTreeConfig struct {
	Type       DisplayType
	Characters CharacterType
	Highlight  interface{}
}

func NewDisplayConfig() DisplayTreeConfig {
	return DisplayTreeConfig{Type: Balanced, Characters: Unicode, Highlight: nil}
}

func (config DisplayTreeConfig) WithDisplayType(value DisplayType) DisplayTreeConfig {
	return DisplayTreeConfig{Type: value, Characters: config.Characters, Highlight: config.Highlight}
}

func (config DisplayTreeConfig) WithCharacterType(value CharacterType) DisplayTreeConfig {
	return DisplayTreeConfig{Type: config.Type, Characters: value, Highlight: config.Highlight}
}

func (config DisplayTreeConfig) WithHighlight(value interface{}) DisplayTreeConfig {
	return DisplayTreeConfig{Type: config.Type, Characters: config.Characters, Highlight: value}
}

type DisplayType int
type CharacterType int

const (
	TopDown           DisplayType   = iota
	Balanced          DisplayType   = iota
	BalancedFavourTop DisplayType   = iota
	BottomUp          DisplayType   = iota
	Unicode           CharacterType = iota
)

