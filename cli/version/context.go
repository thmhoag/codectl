package version

type Ctx interface {
	CurrentVersion() *Properties
}