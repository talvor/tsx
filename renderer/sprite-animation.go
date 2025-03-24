package renderer

import (
	"time"
)

type SimpleAnimation struct {
	sprite   SpriteDrawer
	frames   []int
	duration uint32
}

func NewSimpleAnimation(sprite SpriteDrawer, duration uint32, frames []int) *SimpleAnimation {
	return &SimpleAnimation{
		sprite:   sprite,
		frames:   frames,
		duration: duration,
	}
}

func (sa *SimpleAnimation) SetFrames(frames []int) {
	sa.frames = frames
}

func (sa *SimpleAnimation) DrawAnimation(opts *DrawOptions) error {
	animationIdx := int(time.Now().UnixMilli()) / int(sa.duration) % len(sa.frames)
	frame := sa.frames[animationIdx]
	if frame == -1 {
		return nil
	}
	return sa.sprite.Draw(uint32(frame), opts)
}
