package audio

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pilotpirxie/ebiten-test/src/game"
	"os"
	"path"
)

type Audio struct {
	image           *ebiten.Image
	audioContext    *audio.Context
	hurtAudioPlayer *audio.Player
	coinAudioPlayer *audio.Player
}

func NewAudio() game.Entity {
	return &Audio{}
}

func (b *Audio) Start(state *game.StateShape) error {
	b.audioContext = audio.NewContext(44100)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	hurtSfxFileData, err := ebitenutil.OpenFile(path.Join(pwd, "../src/assets/audio/hitHurt.wav"))
	if err != nil {
		return err
	}

	coinSfxFileData, err := ebitenutil.OpenFile(path.Join(pwd, "../src/assets/audio/pickupCoin.wav"))
	if err != nil {
		return err
	}

	hurtAudioPlayer, err := b.audioContext.NewPlayer(hurtSfxFileData)
	if err != nil {
		return err
	}

	b.hurtAudioPlayer = hurtAudioPlayer

	coinAudioPlayer, err := b.audioContext.NewPlayer(coinSfxFileData)
	if err != nil {
		return err
	}

	b.coinAudioPlayer = coinAudioPlayer

	return nil
}

func (b *Audio) Update(_ *game.StateShape) error {
	return nil
}

func (b *Audio) Draw(_ *game.StateShape, _ *ebiten.Image) error {
	return nil
}

func (b *Audio) PlayHurt() error {
	err := b.hurtAudioPlayer.Rewind()
	if err != nil {
		return err
	}

	b.hurtAudioPlayer.Play()
	return nil
}

func (b *Audio) PlayCoin() error {
	err := b.coinAudioPlayer.Rewind()
	if err != nil {
		return err
	}

	b.coinAudioPlayer.Play()
	return nil
}
