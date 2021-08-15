# week08

## homework

### 基础问题

- [转换成小写字母（Easy）](https://leetcode-cn.com/problems/to-lower-case/)

```go
```

- [最后一个单词的长度（Easy）](https://leetcode-cn.com/problems/length-of-last-word/)

```go
```

- [宝石与石头（Easy）](https://leetcode-cn.com/problems/jewels-and-stones/)

```go
```

- [字符串中的第一个唯一字符（Easy）](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)

```go
```

- [最长公共前缀（Easy）](https://leetcode-cn.com/problems/longest-common-prefix/description/)

```go
```

### 字符串操作

- [反转字符串（Easy）](https://leetcode-cn.com/problems/reverse-string/)

```go
```

- [翻转字符串里的单词（Medium）](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

```go
```

- [仅仅反转字母（Easy）](https://leetcode-cn.com/problems/reverse-only-letters/)

```go
```

### 同构 / 异位词系列问题

- [同构字符串（Easy）](https://leetcode-cn.com/problems/isomorphic-strings/)

```go
```

- [有效的字母异位词（Easy）](https://leetcode-cn.com/problems/valid-anagram/)

```go
```

- [字母异位词分组（Medium）](https://leetcode-cn.com/problems/group-anagrams/)

```go
```

- [找到字符串中所有字母异位词（Medium）](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)

```go
```

### 字符串 + 动态规划

- [通配符匹配（Hard）](https://leetcode-cn.com/problems/wildcard-matching/)

```go
```

## Part 1

### 最短路

- [网络延迟时间（Medium）](https://leetcode-cn.com/problems/network-delay-time/)

```go
```

- [阈值距离内邻居最少的城市（Medium）](https://leetcode-cn.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/)

```go
```

- [Dijkstra 求最短路 II （Easy）（ACWing）](https://www.acwing.com/problem/content/852/)

```go
```

### 最小生成树

- [连接所有点的最小费用（Medium）](https://leetcode-cn.com/problems/min-cost-to-connect-all-points/)

```go
```

## Part 2

### 字符串基础知识

http://c.biancheng.net/view/18.html

字符串中的每一个元素叫做“字符”，在遍历或者单个获取字符串元素时可以获得字符。

Go语言的字符有以下两种：
一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

- [字符串转换整数 (atoi) （Medium）](https://leetcode-cn.com/problems/string-to-integer-atoi/)

```go
func myAtoi(s string) int {
    // '0' (type rune) 
    // "0" (type untyped string)
    // fmt.Println(s[0]) // 42
    // fmt.Println(reflect.TypeOf(s[0])) // uint8
    // fmt.Println(rune(s[0]) - '0') // uint8 => int32 
    // 字符串转整数模板
    index := 0
    // 1. 丢弃前导空格
    for index < len(s) && s[index] == ' ' {
        index++
    }
    // 2. 判断符号
    sign := 1
        if index < len(s) && (s[index] == '+' || s[index] == '-') {
        if s[index] == '-' { sign = sign * -1 }
        index++
    }
    // 3. 处理数字
    val := 0
    // ASCII table
    // ASCII码 '0'-'9'是相连的
    for index < len(s) && int(s[index]) >= int('0') && int(s[index]) <= int('9') {
        // if 数值范围
        // if (val * 10 + s[index] - '0') > 2147483648 {
        // 移项
        if val > ( math.MaxInt32 - int(s[index]) + int('0') ) / 10 {
        if sign == -1 { return -2147483648 } // −2^31
        return 2147483647 // 2^31 − 1
        }
        val = val * 10 + int(s[index]) - int('0')
        index++
    }
    // 4. 终止条件：遇到非数字停止
    // 已经体现在while循环中
    return val * sign
}
```

### Rabin-Karp 字符串哈希算法

- 引入

```go
选用的 Hash 函数：

把字符串看作一个 b 进制数（一个多项式），计算它（在十进制下）对 p 取模的值

举例：

取 b = 131，p = 2^64

字符串 foobar 的 Hash 值为（a=1，b=2，f=6，o=15，r=18）

(6*131^5 + 15*131^4 + 15*131^3 + 2*131^2 + 1*131^1 + 18*131^0) mod 2^64


选取b和p的值决定了Hash函数的质量

根据经验，b=131，13331等，p为大质数，冲突概率极小

Hash值相等时，可以再对比一下两个字符串，避免Hash碰撞问题

c++可以自然溢出，其他语言需要取模。
```

如何快速计算一个子串的hash值？

```go
类比十进制数 s = "0123":
计算前 i 个数字的 H 值
H[0] = 0 // 初始化
H[1] = Hash(s[0...1-1]) = Hash[s[0...0]] = 0
H[2] = Hash(s[0...2-1]) = Hash[s[0...1]] = H[1] * 10^0 + 1 = 1
H[3] = Hash(s[0...3-1]) = Hash[s[0...2]] = H[2] * 10^1 + 2 = 12
H[4] = Hash(s[0...4-1]) = Hash[s[0...3]] = H[3] * 10^2 + 3 = 123
H[i] = Hash(s[0...i-1]) = H[i - 1] * 10^(i-2) (i > 1)

// 推导过程
s = "foobar"

先计算6个前缀子串的Hash值，O(n):
H[i] = Hash(s[0...i-1]) = (H[i - 1] * b + s[i - 1]) * mod p

计算子串 oba 的 Hash 值：
相当于b进制下两个数做减法（H[5] - H[2] * b^3）
 fooba
-fo000
------
   oba

假设下标从0开始，
Hash(oba) =>
Hash(s[2..4]) =>（H[5] - H[2] * b^3）
归纳为：
Hash(s[l...r]) => ( H[r+1] - H[l] * b^(r-l+1) ) % p, O(1)
```

- [实现 strStr() （Easy）](https://leetcode-cn.com/problems/implement-strstr/)

```go
func strStr(haystack string, needle string) int {
    s := haystack // 原始字符串
    t := needle // 目标字符串
    // RKHash模版（取模法）
    if len(t) == 0 { return 0 }
    n := len(s)
    m := len(t)
    s = " " + s
    t = " " + t

    var p int64 = 1e9 + 7 // 10^9 + 7 是一个质数
    var tHash int64 = 0 // java long类型，int64
    for i := 1; i <= m; i++ {
        tHash = (tHash * 131 + int64(t[i] - 'a') + 1) % p
    }
    // 模版
    sHash := make([]int64, n + 1)
    sHash[0] = 0
    p131 := make([]int64, n + 1) // 131的次幂
    p131[0] = 1
    for i := 1; i <= n; i++ {
        // 计算s前i个字符的hash值
        sHash[i] = (sHash[i - 1] * 131 + int64(s[i] - 'a') + 1) % p
        // 预存次幂，节省时间
        p131[i] = p131[i - 1] * 131 % p
    }
    // hello
    // ll
    for i := m; i <= n; i++ { // 滑动窗结尾
        // s[i-m+1...i] 与 t[1...m] 是否相等
        //if calcHash(sHash, p131, p, i-m+1, i) == tHash { // 不考虑碰撞
        if calcHash(sHash, p131, p, i-m+1, i) == tHash && s[i-m+1:i+1] == t[1:] { // 考虑碰撞
            return i - m //+ 1 - 1 // 下标变回从0开始
        }
    }
    return -1
}
// 模板：O(1)得到子串[l..r]的Hash值
func calcHash(H,p131 []int64, p int64, l,r int) int64 {
    // 求 hello 的子串的hash值
    //  h  e  l l
    // -h  e  0 0
    // =      l l
    //   (l-1)l r => 下标位置
    // r - l + 1 => 长度次幂
    // ？+ p % p 是为了控制 0 ～ p 之间
    return ( (H[r] - H[l - 1] * p131[r - l + 1]) % p + p) % p
}
```

- [重复叠加字符串匹配（Medium）](https://leetcode-cn.com/problems/repeated-string-match/)

```go
```

### 回文串系列

- [验证回文串（Easy）](https://leetcode-cn.com/problems/valid-palindrome/)

```go
func isPalindrome(s string) bool {
    var t string
    for _,ch := range s {
        //fmt.Println(reflect.TypeOf(ch)) // int32
        //fmt.Println(reflect.TypeOf('0')) // int32
        // 只考虑字符和数字
        if (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || ch >= 'a' && ch <= 'z' {
            if ch >= 'A' && ch <= 'Z' {
                ch = ch - 'A' + 'a'
            }
            t = t + string(ch)
        }
    }
    l := 0
    r := len(t) - 1
    for l < r {
        if t[l] != t[r] {return false}
        l++
        r--
    }
    return true
}
// 效率较低 192ms
```

- [验证回文字符串 Ⅱ（Easy）（贪心 + 验证）](https://leetcode-cn.com/problems/valid-palindrome-ii/)

```go
func validPalindrome(s string) bool {
    return validPalindromeHelper(s, 0, len(s) - 1, true)
}

func validPalindromeHelper(s string, l,r int, canDelete bool) bool {
    for l < r {
        if s[l] == s[r] {
            l++
            r--
        } else {
            if canDelete { // 能删除，尝试删除一个位置
                return validPalindromeHelper(s, l+1, r, false) || validPalindromeHelper(s, l, r-1, false)
            } else {
                return false
            }
        }
    }
    return true
}
```

- [最长回文子串（Medium）](https://leetcode-cn.com/problems/longest-palindromic-substring/)

**解法1**
```go
func longestPalindrome(s string) string {
    if len(s) == 0 {return ""}
    // 枚举中点，向两边扩展，考虑奇偶
    n := len(s)
    s = " " + s
    anslen := 0
    ansstart := 0
    // 中心是一个字符的情况，比如aba
    for center := 1; center <= n; center++ {
        l := center - 1
        r := center + 1
        for l > 0 && r <= n && s[l] == s[r] {
            l--
            r++
        }
        // l+1 ~ r-1 // 最后l减1、r加1之前是回文串，需要加回去
        // (r-1) - (l+1) + 1 = r - l - 1 // 实际长度
        if r - l - 1 > anslen { // 比之前的结果长，则更新
            anslen = r - l - 1
            ansstart = l + 1 // 加回来
        }
    }
    // 中心是两个字符（一个空），比如abba
    for center := 1; center < n; center++ { // n-1,n
        l := center
        r := center + 1
        for l > 0 && r <= n && s[l] == s[r] {
            l--
            r++
        }
        // l+1 ~ r-1 // 最后l减1、r加1之前是回文串，需要加回去
        // (r-1) - (l+1) + 1 = r - l - 1 // 实际长度
        if r - l - 1 > anslen { // 比之前的结果长，则更新
            anslen = r - l - 1
            ansstart = l + 1 // 加回来
        }
    }
    return  s[ansstart:ansstart + anslen] // 前闭后开
}
```

**解法2 二分答案+RKHash**

> 中间向两边扩张 O(n2)
加入二分 + Rabin-Karp 优化，O(nlogn)

```go
func longestPalindrome(s string) string {
    if len(s) == 0 {return ""}
    // 枚举中点，向两边扩展，考虑奇偶
    n := len(s)
    s = " " + s
    
    // RKHash模版
    var p int64 = 1e9 + 7
    preH := make([]int64, n+1) // 前缀Hash
    preH[0] = 0
    p131 := make([]int64, n+1) // 131次幂
    p131[0] = 1
    for i := 1; i <= n; i++ {
        preH[i] = (preH[i - 1] * 131 + int64(s[i] - 'a') + 1) % p
        p131[i] = p131[i - 1] * 131 % p
    }
    sufH := make([]int64, n+2) // 后缀Hash（反着读就是前缀字符串）
    sufH[n+1] = 0
    for i := n; i >= 1; i-- {
        sufH[i] = (sufH[i + 1] * 131 + int64(s[i] - 'a') + 1) % p
    }

    anslen := 0
    ansstart := 0
    // 中心是一个字符的情况，比如aba
    for center := 1; center <= n; center++ {
        // 二分查找，从当前字符串向两边可扩展的最大长度
        lenL := 0 // 下界
        lenR := n // 上界
        l := 0
        r := 0
        len := 0
        for lenL < lenR { // 二分模版部分
            len = (lenL + lenR + 1) >> 1 // mid
            l = center - len
            r = center + len
            if isPalindrome(s, n, preH, sufH, p131, p, l, r) { // 是回文，可以更长
                lenL = len
            } else {
                lenR = len - 1
            }
        }
        // 两侧最多扩LenL，再加一个中心
        if lenL * 2 + 1 > anslen {
            anslen = lenL * 2 + 1
            ansstart = center - lenL
        }
    }
    // 中心是两个字符（一个空），比如abba
    for center := 1; center < n; center++ { // n-1,n
        // 二分查找，从当前字符串向两边可扩展的最大长度
        lenL := 0 // 下界
        lenR := n // 上界
        l := 0
        r := 0
        len := 0
        for lenL < lenR { // 二分模版部分
            len = (lenL + lenR + 1) >> 1 // mid
            // 代入abba // 走len=0步，为空串；
            l = center - len + 1 // l～center
            r = center + len // center+1～r
            if isPalindrome(s, n, preH, sufH, p131, p, l, r) { // 是回文，可以更长
                lenL = len
            } else {
                lenR = len - 1
            }
        }
        // 两侧分别是l～center和center+1～r
        if lenL * 2 > anslen {
            anslen = lenL * 2
            ansstart = center - lenL + 1
        }
    }
    return  s[ansstart:ansstart + anslen] // 前闭后开
}
// RKHash判定是否回文
func isPalindrome(s string, n int, preH,sufH,p131 []int64, p int64, l, r int) bool {
    if l < 1 || r > n { return false }
    if l > r { return true }
    return calcPre(preH, p131, p, l, r) == calcSuf(sufH, p131, p, l, r)
}

// 模板：O(1)得到子串[l..r]的Hash值
func calcPre(preH,p131 []int64, p int64, l,r int) int64 {
    return ((preH[r] - preH[l - 1] * p131[r - l + 1]) %p + p) %p
}

// 模板：O(1)得到子串[l..r]反着读的Hash值
func calcSuf(sufH,p131 []int64, p int64, l,r int) int64 {
    return ((sufH[l] - sufH[r + 1] * p131[r - l + 1]) %p + p) %p
}
```

### 字符串 + 动态规划

- [正则表达式匹配（Hard）](https://leetcode-cn.com/problems/regular-expression-matching/)

```go
// 待补充注释
func isMatch(s string, p string) bool {
    n := len(s)
    m := len(p)
    s = " " + s
    p = " " + p
    // 定义前i个和前j个是否匹配
    f := make([][]bool, n + 1)
    for i := range f {
        f[i] = make([]bool, m + 1)
    }
    f[0][0] = true
    for i := 2; i<= m && p[i] == '*'; i+=2 {
        f[0][i] = true // 
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if p[j] == '.' { // i 和 j 一定能匹配
                f[i][j] = f[i - 1][j - 1]
            } else if p[j] == '*' {
                f[i][j] = f[i][j - 2] // 不要 _* 配0个
                if p[j - 1] == '.' || s[i] == p[j - 1] { // 让 _* 的 _ 去配 s[i] ?
                    f[i][j] = f[i][j] || f[i - 1][j]
                }
            } else {
                // 前i个与前j个匹配，一定是前 i-1 匹配 j-1 同时 i == j
                f[i][j] = f[i - 1][j - 1] && s[i] == p[j]
            }
        }
    }
    return f[n][m]
}
```

- [不同的子序列（Hard）](https://leetcode-cn.com/problems/distinct-subsequences/)

```go
func numDistinct(s string, t string) int {
    // 定义s前i个字符里，有几个子序列等于t的前几个字符
    n := len(s)
    m := len(t)
    s = " " + s
    t = " " + t
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, m + 1)
    }
    for i := 0; i<= n; i++ {
        f[i][0] = 1 // s前i个字符出现空串的次数是1
    }
    // 根据i-1的状态推i的状态
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            f[i][j] = f[i - 1][j] // 不要i这个位置的字符
            // 要i这个位置的字符
            if s[i] == t[j] {
                // 如果字符相等的话，就要加上f[i - 1][j - 1]的次数
                f[i][j] = f[i][j] + f[i - 1][j - 1]
            }
        }
    }
    return f[n][m]
}
```
