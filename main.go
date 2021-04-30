package main

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func main() {
	sel, err := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app": "nginx",
		},
		MatchExpressions: nil,
	})
	if err != nil {
		panic(err)
	}
	reqs := SelectorToArray(sel)
	fmt.Println(reqs)
	fmt.Println(sel)
}

func printJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func SelectorToArray(sel labels.Selector) []string {
	requirements, selectable := sel.Requirements()
	if !selectable {
		fmt.Printf("sel= %v is NOT selectable", sel)
	}
	reqs := make([]string, 0, len(requirements))
	for _, r := range requirements {
		reqs = append(reqs, r.String())
	}
	return reqs
}
