package lcd

import "encoding/json"

type (
	AIPALListBody struct {
		Height string        `json:"height"`
		Result []AIPALResult `json:"result"`
	}
)

func (c *client) QueryAIPALList() (AIPALListBody, error) {
	var r AIPALListBody

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
