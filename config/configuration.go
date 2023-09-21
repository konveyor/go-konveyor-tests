package config

type Configuration struct {
	KUBECONFIG    string `env:"KUBECONFIG"`
	JIRA_USERNAME string `env:"JIRA_USERNAME"`
	JIRA_PASSWORD string `env:"JIRA_PASSWORD"`
	JIRA_URL      string `env:"JIRA_URL"`
	RETRY_NUM     string `env:"RETRY_NUM"`
	KEEP          string `env:"KEEP"`
}
