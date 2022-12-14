package split_yaml

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Document struct {
	metav1.TypeMeta
	Metadata metav1.ObjectMeta `yaml:"metadata"`
}
