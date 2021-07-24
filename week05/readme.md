# week05

## homework

- [在 D 天内送达包裹的能力（Medium）](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/)

```go
```

- [在线选举（Medium）](https://leetcode-cn.com/problems/online-election/)

```go
```

- [爱吃香蕉的珂珂（Medium）](https://leetcode-cn.com/problems/koko-eating-bananas/)

```go
```

- [区间和的个数（选做）（Hard）](https://leetcode-cn.com/problems/count-of-range-sum/)

```go
```

## 排序

### 初级排序算法

- 选择排序（该放哪个数了）

```go
// 每次从未排序数据中找最小值，放到已排序序列的末尾
func selectSort(nums []int) []int {
    length := len(nums)
    if length <= 1 {
        return nums
    }

    for i :=0; i < length; i++ {
        min := i // 从0开始
        for j := length - 1; j > i; j-- { // 从右侧开始检查最小值
            if nums[j] < nums[min] {
                min = j
            }
        }
        nums[i], nums[min] = nums[min], nums[i]
    }

    return nums
}
```

- 插入排序（这个数该放哪）

```go
// 从前到后依次考虑每个未排序数据，在已排序的序列总找到合适位置插入
// 算法设计的思路是，将数组划分成两部分，第一部分是有序的，第二部分是无序的
// 每次从无序部分取一个元素，将这个元素插入到有序部分，保持有序部分的有序性质
// 直到无序部分为空
func insertSort(nums []int) []int {
    var j int
    for i := 1; i < len(nums); i++ {
        tmp := nums[i]
        for  j = i; j > 0 && tmp < nums[j - 1]; j-- {
            nums[j] = nums[j - 1]
        }
        nums[j] = tmp
    }
    return nums
}
```

- 冒泡排序

```go
// 不断循环扫描，每次查看相邻元素，如果逆序则交换
func bubbleSort(nums []int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ { // 做完第一轮就把最小数放到了首位
            if nums[i] > nums[j] {
                nums[i],nums[j] = nums[j],nums[i]
            }
        }
    }
    return nums
}
```

### 重要排序算法

- 堆排序（Heap Sort是对选择排序的优化，利用二叉堆高效选出最小值）

```go
```

- 希尔排序（Shell Sort是对插入排序的优化）（选做）

- 归并排序（基于分治，先排序左右子数组，然后合并两个有序数组）

```go
func sortArray(nums []int) []int {
    mergeSort(nums, 0, len(nums) - 1)
    return nums
}

// 归并排序
func mergeSort(nums []int, left,right int) {
    if right <= left { return }
    mid := (left + right) >> 1
    mergeSort(nums, left, mid)
    mergeSort(nums, mid + 1, right)
    merge(nums, left, mid, right)
    return
}

// 对left...right之间的数进行排序
func merge(nums []int, left,mid,right int) {
    tmp := make([]int, right - left + 1) // 新开的临时数组
    i := left
    j := mid + 1
    k := 0

    for i <= mid && j <= right {
        if nums[i] <= nums[j] { // 哪个小取哪个放入临时数组
            tmp[k] = nums[i]
            i++
        } else {
            tmp[k] = nums[j]
            j++
        }
        k++
    }

    for i <= mid { // 把剩余的数取尽
        tmp[k] = nums[i]
        k++
        i++
    }
    for j <= right {
        tmp[k] = nums[j]
        k++
        j++
    }

    for p := 0; p < len(tmp); p++ {
        nums[left + p] = tmp[p]
    }
}
```

- 快速排序（基于分治，先调配出左右子数组，然后对左右子数组分别进行排序）

```go
func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums) - 1)
    return nums
}

// 解法1：不带pivot 944ms
func quickSort(nums []int, left, right int) {
    if left > right {
        return
    }
    i, j, base := left, right, nums[left]
    for i < j {
        for nums[j] >= base && i < j {
            j--
        }
        for nums[i] <= base && i < j {
            i++
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    nums[i], nums[left] = nums[left], nums[i]
    quickSort(nums, left, i - 1)
    quickSort(nums, i + 1, right)
}

// 解法2：带pivot
func quickSort(nums []int, l,r int) {
    if l >= r { return }
    pivot := partition(nums, l, r)
    quickSort(nums, l, pivot)
    quickSort(nums, pivot + 1, r)
    return
}

func partition(nums []int, l,r int) int {
    // pivot := l+1 // 2-1不随机 480ms
    pivot := l + rand.Intn(r - l + 1) // 2-2随机 48ms
    pivotVal := nums[pivot] 
    for l <= r {
        for nums[l] < pivotVal { l++ }
        for nums[r] > pivotVal { r-- }
        if l <= r {
            nums[l],nums[r] = nums[r],nums[l]
            l++
            r--
        }
    }
    return r
}
```

### 非比较类排序

- 计数排序

- 桶排序

- 基数排序

### 习题

- [排序数组（Medium）](https://leetcode-cn.com/problems/sort-an-array/)

```go
// 使用快排即可
```

- [数组的相对排序](https://leetcode-cn.com/problems/relative-sort-array/)

```go
func relativeSortArray(arr1 []int, arr2 []int) []int {
    // 哈希 + 计数排序
    ans := make([]int, len(arr1))
    count := map[int]int{}

    for _, v := range arr1 { // 统计arr1中数字出现的次数
        count[v]++
    }

    n := 0
    for k,_ := range arr2 { // 按照arr2的顺序去取出
        for count[arr2[k]] > 0 {
            count[arr2[k]]--
            ans[n] = arr2[k]
            n++
        }
    }

    for val := 0; val <= 1000; val++ { // 剩余没有在arr2出现的，按顺序取出
        for count[val] > 0 {
            count[val]--
            ans[n] = val
            n++
        }
    }
    return ans
}
```

- [合并区间（Medium）](https://leetcode-cn.com/problems/merge-intervals/)

```go
// 解法1: 双关键字快排
func merge(intervals [][]int) (ans [][]int) {
    length := len(intervals)
    // 保存数组下标
    var idx []int
    for i := 0; i < length; i++ {
        idx = append(idx, i)
    }
    // 使用快排进行双关键字排序
    quickSort(idx, 0, length - 1, intervals)
    
    left := -1
    far := -1
    for _,newIdx := range idx {
        interval := intervals[newIdx] // 按最新下标顺序取出intervals
        start := interval[0]
        end := interval[1]
        if start <= far {
            far = max(far, end)
        } else {
            if far >= 0 {
                ans = append(ans, []int{left, far})
            }
            left = start
            far = end
        }
    }
    if far >= 0 {
        ans = append(ans, []int{left, far})
    }
    return
}

// 解法2: 差分思想

// 快排，带pivot
func quickSort(nums []int, l,r int, intervals [][]int) {
    if l >= r { return }
    pivot := partition(nums, l, r, intervals)
    quickSort(nums, l, pivot, intervals)
    quickSort(nums, pivot + 1, r, intervals)
    return
}
func partition(nums []int, l,r int, intervals [][]int) int {
    pivot := l + rand.Intn(r - l + 1) // 随机 48ms
    pivotVal := nums[pivot] 
    for l <= r {
        for isALessB(intervals[nums[l]], intervals[pivotVal]) { l++ }
        for isAGreatB(intervals[nums[r]], intervals[pivotVal]) { r-- }
        if l <= r {
            nums[l],nums[r] = nums[r],nums[l]
            l++
            r--
        }
    }
    return r
}

// 定义数组a是否小于数组b
func isALessB (a []int, b []int) bool {
    return a[0] < b[0] || (a[0]==b[0] && a[1] < b[1])
}

// 定义数组a是否大于数组b
func isAGreatB (a []int, b []int) bool {
    return a[0] > b[0] || (a[0]==b[0] && a[1] > b[1])
}

func max(a,b int) int {
    if a > b { return a }
    return b
}
```

- [数组中的第 K 个最大元素（Medium）](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

```go
// 解法1：大根堆
import "container/heap"

func findKthLargest(nums []int, k int) int {
	maxHeap := &MaxIntHeap{}
	heap.Init(maxHeap)
	
	length := len(nums)
	for length > 0 { // 构建大根堆
		heap.Push(maxHeap, nums[length - 1])
		length--
	}
	for k - 1 > 0 { // 取k-1次堆顶
		heap.Pop(maxHeap)
		k--
	}
	return heap.Pop(maxHeap).(int)
}

// 实现大顶堆
type MaxIntHeap []int

func (h MaxIntHeap) Len() int {
    return len(h)
}

func (h MaxIntHeap) Less(i,j int) bool {
    return h[i] > h[j]
}

func (h MaxIntHeap) Swap(i,j int) {
    h[i],h[j] = h[j],h[i]
}

func (h *MaxIntHeap) Push(x interface{}) {
    *h = append(*h, x.(int)) // .(type)
}

func (h *MaxIntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0:n-1] // 前闭后开
    return x
}
```

- [货仓选址](https://www.acwing.com/problem/content/description/106/)

```go
```

- [翻转对（Hard）](https://leetcode-cn.com/problems/reverse-pairs/)

```go
// 解题思路
// 把数组一分为二，左边下标天然小于右边
// 当左右都排好序后，左边、右边分别统计完，只需要加上左边下标和右边下标比较的情况 
// 子问题的最细粒度的情况就是左右各一个数

// 解法：归并排序 + 计算
var ans int // 全局变量
func reversePairs(nums []int) int {
	ans = 0 // 注意清0，不然会存着上一次的结果
    mergeSort(nums, 0, len(nums) - 1)
	return ans
}

// 归并排序part1
func mergeSort(nums []int, left,right int) {
    if right <= left { return }
    mid := (left + right) >> 1
    mergeSort(nums, left, mid)
    mergeSort(nums, mid + 1, right)
    calculate(nums, left, mid, right)
    merge(nums, left, mid, right)
}

// 计算
func calculate(nums []int, left,mid,right int) {
    i := left
    j := mid
    for ; i <= mid; i++ {
        for j < right && nums[i] > 2 * nums[j + 1] {
            j++
        }
        ans += j - mid // 右边符合条件的次数
    }
}

// 归并排序part2:对left...right之间的数进行排序
func merge(nums []int, left,mid,right int) {
    tmp := make([]int, right - left + 1) // 新开的临时数组
    i := left
    j := mid + 1
    k := 0

    for i <= mid && j <= right {
        if nums[i] <= nums[j] { // 哪个小取哪个放入临时数组
            tmp[k] = nums[i]
            i++
        } else {
            tmp[k] = nums[j]
            j++
        }
        k++
    }

    for i <= mid { // 把剩余的数取尽
        tmp[k] = nums[i]
        k++
        i++
    }
    for j <= right {
        tmp[k] = nums[j]
        k++
        j++
    }

    for p := 0; p < len(tmp); p++ {
        nums[left + p] = tmp[p]
    }
}
```

## 贪心

- [零钱兑换（Medium）](https://leetcode-cn.com/problems/coin-change/)

```go
```

- [柠檬水找零（Easy）](https://leetcode-cn.com/problems/lemonade-change/description/)

```go
```

- [分发饼干（Easy）](https://leetcode-cn.com/problems/assign-cookies/description/)

```go
// 效率 32ms；
// 使用快排代替内置排序算法：28ms
func findContentChildren(g []int, s []int) int {
    // 大饼干给大孩子，小饼干给小孩子
    sort.Ints(g) // 胃口
    sort.Ints(s) // 饼干
    j := 0
    ans := 0
    for i := 0; i < len(g); i++ {
        for j < len(s) && s[j] < g[i] { j++ } // 不满足就跳过
        if j < len(s) { 
            j++
            ans++
        }
    }
    return ans
}
```

- [买卖股票的最佳时机 II （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

```go
func maxProfit(prices []int) int {
    // 持有股票，卖不卖？ 往后看一天，如果明天跌，今天肯定卖
    // 没有股票，买不买？ 往后看一天，如果明天涨，那肯定买
    // 最优买入不会损失收益，那么进一步推导出：
    // 获得所有 price[i] - price[i - 1] > 0 的收益
    ans := 0
    for i := 1; i < len(prices); i++ {
        if (prices[i] - prices[i - 1]) > 0 {
            ans += prices[i] - prices[i - 1]
        }
    }
    return ans
}
```

- [跳跃游戏 II （Medium）](https://leetcode-cn.com/problems/jump-game-ii/)

```go
func jump(nums []int) int {
    // 决策包容性：同样是跳1步，从 a 跳到 “能跳得更远位置c”的 b，
    // 未来的可达集合包含了跳到其他b的可达集合，所以这个局部最优决策是正确的
    now := 0
    ans := 0
    right := 0
    //next := 0
    for now < len(nums) - 1 {
        if nums[now] == 0 { return -1 }

        right = now + nums[now]
        if right >= len(nums) - 1 { // 到达最后
            return ans + 1
        }
        // 从now出发可以到[now+1, right]
        next := now + 1
        for i := now + 2; i <= right; i++ { // 搜索[now+1, right]之间能跳到最远c的某个b点
            next_right := i + nums[i]
            if next_right > (next + nums[next]) {
                next = i
            }
        }
        now = next
        ans++
    }
    return ans
}
```

- [完成所有任务的最少初始能量（Hard）](https://leetcode-cn.com/problems/minimum-initial-energy-to-finish-tasks/)

图示

---------|------|-----|----------|

         i   i+1  i+2      n

推导过程：
```
定义相邻两个任务 i 、i+1 完成的顺序是 p->q 或者 q->p
先做 p ，再做 q，那么就如图示的顺序 i -> i+1
定义做 p 之前至少需要的初始能量为 min[p]，做完p任务实际消耗 a[p]
定义做 q 之前至少需要的初始能量为 min[q]，做完q任务实际消耗 a[q]
完成 i+2 ... n 的部分需要的能量设为 S，即达到 i+2 时，至少还有 S 的能量
```

part1
```
那么，先做 p 所需的初始能量 max（min[p], i 到 n 之间的能量条）
如果 min[p] 小于后者，那肯定完不成 p -> q，所以是取较大的能量数;

i 到 n 之间的能量条 ： a[p] + max（min[q], a[q] + S）
第二项同理，如果 min[q] 小于后者，也完不成 q -> i+2；

也就有第一个式子：
 max（min[p], a[p] + max（min[q], a[q] + S））
```

part2
```
先做 q 所需的初始能量，同理可得：
 max（min[q], a[q] + max（min[p], a[p] + S））
```

part3
```
假设先做 p 比较优，则有 part1 < part2

拆括号，
max（min[p], a[p]+min[q], a[p]+a[q]+S） <
max（min[q], a[q]+min[p], a[q]+a[p]+S）

消除第三项后，
max（min[p], a[p]+min[q]）< max（min[q], a[q]+min[p]）
由于前者第二项包含了后者第一项，后者第二项包含了前者第一项，
所以min[p]、min[q]在这个式子里没有决定性，将其消除

那么必定有 a[p]+min[q] < a[q]+min[p]
最终得出 a[p] - min[p] < a[q] - min[q]
```

代码
```go
func minimumEffort(tasks [][]int) int {
    // 贪心策略：按照 actual - min 升序排序，以此顺序完成任务
    // 也可以理解位 消耗小，门槛大的，是先做的条件
    length := len(tasks)
    // 保存数组下标
    var idx []int
    for i := 0; i < length; i++ {
        idx = append(idx, i)
    }
    // 使用快排进行双关键字排序
    quickSort(idx, 0, length - 1, tasks)
    // 正序做任务，但是计算要倒序
    energy := 0
    for i := len(idx) - 1; i >= 0; i-- {
        // 门槛和实际消耗哪个大取哪个
        energy = max(tasks[idx[i]][1], energy + tasks[idx[i]][0])
    }
    return energy
}

// 快排，带pivot
func quickSort(nums []int, l,r int, tasks [][]int) {
    if l >= r { return }
    pivot := partition(nums, l, r, tasks)
    quickSort(nums, l, pivot, tasks)
    quickSort(nums, pivot + 1, r, tasks)
    return
}
func partition(nums []int, l,r int, tasks [][]int) int {
    pivot := l + rand.Intn(r - l + 1) // 随机 48ms
    pivotVal := nums[pivot] 
    for l <= r {
        for isALessB(tasks[nums[l]], tasks[pivotVal]) { l++ }
        for isAGreatB(tasks[nums[r]], tasks[pivotVal]) { r-- }
        if l <= r {
            nums[l],nums[r] = nums[r],nums[l]
            l++
            r--
        }
    }
    return r
}

// 定义数组a是否小于数组b
func isALessB (a []int, b []int) bool {
    return a[0] - a[1] < b[0] - b[1]
}

// 定义数组a是否大于数组b
func isAGreatB (a []int, b []int) bool {
    return a[0] - a[1] > b[0] - b[1]
}

func max(a,b int) int {
    if a > b { return a }
    return b
}
```