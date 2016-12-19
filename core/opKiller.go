package core

import (
  "github.com/opspec-io/sdk-golang/pkg/model"
  "time"
  "sync"
  "fmt"
  "github.com/opspec-io/engine/pkg/pubsubprovider"
  "github.com/opspec-io/engine/pkg/containerprovider"
)

type opKiller interface {
  Start()
}

func newOpKiller(
containerProvider containerprovider.ContainerProvider,
eventsPubSubTopic pubsubprovider.Topic,
opInstanceRepo opInstanceRepo,
opKillsPubSubTopic pubsubprovider.Topic,
) opKiller {

  opKillerPubSubSubscription, err := opKillsPubSubTopic.CreateSubscriptionIfNotExists("opkiller")
  if (nil != err) {
    panic(err)
  }

  return &_opKiller{
    containerProvider:containerProvider,
    eventsPubSubTopic:eventsPubSubTopic,
    opInstanceRepo:opInstanceRepo,
    opKillsPubSubTopic:opKillsPubSubTopic,
    opKillerPubSubSubscription:opKillerPubSubSubscription,
  }
}

type _opKiller struct {
  containerProvider          containerprovider.ContainerProvider
  eventsPubSubTopic          pubsubprovider.Topic
  opInstanceRepo             opInstanceRepo
  opKillsPubSubTopic         pubsubprovider.Topic
  opKillerPubSubSubscription pubsubprovider.Subscription
}

// polls for opkills and handles them
func (this _opKiller) Start() {
  go func() {
    for {
      msg, err := this.opKillerPubSubSubscription.Pull()
      if (nil != err) {
        fmt.Println(err.Error())
        continue
      }

      this.handleMsg(msg)
    }
  }()
  return
}

func (this _opKiller) handleMsg(
msg pubsubprovider.Message,
) {

  opKillDescriptor, ok := msg.Data().(opKillDescriptor)
  if (!ok) {
    msg.Done(true)
    fmt.Printf("OpKiller received msg w/ invalid data: '%v' and discarded the msg", msg.Data())
    return
  }

  var waitGroup sync.WaitGroup

  opInstance := this.opInstanceRepo.getIfExists(opKillDescriptor.OpInstanceId)
  // guard opInstance found
  if (nil == opInstance) {
    this.opInstanceRepo.deleteIfExists(opKillDescriptor.OpInstanceId)
  }

  // order of the following matters (hence the numbering)
  // 1) recover resources
  go func() {
    waitGroup.Add(1)
    // this.containerProvider.DeleteContainerIfExists()
    defer waitGroup.Done()
  }()

  // 2) queue child kills
  for _, childOpInstance := range this.opInstanceRepo.listWithParentId(opKillDescriptor.OpInstanceId) {
    this.opKillsPubSubTopic.Publish(childOpInstance.Id)
  }

  // 3) wait for 2 & 3
  waitGroup.Wait()

  // 4) delete from storage
  this.opInstanceRepo.deleteIfExists(opKillDescriptor.OpInstanceId)

  // 5) send OpEndedEvent last to ensure OpEndedEvent's send in reverse order of OpStartedEvent's
  this.eventsPubSubTopic.Publish(
    model.Event{
      Timestamp:time.Now().UTC(),
      OpEnded:&model.OpEndedEvent{
        OpInstanceId:opKillDescriptor.OpInstanceId,
        Outcome:model.OpOutcomeKilled,
      },
    },
  )

  msg.Done(true)

}
