# Windup report notes

Executed on Konveyor 0.2.

# TIER0

```
$ KEEP=1 go test -v -count=1 ./hack/analysis-windup/
time=2023-09-08T14:54:09+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T14:54:09+02:00 level=info msg=[addon] Addon (adapter) created.
time=2023-09-08T14:54:09+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T14:54:10+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-08T14:54:10+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T14:54:10+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-08T14:54:10+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T14:54:11+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
=== RUN   TestApplicationAnalysis
=== RUN   TestApplicationAnalysis/Customer_Tomcat_Legacy_-_shoud_never_fail
time=2023-09-08T14:54:11+02:00 level=info msg=[binding] |201|  POST /hub/applications
time=2023-09-08T14:54:11+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-08T14:54:12+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:17+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:22+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:27+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:32+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:38+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:43+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:48+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:53+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:54:59+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:04+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:09+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:14+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:19+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:24+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:30+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:35+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:40+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:45+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T14:55:47+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/index.html
time=2023-09-08T14:55:48+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/migration_issues.html
time=2023-09-08T14:55:48+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/ApplicationDetails_Order_Management.html
time=2023-09-08T14:55:49+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/pom_xml.html
time=2023-09-08T14:55:49+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/persistence_properties.html
Cannot match Incident $("<a name='126976' class='windup-file-location'></a><div class='inline-source-comment green tag-none'><div class='inline-comment'><div class='inline-comment-heading'><strong class='notification warning'>Hard-coded IP address</strong><a title='View Rule: DiscoverHardcodedIPAddressRuleProvider' href='windup_ruleproviders.html#DiscoverHardcodedIPAddressRuleProvider'><span class='glyphicon glyphicon-link rule-link floatRight'></span></a></div><div class='inline-comment-body'><p><strong>Hard-coded IP: 169.60.225.216<\/strong><\/p><p>When migrating environments, hard-coded IP addresses may need to be modified or eliminated.<\/p></div></div></div>").appendTo('#2-inlines'); with 'windup_ruleproviders.html#([a-z0-9-]+)'
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:persistence_properties.html Line:2 Message: CodeSnip: Facts:map[]}
time=2023-09-08T14:55:49+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/OrderManagementAppInitializer_java.html
time=2023-09-08T14:55:49+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/Customer_java.html
WINDUP ANALYSIS OUTPUT FOR "Customer Tomcat Legacy - shoud never fail":api.Analysis{
    Effort: 1,
    Issues: []api.Issue{
        {
            Category: "cloud-mandatory",
            Description: "Trivial change or 1-1 library swap",
            Effort: 1,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            {
                File: "persistence_properties.html",
                Line: "2",
                Message: "",
            },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
    }
}
    analysis_test.go:108: Different effort error. Got 1, expected 2
=== RUN   TestApplicationAnalysis/Pathfinder_example1_cloud-readiness_with_tagger
time=2023-09-08T14:55:49+02:00 level=info msg=[binding] |201|  POST /hub/applications
time=2023-09-08T14:55:50+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-08T14:55:50+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:55:55+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:00+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:05+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:11+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:16+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:21+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:26+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:31+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:37+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:42+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:47+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:52+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:56:57+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:03+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:08+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:13+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:18+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:23+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:29+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:34+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T14:57:35+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/index.html
time=2023-09-08T14:57:37+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/migration_issues.html
time=2023-09-08T14:57:37+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/ApplicationDetails_tackle_pathfinder.html
time=2023-09-08T14:57:37+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/pom_xml.html
time=2023-09-08T14:57:38+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/MavenWrapperDownloader_java.html
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:111 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:78 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:50 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:60 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:55 Message: CodeSnip: Facts:map[]}
time=2023-09-08T14:57:38+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/maven_wrapper_properties.html
time=2023-09-08T14:57:38+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/application_properties.html
time=2023-09-08T14:57:38+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/microprofile_config_properties.html
time=2023-09-08T14:57:38+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentsResource_java.html
time=2023-09-08T14:57:39+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/QuestionnairesResource_java.html
time=2023-09-08T14:57:39+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/TranslatorResource_java.html
time=2023-09-08T14:57:39+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/ApplicationDto_java.html
time=2023-09-08T14:57:39+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentCreateCommand_java.html
time=2023-09-08T14:57:39+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentSvc_java.html
time=2023-09-08T14:57:40+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/BulkCreateSvc_java.html
time=2023-09-08T14:57:40+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/TranslatorSvc_java.html
time=2023-09-08T14:57:40+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentCategory_java.html
time=2023-09-08T14:57:40+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentQuestion_java.html
time=2023-09-08T14:57:41+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentQuestionnaire_java.html
time=2023-09-08T14:57:41+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentSingleOption_java.html
time=2023-09-08T14:57:41+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentStakeholder_java.html
time=2023-09-08T14:57:41+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentStakeholdergroup_java.html
time=2023-09-08T14:57:41+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentBulk_java.html
time=2023-09-08T14:57:42+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/AssessmentBulkApplication_java.html
time=2023-09-08T14:57:42+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/TranslatedText_java.html
time=2023-09-08T14:57:42+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Category_java.html
time=2023-09-08T14:57:42+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Question_java.html
time=2023-09-08T14:57:42+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Questionnaire_java.html
time=2023-09-08T14:57:43+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/SingleOption_java.html
WINDUP ANALYSIS OUTPUT FOR "Pathfinder example1 cloud-readiness with tagger":api.Analysis{
    Effort: 5,
    Issues: []api.Issue{
        {
            Category: "cloud-mandatory",
            Description: "Trivial change or 1-1 library swap",
            Effort: 5,
            RuleSet: "",
            Rule: "local-storage-00001",
            Incidents: []api.Incident{
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "111",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "78",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "50",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "60",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "55",
                Message: "",
            },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
    }
}
```

## TIER1

```
$ TIER1=1 KEEP=1 go test -v -count=1 ./hack/analysis-windup/
time=2023-09-08T15:02:06+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:02:06+02:00 level=info msg=[addon] Addon (adapter) created.
time=2023-09-08T15:02:06+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:02:07+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-08T15:02:07+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:02:08+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-08T15:02:08+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:02:08+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
=== RUN   TestApplicationAnalysis
=== RUN   TestApplicationAnalysis/Petclinic_cloud-readiness_with_tagger_main
time=2023-09-08T15:02:08+02:00 level=info msg=[binding] |201|  POST /hub/applications
time=2023-09-08T15:02:09+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-08T15:02:09+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:14+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:19+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:25+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:30+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:35+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:40+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:46+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:51+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:02:56+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:01+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:06+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:12+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:17+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:22+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:27+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:32+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:38+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:43+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:48+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:53+02:00 level=info msg=[binding] |200|  GET /hub/tasks/3
time=2023-09-08T15:03:54+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/index.html
time=2023-09-08T15:03:58+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/index.html
time=2023-09-08T15:03:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/migration_issues.html
time=2023-09-08T15:03:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/ApplicationDetails_Spring_Framework_Petclinic.html
time=2023-09-08T15:03:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/pom_xml.html
time=2023-09-08T15:04:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/MavenWrapperDownloader_java.html
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:60 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:50 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:111 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:55 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:78 Message: CodeSnip: Facts:map[]}
time=2023-09-08T15:04:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/maven_wrapper_properties.html
time=2023-09-08T15:04:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/messages_properties.html
time=2023-09-08T15:04:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/messages_de_properties.html
time=2023-09-08T15:04:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/messages_en_properties.html
time=2023-09-08T15:04:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/messages_es_properties.html
time=2023-09-08T15:04:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/business_config_xml.html
time=2023-09-08T15:04:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/data_access_properties.html
time=2023-09-08T15:04:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/datasource_config_xml.html
time=2023-09-08T15:04:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/mvc_core_config_xml.html
time=2023-09-08T15:04:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/mvc_view_config_xml.html
time=2023-09-08T15:04:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/tools_config_xml.html
time=2023-09-08T15:04:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/mvc_test_config_xml.html
time=2023-09-08T15:04:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/Owner_java.html
time=2023-09-08T15:04:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/Pet_java.html
time=2023-09-08T15:04:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/PetType_java.html
time=2023-09-08T15:04:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/Specialty_java.html
time=2023-09-08T15:04:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/Vet_java.html
time=2023-09-08T15:04:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/Visit_java.html
time=2023-09-08T15:04:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/3/bucket/windup/report/reports/CallMonitoringAspect_java.html
WINDUP ANALYSIS OUTPUT FOR "Petclinic cloud-readiness with tagger main":api.Analysis{
    Effort: 5,
    Issues: []api.Issue{
        {
            Category: "cloud-mandatory",
            Description: "Trivial change or 1-1 library swap",
            Effort: 5,
            RuleSet: "",
            Rule: "local-storage-00001",
            Incidents: []api.Incident{
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "60",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "50",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "111",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "55",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "78",
                Message: "",
            },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
    }
}
=== RUN   TestApplicationAnalysis/Petclinic_legacy_cloud-readiness_with_tagger_and_hazelcast_custom_rules
time=2023-09-08T15:04:03+02:00 level=info msg=[binding] |201|  POST /hub/applications
    error.go:12: stat ./data/hz.windup.xml: no such file or directory
time=2023-09-08T15:04:03+02:00 level=info msg=[binding] |400|  POST /hub/rulesets
    error.go:12: POST /hub/rulesets failed: 400(Bad Request) body: {"error":"Key: 'RuleSet.Image.ID' Error:Field validation for 'ID' failed on the 'required' tag"}
time=2023-09-08T15:04:04+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-08T15:04:04+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:09+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:14+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:20+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:25+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:30+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:35+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:42+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:47+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:52+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:04:58+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:03+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:08+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:13+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:19+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:24+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:29+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:34+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:39+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:45+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:50+02:00 level=info msg=[binding] |200|  GET /hub/tasks/4
time=2023-09-08T15:05:51+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/index.html
    analyzer_legacy_windup.go:78: Error report contect check for /windup/report/index.html. Cannot find 12
        story points in 
        
        
        
        
        
        
        Application List
        
        
        
        
        
        
        body.viewAppList .apps{ margin: 0 2ex; }
        
        body.viewAppList .apps .appInfo {
        border-bottom: 1px solid gray;
        overflow: auto; width: 100%; /* clearing */
        padding: 1ex 0 1ex;
        }
        body.viewAppList .apps .appInfo .stats { float: right; width: 30%; padding: 0.4ex 0; }
        body.viewAppList .apps .appInfo .stats .effortPoints { float: left; width: 160px; padding: 0.3ex 0.2em 0; font-size: 33pt; }
        body.viewAppList .apps .appInfo .stats .effortPointsspan { display: block; margin: auto; text-align: center; }
        body.viewAppList .apps .appInfo .stats .effortPoints.points { line-height: 1; color: #294593; }
        body.viewAppList .apps .appInfo .stats .effortPoints.legend { font-size: 7pt; }
        body.viewAppList .apps .appInfo .stats .effortPoints.shared,
        body.viewAppList .apps .appInfo .stats .effortPoints.unique { width: 90px; font-size: 18pt; margin-top: 23px; }
        /* Hide the "cell" if the app has 0 shared points". */
        body.viewAppList .apps .appInfo.pointsShared0 .stats .effortPoints.shared,
        body.viewAppList .apps .appInfo.pointsShared0 .stats .effortPoints.unique { visibility: hidden; }
        /* Hide the whole "column" if there's no virtual app (i.e. no shared-libs app). */
        body.viewAppList.noVirtualApp .apps .appInfo.stats .effortPoints.shared,
        body.viewAppList.noVirtualApp .apps .appInfo.stats .effortPoints.unique { display: none; }
        body.viewAppList .apps .appInfo .stats .effortPoints.shared .points,
        body.viewAppList .apps .appInfo .stats .effortPoints.unique .points { color: #8491a8; /* Like normal, but grayed. */ }
        
        body.viewAppList .apps .appInfo .stats .incidentsCount { float: left; margin:0 2ex;}
        body.viewAppList .apps .appInfo .stats .incidentsCount table tr.total td { border-top: 1px solid silver; }
        body.viewAppList .apps .appInfo .stats .incidentsCount .count { text-align: right; padding-right: 1ex; min-width: 7.4ex; }
        body.viewAppList .apps .appInfo .traits { margin-left: 0px; width: 70%;}
        body.viewAppList .apps .appInfo .traits .fileName { padding: 0.0ex 0em 0.2ex; font-size: 18pt; /* color: #008cba; (Default BS link color) */ }
        body.viewAppList .apps .appInfo .traits .techs { }
        
        /* Specifics for virtual apps. */
        body.viewAppList .apps .virtual .appInfo .traits .fileName { color: #477280; }
        
        body.viewAppList .apps .appInfo:first-of-type { border-top: 1px solid gray; }
        
        
        
        
        var TARGET_RUNTIME = JSON.parse('[{"id":"eap","name":"JBoss EAP","description":"Red Hat JBoss Enterprise Application Platform","supported":["Java EE","JBoss Web XML","JMS Connection Factory","JAXB","Enterprise Web Services","CDI","JMS Topic","JSF Page","WS Metadata","Java EE Security","JAX-RS","Clustering EJB","Common Annotations","JAXR","Clustering Web Session","Stateful (SFSB)","SOAP (SAAJ)","Servlet","JPA named queries","EJB XML*","Java EE Batch","JSON-B","JSP Page","Persistence units","Stateless (SLSB)","Message (MDB)","JCA","JAX-WS","Java EE JSON-P","JMS Queue","Java EE Batch API","JMS","Bean Validation","JavaMail","MEJB","JACC","JBoss EJB XML","JTA","EAR","JPA XML*","JPA","EJB","JPA entities"],"unsuitable":["WebSphere Web XML","WebLogic Web XML","WebSphere EJB Ext","JAX-RPC","WebSphere WS Extension","WebSphere WS Binding","RMI","WebSphere EJB"],"neutral":["Spring MVC","JDBC XA datasources","Bouncy Castle","Manifest","Properties","Spring","WSDL","Java Source","JDBC datasources","Swagger","Maven XML","Web XML*","Spring Security"]},{"id":"jws","name":"JWS","description":"Red Hat JBoss Web Server","supported":["JSP Page","Bean Validation","JAXB","Servlet","JSF Page"],"unsuitable":["Java EE","WebSphere Web XML","JBoss Web XML","JMS Connection Factory","Enterprise Web Services","CDI","JMS Topic","WS Metadata","Java EE Security","JAX-RS","Clustering EJB","Common Annotations","JAXR","Clustering Web Session","Stateful (SFSB)","SOAP (SAAJ)","WebSphere EJB Ext","JPA named queries","EJB XML*","JAX-RPC","Java EE Batch","WebSphere EJB","Stateless (SLSB)","Message (MDB)","JCA","JAX-WS","Java EE JSON-P","WebSphere WS Extension","JMS Queue","WebSphere WS Binding","RMI","Java EE Batch API","JMS","JavaMail","MEJB","JACC","JBoss EJB XML","JTA","WebLogic Web XML","EAR","JPA XML*","JPA","EJB","JPA entities"],"neutral":["Java Source","JDBC datasources","Manifest","Maven XML","Properties","Web XML*"]}]');
        
        
        
        
        
        
        
        
        
        
        
        
        
        Migration Toolkit for Applications
        
        
        
        
        
        All Applications
        
        
        
        
        
        
        
        Technologies
        
        
        
        
        
        
        
        About
        
        
        
        
        
        Send Feedback 
        
        
        
        
        
        
        
        var FEEDBACK_JS_ADDED = false;
        var FEEDBACK_FORM_TRIGGER = null;
        
        function displayFeedbackForm() {
        FEEDBACK_FORM_TRIGGER();
        }
        
        window.ATL_JQ_PAGE_PROPS = {
        "triggerFunction": function(showCollectorDialog) {
        FEEDBACK_FORM_TRIGGER = showCollectorDialog;
        }
        };
        
        document.addEventListener("DOMContentLoaded", function(event) {
        jQuery(".jiraFeedbackTrigger").click(function(e) {
        e.preventDefault();
        displayFeedbackForm();
        });
        });
        
        
        
        var script= document.createElement("script");
        script.type= "text/javascript";
        script.src= "reports/resources/js/navbar.js";
        document.body.appendChild(script);
        
        
        
        
        
        
        
        
        
        Application List
        
        
        
        
        
        
        
        
        
        
        
        
        
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"> 
        
        
        
        
        
        
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"> 
        
        
        
        
        
        
        
        
        
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
        > 
        
        
        
        
        
        
        
        
        
        
        
        &times;
        Warning! Invalid search regular expression!
        
        
        
        
        Results
        Active filters:
        
        
        <!--
        
        
        Name: nameofthething
        
        
        
        -->
        
        Clear All Filters
        
        
        
        
        
        
        
        
        
        
        
        
        
        Runtime labels legend
        
        Supported
        Partially supported
        Unsuitable
        Neutral
        
        
        
        
        
        
        
        
        
        
        
        
        
        5
        story points
        
        
        
        
        Number of incidents
        
        
        5
        Cloud Mandatory
        
        
        13
        Information
        
        
        18 
        Total 
        
        
        
        
        
        
        
        
        spring-framework-petclinic
        
        
        
        
        Bean Validation 
        
        
        Common Annotations 
        
        
        JAXB 
        
        
        JPA entities 
        
        
        JSP Page 
        
        
        Java Source 
        
        
        Maven XML 
        
        
        Properties 
        
        
        Servlet 
        
        
        Spring Data JPA 
        
        
        Spring JMX 
        
        
        Spring MVC 
        
        
        
        
        
        
        
        $("body").addClass("noVirtualApp");
        
        
        Rule providers execution overview
        |
        FreeMarker methods
        
        
        Page generated: Sep 8, 2023, 1:05:34 PM
        
        
        
        
        $("body.viewAppList .apps .real .appInfo").sortElements(function(a, b){
        return $(a).find(".traits .fileName").first().text().trim().toLowerCase() > $(b).find(".traits .fileName").first().text().trim().toLowerCase() ? 1 : -1;
        });
        
        
        $(document).ready(function(){$('[data-toggle="tooltip"]').tooltip();});
        
        
    analyzer_legacy_windup.go:78: Error report contect check for /windup/report/index.html. Cannot find 8
        Cloud Mandatory in 
        
        
        
        
        
        
        Application List
        
        
        
        
        
        
        body.viewAppList .apps{ margin: 0 2ex; }
        
        body.viewAppList .apps .appInfo {
        border-bottom: 1px solid gray;
        overflow: auto; width: 100%; /* clearing */
        padding: 1ex 0 1ex;
        }
        body.viewAppList .apps .appInfo .stats { float: right; width: 30%; padding: 0.4ex 0; }
        body.viewAppList .apps .appInfo .stats .effortPoints { float: left; width: 160px; padding: 0.3ex 0.2em 0; font-size: 33pt; }
        body.viewAppList .apps .appInfo .stats .effortPointsspan { display: block; margin: auto; text-align: center; }
        body.viewAppList .apps .appInfo .stats .effortPoints.points { line-height: 1; color: #294593; }
        body.viewAppList .apps .appInfo .stats .effortPoints.legend { font-size: 7pt; }
        body.viewAppList .apps .appInfo .stats .effortPoints.shared,
        body.viewAppList .apps .appInfo .stats .effortPoints.unique { width: 90px; font-size: 18pt; margin-top: 23px; }
        /* Hide the "cell" if the app has 0 shared points". */
        body.viewAppList .apps .appInfo.pointsShared0 .stats .effortPoints.shared,
        body.viewAppList .apps .appInfo.pointsShared0 .stats .effortPoints.unique { visibility: hidden; }
        /* Hide the whole "column" if there's no virtual app (i.e. no shared-libs app). */
        body.viewAppList.noVirtualApp .apps .appInfo.stats .effortPoints.shared,
        body.viewAppList.noVirtualApp .apps .appInfo.stats .effortPoints.unique { display: none; }
        body.viewAppList .apps .appInfo .stats .effortPoints.shared .points,
        body.viewAppList .apps .appInfo .stats .effortPoints.unique .points { color: #8491a8; /* Like normal, but grayed. */ }
        
        body.viewAppList .apps .appInfo .stats .incidentsCount { float: left; margin:0 2ex;}
        body.viewAppList .apps .appInfo .stats .incidentsCount table tr.total td { border-top: 1px solid silver; }
        body.viewAppList .apps .appInfo .stats .incidentsCount .count { text-align: right; padding-right: 1ex; min-width: 7.4ex; }
        body.viewAppList .apps .appInfo .traits { margin-left: 0px; width: 70%;}
        body.viewAppList .apps .appInfo .traits .fileName { padding: 0.0ex 0em 0.2ex; font-size: 18pt; /* color: #008cba; (Default BS link color) */ }
        body.viewAppList .apps .appInfo .traits .techs { }
        
        /* Specifics for virtual apps. */
        body.viewAppList .apps .virtual .appInfo .traits .fileName { color: #477280; }
        
        body.viewAppList .apps .appInfo:first-of-type { border-top: 1px solid gray; }
        
        
        
        
        var TARGET_RUNTIME = JSON.parse('[{"id":"eap","name":"JBoss EAP","description":"Red Hat JBoss Enterprise Application Platform","supported":["Java EE","JBoss Web XML","JMS Connection Factory","JAXB","Enterprise Web Services","CDI","JMS Topic","JSF Page","WS Metadata","Java EE Security","JAX-RS","Clustering EJB","Common Annotations","JAXR","Clustering Web Session","Stateful (SFSB)","SOAP (SAAJ)","Servlet","JPA named queries","EJB XML*","Java EE Batch","JSON-B","JSP Page","Persistence units","Stateless (SLSB)","Message (MDB)","JCA","JAX-WS","Java EE JSON-P","JMS Queue","Java EE Batch API","JMS","Bean Validation","JavaMail","MEJB","JACC","JBoss EJB XML","JTA","EAR","JPA XML*","JPA","EJB","JPA entities"],"unsuitable":["WebSphere Web XML","WebLogic Web XML","WebSphere EJB Ext","JAX-RPC","WebSphere WS Extension","WebSphere WS Binding","RMI","WebSphere EJB"],"neutral":["Spring MVC","JDBC XA datasources","Bouncy Castle","Manifest","Properties","Spring","WSDL","Java Source","JDBC datasources","Swagger","Maven XML","Web XML*","Spring Security"]},{"id":"jws","name":"JWS","description":"Red Hat JBoss Web Server","supported":["JSP Page","Bean Validation","JAXB","Servlet","JSF Page"],"unsuitable":["Java EE","WebSphere Web XML","JBoss Web XML","JMS Connection Factory","Enterprise Web Services","CDI","JMS Topic","WS Metadata","Java EE Security","JAX-RS","Clustering EJB","Common Annotations","JAXR","Clustering Web Session","Stateful (SFSB)","SOAP (SAAJ)","WebSphere EJB Ext","JPA named queries","EJB XML*","JAX-RPC","Java EE Batch","WebSphere EJB","Stateless (SLSB)","Message (MDB)","JCA","JAX-WS","Java EE JSON-P","WebSphere WS Extension","JMS Queue","WebSphere WS Binding","RMI","Java EE Batch API","JMS","JavaMail","MEJB","JACC","JBoss EJB XML","JTA","WebLogic Web XML","EAR","JPA XML*","JPA","EJB","JPA entities"],"neutral":["Java Source","JDBC datasources","Manifest","Maven XML","Properties","Web XML*"]}]');
        
        
        
        
        
        
        
        
        
        
        
        
        
        Migration Toolkit for Applications
        
        
        
        
        
        All Applications
        
        
        
        
        
        
        
        Technologies
        
        
        
        
        
        
        
        About
        
        
        
        
        
        Send Feedback 
        
        
        
        
        
        
        
        var FEEDBACK_JS_ADDED = false;
        var FEEDBACK_FORM_TRIGGER = null;
        
        function displayFeedbackForm() {
        FEEDBACK_FORM_TRIGGER();
        }
        
        window.ATL_JQ_PAGE_PROPS = {
        "triggerFunction": function(showCollectorDialog) {
        FEEDBACK_FORM_TRIGGER = showCollectorDialog;
        }
        };
        
        document.addEventListener("DOMContentLoaded", function(event) {
        jQuery(".jiraFeedbackTrigger").click(function(e) {
        e.preventDefault();
        displayFeedbackForm();
        });
        });
        
        
        
        var script= document.createElement("script");
        script.type= "text/javascript";
        script.src= "reports/resources/js/navbar.js";
        document.body.appendChild(script);
        
        
        
        
        
        
        
        
        
        Application List
        
        
        
        
        
        
        
        
        
        
        
        
        
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"> 
        
        
        
        
        
        
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"> 
        
        
        
        
        
        
        
        
        
        <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
        > 
        
        
        
        
        
        
        
        
        
        
        
        &times;
        Warning! Invalid search regular expression!
        
        
        
        
        Results
        Active filters:
        
        
        <!--
        
        
        Name: nameofthething
        
        
        
        -->
        
        Clear All Filters
        
        
        
        
        
        
        
        
        
        
        
        
        
        Runtime labels legend
        
        Supported
        Partially supported
        Unsuitable
        Neutral
        
        
        
        
        
        
        
        
        
        
        
        
        
        5
        story points
        
        
        
        
        Number of incidents
        
        
        5
        Cloud Mandatory
        
        
        13
        Information
        
        
        18 
        Total 
        
        
        
        
        
        
        
        
        spring-framework-petclinic
        
        
        
        
        Bean Validation 
        
        
        Common Annotations 
        
        
        JAXB 
        
        
        JPA entities 
        
        
        JSP Page 
        
        
        Java Source 
        
        
        Maven XML 
        
        
        Properties 
        
        
        Servlet 
        
        
        Spring Data JPA 
        
        
        Spring JMX 
        
        
        Spring MVC 
        
        
        
        
        
        
        
        $("body").addClass("noVirtualApp");
        
        
        Rule providers execution overview
        |
        FreeMarker methods
        
        
        Page generated: Sep 8, 2023, 1:05:34 PM
        
        
        
        
        $("body.viewAppList .apps .real .appInfo").sortElements(function(a, b){
        return $(a).find(".traits .fileName").first().text().trim().toLowerCase() > $(b).find(".traits .fileName").first().text().trim().toLowerCase() ? 1 : -1;
        });
        
        
        $(document).ready(function(){$('[data-toggle="tooltip"]').tooltip();});
        
        
time=2023-09-08T15:05:57+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/index.html
time=2023-09-08T15:05:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/migration_issues.html
time=2023-09-08T15:05:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/ApplicationDetails_Spring_Framework_Petclinic.html
time=2023-09-08T15:05:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/pom_xml.html
time=2023-09-08T15:05:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/MavenWrapperDownloader_java.html
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:50 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:60 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:55 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:111 Message: CodeSnip: Facts:map[]}
FOUND INCIDENT: {Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} File:MavenWrapperDownloader_java.html Line:78 Message: CodeSnip: Facts:map[]}
time=2023-09-08T15:05:59+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/maven_wrapper_properties.html
time=2023-09-08T15:06:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/wro_properties.html
time=2023-09-08T15:06:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/messages_properties.html
time=2023-09-08T15:06:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/messages_de_properties.html
time=2023-09-08T15:06:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/messages_en_properties.html
time=2023-09-08T15:06:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/messages_es_properties.html
time=2023-09-08T15:06:00+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/session_properties.html
time=2023-09-08T15:06:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/data_access_properties.html
time=2023-09-08T15:06:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/PetclinicInitializer_java.html
time=2023-09-08T15:06:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/InitDataSourceConfig_java.html
time=2023-09-08T15:06:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/RootApplicationContextConfig_java.html
time=2023-09-08T15:06:01+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/SessionConfiguration_java.html
time=2023-09-08T15:06:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Owner_java.html
time=2023-09-08T15:06:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Person_java.html
time=2023-09-08T15:06:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Pet_java.html
time=2023-09-08T15:06:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/PetType_java.html
time=2023-09-08T15:06:02+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Specialty_java.html
time=2023-09-08T15:06:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Vet_java.html
time=2023-09-08T15:06:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Vets_java.html
time=2023-09-08T15:06:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/Visit_java.html
time=2023-09-08T15:06:03+02:00 level=info msg=[binding] |200|  GET /hub/applications/4/bucket/windup/report/reports/CallMonitoringAspect_java.html
WINDUP ANALYSIS OUTPUT FOR "Petclinic legacy cloud-readiness with tagger and hazelcast custom rules":api.Analysis{
    Effort: 5,
    Issues: []api.Issue{
        {
            Category: "cloud-mandatory",
            Description: "Trivial change or 1-1 library swap",
            Effort: 5,
            RuleSet: "",
            Rule: "local-storage-00001",
            Incidents: []api.Incident{
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "50",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "60",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "55",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "111",
                Message: "",
            },
            {
                File: "MavenWrapperDownloader_java.html",
                Line: "78",
                Message: "",
            },
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
    }
}
    analysis_test.go:108: Different effort error. Got 5, expected 12
=== RUN   TestApplicationAnalysis/Apache_Wicket
time=2023-09-08T15:06:03+02:00 level=info msg=[binding] |201|  POST /hub/applications
time=2023-09-08T15:06:03+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-08T15:06:04+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:09+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:14+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:19+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:25+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:30+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:35+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:40+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:45+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:51+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:06:56+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:07:01+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:07:06+02:00 level=info msg=[binding] |200|  GET /hub/tasks/5
time=2023-09-08T15:07:11+02:00 level=info msg=[binding] |401|  GET /hub/tasks/5
    analysis_test.go:94: Analyze Task failed. Details: &{Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} Name: Locator: Priority:0 Variant: Policy: TTL:<nil> Addon: Data:<nil> Application:<nil> State: Image: Pod: Retries:0 Started:<nil> Terminated:<nil> Canceled:false Bucket:<nil> Purged:false Errors:[] Activity:[]}
time=2023-09-08T15:07:12+02:00 level=info msg=[binding] |200|  GET /hub/applications/5/bucket/windup/report/index.html
    analyzer_legacy_windup.go:78: Error report contect check for /windup/report/index.html. Cannot find 5
        story points in Migration Toolkit for Applicationswindow._env = JSON.parse(atob("eyJBVVRIX1JFUVVJUkVEIjoidHJ1ZSIsIktFWUNMT0FLX1JFQUxNIjoibXRhIiwiS0VZQ0xPQUtfQ0xJRU5UX0lEIjoibXRhLXVpIiwiS0VZQ0xPQUtfU0VSVkVSX1VSTCI6Imh0dHBzOi8va2V5Y2xvYWsub3BlbnNoaWZ0LW10YS5zdmM6ODQ0MyIsIlVJX0lOR1JFU1NfUFJPWFlfQk9EWV9TSVpFIjoiNTAwbSIsIlBST0ZJTEUiOiJtdGEiLCJWRVJTSU9OIjoiNi4yLjAiLCJSV1hfU1VQUE9SVEVEIjoiZmFsc2UifQ=="));You need to enable JavaScript to run this app.
time=2023-09-08T15:07:13+02:00 level=info msg=[binding] |200|  GET /hub/applications/5/bucket/windup/report/index.html
time=2023-09-08T15:07:13+02:00 level=info msg=[binding] |200|  GET /hub/applications/5/bucket/windup/report/reports/migration_issues.html
time=2023-09-08T15:07:13+02:00 level=info msg=[binding] |200|  GET /hub/applications/5/bucket/windup/report/reports
WINDUP ANALYSIS OUTPUT FOR "Apache Wicket":api.Analysis{
    Effort: 0,
    Issues: []api.Issue{
    }
}
    analysis_test.go:108: Different effort error. Got 0, expected 5
=== RUN   TestApplicationAnalysis/Seam_booking
time=2023-09-08T15:07:14+02:00 level=info msg=[binding] |401|  POST /hub/applications
    error.go:12: POST /hub/applications failed: 401(Unauthorized)
time=2023-09-08T15:07:14+02:00 level=info msg=[binding] |401|  POST /hub/tasks
    error.go:12: POST /hub/tasks failed: 401(Unauthorized)
time=2023-09-08T15:07:14+02:00 level=info msg=[binding] |401|  GET /hub/tasks/0
    analysis_test.go:94: Analyze Task failed. Details: &{Resource:{ID:0 CreateUser: UpdateUser: CreateTime:0001-01-01 00:00:00 +0000 UTC} Name: Locator: Priority:0 Variant: Policy: TTL:<nil> Addon: Data:<nil> Application:<nil> State: Image: Pod: Retries:0 Started:<nil> Terminated:<nil> Canceled:false Bucket:<nil> Purged:false Errors:[] Activity:[]}
time=2023-09-08T15:07:14+02:00 level=info msg=[binding] |200|  GET /hub/applications/0/bucket/windup/report/index.html
    analyzer_legacy_windup.go:78: Error report contect check for /windup/report/index.html. Cannot find 0
        story points in Migration Toolkit for Applicationswindow._env = JSON.parse(atob("eyJBVVRIX1JFUVVJUkVEIjoidHJ1ZSIsIktFWUNMT0FLX1JFQUxNIjoibXRhIiwiS0VZQ0xPQUtfQ0xJRU5UX0lEIjoibXRhLXVpIiwiS0VZQ0xPQUtfU0VSVkVSX1VSTCI6Imh0dHBzOi8va2V5Y2xvYWsub3BlbnNoaWZ0LW10YS5zdmM6ODQ0MyIsIlVJX0lOR1JFU1NfUFJPWFlfQk9EWV9TSVpFIjoiNTAwbSIsIlBST0ZJTEUiOiJtdGEiLCJWRVJTSU9OIjoiNi4yLjAiLCJSV1hfU1VQUE9SVEVEIjoiZmFsc2UifQ=="));You need to enable JavaScript to run this app.
    analyzer_legacy_windup.go:78: Error report contect check for /windup/report/index.html. Cannot find 3
        Information in Migration Toolkit for Applicationswindow._env = JSON.parse(atob("eyJBVVRIX1JFUVVJUkVEIjoidHJ1ZSIsIktFWUNMT0FLX1JFQUxNIjoibXRhIiwiS0VZQ0xPQUtfQ0xJRU5UX0lEIjoibXRhLXVpIiwiS0VZQ0xPQUtfU0VSVkVSX1VSTCI6Imh0dHBzOi8va2V5Y2xvYWsub3BlbnNoaWZ0LW10YS5zdmM6ODQ0MyIsIlVJX0lOR1JFU1NfUFJPWFlfQk9EWV9TSVpFIjoiNTAwbSIsIlBST0ZJTEUiOiJtdGEiLCJWRVJTSU9OIjoiNi4yLjAiLCJSV1hfU1VQUE9SVEVEIjoiZmFsc2UifQ=="));You need to enable JavaScript to run this app.
time=2023-09-08T15:07:15+02:00 level=info msg=[binding] |200|  GET /hub/applications/0/bucket/windup/report/index.html
time=2023-09-08T15:07:15+02:00 level=info msg=[binding] |200|  GET /hub/applications/0/bucket/windup/report/reports/migration_issues.html
time=2023-09-08T15:07:15+02:00 level=info msg=[binding] |200|  GET /hub/applications/0/bucket/windup/report/reports
WINDUP ANALYSIS OUTPUT FOR "Seam booking":api.Analysis{
    Effort: 0,
    Issues: []api.Issue{
    }
}
```

## TIER2

```
$ TIER2=1 KEEP=1 go test -v -count=1 ./hack/analysis-windup/
time=2023-09-08T15:07:30+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:07:30+02:00 level=info msg=[addon] Addon (adapter) created.
time=2023-09-08T15:07:30+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:07:31+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-08T15:07:31+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:07:31+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
time=2023-09-08T15:07:31+02:00 level=info msg=[binding] Hub RichClient created.
time=2023-09-08T15:07:32+02:00 level=info msg=[binding] |201|  POST /hub/auth/login
=== RUN   TestApplicationAnalysis
=== RUN   TestApplicationAnalysis/Daytrader
time=2023-09-08T15:07:32+02:00 level=info msg=[binding] |201|  POST /hub/applications
time=2023-09-08T15:07:32+02:00 level=info msg=[binding] |201|  POST /hub/tasks
time=2023-09-08T15:07:32+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:07:38+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:07:43+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:07:48+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:07:53+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:07:58+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:04+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:09+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:14+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:19+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:24+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:30+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:35+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:40+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:45+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:50+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:08:56+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:09:01+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:09:06+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:09:11+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:09:16+02:00 level=info msg=[binding] |200|  GET /hub/tasks/6
time=2023-09-08T15:09:17+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/index.html
time=2023-09-08T15:09:22+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/index.html
time=2023-09-08T15:09:23+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/migration_issues.html
time=2023-09-08T15:09:24+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/ApplicationDetails_WAS_Liberty_Sample___Java_EE7_Benchmark_Sample.html
time=2023-09-08T15:09:24+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/pom_xml.html
time=2023-09-08T15:09:24+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/pom_xml.3.html
time=2023-09-08T15:09:24+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/service_properties.html
time=2023-09-08T15:09:24+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/pom_xml.4.html
time=2023-09-08T15:09:25+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/ejb_jar_xml.html
time=2023-09-08T15:09:25+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/pom_xml.1.html
time=2023-09-08T15:09:25+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/build_properties.html
time=2023-09-08T15:09:25+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/web_xml.html
time=2023-09-08T15:09:25+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/daytrader_properties.html
time=2023-09-08T15:09:25+02:00 level=info msg=[binding] |200|  GET /hub/applications/6/bucket/windup/report/reports/build_properties.2.html
WINDUP ANALYSIS OUTPUT FOR "Daytrader":api.Analysis{
    Effort: 0,
    Issues: []api.Issue{
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
        {
            Category: "information",
            Description: "Info",
            Effort: 0,
            RuleSet: "",
            Rule: "",
            Incidents: []api.Incident{
            },
        },
    }
}
```

