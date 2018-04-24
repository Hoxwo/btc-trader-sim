package exchange

type Exchange struct {
    name string
    valueAdded float64
    maxValueAdded float64
    launchDay
}

func New(name string, valueAdded float64, maxValueAdded float64, launchDay int) Exchange {
    e := Exchange {name, valueAdded, maxValueAdded, launchDay}
    return e
}

func (e *Exchange) SetName(name string) string {
    e.name = name
}

func (e Exchange) Name() string {
    return e.name
}

func (e Exchange) ValueAdded() float64 {
    return e.valueAdded
}

func (e *Exchange) SetValueAdded(valueAdded float64) {
    e.valueAdded = valueAdded
}

func (e Exchange) MaxValueAdded() float64 {
    return e.maxValueAdded
}

func (e *Exchange) SetMaxValueAdded(maxValueAdded float64) {
    e.maxValueAdded = maxValueAdded
}

func (e Exchange) LaunchDay() int {
    return e.launchDay
}

func (e *Exchange) SetLaunchDay(launchDay int) {
    e.launchDay = launchDay
}

func (e *Exchange) DailyValueAdjustment(totalMarketCap float64, trend int) {
	if (trend == 1) {

	} else if (trend == 2) {

	} else if (trend == 3) {

	} else {

	}
}
