package validator

import (
	"testing"

	ii18n "github.com/adityarifqyfauzan/go-boilerplate/pkg/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestValidator(t *testing.T) {
	type User struct {
		Name  string `validate:"required"`
		Age   int    `validate:"required,gte=18,lte=100"`
		Email string `validate:"required,email"`
	}
	bundle := ii18n.Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	localizer := i18n.NewLocalizer(bundle, "en")
	validator := New(localizer)

	data := User{
		Name:  "",
		Age:   10,
		Email: "aditya",
	}

	errors := validator.Validate(data)
	if errors["Name"] != "Name is required" {
		t.Errorf("expected %s, got %s", "Name is required", errors["Name"])
	}

	if errors["Age"] != "Age must be greater than or equal to 18" {
		t.Errorf("expected %s, got %s", "Age must be greater than or equal to 18", errors["Age"])
	}

	if errors["Email"] != "Email must be a valid email address" {
		t.Errorf("expected %s, got %s", "Email must be a valid email address", errors["Email"])
	}
}

func TestValidatorInID(t *testing.T) {
	type User struct {
		Name  string `validate:"required"`
		Age   int    `validate:"required,gte=18,lte=100"`
		Email string `validate:"required,email"`
	}
	bundle := ii18n.Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	localizer := i18n.NewLocalizer(bundle, "id")
	validator := New(localizer)

	data := User{
		Name:  "",
		Age:   10,
		Email: "aditya",
	}

	errors := validator.Validate(data)
	if errors["Name"] != "Nama wajib diisi" {
		t.Errorf("expected %s, got %s", "Nama wajib diisi", errors["Name"])
	}

	if errors["Age"] != "Umur harus lebih besar atau sama dengan 18" {
		t.Errorf("expected %s, got %s", "Umur harus lebih besar atau sama dengan 18", errors["Age"])
	}

	if errors["Email"] != "Email harus berupa alamat email yang valid" {
		t.Errorf("expected %s, got %s", "Email harus berupa alamat email yang valid", errors["Email"])
	}
}

func TestValidatorInJA(t *testing.T) {
	type User struct {
		Name  string `validate:"required"`
		Age   int    `validate:"required,gte=18,lte=100"`
		Email string `validate:"required,email"`
	}
	bundle := ii18n.Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	localizer := i18n.NewLocalizer(bundle, "ja")
	validator := New(localizer)

	data := User{
		Name:  "",
		Age:   10,
		Email: "aditya",
	}

	errors := validator.Validate(data)
	if errors["Name"] != "名前は必須項目です" {
		t.Errorf("expected %s, got %s", "名前は必須項目です", errors["Name"])
	}

	if errors["Age"] != "年齢は18以上でなければなりません" {
		t.Errorf("expected %s, got %s", "年齢は18以上でなければなりません", errors["Age"])
	}

	if errors["Email"] != "メールアドレスは有効なメールアドレスである必要があります" {
		t.Errorf("expected %s, got %s", "メールアドレスは有効なメールアドレスである必要があります", errors["Email"])
	}
}

func TestValidatorPasswordConfirmation(t *testing.T) {
	type User struct {
		Password             string `json:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	}
	bundle := ii18n.Init("../../locales")
	if bundle == nil {
		panic("failed to init i18n")
	}

	localizer := i18n.NewLocalizer(bundle, "id")
	validator := New(localizer)

	data := User{
		Password:             "password",
		PasswordConfirmation: "password_confirmation",
	}

	errors := validator.Validate(data)
	if errors["PasswordConfirmation"] != "PasswordConfirmation tidak cocok" {
		t.Errorf("expected %s, got %s", "PasswordConfirmation tidak cocok", errors["PasswordConfirmation"])
	}
}
