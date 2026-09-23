package manipulate

import (
	"image/color"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/boxes-ltd/imaging"
)

func TestResizeIfNeeded(t *testing.T) {
	tmpDir := t.TempDir()
	conf := &config.Config{
		StorageURL:       filepath.Join(tmpDir, "orig"),
		ResizeStorageURL: filepath.Join(tmpDir, "resized"),
	}

	_ = os.MkdirAll(conf.StorageURL, 0755)
	origPath := filepath.Join(conf.StorageURL, "test.jpg")
	createTestImage(t, origPath)

	tests := []struct {
		name       string
		objectID   string
		parameters url.Values
		wantErr    bool
	}{
		{
			name:       "no dimensions",
			objectID:   "test.jpg",
			parameters: url.Values{},
			wantErr:    false,
		},
		{
			name:       "invalid width string",
			objectID:   "test.jpg",
			parameters: url.Values{"width": []string{"abc"}},
			wantErr:    true,
		},
		{
			name:       "invalid height string",
			objectID:   "test.jpg",
			parameters: url.Values{"height": []string{"abc"}},
			wantErr:    true,
		},
		{
			name:       "width below minimum",
			objectID:   "test.jpg",
			parameters: url.Values{"width": []string{"29"}},
			wantErr:    true,
		},
		{
			name:       "height below minimum",
			objectID:   "test.jpg",
			parameters: url.Values{"height": []string{"29"}},
			wantErr:    true,
		},
		{
			name:       "width above maximum",
			objectID:   "test.jpg",
			parameters: url.Values{"width": []string{"3001"}},
			wantErr:    true,
		},
		{
			name:       "height above maximum",
			objectID:   "test.jpg",
			parameters: url.Values{"height": []string{"3001"}},
			wantErr:    true,
		},
		{
			name:       "valid width resize",
			objectID:   "test.jpg",
			parameters: url.Values{"width": []string{"100"}},
			wantErr:    false,
		},
		{
			name:       "valid height resize",
			objectID:   "test.jpg",
			parameters: url.Values{"height": []string{"100"}},
			wantErr:    false,
		},
		{
			name:       "valid both dimensions",
			objectID:   "test.jpg",
			parameters: url.Values{"width": []string{"100"}, "height": []string{"200"}},
			wantErr:    false,
		},
		{
			name:       "cached resize returned",
			objectID:   "test.jpg",
			parameters: url.Values{"width": []string{"100"}},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResizeIfNeeded(conf, tt.objectID, tt.parameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResizeIfNeeded() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != nil {
				got.Close()
			}
		})
	}
}

func TestResize(t *testing.T) {
	tmpDir := t.TempDir()
	conf := &config.Config{
		StorageURL:       filepath.Join(tmpDir, "orig"),
		ResizeStorageURL: filepath.Join(tmpDir, "resized"),
	}

	_ = os.MkdirAll(conf.StorageURL, 0755)
	origPath := filepath.Join(conf.StorageURL, "test.jpg")
	createTestImage(t, origPath)

	t.Run("successful resize", func(t *testing.T) {
		f, err := resize(conf, "test.jpg", 100, 200)
		if err != nil {
			t.Fatalf("Resize() error = %v", err)
		}
		defer f.Close()

		info, _ := f.Stat()
		if info.Size() == 0 {
			t.Error("resized image is empty")
		}
	})

	t.Run("missing original", func(t *testing.T) {
		_, err := resize(conf, "nonexistent.jpg", 100, 100)
		if err == nil {
			t.Error("expected error for missing original")
		}
	})

	t.Run("width prefix when width > height", func(t *testing.T) {
		// cleanup vorherige Resizes
		os.RemoveAll(conf.ResizeStorageURL)
		_ = os.MkdirAll(conf.ResizeStorageURL, 0755)

		resize(conf, "test.jpg", 200, 100)
		expected := filepath.Join(conf.ResizeStorageURL, "w200_test.jpg")
		if _, err := os.Stat(expected); os.IsNotExist(err) {
			t.Error("expected w-prefixed filename")
		}
	})

	t.Run("height prefix when height >= width", func(t *testing.T) {
		// cleanup vorherige Resizes
		os.RemoveAll(conf.ResizeStorageURL)
		_ = os.MkdirAll(conf.ResizeStorageURL, 0755)

		resize(conf, "test.jpg", 100, 200)
		expected := filepath.Join(conf.ResizeStorageURL, "h200_test.jpg")
		if _, err := os.Stat(expected); os.IsNotExist(err) {
			t.Error("expected h-prefixed filename")
		}
	})

	t.Run("equal dimensions uses height prefix", func(t *testing.T) {
		// cleanup vorherige Resizes
		os.RemoveAll(conf.ResizeStorageURL)
		_ = os.MkdirAll(conf.ResizeStorageURL, 0755)

		resize(conf, "test.jpg", 100, 100)
		expected := filepath.Join(conf.ResizeStorageURL, "h100_test.jpg")
		if _, err := os.Stat(expected); os.IsNotExist(err) {
			t.Error("expected h-prefixed filename for equal dimensions")
		}
	})
}

func createTestImage(t *testing.T, path string) {
	t.Helper()
	img := imaging.New(10, 10, color.NRGBA{128, 128, 128, 255})
	if err := imaging.Save(img, path); err != nil {
		t.Fatalf("failed to create test image: %v", err)
	}
}
