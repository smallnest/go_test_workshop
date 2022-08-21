package s2

import (
	"net/http"
	"testing"
)

func TestParallel(t *testing.T) {
	var urls = map[string]string{"baidu": "http://baidu.com", "bing": "http://bing.com", "google": "http://google.com"}
	for k, v := range urls {
		v := v
		ok := t.Run(k, func(t *testing.T) {
			t.Parallel()

			t.Logf("start to get %s", v)
			resp, err := http.Get(v)
			t.Logf("end to get %s", v)
			if err != nil {
				t.Errorf("failed to get %s: %v", v, err)
				return
			}

			resp.Body.Close()
		})

		t.Logf("run: %t", ok)
	}
}
