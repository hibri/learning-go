package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyCheckout(t *testing.T) {

	assert.Equal(t, 0, scan(""))
}

func TestCheckoutA(t *testing.T) {

	assert.Equal(t, 50, scan("A"))
}

func TestCheckoutAA(t *testing.T) {

	assert.Equal(t, 100, scan("AA"))
}

func TestCheckoutAAA(t *testing.T) {

	assert.Equal(t, 130, scan("AAA"))
}

func TestCheckoutB(t *testing.T) {

	assert.Equal(t, 30, scan("B"))
}

func TestCheckoutBB(t *testing.T) {

	assert.Equal(t, 45, scan("BB"))
}
