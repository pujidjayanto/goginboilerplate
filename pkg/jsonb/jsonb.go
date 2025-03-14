package jsonb

import (
	"database/sql/driver"
	"fmt"

	"github.com/bytedance/sonic"
)

// JSONB stands for json binary or json better
// https://gorm.io/docs/data_types.html#Scanner-x2F-Valuer
type JSON map[string]interface{}

// Scan scan value into jsonb.JSON, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}

	return sonic.Unmarshal(b, &j)
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return sonic.Marshal(j)
}
