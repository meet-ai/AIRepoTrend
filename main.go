package main

import "AIRepoTrend/codegpt"

func main() {
	println("AI Agents to get github code, to generate data")
	repos := getTodayTrends()

	for _, r := range repos {
		filepath := gitclone(r)
		codegpt.Codegpt(r, filepath)
	}
}

func getTodayTrends() []string {
	return []string{""}
}

// 拷贝到本地形成目录
func gitclone(repo string) string {
	return ""
}
