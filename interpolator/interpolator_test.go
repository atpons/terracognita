package interpolator_test

import (
	"testing"

	"github.com/cycloidio/terracognita/interpolator"
	"github.com/stretchr/testify/assert"
)

func TestInterpolate(t *testing.T) {

	t.Run("normal cases", func(t *testing.T) {
		i := interpolator.New("azurerm")
		i.AddResourceAttributes("azurerm_virtual_machine.front", map[string]string{
			"id":    "secretid",
			"other": "ovalue",
		})
		i.AddResourceAttributes("azurerm_virtual_something.front", map[string]string{
			"id": "secretid",
		})

		s, ok := i.Interpolate("virtual_machine_id", "secretid")
		assert.Equal(t, s, "${azurerm_virtual_machine.front.id}")
		assert.True(t, ok)

		s, ok = i.Interpolate("virtual_machine_potato", "ovalue")
		assert.Equal(t, s, "${azurerm_virtual_machine.front.other}")
		assert.True(t, ok)

		s, ok = i.Interpolate("totally_random", "secretid")
		assert.Equal(t, s, "${azurerm_virtual_something.front.id}")
		assert.True(t, ok)

		s, ok = i.Interpolate("virtual_id", "secretid")
		assert.Equal(t, s, "${azurerm_virtual_machine.front.id}")
		assert.True(t, ok)

		s, ok = i.Interpolate("virtual_machine_potato", "none")
		assert.Equal(t, s, "")
		assert.False(t, ok)

	})

	t.Run("resource id is uses hyphen, but but resource is underscore", func(t *testing.T) {
		aws := interpolator.New("aws")
		aws.AddResourceAttributes("aws_security_group.sg_test", map[string]string{
			"id":          "sg-test",
			"description": "security groups",
		})
		s, ok := aws.Interpolate("security_group_id", "sg_test")
		assert.Equal(t, s, "${aws_security_group.sg_test.id}")
		assert.True(t, ok)
	})
}
