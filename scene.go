package mizu

type Scene interface {
	Setup(w World)
	//InitEntities(w World)
	//InitComponents(w World)
	//InitSystems(w World)
}
