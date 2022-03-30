package configread

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AwsConfig Aws `yaml:"aws"`
	// Need to think of what information is required
	// for each LDAP HOST and if they belong to a region.
	// Ldap      struct {
}

type Aws struct {
	Profiles []Profile `yaml:"profiles"`
}

type Profile struct {
	Region string `yaml:"region"`
	Name   string `yaml:"name"`
}

func unmarshalYAMLFile(path string, v interface{}) error {
	if path == "" {
		path = "conf/cool-api..yaml"
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(v)
}

// read the Yaml into struct(s)
func ParseYamlConfig(pathYaml string) Config {
	conf1 := Config{}
	if err := unmarshalYAMLFile(pathYaml, &conf1); err != nil {
		panic(err)
	}
	return conf1
}
