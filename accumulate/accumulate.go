package accumulate

// Accumulate maps over the given data
func Accumulate(input []string, f func(string) string) []string {
	output := make([]string, len(input))
	for index, value := range input {
		output[index] = f(value)
	}
	return output
}
