package dualy_test

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func matchSnapshot(t *testing.T, actual image.Image, snapshot string) {
	t.Helper()

	if os.Getenv("UPDATE_SNAPSHOTS") == "1" {
		updateSnapshot(t, actual, snapshot)
	}

	snapshotFile, err := os.Open(snapshot)
	if err != nil {
		t.Fatalf("could not open snapshot file: %v", err)
	}
	defer snapshotFile.Close()

	expected, _, err := image.Decode(snapshotFile)
	if err != nil {
		t.Fatalf("could not decode snapshot file: %v", err)
	}

	ab := actual.Bounds()
	eb := expected.Bounds()
	if ab.Size() != eb.Size() {
		t.Errorf("snapshot size mismatch: got %s, want %s", ab.Size(), eb.Size())
		return
	}

	for y := ab.Min.Y; y < ab.Max.Y; y++ {
		for x := ab.Min.X; x < ab.Max.X; x++ {
			if actual.At(x, y) != expected.At(x, y) {
				t.Error("snapshot mismatch")
				return
			}
		}
	}
}

func updateSnapshot(t *testing.T, actual image.Image, snapshot string) {
	t.Helper()

	snapshotFile, err := os.Create(snapshot)
	if err != nil {
		t.Fatalf("could not create snapshot file: %v", err)
	}
	defer snapshotFile.Close()

	err = png.Encode(snapshotFile, actual)
	if err != nil {
		t.Fatalf("could not encode snapshot file: %v", err)
	}
}
