package core

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ Core

import (
	"github.com/opspec-io/opctl/pkg/containerengine"
	"github.com/opspec-io/opctl/util/eventbus"
	"github.com/opspec-io/opctl/util/pathnormalizer"
	"github.com/opspec-io/opctl/util/uniquestring"
	"github.com/opspec-io/sdk-golang/pkg/bundle"
	"github.com/opspec-io/sdk-golang/pkg/model"
	"github.com/opspec-io/sdk-golang/pkg/validate"
)

type Core interface {
	GetEventStream(
		req *model.GetEventStreamReq,
		eventChannel chan model.Event,
	) (err error)

	KillOp(
		req model.KillOpReq,
	)

	StartOp(
		req model.StartOpReq,
	) (
		callId string,
		err error,
	)
}

func New(
	containerEngine containerengine.ContainerEngine,
) (core Core) {
	uniqueStringFactory := uniquestring.NewUniqueStringFactory()

	eventBus := eventbus.New()

	_bundle := bundle.New()

	dcgNodeRepo := newDcgNodeRepo()

	caller := newCaller(
		newContainerCaller(
			_bundle,
			containerEngine,
			eventBus,
			dcgNodeRepo,
		),
	)

	caller.setParallelCaller(
		newParallelCaller(
			caller,
			uniqueStringFactory,
		),
	)

	caller.setSerialCaller(
		newSerialCaller(
			caller,
			uniqueStringFactory,
		),
	)

	opCaller := newOpCaller(
		_bundle,
		eventBus,
		dcgNodeRepo,
		caller,
		uniqueStringFactory,
		validate.New(),
	)

	caller.setOpCaller(
		opCaller,
	)

	core = _core{
		containerEngine:     containerEngine,
		eventBus:            eventBus,
		opCaller:            opCaller,
		pathNormalizer:      pathnormalizer.NewPathNormalizer(),
		dcgNodeRepo:         dcgNodeRepo,
		uniqueStringFactory: uniqueStringFactory,
	}

	return
}

type _core struct {
	containerEngine     containerengine.ContainerEngine
	eventBus            eventbus.EventBus
	caller              caller
	pathNormalizer      pathnormalizer.PathNormalizer
	dcgNodeRepo         dcgNodeRepo
	uniqueStringFactory uniquestring.UniqueStringFactory
	opCaller            opCaller
}
