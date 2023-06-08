package max_dot_product

import "testing"

func TestMaxDotProduct(t *testing.T) {
	type args struct {
		clicks []int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				clicks: []int{23},
				prices: []int{39},
			},
			want: 897,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := max_dot_product(test.args.clicks, test.args.prices); got != test.want {
				t.Errorf("maxDotProduct() = %v, want %v", got, test.want)
			}
		})
	}
}
