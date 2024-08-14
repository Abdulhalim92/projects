package quiz

import "sync"

type ScoreManager struct {
	l      sync.RWMutex
	scores map[string]int
}

func NewScoreManager() *ScoreManager {
	return &ScoreManager{
		scores: map[string]int{},
	}
}

func (sm *ScoreManager) Update(login string, val int) {
	sm.l.Lock()
	defer sm.l.Unlock()
	sm.scores[login] = val
}

func (sm *ScoreManager) Read(login string) (int, bool) {
	sm.l.RLock()
	defer sm.l.RUnlock()
	val, ok := sm.scores[login]
	return val, ok
}
