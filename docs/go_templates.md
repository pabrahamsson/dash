# Golang Template Processor

This project supports [Golang Templates](https://golang.org/pkg/text/template/).

The spec for Go Templates as a resource type looks like:

```
- name: <Resource name>
  goTemplate:
    template: <path to file> # Required
    params: # Optional
      <key>: <value>
      ...
```

## Processing Modes

Go Templates can currently be processed in one mode only, single template with single parameter set.

## Single Template, Single Parameter Set

When `.goTemplate.template` points to a file, we will process that template, passing the parameters to it.

```yaml
- name: <Resource name>
  goTemplate:
    template: templates/app-service.tpl
    params:
      name: dash
      ports:
        - name: http
          port: 8080
          targetPort: 80
        - name: https
          port: 8443
          targetPort: 443
```
Below is the go template file `templates/app-service.tpl`
```
---
apiVersion: v1
kind: Service
metadata:
  name: {{ .name }}-service
spec:
  ports: {{ range .ports }}
    - name: {{ .portName }}
      protocol: TCP
      port: {{ .port }}
      targetPort: {{ .targetPort }} {{ end }}
  selector:
    app: {{ .name }}
```
This will generate a the `dash-service` below
```
---
apiVersion: v1
kind: Service
metadata:
  name: dash-service
spec:
  ports: 
    - name: http
      protocol: TCP
      port: 8080
      targetPort: 80 
    - name: https
      protocol: TCP
      port: 8443
      targetPort: 443 
  selector:
    app: dash
```
