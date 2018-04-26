package exchange

import "time"
import "math/rand"
import percent "github.com/dariubs/percent"

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
	//modifier to make exchanges less prone to crashing for a long time
	modifier := 10
	if(e.ValueAdded() < 10 && (marketTrend == 1 || marketTrend == 2)) {
		modifier = 12		
	} else if (e.ValueAdded() < 20 && (marketTrend == 1 || marketTrend == 2)) {
		modifier = 10
	} else if (e.ValueAdded() < 20 && (marketTrend == 3)) {
		modifier = 7
	} else if (e.ValueAdded() < 20 && (marketTrend == 4)) {
		modifier = 4
	}

	if (marketTrend == 1) {
	    dailyGains := random(5, 22)
	    if(dailyGains == 17 || dailyGains == 19) {
		dailyGains = 3 //nerfing the top end
	    }
	    if ((e.ValueAdded() + 1) + int(percent.Percent(dailyGains, int(e.ValueAdded()))) < e.MaxValueAdded()) {
	    	e.SetValueAdded(e.ValueAdded() + int(percent.Percent(dailyGains, int(e.ValueAdded()+modifier))))
	    }
	} else if (marketTrend == 2) {
	    dailyGains := random(1, 5)
		if(e.MaxValueAdded() > 200) {  //buffing larger exchanges
			dailyGains = dailyGains + 2
			modifier = modifier + 2
		}
	    if ((e.ValueAdded() + 1) + int(percent.Percent(dailyGains, int(e.ValueAdded()))) < e.MaxValueAdded()) {
	     	e.SetValueAdded(e.ValueAdded() + int(percent.Percent(dailyGains, int(e.ValueAdded()+modifier))))
	    }
	} else if (marketTrend == 3) {
	    dailyGains := random(1, 5)
	    if (e.ValueAdded() - int(percent.Percent(dailyGains, int(e.ValueAdded()))) <= 0) {
		e.SetValueAdded(0)
	    } else {
		e.SetValueAdded(e.ValueAdded() - int(percent.Percent(dailyGains, int(e.ValueAdded()+modifier))))
	    }
	} else {	
	    dailyGains := random(3, 16)
	    if (e.ValueAdded() - int(percent.Percent(dailyGains, int(e.ValueAdded()))) <= 0) {
		e.SetValueAdded(0)
	    } else {
		e.SetValueAdded(e.ValueAdded() - int(percent.Percent(dailyGains, int(e.ValueAdded()+modifier))))
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
		e.SetLaunchDay(e.LaunchDay() + dailySlide) 
	}
    } else {
	dailySlide := random(1, 3)
	e.SetLaunchDay(e.LaunchDay() + dailySlide) 
    }

    return e.LaunchDay()
}
