package docker

import (
	"bytes"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/reference"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/opctl/util/eventbus"
	"github.com/pkg/errors"
	"io/ioutil"
)

var _ = Context("pullImage", func() {
	Context("imageRef invalid", func() {
		It("should return expected error", func() {
			/* arrange */
			providedImageRef := "%$^"
			_, _, expectedError := reference.Parse(providedImageRef)

			objectUnderTest := _containerEngine{
				dockerClient: new(fakeDockerClient),
			}

			/* act */
			actualError := objectUnderTest.pullImage(
				providedImageRef,
				"",
				"",
				new(eventbus.FakeEventPublisher),
			)

			/* assert */
			Expect(actualError).To(Equal(expectedError))
		})
	})
	Context("imageRef valid", func() {
		It("should call dockerClient.ImagePull w/ expected args", func() {
			/* arrange */
			providedImageRef := "dummy-ref"
			expectedImageRef := fmt.Sprintf("%v:latest", providedImageRef)
			expectedImagePullOptions := types.ImagePullOptions{}

			imagePullResponse := ioutil.NopCloser(bytes.NewBufferString(""))

			_fakeDockerClient := new(fakeDockerClient)
			_fakeDockerClient.ImagePullReturns(imagePullResponse, nil)

			objectUnderTest := _containerEngine{
				dockerClient: _fakeDockerClient,
			}

			/* act */
			err := objectUnderTest.pullImage(
				providedImageRef,
				"",
				"",
				new(eventbus.FakeEventPublisher),
			)
			if nil != err {
				panic(err)
			}

			/* assert */
			_, actualImageRef, actualImagePullOptions := _fakeDockerClient.ImagePullArgsForCall(0)
			Expect(actualImageRef).To(Equal(expectedImageRef))
			Expect(actualImagePullOptions).To(Equal(expectedImagePullOptions))
		})
		Context("dockerClient.ImagePull errors", func() {
			It("should return expected error", func() {
				/* arrange */
				imagePullError := errors.New("dummyerror")
				expectedError := imagePullError
				imagePullResponse := ioutil.NopCloser(bytes.NewBufferString(""))

				_fakeDockerClient := new(fakeDockerClient)
				_fakeDockerClient.ImagePullReturns(imagePullResponse, imagePullError)

				objectUnderTest := _containerEngine{
					dockerClient: _fakeDockerClient,
				}

				/* act */
				actualError := objectUnderTest.pullImage(
					"dummy-ref",
					"",
					"",
					new(eventbus.FakeEventPublisher),
				)

				/* assert */
				Expect(actualError).To(Equal(expectedError))
			})
		})
	})
})
