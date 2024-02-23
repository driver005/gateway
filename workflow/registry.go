package workflow

import (
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
	return []byte(result), nil
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
	w.Workflow.Register(w.Name, DefineWorkflow)
	w.Workflow.Start()
}

func conditionCheck(response []byte) []string {
	fmt.Println(response)
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
