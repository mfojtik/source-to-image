package api

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"

	"github.com/openshift/source-to-image/pkg/util/user"
)

// Image label namespace constants
const (
	DefaultNamespace    = "io.openshift.s2i."
	KubernetesNamespace = "io.k8s."
)

const (
	// PullAlways means that we always attempt to pull the latest image.
	PullAlways PullPolicy = "always"

	// PullNever means that we never pull an image, but only use a local image.
	PullNever PullPolicy = "never"

	// PullIfNotPresent means that we pull if the image isn't present on disk.
	PullIfNotPresent PullPolicy = "if-not-present"

	// DefaultBuilderPullPolicy specifies the default pull policy to use
	DefaultBuilderPullPolicy = PullIfNotPresent

	// DefaultPreviousImagePullPolicy specifies policy for pulling the previously
	// build Docker image when doing incremental build
	DefaultPreviousImagePullPolicy = PullAlways
)

// SourceConfig describes the repository configuration
type SourceConfig struct {
	// URI describing the location of sources used to build the result image.
	URI string

	// Ref is a tag/branch to be used for build.
	Ref string

	// NoRecursive disables the --recursive option for the git clone that
	// allows to use the GIT without requiring the git submodule to be called.
	NoRecursive bool

	// Specify a relative directory inside the application repository that should
	// be used as a root directory for the application.
	ContextDir string

	// WorkingDir describes temporary directory used for downloading sources,
	// scripts and tar operations.
	WorkingDir string

	// WorkingSourceDir describes the subdirectory off of WorkingDir set up during
	// the repo download that is later used as the root for ignore processing
	WorkingSourceDir string
}

type OutputConfig struct {
	// Tag is a result image tag name.
	Tag string

	// DisplayName is a result image display-name label. This defaults to the
	// output image name.
	DisplayName string

	// Description is a result image description label. The default is no description.
	Description string

	// CallbackURL is a URL which is called upon successful build to inform about that fact.
	CallbackURL string

	// LabelNamespace provides the namespace under which the labels will be generated.
	LabelNamespace string
}

type BuildConfig struct {
	// Image is the name of the Docker image we use as a builder image
	Image string

	// ImageVersion provides optional version information about the builder image.
	ImageVersion string

	// BaseImageVersion provides optional version information about the builder base image.
	BaseImageVersion string

	// BuilderPullPolicy specifies when to pull the builder image
	BuilderPullPolicy PullPolicy

	// PreviousImagePullPolicy specifies when to pull the previously build image
	// when doing incremental build
	PreviousImagePullPolicy PullPolicy

	// Incremental describes whether to try to perform incremental build.
	IsIncremental bool

	// RemovePreviousImage describes if previous image should be removed after successful build.
	// This applies only to incremental builds.
	RemovePreviousImage bool

	// ForcePull defines if the builder image should be always pulled or not.
	// DEPRECATED: Use the BuilderPullPolicy instead of ForcePull
	// Setting this to 'true' equals setting BuilderPullPolicy to 'PullAlways'.
	// Setting this to 'false' equals setting BuilderPullPolicy to 'PullIfNotPresent'
	ForcePull bool

	// Environment is a map of environment variables to be passed to the image.
	Environment map[string]string

	// AssembleUser specifies the user to run the assemble script in container
	AssembleUser string

	// LayeredBuild describes if this is build which layered scripts and sources
	// on top of BuilderImage.
	IsLayered bool

	// AllowedUIDs is a list of user ranges of users allowed to run the builder image.
	// If a range is specified and the builder image uses a non-numeric user or a user
	// that is outside the specified range, then the build fails.
	AllowedUIDs user.RangeList

	// ScriptsURI is a URI describing the location of the S2I scripts used during
	// build process.
	ScriptsURI string

	// Destination specifies a location where the untar operation will place its artifacts.
	Destination string
}

// Config contains essential fields for performing build.
type Config struct {
	// Source is a repository configuration
	Source *SourceConfig

	// Build contains the build specific configuration
	Build *BuildConfig

	// Result describes the result image
	Output *OutputConfig

	// DockerConfig describes how to access host docker daemon.
	DockerConfig *DockerConfig

	// Operate quietly. Progress and assemble script output are not reported, only fatal errors.
	// (default: false).
	Quiet bool

	// RunImage will trigger a "docker run ..." invocation of the produced image so the user
	// can see if it operates as he would expect
	RunImage bool

	// Usage allows for properly shortcircuiting s2i logic when `s2i usage` is invoked
	Usage bool
}

// DockerConfig contains the configuration for a Docker connection
type DockerConfig struct {
	// NetworkMode is used to set the docker network setting to --net=container:<id>
	// when the builder is invoked from a container.
	NetworkMode DockerNetworkMode

	// PullAuthentication holds the authentication information for pulling the
	// Docker images from private repositories
	PullAuthentication docker.AuthConfiguration

	// IncrementalAuthentication holds the authentication information for pulling the
	// previous image from private repositories
	IncrementalAuthentication docker.AuthConfiguration

	// Path provides the path to the .dockercfg file
	Path string

	// Endpoint is the docker network endpoint or socket
	Endpoint string

	// CertFile is the certificate file path for a TLS connection
	CertFile string

	// KeyFile is the key file path for a TLS connection
	KeyFile string

	// CAFile is the certificate authority file path for a TLS connection
	CAFile string
}

// Result structure contains information from build process.
type Result struct {

	// Success describes whether the build was successful.
	Success bool

	// Messages is a list of messages from build process.
	Messages []string

	// WorkingDir describes temporary directory used for downloading sources, scripts and tar operations.
	WorkingDir string

	// ImageID describes resulting image ID.
	ImageID string
}

// InstallResult structure describes the result of install operation
type InstallResult struct {
	// Script describes which script this result refers to
	Script string

	// URL describes from where the script was taken
	URL string

	// Downloaded describes if download operation happened, this will be true for
	// external scripts, but false for scripts from inside the image
	Downloaded bool

	// Installed describes if script was installed to upload directory
	Installed bool

	// Error describes last error encountered during install operation
	Error error
}

// SourceInfo stores information about the source code
type SourceInfo struct {
	// Ref represents a commit SHA-1, valid GIT branch name or a GIT tag
	// The output image will contain this information as 'io.openshift.build.commit.ref' label.
	Ref string

	// CommitID represents an arbitrary extended object reference in GIT as SHA-1
	// The output image will contain this information as 'io.openshift.build.commit.id' label.
	CommitID string

	// Date contains a date when the committer created the commit.
	// The output image will contain this information as 'io.openshift.build.commit.date' label.
	Date string

	// AuthorName contains the name of the author
	// The output image will contain this information (along with AuthorEmail) as 'io.openshift.build.commit.author' label.
	AuthorName string

	// AuthorEmail contains the e-mail of the author
	// The output image will contain this information (along with AuthorName) as 'io.openshift.build.commit.author' lablel.
	AuthorEmail string

	// CommitterName contains the name of the committer
	CommitterName string

	// CommitterEmail contains the e-mail of the committer
	CommitterEmail string

	// Message represents the first 80 characters from the commit message.
	// The output image will contain this information as 'io.openshift.build.commit.message' label.
	Message string

	// Location contains a valid URL to the original repository.
	// The output image will contain this information as 'io.openshift.build.source-location' label.
	Location string

	// ContextDir contains path inside the Location directory that
	// contains the application source code.
	// The output image will contain this information as 'io.openshift.build.source-context-dir'
	// label.
	ContextDir string
}

// CloneConfig specifies the options used when cloning the application source
// code.
type CloneConfig struct {
	Recursive bool
	Quiet     bool
}

// DockerNetworkMode specifies the network mode setting for the docker container
type DockerNetworkMode string

const (
	// DockerNetworkModeHost places the container in the default (host) network namespace.
	DockerNetworkModeHost DockerNetworkMode = "host"
	// DockerNetworkModeBridge instructs docker to create a network namespace for this container connected to the docker0 bridge via a veth-pair.
	DockerNetworkModeBridge DockerNetworkMode = "bridge"
	// DockerNetworkModeContainerPrefix is the string prefix used by NewDockerNetworkModeContainer.
	DockerNetworkModeContainerPrefix string = "container:"
)

// NewDockerNetworkModeContainer creates a DockerNetworkMode value which instructs docker to place the container in the network namespace of an existing container.
// It can be used, for instance, to place the s2i container in the network namespace of the infrastructure container of a k8s pod.
func NewDockerNetworkModeContainer(id string) DockerNetworkMode {
	return DockerNetworkMode(DockerNetworkModeContainerPrefix + id)
}

// PullPolicy specifies a type for the method used to retrieve the Docker image
type PullPolicy string

// String implements the String() function of pflags.Value so this can be used as
// command line parameter.
// This method is really used just to show the default value when printing help.
// It will not default the configuration.
func (p *PullPolicy) String() string {
	if len(string(*p)) == 0 {
		return string(DefaultBuilderPullPolicy)
	}
	return string(*p)
}

// Type implements the Type() function of pflags.Value interface
func (p *PullPolicy) Type() string {
	return "string"
}

// Set implements the Set() function of pflags.Value interface
// The valid options are "always", "never" or "if-not-present"
func (p *PullPolicy) Set(v string) error {
	switch v {
	case "always":
		*p = PullAlways
	case "never":
		*p = PullNever
	case "if-not-present":
		*p = PullIfNotPresent
	default:
		return fmt.Errorf("invalid value %q, valid values are: always, never or if-not-present")
	}
	return nil
}
