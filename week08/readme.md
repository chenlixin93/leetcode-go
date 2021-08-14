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

- [实现 strStr() （Easy）](https://leetcode-cn.com/problems/implement-strstr/)

```go
```

- [重复叠加字符串匹配（Medium）](https://leetcode-cn.com/problems/repeated-string-match/)

```go
```

### 回文串系列

- [验证回文串（Easy）](https://leetcode-cn.com/problems/valid-palindrome/)

```go
```

- [验证回文字符串 Ⅱ（Easy）（贪心 + 验证）](https://leetcode-cn.com/problems/valid-palindrome-ii/)

```go
```

- [最长回文子串（Medium）](https://leetcode-cn.com/problems/longest-palindromic-substring/)

> 中间向两边扩张 O(n2)
加入二分 + Rabin-Karp 优化，O(nlogn)

```go
```

### 字符串 + 动态规划

- [正则表达式匹配（Hard）](https://leetcode-cn.com/problems/regular-expression-matching/)

```go
```

- [不同的子序列（Hard）](https://leetcode-cn.com/problems/distinct-subsequences/)

```go
```
