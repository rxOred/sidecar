package models

type Pipeline struct {
}

type BitbucketPipeline struct {
	Image     string     `yaml:"image"`
	Pipelines []Pipeline `yaml:"pipelines"`
}
