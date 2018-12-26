package types

var (
	//InitMoves = []Move{}
	InitMoves = []Move{
		{A: State{X: 1, Y: 9}, B: State{X: 2, Y: 10}},
		{A: State{X: 2, Y: 9}, B: State{X: 2, Y: 10}},
		{A: State{X: 9, Y: 13}, B: State{X: 10, Y: 13}},
		{A: State{X: 9, Y: 1}, B: State{X: 8, Y: 2}},
		{A: State{X: 11, Y: 9}, B: State{X: 11, Y: 10}},
		{A: State{X: 1, Y: 8}, B: State{X: 1, Y: 9}},
		{A: State{X: 2, Y: 10}, B: State{X: 1, Y: 11}},
		{A: State{X: 5, Y: 1}, B: State{X: 4, Y: 2}},
		{A: State{X: 1, Y: 4}, B: State{X: 2, Y: 4}},
		{A: State{X: 10, Y: 12}, B: State{X: 10, Y: 13}},
		{A: State{X: 7, Y: 12}, B: State{X: 8, Y: 12}},
		{A: State{X: 9, Y: 1}, B: State{X: 10, Y: 1}},
		{A: State{X: 10, Y: 2}, B: State{X: 11, Y: 3}},
		{A: State{X: 1, Y: 12}, B: State{X: 1, Y: 13}},
		{A: State{X: 1, Y: 1}, B: State{X: 2, Y: 2}},
		{A: State{X: 2, Y: 2}, B: State{X: 3, Y: 2}},
		{A: State{X: 5, Y: 1}, B: State{X: 6, Y: 1}},
		{A: State{X: 7, Y: 1}, B: State{X: 8, Y: 2}},
		{A: State{X: 9, Y: 2}, B: State{X: 10, Y: 2}},
		{A: State{X: 10, Y: 8}, B: State{X: 10, Y: 9}},
		{A: State{X: 2, Y: 1}, B: State{X: 2, Y: 2}},
		{A: State{X: 11, Y: 10}, B: State{X: 10, Y: 11}},
		{A: State{X: 2, Y: 2}, B: State{X: 1, Y: 3}},
		{A: State{X: 1, Y: 5}, B: State{X: 2, Y: 6}},
		{A: State{X: 10, Y: 4}, B: State{X: 10, Y: 5}},
		{A: State{X: 2, Y: 9}, B: State{X: 1, Y: 10}},
		{A: State{X: 10, Y: 12}, B: State{X: 11, Y: 13}},
		{A: State{X: 10, Y: 10}, B: State{X: 10, Y: 11}},
		{A: State{X: 2, Y: 12}, B: State{X: 3, Y: 12}},
		{A: State{X: 3, Y: 12}, B: State{X: 4, Y: 13}},
		{A: State{X: 11, Y: 8}, B: State{X: 10, Y: 9}},
		{A: State{X: 8, Y: 2}, B: State{X: 9, Y: 2}},
		{A: State{X: 1, Y: 6}, B: State{X: 2, Y: 7}},
		{A: State{X: 8, Y: 12}, B: State{X: 7, Y: 13}},
		{A: State{X: 3, Y: 1}, B: State{X: 3, Y: 2}},
		{A: State{X: 4, Y: 2}, B: State{X: 5, Y: 2}},
		{A: State{X: 11, Y: 2}, B: State{X: 11, Y: 3}},
		{A: State{X: 2, Y: 12}, B: State{X: 1, Y: 13}},
		{A: State{X: 1, Y: 3}, B: State{X: 2, Y: 3}},
		{A: State{X: 7, Y: 2}, B: State{X: 8, Y: 2}},
		{A: State{X: 10, Y: 3}, B: State{X: 10, Y: 4}},
		{A: State{X: 10, Y: 9}, B: State{X: 11, Y: 10}},
		{A: State{X: 3, Y: 13}, B: State{X: 4, Y: 13}},
		{A: State{X: 4, Y: 1}, B: State{X: 3, Y: 2}},
		{A: State{X: 11, Y: 6}, B: State{X: 11, Y: 7}},
		{A: State{X: 1, Y: 7}, B: State{X: 2, Y: 7}},
		{A: State{X: 1, Y: 8}, B: State{X: 2, Y: 9}},
		{A: State{X: 8, Y: 13}, B: State{X: 9, Y: 13}},
		{A: State{X: 10, Y: 5}, B: State{X: 11, Y: 6}},
		{A: State{X: 2, Y: 6}, B: State{X: 2, Y: 7}},
		{A: State{X: 11, Y: 10}, B: State{X: 11, Y: 11}},
		{A: State{X: 2, Y: 4}, B: State{X: 1, Y: 5}},
		{A: State{X: 10, Y: 8}, B: State{X: 11, Y: 9}},
		{A: State{X: 9, Y: 12}, B: State{X: 8, Y: 13}},
		{A: State{X: 1, Y: 10}, B: State{X: 1, Y: 11}},
		{A: State{X: 8, Y: 1}, B: State{X: 9, Y: 1}},
		{A: State{X: 10, Y: 1}, B: State{X: 11, Y: 2}},
		{A: State{X: 1, Y: 5}, B: State{X: 1, Y: 6}},
		{A: State{X: 2, Y: 2}, B: State{X: 2, Y: 3}},
		{A: State{X: 8, Y: 1}, B: State{X: 7, Y: 2}},
		{A: State{X: 11, Y: 12}, B: State{X: 10, Y: 13}},
		{A: State{X: 4, Y: 1}, B: State{X: 5, Y: 1}},
		{A: State{X: 11, Y: 9}, B: State{X: 10, Y: 10}},
		{A: State{X: 11, Y: 7}, B: State{X: 11, Y: 8}},
		{A: State{X: 10, Y: 10}, B: State{X: 11, Y: 11}},
		{A: State{X: 8, Y: 12}, B: State{X: 9, Y: 13}},
		{A: State{X: 1, Y: 2}, B: State{X: 1, Y: 3}},
		{A: State{X: 11, Y: 11}, B: State{X: 11, Y: 12}},
		{A: State{X: 2, Y: 7}, B: State{X: 1, Y: 8}},
		{A: State{X: 5, Y: 12}, B: State{X: 4, Y: 13}},
		{A: State{X: 11, Y: 2}, B: State{X: 10, Y: 3}},
		{A: State{X: 1, Y: 12}, B: State{X: 2, Y: 13}},
		{A: State{X: 10, Y: 7}, B: State{X: 11, Y: 7}},
		{A: State{X: 3, Y: 1}, B: State{X: 4, Y: 2}},
		{A: State{X: 1, Y: 6}, B: State{X: 1, Y: 7}},
		{A: State{X: 10, Y: 3}, B: State{X: 11, Y: 4}},
		{A: State{X: 1, Y: 9}, B: State{X: 2, Y: 9}},
		{A: State{X: 10, Y: 6}, B: State{X: 11, Y: 6}},
		{A: State{X: 10, Y: 3}, B: State{X: 11, Y: 3}},
		{A: State{X: 10, Y: 5}, B: State{X: 10, Y: 6}},
		{A: State{X: 1, Y: 10}, B: State{X: 2, Y: 10}},
		{A: State{X: 11, Y: 5}, B: State{X: 11, Y: 6}},
		{A: State{X: 1, Y: 13}, B: State{X: 2, Y: 13}},
		{A: State{X: 1, Y: 4}, B: State{X: 2, Y: 5}},
		{A: State{X: 1, Y: 11}, B: State{X: 2, Y: 11}},
		{A: State{X: 7, Y: 12}, B: State{X: 8, Y: 13}},
		{A: State{X: 2, Y: 12}, B: State{X: 2, Y: 13}},
		{A: State{X: 11, Y: 3}, B: State{X: 10, Y: 4}},
		{A: State{X: 1, Y: 1}, B: State{X: 2, Y: 1}},
		{A: State{X: 8, Y: 1}, B: State{X: 8, Y: 2}},
		{A: State{X: 10, Y: 11}, B: State{X: 10, Y: 12}},
		{A: State{X: 7, Y: 1}, B: State{X: 8, Y: 1}},
		{A: State{X: 4, Y: 12}, B: State{X: 4, Y: 13}},
		{A: State{X: 4, Y: 1}, B: State{X: 4, Y: 2}},
		{A: State{X: 1, Y: 2}, B: State{X: 2, Y: 2}},
		{A: State{X: 2, Y: 8}, B: State{X: 2, Y: 9}},
		{A: State{X: 1, Y: 5}, B: State{X: 2, Y: 5}},
		{A: State{X: 3, Y: 2}, B: State{X: 4, Y: 2}},
		{A: State{X: 10, Y: 7}, B: State{X: 10, Y: 8}},
		{A: State{X: 3, Y: 12}, B: State{X: 2, Y: 13}},
		{A: State{X: 7, Y: 13}, B: State{X: 8, Y: 13}},
		{A: State{X: 2, Y: 4}, B: State{X: 2, Y: 5}},
		{A: State{X: 11, Y: 11}, B: State{X: 10, Y: 12}},
		{A: State{X: 2, Y: 12}, B: State{X: 3, Y: 13}},
		{A: State{X: 11, Y: 4}, B: State{X: 10, Y: 5}},
		{A: State{X: 2, Y: 10}, B: State{X: 2, Y: 11}},
		{A: State{X: 11, Y: 1}, B: State{X: 10, Y: 2}},
		{A: State{X: 10, Y: 13}, B: State{X: 11, Y: 13}},
		{A: State{X: 11, Y: 5}, B: State{X: 10, Y: 6}},
		{A: State{X: 1, Y: 3}, B: State{X: 2, Y: 4}},
		{A: State{X: 9, Y: 1}, B: State{X: 9, Y: 2}},
		{A: State{X: 10, Y: 9}, B: State{X: 11, Y: 9}},
		{A: State{X: 6, Y: 1}, B: State{X: 7, Y: 1}},
		{A: State{X: 11, Y: 3}, B: State{X: 11, Y: 4}},
		{A: State{X: 2, Y: 3}, B: State{X: 2, Y: 4}},
		{A: State{X: 5, Y: 1}, B: State{X: 5, Y: 2}},
		{A: State{X: 7, Y: 12}, B: State{X: 7, Y: 13}},
		{A: State{X: 10, Y: 5}, B: State{X: 11, Y: 5}},
		{A: State{X: 10, Y: 1}, B: State{X: 9, Y: 2}},
		{A: State{X: 10, Y: 11}, B: State{X: 11, Y: 12}},
		{A: State{X: 10, Y: 11}, B: State{X: 11, Y: 11}},
		{A: State{X: 10, Y: 4}, B: State{X: 11, Y: 4}},
		{A: State{X: 10, Y: 1}, B: State{X: 11, Y: 1}},
		{A: State{X: 4, Y: 12}, B: State{X: 3, Y: 13}},
		{A: State{X: 2, Y: 1}, B: State{X: 3, Y: 2}},
		{A: State{X: 4, Y: 12}, B: State{X: 5, Y: 12}},
		{A: State{X: 2, Y: 11}, B: State{X: 1, Y: 12}},
		{A: State{X: 9, Y: 12}, B: State{X: 10, Y: 12}},
		{A: State{X: 3, Y: 1}, B: State{X: 4, Y: 1}},
		{A: State{X: 10, Y: 7}, B: State{X: 11, Y: 8}},
		{A: State{X: 1, Y: 3}, B: State{X: 1, Y: 4}},
		{A: State{X: 3, Y: 1}, B: State{X: 2, Y: 2}},
		{A: State{X: 10, Y: 6}, B: State{X: 11, Y: 7}},
		{A: State{X: 5, Y: 13}, B: State{X: 6, Y: 13}},
		{A: State{X: 10, Y: 9}, B: State{X: 10, Y: 10}},
		{A: State{X: 1, Y: 10}, B: State{X: 2, Y: 11}},
		{A: State{X: 2, Y: 5}, B: State{X: 2, Y: 6}},
		{A: State{X: 1, Y: 7}, B: State{X: 1, Y: 8}},
		{A: State{X: 2, Y: 3}, B: State{X: 1, Y: 4}},
		{A: State{X: 11, Y: 6}, B: State{X: 10, Y: 7}},
		{A: State{X: 11, Y: 8}, B: State{X: 11, Y: 9}},
		{A: State{X: 2, Y: 6}, B: State{X: 1, Y: 7}},
		{A: State{X: 9, Y: 12}, B: State{X: 9, Y: 13}},
		{A: State{X: 9, Y: 1}, B: State{X: 10, Y: 2}},
		{A: State{X: 10, Y: 2}, B: State{X: 11, Y: 2}},
		{A: State{X: 1, Y: 11}, B: State{X: 2, Y: 12}},
		{A: State{X: 10, Y: 1}, B: State{X: 10, Y: 2}},
		{A: State{X: 11, Y: 12}, B: State{X: 11, Y: 13}},
		{A: State{X: 2, Y: 11}, B: State{X: 2, Y: 12}},
		{A: State{X: 1, Y: 2}, B: State{X: 2, Y: 3}},
		{A: State{X: 8, Y: 12}, B: State{X: 8, Y: 13}},
		{A: State{X: 2, Y: 7}, B: State{X: 2, Y: 8}},
		{A: State{X: 11, Y: 7}, B: State{X: 10, Y: 8}},
		{A: State{X: 10, Y: 12}, B: State{X: 11, Y: 12}},
		{A: State{X: 1, Y: 9}, B: State{X: 1, Y: 10}},
		{A: State{X: 5, Y: 12}, B: State{X: 5, Y: 13}},
		{A: State{X: 3, Y: 12}, B: State{X: 4, Y: 12}},
		{A: State{X: 1, Y: 6}, B: State{X: 2, Y: 6}},
		{A: State{X: 10, Y: 12}, B: State{X: 9, Y: 13}},
		{A: State{X: 2, Y: 13}, B: State{X: 3, Y: 13}},
		{A: State{X: 7, Y: 1}, B: State{X: 7, Y: 2}},
		{A: State{X: 2, Y: 8}, B: State{X: 1, Y: 9}},
		{A: State{X: 1, Y: 4}, B: State{X: 1, Y: 5}},
		{A: State{X: 4, Y: 13}, B: State{X: 5, Y: 13}},
		{A: State{X: 1, Y: 1}, B: State{X: 1, Y: 2}},
		{A: State{X: 2, Y: 5}, B: State{X: 1, Y: 6}},
		{A: State{X: 1, Y: 8}, B: State{X: 2, Y: 8}},
		{A: State{X: 10, Y: 6}, B: State{X: 10, Y: 7}},
		{A: State{X: 1, Y: 11}, B: State{X: 1, Y: 12}},
		{A: State{X: 2, Y: 1}, B: State{X: 1, Y: 2}},
		{A: State{X: 10, Y: 8}, B: State{X: 11, Y: 8}},
		{A: State{X: 10, Y: 4}, B: State{X: 11, Y: 5}},
		{A: State{X: 1, Y: 7}, B: State{X: 2, Y: 8}},
		{A: State{X: 10, Y: 2}, B: State{X: 10, Y: 3}},
		{A: State{X: 8, Y: 1}, B: State{X: 9, Y: 2}},
		{A: State{X: 2, Y: 1}, B: State{X: 3, Y: 1}},
		{A: State{X: 4, Y: 12}, B: State{X: 5, Y: 13}},
		{A: State{X: 3, Y: 12}, B: State{X: 3, Y: 13}},
		{A: State{X: 4, Y: 1}, B: State{X: 5, Y: 2}},
		{A: State{X: 6, Y: 13}, B: State{X: 7, Y: 13}},
		{A: State{X: 10, Y: 10}, B: State{X: 11, Y: 10}},
		{A: State{X: 8, Y: 12}, B: State{X: 9, Y: 12}},
		{A: State{X: 11, Y: 4}, B: State{X: 11, Y: 5}},
		{A: State{X: 11, Y: 1}, B: State{X: 11, Y: 2}},
		{A: State{X: 9, Y: 12}, B: State{X: 10, Y: 13}},
		{A: State{X: 1, Y: 12}, B: State{X: 2, Y: 12}},
	}

	UpperSide = Side{
		Pos: UpperPos,
		WinStates: []State{
			{5, 13},
			{6, 13},
			{7, 13},
		},
		LoseStates: []State{
			{5, 1},
			{6, 1},
			{7, 1},
		},
	}
	LowerSide = Side{
		Pos: LowerPos,
		WinStates: []State{
			{5, 1},
			{6, 1},
			{7, 1},
		},
		LoseStates: []State{
			{5, 13},
			{6, 13},
			{7, 13},
		},
	}
)
