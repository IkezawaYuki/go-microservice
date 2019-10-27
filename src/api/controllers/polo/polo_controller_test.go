package polo

import (
	"github.com/stretchr/testify/assert"
	"go-microservice/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T){

}

func TestPolo(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	c := test_utils.GetMockedContext(request, response)

	Marco(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}