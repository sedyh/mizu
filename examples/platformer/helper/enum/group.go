package enum

type CollisionGroup int

const (
	CollisionGroupNone = CollisionGroup(iota)
	CollisionGroupPlayer
	CollisionGroupTile
	CollisionGroupCrate
)
