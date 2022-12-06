package main

import "testing"

func Test_zonesFullyContained(t *testing.T) {
	type args struct {
		z1 zones
		z2 zones
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"no overlap", args{zones{0, 10}, zones{11, 20}}, false},
		{"overlap but not contained", args{zones{0, 10}, zones{5, 15}}, false},
		{"contained (left)", args{zones{5, 15}, zones{0, 20}}, true},
		{"contained (right)", args{zones{0, 20}, zones{20, 20}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zonesFullyContained(tt.args.z1, tt.args.z2); got != tt.want {
				t.Errorf("zonesFullyContained() = %v, want %v", got, tt.want)
			}
		})
	}
}
