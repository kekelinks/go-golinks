package golinks

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testLinkName = "test-link"
)

func getLinkJSON(gid, cid int) string {
	return fmt.Sprintf(`{
	"gid": %d,
	"cid": %d,
	"name": "%s"
}`, gid, cid, testLinkName)
}

func getLink(gid, cid int) *Link {
	return &Link{
		GID:  gid,
		CID:  cid,
		Name: testLinkName,
	}
}

func TestTeamsService_Retrieve(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	tcs := map[string]struct {
		id   int
		want *Link
	}{
		"ok": {1, getLink(1, 1)},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			mux.HandleFunc(fmt.Sprintf("/%s/%d", linksPath, tc.id), func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, fmt.Sprintf(getLinkJSON(tc.id, tc.id)))
			})

			got, err := client.Links.Retrieve(context.Background(), tc.id)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Fatalf("Diff: %s(-got +want)", diff)
			}
		})
	}
}
