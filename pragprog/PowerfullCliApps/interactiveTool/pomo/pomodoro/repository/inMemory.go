package repository

import (
	"sync"

	"github.com/roca/GO/tree/staging/pragprog/PowerfullCliApps/interactiveTool/pomo/pomodoro"
)

type inMemoryRepo struct {
	sync.RWMutex
	intervals []pomodoro.Interval
}

func NewInMemoryRepo() *inMemoryRepo {
	return &inMemoryRepo{
		intervals: []pomodoro.Interval{},
	}
}

func (r *inMemoryRepo) Create(i pomodoro.Interval) (int64, error) 

func (r *inMemoryRepo) Update(i pomodoro.Interval) error          

func (r *inMemoryRepo) ByID(id int64) (pomodoro.Interval, error)  

func (r *inMemoryRepo) Last() (pomodoro.Interval, error)          

func (r *inMemoryRepo) Breaks(n int) ([]pomodoro.Interval, error) 

