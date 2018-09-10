package main

func main() {
	friendRelation := createFriendCircle()
}

func createFriendCircle() map[string][]string {
	relationMap := make(map[string][]string)

	relationMap["you"] = []string{"bob", "alice", "claire"}
	relationMap["bob"] = []string{"anuj", "peggy"}
	relationMap["alice"] = []string{"peggy"}
	relationMap["claire"] = []string{"jonny", "thom"}

	return relationMap
}

func findNeedFriend(friendRelation map[string][]string) {
	if len(friendRelation) == 0 {
		return
	}

	searched := make(map[string]bool)
	for {

	}
}
