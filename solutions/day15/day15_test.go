package day15

import "testing"

func TestProblem1(t *testing.T) {
	testP1(t, []int{1,3,2}, 1)
	testP1(t, []int{2,1,3}, 10)
	testP1(t, []int{1,2,3}, 27)
	testP1(t, []int{2,3,1}, 78)
	testP1(t, []int{0,3,6}, 436)
	testP1(t, []int{3,2,1}, 438)
	testP1(t, []int{3,1,2}, 1836)
}

func testP1(t *testing.T, input []int, expected int) {
	result := problem1(input)
	if result != expected {
		t.Errorf("Failure - received: %d, expected: %d.", result, expected)
	}
}

func TestProblem2(t *testing.T) {
	testP2(t, []int{0,3,6}, 175594)
	testP2(t, []int{1,3,2}, 2578)
	testP2(t, []int{2,1,3}, 3544142)
	testP2(t, []int{1,2,3}, 261214)
	testP2(t, []int{2,3,1}, 6895259)
	testP2(t, []int{3,2,1}, 18)
	testP2(t, []int{3,1,2}, 362)
}

func testP2(t *testing.T, input []int, expected int) {
	result := problem2(input)
	if result != expected {
		t.Errorf("Failure - received: %d, expected: %d.", result, expected)
	}
}

