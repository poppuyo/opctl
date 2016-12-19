package main

import (
  "github.com/opspec-io/engine/tcp"
  "github.com/opspec-io/engine/core"
  containerprovider "github.com/opspec-io/engine/pkg/containerprovider/providers/fake"
  pubsubprovider "github.com/opspec-io/engine/pkg/pubsubprovider/providers/fake"
)

func main() {

  tcp.New(
    core.New(
      new(containerprovider.ContainerProvider),
      new(pubsubprovider.PubSubProvider),
    ),
  ).Start()

}
