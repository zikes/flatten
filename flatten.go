package flatten

import "fmt"

// Int flattens an arbitrarily nested array of integers into a flat array
// of integers.
func Int(input []interface{}) ([]int, error) {
	output := []int{}

	for _, val := range input {
		// Check if element is an int, append if so
		if i, ok := val.(int); ok {
			output = append(output, i)
			continue
		}

		// Check if element is an array, recursively flatten and append if so
		if arr, ok := val.([]interface{}); ok {
			ints, err := Int(arr)
			if err != nil {
				return nil, err
			}
			output = append(output, ints...)
			continue
		}
		return nil, fmt.Errorf("unsupported element type <%T>", val)
	}

	return output, nil
}
