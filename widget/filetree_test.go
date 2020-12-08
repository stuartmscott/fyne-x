package widget_test

import (
	"fyne.io/fyne/storage"
	"fyne.io/fyne/test"
	"github.com/fyne-io/fyne-x/widget"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestFileTree(t *testing.T) {
	tree := &widget.FileTree{}
	tree.Refresh() // Should not crash
}

func Test_NewFileTree(t *testing.T) {
	test.NewApp()

	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)

	root := storage.NewURI("file://" + tempDir)
	tree := widget.NewFileTree(root)
	tree.OpenAllBranches()

	assert.True(t, tree.IsBranchOpen(root.String()))
	b1, err := storage.Child(root, "A")
	assert.NoError(t, err)
	assert.True(t, tree.IsBranchOpen(b1.String()))
	b2, err := storage.Child(root, "B")
	assert.NoError(t, err)
	assert.True(t, tree.IsBranchOpen(b2.String()))
	l1, err := storage.Child(b2, "C")
	assert.NoError(t, err)
	assert.False(t, tree.IsBranchOpen(l1.String()))
}

func createTempDir(t *testing.T) string {
	t.Helper()
	tempDir, err := ioutil.TempDir("", "test")
	assert.NoError(t, err)
	err = os.MkdirAll(path.Join(tempDir, "A"), os.ModePerm)
	assert.NoError(t, err)
	err = os.MkdirAll(path.Join(tempDir, "B"), os.ModePerm)
	assert.NoError(t, err)
	err = ioutil.WriteFile(path.Join(tempDir, "B", "C"), []byte("c"), os.ModePerm)
	assert.NoError(t, err)
	return tempDir
}
