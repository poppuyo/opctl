package core

//go:generate counterfeiter -o ./fakeCore.go --fake-name FakeCore ./ Core

import (
  "github.com/opspec-io/sdk-golang/pkg/model"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
  "github.com/opspec-io/engine/pkg/containerprovider"
  "github.com/opspec-io/engine/pkg/pubsubprovider"
  "github.com/opspec-io/engine/util/uniquestring"
  "github.com/opspec-io/engine/util/pathnormalizer"
)

type Core interface {
  GetEventStream(
  eventChannel chan model.Event,
  ) (err error)

  KillOp(
  req model.KillOpReq,
  ) (err error)

  StartOp(
  req model.StartOpReq,
  ) (
  opInstanceId string,
  err error,
  )
}

const (
  PubSubTopicIdPrefix = "opspec.engine"
)

func New(
containerProvider containerprovider.ContainerProvider,
pubSubProvider pubsubprovider.PubSubProvider,
) (core Core) {

  uniqueStringFactory := uniquestring.NewUniqueStringFactory()

  _bundle := bundle.New()

  opInstanceRepo := newOpInstanceRepo()

  // init pubsub
  eventsPubSubTopic, err := pubSubProvider.CreateTopicIfNotExists(PubSubTopicIdPrefix + ".events")
  if (nil != err) {
    panic(err)
  }
  opKillsPubSubTopic, err := pubSubProvider.CreateTopicIfNotExists(PubSubTopicIdPrefix + ".opkills")
  if (nil != err) {
    panic(err)
  }
  opStartsPubSubTopic, err := pubSubProvider.CreateTopicIfNotExists(PubSubTopicIdPrefix + ".opstarts")
  if (nil != err) {
    panic(err)
  }

  newOpKiller(containerProvider, eventsPubSubTopic, opInstanceRepo, opKillsPubSubTopic).Start()

  core = &_core{
    bundle:_bundle,
    containerProvider:containerProvider,
    eventsPubSubTopic:eventsPubSubTopic,
    opKillsPubSubTopic:opKillsPubSubTopic,
    opStartsPubSubTopic:opStartsPubSubTopic,
    opInstanceRepo:opInstanceRepo,
    pathNormalizer:pathnormalizer.NewPathNormalizer(),
    pubSubProvider:pubSubProvider,
    uniqueStringFactory:uniqueStringFactory,
  }
  return
}

type _core struct {
  bundle              bundle.Bundle
  containerProvider   containerprovider.ContainerProvider
  eventsPubSubTopic   pubsubprovider.Topic
  opKillsPubSubTopic  pubsubprovider.Topic
  opStartsPubSubTopic pubsubprovider.Topic
  opInstanceRepo      opInstanceRepo
  pathNormalizer      pathnormalizer.PathNormalizer
  pubSubProvider      pubsubprovider.PubSubProvider
  uniqueStringFactory uniquestring.UniqueStringFactory
}

func (this _core) KillOp(
req model.KillOpReq,
) (err error) {
  opKillDescriptor := opKillDescriptor{req}
  err = this.opKillsPubSubTopic.Publish(opKillDescriptor)
  return
}

func (this _core) StartOp(
req model.StartOpReq,
) (
opInstanceId string,
err error,
) {
  opInstanceId = this.uniqueStringFactory.Construct()
  this.opStartsPubSubTopic.Publish()
  return
}
