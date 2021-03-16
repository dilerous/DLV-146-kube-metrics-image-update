apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: {{ .Spec.Logging.Kibana.SvcName }}
  namespace: {{ .Namespace }}
spec:
  hosts:
    - "{{.Spec.Logging.Kibana.SvcName}}.{{ .Spec.ClusterDomain }}"
  gateways:
  - {{ .Spec.Ingress.IstioGwName }}
  http:
  - retries:
      attempts: {{ .Spec.Ingress.RetriesAttempts }}
      perTryTimeout: {{ .Spec.Ingress.PerTryTimeout }}
    timeout: {{ .Spec.Ingress.Timeout }}
    route:
    - destination:
        port:
          number: {{.Spec.Logging.Kibana.Port}}
        host: "{{ .Spec.Logging.Kibana.SvcName }}.{{ .Namespace }}.svc.cluster.local"