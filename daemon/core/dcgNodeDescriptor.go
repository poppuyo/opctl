package core

// descriptor for a DCG (dynamic call graph) node
type dcgNodeDescriptor struct {
	Id        string
	OpRef     string
	OpGraphId string
	Op        *dcgOpDescriptor
	Container *dcgContainerDescriptor
}

type dcgOpDescriptor struct{}

type dcgContainerDescriptor struct{}
