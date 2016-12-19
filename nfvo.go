package openbaton

type ConfigurationParameter struct {
	ID          string `json:"id"`
	Version     int    `json:"version"`
	Description string `json:"description"`
	ConfKey     string `json:"confKey"`
	Value       string `json:"value"`
}

type Configuration struct {
	ID                      string                    `json:"id"`
	Version                 int                       `json:"version"`
	ProjectID               string                    `json:"projectId"`
	ConfigurationParameters []*ConfigurationParameter `json:"configurationParameters"`
	Name                    string                    `json:"name"`
}

type HistoryLifecycleEvent struct {
	ID          string `json:"id"`
	Event       string `json:"event"`
	Description string `json:"description"`
	ExecutedAt  string `json:"executedAt"`
}