type KthLargest struct {
    k int
	arr []int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest{k: k, arr: nums}
}


func (this *KthLargest) Add(val int) int {
    this.arr = append(this.arr, val)
	sort.Ints(this.arr)
	return this.arr[len(this.arr)-this.k]
}