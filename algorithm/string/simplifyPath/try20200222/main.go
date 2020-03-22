package main

import "strings"

func main() {

}

/*
思路：按/分隔字符串，对不同的子串进行处理，然后join by /
伪代码：
strSli := strings.Split(path, "/")
rstSli := make([]string,len(strSli),0)
for _,str := range strSli {
	switch str {
	case "":
	case ".":
	case "..":
		if len(rstSli) > 0 {
			rstSli = rstSli[:len(rstSli)-1]
		}
	default:
		rstSli = append(rstSli,str)
	}
}
return "/"+strings.Join(rstSli,"/")
*/
func simplifyPath(path string) string {
	strSli := strings.Split(path, "/")
	rstSli := make([]string, 0, len(strSli))
	for _, str := range strSli {
		switch str {
		case "":
		case ".":
		case "..":
			if len(rstSli) > 0 {
				rstSli = rstSli[:len(rstSli)-1]
			}
		default:
			rstSli = append(rstSli, str)
		}
	}
	return "/" + strings.Join(rstSli, "/")
}
