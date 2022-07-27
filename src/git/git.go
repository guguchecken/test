package git

import (
	"fmt"
	"net/http"
	"strings"
)

type pr struct {
	owner string
	name  string
	id_pr string

	token string
}

func NewPR(owner string, name string, id_pr string, token string) pr {
	return pr{
		owner: owner,
		name:  name,
		id_pr: id_pr,
		token: token,
	}
}

func (p pr) GetPaths() []string {
	query := `{
                repository(name: "matrixone", owner: "matrixorigin") {
                    pullRequest(number: 4174) {
                        files(first: 100) {
                            edges {
                                node {
                                    path
                                }
                            }
                        }
                    }
                }
            }`
	que := strings.NewReader(query)
	req, err := http.NewRequest("POST", `https://api.github.com/graphql`, que)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	req.Header.Set(`Authorization`, `bearer secret123`)
	req.Header.Set(`Content-Type`, `application/json`)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	t := make([]byte, 1024*4)
	_, err = resp.Body.Read(t)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	fmt.Printf("t: %v\n", string(t))
	return nil
}

func (p pr) onlyPaths() (ans []string) {
	return nil
}
