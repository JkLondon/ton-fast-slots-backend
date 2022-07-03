package slot

import (
	"Casino/pkg/utils"
)

type block struct {
	Name        string
	Multiplayer float64
	Rate        int
}

type Manager struct {
	Blocks []block
}

func InitManager() Manager {
	m := Manager{}
	m.Blocks = append(m.Blocks, block{
		Name:        "TON",
		Multiplayer: 70,
		Rate:        5,
	})
	m.Blocks = append(m.Blocks, block{
		Name:        "ETH",
		Multiplayer: 30,
		Rate:        10,
	})
	m.Blocks = append(m.Blocks, block{
		Name:        "BTC",
		Multiplayer: 20,
		Rate:        15,
	})
	m.Blocks = append(m.Blocks, block{
		Name:        ".sol",
		Multiplayer: 10,
		Rate:        20,
	})
	m.Blocks = append(m.Blocks, block{
		Name:        ".rs",
		Multiplayer: 5,
		Rate:        25,
	})
	m.Blocks = append(m.Blocks, block{
		Name:        ".go",
		Multiplayer: 4,
		Rate:        30,
	})
	m.Blocks = append(m.Blocks, block{
		Name:        ".cpp",
		Multiplayer: 3,
		Rate:        35,
	})
	m.Blocks = append(m.Blocks, block{
		Name:        ".js",
		Multiplayer: 2,
		Rate:        40,
	})
	return m
}

type ScrollResult struct {
	Name string
	Rate int
}

func (m *Manager) Scroll() (totalMultiply float64, res []ScrollResult) {
	res = make([]ScrollResult, 3)
	toMultiply := make([]string, 3)
	for i := 0; i < 3; i++ {
		randomizedBlock := m.getBlock(int(utils.GenRandNum()))
		toMultiply[i] = randomizedBlock.Name
		res[i] = ScrollResult{
			Name: randomizedBlock.Name,
			Rate: randomizedBlock.Rate,
		}
	}

	totalMultiply = m.getMultiply(toMultiply)
	return
}

func (m *Manager) getMultiply(v []string) float64 {
	start := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] != start {
			return 0
		}
	}
	for _, r := range m.Blocks {
		if r.Name == start {
			return r.Multiplayer
		}
	}
	return 0
}

func (m *Manager) getBlock(value int) block {
	for _, r := range m.Blocks {
		if r.Rate >= value {
			return r
		}
	}
	return m.Blocks[len(m.Blocks)-1]
}

//func (m *Manager) Scroll() (totalMultiply float64, res [][]ScrollResult) {
//	res = make([][]ScrollResult, 3)
//	horizontal := make([][]string, 3)
//	vertical := make([][]string, 3)
//	diagonalLeft := make([]string, 0)
//	diagonalRight := make([]string, 0)
//	for i := 0; i < 9; i++ {
//		randomizedBlock := m.getBlock(int(utils.GenRandNum()))
//		row := 0
//		if i >= 3 && i <= 5 {
//			row = 1
//		} else if i > 5 && i <= 8 {
//			row = 2
//		}
//		switch i {
//		case 0:
//			horizontal[0] = append(horizontal[0], randomizedBlock.Name)
//			vertical[0] = append(vertical[0], randomizedBlock.Name)
//			diagonalLeft = append(diagonalLeft, randomizedBlock.Name)
//		case 1:
//			horizontal[0] = append(horizontal[0], randomizedBlock.Name)
//			vertical[1] = append(vertical[1], randomizedBlock.Name)
//		case 2:
//			horizontal[0] = append(horizontal[0], randomizedBlock.Name)
//			vertical[2] = append(vertical[2], randomizedBlock.Name)
//			diagonalRight = append(diagonalRight, randomizedBlock.Name)
//		case 3:
//			horizontal[1] = append(horizontal[1], randomizedBlock.Name)
//			vertical[0] = append(vertical[0], randomizedBlock.Name)
//		case 4:
//			horizontal[1] = append(horizontal[1], randomizedBlock.Name)
//			vertical[1] = append(vertical[1], randomizedBlock.Name)
//			diagonalLeft = append(diagonalLeft, randomizedBlock.Name)
//			diagonalRight = append(diagonalRight, randomizedBlock.Name)
//		case 5:
//			horizontal[1] = append(horizontal[1], randomizedBlock.Name)
//			vertical[2] = append(vertical[2], randomizedBlock.Name)
//		case 6:
//			horizontal[2] = append(horizontal[2], randomizedBlock.Name)
//			vertical[0] = append(vertical[0], randomizedBlock.Name)
//		case 7:
//			horizontal[2] = append(horizontal[2], randomizedBlock.Name)
//			vertical[1] = append(vertical[1], randomizedBlock.Name)
//		case 8:
//			horizontal[2] = append(horizontal[2], randomizedBlock.Name)
//			vertical[2] = append(vertical[2], randomizedBlock.Name)
//			diagonalLeft = append(diagonalLeft, randomizedBlock.Name)
//			diagonalRight = append(diagonalRight, randomizedBlock.Name)
//		}
//		res[row] = append(res[row], ScrollResult{
//			Name: randomizedBlock.Name,
//			Rate: randomizedBlock.Rate,
//		})
//	}
//	totalMultiply += m.getMultiply(horizontal[0])
//	totalMultiply += m.getMultiply(horizontal[1])
//	totalMultiply += m.getMultiply(horizontal[2])
//	totalMultiply += m.getMultiply(vertical[0])
//	totalMultiply += m.getMultiply(vertical[1])
//	totalMultiply += m.getMultiply(vertical[2])
//	totalMultiply += m.getMultiply(diagonalLeft)
//	totalMultiply += m.getMultiply(diagonalRight)
//	fmt.Printf("horizontal 0 %v %f\n", horizontal[0], m.getMultiply(horizontal[0]))
//	fmt.Printf("horizontal 1 %v %f\n", horizontal[1], m.getMultiply(horizontal[1]))
//	fmt.Printf("horizontal 2 %v %f\n", horizontal[2], m.getMultiply(horizontal[2]))
//	fmt.Printf("vertical 0 %v %f\n", vertical[0], m.getMultiply(vertical[0]))
//	fmt.Printf("vartical 1 %v %f\n", vertical[1], m.getMultiply(vertical[1]))
//	fmt.Printf("vartical 2 %v %f\n", vertical[2], m.getMultiply(vertical[2]))
//	fmt.Printf("diagonalLeft %v %f\n", diagonalLeft, m.getMultiply(diagonalLeft))
//	fmt.Printf("diagonalRight %v %f\n", diagonalRight, m.getMultiply(diagonalRight))
//	fmt.Printf("totalMultiply %f\n", totalMultiply)
//	return
//}
