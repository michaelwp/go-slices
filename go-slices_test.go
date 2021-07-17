package go_slices

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("=== Testing Start ===")

	code := m.Run()

	log.Println("=== Testing End ===")
	os.Exit(code)
}

func TestSlices_Shift(t *testing.T) {
	tests := []struct {
		name           string
		data           Slices
		expectedResult Slices
	}{
		{
			name:           "Slices{1,2,3,4,5,6}.Shift()",
			data:           Slices{1, 2, 3, 4, 5, 6},
			expectedResult: Slices{2, 3, 4, 5, 6},
		},
		{
			name:           "Slices{'a','b','c','d','e','f'}.Shift()",
			data:           Slices{"a", "b", "c", "d", "e", "f"},
			expectedResult: Slices{"b", "c", "d", "e", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actualResult Slices

			if actualResult = tt.data.Shift(); !reflect.DeepEqual(actualResult, tt.expectedResult) {
				t.Errorf("actual result = %v, expected result %v", actualResult, tt.expectedResult)
			}
		})
	}
}

func TestSlices_Pop(t *testing.T) {
	tests := []struct {
		name           string
		data           Slices
		expectedResult Slices
	}{
		{
			name:           "Slices{1,2,3,4,5,6}.Pop()",
			data:           Slices{1, 2, 3, 4, 5, 6},
			expectedResult: Slices{1, 2, 3, 4, 5},
		},
		{
			name:           "Slices{'a','b','c','d','e','f'}.Pop()",
			data:           Slices{"a", "b", "c", "d", "e", "f"},
			expectedResult: Slices{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actualResult Slices

			if actualResult = tt.data.Pop(); !reflect.DeepEqual(actualResult, tt.expectedResult) {
				t.Errorf("actual result = %v, expected result %v", actualResult, tt.expectedResult)
			}
		})
	}
}

func TestSlices_Remove(t *testing.T) {
	tests := []struct {
		name            string
		data            Slices
		elementToRemove int
		expectedResult  Slices
	}{
		{
			name:            "Slices{1,2,3,4,5,6}.Remove(0)",
			data:            Slices{1, 2, 3, 4, 5, 6},
			elementToRemove: 0,
			expectedResult:  Slices{2, 3, 4, 5, 6},
		},
		{
			name:            "Slices{1,2,3,4,5,6}.Remove(5)",
			data:            Slices{1, 2, 3, 4, 5, 6},
			elementToRemove: 5,
			expectedResult:  Slices{1, 2, 3, 4, 5},
		},
		{
			name:            "Slices{1,2,3,4,5,6}.Remove(3)",
			data:            Slices{1, 2, 3, 4, 5, 6},
			elementToRemove: 3,
			expectedResult:  Slices{1, 2, 3, 5, 6},
		},
		{
			name:            "Slices{'a','b','c','d','e','f'}.Remove(0)",
			data:            Slices{"a", "b", "c", "d", "e", "f"},
			elementToRemove: 0,
			expectedResult:  Slices{"b", "c", "d", "e", "f"},
		},
		{
			name:            "Slices{'a','b','c','d','e','f'}.Remove(5)",
			data:            Slices{"a", "b", "c", "d", "e", "f"},
			elementToRemove: 5,
			expectedResult:  Slices{"a", "b", "c", "d", "e"},
		},
		{
			name:            "Slices{'a','b','c','d','e','f'}.Remove(2)",
			data:            Slices{"a", "b", "c", "d", "e", "f"},
			elementToRemove: 2,
			expectedResult:  Slices{"a", "b", "d", "e", "f"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actualResult Slices

			if actualResult = tt.data.Remove(tt.elementToRemove); !reflect.DeepEqual(actualResult, tt.expectedResult) {
				t.Errorf("actual result = %v, expected result %v", actualResult, tt.expectedResult)
			}
		})
	}
}

func BenchmarkSlices(b *testing.B) {
	slice := Slices{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	b.Run("slice Shift()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice.Shift()
		}
	})

	b.Run("slice Pop()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice.Pop()
		}
	})

	b.Run("slice Remove(e)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice.Remove(0)
		}
	})
}
