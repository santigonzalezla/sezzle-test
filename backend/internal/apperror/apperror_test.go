package apperror

import "testing"

func TestAppError_Error(t *testing.T) {
	err := BadRequest("SOME_CODE", "something went wrong")

	if err.Error() != "something went wrong" {
		t.Errorf("Error() = %q, want %q", err.Error(), "something went wrong")
	}
}
