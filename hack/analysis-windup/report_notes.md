# Windup report notes

Executed on MTA 6.2.

# TIER0 applications analysis

```
$ KEEP=1 go test -v -count=1 ./hack/analysis-windup/
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-14T10:56:23+02:00 level=info msg=[addon] Addon (adapter) created.
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
=== RUN   TestApplicationAnalysis
=== RUN   TestApplicationAnalysis/WindupCustomer_Tomcat_Legacy_-_shoud_never_fail
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] |201|  POST /hub/applications
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-14T10:56:23+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:56:28+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:56:33+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:56:38+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:56:43+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:56:48+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:56:53+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:56:58+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:03+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:08+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:13+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:18+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:23+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:28+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/tasks/1
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/index.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/migration_issues.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/ApplicationDetails_Order_Management.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/pom_xml.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/persistence_properties.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/OrderManagementAppInitializer_java.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1/bucket/windup/report/reports/Customer_java.html
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/applications/1
## GOT ANALYSIS OUTPUT FOR "Customer Tomcat Legacy - shoud never fail":
api.Analysis{
    Effort: 1,
    Issues: []api.Issue{
        {
            Category: "cloud-mandatory",
            Description: "Trivial change or 1-1 library swap",
            Effort: 1,
            RuleSet: "DiscoverHardcodedIPAddressRuleProvider",
            Rule: "DiscoverHardcodedIPAddressRuleProvider",
            Incidents: []api.Incident{
                {
                    File: "persistence.properties",
                    Line: 2,
                    Message: "Hard-coded IP: 169.60.225.216When migrating environments, hard-coded IP addresses may need to be modified or eliminated.",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "database",
            Rule: "database-03200",
            Incidents: []api.Incident{
                {
                    File: "pom.xml",
                    Line: 0,
                    Message: "The application embeds Spring Data JPA",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "mvc",
            Rule: "mvc-01220",
            Incidents: []api.Incident{
                {
                    File: "pom.xml",
                    Line: 0,
                    Message: "The application embeds a Spring MVC library.",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "embedded-framework",
            Rule: "embedded-framework-08400",
            Incidents: []api.Incident{
                {
                    File: "pom.xml",
                    Line: 0,
                    Message: "The application embeds the Spring Web framework.",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "observability",
            Rule: "observability-0100",
            Incidents: []api.Incident{
                {
                    File: "pom.xml",
                    Line: 0,
                    Message: "The application embeds Spring Boot Actuator.",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "technology-usage-javaee",
            Rule: "javaee-technology-usage-00120",
            Incidents: []api.Incident{
                {
                    File: "OrderManagementAppInitializer.java",
                    Line: 0,
                    Message: "The application uses Java Servlets",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "DiscoverMavenProjectsRuleProvider",
            Rule: "DiscoverMavenProjectsRuleProvider_1",
            Incidents: []api.Incident{
                {
                    File: "pom.xml",
                    Line: 0,
                    Message: "Maven Project Object Model (POM) File",
                },
            },
        },
    },
}
## GOT TAGS FOR "Customer Tomcat Legacy - shoud never fail":
[]api.Tag{
    {Name: "Servlet"},
    {Name: "Spring MVC"},
    {Name: "Spring Boot Actuator"},
    {Name: "Properties"},
    {Name: "JPA entities"},
    {Name: "Spring Data JPA"},
    {Name: "Spring Web"},
}
=== RUN   TestApplicationAnalysis/WindupPathfinder_example1_cloud-readiness_with_tagger
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |201|  POST /hub/applications
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-14T10:57:33+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:57:38+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:57:43+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:57:48+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:57:53+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:57:58+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:03+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:08+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:13+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:18+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:23+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:28+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:33+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:38+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:43+02:00 level=info msg=[binding] |200|  GET /hub/tasks/2
time=2023-09-14T10:58:43+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/index.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/migration_issues.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/ApplicationDetails_tackle_pathfinder.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/pom_xml.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/MavenWrapperDownloader_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/maven_wrapper_properties.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/application_properties.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/microprofile_config_properties.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentsResource_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/QuestionnairesResource_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/TranslatorResource_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/ApplicationDto_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentCreateCommand_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentSvc_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/BulkCreateSvc_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/TranslatorSvc_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/windup_ruleproviders.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentCategory_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentQuestion_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentQuestionnaire_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentSingleOption_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentStakeholder_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentStakeholdergroup_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentBulk_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/AssessmentBulkApplication_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/TranslatedText_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/Category_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/Question_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/Questionnaire_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2/bucket/windup/report/reports/SingleOption_java.html
time=2023-09-14T10:58:44+02:00 level=info msg=[binding] |200|  GET /hub/applications/2
## GOT ANALYSIS OUTPUT FOR "Pathfinder example1 cloud-readiness with tagger":
api.Analysis{
    Effort: 5,
    Issues: []api.Issue{
        {
            Category: "cloud-mandatory",
            Description: "Trivial change or 1-1 library swap",
            Effort: 5,
            RuleSet: "local-storage",
            Rule: "local-storage-00001",
            Incidents: []api.Incident{
                {
                    File: "MavenWrapperDownloader.java",
                    Line: 111,
                    Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nTwelve-Factor App: Backing servicesTwelve-Factor App: ConfigOpenShift Container Platform: Input secrets and ConfigMapsOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: LogsOpenShift Container Platform: Understanding persistent storage",
                },
                {
                    File: "MavenWrapperDownloader.java",
                    Line: 50,
                    Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nTwelve-Factor App: Backing servicesOpenShift Container Platform: Input secrets and ConfigMapsTwelve-Factor App: LogsOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: ConfigOpenShift Container Platform: Understanding persistent storage",
                },
                {
                    File: "MavenWrapperDownloader.java",
                    Line: 78,
                    Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nOpenShift Container Platform: Input secrets and ConfigMapsTwelve-Factor App: ConfigOpenShift Container Platform: Understanding persistent storageTwelve-Factor App: LogsOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: Backing services",
                },
                {
                    File: "MavenWrapperDownloader.java",
                    Line: 60,
                    Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nTwelve-Factor App: Backing servicesTwelve-Factor App: LogsTwelve-Factor App: ConfigOpenShift Container Platform: Input secrets and ConfigMapsOpenShift Container Platform: Understanding cluster loggingOpenShift Container Platform: Understanding persistent storage",
                },
                {
                    File: "MavenWrapperDownloader.java",
                    Line: 55,
                    Message: "An application running inside a container could lose access to a file in local storage.RecommendationsThe following recommendations depend on the function of the file in local storage:\n\n  Logging: Log to standard output and use a centralized log collector to analyze the logs.\n  Caching: Use a cache backing service.\n  Configuration: Store configuration settings in environment variables so that they can be updated without code changes.\n  Data storage: Use a database backing service for relational data or use a persistent data storage system.\n  Temporary data storage: Use the file system of a running container as a brief, single-transaction cache.\nOpenShift Container Platform: Understanding persistent storageOpenShift Container Platform: Understanding cluster loggingTwelve-Factor App: LogsTwelve-Factor App: Backing servicesTwelve-Factor App: ConfigOpenShift Container Platform: Input secrets and ConfigMaps",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "configuration-management",
            Rule: "configuration-management-0200",
            Incidents: []api.Incident{
                {
                    File: "application.properties",
                    Line: 0,
                    Message: "Application properties file detected",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "technology-usage-javaee",
            Rule: "javaee-technology-usage-00110",
            Incidents: []api.Incident{
                {
                    File: "AssessmentsResource.java",
                    Line: 0,
                    Message: "The application uses Bean Validation",
                },
                {
                    File: "TranslatorResource.java",
                    Line: 0,
                    Message: "The application uses Bean Validation",
                },
                {
                    File: "ApplicationDto.java",
                    Line: 0,
                    Message: "The application uses Bean Validation",
                },
                {
                    File: "AssessmentCreateCommand.java",
                    Line: 0,
                    Message: "The application uses Bean Validation",
                },
                {
                    File: "AssessmentSvc.java",
                    Line: 0,
                    Message: "The application uses Bean Validation",
                },
                {
                    File: "BulkCreateSvc.java",
                    Line: 0,
                    Message: "The application uses Bean Validation",
                },
                {
                    File: "TranslatorSvc.java",
                    Line: 0,
                    Message: "The application uses Bean Validation",
                },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "DiscoverMavenProjectsRuleProvider",
            Rule: "DiscoverMavenProjectsRuleProvider_1",
            Incidents: []api.Incident{
                {
                    File: "pom.xml",
                    Line: 0,
                    Message: "Maven Project Object Model (POM) File",
                },
            },
        },
    },
}
## GOT TAGS FOR "Pathfinder example1 cloud-readiness with tagger":
[]api.Tag{
    {Name: "CDI"},
    {Name: "HTML"},
    {Name: "Properties"},
    {Name: "JPA entities"},
    {Name: "JAX-RS"},
    {Name: "Bean Validation"},
    {Name: "Application Properties File"},
}
--- PASS: TestApplicationAnalysis (141.25s)
    --- PASS: TestApplicationAnalysis/WindupCustomer_Tomcat_Legacy_-_shoud_never_fail (70.51s)
    --- PASS: TestApplicationAnalysis/WindupPathfinder_example1_cloud-readiness_with_tagger (70.73s)
PASS
ok  	github.com/konveyor/go-konveyor-tests/hack/analysis-windup	141.310s
```

