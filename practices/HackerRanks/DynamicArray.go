package main

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	arr := make([][]int32,n)
	answerArr := make([]int32, 0)
	lastAnswer := int32(0)
	for _, q  := range queries {
		qtype := q[0]
		x := q[1]
		y := q[2]

		switch qtype{
			case 1:
				indx := (lastAnswer ^ x) % n
				arr[indx] = append(arr[indx], y)
			case 2:
				xor := lastAnswer ^ x
				indx := xor % n
				lastAnswer = arr[indx][y % int32(len(arr[indx]))]
				answerArr = append(answerArr, lastAnswer)
		}
	}
	return  answerArr
}

func main() {
	queries := [][]int32{ {1,0,5}, {1,1,7}, {1,0,3}, {2,1,0}, {2,1,1}}

	print(dynamicArray(2, queries))

}


