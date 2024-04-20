package providers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonVersioned "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	cliconfig "sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/fandujar/choregate/pkg/entities"
)

// TektonClient is a client for interacting with Tekton.
type TektonClient interface {
	// GetTaskRun returns a task run by name.
	GetTaskRun(ctx context.Context, id uuid.UUID) (*entities.TaskRun, error)
	// RunTask runs a task.
	RunTask(ctx context.Context, task *entities.TaskRun) error
}

type TektonClientImpl struct {
	kubeClient   *kubernetes.Clientset
	tektonClient *tektonVersioned.Clientset
}

// NewTektonClient creates a new TektonClient.
func NewTektonClient() (TektonClient, error) {
	// Try in-cluster config or fallback to default config.
	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Warn().Err(err).Msg("failed to get in-cluster config, falling back to default config")
		cfg, err = cliconfig.GetConfig()
		if err != nil {
			return nil, err
		}
	}

	k, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	t, err := tektonVersioned.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &TektonClientImpl{
		kubeClient:   k,
		tektonClient: t,
	}, nil
}

func (c *TektonClientImpl) GetTaskRun(ctx context.Context, id uuid.UUID) (*entities.TaskRun, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *TektonClientImpl) RunTask(ctx context.Context, task *entities.TaskRun) error {

	t := &tekton.TaskRun{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", task.ID.String()),
		},
		Spec: tekton.TaskRunSpec{
			TaskSpec: &tekton.TaskSpec{
				Steps: task.Steps,
			},
		},
	}

	_, err := c.tektonClient.TektonV1().TaskRuns(task.Namespace).Create(ctx, t, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
