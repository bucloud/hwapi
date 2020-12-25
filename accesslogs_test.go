package hwapi_test

import (
	"flag"
	"net/http"
	"testing"
	"time"

	"github.com/bucloud/hwapi"
)

var (
	pj   *string
	ak   *string
	sk   *string
	user *string
	pwd  *string
)

func init() {
	pj = flag.String("privatekey", "", "provide private key json string, note this should be an base64 encoded string")
	ak = flag.String("keyid", "", "access key id")
	sk = flag.String("secretkey", "", "secret key string")
	user = flag.String("user", "", "striketracker username")
	pwd = flag.String("pwd", "", "striketracker password")
	// flag.Parse()
}
func TestSearchLogsV2(t *testing.T) {
	api := hwapi.Init(&http.Transport{})
	startDate := time.Now().UTC().Add(-time.Hour * 240)
	endDate := time.Now().UTC().Add(-time.Hour * 72)

	if _, err := api.SearchLogsV2(&hwapi.SearchLogsOptions{
		AccountHash: "b9h3f3y8",
		HostHash:    "f6g4s8v3",
		StartDate:   startDate,
		EndDate:     endDate,
		HCSCredentials: hwapi.HCSCredentials{
			AccessKeyID: *ak,
			SecretKey:   *sk,
		},
	}); err != nil {
		t.Error(err)
	}

	if _, err := api.SearchLogsV2(&hwapi.SearchLogsOptions{
		AccountHash: "b9h3f3y8",
		HostHash:    "f6g4s8v3",
		StartDate:   startDate,
		EndDate:     endDate,
		HCSCredentials: hwapi.HCSCredentials{
			PrivateKeyJSON: *pj,
		},
	}); err != nil {
		t.Error(err)
	}

	api.Auth(*user, *pwd, true)
	if _, err := api.SearchLogs("f6g4s8v3", "cds", startDate, endDate); err != nil {
		t.Error(err)
	}
}

func BenchmarkSearchLogsGCS(b *testing.B) {
	api := hwapi.Init(&http.Transport{})
	startDate := time.Now().UTC().Add(-time.Hour * 240)
	endDate := time.Now().UTC().Add(-time.Hour * 72)
	// pj := flag.String("privatekey", "", "provide private key json string, note this should be an base64 encoded string")
	if *pj == "" {
		b.Fatalf("privatekey missed\n")
	}
	for i := 0; i < b.N; i++ {
		if _, err := api.SearchLogsV2(&hwapi.SearchLogsOptions{
			AccountHash: "b9h3f3y8",
			HostHash:    "f6g4s8v3",
			StartDate:   startDate,
			EndDate:     endDate,
			HCSCredentials: hwapi.HCSCredentials{
				PrivateKeyJSON: *pj,
			},
		}); err != nil {
			b.Fatal(err)
		}

		// fmt.Printf("%d %v\n", len(gcsURLs), err)
	}
}

func BenchmarkSearchLogsAWS(b *testing.B) {
	api := hwapi.Init(&http.Transport{})
	startDate := time.Now().UTC().Add(-time.Hour * 240)
	endDate := time.Now().UTC().Add(-time.Hour * 72)
	// ak := flag.String("keyid", "", "access key id")
	// sk := flag.String("secretkey", "", "secret key string")
	if *ak == "" || *sk == "" {
		b.Fatalf("accesskey & secretkey missed\n")
	}
	for i := 0; i < b.N; i++ {
		if _, err := api.SearchLogsV2(&hwapi.SearchLogsOptions{
			AccountHash: "b9h3f3y8",
			HostHash:    "f6g4s8v3",
			StartDate:   startDate,
			EndDate:     endDate,
			HCSCredentials: hwapi.HCSCredentials{
				AccessKeyID: *ak,
				SecretKey:   *sk,
			},
		}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSearchLogsHCS(b *testing.B) {
	api := hwapi.Init(&http.Transport{})
	startDate := time.Now().UTC().Add(-time.Hour * 240)
	endDate := time.Now().UTC().Add(-time.Hour * 72)
	if _, err := api.Auth(*user, *pwd, true); err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		if _, err := api.SearchLogs("f6g4s8v3", "cds", startDate, endDate); err != nil {
			b.Fatal(err)
		}
	}
}
