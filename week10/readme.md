# week10

## homework

### [掉落的方块（Hard）](https://leetcode-cn.com/problems/falling-squares/) 

```go
```

## 树状数组

**定义**

```go
树状数组是一种维护数组前缀和、区间和的数据结构
思想和跳表有点类似
跳表：添加索引，高效维护链表
树状数组：添加索引，高效维护数组
```

**如何添加索引**

```go
树状数组的一个结点索引的原始数据数量，与该结点编号在二进制最低位的1有关
1、3、5、7...二进制下以1结尾，仅索引1个数据（自身）
2、6、10、14...二进制下以10为结尾，索引2个数据（自身、它前面那个）
4、12...二进制下以100结尾，索引4个数据（自身、前面3个）
```

**二进制分解与lowbit**

```go
// 任意正整数可以唯一分解为若干个不重复的2的次幂之和
7 =  2^2 + 2^1 + 2^0
12 = 2^3 + 2^2

// lowbit(x)定义为x二进制下最低位的1和后面的0组成的数值（或者说x二进制分解下的最小次幂）
lowbit(7) = lowbit(111)2 = 1(2) = 2^0 = 1
lowbit(12) = lowbit(1100)2 = 100(2) = 2^2 = 4

// 【推导】
-x = ～x + 1
lowbit(x) = x & (~x + 1) = x & -x

//树状数组c的结点c[x] 存储x前面lowbit(x)个数据（包括x）的和
c[7] = a[7]
c[12] = a[9] + a[10] + a[11] + a[12]
```

**树状数组的性质**

```go
// 每个内部结点c[x] 保存以它为根的子树中的所有叶结点的和
// 除树根外，每个内部结点c[x] 的父亲是 c[x + lowbit[x]]
// 树的深度为 O(logN)
// 如果N不是2的整次幂，那么树状数组就是一个具有同样性质的森林结构

**查询**
// 树状数组支持的第一个基本操作 -- 【查询前缀和】
// 根据整数的二进制分解，可以把任意区间[1,x]拆成O(logN)个小区间
13 = 8 + 4 + 1 对应二进制 1101(2) = 1000(2) + 100(2) + 1(2)
[1, 13] 可以拆成 [1,8], [9,12], [13,13]
对应二进制：[1, 1101]拆成[1, 1000],[1001,1100],[1101,1101]
// 规律：
13 前面的lowbit(13) = 1个数，对应区间[13,13]； 13 - 1 = 12；
12 前面的lowbit(12) = 4个数，对应区间[9,12]； 12 - 4 = 8；
8  前面的lowbit(8) = 8个数，对应区间[1,8]; 8 - 8 = 0；结束。

// 代码：
int query(int x) {
    int ans = 0;
    for (; x > 0; x -= x & -x)  ans += c[x]
    return ans;
}
前缀和知道了，区间和（第l～r个数据的和）可以直接有query(r) - query(l - 1)得到
时间复杂度：O(logN) -- 循环次数不超过二进制位数

**更新**
// 树状数组支持的第二个基本操作是单点增加，即把某个数据x增加一个值delta
// 需要更新的索引就是c[x] 以及它的所有祖先结点，至多O(logN)个
// 利用性质：每个内部结点c[x] 的父亲是c[x + lowbit(x)]

// 代码：
void add(int x, int y) {
    for(; x <= N; x += x & -x) c[x] += y
}
// 如果要修改一个数据，可以先算出差值，再执行add操作。
```

**树状数组的局限性**

```go
// 实现简单、效率高、省空间，但也有很大局限性

// 维护的信息需要满足区间减法性质
不然无法通过前缀和相减得到区间和
例如无法直接拿来维护区间最值

// 不能很好地支持修改操作
单点修改需要先求出差值，转化为增加操作
基本上难以支持区间修改（修改连续的一段数据）
```

### [区域和检索 - 数组可修改（Medium）](https://leetcode-cn.com/problems/range-sum-query-mutable/)

**解法1: 树状数组**

```go
type NumArray struct {
    n int
    a []int // 单点数组
    c []int // 前缀和数组
}


func Constructor(nums []int) NumArray {
    n := len(nums)
    a := make([]int, n + 1) // 下标从1开始
    c := make([]int, n + 1) // 下标从1开始

    res := NumArray{
        n : n,
        a : a,
        c : c,
    }
    //fmt.Println(res)
    for i := 1; i <= n; i++ {
        res.a[i] = nums[i - 1]
        add(i, res.a[i], &res.c, n)
    }
    //fmt.Println(res)
    return res
}

func (this *NumArray) Update(index int, val int)  {
    index++ // 下标从1开始
    delta := val - this.a[index]
    add(index, delta, &this.c, this.n)
    this.a[index] = val
}


func (this *NumArray) SumRange(left int, right int) int {
    left++ // 下标从1开始
    right++ // 下标从1开始
    return query(right, &this.c) - query(left - 1, &this.c)
}

func query(x int, c *[]int) int {
    ans := 0
    for ; x > 0; x -= lowbit(x) {
        ans += (*c)[x]
    }
    return ans
}

// 单点增加，所有祖先结点都应该加上变化
func add(x,delta int, c *[]int, n int) {
    for ; x <= n; x += lowbit(x) {
        (*c)[x] += delta
    }
}

// 求出二进制的最低为1和后面0值组成的数
func lowbit(x int) int {
    return x & -x
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
```

## 线段树

- 引入

```go
**定义**
// 线段树是一种基于分治思想的二叉树结构，用于在区间上进行信息统计
线段树的每个节点都代表一个闭区间
线段树具有唯一的根节点，代表的区间是整个统计范围，如[1,N]
线段树的每个叶结点都代表一个长度为1的元区间[x, x]
对于每个内部结点[l, r]，它的左子节点是[l, mid]，右子节点是[mid+1, r]，其中 mid = (l + r) / 2（向下取整）

// 除去树最后一层，整棵线段树一定是一棵完全二叉树
// 树的深度为O(logN)
// 可以按照与二叉堆类似的“父子2倍”节点编号方法
根节点编号为1
编号为p的节点左子节点编号为 p*2，右子节点编号为 p*2 + 1
这样一来，就能简单地使用数组来保存线段树

由于最后一不一定是连续的，保存线段树的数组长度不要小于4N【？】

**区间最值问题**
// 维护一个序列，支持：
查询区间最值（第l个数到第r个数的最大值）
单点修改（更新第x个数据）
*（选做）区间统一修改（把第l个数到第r个数都置为val）

**建树**
Build(1, 1, n)
// 时间复杂度 O(n), 不超过结点数4n

**单点修改**
Change(1, x, v)
// 从根（1号）出发，递归找到代表区间[x, x]的叶子结点
// 自底向上更新保存的信息
// 时间复杂度 O(log(n)), 每层一个结点，更新次数不超过树高

**区间查询**
Query(1, l, r)，从根结点开始递归查询
// 若[l, r]完全覆盖了当前结点代表的区间，则立即回溯，并且该结点的dat值为候选答案
// 若左（右）子结点与[l, r]有重叠部分，则递归访问左（右）子节点
// 时间复杂度 O(log(n)), l, r各在树上划分出一条边界，最多形成2logn个候选区间

**区间修改（选修）**
```

### [区域和检索 - 数组可修改（Medium）](https://leetcode-cn.com/problems/range-sum-query-mutable/)

**解法2: 线段树**

```go
type NumArray struct {
    tree SegmentTree
}


func Constructor(nums []int) NumArray {
    // 构建线段树
    ST := ConstructorST(nums)
    return NumArray{
        tree: ST,
    }
}


func (this *NumArray) Update(index int, val int)  {
    this.tree.change(1, index, val)
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.tree.query(1, left, right)
}

// ###简易版线段树###
// 线段树结点，维护区间、区间和
type Node struct {
    l, r int
    sum int
}
// 线段树
type SegmentTree struct {
    a []Node
}

// 构建线段树
func ConstructorST(nums []int) SegmentTree {
    n := len(nums)
    a := make([]Node, 4*n)
    ST := SegmentTree{
        a : a,
    }
    ST.build(1, 0, n - 1, nums)
    return ST
}

// 递归建树
func (ST *SegmentTree) build(curr int, l,r int, nums []int) {
    ST.a[curr] = Node{
        l: l,
        r: r,
    }
    // 递归边界
    if l == r {
        ST.a[curr].sum = nums[l]
        return
    }
    mid := (l + r) >> 1
    // 分两半，递归
    ST.build(curr * 2, l, mid, nums)
    ST.build(curr * 2 + 1, mid + 1, r, nums)
    // 回溯时，自底向上统计信息
    ST.a[curr].sum = ST.a[curr * 2].sum + ST.a[curr * 2 + 1].sum
}

// 单点修改：先递归找到，然后自底向上统计信息
func (ST *SegmentTree) change(curr int, index int, val int) {
    // 递归边界：叶子结点[index,index]
    if ST.a[curr].l == ST.a[curr].r {
        ST.a[curr].sum = val
        return
    }
    mid := (ST.a[curr].l + ST.a[curr].r) >> 1
    if index <= mid {
        ST.change(curr * 2, index, val)
    } else {
        ST.change(curr * 2 + 1, index, val)
    }
    ST.a[curr].sum = ST.a[curr * 2].sum + ST.a[curr * 2 + 1].sum
}
// 查询[l,r]区间和
// 完全包含：直接返回
// 否则：左右划分
func (ST *SegmentTree) query(curr int, l,r int) int {
    if l <= ST.a[curr].l && r >= ST.a[curr].r {
        return ST.a[curr].sum // 当前结点的区间在[l,r]里面，直接返回区间和到上层
    }
    mid := (ST.a[curr].l + ST.a[curr].r) >> 1
    ans := 0
    if l <= mid { // 累加左子树的区间和
        ans += ST.query(curr * 2, l, r)
    }
    if r > mid { // 累加右子树的区间和
        ans += ST.query(curr * 2 + 1, l, r)
    }
    return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
```

**解法3: 线段树(懒惰标记/延迟标记)**

```go
// TODO
```

### [一个简单的整数问题 2 （Hard）（AcWing）](https://www.acwing.com/problem/content/description/244/)

```go
```

## 离散化

### [区间和的个数（Hard）](https://leetcode-cn.com/problems/count-of-range-sum/)

```go
```

## 位运算

### [位 1 的个数（Easy）](https://leetcode-cn.com/problems/number-of-1-bits/)

```go
// 解法1：获取x在二进制下的第n位 (x >> n) & 1
func hammingWeight(num uint32) int {
    count := 0
    for i := 0; i < 32; i++ {
        if ((num >> i) & 1) == 1 { 
            // 假设num = 0010，右移1位得到0001，(0001&0001) == 1，证明第1位是1
            count++
        }
    }
    return count
}

// 解法2：
func hammingWeight(num uint32) int {
    count := 0
    for num > 0 { // 二进制还有1
        count++
        num -= num & (-num) // lowbit(x)，取出最低位1和后面0代表的数据，每次减掉这个1
    }
    return count
}
```

### [2 的幂（Easy）](https://leetcode-cn.com/problems/power-of-two/)

```go
// lowbit
func isPowerOfTwo(n int) bool {
    // 00001000
    //     1000
    return n > 0 && (n == (n & -n))
}
```

### [颠倒二进制位（Easy）](https://leetcode-cn.com/problems/reverse-bits/)

```go
func reverseBits(num uint32) uint32 {
    var ans uint32
    var one uint32
    one = 1
    for i := 0; i < 32; i++ {
        // 答案左移，然后加上低位的值（0或1）
        ans = (ans << one) | ((num >> i) & one)
    }
    return ans
}

// 优化版本 TODO
```

### [比特位计数（Easy）](https://leetcode-cn.com/problems/counting-bits/)

```go
//解法1: 暴力解法
func countBits(n int) []int {
    var ans []int
    for i := 0; i <= n; i++ {
        ans = append(ans, hammingWeight(uint32(i)))
    }
    return ans
}
// 位 1 的个数（Easy）
func hammingWeight(num uint32) int {
    count := 0
    for num > 0 { // 二进制还有1
        count++
        num -= num & (-num) // lowbit(x)，取出最低位1和后面0代表的数据，每次减掉这个1
    }
    return count
}

//解法2: dp解法
func countBits(n int) []int {
    ans := make([]int, n+1)
    ans[0] = 0
    for i := 1; i <= n; i++ {
        //ans[i] = ans[i - (i & -i)] + 1 // 减lowbit
        ans[i] = ans[i & (i - 1)] + 1 // 最低位清0
    }
    return ans
}
```

### [Pow(x, n) （Medium）](https://leetcode-cn.com/problems/powx-n/)

```go
// 快速幂解法
func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1 / myPow(x, -n)
    }
    if n == 0 {
        return 1
    }
    var ans float64
    ans = 1
    for n > 0 {
        // x的21次幂 = x的（16+4+1）次幂
        // 10101
        if (n & 1) == 1 { ans = ans * x }
        n = n >> 1 // 用过的位置右移去掉
        x = x * x // 【关键】
    }
    return ans
}

// 分治解法
func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n) // 负负为正
}

func quickMul(x float64, n int) float64 {
	if n == 0 { return 1 }

	y := quickMul(x, n/2)
	if n%2 == 0 { return y * y }
	return y * y * x // n为奇数，n/2 向下取整，所以这里还得乘回1个x
}
```

### [N 皇后（Hard）](https://leetcode-cn.com/problems/n-queens/)

```go

```

### [解数独（Hard）](https://leetcode-cn.com/problems/sudoku-solver/)

**[回溯+剪枝（bool数组版本）](https://github.com/chenlixin93/leetcode-go/tree/main/week09#%E8%A7%A3%E6%95%B0%E7%8B%AChard)**

**回溯+剪枝（位运算版本）**

```go
func solveSudoku(board [][]byte)  {    
    // bool数组改为二进制版本
    row := make([]int, 9)
    col := make([]int, 9)
    box := make([]int, 9)
    initial := int(1022) // 二进制 1111111110（0位浪费）
    for i := 0; i < 9; i++ {
        row[i] = initial
        col[i] = initial
        box[i] = initial
    }

	// 预处理数字的可用情况
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' { // 初始化时false
				digit := board[i][j] - byte('0')
				row[i] ^= (1 << digit) // 异或，相同为0
				col[j] ^= (1 << digit)
				box[i/3*3+j/3] ^= (1 << digit)
			}
		}
	}
	dfs(board, row, col, box)
    return
}

func dfs(board [][]byte, row,col,box []int) bool {
	// 找到第一个可填的数
	location := getLeastPossibleLocation(board, row, col, box)
	x := location[0]
	y := location[1]
	if x == -1 { return true } // 填满了，有解
	// 尝试填入1-9
    boxId := x/3*3+y/3
    availability := row[x] & col[y] & box[boxId]
	for digit := 1; digit <= 9; digit++ {
        if ((availability >> digit) & 1) == 1 {
            // 初始都是true
            row[x] ^= (1 << digit) // 取反
            col[y] ^= (1 << digit) 
            box[boxId] ^= (1 << digit)
            board[x][y] = byte('0') + byte(digit)
            if dfs(board, row, col, box) { 
                return true 
            }
            board[x][y] = '.'
            row[x] ^= (1 << digit) // 取反
            col[y] ^= (1 << digit) 
            box[boxId] ^= (1 << digit)
        }
	}
	return false
}

// 寻找可填的位置（决策最少的一个）
func getLeastPossibleLocation(board [][]byte, row,col,box []int) [2]int {
	ansCnt := 10
	ans := [2]int{-1, -1}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 如果是小数点
			if board[i][j] == '.' {
				// 检查能填的可能性
                // 要能填入，一定是行、列、盒子都还没填过的，直接统计1的个数
                boxId := i/3*3+j/3
                cnt := hammingWeight(row[i] & col[j] & box[boxId])
				if cnt < ansCnt {
					ansCnt = cnt
					ans = [2]int{i, j}
				}
			}
		}
	}
	return ans
}

func hammingWeight(num int) int {
    count := 0
    for num > 0 { // 二进制还有1
        count++
        num -= num & (-num) // lowbit(x)，取出最低位1和后面0代表的数据，每次减掉这个1
    }
    return count
}
```

## 随机题目

### [天际线问题（Hard）](https://leetcode-cn.com/problems/the-skyline-problem/)

```go

```

### [包含每个查询的最小区间](https://leetcode-cn.com/problems/minimum-interval-to-include-each-query/)

```go

```