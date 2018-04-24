package exchange

import "time"
import "math/rand"

type Exchange struct {
    name string
    valueAdded int
    maxValueAdded int
    launchDay int
}

func New(name string, valueAdded int, maxValueAdded int, launchDay int) Exchange {
    e := Exchange {name, valueAdded, maxValueAdded, launchDay}
    return e
}

func (e *Exchange) SetName(name string) {
    e.name = name
}

func (e Exchange) Name() string {
    return e.name
}

func (e Exchange) ValueAdded() int {
    return e.valueAdded
}

func (e *Exchange) SetValueAdded(valueAdded int) {
    e.valueAdded = valueAdded
}

func (e Exchange) MaxValueAdded() int {
    return e.maxValueAdded
}

func (e *Exchange) SetMaxValueAdded(maxValueAdded int) {
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

func (e *Exchange) DailyValueAdjustment(totalMarketCap int, marketTrend int) int {
	if (marketTrend == 1) {
	    dailyIncrease := 10
	    e.SetValueAdded(e.ValueAdded() + dailyIncrease)
	} else if (marketTrend == 2) {
	    dailyIncrease := 10
	    e.SetValueAdded(e.ValueAdded() + dailyIncrease)
	} else if (marketTrend == 3) {
	    dailyDecrease := 1
	    if (dailyDecrease > e.ValueAdded()) {
		e.SetValueAdded(0)
	    } else {
		e.SetValueAdded(e.ValueAdded() - dailyDecrease)
	    }
	} else {	
	    dailyDecrease := 1
	    if (dailyDecrease > e.ValueAdded()) {
		e.SetValueAdded(0)
	    } else {
		e.SetValueAdded(e.ValueAdded() - dailyDecrease)
	    }
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
