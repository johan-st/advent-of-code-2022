package search

import (
	"reflect"
	"testing"
)

func Test_graph_getNeighboursOf(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		g       graph
		args    args
		want    []int
		wantErr bool
	}{
		{"error on invalid id", graph{}, args{1}, []int{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.getNeighboursOf(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("graph.getNeighboursOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("graph.getNeighboursOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func exampleGraph() graph {
	return graph{}
}
