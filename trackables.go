package gomotion

type Hand struct {
	Id                     int
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

type InteractionBox struct {
	Center []float32 `json:center`
	Size   []float32 `json:size`
}

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

type Finger struct {
	Id int
	*Pointable
}

type Tool struct {
	Id int
	*Pointable
}

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
