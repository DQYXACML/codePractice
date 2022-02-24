package main

type unionSet struct {
	set  []int // 数据集合
	rank []int // 节点的树高，合并时少的节点挂在高的节点下面
}

func NewUnionSet(size int) *unionSet {
	buf := make([]int, size)
	buf2 := make([]int, size)
	for i := 0; i < size; i++ { // 节点初始化指向自己
		buf[i] = i
	}
	// 初始化节点高度为 1
	for i := 0; i < size; i++ {
		buf2[i] = 1
	}
	return &unionSet{set: buf}
}

func (u *unionSet) GetSize() int {
	return len(u.set)
}

// getRoot: 找出p的根节点,时间复杂度为O(h),h为树的高度
func (u *unionSet) getRoot(p int) int {
	if u.set[p] == p {
		return p
	}
	return u.getRoot(u.set[p])
}

func (u *unionSet) IsConnected(p, q int) bool {
	if p < 0 || q < 0 || p > len(u.set) || q > len(u.set) {
		return false
	}
	return u.getRoot(u.set[p]) == u.getRoot(u.set[q])
}

// Union 找到p的值，将其改为q
func (u *unionSet) Union(p, q int) {
	if p < 0 || q < 0 || p > len(u.set) || q > len(u.set) {
		return
	}
	pRoot := u.getRoot(p)
	qRoot := u.getRoot(q)

	if pRoot != qRoot {
		if u.rank[pRoot] < u.rank[qRoot] {
			u.set[pRoot] = qRoot // 矮树挂在高数下
		} else if u.rank[pRoot] > u.rank[qRoot] {
			u.set[qRoot] = pRoot
		} else {
			u.set[qRoot] = pRoot
			u.rank[qRoot]++
		}
	}
}
