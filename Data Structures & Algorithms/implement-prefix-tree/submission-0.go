type PrefixTree struct {
	Node [26]*PrefixTree
	IsEndOfWord bool
}

func Constructor() *PrefixTree {
	return &PrefixTree{}
    
}

func (this *PrefixTree) Insert(word string) {
	for _, c := range word {
		index := c - 'a'

		if this.Node[index] == nil {
			this.Node[index] = &PrefixTree{}
		}

		this = this.Node[index] 
	}

	this.IsEndOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	for _, c := range word {
		index := c - 'a'
		if this.Node[index] == nil {
			return false
		}

		this = this.Node[index]
	}

	return this.IsEndOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	for _, c := range prefix {
		index := c - 'a'
		if this.Node[index] == nil {
			return false
		}
		this = this.Node[index]
	}

	return true
}
