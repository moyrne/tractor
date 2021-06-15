package netx

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestQueryIP(t *testing.T) {
	ip, err := QueryIP("182.254.116.116:53", "www.baidu.com")
	assert.Equal(t, err, nil)
	ipv4, err := ip.Ipv4()
	assert.Equal(t, err, nil)
	assert.Equal(t, string(ip), ipv4.String())
}
