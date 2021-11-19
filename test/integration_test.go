package test

import (
	"bytes"
	"encoding/json"
	"github.com/wwt/go-mongo-rest/lib/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"reflect"
	"testing"
)

// TODO build a docker container, startup mongo with docker compose, run tests
// For now, start mongo and this server manually.

const server = "http://127.0.0.1:8080"

func Test_Authors(t *testing.T) {
	var insertedAuthorId primitive.ObjectID

	t.Run("Create author fails due to required fields", func(t *testing.T) {
		var author models.Author
		res := Update("POST", "/authors", author)
		if res.StatusCode != 400 {
			t.Error("Expected an error since name is required")
		}
	})
	t.Run("Create author works", func(t *testing.T) {
		var author, response models.Author
		author.Name = "John Doe"

		res := Update("POST", "/authors", author)
		if res.StatusCode != 200 {
			t.Error(res.StatusCode)
		}
		unmarshal(res, &response)
		author.ID = response.ID
		author.Created = response.Created
		author.Updated = response.Updated
		if !reflect.DeepEqual(author, response) {
			t.Error("Expected", author, "got", response)
		}
		insertedAuthorId = response.ID
	})
	t.Run("List author works", func(t *testing.T) {
		var response []models.Author

		res := GET("/authors", &response)
		if res.StatusCode != 200 {
			t.Error(res.StatusCode)
		}
		if len(response) != 1 {
			t.Fatal("Expected 1 got", len(response))
		}
		if response[0].Name != "John Doe" {
			t.Error(response[0].Name)
		}

		res = GET("/authors?name=hello", &response)
		if res.StatusCode != 200 {
			t.Error(res.StatusCode)
		}
		if len(response) != 0 {
			t.Error("Expected 0 got", len(response))
		}

		res = GET("/authors?name=John+Doe", &response)
		if res.StatusCode != 200 {
			t.Error(res.StatusCode)
		}
		if len(response) != 1 {
			t.Error("Expected 1 got", len(response))
		}
	})
	t.Run("Patch author works", func(t *testing.T) {
		author := models.Author{
			Books: []models.Book{
				{Title: "Adventure"},
			},
		}
		{
			res := Update("PATCH", "/authors/"+insertedAuthorId.Hex(), &author)
			if res.StatusCode != 200 {
				var err string
				unmarshal(res, &err)
				t.Fatal(res.StatusCode, err)
			}

			var response1 models.Author
			unmarshal(res, &response1)
			// verify the patch didn't clear the name
			if author.Name != response1.Name && response1.Name != "John Doe" {
				t.Error(response1.Name)
			}
			if len(author.Books) != 1 {
				t.Error(len(author.Books))
			}
			if author.Books[0].Genre != "" {
				t.Error(author.Books[0].Genre)
			}
		}

		{
			// now try patching a nested value
			author.Books[0].Genre = "fantasy"
			res2 := Update("PATCH", "/authors/"+insertedAuthorId.Hex(), &author)
			if res2.StatusCode != 200 {
				t.Error(res2.StatusCode)
			}

			var response2 models.Author
			unmarshal(res2, &response2)
			if !reflect.DeepEqual(author.Books, response2.Books) {
				t.Error(response2.Books)
			}
			if author.Books[0].Title != "Adventure" {
				t.Error(author.Books[0].Title)
			}
		}
	})
	t.Run("Delete author works", func(t *testing.T) {
		res := DELETE("/authors/"+insertedAuthorId.Hex(), nil)
		if res.StatusCode != 200 {
			t.Error(res.StatusCode)
		}
		var response []models.Author
		res = GET("/authors", &response)
		if res.StatusCode != 200 {
			t.Error(res.StatusCode)
		}
		if len(response) != 0 {
			t.Error(len(response))
		}
	})
}

func GET(path string, response interface{}) *http.Response {
	req, err := http.NewRequest("GET", server+path, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err)
	}
	return res
}

func DELETE(path string, response interface{}) *http.Response {
	req, err := http.NewRequest("DELETE", server+path, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return res
}

func Update(method, path string, payload interface{}) *http.Response {
	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(method, server+path, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return res
}

func unmarshal(res *http.Response, response interface{}) {
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err)
	}
}
