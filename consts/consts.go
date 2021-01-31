package consts

// MoveDirection represents a movement direction
type MoveDirection int

// Movement Directions
const (
	MoveDirectionNone = iota
	MoveDirectionRight
	MoveDirectionUp
	MoveDirectionLeft
	MoveDirectionDown
)

func (d MoveDirection) String() string {
	return [...]string{"None", "Right", "Up", "Left", "Down"}[d]
}

// Main game Image sizes
const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

// GameState represents the state of the overall game (Intro, Ongoing, Paused, Won...)
type GameState int

// Game States
const (
	IntroState = iota
	MainGameState
	PausedGameState
)

func (state GameState) String() string {
	return [...]string{"Intro", "Main", "Paused"}[state]
}

// Runner animation frame sizes
const (
	FrameOX     = 0
	FrameOY     = 32
	FrameWidth  = 32
	FrameHeight = 32
	FrameNum    = 8
)

// Empty padding pixels on the main game image (To keep Entities visible inside the frame)
const (
	TopPadding    = 0
	RightPadding  = 30
	BottomPadding = 30
	LeftPadding   = 0
)

//MoveSpeed is the default movement speed of Moveables per frame
const MoveSpeed = 6

//UpdateNPCIntentInterval is the interval (amount of ticks) every which the moveState of NPCs is updated
const UpdateNPCIntentInterval = 60

//MinimumCollisionDistance is the minimum distance between two Movables before they can be considered to be colliding
const MinimumCollisionDistance = 30.0
