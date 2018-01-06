package compact

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"bitbucket.org/dnovikoff/tempai-core/tile"
)

func TestPackedMasks(t *testing.T) {
	packed := PackedMasks(0)
	assert.Equal(t, 0, packed.CountBits())
	packed.Set(NewMask(15, tile.Pin6), 1)
	// Unchanged
	assert.Equal(t, 0, packed.CountBits())
	packed = packed.Set(NewMask(15, tile.Pin6), 1)
	assert.Equal(t, 4, packed.CountBits())

	packed = packed.Set(NewMask(15, tile.Pin6), 0)
	assert.Equal(t, 8, packed.CountBits())

	packed = packed.Set(NewMask(2, tile.Pin6), 0)
	assert.Equal(t, 5, packed.CountBits())
}

func TestInstanceCountingOne(t *testing.T) {
	compact := NewInstances()
	assert.Equal(t, 0, compact.Count())
	mask := NewMask(15, tile.Man1)
	assert.Equal(t, 4, mask.Count())
	compact.SetMask(mask)
	assert.Equal(t, 4, compact.Count())

	compact.SetMask(NewMask(15, tile.Pin6))
	assert.Equal(t, 8, compact.Count())
}

func TestInstanceCounting(t *testing.T) {
	compact := NewInstances()
	assert.Equal(t, 0, compact.Count())
	compact.SetMask(NewMask(15, tile.Pin6))
	assert.Equal(t, 4, compact.Count())
	compact.SetMask(NewMask(2, tile.Pin6))
	assert.Equal(t, 1, compact.Count())

	compact.SetMask(NewMask(3, tile.Red))
	assert.Equal(t, 3, compact.Count())

	assert.Equal(t, "6p77z", compact.Instances().String())
	assert.Equal(t, "6p7z", compact.UniqueTiles().Tiles().String())
}

func TestInstanceMerge(t *testing.T) {
	first := NewInstances().AddCount(tile.Man1, 2)
	second := NewInstances().AddCount(tile.Sou8, 3)

	third := first.Merge(second)

	assert.Equal(t, "11m888s", third.Instances().String())
	assert.Equal(t, 5, third.Count())
}

func TestMaskError1(t *testing.T) {
	st := NewInstances()
	assert.Equal(t, "", st.Instances().String())
	st.Set(tile.Sou4.Instance(0))
	assert.Equal(t, 1, st.Count())
	assert.Equal(t, "4s", st.Instances().String())
}

func TestMaskErrors(t *testing.T) {
	tg := NewTileGenerator()
	str := "22223333444s55z"
	inst, err := tg.InstancesFromString(str)
	require.NoError(t, err)

	st := NewInstances()
	st.Add(inst)

	assert.Equal(t, len(inst), st.Count())
	assert.Equal(t, str, st.Instances().String())

	assert.Equal(t, len(inst), st.Count())
}

func TestInstancetCounters(t *testing.T) {
	st := NewInstances()
	assert.Equal(t, 0, st.GetCount(tile.Man3))
	assert.False(t, st.GetMask(tile.Man3).IsFull())
	assert.True(t, st.GetMask(tile.Man3).IsEmpty())

	st.SetCount(tile.Man3, 3)
	assert.Equal(t, 3, st.GetCount(tile.Man3))
	assert.False(t, st.GetMask(tile.Man3).IsFull())
	assert.False(t, st.GetMask(tile.Man3).IsEmpty())

	st.SetCount(tile.Man3, 4)
	assert.Equal(t, 4, st.GetCount(tile.Man3))
	assert.True(t, st.GetMask(tile.Man3).IsFull())
	assert.False(t, st.GetMask(tile.Man3).IsEmpty())
}