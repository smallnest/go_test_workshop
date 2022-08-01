package mock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := NewMockReader(ctrl)

	r.EXPECT().Read(gomock.All()).Do(func(arg0 []byte) {
		for i := 0; i < 10; i++ {
			arg0[i] = 'a' + byte(i)
		}
	}).Return(10, nil).AnyTimes()

	data := make([]byte, 1024)
	n, err := r.Read(data)
	assert.NoError(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, "abcdefghij", string(data[:n]))

}
