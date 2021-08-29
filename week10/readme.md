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

**解法2: 线段树**

```go

```

## 线段树

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

```

### [2 的幂（Easy）](https://leetcode-cn.com/problems/power-of-two/)

```go

```

### [颠倒二进制位（Easy）](https://leetcode-cn.com/problems/reverse-bits/)

```go

```

### [比特位计数（Easy）](https://leetcode-cn.com/problems/counting-bits/)

```go

```

### [Pow(x, n) （Medium）](https://leetcode-cn.com/problems/powx-n/)

```go

```

### [N 皇后（Hard）](https://leetcode-cn.com/problems/n-queens/)

```go

```

### [解数独（Hard）](https://leetcode-cn.com/problems/sudoku-solver/)

```go

```

## 随机题目

### [天际线问题（Hard）](https://leetcode-cn.com/problems/the-skyline-problem/)

```go

```

### [包含每个查询的最小区间](https://leetcode-cn.com/problems/minimum-interval-to-include-each-query/)

```go

```