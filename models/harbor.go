package models

type (
	// HookMessage is the message we receive from Harbor
	HookMessage struct {
		Type    string `json:"type"`
		OccurAT *Time  `json:"occur_at"`
		// OccurAT   int64     `json:"occur_at"`
		Operator  string    `json:"operator"`
		EventData EventData `json:"event_data"`
	}

	// EventData -
	EventData struct {
		Resources  []Resource `json:"resources"`
		Repository Repository `json:"repository"`
	}

	// Repository -
	Repository struct {
		DataCreated *Time `json:"date_created"`
		// DataCreated  int64  `json:"date_created"`
		Name         string `json:"name" yaml:"name,omitempty"`
		Namespace    string `json:"namespace" yaml:"namespace,omitempty"`
		RepoFullName string `json:"repo_full_name" yaml:"repoFullName,omitempty"`
		RepoType     string `json:"repo_type" yaml:"repoType,omitempty"`
	}

	// Resource -
	Resource struct {
		Digest      string `json:"digest"`
		Tag         string `json:"tag"`
		ResourceURL string `json:"resource_url"`
	}

	// ---------------------------------------

	// Rule -
	Rule struct {
		Pushimage []Pushimage `json:"pushimage" yaml:"pushimage,omitempty"`
		Pullimage []Pullimage `json:"pullimage" yaml:"pullimage,omitempty"`
	}

	// Pushimage -
	Pushimage struct {
		Action     string       `json:"action" yaml:"action,omitempty"` // [updateDB, none]
		Repository []Repository `json:"repository" yaml:"repository,omitempty"`
	}

	// Pullimage -
	Pullimage struct {
		Action     string       `json:"action" yaml:"action,omitempty"` // [updateDB, none]
		Repository []Repository `json:"repository" yaml:"repository,omitempty"`
	}
)
