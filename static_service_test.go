package _023_1_CashMap

import (
	"bufio"
	"bytes"
	"depeche/static/entities"
	"depeche/static/repository"
	service2 "depeche/static/service"
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
				"test_img_1.png",
				"test_img_2.png",
				"test_img_3.png",
				"test_img_4.png",
			},
			ExpectedError: nil,

			ExpectedWrittenFilesAmount: 4,
		},

		{
			Name: "Invalid file load",

			FilesToWrite: []string{
				"test_img_1.png",
				"test_img_2.png",
				"test_img_3.png",
				"test_incorrect_file",
			},

			ExpectedError:              errors.New("invalid file type (empty extension)"),
			ExpectedWrittenFilesAmount: 0,
		},
	}

	repo := repository.NewFileRepository()

	service := service2.NewFileUsecase(repo)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			var filesStreams []io.ReadCloser
			var files []*entities.UserFile
			for _, fileName := range test.FilesToWrite {
				stream, err := os.Open("./static/service/test_files/" + fileName)
				if err != nil {
					t.Error(err)
				}

				filesStreams = append(filesStreams, stream)

				newUserFile := &entities.UserFile{
					Name: fileName,
					Type: "img",
				}

				files = append(files, newUserFile)
			}

			filesOut, err := service.CreateFile(files, filesStreams)
			fmt.Println(err)
			if test.ExpectedWrittenFilesAmount == 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			fmt.Println(len(filesOut))
			fmt.Println(test.ExpectedWrittenFilesAmount)
			assert.Equal(t, len(filesOut), test.ExpectedWrittenFilesAmount)
		})

	}

	err := os.RemoveAll("./files/img/")
	if err != nil {
		t.Error(err)
	}

	err = os.RemoveAll("./files/doc/")
	if err != nil {
		t.Error(err)
	}
}

func TestStaticService_ReadFile(t *testing.T) {
	tests := []struct {
		Name string

		FilesToWrite []string

		ExpectedError error
	}{
		{
			Name: "Read existing file",

			FilesToWrite: []string{"test_img_1.png"},

			ExpectedError: nil,
		},
	}

	repo := repository.NewFileRepository()

	service := service2.NewFileUsecase(repo)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			var filesStreams []io.ReadCloser
			var files []*entities.UserFile
			for _, fileName := range test.FilesToWrite {
				stream, err := os.Open("./static//service/test_files/" + fileName)
				if err != nil {
					t.Error(err)
				}

				filesStreams = append(filesStreams, stream)

				newUserFile := &entities.UserFile{
					Name: fileName,
					Type: "img",
				}

				files = append(files, newUserFile)
			}

			filesOut, _ := service.CreateFile(files, filesStreams)
			var data []byte
			writer := bufio.NewWriter(bytes.NewBuffer(data))
			assert.NoError(t, service.ReadFile(filesOut[0], writer))
		})

	}

	err := os.RemoveAll("./files/img/")
	if err != nil {
		t.Error(err)
	}

	err = os.RemoveAll("./files/doc/")
	if err != nil {
		t.Error(err)
	}
}
