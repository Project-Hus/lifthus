package act

import (
	"routine/components/domain/aggregates/user"
	"testing"
)

func TestCreateActFailByTooLongName(t *testing.T) {
	author := user.UserFrom(42)
	_, err := CreateWeightAct("", *author, ActDescriptionFrom(nil, "", nil))
	_, err := CreateAct()
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
