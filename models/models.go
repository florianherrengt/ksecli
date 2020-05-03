package models

// Secrets struct for secrets.yml
type Secrets struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
	}
	Data struct {
		DatabaseUsername string `yaml:"DATABASE_USERNAME"`
		DatabasePassword string `yaml:"DATABASE_PASSWORD"`
	}
}
