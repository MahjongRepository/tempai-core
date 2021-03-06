package yaku_test

import (
	"fmt"
	"log"

	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/hand/tempai"
	"github.com/dnovikoff/tempai-core/tile"
	"github.com/dnovikoff/tempai-core/yaku"
)

func ExampleWin() {
	generator := compact.NewTileGenerator()
	tiles, err := generator.CompactFromString("33z123m456p66778s")
	if err != nil {
		log.Fatal(err)
	}
	winTile := generator.Instance(tile.Sou5)

	results := tempai.Calculate(tiles)
	ctx := &yaku.Context{
		Tile:      winTile,
		Rules:     yaku.RulesEMA(),
		IsTsumo:   true,
		IsChankan: true,
	}
	yakuResult := yaku.Win(results, ctx, nil)
	fmt.Printf("%v\n", yakuResult.Yaku.String())
	fmt.Printf("Value: %v.%v\n", yakuResult.Sum(), yakuResult.Fus.Sum())
	// Output:
	// YakuChankan: 1, YakuPinfu: 1, YakuTsumo: 1
	// Value: 3.20
}
