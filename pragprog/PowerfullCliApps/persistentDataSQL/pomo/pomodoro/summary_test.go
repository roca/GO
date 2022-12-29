package pomodoro_test

import (
	"testing"
	"time"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/pomodoro"
)

func TestDailySummary(t *testing.T) {
	repo, cleanup := getRepo(t)
	defer cleanup()

	days := []time.Time{
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC),
	}

	intervals := []pomodoro.Interval{
		{ID: 1, StartTime: days[0], PlannedDuration: 25 * time.Minute, ActualDuration: 25 * time.Minute, Category: pomodoro.CategoryPomodoro, State: pomodoro.StateDone},
		{ID: 2, StartTime: days[1], PlannedDuration: 5 * time.Minute, ActualDuration: 5 * time.Minute, Category: pomodoro.CategoryShortBreak, State: pomodoro.StateDone},
		{ID: 3, StartTime: days[2], PlannedDuration: 15 * time.Minute, ActualDuration: 15 * time.Minute, Category: pomodoro.CategoryLongBreak, State: pomodoro.StateDone},
		{ID: 4, StartTime: days[3], PlannedDuration: 25 * time.Minute, ActualDuration: 19 * time.Minute, Category: pomodoro.CategoryPomodoro, State: pomodoro.StateCancelled},
		{ID: 5, StartTime: days[0], PlannedDuration: 25 * time.Minute, ActualDuration: 25 * time.Minute, Category: pomodoro.CategoryPomodoro, State: pomodoro.StateDone},
	}

	for i := range intervals {
		_, err := repo.Create(intervals[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	lastInterval, err := repo.Last()
	if err != nil {
		t.Fatal(err)
	}

	if lastInterval.ID != 5 {
		t.Fatalf("Last interval ID should be 5, got %d", lastInterval.ID)
	}

	if lastInterval.State != pomodoro.StateDone {
		t.Fatalf("Last interval state should be %q, got %q", pomodoro.StateDone, lastInterval.State)
	}

	categorySummary, err := repo.CategorySummary(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), pomodoro.CategoryPomodoro)
	if err != nil {
		t.Fatal(err)
	}
	if categorySummary != 50*time.Minute {
		t.Fatalf("Category summary should be %s, got %s", 50*time.Minute, categorySummary)
	}

	config := pomodoro.NewConfig(repo, 25*time.Minute, 5*time.Minute, 15*time.Minute)



	summary, err := pomodoro.DailySummary(days[0], config)
	if err != nil {
		t.Fatal(err)
	}
	dPomo := summary[0]
	if dPomo != intervals[0].ActualDuration+intervals[4].ActualDuration {
		t.Fatalf("Daily summary should be %s, got %s", 25*time.Minute, dPomo)
	}

	summary, err = pomodoro.DailySummary(days[1], config)
	if err != nil {
		t.Fatal(err)
	}
	dPomo = summary[0]
	if dPomo != 0 {
		t.Fatalf("Daily summary should be %s, got %s", 0*time.Minute, dPomo)
	}

	dBreaks := summary[1]
	if dBreaks != intervals[1].ActualDuration {
		t.Fatalf("Daily breaks summary should be %s, got %s", 5*time.Minute, dBreaks)
	}


}
