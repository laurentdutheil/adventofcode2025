package day5

import (
	"sort"
)

type Kitchen struct {
	Intervals       []FreshnessInterval
	Count           int
	MergedIntervals []FreshnessInterval
}

func (k *Kitchen) AddInterval(interval FreshnessInterval) {
	k.Intervals = append(k.Intervals, interval)
}

func (k *Kitchen) CheckFreshness(ingredient *Ingredient) {
	for _, interval := range k.Intervals {
		if !ingredient.Fresh && ingredient.id >= interval.min && ingredient.id <= interval.max {
			k.Count += 1
			ingredient.Fresh = true
		}
	}
}

func (k *Kitchen) MergeIntervals() {
	sort.Slice(k.Intervals, func(i, j int) bool {
		if k.Intervals[i].min == k.Intervals[j].min {
			return k.Intervals[i].max < k.Intervals[j].max
		}
		return k.Intervals[i].min < k.Intervals[j].min
	})

	currentInterval := k.Intervals[0]
	for _, interval := range k.Intervals {
		if interval.min <= currentInterval.max {
			if interval.max >= currentInterval.max {
				currentInterval = NewFreshnessInterval(currentInterval.min, interval.max)
			}
		} else {
			k.MergedIntervals = append(k.MergedIntervals, currentInterval)
			currentInterval = NewFreshnessInterval(interval.min, interval.max)
		}
	}
	k.MergedIntervals = append(k.MergedIntervals, currentInterval)
}

func (k *Kitchen) CountIngredients() int {
	count := 0
	for _, interval := range k.MergedIntervals {
		count = count + interval.max - interval.min + 1
	}
	return count
}

func NewKitchen() *Kitchen {
	return &Kitchen{}
}

type ID = int

type Ingredient struct {
	id    ID
	Fresh bool
}

func NewIngredient(id ID) *Ingredient {
	return &Ingredient{id: id}
}

type FreshnessInterval struct {
	min ID
	max ID
}

func NewFreshnessInterval(min ID, max ID) FreshnessInterval {
	return FreshnessInterval{max: max, min: min}
}
