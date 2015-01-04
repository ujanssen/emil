package emil

type positionKey string

// PositionEntry is an entry in the PositionDb
type PositionEntry struct {
	Position      *position
	Dtm           int
	PrevPositions map[positionKey]*PositionEntry
	NextPositions map[positionKey]*PositionEntry
}

// NewPositionEntry ceates a new *PositionEntry
func NewPositionEntry(p *position) *PositionEntry {
	return &PositionEntry{
		Position:      position,
		Dtm:           initial,
		PrevPositions: make(map[positionKey]*PositionEntry),
		NextPositions: make(map[positionKey]*PositionEntry)}
}

// PositionDb to query for mate in 1,2, etc.
type PositionDb struct {
	positions map[positionKey]*PositionEntry
}

// NewPositionDB creates a new *PositionDB
func NewPositionDB() *PositionDb {
	return &PositionDb{
		positions: make(map[positionKey]*PositionEntry)}
}
