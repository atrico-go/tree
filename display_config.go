package tree

type DisplayTreeConfig struct {
	Type       DisplayType
	Characters CharacterType
}

func NewDisplayConfig() DisplayTreeConfig {
	return DisplayTreeConfig{Type: Balanced, Characters: Unicode}
}

func (config DisplayTreeConfig) WithDisplayType(value DisplayType) DisplayTreeConfig {
	return DisplayTreeConfig{Type: value, Characters: config.Characters}
}

func (config DisplayTreeConfig) WithCharacterType(value CharacterType) DisplayTreeConfig {
	return DisplayTreeConfig{Type: config.Type, Characters: value}
}

type DisplayType int
type CharacterType int

const (
	TopDown           DisplayType   = iota
	Balanced          DisplayType   = iota
	BalancedFavourTop DisplayType   = iota
	BottomUp          DisplayType   = iota
	ASCII             CharacterType = iota
	Unicode           CharacterType = iota
)

