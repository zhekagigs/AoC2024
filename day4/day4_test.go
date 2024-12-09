package main

import "testing"

func TestGetAnswer_EmptyBoard(t *testing.T) {
	lines := []string{}
	result := getAnswer(lines)
	if result != 0 {
		t.Errorf("Expected 0 for empty board, got %d", result)
	}
}

func TestGetAnswer_OneLine(t *testing.T) {
	lines := []string{
		"XMAS",
	}
	result := getAnswer(lines)
	if result != 1 {
		t.Errorf("Expected 1 for single X XMAS pattern, got %d", result)
	}
}

func TestGetAnswer_MultipleX(t *testing.T) {
	lines := []string{
		"XMAS",
		"XMAS",
		"XMAS",
	}
	result := getAnswer(lines)
	if result != 3 {
		t.Errorf("Expected 3 for multiple horizontal XMAS, got %d", result)
	}
}

func TestGetAnswer_XMASVert(t *testing.T) {
	lines := []string{
		"X",
		"M",
		"A",
		"S",
	}
	result := getAnswer(lines)
	if result != 1 {
		t.Errorf("Expected 1 for single XMAS pattern, got %d", result)
	}
}

func TestGetAnswer_XMASVert2(t *testing.T) {
	lines := []string{
		"XX",
		"MM",
		"AA",
		"SS",
	}
	result := getAnswer(lines)
	if result != 2 {
		t.Errorf("Expected 2 for single XMAS pattern, got %d", result)
	}
}

func TestGetAnswer_XMASVertHor(t *testing.T) {
	lines := []string{
		"XXMAS",
		"MMMMM",
		"AAMMM",
		"SSMMM",
	}
	result := getAnswer(lines)
	if result != 3 {
		t.Errorf("Expected 3 for single XMAS pattern, got %d", result)
	}
}

func TestGetAnswer_DownRight(t *testing.T) {
	lines := []string{
		"XXXXX",
		"SMSSS",
		"SSASS",
		"SSSSS",
	}
	result := getAnswer(lines)
	if result != 1 {
		t.Errorf("Expected 1 for multiple XMAS patterns, got %d", result)
	}
}

func TestGetAnswer_UpLeft(t *testing.T) {
	lines := []string{
		"XSXXX",
		"SMASS",
		"SSAMS",
		"SSSSX",
	}
	result := getAnswer(lines)
	if result != 2 {
		t.Errorf("Expected 1 for multiple XMAS patterns, got %d", result)
	}
}

func TestGetAnswer_BorderXMAS(t *testing.T) {
	lines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	result := getAnswer(lines)
	if result != 18 {
		t.Errorf("Expected 18 for XMAS patterns at borders, got %d", result)
	}
}

func TestGetAnswer2(t *testing.T) {
	lines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	result := getAnswer2(lines)
	if result != 9 {
		t.Errorf("Expected 9 for MAS in x shape, got %d", result)
	}
}

func TestGetAnswer2_small(t *testing.T) {
	lines := []string{
		"MXM",
		"XAX",
		"SXS",
	}
	result := getAnswer2(lines)
	if result != 1 {
		t.Errorf("Expected 1 for MAS in x shape, got %d", result)
	}
}
func TestGetAnswer2_big(t *testing.T) {
	lines := []string{
		"SMMSSSSXSMMXSXMMSASXSMSMXMXMMMSASXMASXMMMSSMMXMSMXAAMSAMXMSAXSXMMAMXXAXXAAMAXSMSAMXSXXXXXXMXSXAXSXXSSSMAXMMSMMMSMSAMXSAMXMAMSXSASMXXAMXMXSSX",
		"XAAAAAAASASAMXSAMXMAMAAXAMMSAXSAMAMASXMASAAAMMMMMSAMXMAXAMMMMXASMXMSMSMMMSMSXMASXMASXMSSSXAASXMASMMSASMMSMAAAAAMMXMXSMASMSSXAAMAXAAMMSXSAMXM",
		"SMMSMMMMSAMAMAMASAMXMXMMMSAMMMMAMMMAMXSXMXMMMAAAMAMXSXSSXSAAASMMMMAAAAXAAAXMAMAMXMASAXAAXMMAMAXAMAMMAMXAAMSSSMSSMAMSXMAMMAMMMSMMMSMSAAMMAMSA",
	}
	result := getAnswer2(lines)
	if result != 1 {
		t.Errorf("Expected 1 for MAS in x shape, got %d", result)
	}
	t.Log(result)
}
func TestGetAnswer_BorderXMASsmall(t *testing.T) {
	lines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
	}
	result := getAnswer(lines)
	if result != 18 {
		t.Errorf("Expected 18 for XMAS patterns at borders, got %d", result)
	}
}

func Test_NewBoard(t *testing.T) {
	lines := []string{
		"SXMASS",
		"MXMASM",
	}
	b := NewBoard(lines)
	t.Log(b)
	if len(b.board) != 2 {
		t.Errorf("Not expected board len %d", len(b.board))
	}
	if len(b.board[0]) != 6 {
		t.Errorf("Not expected row len %d", len(b.board[0]))
	}
	if b.lenX != 6 {
		t.Errorf("lenX is %d", b.lenX)
	}
	if b.lenY != 2 {
		t.Errorf("lenY is %d", b.lenY)
	}
}

func TestBoard_GetValue(t *testing.T) {
	lines := []string{
		"ABC",
		"DEF",
		"GHI",
	}
	b := NewBoard(lines)

	tests := []struct {
		name     string
		position Position
		want     rune
	}{
		{
			name:     "valid position middle",
			position: Position{x: 1, y: 1},
			want:     'E',
		},
		{
			name:     "valid position top left",
			position: Position{x: 0, y: 0},
			want:     'A',
		},
		{
			name:     "valid position bottom right",
			position: Position{x: 2, y: 2},
			want:     'I',
		},
		{
			name:     "invalid position x out of bounds",
			position: Position{x: 3, y: 1},
			want:     -1,
		},
		{
			name:     "invalid position y out of bounds",
			position: Position{x: 1, y: 3},
			want:     -1,
		},
		{
			name:     "invalid position negative x",
			position: Position{x: -1, y: 1},
			want:     -1,
		},
		{
			name:     "invalid position negative y",
			position: Position{x: 1, y: -1},
			want:     -1,
		},
		{
			name:     "invalid position both out of bounds",
			position: Position{x: 3, y: 3},
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.getValue(tt.position); got != tt.want {
				t.Errorf("Board.getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
