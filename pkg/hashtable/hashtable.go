package hashtable

const (
	divisorHash = 100
)

// Hash loop each letter and return ASCII code, divided by divisorHash
// return will be used like indices for the array
func Hash(str string) (hash int32){
	for _, v := range str {
		hash +=v
	}
	// return hash % divisorHash
	return hash / divisorHash
}