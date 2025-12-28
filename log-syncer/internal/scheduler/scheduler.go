// Package scheduler 定时调度器
// 每小时的 01 分和 31 分触发任务
package scheduler

import (
	"log"
	"time"
)

// Task 定时任务函数
type Task func()

// Scheduler 调度器
type Scheduler struct {
	task Task
	quit chan struct{}
}

// NewScheduler 创建调度器
func NewScheduler(task Task) *Scheduler {
	return &Scheduler{
		task: task,
		quit: make(chan struct{}),
	}
}

// Start 启动调度器
func (s *Scheduler) Start() {
	// 启动定时调度
	go s.run()
}

// Stop 停止调度器
func (s *Scheduler) Stop() {
	close(s.quit)
}

// run 运行调度循环
func (s *Scheduler) run() {
	for {
		nextTick := s.nextTick()
		log.Printf("调度器: 等待下一次执行，时间: %s", nextTick.Format("2006-01-02 15:04:05"))

		select {
		case <-time.After(time.Until(nextTick)):
			log.Printf("调度器: 触发任务")
			s.task()
		case <-s.quit:
			log.Printf("调度器: 已停止")
			return
		}
	}
}

// nextTick 计算下一次执行时间
// 规则: 每小时的 01 分和 31 分
func (s *Scheduler) nextTick() time.Time {
	now := time.Now()

	// 计算当前小时的 01 分和 31 分
	hour := now.Hour()
	minute := now.Minute()

	var nextTime time.Time

	// 本小时的 01 分
	t01 := time.Date(now.Year(), now.Month(), now.Day(), hour, 1, 0, 0, now.Location())
	// 本小时的 31 分
	t31 := time.Date(now.Year(), now.Month(), now.Day(), hour, 31, 0, 0, now.Location())
	// 下一小时的 01 分
	tNext01 := time.Date(now.Year(), now.Month(), now.Day(), hour+1, 1, 0, 0, now.Location())

	// 根据当前时间选择下一个执行点
	switch {
	case minute < 1:
		nextTime = t01
	case minute < 31:
		nextTime = t31
	default:
		nextTime = tNext01
	}

	return nextTime
}
