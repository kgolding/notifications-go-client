package notify

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func jsonResponse(body io.ReadCloser, obj interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	err = json.Unmarshal([]byte(b), &obj)
	if err != nil {
		return fmt.Errorf("unmarshal: %w `%s`", err, string(b))
	}

	return nil
}
