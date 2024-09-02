package tektoncd

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
	// WatchTasks watches all tasks inside a namespace.
	WatchTasks(ctx context.Context, namespace string) (<-chan watch.Event, error)
	// WatchTask watches a task.
	WatchTask(ctx context.Context, task *tektonAPI.Task, id uuid.UUID) (<-chan watch.Event, error)
	// SetTaskLabels sets the labels of a task.
	SetTaskLabels(ctx context.Context, task *tektonAPI.Task, labels map[string]string) error
	// GetTaskRun returns a task run by name.
	GetTaskRun(ctx context.Context, namespace string, id uuid.UUID) (*tektonAPI.TaskRun, error)
	// RunTask runs a task.
	RunTaskRun(ctx context.Context, taskRun *tektonAPI.TaskRun) error
	// WatchTaskRun watches a task run.
	WatchTaskRun(ctx context.Context, taskRun *tektonAPI.TaskRun, id uuid.UUID) (<-chan watch.Event, error)
	// GetTaskRunLogs returns a stream of logs for a task run.
	GetTaskRunLogs(ctx context.Context, taskRun *tektonAPI.TaskRun) (map[string]string, error)
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

// WatchTasks watches all tasks inside a namespace.
func (c *TektonClientImpl) WatchTasks(ctx context.Context, namespace string) (<-chan watch.Event, error) {
	// Watch the tasks.
	watcher, err := c.tektonClient.TektonV1().Tasks(namespace).Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return watcher.ResultChan(), nil
}

// WatchTask watches a task.
func (c *TektonClientImpl) WatchTask(ctx context.Context, task *tektonAPI.Task, id uuid.UUID) (<-chan watch.Event, error) {
	namespace := task.Namespace
	timeout := int64(5)
	// Watch the task.
	watcher, err := c.tektonClient.TektonV1().Tasks(namespace).Watch(ctx, metav1.ListOptions{
		TimeoutSeconds: &timeout,
		LabelSelector:  "choregate.fandujar.dev/task-id=" + id.String(),
		Watch:          true,
	})
	if err != nil {
		return nil, err
	}

	return watcher.ResultChan(), nil
}

// SetTaskLabels sets the labels of a task.
func (c *TektonClientImpl) SetTaskLabels(ctx context.Context, task *tektonAPI.Task, labels map[string]string) error {
	namespace := task.Namespace

	// Set the labels.
	task.Labels = labels
	_, err := c.tektonClient.TektonV1().Tasks(namespace).Update(ctx, task, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
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
func (c *TektonClientImpl) Logs(ctx context.Context, taskRun *tektonAPI.TaskRun) (map[string]string, error) {

	pod, err := c.kubeClient.CoreV1().Pods(taskRun.Namespace).Get(ctx, taskRun.Status.PodName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if pod.Status.Phase == v1.PodPending {
		return nil, fmt.Errorf("pod is pending")
	}

	var logs map[string]string = make(map[string]string)

	for _, container := range pod.Spec.Containers {
		raw, err := c.kubeClient.CoreV1().Pods(taskRun.Namespace).GetLogs(pod.Name, &v1.PodLogOptions{
			Container: container.Name,
		}).DoRaw(ctx)

		if err != nil {
			return nil, err
		}

		logs[container.Name] = string(raw)
	}

	return logs, nil
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

func (c *TektonClientImpl) GetTaskRunLogs(ctx context.Context, taskRun *tektonAPI.TaskRun) (map[string]string, error) {
	return c.Logs(ctx, taskRun)
}
