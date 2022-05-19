package b

import c "github.com/AppliedGoCourses/C"

func B() string {
	return "b.B uses package c at version " + c.Version()
}
