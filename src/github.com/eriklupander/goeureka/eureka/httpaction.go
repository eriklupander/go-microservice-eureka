package eureka

type HttpAction struct {
        Method          string              `yaml:"method"`
        Url             string              `yaml:"url"`
        Body            string              `yaml:"body"`
        Template        string              `yaml:"template"`
        Accept          string              `yaml:"accept"`
        ContentType     string              `yaml:"contentType"`
        Title           string              `yaml:"title"`
        StoreCookie     string              `yaml:"storeCookie"`
}
