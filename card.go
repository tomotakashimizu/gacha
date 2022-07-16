package gacha

import "fmt"

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityXR Rarity = "XR"
)

type Card struct {
	Rarity Rarity // レア度
	Name   string // 名前
}

func (c *Card) String() string {
	return fmt.Sprintf("%v:%v", c.Rarity, c.Name)
}
