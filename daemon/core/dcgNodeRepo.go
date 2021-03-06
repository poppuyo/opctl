package core

//go:generate counterfeiter -o ./fakeDcgNodeRepo.go --fake-name fakeDcgNodeRepo ./ dcgNodeRepo

import (
	"sync"
)

// stores DCG (dynamic call graph) nodes
type dcgNodeRepo interface {
	// adds the provided node
	Add(node *dcgNodeDescriptor)
	// deletes the node with the provided id
	DeleteIfExists(nodeId string)
	// lists all nodes with the provided opGraphId
	ListWithOpGraphId(opGraphId string) []*dcgNodeDescriptor
	// tries to get the node with the provided id; returns nil if not found
	GetIfExists(nodeId string) *dcgNodeDescriptor
}

func newDcgNodeRepo() dcgNodeRepo {

	return &_dcgNodeRepo{
		byIdIndex:      make(map[string]*dcgNodeDescriptor),
		byIdIndexMutex: sync.RWMutex{},
	}

}

type _dcgNodeRepo struct {
	// synchronize access via mutex
	byIdIndex      map[string]*dcgNodeDescriptor
	byIdIndexMutex sync.RWMutex
}

// O(1) complexity; thread safe
func (this *_dcgNodeRepo) Add(node *dcgNodeDescriptor) {
	this.byIdIndexMutex.Lock()
	defer this.byIdIndexMutex.Unlock()
	this.byIdIndex[node.Id] = node
}

// O(1) complexity; thread safe
func (this *_dcgNodeRepo) DeleteIfExists(nodeId string) {
	this.byIdIndexMutex.Lock()
	defer this.byIdIndexMutex.Unlock()

	delete(this.byIdIndex, nodeId)
}

// O(n) complexity (n being active node count); thread safe
func (this *_dcgNodeRepo) ListWithOpGraphId(opGraphId string) []*dcgNodeDescriptor {
	this.byIdIndexMutex.RLock()
	defer this.byIdIndexMutex.RUnlock()

	nodesWithGraphIdSlice := []*dcgNodeDescriptor{}

	for _, node := range this.byIdIndex {
		if node.OpGraphId == opGraphId {
			nodesWithGraphIdSlice = append(nodesWithGraphIdSlice, node)
		}
	}
	return nodesWithGraphIdSlice
}

// O(1) complexity; thread safe
func (this *_dcgNodeRepo) GetIfExists(nodeId string) *dcgNodeDescriptor {
	this.byIdIndexMutex.RLock()
	defer this.byIdIndexMutex.RUnlock()

	return this.byIdIndex[nodeId]
}
