package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/driver005/gateway/models"
	// "github.com/gofiber/fiber/v3"
	// "github.com/gofiber/fiber/v3/middleware/cors"
	// "github.com/gofiber/fiber/v3/middleware/logger"
)

//Swagger: swag init -g main.go  --parseDependency --parseInternal --parseDepth 1  --output docs/

var ctx = context.Background()

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api
// @schemes http
func main() {
	// r := registry.New(ctx)

	// r.Setup()

	model := models.Address{
		Address1: "TEST",
	}

	fmt.Println(reflect.ValueOf(model.Address1).IsZero())
	fmt.Println(reflect.ValueOf(model.Address2).IsZero())
}

// var testUUID = uuid.Must(uuid.Parse("f47ac10b-58cc-0372-8567-0e02b2c3d479"))

// type S struct {
// 	ID1 uuid.UUID
// 	ID2 uuid.UUID `json:"ID2,omitempty"`
// }

// testCases := map[string]struct {
// 	data           []byte
// 	expectedError  error
// 	expectedResult uuid.UUID
// }{
// 	"success": {
// 		data:           []byte(`{"ID1": "f47ac10b-58cc-0372-8567-0e02b2c3d479"}`),
// 		expectedError:  nil,
// 		expectedResult: testUUID,
// 	},
// 	"zero": {
// 		data:           []byte(`{"ID1": "00000000-0000-0000-0000-000000000000"}`),
// 		expectedError:  nil,
// 		expectedResult: uuid.Nil,
// 	},
// 	"null": {
// 		data:           []byte(`{"ID1": null}`),
// 		expectedError:  nil,
// 		expectedResult: uuid.Nil,
// 	},
// 	"empty": {
// 		data:           []byte(`{"ID1": ""}`),
// 		expectedError:  errors.New("inavlaid length"),
// 		expectedResult: uuid.Nil,
// 	},
// 	"omitempty": {
// 		data:           []byte(`{"ID2": ""}`),
// 		expectedError:  errors.New("inavlaid length"),
// 		expectedResult: uuid.Nil,
// 	},
// }

// for name, tc := range testCases {
// 	fmt.Println(name)
// 	var s S
// 	if err := json.Unmarshal(tc.data, &s); err != nil {
// 		fmt.Printf("unexpected error: got %v, want %v", err, tc.expectedError)
// 	}
// 	if !reflect.DeepEqual(s.ID1, tc.expectedResult) {
// 		fmt.Printf("got %#v, want %#v", s.ID1, tc.expectedResult)
// 	}

// 	fmt.Print(s)
// }
