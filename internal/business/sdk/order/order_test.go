package order

import (
	"errors"
	"testing"
)

func TestOrder_Parse(t *testing.T) {
	t.Parallel()

	mappings := map[string]string{
		"F1": "field1",
		"F2": "field2",
	}
	defaultOrder := NewBy("field1", ASC)

	cases := []struct {
		name    string
		orderBy string
		err     error
		want    By
	}{
		{
			name: "empty order by, should return default order",
			want: defaultOrder,
		},
		{
			name:    "valid order by, should return order",
			orderBy: "F2",
			want:    NewBy("field2", ASC),
		},
		{
			name:    "valid order by with direction, should return order",
			orderBy: "F2,DESC",
			want:    NewBy("field2", DESC),
		},
		{
			name:    "invalid field, should return error",
			orderBy: "F3",
			err:     errors.New("unknown order: F3"),
		},
		{
			name:    "invalid direction, should return error",
			orderBy: "F1,INVALID",
			err:     errors.New("unknown direction: INVALID"),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := Parse(mappings, tc.orderBy, defaultOrder)

			if tc.err != nil && err == nil {
				t.Errorf("expected error, got nil")
			} else if tc.err == nil && err != nil {
				t.Errorf("unexpected error %v", err)
			} else if tc.err != nil && err != nil && err.Error() != tc.err.Error() {
				t.Errorf("got error %v, want %v", err, tc.err)
			}

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
