package packet

import "github.com/thecubenoob/aimjelnetlib/protocol/types"

type SetDefaultSpawnPosition struct {
	Location types.Position
	Angle    float32
}

func (s SetDefaultSpawnPosition) ID() int32 {
	return 0x50
}

func (s SetDefaultSpawnPosition) Decode(r *Reader) error {
	_ = r.Int64((*int64)(&s.Location))
	return r.Float32(&s.Angle)
}

func (s SetDefaultSpawnPosition) Encode(w Writer) error {
	_ = w.Int64(int64(s.Location))
	return w.Float32(s.Angle)
}
