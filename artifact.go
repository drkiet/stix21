package stix21

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Artifact struct {
	Type string                             `json:"type" binding:"required"`
	SpecVersion string                      `json:"spec_version,omitempty"`
	ID string                               `json:"id" binding:"required"`
	ObjectMarkingRefs []string              `json:"object_marking_refs,omitempty"`
	GranularMarkings []GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool                           `json:"defanged,omitempty"`
	Extensions map[string]interface{}       `json:"extensions,omitempty"`

	MimeType string `json:"mime_type,omitempty"`
	PayloadBin string `json:"payload_bin,omitempty"` // binary type
	Url string `json:"url,omitempty"`
	Hashes map[string]interface{} `json:"hashes,omitempty"`
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"` // enum
	DecryptionAlgorithm string `json:"decryption_key,omitempty"`
}

func (artifact Artifact) isValid() bool {
	return strings.Compare(artifact.Type, ArtifactType) == 0 &&
		strings.Contains(artifact.ID, ArtifactType)
}

//PrintArtifact
func PrintArtifact(artifact Artifact) {
	fmt.Println("Artifact:\n", MarshalArtifact(artifact))
}

func MarshalArtifact(artifact Artifact) (jsonData string){
	data, e := json.MarshalIndent(artifact, "", "  ")
	check(e)
	return string(data)
}

func UnmarshalArtifact(obj json.RawMessage) (artifact Artifact) {
	json.Unmarshal(obj, &artifact)
	return artifact
}

