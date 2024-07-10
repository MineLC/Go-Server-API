package data

type BlockAction int32

const (
	StartDestroy BlockAction = iota
	AbortDestroy
	StopDestroy
	DropAllItems
	DropItem
	ReleaseItem
)
