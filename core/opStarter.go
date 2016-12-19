package core

import (
  "github.com/opspec-io/sdk-golang/pkg/model"
  "time"
  "sync"
  "fmt"
  "github.com/opspec-io/engine/pkg/pubsubprovider"
  "github.com/opspec-io/engine/pkg/containerprovider"
)

type opStarter interface {
  Start()
}

func newOpStarter(
containerProvider containerprovider.ContainerProvider,
eventsPubSubTopic pubsubprovider.Topic,
opInstanceRepo opInstanceRepo,
opStartsPubSubTopic pubsubprovider.Topic,
) opStarter {

  opStarterPubSubSubscription, err := opStartsPubSubTopic.CreateSubscriptionIfNotExists("opstarter")
  if (nil != err) {
    panic(err)
  }

  return &_opStarter{
    containerProvider:containerProvider,
    eventsPubSubTopic:eventsPubSubTopic,
    opInstanceRepo:opInstanceRepo,
    opStartsPubSubTopic:opStartsPubSubTopic,
    opStarterPubSubSubscription:opStarterPubSubSubscription,
  }
}

type _opStarter struct {
  containerProvider          containerprovider.ContainerProvider
  eventsPubSubTopic          pubsubprovider.Topic
  opInstanceRepo             opInstanceRepo
  opStartsPubSubTopic         pubsubprovider.Topic
  opStarterPubSubSubscription pubsubprovider.Subscription
}

// polls for opstarts and handles them
func (this _opStarter) Start() {
  go func() {
    for {
      msg, err := this.opStarterPubSubSubscription.Pull()
      if (nil != err) {
        fmt.Println(err.Error())
        continue
      }

      this.handleMsg(msg)
    }
  }()
  return
}

func (this _opStarter) handleMsg(
msg pubsubprovider.Message,
) {

  opStartDescriptor, ok := msg.Data().(opStartDescriptor)
  if (!ok) {
    msg.Done(true)
    fmt.Printf("OpStarter received msg w/ invalid data: '%v' and discarded the msg", msg.Data())
    return
  }

  msg.Done(true)

}
