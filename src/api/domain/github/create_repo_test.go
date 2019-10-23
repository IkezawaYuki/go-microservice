package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoResqestAsJson(t *testing.T){
	request := CreateRepoRequest{
		Name: "golang introduction",
		Description: "a golang introduction repository",
		Homepage: "https://github.com",
		Private:true,
		HasIssue:false,
		HasProjects:true,
		HasWiki:false,
	}

	if request.Private{

	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	var target CreateRepoRequest

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
}