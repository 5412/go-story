package main

import (
	"io"
	"strings"
)

// 对外的接口
type Replacer struct {
	r replacer
}

type replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}

// 实现接口 replacer
func (r *Replacer) Replace(s string) string {
	return r.r.Replace(s)
}

func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) {
	return r.r.WriteString(w, s)
}

// 字典树的节点类型
// 比如：键 "ax", "ay", "bcbc", "x" 和 "xy" 形成有八个节点的树
// n0 -
// n1 a-
// n2 .x+
// n3 .y+
// n4 b-
// n5 .cbc+
// n6 x+
// n7 .y+
type trieNode struct {

	// value 是字典数节点的 key/value 的键值对的值
	value string

	priority int

	prefix string //  此节点与next节点不同的地方

	next *trieNode // 单一子节点，

	table []*trieNode // 所有子节点
}
func (t *trieNode) add(key, val string, priority int, r *genericReplacer) {
	if key == "" {
		if t.priority == 0 {
			t.value = val
			t.priority = priority
		}
		return
	}

	if t.prefix != "" {
		// 需要分裂多个节点间的前缀
		var n int
		for ; n < len(t.prefix) && n < len(key); n++ {
			if t.prefix[n] != key[n] {
				break // 找到前缀与key第一个不相等的byte
			}
		}

		if n == len(t.prefix) {
			// key 前前N个字符 与前缀相等，
			// 则向自己的子节点添加节点
			t.next.add(key[n:], val, priority, r)
		} else if n == 0 {
			// 第一个字节就不相等的话，在这个节点开始一个新的检索表，
			// 查找t.prefix[0]将会通向 prefixNode,
			// 查找key[0]将会通向 keyNode
			var prefixNode *trieNode
			if len(t.prefix) == 1 {
				// 如果当前节点的前缀就一个字节，
				// 那么前缀节点就是当前节点的next节点
				prefixNode = t.next
			} else {
				// 如果当前节点的前缀大于一个直接，
				// 那么前缀节点是一个新的节点，
				// 其前缀是当前节点的前缀的[1:],
				// 其next节点是当前节点的next节点
				prefixNode = &trieNode{ // 指针
					prefix: t.prefix[1:],
					next:   t.next,
				}
			}

			// 新建键节点，初始化当前节点的搜索table，
			// 向table中写入前缀节点，
			// 写入keyNode, 重制当前节点的前缀为空，next节点为空
			// keyNode节点继续添加新节点
			keyNode := new(trieNode)
			t.table = make([]*trieNode, r.tableSize)
			t.table[r.mapping[t.prefix[0]]] = prefixNode
			t.table[r.mapping[key[0]]] = keyNode
			t.prefix = ""
			t.next = nil
			keyNode.add(key[1:], val, priority, r)
		} else {
			// 插入新的节点在与prefix公共区域之后
			// n比prefix的长度短
			// 新建next节点，其prefix为当前节点prefix[n:]
			// 其next节点为当前节点的next节点
			// 修改当前节点的prefix为其前n个字节
			// 当前节点的next节点为新建的next节点
			// next节点继续向下添加节电
			next := &trieNode{
				prefix:	t.prefix[n:],
				next:	t.next,
			}

			t.prefix = t.prefix[:n]
			t.next = next
			next.add(key[n:], val, priority, r)
		}
	} else if t.table != nil {
		// 如果当前节点的前缀为空，且有搜索表，那么直接插入此搜索表
		m := r.mapping[key[0]]
		if t.table[m] == nil {
			t.table[m] = new(trieNode)
		}
		t.table[m].add(key[1:], val, priority, r)
	} else {
		t.prefix = key
		t.next = new(trieNode)
		t.next.add("", val, priority, r)
	}
}

type genericReplacer struct {
	root trieNode

	tableSize int //  字典树节点的搜索表的大小，唯一字节键的个数

	mapping [256]byte // key bytes 的一个紧凑的映射表
}

func (r *genericReplacer) lookup(s string, ignoreRoot bool) (val string, keylen int, found bool){
	// 沿着字典树向下迭代，根据最高优先级获取value 和 keylen
	bestPriority := 0
	node := &r.root
	n := 0

	for node != nil {
		if node.priority > bestPriority && !(ignoreRoot && node == &r.root) {
			bestPriority = node.priority
			val = node.value
			keylen = n
			found = true
		}

		if s == "" {
			break
		}

		if node.table != nil {
			index := r.mapping[s[0]]
			if int(index) == r.tableSize {
				break
			}
			node = node.table[index]
			s = s[1:]
			n++
		} else if node.prefix != "" && strings.HasPrefix(s, node.prefix) {
			n += len(node.prefix)
			s = s[len(node.prefix):]
			node = node.next
		} else {
			break
		}
	}
	return
}

func makeGenericReplacer(oldnew []string) *genericReplacer {
	r := new(genericReplacer)
	// 查找每个用到的byte，并给它们分配一个index
	// 先将用到的byte作为mapping的key，value设为1,
	// 因mapping是一个256长度的数组，数组默认值为0
	// 循环数组将所有值累加，则为用到的byte个数
	for i := 0; i < len(oldnew) ; i += 2 {
		key := oldnew[i]
		for j := 0 ; j < len(key); j++ {
			r.mapping[key[j]] = 1
		}
	}

	for _, b := range r.mapping {
		r.tableSize += int(b)
	}

	var index byte
	// 重新给mapping的每个元素赋值，
	// 如果原值为0那么值为tableSize，
	// 如果原值为1其值为新的自增index e.g. [0,1,2,3,4,.....]
	for i, b := range r.mapping {
		if b == 0 {
			r.mapping[i] = byte(r.tableSize)
		} else {
			r.mapping[i] = index
			index++
		}
	}

	// 确保更节点用了一个搜索表（处于性能的原因）
	r.root.table = make([]*trieNode, r.tableSize)

	// 构造字典树，并设置优先级，oldnew中越早出现的键值对，优先级越高
	for i := 0; i < len(oldnew); i += 2 {
		r.root.add(oldnew[i], oldnew[i+1], len(oldnew) - i, r)
	}

	return r
}