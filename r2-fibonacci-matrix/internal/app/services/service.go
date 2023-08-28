package services

type (
	Service struct{}

	FibonacciMatrixService interface {
		GenerateMatrix(rows, cols int) [][]int
	}
)

func NewMatrixService() *Service {
	return &Service{}
}

func (s *Service) GenerateMatrix(rows, cols int) [][]int {
	// obtain fibonacci sequence
	fibonacciSequence := getFibonacciNumbers(rows*cols - 1)

	// create matrix initial structure
	matrix := createMatrix(rows, cols)

	// fill matrix with fibonacci sequence
	fillMatrix(matrix, fibonacciSequence, rows, cols)

	return matrix
}

func createMatrix(rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	return matrix
}

func getFibonacciNumbers(depth int) []int {
	sequence := make([]int, depth+1, depth+2)
	if depth < 2 {
		sequence = sequence[0:2]
	}
	sequence[0] = 0
	sequence[1] = 1
	for i := 2; i <= depth; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}
	return sequence
}

func fillMatrix(matrix [][]int, sequence []int, rows int, cols int) {
	top, bottom, left, right := 0, rows-1, 0, cols-1
	index := 0

	for {
		if top > bottom || left > right {
			break
		}

		for i := left; i <= right; i++ {
			if index < len(sequence) {
				matrix[top][i] = sequence[index]
				index++
			}
		}
		top++

		for i := top; i <= bottom; i++ {
			if index < len(sequence) {
				matrix[i][right] = sequence[index]
				index++
			}
		}
		right--

		for i := right; i >= left; i-- {
			if index < len(sequence) {
				matrix[bottom][i] = sequence[index]
				index++
			}
		}
		bottom--

		for i := bottom; i >= top; i-- {
			if index < len(sequence) {
				matrix[i][left] = sequence[index]
				index++
			}
		}
		left++
	}
}
