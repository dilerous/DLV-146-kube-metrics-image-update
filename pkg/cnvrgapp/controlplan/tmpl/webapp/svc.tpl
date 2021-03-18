apiVersion: v1
kind: Service
metadata:
  name: {{ .Spec.ControlPlan.WebApp.SvcName }}
  namespace: {{ .Namespace }}
  labels:
    app: {{ .Spec.ControlPlan.WebApp.SvcName }}
spec:
  {{- if eq .Spec.Ingress.IngressType "nodeport" }}
  type: NodePort
  {{- end }}
  ports:
  - port: {{.Spec.ControlPlan.WebApp.Port}}
    {{- if eq .Spec.Ingress.IngressType "nodeport" }}
    nodePort: {{ .Spec.ControlPlan.WebApp.NodePort }}
    {{- end }}
  selector:
    app: {{ .Spec.ControlPlan.WebApp.SvcName }}