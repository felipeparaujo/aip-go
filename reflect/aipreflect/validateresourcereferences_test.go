package aipreflect

import (
	"strconv"
	"testing"

	examplefreightv1 "go.einride.tech/aip/proto/gen/einride/example/freight/v1"
	"go.einride.tech/aip/resourcename"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"
)

func TestValidateResourceReferences(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name          string
		message       proto.Message
		errorContains string
	}{
		{
			name:    "empty",
			message: &examplefreightv1.Shipment{},
		},

		{
			name: "valid",
			message: &examplefreightv1.Shipment{
				OriginSite:      "shippers/1/sites/1",
				DestinationSite: "shippers/1/sites/2",
			},
		},

		{
			name:    "valid repeated empty",
			message: &examplefreightv1.BatchGetSitesRequest{},
		},

		{
			name: "valid repeated",
			message: &examplefreightv1.BatchGetSitesRequest{
				Names: []string{
					"shippers/1/sites/1",
					"shippers/1/sites/2",
				},
			},
		},

		{
			name: "invalid",
			message: &examplefreightv1.Shipment{
				OriginSite:      "shippers/1",
				DestinationSite: "shippers/1/sites/2",
			},
			errorContains: "value 'shippers/1' of field origin_site is not a valid resource reference" +
				" for freight-example.einride.tech/Site",
		},

		{
			name: "invalid nested",
			message: &examplefreightv1.CreateShipmentRequest{
				Shipment: &examplefreightv1.Shipment{
					OriginSite:      "shippers/1",
					DestinationSite: "shippers/1/sites/2",
				},
			},
			errorContains: "value 'shippers/1' of field shipment.origin_site is not a valid resource reference" +
				" for freight-example.einride.tech/Site",
		},

		{
			name: "invalid repeated",
			message: &examplefreightv1.BatchGetSitesRequest{
				Names: []string{
					"shippers/1/sites/1",
					"shippers/1",
				},
			},
			errorContains: "value 'shippers/1' of field names[1] is not a valid resource reference" +
				" for freight-example.einride.tech/Site",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := ValidateResourceReferences(tt.message)
			if tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
			} else {
				assert.NilError(t, err)
			}
		})
	}
}

func TestValidateResourceReferencesWithValidator(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name          string
		message       proto.Message
		errorContains string
	}{
		{
			name:    "empty",
			message: &examplefreightv1.Shipment{},
		},

		{
			name: "valid",
			message: &examplefreightv1.Shipment{
				OriginSite:      "shippers/1/sites/1",
				DestinationSite: "shippers/1/sites/2",
			},
		},

		{
			name:    "valid repeated empty",
			message: &examplefreightv1.BatchGetSitesRequest{},
		},

		{
			name: "valid repeated",
			message: &examplefreightv1.BatchGetSitesRequest{
				Names: []string{
					"shippers/1/sites/1",
					"shippers/1/sites/2",
				},
			},
		},

		{
			name: "invalid",
			message: &examplefreightv1.Shipment{
				OriginSite:      "shippers/1",
				DestinationSite: "shippers/1/sites/2",
			},
			errorContains: "value 'shippers/1' of field origin_site is not a valid resource reference" +
				" for freight-example.einride.tech/Site",
		},

		{
			name: "invalid nested",
			message: &examplefreightv1.CreateShipmentRequest{
				Shipment: &examplefreightv1.Shipment{
					OriginSite:      "shippers/1",
					DestinationSite: "shippers/1/sites/2",
				},
			},
			errorContains: "value 'shippers/1' of field shipment.origin_site is not a valid resource reference" +
				" for freight-example.einride.tech/Site",
		},

		{
			name: "invalid repeated",
			message: &examplefreightv1.BatchGetSitesRequest{
				Names: []string{
					"shippers/1/sites/1",
					"shippers/1",
				},
			},
			errorContains: "value 'shippers/1' of field names[1] is not a valid resource reference" +
				" for freight-example.einride.tech/Site",
		},

		{
			name: "failed validator",
			message: &examplefreightv1.BatchGetSitesRequest{
				Names: []string{
					"shippers/notinteger/sites/1",
					"shippers/1",
				},
			},
			errorContains: "strconv.Atoi: parsing \"notinteger\": invalid syntax",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := ValidateResourceReferencesWithValidator(tt.message, isInteger)
			if tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
			} else {
				assert.NilError(t, err)
			}
		})
	}

}

func isInteger(pattern, value, messagePath string) error {
	patternScanner := resourcename.Scanner{}
	valueScanner := resourcename.Scanner{}

	patternScanner.Init(pattern)
	valueScanner.Init(value)

	for patternScanner.Scan() && valueScanner.Scan() {
		if patternScanner.Segment().IsVariable() {
			if _, err := strconv.Atoi(valueScanner.Segment().Literal().ResourceID()); err != nil {
				return err
			}
		}
	}

	return nil
}
