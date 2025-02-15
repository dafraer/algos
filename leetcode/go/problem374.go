/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	r := n + 1
	l := 0
	for r-l > 1 {
		switch guess(l + (r-l)/2) {
		case -1:
			r = l + (r-l)/2 - 1
		case 1:
			l = l + (r-l)/2
		case 0:
			return l + (r-l)/2
		}
	}
	return r
}
