package service

import (
	"depeche/internal/static/entities"
	"depeche/internal/static/repository"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestStaticService_CreateFile(t *testing.T) {
	tests := []struct {
		Name string

		FilesToWrite []string

		ExpectedError              error
		ExpectedWrittenFilesAmount int
	}{
		{
			Name: "Parallel 4 files load",

			FilesToWrite: []string{
				"./test_files/test_img_1.png",
				"./test_files/test_img_2.png",
				"./test_files/test_img_3.png",
				"./test_files/test_img_4.png",
			},
			ExpectedError: nil,

			ExpectedWrittenFilesAmount: 0,
		},

		{
			Name: "Invalid file load",

			FilesToWrite: []string{
				"./test_files/test_img_1.png",
				"./test_files/test_img_2.png",
				"./test_files/test_img_3.png",
				"./test_files/test_incorrect_file",
			},

			ExpectedError:              errors.New("invalid file type (empty extension)"),
			ExpectedWrittenFilesAmount: 0,
		},
	}

	repo := repository.NewFileRepository()

	service := NewFileUsecase(repo)

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			var filesStreams []io.ReadCloser
			var files []*entities.UserFile
			for _, fileName := range test.FilesToWrite {
				stream, err := os.Open(fileName)
				if err != nil {
					t.Error(err)
				}

				filesStreams = append(filesStreams, stream)

				newUserFile := &entities.UserFile{
					Name: stream.Name(),
					Type: "img",
				}

				files = append(files, newUserFile)
			}

			files, err := service.CreateFile(files, filesStreams)
			fmt.Println(err)
			if test.ExpectedWrittenFilesAmount == 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, len(files), test.ExpectedWrittenFilesAmount)
		})
	}
}
