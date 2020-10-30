package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//NamespaceValue used to automatically distinct namespace/name string
type NamespaceValue struct {
	Namespace, Name string
}

//UnmarshalFlag Unmarshal flag
func (n *NamespaceValue) UnmarshalFlag(value string) error {
	parts := strings.Split(value, "/")

	if len(parts) != 2 {
		return errors.New("expected two strings separated by a /")
	}
	n.Namespace = parts[0]
	n.Name = parts[1]
	return nil
}

//MarshalFlag Marshals flag
func (n NamespaceValue) MarshalFlag() (string, error) {
	return fmt.Sprintf("%s/%s", n.Namespace, n.Name), nil
}

//LogLevel used to automatically distinct namespace/name string
type LogLevelValue struct {
	LogLevel LogLevel
}

//UnmarshalFlag Unmarshal flag
func (n *LogLevelValue) UnmarshalFlag(value string) error {
	switch value {
	case "trace":
		n.LogLevel = Trace
		return nil
	case "debug":
		n.LogLevel = Debug
		return nil
	case "info":
		n.LogLevel = Info
		return nil
	case "warning":
		n.LogLevel = Warning
		return nil
	case "error":
		n.LogLevel = Error
		return nil
	}

	return fmt.Errorf("value %s not permitted", value)
}

//OSArgs contains arguments that can be sent to controller
type OSArgs struct {
	Version               []bool         `short:"v" long:"version" description:"version"`
	DefaultBackendService NamespaceValue `long:"default-backend-service" default:"" description:"default service to serve 404 page. If not specified HAProxy serves http 400"`
	DefaultCertificate    NamespaceValue `long:"default-ssl-certificate" default:"" description:"secret name of the certificate"`
	ConfigMap             NamespaceValue `long:"configmap" description:"configmap designated for HAProxy" default:""`
	ConfigMapTCPServices  NamespaceValue `long:"configmap-tcp-services" description:"configmap used to define tcp services" default:""`
	ConfigMapErrorfiles   NamespaceValue `long:"configmap-errorfiles" description:"configmap used to define custom error pages associated to HTTP error codes" default:""`
	KubeConfig            string         `long:"kubeconfig" default:"" description:"combined with -e. location of kube config file"`
	NamespaceWhitelist    []string       `long:"namespace-whitelist" description:"whitelisted namespaces"`
	NamespaceBlacklist    []string       `long:"namespace-blacklist" description:"blacklisted namespaces"`
	PprofEnabled          bool           `short:"p" description:"enable pprof over https"`
	CfgDir                string         `long:"config-dir" description:"path to HAProxy configuration directory NOTE: works only with OutOfCluster"`
	OutOfCluster          bool           `short:"e" description:"use as out of cluster controller NOTE: experimental"`
	Test                  bool           `short:"t" description:"simulate running HAProxy"`
	Help                  []bool         `short:"h" long:"help" description:"show this help message"`
	IngressClass          string         `long:"ingress.class" default:"" description:"ingress.class to monitor in multiple controllers environment"`
	PublishService        string         `long:"publish-service" default:"" description:"Takes the form namespace/name. The controller mirrors the address of this service's endpoints to the load-balancer status of all Ingress objects it satisfies"`
	SyncPeriod            time.Duration  `long:"sync-period" default:"5s" description:"Sets the period at which the controller syncs HAProxy configuration file"`
	CacheResyncPeriod     time.Duration  `long:"cache-resync-period" default:"10m" description:"Sets the underlying Shared Informer resync period: resyncing controller with informers cache"`
	LogLevel              LogLevelValue  `long:"log" default:"info" description:"level of log messages you can see"`
	DisableIPV4           bool           `long:"disable-ipv4" description:"toggle to disable the IPv4 protocol from all frontends"`
	IPV4BindAddr          string         `long:"ipv4-bind-address" default:"0.0.0.0" description:"IPv4 address the Ingress Controller listens on (if enabled)"`
	DisableIPV6           bool           `long:"disable-ipv6" description:"toggle to disable the IPv6 protocol from all frontends"`
	IPV6BindAddr          string         `long:"ipv6-bind-address" default:"::" description:"IPv6 address the Ingress Controller listens on (if enabled)"`
	HTTPBindPort          int64          `long:"http-bind-port" default:"80" description:"port to listen on for HTTP traffic"`
	DisableHTTP           bool           `long:"disable-http" description:"toggle to disable the HTTP frontend"`
	HTTPSBindPort         int64          `long:"https-bind-port" default:"443" description:"port to listen on for HTTPS traffic"`
	DisableHTTPS          bool           `long:"disable-https" description:"toggle to disable the HTTPs frontend"`
}
