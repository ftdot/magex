package tilemap

import (
	"math"

	"github.com/ftdot/magex/components/transform"
	"github.com/ftdot/magex/interfaces"
	"github.com/ftdot/magex/utils"
	"github.com/ftdot/magex/utils/vector2"
)

type TilePosition struct {
	X, Y  int
	Layer float64
}

// Primitive TileMap contains the square tiles. This
// makes work with tiles easier. Also, creates its
// own coordinate system depending on the size of one
// tile.
//
// This component does not need to be added to the GetComponents() function!
type TileMap[T any] struct {
	Transform      interfaces.ITransform
	tiles          map[TilePosition]T // All the tiles
	tileSizeScaled float64            // Scaled size of the tiles
	ID             string
}

func New[T any](tf interfaces.ITransform, tileSize float64) *TileMap[T] {
	return &TileMap[T]{
		Transform:      tf,
		tiles:          map[TilePosition]T{},
		tileSizeScaled: tileSize * tf.GetScale().X,
		ID:             utils.GenerateComponentID(),
	}
}

////

func (tm *TileMap[T]) GetTransform() interfaces.ITransform {
	return tm.Transform
}

////

// Places a tile at the given position. Returns the transform,
// that must be set to this tile.
func (tm *TileMap[T]) PlaceTile(localX, localY int, localRot, localLayer float64, tile T) interfaces.ITransform {
	t := transform.New(vector2.NewInt(localX, localY).MulScalar(tm.tileSizeScaled), vector2.Identity.Copy(), localRot, localLayer)
	t.SetParent(tm.Transform)

	tm.tiles[TilePosition{localX, localY, localLayer}] = tile

	return t
}

// Removes tile at the given position.
func (tm *TileMap[T]) RemoveTile(localX, localY int, localLayer float64) bool {
	tp := TilePosition{localX, localY, localLayer}
	_, ok := tm.tiles[tp]
	if !ok {
		return false
	}

	delete(tm.tiles, tp)

	return true
}

// Tries to get a tile with given position. If there is no
// tile, ok is equals to "false".
func (tm *TileMap[T]) TileAt(localX, localY int, localLayer float64) (tile T, ok bool) {
	tile, ok = tm.tiles[TilePosition{localX, localY, localLayer}]
	return
}

func (tm *TileMap[T]) OffsetVector2(point *vector2.Vector2) *vector2.Vector2 {
	return point.DivScalar(tm.tileSizeScaled)
}

func (tm *TileMap[T]) OffsetCoordinates(point *vector2.Vector2) (int, int) {
	return int(math.Floor(point.X / tm.tileSizeScaled)), int(math.Floor(point.Y / tm.tileSizeScaled))
}

////

func (tm *TileMap[T]) GetID() string {
	return tm.ID
}
