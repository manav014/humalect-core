package constants

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1 "k8s.io/api/apps/v1"

	corev1 "k8s.io/api/core/v1"

	networkingv1 "k8s.io/api/networking/v1"
)

const (
	TempDirectoryName = "temp"
)

type DeploymentYamlManifestType struct {
	Metadata metav1.ObjectMeta     `json:"metadata"`
	Spec     appsv1.DeploymentSpec `json:"spec"`
}

type ServiceYamlManifestType struct {
	Metadata metav1.ObjectMeta  `json:"metadata"`
	Spec     corev1.ServiceSpec `json:"spec"`
}
type IngressYamlManifestType struct {
	Metadata metav1.ObjectMeta        `json:"metadata"`
	Spec     networkingv1.IngressSpec `json:"spec"`
}

type EcrCredentials struct {
	AWSTempAccessKey string `json:"awsTempAccessKey,omitempty"`
	AWSTempSecretKey string `json:"awsTempSecretKey,omitempty"`
}

type DockerHubCredentials struct {
	Username   string `json:"username,omitempty"`
	SecretName string `json:"secretName,omitempty"`
}

type AcrCredentials struct {
	AzureManagementScopeToken string `json:"azureManagementScopeToken,omitempty"`
	AzureAcrRegistryName      string `json:"azureAcrRegistryName,omitempty"`
	AzureSubscriptionId       string `json:"azureSubscriptionId,omitempty"`
	AzureResourceGroupName    string `json:"azureResourceGroupName,omitempty"`
}

type AwsSecretCredentials struct {
	AWSTempAccessKey string `json:"awsTempAccessKey,omitempty"`
	AWSTempSecretKey string `json:"awsTempSecretKey,omitempty"`
	AWSRegion        string `json:"awsRegion,omitempty"`
}

type AzureSecretCredentials struct {
	AzureVaultToken string `json:"azureVaulttoken,omitempty"`
	AzureVaultName  string `json:"azureVaultName,omitempty"`
}

type ParamsConfig struct {
	SecretsProvider           string
	EcrCredentials            string
	DockerHubCredentials      string
	AcrCredentials            string
	AwsSecretCredentials      string
	AzureSecretCredentials    string
	RegistryProvider          string
	CloudProvider             string
	AzureManagementScopeToken string
	SourceCodeRepositoryName  string
	SourceCodeProvider        string
	SourceCodeToken           string
	SourceCodeOrgName         string
	CommitId                  string
	DockerManifest            string
	ArtifactsRepositoryName   string
	AwsEcrRegistryUrl         string
	K8sAppName                string
	AzureAcrRegistryName      string
	AwsEcrUserName            string
	AzureSubscriptionId       string
	AzureResourceGroupName    string
	UseDockerFromCodeFlag     bool
	DeploymentYamlManifest    string
	IngressYamlManifest       string
	ServiceYamlManifest       string
	ManagedBy                 string `default:"Humalect"`
	CloudRegion               string
	K8sResourcesIdentifier    string
	AzureVaultToken           string
	SecretManagerName         string
	Namespace                 string `default:"default"`
	AzureVaultName            string
	DeploymentId              string
	WebhookEndpoint           string
	WebhookData               string
}

const (
	RegistryIdAWS                     = "ecr"
	RegistryIdAzure                   = "acr"
	RegistryIdDockerhub               = "dockerhub"
	CloudIdCivo                       = "civo"
	CloudIdAzure                      = "azure"
	CloudIdAWS                        = "aws"
	KanikoWorkspaceName               = "workspace"
	SourceGithub                      = "github"
	SourceGitlab                      = "gitlab"
	SourceBitbucket                   = "bitbucket"
	DockerConfigMountPath             = "/docker-config"
	CreatedKanikoJob                  = "CREATED_KANIKO_JOB"
	WebhookTypeDeploymentStatusUpdate = "TYPE_DEPLOYMENT_STATUS_UPDATE"
	DeploymentFailed                  = "DEPLOYMENT_FAILED"
	KanikoJobExecuted                 = "KANIKO_JOB_EXECUTED"
	CreatedApplicationCrd             = "CREATED_APPLICATION_CRD"
)