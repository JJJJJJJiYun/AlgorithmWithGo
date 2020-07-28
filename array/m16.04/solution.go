package m16_04

// 设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。
// 以下是井字游戏的规则：
// 玩家轮流将字符放入空位（" "）中。
// 第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
// "X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
// 当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
// 当所有位置非空时，也算为游戏结束。
// 如果游戏结束，玩家不允许再放置字符。
// 如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。
func tictactoe(board []string) string {
	dia, adia := 0, 0
	isFull := true
	// 使用一次遍历，计算行、列、对角线的和
	// 字符串的比较可以用相加或者异或来做
	for i := 0; i < len(board); i++ {
		row, col := 0, 0
		for j := 0; j < len(board); j++ {
			row += int([]byte(board[i])[j])
			col += int([]byte(board[j])[i])
			if []byte(board[i])[j] == ' ' {
				isFull = false
			}
		}
		if row == int('X')*len(board) || col == int('X')*len(board) {
			return "X"
		}
		if row == int('O')*len(board) || col == int('O')*len(board) {
			return "O"
		}
		dia += int([]byte(board[i])[i])
		adia += int([]byte(board[i])[len(board)-i-1])
	}
	if dia == int('X')*len(board) || adia == int('X')*len(board) {
		return "X"
	}
	if dia == int('O')*len(board) || adia == int('O')*len(board) {
		return "O"
	}
	if isFull {
		return "Draw"
	}
	return "Pending"
}
