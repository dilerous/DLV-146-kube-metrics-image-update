
apiVersion: v1
kind: Service
metadata:
  name: {{ .Spec.Logging.Kibana.SvcName }}
  namespace: {{ .Namespace }}
  labels:
    app: {{ .Spec.Logging.Kibana.SvcName }}
spec:
  {{- if eq .Spec.Ingress.IngressType "nodeport" }}
  type: NodePort
  {{- end }}
  selector:
    app: {{ .Spec.Logging.Kibana.SvcName }}
  ports:
    - port: {{ .Spec.Logging.Kibana.Port }}
      protocol: TCP
      {{- if eq .Spec.Ingress.IngressType "nodeport" }}
      nodePort: {{ .Spec.Logging.Kibana.NodePort }}
      {{- end }}