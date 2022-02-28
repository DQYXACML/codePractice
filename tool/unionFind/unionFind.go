package unionFind

type unionSet struct {
	parent []int // 数据集合
	rank   []int // 节点的树高，合并时少的节点挂在高的节点下面；如果不是根节点，其rank相当于以它作为根节点的子树的深度
	count  int
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
	return &unionSet{parent: buf, rank: buf2, count: size}
}

func NewUnionSetByDoubleList(arr [][]byte) *unionSet {
	row := len(arr)
	col := len(arr[0])
	N := row * col
	buf := make([]int, N)
	buf2 := make([]int, N)
	var size int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if arr[i][j] == '1' {
				index := double2Single(i, j)
				buf[index] = index
				buf2[index] = 1
				size++
			}
		}
	}
	return &unionSet{
		parent: buf,
		rank:   buf2,
		count:  size,
	}
}

// getRoot: 找出p的根节点,时间复杂度为O(h),h为树的高度
// 路径压缩，找根的过程中将节点指向根节点，优化下次查找
func (u *unionSet) getRoot(p int) int {
	if u.parent[p] == p {
		return p
	}
	u.parent[p] = u.getRoot(u.parent[p])
	return u.parent[p]
}

func (u *unionSet) IsConnected(p, q int) bool {
	if p < 0 || q < 0 || p > len(u.parent) || q > len(u.parent) {
		return false
	}
	return u.getRoot(u.parent[p]) == u.getRoot(u.parent[q])
}

// Union 找到p的值，将其改为q
func (u *unionSet) Union(p, q int) {
	if p < 0 || q < 0 || p > len(u.parent) || q > len(u.parent) {
		return
	}
	pRoot := u.getRoot(p)
	qRoot := u.getRoot(q)

	if pRoot != qRoot {
		// 高度低的挂在高度高的下面，高度高的高度不需要调整
		if u.rank[pRoot] < u.rank[qRoot] {
			u.parent[pRoot] = qRoot // 矮树挂在高树下
		} else if u.rank[pRoot] > u.rank[qRoot] {
			u.parent[qRoot] = pRoot
		} else {
			// 高度相等时，随便挂一个，需要调整新根的高度，即+1
			u.parent[qRoot] = pRoot
			u.rank[pRoot]++
		}
	}
	u.count-- // 一共分了多少组
}

// 二维数组合并
func (u *unionSet) UnionByDouble(r1, c1, r2, c2 int) {
	if c1 < 0 || r1 < 0 || c1 > len(u.parent) || r1 > len(u.parent) ||
		c2 < 0 || r2 < 0 || c2 > len(u.parent) || r2 > len(u.parent) {
		return
	}
	index1 := r1*c1 + c1
	index2 := r2*c2 + c2
	qRoot := u.getRoot(index1)
	pRoot := u.getRoot(index2)

	if qRoot != pRoot {
		if u.rank[qRoot] < u.rank[pRoot] {
			u.rank[qRoot] = pRoot
		} else if u.rank[qRoot] > u.rank[pRoot] {
			u.rank[pRoot] = qRoot
		} else {
			u.rank[pRoot] = qRoot
			u.rank[pRoot]++
		}
		u.count--
	}
}

func (u *unionSet) TotalCount() int {
	return u.count
}

func double2Single(row, col int) int {
	return row*col + col
}
