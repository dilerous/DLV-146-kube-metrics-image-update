apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Spec.ControlPlane.Hyper.SvcName }}
  namespace: {{ ns . }}
  labels:
    app: {{ .Spec.ControlPlane.Hyper.SvcName }}
spec:
  replicas: {{ .Spec.ControlPlane.Hyper.Replicas }}
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 25%
      maxSurge: 1
  selector:
    matchLabels:
      app: {{ .Spec.ControlPlane.Hyper.SvcName }}
  template:
    metadata:
      labels:
        app: {{ .Spec.ControlPlane.Hyper.SvcName }}
    spec:
      serviceAccountName: {{ .Spec.ControlPlane.Rbac.ServiceAccountName }}
      {{- if eq .Spec.ControlPlane.Tenancy.Enabled "true" }}
      nodeSelector:
        {{ .Spec.ControlPlane.Tenancy.Key }}: "{{ .Spec.ControlPlane.Tenancy.Value }}"
      {{- end }}
      tolerations:
        - key: "{{ .Spec.ControlPlane.Tenancy.Key }}"
          operator: "Equal"
          value: "{{ .Spec.ControlPlane.Tenancy.Value }}"
          effect: "NoSchedule"
      containers:
        - image: {{ .Spec.ControlPlane.Hyper.Image }}
          name: {{ .Spec.ControlPlane.Hyper.SvcName }}
          envFrom:
            - configMapRef:
                name: cp-base-config
            - configMapRef:
                name: cp-networking-config
            - secretRef:
                name: cp-base-secret
            - secretRef:
                name: cp-ldap
            - secretRef:
                name: cp-object-storage
          ports:
            - containerPort: {{ .Spec.ControlPlane.Hyper.Port }}
          readinessProbe:
            failureThreshold: 3
            httpGet:
              path: "/?key={{.Spec.ControlPlane.Hyper.Token}}"
              port: {{.Spec.ControlPlane.Hyper.Port}}
              scheme: HTTP
            initialDelaySeconds: 20
            successThreshold: 1
            periodSeconds: {{ .Spec.ControlPlane.Hyper.ReadinessPeriodSeconds }}
            timeoutSeconds: {{ .Spec.ControlPlane.Hyper.ReadinessTimeoutSeconds }}
          resources:
            requests:
              cpu: {{.Spec.ControlPlane.Hyper.CPURequest}}
              memory: {{.Spec.ControlPlane.Hyper.MemoryRequest}}
            limits:
              cpu: {{ .Spec.ControlPlane.Hyper.CPULimit }}
              memory: {{ .Spec.ControlPlane.Hyper.MemoryLimit }}