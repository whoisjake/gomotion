package gomotion

// The hand is a basic trackable object through the LeapMotionDevice.
// There can be zero, one, or two that come through each Frame.
// A Hand will carry the same ID across the frames in which it is visible.
type Hand struct {
	Id                     int         `json:id`
	Direction              []float32   `json:direction`
	PalmNormal             []float32   `json:palmNormal`
	PalmPosition           []float32   `json:palmPosition`
	PalmVelocity           []float32   `json:palmVelocity`
	SphereCenter           []float32   `json:sphereCenter`
	SphereRadius           float32     `json:"sphereRadius"`
	StabilizedPalmPosition []float32   `json:StabilizedPalmPosition`
	R                      [][]float32 `json:"r"`
	S                      float32     `json:"s"`
	T                      []float32   `json:"t"`
	TimeVisible            float32     `json:timeVisible`
}

// A Frame can have a list of Gestures that are starting, updating, or ending.
// A Gesture will carry the same ID across the frames in which it is visible.
// *NOTE* Pay attention to the State of the Gesture, as it can give you an idea
// of where the gesture is going.
type Gesture struct {
	Id            int       `json:id`
	State         string    `json:state`
	GestureType   string    `json:type`
	Duration      int       `json:duration`
	HandIds       []int     `json:handIds`
	PointableIds  []int     `json:pointableIds`
	Speed         float32   `json:speed`
	Radius        float32   `json:radius`
	Progress      float32   `json:progress`
	Center        []float32 `json:center`
	Normal        []float32 `json:normal`
	StartPosition []float32 `json:startPosition`
	Position      []float32 `json:position`
	Direction     []float32 `json:direction`
}

// An InteractionBox gives you the physical bounding box in which
// the device is detecting the Hands, Gestures, etc. It's provided
// to create a perspective to map your screen view box to.
type InteractionBox struct {
	Center []float32 `json:center`
	Size   []float32 `json:size`
}

// A Finger or Tool can be a Pointable, as well as pointables themselves.
// A Frame will have a list of pointables associated with each hand.
type Pointable struct {
	Id                    int       `json:id`
	HandId                int       `json:handId`
	Direction             []float32 `json:direction`
	Length                float32   `json:length`
	StabilizedTipPosition []float32 `json:stabilizedTipPosition`
	TimeVisible           float32   `json:timeVisible`
	TipPosition           []float32 `json:tipPosition`
	TipVelocity           []float32 `json:tipVelocity`
	Tool                  bool      `json:tool`
	TouchDistance         float32   `json:touchDistance`
	TouchZone             string    `json:touchZone`
}

// A Frame can have a list of Fingers, which are structured like Pointables
type Finger struct {
	Id int `json:id`
	*Pointable
}

// A Frame can have a list of Tools, which are structured like Pointables
type Tool struct {
	Id int `json:id`
	*Pointable
}

// This struct represents each Frame presented on the LeapMotionDevice WebSocket.
// The base structure will fill in with information available or for the arrays,
// empty arrays, so that you can quickly iterate over each array in the struct.
type Frame struct {
	Id             int            `json:"id"`
	Timestamp      int            `json:"timestamp"`
	InteractionBox InteractionBox `json:interactionBox`
	Hands          []Hand         `json:"hands"`
	Pointables     []Pointable    `json:"pointables"`
	Fingers        []Finger       `json:"fingers"`
	Tools          []Tool         `json:"tools"`
	Gestures       []Gesture      `json:"gestures"`
	R              [][]float32    `json:"r"`
	S              float32        `json:"s"`
	T              []float32      `json:"t"`
}
