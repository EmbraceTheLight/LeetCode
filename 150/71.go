// 71. 简化路径
/*
 '..' 表示将目录切换到上一级（指向父目录）。
任意多个连续的斜杠（即，'//' 或 '///'）都被视为单个斜杠 '/'。
任何其他格式的点（例如，'...' 或 '....'）均被视为有效的文件/目录名称。
返回的 简化路径 必须遵循下述格式：
始终以斜杠 '/' 开头。
两个目录名之间必须只有一个斜杠 '/' 。
最后一个目录名（如果存在）不能 以 '/' 结尾。
此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
返回简化后得到的 规范路径 。

示例 1：
输入：path = "/home/"
输出："/home"
解释：
应删除尾随斜杠。

示例 2：
输入：path = "/home//foo/"
输出："/home/foo"
解释：
多个连续的斜杠被单个斜杠替换。

示例 3：
输入：path = "/home/user/Documents/../Pictures"
输出："/home/user/Pictures"
解释：
两个点 ".." 表示上一级目录（父目录）。

示例 4：
输入：path = "/../"
输出："/"
解释：
不可能从根目录上升一级目录。

示例 5：
输入：path = "/.../a/../b/c/../d/./"
输出："/.../b/d"
解释：
"..." 在这个问题中是一个合法的目录名。

提示：
1 <= path.length <= 3000
path 由英文字母，数字，'.'，'/' 或 '_' 组成。
path 是一个有效的 Unix 风格绝对路径。
*/
package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	paths := make([]string, 0)
	p := strings.Split(path, "/")
	for _, part := range p {
		if part == "" || part == "." {
			continue
		} else if part == ".." { // 上一级目录
			if len(paths) > 0 {
				paths = paths[:len(paths)-1]
			}
		} else {
			paths = append(paths, part)
		}
	}
	return "/" + strings.Join(paths, "/")
}

// Test Case1: "/home/"	Output: "/home"
// Test Case2: "/home//foo/"	Output: "/home/foo"
// Test Case3: "/home/user/Documents/../Pictures"	Output: "/home/user/Pictures"
// Test Case4: "/../"	Output: "/"
// Test Case5: "/.../a/../b/c/../d/./"	Output: "/.../b/d"
func main() {
	var path string
	fmt.Scan(&path)
	fmt.Println(simplifyPath(path))
}
