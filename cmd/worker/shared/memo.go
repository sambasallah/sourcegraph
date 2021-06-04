package shared

import "sync"

// TODO - document
type MemoizedConstructor struct {
	ctor  func() (interface{}, error)
	value interface{}
	err   error
	once  sync.Once
}

// TODO - document
func NewMemo(ctor func() (interface{}, error)) *MemoizedConstructor {
	return &MemoizedConstructor{ctor: ctor}
}

// TODO - document
func (m *MemoizedConstructor) Init() (interface{}, error) {
	m.once.Do(func() { m.value, m.err = m.ctor() })
	return m.value, m.err
}
