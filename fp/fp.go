package fp

// Some example usage:
// func Example() {
//     numbers := []int{1, 2, 3, 4, 5}

//     // Map example: double all numbers
//     doubled := Map(numbers, func(x int) int {
//         return x * 2
//     })
//     // doubled = [2, 4, 6, 8, 10]

//     // Filter example: keep even numbers
//     evens := Filter(numbers, func(x int) bool {
//         return x%2 == 0
//     })
//     // evens = [2, 4]

//     // Reduce example: sum all numbers
//     sum := Reduce(numbers, 0, func(acc, curr int) int {
//         return acc + curr
//     })
//     // sum = 15

//     // Compose example
//     addOne := func(x int) int { return x + 1 }
//     double := func(x int) int { return x * 2 }
//     addOneThenDouble := Compose(double, addOne)
//     // addOneThenDouble(3) = 8

//     // Curry example
//     add := func(x, y int) int { return x + y }
//     addCurried := Curry2(add)
//     addFive := addCurried(5)
//     // addFive(3) = 8
// }

// Map applies function f to each element in slice
func Map[T, U any](slice []T, f func(T) U) []U {
	mapped := make([]U, len(slice))
	for i, elem := range slice {
		mapped[i] = f(elem)
	}
	return mapped
}

// Filter returns slice of elements that satisfy predicate f
func Filter[T any](slice []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, elem := range slice {
		if f(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

// Reduce accumulates result by applying f to each element
func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
	result := initial
	for _, elem := range slice {
		result = f(result, elem)
	}
	return result
}

// Compose creates a new function that is the composition of f and g
func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	return func(a A) C {
		return f(g(a))
	}
}

// Curry2 converts a function of 2 arguments into a series of 1 argument functions
func Curry2[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

//

// Option type for handling optional values
type Option[T any] struct {
	value T
	valid bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, valid: true}
}

func None[T any]() Option[T] {
	return Option[T]{valid: false}
}

func (o Option[T]) GetOrElse(defaultValue T) T {
	if !o.valid {
		return defaultValue
	}
	return o.value
}
