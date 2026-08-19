package urlshortner

import (
	"encoding/json"
	"errors"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		dest, found := pathsToUrls[path]
		if found {
			http.Redirect(writer, request, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(writer, request)
	}
}

func DataHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {

	pathUrls, err := parseData(data)
	if err != nil {
		return nil, err
	}

	pathsToUrls := makeMap(pathUrls)

	// make use of our existing MapHandler
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func parseData(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	switch {
	case json.Valid(data):
		err := json.Unmarshal(data, &pathUrls)
		if err != nil {
			return nil, err
		}
	case validYAML(data):
		err := yaml.Unmarshal(data, &pathUrls)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Unsupported or invalid data format: expected JSON or YAML")
	}
	return pathUrls, nil
}

func makeMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

func validYAML(data []byte) bool {
	var out yaml.MapSlice
	return yaml.Unmarshal(data, &out) == nil
}
