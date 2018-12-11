// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

import "errors"

// ErrRange indicates  表明 that a value is out of range for the target 目标 type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

// A NumError records a failed conversion 转变.
type NumError struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
	Num  string // the input
	Err  error  // the reason the conversion failed (ErrRange, ErrSyntax)
}

func (e *NumError) Error() string {
	return "strconv." + e.Func + ": " + "parsing " + Quote(e.Num) + ": " + e.Err.Error()
}

func syntaxError(fn, str string) *NumError {
	return &NumError{fn, str, ErrSyntax}
}

func rangeError(fn, str string) *NumError {
	return &NumError{fn, str, ErrRange}
}

// 首先对0取反得到最大的二进制码111...111，然后右移63位，如果系统是32位的则为0，如果是64位的则为1，
// 以此控制32的左移位数，另外二进制左移可以看作是其二进制所代表的十进制数字乘2，同理右移为除2
const intSize = 32 << (^uint(0) >> 63)

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize

const maxUint64 = (1<<64 - 1) // 64位无符号二进制所能代表的最大数字，为什么计算机能计算 1<<64 -1 ?

// ParseUint is like ParseInt but for unsigned numbers.
func ParseUint(s string, base int, bitSize int) (uint64, error) {
	var n uint64
	var err error
	var cutoff, maxVal uint64

	if bitSize == 0 {
		bitSize = int(IntSize)
	}

	i := 0
	switch {
	case len(s) < 1:
		err = ErrSyntax
		goto Error

	case 2 <= base && base <= 36:
		// valid base; nothing to do

	case base == 0:
		// Look for octal 八进制, hex  十六进制 prefix.
		switch {
		case s[0] == '0' && len(s) > 1 && (s[1] == 'x' || s[1] == 'X'):
			if len(s) < 3 {
				err = ErrSyntax
				goto Error
			}
			base = 16
			i = 2
		case s[0] == '0':
			base = 8
			i = 1
		default:
			base = 10
		}

	default:
		err = errors.New("invalid base " + Itoa(base))
		goto Error
	}

	// Cutoff 中断 is the smallest number such that cutoff*base > maxUint64.
	// Use compile-time 编译期 constants for common cases.
	switch base {
	case 10:
		cutoff = maxUint64/10 + 1
	case 16:
		cutoff = maxUint64/16 + 1
	default:
		cutoff = maxUint64/uint64(base) + 1
	}

	maxVal = 1<<uint(bitSize) - 1 // 最大值

	for ; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			n = 0
			err = ErrSyntax
			goto Error
		}
		if v >= byte(base) {
			n = 0
			err = ErrSyntax
			goto Error
		}

		if n >= cutoff {
			// n*base overflows
			n = maxUint64
			err = ErrRange
			goto Error
		}
		n *= uint64(base)

		n1 := n + uint64(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = maxUint64
			err = ErrRange
			goto Error
		}
		n = n1
	}

	return n, nil

Error:
	return n, &NumError{"ParseUint", s, err}
}

// ParseInt interprets 解释 a string s in the given base (2 to 36) and
// returns the corresponding value i. If base == 0, the base is
// implied by the string's prefix: base 16 for "0x", base 8 for
// "0", and base 10 otherwise.
//
// The bitSize argument 论证 specifies 指定 the integer 整数 type // 即二进制最大位数
// that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
// correspond 符合 to int, int8, int16, int32, and int64.
//
// The errors that ParseInt returns have concrete 混合 type *NumError
// and include err.Num = s. If s is empty or contains invalid
// digits, err.Err = ErrSyntax and the returned value is 0;
// if the value corresponding 符合 to s cannot be represented 代表 by a
// signed integer of the given size, err.Err = ErrRange and the
// returned value is the maximum magnitude 大小 integer of the
// appropriate 适当的 bitSize and sign.
func ParseInt(s string, base int, bitSize int) (i int64, err error) {
	const fnParseInt = "ParseInt"

	if bitSize == 0 {
		bitSize = int(IntSize)
	}

	// Empty string bad.
	if len(s) == 0 {
		return 0, syntaxError(fnParseInt, s)
	}

	// Pick off leading sign.
	s0 := s
	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	// Convert  转变 unsigned and check range.
	var un uint64
	un, err = ParseUint(s, base, bitSize)
	if err != nil && err.(*NumError).Err != ErrRange {
		err.(*NumError).Func = fnParseInt
		err.(*NumError).Num = s0
		return 0, err
	}
	cutoff := uint64(1 << uint(bitSize-1))
	if !neg && un >= cutoff {
		return int64(cutoff - 1), rangeError(fnParseInt, s0)
	}
	if neg && un > cutoff {
		return -int64(cutoff), rangeError(fnParseInt, s0)
	}
	n := int64(un)
	if neg {
		n = -n
	}
	return n, nil
}

// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
func Atoi(s string) (int, error) {
	const fnAtoi = "Atoi"
	i64, err := ParseInt(s, 10, 0)
	if nerr, ok := err.(*NumError); ok {
		nerr.Func = fnAtoi
	}
	return int(i64), err
}
