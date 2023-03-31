package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorLike(t *testing.T) {
	operator, _ := OperatorQuery("like")
	assert.Equal(t, "like", operator)
}
func TestOperatorGreaterThan(t *testing.T) {
	operator, _ := OperatorQuery("gt")
	assert.Equal(t, ">", operator)
}
func TestOperatorGreaterThanEqual(t *testing.T) {
	operator, _ := OperatorQuery("gte")
	assert.Equal(t, ">=", operator)
}
func TestOperatorLowerThan(t *testing.T) {
	operator, _ := OperatorQuery("lt")
	assert.Equal(t, "<", operator)
}

func TestOperatorLowerThanEqual(t *testing.T) {
	operator, _ := OperatorQuery("lte")
	assert.Equal(t, "<=", operator)
}

func TestOperatorEqual(t *testing.T) {
	operator, _ := OperatorQuery("eq")
	assert.Equal(t, "=", operator)
}

func TestOperatorError(t *testing.T) {
	_, err := OperatorQuery("sf")
	assert.Equal(t, "operator symbol paramater is not valid", err.Error())
}
