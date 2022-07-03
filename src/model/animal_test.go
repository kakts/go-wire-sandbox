package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Owner.Animalをモック化したOwnerインスタンスを取得する
// Animalインタフェースを実装したモック用structを作り、DIさせる
func getMockedOwnerDog() Owner {
	mockedDog := MockDog{}
	owner := Owner{Animal: mockedDog}
	return owner
}

func TestOwnerDogBark(t *testing.T) {
	owner := getMockedOwnerDog()

	got := owner.Animal.Bark()
	require.Equal(t, "This is mock dog.", got)
}

func TsetOwnerDogWalk(t *testing.T) {
	owner := getMockedOwnerDog()

	for i := 0; i < 5; i++ {
		got := owner.Animal.Walk()
		require.Equal(t, i+1, got)
	}
}

type MockDog struct {
	X int
}

func (md MockDog) Bark() string {
	return "This is mock dog."
}

func (md MockDog) Walk() int {
	return 10
}
