/**
  For reference: https://cp-algorithms.com/data_structures/fenwick.html
  This is an implementation for 1-Indexed fenwick tree
*/

package Fenwick

type FenwickTree struct {
    bit []int  // Segment tree
    nums []int // Original array
}


func Constructor(nums []int) FenwickTree {
    /*
      Constructs a fenwick tree in O(n) complexity
    */
    bit:=make([]int, len(nums)+1)
    for i:=range nums{
        j:=i+1
        bit[j]+=nums[i]
        r:=j+(j&-j)
        if r<=len(nums) {
            bit[r]+=bit[i+1]
        }
    }
    return FenwickTree{bit, nums}
}


func (this *FenwickTree) Update(index int, val int)  {
    /*
      Updates the sums in the fenwick tree by propagating delta. 
      Complexity: O(logn)
    */
    delta:=val-this.nums[index]
    this.nums[index]=val
    for i:=index+1;i<len(this.bit);i+=i&-i{
        this.bit[i]+=delta
    }
}

func (this *FenwickTree) Sum(index int) int {
    /*
      Returns sum from fenwick tree
      Complexity: O(logn)
    */
    ans:=0
    for i:=index+1;i>0;i-=i&-i{
        ans+=this.bit[i]
    }
    return ans
}

func (this *FenwickTree) SumRange(left int, right int) int {
    /*
      Returns range sum from fenwick tree
      Complexity: O(logn)
    */
    return this.Sum(right)-this.Sum(left)+this.nums[left]
}
