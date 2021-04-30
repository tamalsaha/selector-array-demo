package main

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"os"
)

func main() {
	sel, err := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app":  "nginx",
			"tier": "frontend",
		},
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      "node/region",
				Operator: metav1.LabelSelectorOpIn,
				Values:   []string{"us-east-1a", "us-east-1b"},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	reqs := SelectorToArray(sel)
	printJSON(reqs)
	fmt.Println(sel)

	fmt.Println("Everything =================")

	selEverything := labels.Everything()
	reqsEverything := SelectorToArray(selEverything)
	printJSON(reqsEverything)
	fmt.Println(selEverything)

	fmt.Println("Nothing =================")

	selNothing := labels.Nothing()
	reqsNothing := SelectorToArray(selNothing)
	printJSON(reqsNothing)
	fmt.Println(selNothing)
}

func printJSON(v interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(v)
	if err != nil {
		panic(err)
	}
}

func SelectorToArray(sel labels.Selector) []string {
	requirements, selectable := sel.Requirements()
	if !selectable {
		return []string{"<NONE>"}
	}
	reqs := make([]string, 0, len(requirements))
	for _, r := range requirements {
		reqs = append(reqs, r.String())
	}
	if len(reqs) == 0 {
		reqs = []string{"<ALL>"}
	}
	return reqs
}
