package core

import "github.com/opspec-io/sdk-golang/pkg/model"

type containerCallDescriptor struct {
  model.ContainerCallInstance
  ContainerInstanceId string
}

type opCallDescriptor struct {
  // map of args keyed by param name
  OpArgs  map[string]*model.Arg
  OpRef string
  OpInstanceId string
}

type opKillDescriptor struct {
  model.KillOpReq
}
