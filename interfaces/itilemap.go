package interfaces

type ITileMap[T any] interface {
	GetTransform() ITransform
	PlaceTile(localX, localY int, localRot, localLayer float64, tile T) ITransform
	RemoveTile(localX, localY int, localLayer float64) bool
}
