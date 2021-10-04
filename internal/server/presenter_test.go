package server

import "testing"

func TestJsonPresenter_Format(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"EmptyResp",
			args{data: struct{}{}},
			"{}",
			false,
		},
		{
			"SuccessResp",
			args{data: struct {
				Status string
				Href   string
			}{
				"ok",
				"https://api-m.paypal.com/v2/payments/authorizations/?&*()_)<>0VF52814937998046",
			},
			},
			`{"Status":"ok","Href":"https://api-m.paypal.com/v2/payments/authorizations/?\u0026*()_)\u003c\u003e0VF52814937998046"}`,
			false,
		},
		{
			"UnSuccessResp",
			args{data: struct {
				Status string
				Error  string
				Href   string
			}{
				"fail",
				"unknown payment service",
				"https://headway.onelink.me/8zSH/playstore",
			},
			},
			`{"Status":"fail","Error":"unknown payment service","Href":"https://headway.onelink.me/8zSH/playstore"}`,
			false,
		},
		{
			"MapResponse",
			args{data: map[int]interface{} {
				1: "str",
				2: []int{10, 20, 30},
				3: map[string]string{"x":"yyy"},
			},
			},
			`{"1":"str","2":[10,20,30],"3":{"x":"yyy"}}`,
			false,
		},
		{
			"InvalidType",
			args{data: make(chan int)},
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsonPresenter{}
			got, err := js.Format(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
