package lcd

import "encoding/json"

type (
	IPALListBody struct {
		Height string       `json:"height"`
		Result []IPALResult `json:"result"`
	}
)

func (c *client) QueryIPALList() (IPALListBody, error) {
	var r IPALListBody

	if _, body, err := c.httpClient.Get(UriQueryIPALList, nil); err != nil {
		return r, err
	} else {
		if err := json.Unmarshal(body, &r); err != nil {
			return r, err
		} else {
			return r, nil
		}
	}
}
