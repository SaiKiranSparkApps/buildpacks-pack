package dist

import (
	"os"
	"path/filepath"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/pkg/errors"
)

func LayerDiffID(layerTarPath string) (v1.Hash, error) {
	fh, err := os.Open(filepath.Clean(layerTarPath))
	if err != nil {
		return v1.Hash{}, errors.Wrap(err, "opening tar file")
	}
	defer fh.Close()

	layer, err := tarball.LayerFromFile(layerTarPath)
	if err != nil {
		return v1.Hash{}, errors.Wrap(err, "reading layer tar")
	}

	hash, err := layer.DiffID()
	if err != nil {
		return v1.Hash{}, errors.Wrap(err, "generating diff id")
	}

	return hash, nil
}

func AddToLayersMD(layerMD BuildpackLayers, descriptor BuildpackDescriptor, diffID string) {
	info := descriptor.Info()
	if _, ok := layerMD[info.ID]; !ok {
		layerMD[info.ID] = map[string]BuildpackLayerInfo{}
	}
	layerMD[info.ID][info.Version] = BuildpackLayerInfo{
		API:         descriptor.API,
		Stacks:      descriptor.Stacks,
		Order:       descriptor.Order,
		LayerDiffID: diffID,
		Homepage:    info.Homepage,
		Name:        info.Name,
	}
}
