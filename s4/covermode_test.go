package s4

import "testing"

func TestExample(t *testing.T) {
	for in, out := range map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "carrot",
		"d": "date",
		"e": "eggplant",
		"f": "fig",
		"":  "nevermind",
	} {
		t.Run("input-"+in, func(t *testing.T) {
			// t.Parallel()
			result := example(in)
			if result != out {
				t.Errorf("sent %s and got %s expected %s", in, result, out)
			}
		})
	}
}
