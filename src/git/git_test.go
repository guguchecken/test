package git

import "testing"

func Test_OnlyPath(t *testing.T) {
	pr := NewPR(`guguducken`, `git_test`, `28`, `ahgjskaiogjeijskgjwkgj`)
	pr.onlyPaths()
}
