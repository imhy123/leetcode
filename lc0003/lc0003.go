package lc0003

var Target = lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	var length = len(s)

	if length <= 0 {
		return 0
	}

	var i, j, ret int = 0, 0, 0
	var bitmap [128]int

	bitmap[s[j]] = 1

	for j < length {
		j++
		if j == length {
			break
		}

		if bitmap[s[j]] == 1 {
			if j-i > ret {
				ret = j - i
			}

			var matched = false
			for i < j {
				if matched {
					break
				}

				matched = s[i] == s[j]

				bitmap[s[i]] = 0
				i++
			}
		}

		bitmap[s[j]] = 1
	}

	if j-i > ret {
		ret = j - i
	}

	return ret
}
