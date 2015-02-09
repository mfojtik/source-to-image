package build

import "github.com/openshift/source-to-image/pkg/sti/api"

// Builder defines an arbitrary builder interface that performs the build
// based on the Request
type Builder interface {
	Build(*api.Request) (*api.Result, error)
}
