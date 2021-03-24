apiVersion: apps/v1
kind: Deployment
metadata:
  name: cnvrg-prometheus-operator
  namespace: {{ ns . }}
  labels:
    app: cnvrg-prometheus-operator
    version: v0.44.1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cnvrg-prometheus-operator
      version: v0.44.1
  template:
    metadata:
      labels:
        app: cnvrg-prometheus-operator
        version: v0.44.1
    spec:
      serviceAccountName: cnvrg-prometheus-operator
      securityContext:
        runAsNonRoot: true
        runAsUser: 65534
      containers:
      - args:
        - --kubelet-service=kube-system/kubelet
        - --prometheus-config-reloader={{ .Spec.Monitoring.PrometheusOperator.Images.PrometheusConfigReloaderImage }}
        image: {{ .Spec.Monitoring.PrometheusOperator.Images.OperatorImage }}
        name: prometheus-operator
        ports:
        - containerPort: 8080
          name: http
        resources:
          limits:
            cpu: 200m
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 100Mi
        securityContext:
          allowPrivilegeEscalation: false
      - args:
        - --logtostderr
        - --secure-listen-address=:8443
        - --tls-cipher-suites=TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305
        - --upstream=http://127.0.0.1:8080/
        image: {{ .Spec.Monitoring.PrometheusOperator.Images.KubeRbacProxyImage }}
        name: kube-rbac-proxy
        ports:
        - containerPort: 8443
          name: https
        securityContext:
          runAsGroup: 65532
          runAsNonRoot: true
          runAsUser: 65532

