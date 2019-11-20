package stats

var (
	InsertFindLoopCounter    int
	InsertBalanceLoopCounter int
	InsertRotateCounter      int

	FindLoopCounter int

	DeleteBalanceLoopCounter int
	DeleteRotateCounter      int
)

func Reset() {
	InsertFindLoopCounter = 0
	InsertBalanceLoopCounter = 0
	InsertRotateCounter = 0

	FindLoopCounter = 0

	DeleteBalanceLoopCounter = 0
	DeleteRotateCounter = 0
}
