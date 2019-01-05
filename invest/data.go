package invest

import (
	"math"
	"strconv"
)

type InvestVars struct {
	Init     string
	Invested string
}

func Calculate(init string) string {
	principal, _ := strconv.ParseFloat(init, 64)
	frequency := 12.0
	years := 30.0
	rate := 0.07
	tRate := 1 + rate/frequency
	time := frequency * years
	total := principal * math.Pow(tRate, time)
	return strconv.Itoa(int(total))
}
