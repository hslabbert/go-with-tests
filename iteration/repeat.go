package iteration

// Repeat takes a single character and repeats it 'count' times
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
