package pullsecrets

import (
	"github.com/devspace-cloud/devspace/pkg/devspace/config/generated"
	"github.com/devspace-cloud/devspace/pkg/devspace/config/versions/latest"
	"github.com/devspace-cloud/devspace/pkg/devspace/docker"
	"github.com/devspace-cloud/devspace/pkg/devspace/hook"
	"github.com/devspace-cloud/devspace/pkg/devspace/kubectl"
	"github.com/devspace-cloud/devspace/pkg/util/log"
)

// Client communicates with a registry
type Client interface {
	CreatePullSecrets() error
	CreatePullSecret(options *PullSecretOptions) error
}

// NewClient creates a client for a registry
func NewClient(config *latest.Config, cache *generated.CacheConfig, kubeClient kubectl.Client, dockerClient docker.Client, log log.Logger) Client {
	return &client{
		config:       config,
		cache:        cache,
		kubeClient:   kubeClient,
		dockerClient: dockerClient,
		hookExecuter: hook.NewExecuter(config),
		log:          log,
	}
}

type client struct {
	config       *latest.Config
	cache        *generated.CacheConfig
	kubeClient   kubectl.Client
	dockerClient docker.Client
	hookExecuter hook.Executer
	log          log.Logger
}
