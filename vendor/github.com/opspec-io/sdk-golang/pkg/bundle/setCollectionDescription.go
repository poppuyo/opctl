package bundle

import (
	"github.com/opspec-io/sdk-golang/pkg/model"
	"path"
)

func (this _bundle) SetCollectionDescription(
	req model.SetCollectionDescriptionReq,
) (err error) {

	pathToCollectionManifest := path.Join(req.PathToCollection, NameOfCollectionManifestFile)

	collectionManifestBytes, err := this.fileSystem.GetBytesOfFile(
		pathToCollectionManifest,
	)
	if (nil != err) {
		return
	}

	collectionManifest := model.CollectionManifest{}
	err = this.yaml.To(
		collectionManifestBytes,
		&collectionManifest,
	)
	if (nil != err) {
		return
	}

	collectionManifest.Description = req.Description

	collectionManifestBytes, err = this.yaml.From(&collectionManifest)
	if (nil != err) {
		return
	}

	err = this.fileSystem.SaveFile(
		pathToCollectionManifest,
		collectionManifestBytes,
	)

	return

}
