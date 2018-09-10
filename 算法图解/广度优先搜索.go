package main

import "fmt"

func main() {

	friendRelation := createFriendRelation()
	friend := findNeedFriend(friendRelation)
	if friend {
		fmt.Println("yes")
	}
}

func createFriendRelation() map[string][]string {
	relationMap := make(map[string][]string)

	relationMap["you"] = []string{"alice", "bob", "claire"}
	relationMap["bob"] = []string{"anuj", "peggq"}
	relationMap["alice"] = []string{"pegga"}
	relationMap["claire"] = []string{"thom", "jonne"}
	relationMap["anuj"] = []string{}
	relationMap["jonne"] = []string{"peggy"}
	relationMap["thom"] = []string{}
	relationMap["jonny"] = []string{}

	return relationMap
}

func IsNeedMan(manName string) bool {

	return manName[len(manName)-1:] == "y"
}

func findNeedFriend(friendMap map[string][]string) bool {

	searchList := friendMap["you"]
	if len(searchList) == 0 {
		fmt.Println("not have result")
		return false
	}

	searched := make(map[string]bool)
	for {
		person := searchList[0]
		searchList = searchList[1:]
		_, find := searched[person]

		if !find {
			if IsNeedMan(person) {
				fmt.Println(person, "is seller!")
				return true
			} else {
				// 若map没有元素，其不会append空字符串的
				searchList = append(searchList, friendMap[person]...)
				searched[person] = true
			}
		}

		if len(searchList) == 0 {
			break
		}
	}

	return false
}
