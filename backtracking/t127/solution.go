package t127

import "math"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	var graph [][]int
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

// false but my
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	res := new(int)
	*res = math.MaxInt32
	helper(beginWord, endWord, wordList, 0, res)
	if *res == math.MaxInt32 {
		return 0
	}
	return *res
}

func helper(curWord, endWord string, wordList []string, count int, res *int) {
	if curWord == endWord {
		if count < *res {
			*res = count + 1
		}
	}
	if len(wordList) == 0 {
		return
	}
	for i, word := range wordList {
		if isOneStep(curWord, word) {
			newWordList := make([]string, 0)
			newWordList = append(newWordList, wordList[:i]...)
			newWordList = append(newWordList, wordList[i+1:]...)
			helper(word, endWord, newWordList, count+1, res)
		}
	}
}

func isOneStep(w1, w2 string) bool {
	count := 0
	for i := range w1 {
		if w1[i] != w2[i] {
			count++
		}
	}
	return count == 1
}
