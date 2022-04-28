package gitea

const (
	GiteaImage              = "quay.io/plotly/gitea-gitea"
	GiteaVersion            = "1.15.5-rootless-rootless"
	GiteaConfigName         = "gitea-config"
	GiteaDeploymentName     = "gitea"
	GiteaIngressName        = "gitea-ingress"
	GiteaPgDeploymentName   = "postgres"
	GiteaPgPvcName          = "gitea-postgres-pvc"
	GiteaPgServiceName      = "gitea-postgresql"
	GiteaReposPvcName       = "gitea-repos"
	GiteaServiceAccountName = "gitea-service-account"
	GiteaServiceName        = "gitea-http"
	GiteaServiceSshName     = "gitea-ssh"          // added by TB
	GiteaInitSecretName     = "gitea-init"         // added by TB
	GiteaAdminSecretName    = "gitea-admin-secret" // added by TB

)
