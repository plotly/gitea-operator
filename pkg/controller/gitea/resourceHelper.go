package gitea

import (
	yaml "github.com/ghodss/yaml"
	integreatlyv1alpha1 "github.com/plotly/gitea-operator/pkg/apis/integreatly/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceHelper struct {
	templateHelper *GiteaTemplateHelper
	cr             *integreatlyv1alpha1.Gitea
}

func newResourceHelper(cr *integreatlyv1alpha1.Gitea) *ResourceHelper {
	return &ResourceHelper{
		templateHelper: newTemplateHelper(cr),
		cr:             cr,
	}
}

func (r *ResourceHelper) createResource(template string) (client.Object, error) {
	tpl, err := r.templateHelper.loadTemplate(template)
	if err != nil {
		return nil, err
	}

	resource := unstructured.Unstructured{}
	err = yaml.Unmarshal(tpl, &resource)

	if err != nil {
		return nil, err
	}

	return &resource, nil
}
