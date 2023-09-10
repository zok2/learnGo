package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randomTeam(groupList [][]string) [][]string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 复制原始小组列表，避免修改原始数据
	groups := make([][]string, len(groupList))
	for i := range groupList {
		groups[i] = make([]string, len(groupList[i]))
		copy(groups[i], groupList[i])
	}

	var result [][]string

	for len(groups) >= 2 {
		// 随机选择2个或3个小组来组队
		teamSize := 2 + r.Intn(2) // 随机生成2或3
		var team []string

		for i := 0; i < teamSize; i++ {
			// 随机选择一个小组，并从中随机选择一个人
			selectedGroupIndex := r.Intn(len(groups))
			selectedGroup := groups[selectedGroupIndex]
			personIndex := r.Intn(len(selectedGroup))
			person := selectedGroup[personIndex]

			// 将选中的人从原小组中移除
			selectedGroup = append(selectedGroup[:personIndex], selectedGroup[personIndex+1:]...)
			groups[selectedGroupIndex] = selectedGroup

			// 如果一个小组没有人了，将其从列表中移除
			if len(selectedGroup) == 0 {
				groups = append(groups[:selectedGroupIndex], groups[selectedGroupIndex+1:]...)
			}

			team = append(team, person)
		}

		result = append(result, team)
	}

	return result
}

func main() {
	groupList := [][]string{
		{"小名", "小红", "小马", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"A", "B", "C", "D", "E"},
		{"建国", "建军", "建民", "建超", "建跃"},
		{"一", "二", "三", "四", "五"},
	}

	result := randomTeam(groupList)
	output := make([]string, len(result))

	for i, team := range result {
		output[i] = "[" + strings.Join(team, " ") + "]"
	}

	fmt.Println(strings.Join(output, ""))
}
