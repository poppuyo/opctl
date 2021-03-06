package core

import (
	"github.com/opspec-io/sdk-golang/pkg/model"
)

func (this _core) StartOp(
	req model.StartOpReq,
) (
	opId string,
	err error,
) {

	opId = this.uniqueStringFactory.Construct()

	normalizedOpRef := this.pathNormalizer.Normalize(req.OpRef)

	go func() {
		this.opCaller.Call(
			req.Args,
			opId,
			normalizedOpRef,
			opId,
		)
	}()

	return

}
