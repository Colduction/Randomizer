package randomizer_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/colduction/randomizer"
)

func benchPerCoreConfigs(b *testing.B, f func(b *testing.B)) {
	b.Helper()
	coreConfigs := []int{1, 2, 4, 8, 16, 32}
	for _, n := range coreConfigs {
		name := fmt.Sprintf("%d cores", n)
		b.Run(name, func(b *testing.B) {
			runtime.GOMAXPROCS(n)
			f(b)
		})
	}
}

func BenchmarkHashPool(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				_ = randomizer.DefaultHashPool.Sum64()
			}
		})
	})
}
