package hashtable

const (
	divisorHash = 100
)

var hash = make([]string, 7)

type HashTable struct {
	table map[int32]string
}

// Hash loop each letter and return ASCII code, divided by divisorHash
// return will be used like indices for the array
func Hash(str string) (hash int32){
	for _, v := range str {
		hash +=v
	}
	// return hash % divisorHash
	return hash / divisorHash
}

func NewHashTable() *HashTable {
	hash := HashTable{}
	for i:=0; i < len(hash.table); i++ {
		hash.table[int32(i)] = ""
	}
	return &hash
}

func (h *HashTable) Add(key string, value string) {
	hashIndex := Hash(key)
	h.table[hashIndex] = value
}

func (h *HashTable) Get(key string) string {
	hashIndex := Hash(key)
	return h.table[hashIndex]
}

func (h *HashTable) Remove(key string) {
	hashIndex := Hash(key)
	h.table[hashIndex] = ""
}

