package emitter

import (
	"fmt"

	logs "github.com/thedevflex/kubi8al-webhook/utils/logger"
)

type PackagePublishedEvent struct {
	Event   string  `json:"event"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Action       string       `json:"action"`
	PackageData  Package      `json:"package"`
	Repository   Repository   `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       Sender       `json:"sender"`
}

type Organization struct {
	AvatarURL string `json:"avatar_url"`
	Login     string `json:"login"`
	ID        int    `json:"id"`
	URL       string `json:"url"`
}

type Package struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Namespace      string         `json:"namespace"`
	Ecosystem      string         `json:"ecosystem"`
	PackageType    string         `json:"package_type"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
	Owner          Owner          `json:"owner"`
	Registry       Registry       `json:"registry"`
	PackageVersion PackageVersion `json:"package_version"`
}

type Owner struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Type      string `json:"type"`
}

type Registry struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type PackageVersion struct {
	Version           string            `json:"version"`
	HTMLURL           string            `json:"html_url"`
	CreatedAt         string            `json:"created_at"`
	PackageURL        string            `json:"package_url"`
	InstallCmd        string            `json:"installation_command"`
	ContainerMetadata ContainerMetadata `json:"container_metadata"`
}

type ContainerMetadata struct {
	Tag       Tag      `json:"tag"`
	Manifests Manifest `json:"manifests"`
}

type Tag struct {
	Name   string `json:"name"`
	Digest string `json:"digest"`
}

type Manifest struct {
	Digest    string `json:"digest"`
	MediaType string `json:"media_type"`
	Size      int    `json:"size"`
	URI       string `json:"uri"`
}

type VersionRepository struct {
	Repository RepositoryInfo `json:"repository"`
}

type RepositoryInfo struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Language    string `json:"primary_language_name"`
	PushedAt    string `json:"pushed_at"`
	Public      bool   `json:"public"`
	URL         string `json:"url"`
}

type Repository struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Private       bool   `json:"private"`
	HTMLURL       string `json:"html_url"`
	DefaultBranch string `json:"default_branch"`
	CloneURL      string `json:"clone_url"`
	PushedAt      string `json:"pushed_at"`
}

type Sender struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Type      string `json:"type"`
}

func EmitPackagePayload(payload PackagePublishedEvent) error {

	if payload.Event != "package" && payload.Payload.Action != "published" {
		logs.Info("non-published package event")
		return nil
	}

	type ResponsePackage struct {
		Name            string `json:"name"`
		Version         string `json:"version"`
		Owner           string `json:"owner"`
		CreatedAt       string `json:"createdAt"`
		UpdatedAt       string `json:"updatedAt"`
		PackageType     string `json:"packageType"`
		Package_url     string `json:"package_url"`
		InstallationCmd string `json:"installationCmd"`
		Registry        string `json:"registry"`
	}

	type ResponsePayload struct {
		Event          string          `json:"event"`
		RepositoryName string          `json:"repository_name"`
		Package        ResponsePackage `json:"package"`
	}

	responsePayload := ResponsePayload{
		Event:          "published",
		RepositoryName: payload.Payload.Repository.Name,
		Package: ResponsePackage{
			Name:            payload.Payload.PackageData.Name,
			Version:         payload.Payload.PackageData.PackageVersion.Version,
			Owner:           payload.Payload.PackageData.Owner.Login,
			CreatedAt:       payload.Payload.PackageData.CreatedAt,
			UpdatedAt:       payload.Payload.PackageData.UpdatedAt,
			PackageType:     payload.Payload.PackageData.PackageType,
			Package_url:     payload.Payload.PackageData.PackageVersion.PackageURL,
			InstallationCmd: payload.Payload.PackageData.PackageVersion.InstallCmd,
			Registry:        payload.Payload.PackageData.Registry.Name,
		},
	}

	err := EmitWebhookPayload(responsePayload)
	if err != nil {
		logs.Error("Failed to emit package payload", err)
		return fmt.Errorf("failed to emit package payload: %v", err)
	}
	logs.Info("Successfully emitted package payload")

	return nil
}
