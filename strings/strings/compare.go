// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

// Compare returns an integer comparing two strings lexicographically 字典序.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
//
// Compare is included only for symmetry 对称 with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, <, >, and so on.
func Compare(a, b string) int {
	// NOTE(rsc): This function does NOT call the runtime cmpstring function,
	// because we do not want to provide 提供 any performance justification 理由 for
	// using strings.Compare. Basically 基本上 no one should use strings.Compare.
	// As the comment above says, it is here only for symmetry 对称 with package bytes.
	// If performance is important, the compiler should be changed to recognize 认出
	// the pattern 模式 so that all code doing three-way comparisons, not just code
	// using strings.Compare, can benefit 益处.
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
