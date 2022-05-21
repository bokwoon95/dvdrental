package dvdrental

import (
	"testing"

	"github.com/bokwoon95/dvdrental/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestCmp(t *testing.T) {
	gotFilms := setA()
	wantFilms := setB()
	if diff := testutil.Diff(gotFilms, wantFilms); diff != "" {
		t.Error(testutil.Callers(), diff)
	}
}

func TestTestify(t *testing.T) {
	gotFilms := setA()
	wantFilms := setB()
	assert.Equal(t, gotFilms, wantFilms, "should be equal")
}
