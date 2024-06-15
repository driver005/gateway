// See https://github.com/Courtsite/temporal-go-helpers

package workflow

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

type BasicReceiveWithTimeoutWorkflowResult struct {
	HasTimedOut bool
	IsCancelled bool
	Message     string
}

func BasicReceiveWithTimeoutWorkflow__WithPayload(ctx workflow.Context) (BasicReceiveWithTimeoutWorkflowResult, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// This is to ensure the signal channel gets cleaned up after timing out.
	childCtx, cancel := workflow.WithCancel(ctx)

	sigCh := workflow.GetSignalChannel(childCtx, "signal-receive-with-timeout")
	var message string
	res := ReceiveWithTimeout(ctx, sigCh, &message, time.Minute*30)

	if res.HasTimedOut {
		cancel()
	}

	var result BasicReceiveWithTimeoutWorkflowResult = BasicReceiveWithTimeoutWorkflowResult{
		HasTimedOut: res.HasTimedOut,
		IsCancelled: res.IsCancelled,
		Message:     message,
	}

	return result, nil
}

func BasicReceiveWithTimeoutWorkflow__NoPayload(ctx workflow.Context) (BasicReceiveWithTimeoutWorkflowResult, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// This is to ensure the signal channel gets cleaned up after timing out.
	childCtx, cancel := workflow.WithCancel(ctx)

	sigCh := workflow.GetSignalChannel(childCtx, "signal-receive-with-timeout")
	res := ReceiveWithTimeout(ctx, sigCh, nil, time.Minute*30)

	if res.HasTimedOut {
		cancel()
	}

	var result BasicReceiveWithTimeoutWorkflowResult = BasicReceiveWithTimeoutWorkflowResult{
		HasTimedOut: res.HasTimedOut,
		IsCancelled: res.IsCancelled,
		Message:     "",
	}

	return result, nil
}

func BasicDrainWorkflow(ctx workflow.Context) (int, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	sigCh := workflow.GetSignalChannel(ctx, "signal")

	totalDrained := 0

	for {
		totalDrained += Drain(sigCh)

		var message string
		sigCh.Receive(ctx, &message)

		if message == "OK" {
			// Do some processing
			workflow.Sleep(ctx, time.Second*5)
		} else if message == "STOP" {
			break
		}
	}

	return totalDrained, nil
}

func BasicSagaWorkflow(ctx workflow.Context, initialAmount int) (int, error) {
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	sagaCtx := NewSaga(ctx, SagaOptions{
		ParallelCompensation: false,
		ContinueWithError:    false,
	})

	ca := &CalculatorActivities{}

	currentAmount := initialAmount
	var compensationOrder []int

	err := workflow.SetQueryHandler(ctx, "currentAmount", func(input []byte) (int, error) {
		return currentAmount, nil
	})
	if err != nil {
		return 0, err
	}

	err = workflow.SetQueryHandler(ctx, "compensationOrder", func(input []byte) ([]int, error) {
		return compensationOrder, nil
	})
	if err != nil {
		return 0, err
	}

	err = workflow.ExecuteActivity(ctx, ca.Add, currentAmount, 10).Get(ctx, &currentAmount)
	if err != nil {
		logger.Error("activity failed", zap.Error(err))
		handleSagaErr(ctx, Compensate(sagaCtx))
		return 0, err
	}
	AddCompensation(sagaCtx, func(ctx workflow.Context) error {
		compensationOrder = append(compensationOrder, 5)
		err := workflow.ExecuteActivity(ctx, ca.Minus, currentAmount, 5).Get(ctx, &currentAmount)
		if err != nil {
			logger.Error("compensation activity failed", zap.Error(err))
			return err
		}
		return nil
	})

	err = workflow.ExecuteActivity(ctx, ca.Add, currentAmount, 20).Get(ctx, &currentAmount)
	if err != nil {
		logger.Error("activity failed", zap.Error(err))
		handleSagaErr(ctx, Compensate(sagaCtx))
		return 0, err
	}
	AddCompensation(sagaCtx, func(ctx workflow.Context) error {
		compensationOrder = append(compensationOrder, 10)
		err := workflow.ExecuteActivity(ctx, ca.Minus, currentAmount, 10).Get(ctx, &currentAmount)
		if err != nil {
			logger.Error("compensation activity failed", zap.Error(err))
			return err
		}
		return nil
	})

	err = workflow.ExecuteActivity(ctx, ca.Add, currentAmount, 30).Get(ctx, &currentAmount)
	if err != nil {
		logger.Error("activity failed", zap.Error(err))
		handleSagaErr(ctx, Compensate(sagaCtx))
		return 0, err
	}

	return currentAmount, nil
}

func MultipleCompensateSagaWorkflow(ctx workflow.Context, initialAmount int) (int, error) {
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	sagaCtx := NewSaga(ctx, SagaOptions{
		ParallelCompensation: false,
		ContinueWithError:    false,
	})

	ca := &CalculatorActivities{}

	currentAmount := initialAmount

	err := workflow.SetQueryHandler(ctx, "currentAmount", func(input []byte) (int, error) {
		return currentAmount, nil
	})
	if err != nil {
		return 0, err
	}

	err = workflow.ExecuteActivity(ctx, ca.Add, currentAmount, 10).Get(ctx, &currentAmount)
	if err != nil {
		logger.Error("activity failed", zap.Error(err))
		handleSagaErr(ctx, Compensate(sagaCtx))
		return 0, err
	}
	AddCompensation(sagaCtx, func(ctx workflow.Context) error {
		err := workflow.ExecuteActivity(ctx, ca.Minus, currentAmount, 5).Get(ctx, &currentAmount)
		if err != nil {
			logger.Error("compensation activity failed", zap.Error(err))
			return err
		}
		return nil
	})

	handleSagaErr(ctx, Compensate(sagaCtx))
	handleSagaErr(ctx, Compensate(sagaCtx))

	return currentAmount, nil
}

type CalculatorActivities struct{}

func (ca *CalculatorActivities) Add(ctx context.Context, a int, b int) (int, error) {
	return a + b, nil
}

func (ca *CalculatorActivities) Minus(ctx context.Context, a int, b int) (int, error) {
	return a - b, nil
}

func handleSagaErr(ctx workflow.Context, err error) {
	logger := workflow.GetLogger(ctx)
	if err != nil {
		logger.Warn("Error(s) in saga compensation", "error", err)
	}
}
