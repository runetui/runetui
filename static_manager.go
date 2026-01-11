package runetui

import "strings"

type StaticManager struct {
	staticBuffer []string
	staticKeys   map[string]int
}

func NewStaticManager() *StaticManager {
	return &StaticManager{
		staticBuffer: []string{},
		staticKeys:   make(map[string]int),
	}
}

func (sm *StaticManager) AppendStatic(key string, content []string) int {
	if _, exists := sm.staticKeys[key]; exists {
		return 0
	}
	sm.staticBuffer = append(sm.staticBuffer, content...)
	sm.staticKeys[key] = len(sm.staticBuffer)
	return len(content)
}

func (sm *StaticManager) RenderStatic() string {
	return strings.Join(sm.staticBuffer, "\n")
}

func (sm *StaticManager) Clear() {
	sm.staticBuffer = []string{}
	sm.staticKeys = make(map[string]int)
}
