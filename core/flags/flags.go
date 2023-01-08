package flags

import "flag"

var recordFlow = flag.Bool("recordflow", false, "use unit tests to measure program process flow")

func GetRecordFlow() bool {
	return *recordFlow
}
