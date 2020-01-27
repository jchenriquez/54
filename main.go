package main

/**
*Given a matrix of m x n elements (m rows, n 
*columns), return all elements of the matrix in
*spiral order.
*
*Example 1:
*
*Input:
*[
* [ 1, 2, 3 ],
* [ 4, 5, 6 ],
* [ 7, 8, 9 ]
*]
*Output: [1,2,3,6,9,8,7,4,5]
**/

import ("math"
        "fmt")

const (
  RIGHT int =iota
  DOWN
  LEFT
  UP
)

func toKey (x, y int) string {
  return fmt.Sprintf("%d,%d", x, y)
}

func move (x, y, orientation int) (int, int) {
  switch orientation {
    case RIGHT:
      y++
    case LEFT:
      y--
    case DOWN:
      x++
    case UP:
      x--  
  }

  return x, y
}

func spiralOrderH(matrix[][]int, x int, y int, seenPositions map[string]bool, result *[]int, orientations [] int, currentOrientation int) {

  if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
    return 
  }
  seenPositions[toKey(x, y)] = true
  *result = append(*result, matrix[x][y])
  orientation := orientations[currentOrientation]

  xt, yt := move(x, y, orientation)
  _, seen := seenPositions[toKey(xt, yt)];
  if xt < 0 || xt >= len(matrix) || yt < 0 || yt >= len(matrix[0]) || seen{
    nextOrientation := orientations[int(math.Mod(float64(currentOrientation+1), float64(len(orientations))))]
    xt, yt = move(x, y, nextOrientation)
    if _, seen = seenPositions[toKey(xt, yt)]; !seen {
      spiralOrderH(matrix, xt, yt, seenPositions, result, orientations, nextOrientation)
    }
   } else {
     spiralOrderH(matrix, xt, yt, seenPositions, result, orientations, orientation)
   }
}

func spiralOrder(matrix [][]int) []int {
  seenPositions := make(map[string]bool)
  ordDirect := []int {RIGHT, DOWN, LEFT, UP}

  if len(matrix) == 0 {
    return []int{}
  }
  
  result := make([]int, 0,len(matrix)*len(matrix[0]))

  x := 0
  y := 0
  resultPtr := &result

  spiralOrderH(matrix, x, y, seenPositions, resultPtr, ordDirect, 0)
  return *resultPtr
}

func main() {
  fmt.Println(spiralOrder([][]int{[]int{ 1 }}))
}