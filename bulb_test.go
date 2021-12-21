package lightbrite

import "testing"

func Test_EncodePilotRequest(t *testing.T) {
	tests := []struct {
		name     string
		inp      PilotRequest
		expected []byte
	}{
		{
			name: "Encode successfully",
			inp: PilotRequest{
				Method: "getPath",
				Params: PilotParams{
					"r": 255,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tt.inp.toJSON()
			if len(out) == 0 {
				t.Fail()
			}
		})
	}
}

func Test_DecodePilotResponse(t *testing.T) {
	tests := []struct {
		name     string
		inp      []byte
		expected PilotResponse
	}{
		{
			name: "Decode successfully",
			inp: []byte(`{"result": {"mac": "xxxxx"}}
		`),
			expected: PilotResponse{
				Result: make(map[string]interface{}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := PilotResponseFromJSON(tt.inp)
			if out.Result == nil {
				t.Logf("Got: %v", out)
				t.Fail()
			}
		})
	}
}
