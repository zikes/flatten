package flatten

import (
	"fmt"
	"reflect"
	"testing"
)

// TestFlattenInt will test the FlattenInt function and verify that the output
// matches expectations
func TestFlattenInt(t *testing.T) {

	// flattenEquals is a closured test runner, which makes for a more readable
	// table-driven test structure.
	flattenEquals := func(input []interface{}, wantInts []int, wantErr error) func(t *testing.T) {
		return func(t *testing.T) {
			gotInts, gotErr := Int(input)
			if !reflect.DeepEqual(wantErr, gotErr) {
				t.Fatalf("expected: %v, got: %v", wantErr, gotErr)
			}
			if !reflect.DeepEqual(wantInts, gotInts) {
				t.Fatalf("expected: %v, got: %v", wantErr, gotErr)
			}
		}
	}

	// Test a robust number of different nesting variants
	t.Run("not-nested", flattenEquals([]interface{}{1, 2, 3, 4}, []int{1, 2, 3, 4}, nil))
	t.Run("singly-nested", flattenEquals([]interface{}{[]interface{}{1, 2, 3, 4}}, []int{1, 2, 3, 4}, nil))
	t.Run("doubly-nested", flattenEquals([]interface{}{[]interface{}{[]interface{}{1, 2, 3, 4}}}, []int{1, 2, 3, 4}, nil))
	t.Run("adjacent-nested", flattenEquals([]interface{}{[]interface{}{1, 2}, []interface{}{3, 4}}, []int{1, 2, 3, 4}, nil))
	t.Run("adjacent-singly-nested", flattenEquals([]interface{}{[]interface{}{[]interface{}{1}, []interface{}{2}}, []interface{}{3, 4}}, []int{1, 2, 3, 4}, nil))
	t.Run("adjacent-doubly-nested", flattenEquals([]interface{}{[]interface{}{1, 2}, []interface{}{[]interface{}{[]interface{}{3}}, 4}}, []int{1, 2, 3, 4}, nil))

	// Ensure errors are returned for unsupported types
	t.Run("invalid-type-int64", flattenEquals([]interface{}{int64(1), 2, 3, 4}, nil, fmt.Errorf("unsupported element type <int64>")))
	t.Run("invalid-type-string", flattenEquals([]interface{}{"1", 2, 3, 4}, nil, fmt.Errorf("unsupported element type <string>")))
	t.Run("invalid-type-float64", flattenEquals([]interface{}{float64(1), 2, 3, 4}, nil, fmt.Errorf("unsupported element type <float64>")))
	t.Run("invalid-type-uint8", flattenEquals([]interface{}{uint8(1), 2, 3, 4}, nil, fmt.Errorf("unsupported element type <uint8>")))
}
