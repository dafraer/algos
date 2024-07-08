package main

import (
	"fmt"
	"math"
)
func asteroidCollision(asteroids []int) []int {

	for {
		check := false
		for i := 1; i < len(asteroids); i++ {
			if opposite(asteroids[i-1], asteroids[i]) {
				asteroids = collide(asteroids, i-1, i)
				check = true
				break
			}
		}
		if !check {
			break
		}
	}
	return asteroids
}

func opposite(a, b int) bool {
	if a >= 0 && b < 0 {
		return true
	}
	return false
}

func collide(asteroids []int, i1, i2 int) []int {
	ans := make([]int, 0, len(asteroids))
	asteroid1 := int(math.Abs(float64(asteroids[i1])))
	asteroid2 := int(math.Abs(float64(asteroids[i2])))
	if asteroid1 == asteroid2 {
		//IMPORTANT might go out of boundaries
		ans = append(ans, asteroids[0:i1]...)
		ans = append(ans, asteroids[i2+1:]...)
	} else if asteroid1 > asteroid2 {
		ans = append(ans, asteroids[0:i1+1]...)
		ans = append(ans, asteroids[i2+1:]...)
	} else if asteroid1 < asteroid2 {
		ans = append(ans, asteroids[0:i1]...)
		ans = append(ans, asteroids[i2:]...)
	}
	return ans
}

