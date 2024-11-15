package mo

import (
	"reflect"
	"testing"
)

func TestPipelineBuilder_Match(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		got := NewPipeline().
			Match(Eq("foo", 2)).Value()

		NewPipeline().
			Match(Eq("foo", 2)),

		if !reflect.DeepEqual(got, []interface{}{}) {
			t.Errorf("Match() = %v, want %v", got, []interface{}{})
		}
	})
}
