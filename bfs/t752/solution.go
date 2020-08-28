package t752

func openLock(deadends []string, target string) int {
	secrets := make([]string, 0)
	secrets = append(secrets, "0000")
	visited := make(map[string]struct{})
	step := 0
	for _, deadend := range deadends {
		visited[deadend] = struct{}{}
	}
	for len(secrets) != 0 {
		secretsLen := len(secrets)
		for i := 0; i < secretsLen; i++ {
			cur := secrets[0]
			secrets = secrets[1:]
			if _, ok := visited[cur]; ok {
				continue
			}
			visited[cur] = struct{}{}
			if cur == target {
				return step
			}
			for j := 0; j < 4; j++ {
				if _, ok := visited[turnUp(cur, j)]; !ok {
					secrets = append(secrets, turnUp(cur, j))
				}
				if _, ok := visited[turnDown(cur, j)]; !ok {
					secrets = append(secrets, turnDown(cur, j))
				}
			}
		}
		step++
	}
	return -1
}

func turnUp(s string, i int) string {
	sr := []rune(s)
	if sr[i] == '9' {
		sr[i] = '0'
	} else {
		sr[i]++
	}
	return string(sr)
}

func turnDown(s string, i int) string {
	sr := []rune(s)
	if sr[i] == '0' {
		sr[i] = '9'
	} else {
		sr[i]--
	}
	return string(sr)
}
