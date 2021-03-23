apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: {{ .Spec.Monitoring.Prometheus.SvcName }}
  namespace: {{ ns . }}
spec:
  hosts:
  - "{{ .Spec.Monitoring.Prometheus.SvcName }}.{{ .Spec.ClusterDomain }}"
  gateways:
  - {{ .Spec.Networking.Ingress.IstioGwName }}
  http:
  - retries:
      attempts: {{ .Spec.Networking.Ingress.RetriesAttempts }}
      perTryTimeout: {{ .Spec.Networking.Ingress.PerTryTimeout }}
    timeout: {{ .Spec.Networking.Ingress.Timeout }}
    route:
    - destination:
        port:
          number: {{ .Spec.Monitoring.Prometheus.Port }}
        host: "{{ .Spec.Monitoring.Prometheus.SvcName }}.{{ ns . }}.svc.cluster.local"