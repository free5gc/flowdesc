package flowdesc_test

import (
	"github.com/stretchr/testify/assert"
	"free5gc/lib/flowdesc"
	"testing"
)

func TestIPFilterRuleEncode(t *testing.T) {
	testStr1 := "permit out ip from any to assigned 655"

	rule := flowdesc.NewIPFilterRule()
	if rule == nil {
		t.Fatal("IP Filter Rule Create Error")
	}

	if err := rule.SetAction(true); err != nil {
		assert.Nil(t, err)
	}

	if err := rule.SetDirection(true); err != nil {
		assert.Nil(t, err)
	}

	if err := rule.SetProtocal(0xfc); err != nil {
		assert.Nil(t, err)
	}

	if err := rule.SetSourceIp("any"); err != nil {
		assert.Nil(t, err)
	}

	if err := rule.SetDestinationIp("assigned"); err != nil {
		assert.Nil(t, err)
	}

	if err := rule.SetDestinationPorts("655"); err != nil {
		assert.Nil(t, err)
	}

	result, err := rule.Encode()
	if err != nil {
		assert.Nil(t, err)
	}
	if result != testStr1 {
		t.Fatalf("Encode error, \n\t expect: %s,\n\t    get: %s", testStr1, result)
	}
}
