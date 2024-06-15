package workflow

import (
	"context"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type activity func(ctx context.Context, param interface{}) (interface{}, error)
type flow func(ctx workflow.Context, param interface{}) (interface{}, error)

// Worker is a Worker in the Saga.
type Worker struct {
	name   string
	worker worker.Worker
	client client.Client
}

func NewWorker(
	name string,
	cl client.Client,
) *Worker {
	w := &Worker{
		name:   name,
		worker: worker.New(cl, name, worker.Options{}),
		client: cl,
	}
	return w
}

func (w *Worker) GetName() string {
	return w.name
}

func (w *Worker) AddStep(
	name string,
	invoke activity,
	compensate activity,
) {
	step := func(ctx workflow.Context, param interface{}) (interface{}, error) {
		ctx = NewSaga(ctx, SagaOptions{
			ParallelCompensation: false,
			ContinueWithError:    false,
		})

		AddCompensation(ctx, func(ctx workflow.Context) error {
			if err := workflow.ExecuteActivity(ctx, compensate).Get(ctx, nil); err != nil {
				return err
			}
			return nil
		})

		if err := workflow.ExecuteActivity(ctx, invoke, name).Get(ctx, nil); err != nil {
			Compensate(ctx)
			return nil, err
		}

		return nil, nil
	}

	w.worker.RegisterActivity(invoke)
	w.worker.RegisterActivity(compensate)
	w.worker.RegisterWorkflowWithOptions(step, workflow.RegisterOptions{Name: name})
}

func (w *Worker) AddWorkflow(
	name string,
	flow flow,
) {
	w.worker.RegisterWorkflowWithOptions(flow, workflow.RegisterOptions{Name: name})
}

func (w *Worker) ExecuteStep(
	name string,
	options client.StartWorkflowOptions,
) {
	w.client.ExecuteWorkflow(context.Background(), options, name)
}

func (w *Worker) ExecuteWorker(
	name string,
	options client.StartWorkflowOptions,
) {
	w.client.ExecuteWorkflow(context.Background(), options, name)
}
