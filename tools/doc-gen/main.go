package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	// auditingv1 "github.com/kubeclipper/kubeclipper/pkg/apis/auditing/v1"
	// "github.com/kubeclipper/kubeclipper/pkg/apis/proxy"
	"github.com/kubeclipper/kubeclipper/pkg/logger"

	// "github.com/kubeclipper/kubeclipper/pkg/apis/oauth"

	// iamv1 "github.com/kubeclipper/kubeclipper/pkg/apis/iam/v1"

	corev1 "github.com/kubeclipper/kubeclipper/pkg/apis/core/v1"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/version"

	// configv1 "github.com/kubeclipper/kubeclipper/pkg/apis/config/v1"
)

var output string

func init() {
	flag.StringVar(&output, "output", "./docs/openapi-spec/swagger.json", "--output=./api.json")
}

func main() {
	flag.Parse()
	// 生成api文档
	swaggerSpec := generateSwaggerJSON()

	// 格式校验
	err := validateSpec(swaggerSpec)
	if err != nil {
		logger.Warn("Swagger specification has errors")
	}
}

func validateSpec(apiSpec []byte) error {
	// doc, err := loads.Spec(output, loads.WithDocLoader(loads.JSONDoc))
	doc, err := loads.Analyzed(apiSpec, "2.0", loads.WithDocLoader(loads.JSONDoc))
	if err != nil {
		return err
	}
	validator := validate.NewSpecValidator(doc.Schema(), strfmt.Default)
	validator.SetContinueOnErrors(true)
	result, _ := validator.Validate(doc)
	if result.HasWarnings() {
		log.Printf("See warnings below:\n")
		for _, desc := range result.Warnings {
			log.Printf("- WARNING: %s\n", desc.Error())
		}
	}
	if result.HasErrors() {
		str := fmt.Sprintf("The swagger spec is invalid against swagger specification %s.\nSee errors below:\n", doc.Version())
		for _, desc := range result.Errors {
			str += fmt.Sprintf("- %s\n", desc.Error())
		}
		log.Println(str)
		return errors.New(str)
	}
	return nil
}

func generateSwaggerJSON() []byte {
	container := restful.NewContainer()
	urlruntime.Must(corev1.AddToContainer(container, nil, nil, nil, nil, nil, nil, nil, nil, nil))
	// urlruntime.Must(iamv1.AddToContainer(container, nil, nil, nil))
	// urlruntime.Must(configv1.AddToContainer(container, nil, nil))
	// urlruntime.Must(oauth.AddToContainer(container, nil, nil, nil, nil, nil, nil, nil))
	// urlruntime.Must(proxy.AddToContainer(container, nil))
	// urlruntime.Must(auditingv1.AddToContainer(container, nil))

	config := restfulspec.Config{
		WebServices:                   container.RegisteredWebServices(),
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}

	swagger := restfulspec.BuildSwagger(config)
	// swagger.Info.Extensions = make(spec.Extensions)
	// swagger.Info.Extensions.Add("x-tagGroups", []struct {
	// 	Name string   `json:"name"`
	// 	Tags []string `json:"tags"`
	// }{
	// 	// {
	// 	//	Name: "Cluster",
	// 	//	Tags: []string{
	// 	//		constants.CoreClusterTag,
	// 	//	},
	// 	// },
	// })

	data, _ := json.MarshalIndent(swagger, "", "  ")
	err := os.WriteFile(output, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("successfully written to %s", output)

	return data
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "KubeClipper",
			Description: "KubeClipper OpenAPI",
			Version:     version.Get().GitVersion,
			Contact: &spec.ContactInfo{
				// TODO: add url and email
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "KubeClipper",
					URL:   "https://github.com/kubeclipper-labs/kubeclipper",
					Email: "",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "Apache 2.0",
					URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
				},
			},
		},
	}

	// setup security definitions
	swo.SecurityDefinitions = map[string]*spec.SecurityScheme{
		"jwt": spec.APIKeyAuth("Authorization", "header"),
	}
	swo.Security = []map[string][]string{{"jwt": []string{}}}
}

// func apiTree(container *restful.Container) {
//	buf := bytes.NewBufferString("\n")
//	for _, ws := range container.RegisteredWebServices() {
//		for _, route := range ws.Routes() {
//			buf.WriteString(fmt.Sprintf("%s %s\n", route.Method, route.Path))
//		}
//	}
//	log.Println(buf.String())
// }
