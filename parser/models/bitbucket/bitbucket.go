package bitbucket

type Cache struct {
}

type Step struct {
	Name       string   `yaml:"name"`
	Image      string   `yaml:"image,omitempty"`
	Caches     []Cache  `yaml:"caches,omitempty"`
	Script     []string `yaml:"script"`
	Deployment string   `yaml:"deployment,omitempty"`
}

type Pipeline struct {
	Steps []Step `yaml:"step,omitempty"`
}

type BitbucketPipeline struct {
	Image     string                `yaml:"image,omitempty"`
	Pipelines map[string][]Pipeline `yaml:"pipelines,omitempty"`
}
