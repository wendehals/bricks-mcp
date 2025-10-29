package tools

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/wendehals/bricks-cli/api"
	"github.com/wendehals/bricks-mcp/utils"
)

// fakeRoundTripper is an http.RoundTripper that returns canned responses based on the request URL
type fakeRoundTripper struct{}

func (f *fakeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	url := req.URL.String()
	var body string
	if url == "https://rebrickable.com/api/v3/users/_token/" {
		// token response
		body = `{"user_token":"TESTTOKEN"}`
	} else if bytes.Contains([]byte(url), []byte("/allparts")) {
		// parts page response
		body = `{"next":"","results":[{"quantity":2,"part":{"part_num":"3001","name":"Brick 2 x 4","part_url":"https://rebrickable.com/parts/3001","part_img_url":""},"color":{"id":1,"name":"Blue"},"is_spare":false,"set_num":""}]}`
	} else {
		// default: empty page
		body = `{"next":"","results":[]}`
	}

	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(body))),
		Header:     make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp, nil
}

func TestGetAllUserPartsToolWithFakeUsersAPI(t *testing.T) {
	// Install override to create UsersAPI with a fake HTTP client
	prev := utils.GetUsersAPIOverride
	defer func() { utils.GetUsersAPIOverride = prev }()

	utils.GetUsersAPIOverride = func() (*api.UsersAPI, error) {
		client := &http.Client{Transport: &fakeRoundTripper{}}
		creds := &api.Credentials{UserName: "u", Password: "p", APIKey: "k"}
		return api.NewUsersAPI(client, creds, false), nil
	}

	// Call the tool
	_, coll, err := GetAllUserParts(context.Background(), nil, GetAllUserPartsInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if coll.Parts == nil || len(coll.Parts) != 1 {
		t.Fatalf("expected 1 part, got %d", len(coll.Parts))
	}

	p := coll.Parts[0]
	if p.Quantity != 2 {
		t.Fatalf("expected quantity 2, got %d", p.Quantity)
	}
	if p.Shape.Number != "3001" {
		t.Fatalf("expected part_num 3001, got %s", p.Shape.Number)
	}
	if p.Color.Name != "Blue" {
		t.Fatalf("expected color Blue, got %s", p.Color.Name)
	}

	_ = coll
}
