package api_test

import (
	. "api"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ShowPointHandler_Input_User_Id_Should_Be_Point(t *testing.T) {
	expected := `{"point":"100"}`
	url := "/inquiry?user_id=1"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()

	ShowPointHandler(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Expected %s but it got %s", expected, string(actual))
	}
}
