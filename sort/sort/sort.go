// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run genzfunc.go

// Package sort provides primitives 原始的 for sorting slices and user-defined
// collections 集合、收集.
package sort

import "reflect"

// A type, typically 代表性的 a collection, that satisfies 满足 sort.Interface can be
// sorted by the routines 惯例 in this package. The methods require that the
// elements of the collection be enumerated 枚举 by an integer 整数 index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps 交换 the elements with indexes i and j.
	Swap(i, j int)
}

// Insertion 插入 sort
func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
			// 之前以为此处应该有一个 break ，其实 在for条件里 j>a && data.Less(j, j-1)这里就已经
			// 将不满足条件的已排序好的过滤掉了，完全不会执行
		}
	}
}

// siftDown implements the heap propertyw 性质 on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top. (h1 - 1) / 2 取得最后一个非叶子节点
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data. // 将堆顶元素与堆末尾元素置换，然后重新调整无序部位为堆
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}

// Quicksort, loosely 宽松的 following Bentley and McIlroy,
// ``Engineering 建造 a Sort Function,'' SP&E November 1993.

// medianOfThree moves the median 中央 of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(data Interface, m1, m0, m2 int) {
	// sort 3 elements
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	// data[m0] <= data[m1]
	if data.Less(m2, m1) {
		data.Swap(m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if data.Less(m1, m0) {
			data.Swap(m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func swapRange(data Interface, a, b, n int) {
	for i := 0; i < n; i++ {
		data.Swap(a+i, b+i)
	}
}

// pivot 中心点 求lo 跟 hi 的 中心点  并将其左右两边的数据以某一值的大小进行排序分离
func doPivot(data Interface, lo, hi int) (midlo, midhi int) {
	// 计算中间点位置
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.

	// 当需要处理的元素大于40个的时候， 想办法取一个更靠谱点的中间值
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.

		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+2*s)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	// 将中间元素的值交换为 lo位置， m位置， hi - 1 位置的中间值即排序
	medianOfThree(data, lo, m, hi-1)

	// Invariants 约束条件 are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined 未经检查的
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && data.Less(a, pivot); a++ {  // data[a] < data[pivot]
	}
	b := a
	for {
		for ; b < c && !data.Less(pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && data.Less(pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		data.Swap(b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates 副本 (by property of median of nine).
	// Let be a bit more conservative 保守派, and set border 边界 to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot 中心点
		dups := 0
		if !data.Less(pivot, hi-1) { // data[hi-1] = pivot
			data.Swap(c, hi-1)
			c++
			dups++
		}
		if !data.Less(b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !data.Less(m, pivot) { // data[m] = pivot
			data.Swap(m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed 歪斜的 distribution 分布
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant 约束条件:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !data.Less(b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && data.Less(a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			data.Swap(a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	data.Swap(pivot, b-1)
	return b - 1, c
}

// 快速排序算法，思想，一趟排序将待排序部分分成两个部分使得左部分总比另一部分小，然后对两部分再次进行快速排序，递归进行直到排序完成
func quickSort(data Interface, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(data, a, b)
		// Avoiding recursion 递归 on the larger subproblem guarantees 保证 // 避免很多的递归
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(data, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort(data, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort 希尔排序 pass with gap 间隔 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if data.Less(i, i-6) {
				data.Swap(i, i-6)
			}
		}
		insertionSort(data, a, b)
	}
}

// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}

// maxDepth returns a threshold 极限 at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)). 进一取整
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

// lessSwap is a pair of Less and Swap function for use with the
// auto-generated func-optimized 最佳化 variant 变体 of sort.go in
// zfuncversion.go.
type lessSwap struct {
	Less func(i, j int) bool
	Swap func(i, j int)
}

// Slice sorts the provided slice given the provided less function.
//
// The sort is not guaranteed 担保 to be stable 稳定的. For a stable sort, use
// SliceStable.
//
// The function panics if the provided interface is not a slice.
func Slice(slice interface{}, less func(i, j int) bool) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	quickSort_func(lessSwap{less, swap}, 0, length, maxDepth(length))
}

// SliceStable sorts the provided slice given the provided less
// function while keeping the original 原始的 order of equal elements.
//
// The function panics if the provided interface is not a slice.
func SliceStable(slice interface{}, less func(i, j int) bool) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	stable_func(lessSwap{less, swap}, rv.Len())
}

// SliceIsSorted tests whether a slice is sorted.
//
// The function panics if the provided interface is not a slice.
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool {
	rv := reflect.ValueOf(slice)
	n := rv.Len()
	for i := n - 1; i > 0; i-- {
		if less(i, i-1) {
			return false
		}
	}
	return true
}

type reverse struct {
	// This embedded 嵌入 Interface permits 许可 Reverse 倒转 to use the methods of
	// another Interface implementation.
	Interface
}

// Less returns the opposite 相反的 of the embedded implementation's Less method.
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
	return &reverse{data}
}

// IsSorted reports whether data is sorted.
func IsSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// Convenience 便利店 types for common cases

// IntSlice attaches 附加 the methods of Interface to []int, sorting in increasing 渐增的 order.
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p IntSlice) Sort() { Sort(p) }

// Float64Slice attaches the methods of Interface to []float64, sorting in increasing order
// (not-a-number values are treated as less than other values).
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// isNaN is a copy of math.IsNaN to avoid a dependency 从属 on the math package.
func isNaN(f float64) bool {
	return f != f
}

// Sort is a convenience method.
func (p Float64Slice) Sort() { Sort(p) }

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p StringSlice) Sort() { Sort(p) }

// Convenience wrappers for common cases

// Ints sorts a slice of ints in increasing order.
func Ints(a []int) { Sort(IntSlice(a)) }

// Float64s sorts a slice of float64s in increasing order
// (not-a-number values are treated as less than other values).
func Float64s(a []float64) { Sort(Float64Slice(a)) }

// Strings sorts a slice of strings in increasing order.
func Strings(a []string) { Sort(StringSlice(a)) }

// IntsAreSorted tests whether a slice of ints is sorted in increasing order.
func IntsAreSorted(a []int) bool { return IsSorted(IntSlice(a)) }

// Float64sAreSorted tests whether a slice of float64s is sorted in increasing order
// (not-a-number values are treated as less than other values).
func Float64sAreSorted(a []float64) bool { return IsSorted(Float64Slice(a)) }

// StringsAreSorted tests whether a slice of strings is sorted in increasing order.
func StringsAreSorted(a []string) bool { return IsSorted(StringSlice(a)) }

// Notes on stable 稳定的 sorting:
// The used algorithms are simple and provable 可证明的 correct 正确 on all input and use
// only logarithmic 对数 additional stack space. They perform well if compared
// experimentally 实验性的 to other stable in-place sorting algorithms.
//
// Remarks 备注 on other algorithms 算法 evaluated 评估:
//  - GCC's 4.6.3 stable_sort with merge_without_buffer from libstdc++:
//    Not faster.
//  - GCC's __rotate 旋转 for block rotations: Not faster.
//  - "Practical in-place mergesort" from  Jyrki Katajainen, Tomi A. Pasanen
//    and Jukka Teuhola; Nordic Journal of Computing 3,1 (1996), 27-40:
//    The given algorithms are in-place, number of Swap and Assignments
//    grow as n log n but the algorithm is not stable.
//  - "Fast Stable In-Place Sorting with O(n) Data Moves" J.I. Munro and
//    V. Raman in Algorithmica (1996) 16, 115-160:
//    This algorithm either needs additional 2n bits or works only if there
//    are enough different elements available to encode some permutations
//    which have to be undone later (so not stable on any input).
//  - All the optimal in-place sorting/merging algorithms I found are either
//    unstable or rely on enough different elements in each step to encode the
//    performed block rearrangements. See also "In-Place Merging Algorithms",
//    Denham Coates-Evely, Department of Computer Science, Kings College,
//    January 2004 and the references in there.
//  - Often "optimal" algorithms are optimal in the number of assignments
//    but Interface has only Swap as operation.

// Stable sorts data while keeping the original order of equal elements.
//
// It makes one call to data.Len to determine n, O(n*log(n)) calls to
// data.Less and O(n*log(n)*log(n)) calls to data.Swap.
func Stable(data Interface) {
	stable(data, data.Len())
}

func stable(data Interface, n int) {
	blockSize := 20 // must be > 0
	a, b := 0, blockSize
	for b <= n {
		insertionSort(data, a, b)
		a = b
		b += blockSize
	}
	insertionSort(data, a, n)

	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			symMerge(data, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			symMerge(data, a, m, n)
		}
		blockSize *= 2
	}
}

// SymMerge merges 合并 the two sorted subsequences 子序列 data[a:m] and data[m:b] using
// the SymMerge algorithm from Pok-Son Kim and Arne Kutzner, "Stable Minimum 最小值
// Storage 存储 Merging by Symmetric 对称 Comparisons 比较", in Susanne Albers and Tomasz
// Radzik, editors, Algorithms - ESA 2004, volume 量、卷 3221 of Lecture 讲稿 Notes 注释 in
// Computer Science, pages 714-723. Springer 蹦跳的人, 2004.
//
// Let M = m-a and N = b-n. Wolog M < N.
// The recursion 递归 depth is bound 范围 by ceil(log(N+M)). ceil 进一取整
// The algorithm needs O(M*log(N/M + 1)) calls to data.Less.
// The algorithm needs O((M+N)*log(M)) calls to data.Swap.
//
// The paper 论文 gives O((M+N)*log(M)) as the number of assignments 作业 assuming 假设 a
// rotation 轮转 algorithm which uses O(M+N+gcd(M+N)) assignments. The argumentation 论证
// in the paper carries through for Swap operations, especially as the block
// swapping rotate uses only O(M+N) Swaps.
//
// symMerge assumes non-degenerate 退化的 arguments 参数: a < m && m < b.
// Having the caller check this condition eliminates 消除 many leaf 叶子 recursion 递归 calls,
// which improves performance 性能.
func symMerge(data Interface, a, m, b int) {
	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[a] into data[m:b]
	// if data[a:m] only contains one element.
	if m-a == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] >= data[a] for m <= i < b.
		// Exit the search loop with i == b in case no such index exists.
		i := m
		j := b
		for i < j {
			h := int(uint(i+j) >> 1)
			if data.Less(h, a) {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[a] reaches the position before i.
		for k := a; k < i-1; k++ {
			data.Swap(k, k+1)
		}
		return
	}

	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[m] into data[a:m]
	// if data[m:b] only contains one element.
	if b-m == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] > data[m] for a <= i < m.
		// Exit the search loop with i == m in case no such index exists.
		i := a
		j := m
		for i < j {
			h := int(uint(i+j) >> 1)
			if !data.Less(m, h) {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[m] reaches the position i.
		for k := m; k > i; k-- {
			data.Swap(k, k-1)
		}
		return
	}

	mid := int(uint(a+b) >> 1)
	n := mid + m
	var start, r int
	if m > mid {
		start = n - b
		r = mid
	} else {
		start = a
		r = m
	}
	p := n - 1

	for start < r {
		c := int(uint(start+r) >> 1)
		if !data.Less(p-c, c) {
			start = c + 1
		} else {
			r = c
		}
	}

	end := n - start
	if start < m && m < end {
		rotate(data, start, m, end)
	}
	if a < start && start < mid {
		symMerge(data, a, start, mid)
	}
	if mid < end && end < b {
		symMerge(data, mid, end, b)
	}
}

// Rotate two consecutives 连续 blocks u = data[a:m] and v = data[m:b] in data:
// Data of the form 'x u v y' is changed to 'x v u y'.
// Rotate performs at most b-a many calls to data.Swap.
// Rotate assumes non-degenerate 退化 arguments: a < m && m < b.
func rotate(data Interface, a, m, b int) {
	i := m - a
	j := b - m

	for i != j {
		if i > j {
			swapRange(data, m-i, m, j)
			i -= j
		} else {
			swapRange(data, m-i, m+j-i, i)
			j -= i
		}
	}
	// i == j
	swapRange(data, m-i, m, i)
}

/*
Complexity 复杂性 of Stable Sorting


Complexity of block swapping rotation 轮转、循环、轮流

Each Swap puts one new element into its correct, final position.
Elements which reach their final position are no longer moved.
Thus block swapping rotation needs |u|+|v| calls to Swaps.
This is best possible as each element might need a move.

Pay attention when comparing to other optimal 最佳 algorithms which
typically count the number of assignments 作业 instead of swaps:
E.g. the optimal algorithm of Dudzinski and Dydek for in-place
rotations uses O(u + v + gcd(u,v)) assignments which is
better than our O(3 * (u+v)) as gcd(u,v) <= u.


Stable sorting by SymMerge and BlockSwap rotations

SymMerg complexity for same size input M = N:
Calls to Less:  O(M*log(N/M+1)) = O(N*log(2)) = O(N)
Calls to Swap:  O((M+N)*log(M)) = O(2*N*log(N)) = O(N*log(N))

(The following argument does not fuzz 柔毛、模糊 over a missing -1 or
other stuff 东西、材料 which does not impact 影响 the final result).

Let n = data.Len(). Assume n = 2^k.

Plain 简单的 merge sort performs 执行 log(n) = k iterations 迭代.
On iteration i the algorithm merges 2^(k-i) blocks, each of size 2^i.

Thus iteration i of merge sort performs:
Calls to Less  O(2^(k-i) * 2^i) = O(2^k) = O(2^log(n)) = O(n)
Calls to Swap  O(2^(k-i) * 2^i * log(2^i)) = O(2^k * i) = O(n*i)

In total k = log(n) iterations are performed; so in total:
Calls to Less O(log(n) * n)
Calls to Swap O(n + 2*n + 3*n + ... + (k-1)*n + k*n)
   = O((k/2) * k * n) = O(n * k^2) = O(n * log^2(n))


Above results should generalize 概括 to arbitrary 任意的 n = 2^k + p
and should not be influenced by the initial insertion sort phase 月相、相位:
Insertion sort is O(n^2) on Swap and Less, thus O(bs^2) per block of
size bs at n/bs blocks:  O(bs*n) Swaps and Less during insertion sort.
Merge sort iterations start at i = log(bs). With t = log(bs) constant 常量、恒量:
Calls to Less O((log(n)-t) * n + bs*n) = O(log(n)*n + (bs-t)*n)
   = O(n * log(n))
Calls to Swap O(n * log^2(n) - (t^2+t)/2*n) = O(n * log^2(n))

*/
