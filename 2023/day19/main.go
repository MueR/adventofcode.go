package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MueR/adventofcode.go/cast"
)

var (
	//go:embed input.txt
	input string
)

type (
	Condition    int
	Result       int
	Part         map[string]int
	stepFn       func(Part) (bool, Result, string)
	WorkflowStep struct {
		f         stepFn
		result    Result
		condition Condition
		condVar   string
		condVal   int
		next      string
	}
)

const (
	StrAccepted = "A"
	StrRejected = "R"
)

const (
	Accepted Result = iota
	Rejected
	Undecided
)

const (
	AlwaysTrue Condition = iota
	LessThan
	GreaterThan
)

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()

	s := time.Now()
	if part != 2 {
		ans := part1(input)
		fmt.Printf("Part 1 output: %v  (%v)\n", ans, time.Since(s))
	}
	s = time.Now()
	if part != 1 {
		ans := part2(input)
		fmt.Printf("Part 2 output: %v  (%v)\n", ans, time.Since(s))
	}
}

func part1(input string) int {
	workflows, parts := parseInput(input)
	res := 0

nextPart:
	for _, part := range parts {
		workflow := "in"
	flow:
		for {
			steps := workflows[workflow]
			for _, step := range steps {
				ok, result, next := step.f(part)
				if !ok {
					continue
				}
				switch result {
				case Accepted:
					res += part["x"] + part["m"] + part["a"] + part["s"]
					continue nextPart
				case Rejected:
					continue nextPart
				case Undecided:
					workflow = next
					continue flow
				default:
					m := fmt.Sprintf("invalid result: %v", result)
					panic(m)
				}
			}
		}
	}
	return res
}

func part2(input string) (res int) {
	workflows, _ := parseInput(input)
	node := buildDfsTree(workflows, "in")
	defaultRange := Range{from: 1, to: 4000}
	intervals := dfs(node, Rating{
		"x": defaultRange,
		"m": defaultRange,
		"a": defaultRange,
		"s": defaultRange,
	})
	for _, interval := range intervals {
		v := 1
		for _, rating := range interval {
			v *= rating.to - rating.from + 1
		}
		res += v
	}

	return res
}

func parseInput(input string) (map[string][]WorkflowStep, []Part) {
	text := strings.Split(input, "\n\n")
	workflows := parseWorkflows(text[0])
	parts := parseRatings(text[1])
	return workflows, parts
}

func parseWorkflows(s string) map[string][]WorkflowStep {
	wfs := make(map[string][]WorkflowStep)
	for _, line := range strings.Split(s, "\n") {
		start := strings.Index(line, "{")
		name := line[:start]
		var flows []WorkflowStep
		for _, flow := range strings.Split(line[start+1:len(line)-1], ",") {
			step := parseStep(flow)
			flows = append(flows, step)
		}
		wfs[name] = flows
	}
	return wfs
}

func parseStep(s string) WorkflowStep {
	var res Result
	var next string
	nextIdx := strings.Index(s, ":")
	if nextIdx == -1 {
		switch s {
		case StrAccepted:
			res = Accepted
		case StrRejected:
			res = Rejected
		default:
			res = Undecided
			next = s
		}
		return WorkflowStep{
			f: func(_ Part) (bool, Result, string) {
				return true, res, next
			},
			result:    res,
			condition: AlwaysTrue,
			next:      next,
		}
	}

	next = s[nextIdx+1:]
	before := s[:nextIdx]
	condIdx := strings.Index(before, "<")
	condOp := LessThan
	if condIdx == -1 {
		condIdx = strings.Index(before, ">")
		condOp = GreaterThan
	}
	if condIdx == -1 {
		panic("invalid condition")
	}
	condVar := before[:condIdx]
	condVal := cast.ToInt(before[condIdx+1:])
	switch next {
	case StrAccepted:
		res = Accepted
	case StrRejected:
		res = Rejected
	default:
		res = Undecided
	}
	if condOp == LessThan {
		return WorkflowStep{
			f: func(p Part) (bool, Result, string) {
				if p[condVar] < condVal {
					return true, res, next
				}
				return false, res, ""
			},
			result:    res,
			condition: condOp,
			condVal:   condVal,
			condVar:   condVar,
			next:      next,
		}
	}
	return WorkflowStep{
		f: func(p Part) (bool, Result, string) {
			if p[condVar] > condVal {
				return true, res, next
			}
			return false, res, ""
		},
		result:    res,
		condition: condOp,
		condVal:   condVal,
		condVar:   condVar,
		next:      next,
	}
}

func parseRatings(input string) (parts []Part) {
	for _, line := range strings.Split(input, "\n") {
		part := make(Part)
		for _, s := range strings.Split(line[1:len(line)-1], ",") {
			part[s[0:1]] = cast.ToInt(s[2:])
		}
		parts = append(parts, part)
	}
	return parts
}

type Node struct {
	accepted, rejected bool
	name               string
	children           []*Node
	steps              []WorkflowStep
}

type Range struct {
	from, to int
}
type Rating map[string]Range

func (r Rating) copy() Rating {
	res := make(map[string]Range, len(r))
	for k, v := range r {
		res[k] = v
	}
	return res
}

func buildDfsTree(workflows map[string][]WorkflowStep, name string) *Node {
	var children []*Node
	var steps []WorkflowStep
	for _, step := range workflows[name] {
		switch step.result {
		case Accepted:
			children = append(children, &Node{accepted: true})
			steps = append(steps, step)
			continue
		case Rejected:
			children = append(children, &Node{rejected: true})
			steps = append(steps, step)
			continue
		default:
			// do nothing
		}

		next := buildDfsTree(workflows, step.next)
		if next != nil {
			children = append(children, next)
			steps = append(steps, step)
		}
	}
	return &Node{
		name:     name,
		children: children,
		steps:    steps,
	}
}

func dfs(node *Node, r Rating) []Rating {
	if node.rejected {
		return nil
	}
	if node.accepted {
		for _, val := range r {
			if val.to < val.from {
				return nil
			}
		}
		return []Rating{r}
	}
	parent := r.copy()
	var res []Rating
	for i, child := range node.children {
		current := parent.copy()
		step := node.steps[i]
		v := step.condVar
		rr := current[v]
		switch step.condition {
		case AlwaysTrue:
			res = append(res, dfs(child, current)...)
		case LessThan:
			rr.to = min(rr.to, step.condVal-1)
			current[v] = rr
			res = append(res, dfs(child, current)...)

			rr = parent[v]
			rr.from = max(rr.from, step.condVal)
			parent[v] = rr
		case GreaterThan:
			rr.from = max(rr.from, step.condVal+1)
			current[v] = rr
			res = append(res, dfs(child, current)...)

			rr = parent[v]
			rr.to = min(rr.to, step.condVal)
			parent[v] = rr
		}
	}
	return res
}
