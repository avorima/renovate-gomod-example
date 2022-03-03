package main

import (
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "sigs.k8s.io/controller-runtime/pkg/config"
)

func main() {
}
