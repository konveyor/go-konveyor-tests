package config

type Configuration struct {
	KUBECONFIG string `env:"KUBECONFIG"`

	JIRA_CLOUD_URL       string `env:"JIRA_CLOUD_URL"`
	JIRA_CLOUD_USERNAME  string `env:"JIRA_CLOUD_USERNAME"`
	JIRA_CLOUD_PASSWORD  string `env:"JIRA_CLOUD_PASSWORD"`
	JIRA_SERVER_URL      string `env:"JIRA_SERVER_URL"`
	JIRA_SERVER_USERNAME string `env:"JIRA_SERVER_USERNAME"`
	JIRA_SERVER_PASSWORD string `env:"JIRA_SERVER_PASSWORD"`
	JIRA_SERVER_TOKEN    string `env:"JIRA_SERVER_TOKEN"`
	MAVEN_TESTAPP_USER          string `env:"MAVEN_TESTAPP_USER"`
	MAVEN_TESTAPP_TOKEN         string `env:"MAVEN_TESTAPP_TOKEN"`

	RETRY_NUM string `env:"RETRY_NUM"`
	KEEP      string `env:"KEEP"`
}
