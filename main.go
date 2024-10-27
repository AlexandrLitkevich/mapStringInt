package main

import "encoding/json"

type MapStringInt map[string]int

func MarshalMapStringInt(m MapStringInt) ([]byte, error) {
	return json.Marshal(m)
}

func UnmarshalMapStringInt(data []byte) (MapStringInt, error) {
	var m MapStringInt
	err := json.Unmarshal(data, &m)
	return m, err
}
