package exchange

import "time"
import "math/rand"

type Exchange struct {
    name string
    valueAdded float64
    maxValueAdded float64
    launchDay int
}

func New(name string, valueAdded float64, maxValueAdded float64, launchDay int) Exchange {
    e := Exchange {name, valueAdded, maxValueAdded, launchDay}
    return e
}

func (e *Exchange) SetName(name string) {
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

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func (e *Exchange) DailyValueAdjustment(totalMarketCap float64, marketTrend int) float64 {
	if (marketTrend == 1) {

	} else if (marketTrend == 2) {

	} else if (marketTrend == 3) {

	} else {

	}
	return e.ValueAdded()
}

func (e *Exchange) DailyLaunchAdjustment(marketTrend int) int {
    //depending on overall market trend, this exchange's launch will slide forward or backward
    if (marketTrend == 1) {
	dailySlide := random(1, 3)
	e.SetLaunchDay(e.LaunchDay() - dailySlide) 
    } else if (marketTrend == 2) {
	dailySlide := random(1, 3)
	chance := random(1,5)
	if (chance == 4) {
		e.SetLaunchDay(e.LaunchDay() - dailySlide) 
	}
    } else if (marketTrend == 3) {
	dailySlide := random(1, 3)
	chance := random(1,5)
	if (chance == 4) {
		e.SetLaunchDay(e.LaunchDay() - dailySlide) 
	}
    } else {
	dailySlide := random(1, 3)
	e.SetLaunchDay(e.LaunchDay() + dailySlide) 
    }

    return e.LaunchDay()
}
