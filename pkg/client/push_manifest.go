package client

import (
	"context"

	runtime "github.com/buildpacks/pack/internal/runtime"
)

type PushManifestOptions struct {
	Format          string
	Insecure, Purge bool
}

// PushManifest implements commands.PackClient.
func (c *Client) PushManifest(ctx context.Context, index string, opts PushManifestOptions) (imageID string, err error) {
	manifestList, err := c.runtime.LookupImageIndex(index)
	if err != nil {
		return
	}

	_, err = manifestList.Push(ctx, parseFalgsForImgUtil(opts))

	if err == nil && opts.Purge {
		c.runtime.RemoveManifests(ctx, []string{index})
	}

	return imageID, err
}

func parseFalgsForImgUtil(opts PushManifestOptions) (idxOptions runtime.PushOptions) {
	return idxOptions
}