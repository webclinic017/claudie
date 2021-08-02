apiVersion: kubeone.io/v1beta1
kind: KubeOneCluster
name: cluster

versions:
  kubernetes: '{{ .Kubernetes }}'

clusterNetwork:
  cni:
    external: {}

cloudProvider:
  none: {}
  external: false

addons:
  enable: true
  # In case when the relative path is provided, the path is relative
  # to the KubeOne configuration file.
  path: "addons"

apiEndpoint:
  host: '{{ .ApiEndpoint }}'
  port: 6443

controlPlane:
  hosts:
{{- $privateKey := "./private.pem" }}
{{- range $value := .Nodes }}
{{- if eq $value.IsControl true}}
  - publicAddress: '{{ $value.Public }}'
    privateAddress: '{{ $value.Private }}'
    sshPrivateKeyFile: '{{ $privateKey }}'
{{- end}}
{{- end}}

staticWorkers:
  hosts:
{{- range $value := .Nodes }}
{{- if eq $value.IsControl false}}
  - publicAddress: '{{ $value.Public }}'
    privateAddress: '{{ $value.Private }}'
    sshPrivateKeyFile: '{{ $privateKey }}'
{{- end}}
{{- end}}

machineController:
  deploy: false