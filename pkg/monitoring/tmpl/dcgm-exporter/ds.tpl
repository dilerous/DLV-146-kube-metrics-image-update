apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: dcgm-exporter
  namespace: {{ ns . }}
  labels:
    app: "dcgm-exporter"
spec:
  selector:
    matchLabels:
      app: "dcgm-exporter"
  template:
    metadata:
      labels:
        app: "dcgm-exporter"
    spec:
      serviceAccountName: dcgm-exporter
      volumes:
        - name: "pod-gpu-resources"
          hostPath:
            path: "/var/lib/kubelet/pod-resources"
      tolerations:
        - operator: Exists
      nodeSelector:
        accelerator: nvidia
      containers:
        - name: exporter
          securityContext:
            capabilities:
              add:
                - SYS_ADMIN
            runAsNonRoot: false
            runAsUser: 0
          image: {{ .Spec.Monitoring.DcgmExporter.Image }}
          imagePullPolicy: "IfNotPresent"
          args:
            - -f
            - /etc/dcgm-exporter/dcp-metrics-included.csv
          env:
            - name: "DCGM_EXPORTER_KUBERNETES"
              value: "true"
            - name: "DCGM_EXPORTER_LISTEN"
              value: ":9400"
          ports:
            - name: "metrics"
              containerPort: 9400
          volumeMounts:
            - name: "pod-gpu-resources"
              readOnly: true
              mountPath: "/var/lib/kubelet/pod-resources"
          livenessProbe:
            httpGet:
              path: /health
              port: 9400
            initialDelaySeconds: 5
            periodSeconds: 5
          readinessProbe:
            httpGet:
              path: /health
              port: 9400
            initialDelaySeconds: 5