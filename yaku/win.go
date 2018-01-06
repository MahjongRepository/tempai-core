package yaku

import (
	"bitbucket.org/dnovikoff/tempai-core/meld"
	"bitbucket.org/dnovikoff/tempai-core/shanten"
)

func Win(tempai shanten.IndexedTempaiResult, ctx *Context) (ret *YakuResult) {
	current := tempai[ctx.Tile.Tile()]

	if len(current) == 0 {
		return nil
	}
	top := 0
	for _, v := range current {
		waiting := append(meld.Melds{}, v...)
		winMeld := waiting.Win(ctx.Tile.Tile())
		if winMeld.IsNull() {
			return nil
		}
		calc := NewYakuCalculator(ctx, waiting).Calculate()
		if calc == nil {
			return nil
		}

		if len(calc.Yakuman) > 0 {
			top = 14
			return calc
		}
		sum := int(calc.Fus.Sum()) + int(calc.Yaku.Sum()*1000)
		if sum > top && sum > 1000 {
			ret = calc
			top = sum
		}
	}
	return
}