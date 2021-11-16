package empty1

import (
	"sort"
)

type Tempty1 struct{}
type Tempty1Int int
type Tempty1StringList []string

var Tempty1Var = Tempty1{}
var Tempty1IntVar = Tempty1Int(1)
var Tempty1StringListVar = Tempty1StringList{"a", "23"}

func (receiver Tempty1StringList) Sort() {
	sort.Strings(receiver)
}
