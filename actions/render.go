package actions

import (
	"encoding/json"
	"os"

	"github.com/duckbrain/shiboleet/services"
	"github.com/gobuffalo/tags/v3"
	"github.com/pkg/errors"
)

func NewAssetHelper(provider services.Provider) (AssetHelper, error) {
	helper := AssetHelper{}

	switch provider.Assets.Type {
	case "vite":
		helper.AssetPather = &VitePather{DevelopmentServerPath: provider.Assets.DevelopmentServerPath}
	default:
		return helper, errors.Errorf("unknown assets type %v", provider.Assets.Type)
	}

	if provider.Environment == "development" {
		return helper, nil
	}
	m := helper.Manifest()
	if m != nil {
		b, err := os.ReadFile(provider.Assets.ManifestPath)
		if err != nil {
			return helper, err
		}
		err = json.Unmarshal(b, m)
		if err != nil {
			return helper, err
		}
	}

	return helper, nil
}

type AssetPather interface {
	Path(path string) string
	Manifest() interface{}
}

type AssetHelper struct {
	AssetPather
}

type VitePather struct {
	manifest map[string]struct {
		File string `json:"file"`
	}
	DevelopmentServerPath string
}

func (p VitePather) Path(path string) string {
	if p.manifest == nil {
		return p.DevelopmentServerPath + path
	}
	s := p.manifest[path].File
	if s == "" {
		return path
	}
	return s
}

func (p *VitePather) Manifest() interface{} {
	return &p.manifest
}

func (h AssetHelper) JS(p string) *tags.Tag {
	return tags.JavascriptTag(tags.Options{"src": h.Path(p)})
}

func (h AssetHelper) CSS(p string) *tags.Tag {
	return tags.StylesheetTag(tags.Options{"href": h.Path(p)})
}
