---
apiVersion: v1
kind: Service
metadata:
  name: {{.name }}
spec:
  ports: {{ range .ports}}
    - name: {{.portName}}
      protocol: TCP
      port: {{.port}}
      targetPort: {{.targetPort}} {{ end }}
  selector:
    app: nginx
