package workflow

import (
	"errors"
	"fmt"

	flow "github.com/s8sg/goflow/flow/v1"
	"github.com/s8sg/goflow/operation"
	goflow "github.com/s8sg/goflow/v1"
)

const (
	PASS  = "pass"
	ERROR = "error"
)

type Workflow struct {
	Workflow *goflow.FlowService
	Dag      *flow.Dag
	// context  *flow.Context
	Name string
}

func NewWorkflow() *Workflow {
	return &Workflow{
		Workflow: &goflow.FlowService{},
		Dag:      flow.NewDag(),
		Name:     "workflow",
	}
}

func (w *Workflow) SetContext(ctx *flow.Context) {

}

func (w *Workflow) Create() {
	// fs := &goflow.FlowService{}
	// fs.Register(w.Name, func(flow *flow.Workflow, context *flow.Context) error {

	// })
}

// Workload function to handle case1
func handleCase1(data []byte, option map[string][]string) ([]byte, error) {
	result := fmt.Sprintf("(Executing case 1 with data (%s))", string(data))
	fmt.Println(result)
	return nil, errors.New("test")
}

// Workload function to handle case 2
func handleCase2(data []byte, option map[string][]string) ([]byte, error) {
	result := fmt.Sprintf("(Executing case 2 with data (%s))", string(data))
	fmt.Println(result)
	return []byte(result), nil
}

func DefineWorkflow(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	branches := dag.ConditionalBranch(
		"handel-workflow",
		[]string{PASS, ERROR},
		conditionCheck,
		flow.Aggregator(conditionAggregator),
	)
	branches[PASS].Node("invoke", handleCase1)
	branches[ERROR].Node("compensate", handleCase2)
	// w.Dag = dag
	return nil
}

func (w *Workflow) GetWorkflow() {
	fs := &goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		OpenTraceUrl:      "localhost:5775",
		WorkerConcurrency: 5,
		EnableMonitoring:  true,
	}
	fmt.Println(fs.Register("myflow", DefineWorkflow))
	fmt.Println(fs.Start())
	// fmt.Println(fs.StartWorker())
	fmt.Println(fs.Execute("myflow", &goflow.Request{
		Body: []byte("hallo"),
	}))
}

func conditionCheck(response []byte) []string {
	fmt.Println(string(response))
	return []string{PASS}
}

func conditionAggregator(data map[string][]byte) ([]byte, error) {
	// case1Data, _ := data["case1"]
	// case2Data, _ := data["case2"]
	// aggregatedResult := fmt.Sprintf("(case1: %s, case2: %s)", case1Data, case2Data)
	aggregatedResult := "aggregatedResult"
	return []byte(aggregatedResult), nil
}

func (w *Workflow) Step(
	invoke operation.Modifier,
	compensate operation.Modifier,
) {
	// dag := w.Workflow.Dag()
	// branches := dag.ConditionalBranch(
	// 	"handel-workflow",
	// 	[]string{PASS, ERROR},
	// 	w.conditionCheck,
	// 	flow.Aggregator(conditionAggregator),
	// )
	// branches[PASS].Node("invoke", invoke)
	// branches[ERROR].Node("compensate", compensate)
	// w.Dag = dag
}
