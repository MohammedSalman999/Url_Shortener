package main

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter , req *http.Request) {
		path := req.URL.Path
		if dest , ok := pathToUrls[path];ok {
			http.Redirect(w,req,dest,http.StatusFound)
			return
		}	
		fallback.ServeHTTP(w,req)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc,error) {
	pathUrls , err := parseYaml(yamlBytes)
	if err != nil {
		return nil , err
	}
	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback) , nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url string  `yaml:"url"`
}

func buildMap(pathUrls []pathUrl) (map[string]string) {
	pathToUrls := map[string]string{}
	for _ , pu := range pathUrls{
		pathToUrls[pu.Path] = pu.Url // key is the path and value is the url
	}
	return pathToUrls
}

// this function will parse the yaml 
func parseYaml(data []byte) ([]pathUrl , error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil , err
	}
	return pathUrls , nil
}