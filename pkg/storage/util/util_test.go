package util

import (
	"fmt"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	configv1 "github.com/openshift/api/config/v1"
	imageregistryv1 "github.com/openshift/cluster-image-registry-operator/pkg/apis/imageregistry/v1"
	regopclient "github.com/openshift/cluster-image-registry-operator/pkg/client"
)

var (
	infrastructureName = "jsmith-xhv4"
	prefix             = fmt.Sprintf("%s-%s", infrastructureName, imageregistryv1.ImageRegistryName)
)

type MockInfrastructureLister struct {
}

func (m MockInfrastructureLister) Get(name string) (*configv1.Infrastructure, error) {
	return &configv1.Infrastructure{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster",
		},
		Status: configv1.InfrastructureStatus{
			InfrastructureName: infrastructureName,
			PlatformStatus:     &configv1.PlatformStatus{},
		},
	}, nil
}

func (m MockInfrastructureLister) List(selector labels.Selector) ([]*configv1.Infrastructure, error) {
	var list []*configv1.Infrastructure
	return list, nil
}

func TestGenerateStorageName(t *testing.T) {
	l := regopclient.Listers{Infrastructures: MockInfrastructureLister{}}
	tests := [][]string{
		// test with no additionals
		{},
		// test one additional
		{"test1"},
		// test two additionals
		{"test1", "test2"},
		// test really long additionals
		{"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"},
	}

	for _, test := range tests {
		n, err := GenerateStorageName(&l, test...)
		if err != nil {
			t.Errorf("%v", err)
		}
		wantedPrefix := fmt.Sprintf("%s-%s", prefix, strings.Join(test, "-"))
		if len(wantedPrefix) > 62 {
			wantedPrefix = wantedPrefix[0:62]
		}
		if !strings.HasPrefix(n, wantedPrefix) {
			t.Errorf("name should have the prefix %s, but was %s instead", wantedPrefix, n)
		}
		if len(n) != 62 {
			t.Errorf("name should be exactly 62 characters long, but was %d instead: %s", len(n), n)
		}
	}
}
