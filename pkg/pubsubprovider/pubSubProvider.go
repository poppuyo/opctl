package pubsubprovider

//go:generate counterfeiter -o providers/fake/pubSubProvider.go --fake-name PubSubProvider ./ PubSubProvider

type Message interface {
  Done(ack bool)
  Data() interface{}
}

type Subscription interface {
  Pull() (msg Message, err error)
}

type Topic interface {
  Publish(data interface{}) (err error)
  CreateSubscriptionIfNotExists(subscriptionId string) (subscription Subscription, err error)
  DeleteSubscriptionIfExists(subscriptionId string) (err error)
  GetSubscription(subscriptionId string) (subscription Subscription, err error)
}

type PubSubProvider interface {
  CreateTopicIfNotExists(topicId string) (topic Topic, err error)
  DeleteTopicIfExists(topicId string) (err error)
  GetTopic(topicId string) (topic Topic, err error)
}
