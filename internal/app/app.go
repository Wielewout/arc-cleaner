package app

import (
	"context"
	"time"

	githubarcv1alpha1 "github.com/actions/actions-runner-controller/apis/actions.github.com/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	actionsgithubcom "github.com/wielewout/arc-cleaner/internal/actions.github.com"
	"github.com/wielewout/arc-cleaner/internal/kubernetes"
	"github.com/wielewout/arc-cleaner/internal/logging"
)

type App struct {
	k8sClient *kubernetes.Client
	namespace string
	period    time.Duration
	dryRun    bool
}

func New(k8sClient *kubernetes.Client, opts ...Option) *App {
	app := &App{
		k8sClient: k8sClient,
		namespace: "default",
		period:    30 * time.Second,
		dryRun:    false,
	}

	for _, opt := range opts {
		opt.apply(app)
	}

	return app
}

func (app App) Start(ctx context.Context) {
	logger := logging.FromContext(ctx)

	ticker := time.NewTicker(app.period)
	logger.Debug("started periodic timer", "period", app.period)

	app.reconcile(ctx)

	for {
		select {
		case <-ticker.C:
			logger.Debug("triggered periodic timer")
			app.reconcile(ctx)
		case <-ctx.Done():
			ticker.Stop()
			logger.Debug("stopped periodic timer")
			return
		}
	}
}

func (app App) reconcile(ctx context.Context) {
	ephemeralRunnerList := app.getEphemeralRunnerList(ctx)
	for _, ephemeralRunner := range ephemeralRunnerList.Items {
		controller := actionsgithubcom.NewEphemeralRunnerReconciler(
			app.k8sClient,
			actionsgithubcom.WithDryRun(app.dryRun),
		)
		_ = controller.Reconcile(ctx, types.NamespacedName{
			Name:      ephemeralRunner.GetName(),
			Namespace: ephemeralRunner.GetNamespace(),
		})
	}
}

func (app App) getEphemeralRunnerList(ctx context.Context) *githubarcv1alpha1.EphemeralRunnerList {
	logger := logging.FromContext(ctx).
		With("namespace", app.namespace)

	ephemeralRunnerList := new(githubarcv1alpha1.EphemeralRunnerList)
	err := app.k8sClient.List(
		ctx,
		ephemeralRunnerList,
		client.InNamespace(app.namespace),
	)

	if err != nil {
		logger.Error("failed to list ephemeral runners", "error", err.Error())
		return ephemeralRunnerList
	}

	logger.Debug("listed ephemeral runners", "length", len(ephemeralRunnerList.Items))
	return ephemeralRunnerList
}
