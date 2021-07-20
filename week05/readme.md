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
```

- [货仓选址](https://www.acwing.com/problem/content/description/106/)

```go
```

- [翻转对（Hard）](https://leetcode-cn.com/problems/reverse-pairs/)

```go
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
```

- [买卖股票的最佳时机 II （Easy）](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

```go
```

- [跳跃游戏 II （Medium）](https://leetcode-cn.com/problems/jump-game-ii/)

```go
```

- [完成所有任务的最少初始能量（Hard）](https://leetcode-cn.com/problems/minimum-initial-energy-to-finish-tasks/)

```go
```