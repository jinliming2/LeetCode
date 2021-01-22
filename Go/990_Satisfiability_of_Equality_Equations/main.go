package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Satisfiability of Equality Equations.
// Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Satisfiability of Equality Equations.
// https://leetcode.com/submissions/detail/446345430/
func equationsPossible(equations []string) bool {
	db := [26]byte{}
	for i := range db {
		db[i] = byte(i)
	}
	for _, equation := range equations {
		if equation[1] == '=' {
			a, b := equation[0]-'a', equation[3]-'a'
			for db[a] != a {
				a, db[a] = db[a], db[db[a]]
			}
			for db[b] != b {
				b, db[b] = db[b], db[db[b]]
			}
			db[a] = b
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			a, b := equation[0]-'a', equation[3]-'a'
			for db[a] != a {
				a, db[a] = db[a], db[db[a]]
			}
			for db[b] != b {
				b, db[b] = db[b], db[db[b]]
			}
			if a == b {
				return false
			}
		}
	}
	return true
}
