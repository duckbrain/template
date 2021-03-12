package assets

import (
	"encoding/json"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/tags/v3"
	"github.com/pkg/errors"
)

type Config struct {
	Type         string
	ManifestPath string
	ServerPath   string
}

func MustNewHelper(c Config) Helper {
	h, err := NewHelper(c)
	if err != nil {
		panic(err)
	}
	return h
}

func NewReleaseHelper(c Config) (Helper, error) {
	helper := Helper{}

	switch c.Type {
	case "vite":
		helper.pather = &vite{DevelopmentServerPath: c.ServerPath}
	default:
		return helper, errors.Errorf("unknown assets type %v", c.Type)
	}

	m := helper.Manifest()
	if m != nil {
		b, err := os.ReadFile(c.ManifestPath)
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

func NewDebugHelper(c Config) (Helper, error) {
	helper := Helper{}

	switch c.Type {
	case "vite":
		helper.pather = &vite{DevelopmentServerPath: c.ServerPath}
	default:
		return helper, errors.Errorf("unknown assets type %v", c.Type)
	}
	return helper, nil
}

type pather interface {
	Path(path string) string
	Manifest() interface{}
}

type Helper struct {
	pather
}

type vite struct {
	manifest map[string]struct {
		File string `json:"file"`
	}
	DevelopmentServerPath string
}

func (p vite) Path(path string) string {
	if p.manifest == nil {
		return p.DevelopmentServerPath + path
	}
	s := p.manifest[path].File
	if s == "" {
		return path
	}
	return s
}

func (p *vite) Manifest() interface{} {
	return &p.manifest
}

func (h Helper) JS(p string) *tags.Tag {
	return tags.JavascriptTag(tags.Options{"src": h.Path(p)})
}

func (h Helper) CSS(p string) *tags.Tag {
	return tags.StylesheetTag(tags.Options{"href": h.Path(p)})
}

func (h Helper) Middleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("asset", h)
		return next(c)
	}
}
