package builder

type BuildManager struct {
	builder BuilderInterface
}

func newBuildManager(b BuilderInterface) *BuildManager {
	return &BuildManager{
		builder: b,
	}
}
func (b *BuildManager) createBuild() Build {
	b.builder.setHeight()
	b.builder.setLength()
	b.builder.setWidth()
	b.builder.setName()
	return b.builder.getBuild()
}
