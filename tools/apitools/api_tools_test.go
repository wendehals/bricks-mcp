package apitools

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/wendehals/bricks-cli/api"
	"github.com/wendehals/bricks-mcp/utils"
)

// fakeRoundTripperV2 is an http.RoundTripper that returns canned responses for various UsersAPI endpoints
type fakeRoundTripperV2 struct{}

func (f *fakeRoundTripperV2) RoundTrip(req *http.Request) (*http.Response, error) {
	url := req.URL.String()
	var body string

	switch {
	case url == "https://rebrickable.com/api/v3/users/_token/":
		body = `{"user_token":"TESTTOKEN"}`

	// Parts endpoint
	case bytes.Contains([]byte(url), []byte("/allparts")):
		body = `{"next":"","results":[{"quantity":2,"part":{"part_num":"3001","name":"Brick 2 x 4","part_url":"https://rebrickable.com/parts/3001","part_img_url":""},"color":{"id":1,"name":"Blue"},"is_spare":false,"set_num":""}]}`

	// User sets endpoint
	case bytes.Contains([]byte(url), []byte("/sets")):
		body = `{"count":1,"next":null,"previous":null,"results":[{"set":{"set_num":"10001","name":"Test Set","year":2020,"theme_id":1,"num_parts":100,"set_img_url":"","set_url":""},"quantity":1,"include_spares":false}]}`

	// Set lists endpoint
	case bytes.Contains([]byte(url), []byte("/setlists")):
		if bytes.Contains([]byte(url), []byte("/1234/sets")) {
			// GetSetListSets endpoint for ID 1234
			body = `{"count":1,"next":null,"previous":null,"results":[{"set":{"set_num":"10002","name":"Set List Set","year":2020,"theme_id":1,"num_parts":50,"set_img_url":"","set_url":""},"quantity":1,"include_spares":false}]}`
		} else if bytes.Contains([]byte(url), []byte("/1234")) {
			// GetSetList endpoint for ID 1234
			body = `{"id":1234,"is_buildable":true,"name":"Test Set List","num_sets":1}`
		} else {
			// GetSetLists endpoint (list of all)
			body = `{"count":1,"next":null,"previous":null,"results":[{"id":1234,"is_buildable":true,"name":"Test Set List","num_sets":1}]}`
		}

	// Part lists endpoint
	case bytes.Contains([]byte(url), []byte("/partlists")):
		body = `{"count":1,"next":null,"previous":null,"results":[{"id":5678,"name":"Test Part List","num_parts":1}]}`

	default:
		body = `{"count":0,"next":null,"previous":null,"results":[]}`
	}

	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(body))),
		Header:     make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp, nil
}

// Helper to set up a test UsersAPI with our fake transport
func setupFakeUsersAPI(_ *testing.T) func() {
	prev := utils.GetUsersAPIOverride
	utils.GetUsersAPIOverride = func() (*api.UsersAPI, error) {
		client := &http.Client{Transport: &fakeRoundTripperV2{}}
		creds := &api.Credentials{UserName: "u", Password: "p", APIKey: "k"}
		return api.NewUsersAPI(client, creds, false), nil
	}
	return func() { utils.GetUsersAPIOverride = prev }
}

func TestGetAllUserPartsToolWithFakeUsersAPI(t *testing.T) {
	cleanup := setupFakeUsersAPI(t)
	defer cleanup()

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
}

func TestGetAllUserSetsToolWithFakeUsersAPI(t *testing.T) {
	cleanup := setupFakeUsersAPI(t)
	defer cleanup()

	_, sets, err := GetAllUserSets(context.Background(), nil, GetAllUserSetsInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sets == nil || len(sets.Sets) != 1 {
		t.Fatalf("expected 1 set, got %+v", sets)
	}

	s := sets.Sets[0]
	if s.Set.Number != "10001" {
		t.Fatalf("expected set number 10001, got %s", s.Set.Number)
	}
	if s.Set.Name != "Test Set" {
		t.Fatalf("expected set name 'Test Set', got %s", s.Set.Name)
	}
}

func TestGetUserSetListsToolWithFakeUsersAPI(t *testing.T) {
	cleanup := setupFakeUsersAPI(t)
	defer cleanup()

	_, lists, err := GetUserSetLists(context.Background(), nil, GetUserSetListsInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if lists == nil {
		t.Fatal("expected lists, got nil")
	}

	l := lists.SetLists[0]
	if l.ID != 1234 {
		t.Fatalf("expected list ID 1234, got %d", l.ID)
	}
	if l.Name != "Test Set List" {
		t.Fatalf("expected list name 'Test Set List', got %s", l.Name)
	}
}

func TestGetUserSetListToolWithFakeUsersAPI(t *testing.T) {
	cleanup := setupFakeUsersAPI(t)
	defer cleanup()

	_, list, err := GetUserSetList(context.Background(), nil, GetUserSetListInput{SetListID: 1234})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if list == nil {
		t.Fatal("expected set list, got nil")
	}
	if list.ID != 1234 {
		t.Fatalf("expected list ID 1234, got %d", list.ID)
	}
	if list.Name != "Test Set List" {
		t.Fatalf("expected list name 'Test Set List', got %s", list.Name)
	}
}

func TestGetSetsOfUserSetListToolWithFakeUsersAPI(t *testing.T) {
	cleanup := setupFakeUsersAPI(t)
	defer cleanup()

	_, sets, err := GetSetsOfUserSetList(context.Background(), nil, GetSetsOfUserSetListInput{SetListID: 1234})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sets == nil || len(sets.Sets) != 1 {
		t.Fatalf("expected 1 set, got %+v", sets)
	}

	s := sets.Sets[0]
	if s.Set.Number != "10001" {
		t.Fatalf("expected set number 10001, got %s", s.Set.Number)
	}
	if s.Set.Name != "Test Set" {
		t.Fatalf("expected set name 'Test Set', got %s", s.Set.Name)
	}
}

func TestGetUsersPartListsToolWithFakeUsersAPI(t *testing.T) {
	cleanup := setupFakeUsersAPI(t)
	defer cleanup()

	_, lists, err := GetUsersPartLists(context.Background(), nil, GetUsersPartListsInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if lists == nil {
		t.Fatal("expected lists, got nil")
	}

	l := lists.PartLists[0]
	if l.ID != 5678 {
		t.Fatalf("expected list ID 5678, got %d", l.ID)
	}
	if l.Name != "Test Part List" {
		t.Fatalf("expected list name 'Test Part List', got %s", l.Name)
	}
}
