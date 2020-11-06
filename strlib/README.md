# strlib
字符串操作库

## 功能列表
### 整数字符串转切片
```
var (
    data = "1,2,3"
    var ret []int
)
_ = strlib.SplitInts(data, &ret)
```
只支持整数（包含有符号和无符号）转换
### 整数切片转字符串
```
var (
    data = []int{1, 2, 3}
)
res := strlib.JoinInts(data)
fmt.Println(res) // 输出1,2,3
```
只支持整数（包含有符号和无符号）转换