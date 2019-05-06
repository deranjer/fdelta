package fdelta

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDelta(t *testing.T) {
	for i := 1; i <= 1; i++ {
		origin, err := loadData(i, "origin")
		if err != nil {
			t.Error(err)
		}

		target, err := loadData(i, "target")
		if err != nil {
			t.Error(err)
		}

		goodDelta, err := loadData(i, "delta")
		if err != nil {
			t.Error(err)
		}

		delta := Create(origin, target)

		_, err = Apply(origin, delta)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(delta, goodDelta) {
			t.Error("deltas are not equal")
		}
	}
}

func BenchmarkCreateDelta(b *testing.B) {
	for i := 1; i <= 5; i++ {
		origin, err := loadData(i, "origin")
		if err != nil {
			b.Error(err)
		}

		target, err := loadData(i, "target")
		if err != nil {
			b.Error(err)
		}

		if err != nil {
			b.Error(err)
		}

		b.Run(fmt.Sprintf("CreateDelta%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Create(origin, target)
			}
		})
	}
}

func BenchmarkApplyDelta(b *testing.B) {
	for i := 1; i <= 5; i++ {
		origin, err := loadData(i, "origin")
		if err != nil {
			b.Error(err)
		}

		target, err := loadData(i, "target")
		if err != nil {
			b.Error(err)
		}

		if err != nil {
			b.Error(err)
		}

		delta := Create(origin, target)

		b.Run(fmt.Sprintf("ApplyDelta%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Apply(origin, delta)
			}
		})

	}
}

func loadData(i int, name string) ([]byte, error) {
	return ioutil.ReadFile(
		filepath.FromSlash(fmt.Sprintf("%s/%d/%s", "data", i, name)),
	)
}
