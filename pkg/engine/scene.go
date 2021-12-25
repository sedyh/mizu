package engine

// Scene represents the state of the game with a specific set of components, entities and systems.
// Scenes made in order to implement different game screens: world, pause, settings, menu.
type Scene interface {
	Setup(w World)
}
