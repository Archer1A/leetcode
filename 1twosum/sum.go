package _twosum

func twoSum(nums []int, target int) []int {
	var result []int
	res := map[int]int{}
	for i,num := range nums {
		sub := target - num
		if r,ok := res[sub];ok{
			result = append(result, i)
			result = append(result, r)
			return result
		}
		res[num] = i
	}
	return result
}