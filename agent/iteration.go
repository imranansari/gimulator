package agent

import (
	"github.com/alidadar7676/gimulator/types"
	"log"
	"crypto/md5"
	"encoding/base64"
)

var (
	dirx = []int{1, 1, 0, -1, -1, -1, 0, 1}
	diry = []int{0, 1, 1, 1, 0, -1, -1, -1}
)

type iteration struct {
	moveByte   []byte
	ball       types.State
	winStates  []types.State
	loseStates []types.State
	playground [][]map[int]struct{}
	hashTable map[string]*gameState
}

func (it *iteration) validMoves() []types.Move {
	validMoves := make([]types.Move, 0)
	x := it.ball.X
	y := it.ball.Y

	exist := make([]bool, 10)
	for i := range it.playground[x][y] {
		exist[i] = true
	}

	for i := 0; i < 8; i++ {
		if exist[i] {
			continue
		}
		xx := x + dirx[i]
		yy := y + diry[i]
		validMoves = append(validMoves, types.Move{
			A: it.ball,
			B: types.State{X: xx, Y: yy},
		})
	}
	return validMoves
}

func (it *iteration) next(move types.Move) {
	toggleMoveByte(it.moveByte, move)

	da, db := getDirection(move)
	it.playground[move.A.X][move.A.Y][da] = struct{}{}
	it.playground[move.B.X][move.B.Y][db] = struct{}{}

	it.ball = move.B
}

func (it *iteration) prev(move types.Move) {
	toggleMoveByte(it.moveByte, move)

	da, db := getDirection(move)
	delete(it.playground[move.A.X][move.A.Y], da)
	delete(it.playground[move.B.X][move.B.Y], db)

	it.ball = move.A
}

func (it *iteration) hash() string {
	a := md5.Sum(it.moveByte)
	b := a[0:]
	return base64.StdEncoding.EncodeToString(b)
}

func (it *iteration) hasPrice() bool {
	x := it.ball.X
	y := it.ball.Y
	if len(it.playground[x][y]) > 1 {
		return true
	}
	return false
}

func newIteration(world types.World, name string) *iteration {
	player := types.Player{}
	if world.Player1.Name == name {
		player = world.Player1
	} else {
		player = world.Player2
	}

	pg := newPlayground(world.Moves, world.Width, world.Height)
	mb := newMoveByte(world.Moves)

	return &iteration{
		ball:       world.BallPos,
		winStates:  player.Side.WinStates,
		loseStates: player.Side.LoseStates,
		playground: pg,
		moveByte:   mb,
	}
}

func newPlayground(moves []types.Move, width, height int) [][]map[int]struct{} {
	log.Println("world size: ", width, height)

	pg := make([][]map[int]struct{}, width+5)
	for i := 0; i < width+5; i++ {
		pg[i] = make([]map[int]struct{}, height+5)
		for j := 0; j < height+5; j++ {
			pg[i][j] = make(map[int]struct{})
		}
	}

	for _, mv := range moves {
		da, db := getDirection(mv)

		pg[mv.A.X][mv.A.Y][da] = struct{}{}
		pg[mv.B.X][mv.B.Y][db] = struct{}{}
	}

	return pg
}

func getDirection(mv types.Move) (int, int) {

	var dira int
	for i := 0; i < 8; i++ {
		x := mv.A.X + dirx[i]
		y := mv.A.Y + diry[i]

		if x == mv.B.X && y == mv.B.Y {
			dira = i
		}
	}

	var dirb int
	for i := 0; i < 8; i++ {
		x := mv.B.X + dirx[i]
		y := mv.B.Y + diry[i]

		if x == mv.A.X && y == mv.A.Y {
			dirb = i
		}
	}

	return dira, dirb
}

func newMoveByte(moves []types.Move) []byte {
	moveByte := make([]byte, 200)
	for _, mv := range moves {
		toggleMoveByte(moveByte, mv)
	}
	return moveByte
}

func toggleMoveByte(moveByte []byte, move types.Move) {
	ind := moveToInt(move)
	div := ind / 8
	mod := ind % 8

	var sum byte = 1 << mod
	moveByte[div] ^= sum
}

func moveToInt(move types.Move) uint {
	a := move.A
	b := move.B

	if a.Y > b.Y {
		a, b = b, a
	} else if a.Y == b.Y && a.X > b.X {
		a, b = b, a
	}

	num := uint(0)
	num += uint((a.Y - 1) * 4 * 11)
	num += uint((a.X - 1) * 4)

	switch {
	case a.X+1 == b.X && a.Y == b.Y:
		num += 1
	case a.X+1 == b.X && a.Y+1 == b.Y:
		num += 2
	case a.X == b.X && a.Y+1 == b.Y:
		num += 3
	case a.X-1 == b.X && a.Y+1 == b.Y:
		num += 4
	}
	return num
}
