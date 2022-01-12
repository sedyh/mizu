package enum

type BuildMode int

const (
	BuildModeNone = BuildMode(iota)
	BuildModeDestroy
	BuildModeBuild
)
