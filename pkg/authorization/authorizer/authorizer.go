package authorizer

import "k8s.io/apiserver/pkg/authentication/user"

var (
	_ Attributes = (*AttributesRecord)(nil)
)

// AttributesRecord implements Attributes interface.
type AttributesRecord struct {
	User            user.Info
	Verb            string
	APIGroup        string
	APIVersion      string
	Resource        string
	Subresource     string
	Name            string
	ResourceRequest bool
	Path            string
}

func (a *AttributesRecord) GetVerb() string {
	return a.Verb
}

func (a *AttributesRecord) GetResource() string {
	return a.Resource
}

func (a *AttributesRecord) GetSubresource() string {
	return a.Subresource
}

func (a *AttributesRecord) GetName() string {
	return a.Name
}

func (a *AttributesRecord) GetAPIGroup() string {
	return a.APIGroup
}

func (a *AttributesRecord) GetAPIVersion() string {
	return a.APIVersion
}

func (a *AttributesRecord) IsResourceRequest() bool {
	return a.ResourceRequest
}

func (a *AttributesRecord) GetPath() string {
	return a.Path
}

func (a *AttributesRecord) GetUser() user.Info {
	return a.User
}
