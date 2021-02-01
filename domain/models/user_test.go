package models_test

import (
	validator "github.com/asaskevich/govalidator"
	model "github.com/gilsongabriel/imersao-fullstack-fullcycle/domain/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewBank(t *testing.T) {

	email := "gilson@codemastersolucoes.com"
	name := "Gilson Gabriel"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)
	require.True(t, validator.IsEmail(user.Email))
	require.True(t, validator.IsUUIDv4(user.ID))

	_, err = model.NewUser("", "")
	require.NotNil(t, err)
}