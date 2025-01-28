package dummylogic

func DummyLogic(n int) string {
	if n < 10 {
		return "less than 10"
	} else if n < 20 {
		return "less than 20"
	} else {
		return "more than 20"
	}
}
