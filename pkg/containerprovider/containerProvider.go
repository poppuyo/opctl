package containerprovider

import "github.com/opspec-io/sdk-golang/pkg/model"

//go:generate counterfeiter -o providers/fake/containerProvider.go --fake-name ContainerProvider ./ ContainerProvider


// By design, this interface keeps container execution, networking, and filesystem, the responsibility of the
// container provider.
type ContainerProvider interface {
  DeleteContainerIfExists(
  containerId string,
  ) (err error)

  StartContainer(
  call *model.ContainerCallInstance,
  ) (
  containerId string,
  err error,
  )
}
