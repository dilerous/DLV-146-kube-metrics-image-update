apiVersion: v1
kind: ConfigMap
metadata:
  name: redis-conf
  namespace: {{ ns . }}
data:
  redis.conf: |
    dir /data/
    appendonly "{{ .Spec.ControlPlan.Redis.Appendonly }}"
    appendfilename "appendonly.aof"
    appendfsync everysec
    auto-aof-rewrite-percentage 100
    auto-aof-rewrite-min-size 128mb