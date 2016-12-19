package core

import (
  "github.com/opspec-io/sdk-golang/pkg/model"
  "sync"
)

// repo for op instances
type opInstanceRepo interface {
  // adds the provided op
  add(opInstance *model.OpInstanceView)
  // deletes the op with the provided id if it exists
  deleteIfExists(opInstanceId string)
  // lists all op runs with the provided parentId
  listWithParentId(parentId string) []*model.OpInstanceView
  // tries to get the op with the provided id; returns nil if not found
  getIfExists(opInstanceId string) *model.OpInstanceView
}

func newOpInstanceRepo() opInstanceRepo {

  return &_opInstanceRepo{
    byIdIndex:make(map[string]*model.OpInstanceView),
    byIdIndexMutex:sync.RWMutex{},
  }

}

type _opInstanceRepo struct {
  byIdIndex      map[string]*model.OpInstanceView
  byIdIndexMutex sync.RWMutex
}

// O(1) complexity; thread safe
func (this *_opInstanceRepo) add(opInstance *model.OpInstanceView) {
  this.byIdIndexMutex.Lock()
  defer this.byIdIndexMutex.Unlock()
  this.byIdIndex[opInstance.Id] = opInstance
}

// O(1) complexity; thread safe
func (this *_opInstanceRepo) deleteIfExists(opInstanceId string) {
  this.byIdIndexMutex.Lock()
  defer this.byIdIndexMutex.Unlock()

  delete(this.byIdIndex, opInstanceId)
}

// O(n) complexity (n being running op instance count); thread safe
func (this *_opInstanceRepo) listWithParentId(parentId string) []*model.OpInstanceView {
  this.byIdIndexMutex.RLock()
  defer this.byIdIndexMutex.RUnlock()

  opInstancesWithParentIdSlice := []*model.OpInstanceView{}

  for _, opInstance := range this.byIdIndex {
    if (opInstance.ParentId == parentId) {
      opInstancesWithParentIdSlice = append(opInstancesWithParentIdSlice, opInstance)
    }
  }
  return opInstancesWithParentIdSlice
}

// O(1) complexity; thread safe
func (this *_opInstanceRepo) getIfExists(opInstanceId string) *model.OpInstanceView {
  this.byIdIndexMutex.RLock()
  defer this.byIdIndexMutex.RUnlock()

  return this.byIdIndex[opInstanceId]
}
