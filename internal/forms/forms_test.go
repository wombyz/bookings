package forms

import (
	"net/url"
	"testing"
)

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Errorf("non existent form item passed check")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have and error but did not get one")
	}

	postedData.Add("pass", "long enough!")
	postedData.Add("fail", "no")
	form = New(postedData)

	form.MinLength("pass", 3)
	if !form.Valid() {
		t.Errorf("Error thrown when input is valid. Input was: %s", form.Get("pass"))
	}

	isError = form.Errors.Get("pass")
	if isError != "" {
		t.Error("should not have and error but got one")
	}

	form.MinLength("fail", 3)
	if form.Valid() {
		t.Errorf("Passed when input is too shortInput was: %s", form.Get("fail"))
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")

	if has {
		t.Error("form shows has field when it does not")
	}

	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")

	if !has {
		t.Error("says form does not have field when it should")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Errorf("non existent form item passed check")
	}

	postedData.Add("email-correct", "mcchicken@gmail.com")
	postedData.Add("email-incorrect", "mcchicken@gmail")

	form = New(postedData)

	form.IsEmail("email-correct")
	if !form.Valid() {
		t.Error("rejected a valid email")
	}

	form.IsEmail("email-incorrect")
	if form.Valid() {
		t.Error("accepted an invalid email")
	}
}
