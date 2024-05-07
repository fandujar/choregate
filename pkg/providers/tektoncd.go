package providers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	tektonAPI "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonVersioned "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	cliconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

// TektonClient is a client for interacting with Tekton.
type TektonClient interface {
	// GetTaskRun returns a task run by name.
	GetTaskRun(ctx context.Context, namespace string, id uuid.UUID) (*tektonAPI.TaskRun, error)
	// RunTask runs a task.
	RunTaskRun(ctx context.Context, taskRun *tektonAPI.TaskRun) error
	// WatchTaskRun watches a task run.
	WatchTaskRun(ctx context.Context, taskRun *tektonAPI.TaskRun, id uuid.UUID) (<-chan watch.Event, error)
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

// GetTaskRun returns a task run by ID.
func (c *TektonClientImpl) GetTaskRun(ctx context.Context, namespace string, id uuid.UUID) (*tektonAPI.TaskRun, error) {
	// Find the task run by ID in the labels.
	taskRunList, err := c.tektonClient.TektonV1().TaskRuns(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: "choregate.fandujar.dev/taskrun-id=" + id.String(),
	})
	if err != nil {
		return nil, err
	}
	for _, taskRun := range taskRunList.Items {
		return &taskRun, nil
	}
	return nil, fmt.Errorf("task run not found")
}

// Logs returns the logs of a task run.
func (c *TektonClientImpl) Logs(ctx context.Context, taskRun *tektonAPI.TaskRun) (string, error) {
	// Get the logs of the task run.
	raw, err := c.kubeClient.CoreV1().Pods(taskRun.Namespace).GetLogs(taskRun.Status.PodName, &v1.PodLogOptions{}).Do(ctx).Raw()
	if err != nil {
		return "", err
	}

	if raw == nil {
		return "", fmt.Errorf("failed to get logs")
	}

	return string(raw), nil
}

func (c *TektonClientImpl) RunTaskRun(ctx context.Context, taskRun *tektonAPI.TaskRun) error {
	namespace := taskRun.Namespace

	_, err := c.tektonClient.TektonV1().TaskRuns(namespace).Create(ctx, taskRun, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (c *TektonClientImpl) WatchTaskRun(ctx context.Context, taskRun *tektonAPI.TaskRun, id uuid.UUID) (<-chan watch.Event, error) {
	namespace := taskRun.Namespace

	// Watch the task run.
	watcher, err := c.tektonClient.TektonV1().TaskRuns(namespace).Watch(ctx, metav1.ListOptions{
		LabelSelector: "choregate.fandujar.dev/taskrun-id=" + id.String(),
	})
	if err != nil {
		return nil, err
	}

	return watcher.ResultChan(), nil
}
