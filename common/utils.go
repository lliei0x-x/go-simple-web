package common

import (
	"io/ioutil"
	"regexp"
	"strings"
)

//GetAllFiles ...
func GetAllFiles(dirPth string, temps []string, match string) (files []string, err error) {
	// 拼接上一个目录下的文件路径
	for _, temp := range temps {
		files = append(files, temp)
	}
	fileInfos, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			// 目录, 递归遍历
			files, _ = GetAllFiles(dirPth+`/`+fileInfo.Name(), files, match)
		} else {
			// 过滤指定格式
			if match != "" {
				ok := strings.Contains(fileInfo.Name(), match)
				if ok {
					files = append(files, dirPth+`/`+fileInfo.Name())
				}
			} else {
				files = append(files, fileInfo.Name())
			}
		}
	}
	return
}

// MatchModelName Regular match name and tag by filename
func MatchModelName(fileName string) (name string) {
	pattern := `([^/\\\\]+).html`
	fileNameRegexp := regexp.MustCompile(pattern)
	params := fileNameRegexp.FindStringSubmatch(fileName)
	name = params[1]
	return
}
