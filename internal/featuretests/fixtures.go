package featuretests

import (
	"fmt"

	ingressroutev1 "github.com/projectcontour/contour/apis/contour/v1beta1"
	projcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/tools/cache"
)

// nameOf returns the namespace-qualified name of the give Kubernetes object.
func nameOf(obj interface{}) string {
	name, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		panic(fmt.Sprintf("failed to generate object name: %s", err))
	}

	return name
}

type Fixtures struct {
	secrets       map[string]*v1.Secret
	services      map[string]*v1.Service
	ingresses     map[string]*v1beta1.Ingress
	ingressRoutes map[string]*ingressroutev1.IngressRoute
	httpProxies   map[string]*projcontour.HTTPProxy
}

func (f *Fixtures) Secret(s string) *v1.Secret {
	return f.secrets[s]
}

func (f *Fixtures) Service(s string) *v1.Service {
	return f.services[s]
}

func (f *Fixtures) Ingress(s string) *v1beta1.Ingress {
	return f.ingresses[s]
}

func (f *Fixtures) IngressRoute(s string) *ingressroutev1.IngressRoute {
	return f.ingressRoutes[s]
}

func (f *Fixtures) HTTPProxy(s string) *projcontour.HTTPProxy {
	return f.httpProxies[s]
}

func (f *Fixtures) Add(obj interface{}) {
	name := nameOf(obj)

	switch obj := obj.(type) {
	case *v1.Secret:
		if f.Secret(name) != nil {
			panic(fmt.Sprintf("%T named '%s' already exists", obj, name))
		}
		f.secrets[name] = obj
	case *v1.Service:
		if f.Service(name) != nil {
			panic(fmt.Sprintf("%T named '%s' already exists", obj, name))
		}
		f.services[name] = obj
	case *v1beta1.Ingress:
		if f.Ingress(name) != nil {
			panic(fmt.Sprintf("%T named '%s' already exists", obj, name))
		}
		f.ingresses[name] = obj
	case *ingressroutev1.IngressRoute:
		if f.IngressRoute(name) != nil {
			panic(fmt.Sprintf("%T named '%s' already exists", obj, name))
		}
		f.ingressRoutes[name] = obj
	case *projcontour.HTTPProxy:
		if f.HTTPProxy(name) != nil {
			panic(fmt.Sprintf("%T named '%s' already exists", obj, name))
		}
		f.httpProxies[name] = obj
	default:
		panic(fmt.Sprintf("attempt to add unsupported %T fixture '%s'", obj, name))
	}
}

func NewFixtures() *Fixtures {
	f := Fixtures{
		secrets:       map[string]*v1.Secret{},
		services:      map[string]*v1.Service{},
		ingresses:     map[string]*v1beta1.Ingress{},
		ingressRoutes: map[string]*ingressroutev1.IngressRoute{},
		httpProxies:   map[string]*projcontour.HTTPProxy{},
	}

	f.Add(&v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kuard",
			Namespace: "default",
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{{
				Protocol:   "TCP",
				Port:       8080,
				TargetPort: intstr.FromInt(8080),
			}},
		},
	})

	return &f
}
