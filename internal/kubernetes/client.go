package kubernetes

import (
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	githubv1alpha1 "github.com/actions/actions-runner-controller/apis/actions.github.com/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	client.Client
}

func NewClient() (*Client, error) {
	c := &Client{}
	var scheme *runtime.Scheme
	var err error

	scheme, err = getScheme()
	if err != nil {
		slog.Error("failed to get scheme", "error", err.Error())
		return c, err
	}

	if isRunningInCluster() {
		c.Client, err = getClusterClient(scheme)
	} else {
		c.Client, err = getLocalClient(scheme)
	}

	if err != nil {
		return c, err
	}

	return c, nil
}

func getScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()

	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		slog.Debug("failed to add kubernetes client scheme", "error", err.Error())
		return nil, err
	}
	if err := githubv1alpha1.AddToScheme(scheme); err != nil {
		slog.Debug("failed to add github v1alpha1 scheme", "error", err.Error())
		return nil, err
	}

	return scheme, nil
}

func isRunningInCluster() bool {
	_, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount")
	return !errors.Is(err, fs.ErrNotExist)
}

func getClusterClient(scheme *runtime.Scheme) (client.Client, error) {
	slog.Info("running inside cluster")

	slog.Debug("using mounted service account")
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	return client.New(config, client.Options{
		Scheme: scheme,
	})
}

func getLocalClient(scheme *runtime.Scheme) (client.Client, error) {
	slog.Info("running outside cluster")

	kubeConfigPath := getKubeConfigPath()
	slog.Debug("using local kubeconfig", "path", kubeConfigPath)

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return client.New(config, client.Options{
		Scheme: scheme,
	})
}

func getKubeConfigPath() string {
	kubeConfigPath, isKubeConfigSet := os.LookupEnv("KUBECONFIG")

	if !isKubeConfigSet {
		home, err := os.UserHomeDir()
		if err != nil {
			slog.Error("failed to get user home directory", "error", err.Error())
		}
		kubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	return kubeConfigPath
}
