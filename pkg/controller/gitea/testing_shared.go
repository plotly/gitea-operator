package gitea

import (
	integreatlyv1alpha1 "github.com/plotly/gitea-operator/pkg/apis/integreatly/v1alpha1"
	giteactrl "github.com/plotly/gitea-operator/pkg/controller/gitea"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ExampleNamespace = "example-namespace"

var MockCR = integreatlyv1alpha1.Gitea{
	ObjectMeta: metav1.ObjectMeta{
		Namespace: ExampleNamespace,
	},
	Spec: integreatlyv1alpha1.GiteaSpec{
		Hostname: "gitea.example.com",
	},
}

var Templates = []string{
	giteactrl.GiteaServiceAccountName,
	giteactrl.GiteaConfigName,
	giteactrl.GiteaPgPvcName,
	giteactrl.GiteaReposPvcName,
	giteactrl.GiteaDeploymentName,
	giteactrl.GiteaIngressName,
	giteactrl.GiteaServiceName,
	giteactrl.GiteaPgDeploymentName,
	giteactrl.GiteaPgServiceName,
}
