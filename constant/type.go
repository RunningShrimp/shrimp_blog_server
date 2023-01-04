package constant

import (
	"encoding/json"
	"strconv"
)

type StringUint uint
type StringInt int

func (u *StringUint) UnmarshalJSON(b []byte) error {
	var item any
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*u = StringUint(v)
	case float64:
		*u = StringUint(int(v))
	case string:
		///here convert the string into
		///an integer
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err
		}
		*u = StringUint(i)
	default:
		//log.Logger.Error("没有匹配的类型")
	}
	return nil
}

func (s *StringInt) UnmarshalJSON(b []byte) error {
	var item any
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*s = StringInt(v)
	case float64:
		*s = StringInt(int(v))

	case string:
		///here convert the string into
		///an integer
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err
		}
		*s = StringInt(i)
	default:
		//log.Logger.Error("没有匹配的类型")
	}
	return nil
}
