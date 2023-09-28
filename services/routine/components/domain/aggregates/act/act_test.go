package act

import (
	"routine/components/domain/aggregates/user"
	"testing"
)

func TestCreateActFailByInvalidName(t *testing.T) {
	author := user.UserFrom(42)
	desc := getValidActDescription()
	_, err := CreateWeightAct(getTooLongActName(), *author, desc)
	if err != ErrInvalidActInfo {
		t.Errorf("too long act name is expected to cause ErrInvalidActInfo, but got %v", err)
	}
}

func TestCreateActFailByImages(t *testing.T) {

}

func TestCreateActFailByText(t *testing.T) {

}

func TestUpdateActFailByUnauthorized(t *testing.T) {

}

func TestUpdateActFailByImages(t *testing.T) {

}

func TestUpdateActFailByText(t *testing.T) {

}

func TestDeleteActFailByUnauthorized(t *testing.T) {

}
