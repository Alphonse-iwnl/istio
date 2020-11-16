package model

const (

	// EnvoyJwtFilterName is the name of the Envoy JWT filter. This should be the same as the name defined
	// in https://github.com/envoyproxy/envoy/blob/v1.9.1/source/extensions/filters/http/well_known_names.h#L48
	EnvoyJwtFilterName = "envoy.filters.http.jwt_authn"

	// AuthnFilterName is the name for the Istio AuthN filter. This should be the same
	// as the name defined in
	// https://github.com/istio/proxy/blob/master/src/envoy/http/authn/http_filter_factory.cc#L30
	//// EnvoyJwtFilterName is the name of the Envoy JWT filter. This should be the same as the name defined
	//// in https://github.com/envoyproxy/envoy/blob/v1.9.1/source/extensions/filters/http/well_known_names.h#L48
	//EnvoyJwtFilterName = "envoy.filters.http.jwt_authn"

	//todo: could not find a name for it????

	HTTPMetadataExchangeFilterName = "istio_http_metadataexchange"
	//AuthnFilterName = "istio_authn"

	// KubernetesSecretType is the name of a SDS secret stored in Kubernetes
	KubernetesSecretType    = "kubernetes"
	KubernetesSecretTypeURI = KubernetesSecretType + "://"
)

