package country

import (
	"encoding/json"
	"fghpdf.me/gin-project-template/internal/pkg/thunes/httpClient"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	var server = ServerMock("/v2/money-transfer/countries", listSuccessMock)
	defer server.Close()

	authClient := &httpClient.AuthClient{
		Username: "mock API KEY",
		Password: "mock API SECRET",
		BasicUrl: server.URL,
	}

	params := &ListParams{
		Page:    0,
		PerPage: 50,
	}

	svc := NewServer(authClient)

	res, err := svc.List(params)
	if err != nil {
		t.Error(err)
	}

	if res == nil || len(*res) != 3 {
		t.Errorf("expected 3 countries but got %d\n", len(*res))
	}

	if (*res)[0].IsoCode != "KEN" {
		t.Errorf("expected KEN but got %s\n", (*res)[0].IsoCode)
	}
}

func ServerMock(url string, mockHandler func(http.ResponseWriter, *http.Request)) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(url, mockHandler)

	server := httptest.NewServer(handler)
	return server
}

func listSuccessMock(w http.ResponseWriter, _ *http.Request) {
	countries := &[3]Model{
		{
			Name:    "Kenya",
			IsoCode: "KEN",
		},
		{
			Name:    "China",
			IsoCode: "CHN",
		},
		{
			Name:    "Singapore",
			IsoCode: "SGP",
		},
	}

	res, _ := json.Marshal(countries)
	_, _ = w.Write(res)
}
