package core

import (
  "github.com/opspec-io/sdk-golang/pkg/model"
  "time"
  "github.com/opspec-io/engine/pkg/pubsubprovider"
  "fmt"
)

func (this _core) GetEventStream(
subscriberEventChannel chan model.Event,
) (err error) {

  subscriptionId := this.uniqueStringFactory.Construct()

  // cleanup on exit
  defer this.eventsPubSubTopic.DeleteSubscriptionIfExists(subscriptionId)

  this.eventsPubSubTopic.CreateSubscriptionIfNotExists(subscriptionId)
  var subscription pubsubprovider.Subscription
  subscription, err = this.eventsPubSubTopic.GetSubscription(subscriptionId)
  if (nil != err) {
    return
  }

  for {
    var msg pubsubprovider.Message
    msg, err = subscription.Pull()
    if (nil != err) {
      return
    }

    event, ok := msg.Data().(model.Event)
    if (!ok) {
      msg.Done(true)
      fmt.Printf("GetEventStream received msg w/ invalid data: '%v' and discarded the msg", msg.Data())
      continue
    }

    select {
    case subscriberEventChannel <- event:

    // handle channels that timeout
    case <-time.After(time.Second * 5):
      close(subscriberEventChannel)
      return
    }
    msg.Done(true)
  }

  return
}
