{{/*
Expand the name of the chart.
*/}}
{{- define "typhoon.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "typhoon.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "typhoon.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "typhoon.labels" -}}
helm.sh/chart: {{ include "typhoon.chart" . }}
{{ include "typhoon.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: typhoon
{{- end }}

{{/*
Selector labels
*/}}
{{- define "typhoon.selectorLabels" -}}
app.kubernetes.io/name: {{ include "typhoon.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "typhoon.serviceAccountName" -}}
{{- if .Values.controller.serviceAccount.create }}
{{- default (include "typhoon.fullname" .) .Values.controller.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Controller Service Account Name
*/}}
{{- define "typhoon.controller.serviceAccountName" -}}
{{- $name := include "typhoon.serviceAccountName" . }}
{{- printf "%s-%s" $name "controller" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Controller FQDN
*/}}
{{- define "typhoon.controller.fullname" -}}
{{- $name := include "typhoon.fullname" . }}
{{- printf "%s-%s" $name "controller" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Controller labels
*/}}
{{- define "typhoon.controller.labels" -}}
{{ include "typhoon.labels" . }}
app: typhoon
{{- end }}

{{/*
Controller Selector labels
*/}}
{{- define "typhoon.controller.selectorLabels" -}}
{{ include "typhoon.selectorLabels" . }}
app: typhoon
{{- end }}

{{/*
Webhook Service Account Name
*/}}
{{- define "typhoon.webhook.serviceAccountName" -}}
{{- $name := include "typhoon.serviceAccountName" . }}
{{- printf "%s-%s" $name "webhook" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Webhook FQDN
*/}}
{{- define "typhoon.webhook.fullname" -}}
{{- $name := include "typhoon.fullname" . }}
{{- printf "%s-%s" $name "webhook" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Webhook labels
*/}}
{{- define "typhoon.webhook.labels" -}}
{{ include "typhoon.labels" . }}
app: typhoon-webhook
{{- end }}

{{/*
Webhook Selector labels
*/}}
{{- define "typhoon.webhook.selectorLabels" -}}
{{ include "typhoon.selectorLabels" . }}
app: typhoon-webhook
{{- end }}