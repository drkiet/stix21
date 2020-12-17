package stix21

import (
	"log"
	"testing"
)

func TestUnmarshalArtifact(t *testing.T) {
	log.Println("Test unmarshalling artifact")
	var artifact1 = Artifact{Type: ArtifactType, ID: MakeIdentifier(ArtifactType)}
	var jsonArtifact = MarshalArtifact(artifact1)
	var artifact2 = UnmarshalArtifact([]byte(jsonArtifact))
	if !(artifact1.isValid() && artifact2.isValid()) {
		t.Error("marshalling/unmashalling error")
	}

	log.Println("Test unmarshalling artifact ends", artifact1)
}

var fileNames = []string {
		"examples/sco/artifact_image_jpeg.json",
		"examples/sco/artifact_application_zip.json",
}

func TestArtifactImage(t *testing.T) {
	log.Println("Test image artifact")
	for _, fileName := range fileNames {
		var rawIn = readAll(fileName)
		var artifactIn = UnmarshalArtifact(rawIn)
		var rawOut = MarshalArtifact(artifactIn)
		log.Println("raw-in:\n" + string(rawIn))
		log.Println("raw-out:\n" + rawOut)
		if !artifactIn.isValid() {
			t.Error("marshalling/unmashalling error")
		}
	}
	log.Println("Test image artifact ends")
}