package main

import (
	"github.com/go-logr/zapr"
	"github.com/vrischmann/envconfig"
	"github.com/zhangtbj/KubeOpsAI/internal/engine/controller"
	"github.com/zhangtbj/KubeOpsAI/internal/logger"
	"github.com/zhangtbj/KubeOpsAI/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"log"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

// Config holds application related configuration
type Config struct {
	// LeaderElectionNamespace determines the namespace in which the leader election configmap will be created.
	LeaderElectionNamespace string `envconfig:"optional"`
	// GraphQLAddr is the TCP address the GraphQL endpoint binds to.
	GraphQLAddr string `envconfig:"default=:8080"`
	// MetricsAddr is the TCP address the metric endpoint binds to.
	MetricsAddr string `envconfig:"default=:8081"`
	// HealthzAddr is the TCP address the health probes endpoint binds to.
	HealthzAddr string `envconfig:"default=:8082"`
	// EnableLeaderElection for controller manager. Enabling this will ensure there is only one active controller manager.
	EnableLeaderElection bool `envconfig:"default=false"`
	// MaxConcurrentReconciles is the maximum number of concurrent Reconciles which can be run.
	MaxConcurrentReconciles int `envconfig:"default=1"`

	Logger logger.Config
}

func main() {
	// init configuration
	var cfg Config
	err := envconfig.InitWithPrefix(&cfg, "APP")
	exitOnError(err, "while loading configuration")

	logger, err := logger.New(cfg.Logger)
	exitOnError(err, "while creating zap logger")

	// setup controller
	ctrl.SetLogger(zapr.NewLogger(logger))

	err = clientgoscheme.AddToScheme(scheme)
	exitOnError(err, "while adding k8s scheme")
	err = v1alpha1.AddToScheme(scheme)
	exitOnError(err, "while adding core Question scheme")

	k8sCfg := ctrl.GetConfigOrDie()
	mgr, err := ctrl.NewManager(k8sCfg, ctrl.Options{
		Scheme:                  scheme,
		LeaderElection:          cfg.EnableLeaderElection,
		LeaderElectionNamespace: cfg.LeaderElectionNamespace,
		LeaderElectionID:        "152f0254.kubeopsai.ai",
		MetricsBindAddress:      cfg.MetricsAddr,
		HealthProbeBindAddress:  cfg.HealthzAddr,
	})

	if err = (&controller.QuestionReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Engine"),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ENgine")
		os.Exit(1)
	}

	// Start the manager
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func exitOnError(err error, context string) {
	if err != nil {
		log.Fatalf("%s: %v", context, err)
	}
}
