package requestHandler

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// CreateExerciseImage used to get image from request and create it
func CreateExerciseImage(r *http.Request) (string, error) {
	// Image handling
	r.ParseMultipartForm(10 << 20)

	fileName := ""
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("image")

	if file != nil {
		defer file.Close()
	}

	locationCustomExercises, exists := os.LookupEnv("LocationCustomExercises")

	if !exists {
		return "", errors.New("enviroment variable is not set")
	}

	locationSmallImages := locationCustomExercises + "smallImage/"
	locationBigImages := locationCustomExercises + "bigImage/"

	if err != nil {
		if err.Error() != "http: no such file" {

			return "", err
		}
	} else {
		// Create a temporary file within our temp-images directory that follows
		tempFile, err := ioutil.TempFile(locationBigImages, "upload-*.jpeg")
		if err != nil {
			return "", err
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}
		tempFile.Write(fileBytes)

		smth, err := jpeg.Decode(bytes.NewReader(fileBytes))
		var rotatedImage image.Image
		if err != nil {
			return "", err
		}

		x, err := exif.Decode(bytes.NewReader(fileBytes))
		var rotation float64 = 0

		if err == nil {
			orientationRaw, err := x.Get("Orientation")

			if err == nil {
				orientation := orientationRaw.String()
				if orientation == "3" {
					rotation = 180
				} else if orientation == "6" {
					rotation = 270
				} else if orientation == "8" {
					rotation = 90
				}
			}
		}

		if rotation != 0 {
			rotatedImage = imaging.Rotate(smth, rotation, color.Gray{})
		} else {
			rotatedImage = smth
		}

		f, err := os.Create(locationSmallImages + filepath.Base(tempFile.Name()))

		if err != nil {
			return "", err
		}

		newSize := imaging.Thumbnail(rotatedImage, 256, 256, imaging.ResampleFilter{})

		err = jpeg.Encode(f, newSize, nil)

		if err != nil {
			return "", err
		}

		fileName = filepath.Base(tempFile.Name())
	}
	r.ParseMultipartForm(0)

	return fileName, nil
}

func CreateProfileImage(r *http.Request) (string, error) {
	// Image handling
	r.ParseMultipartForm(10 << 20)

	fileName := ""
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("image")
	if file != nil {
		defer file.Close()
	}

	locationImage, exists := os.LookupEnv("LocationProfileImage")
	if !exists {
		return "", errors.New("enviroment variable is not set")
	}

	if err != nil {
		if err.Error() != "http: no such file" {

			return "", err
		}
	} else {
		// Create a temporary file within our temp-images directory that follows
		tempFile, err := ioutil.TempFile(locationImage, "upload-*.jpeg")

		if err != nil {
			return "", err
		}
		defer tempFile.Close()
		// read all of the contents of our uploaded file into a
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}
		tempFile.Write(fileBytes)

		smth, err := jpeg.Decode(bytes.NewReader(fileBytes))
		var rotatedImage image.Image
		if err != nil {
			return "", err
		}

		x, err := exif.Decode(bytes.NewReader(fileBytes))
		var rotation float64 = 0

		if err == nil {
			orientationRaw, err := x.Get("Orientation")

			if err == nil {
				orientation := orientationRaw.String()
				if orientation == "3" {
					rotation = 180
				} else if orientation == "6" {
					rotation = 270
				} else if orientation == "8" {
					rotation = 90
				}
			}
		}

		if rotation != 0 {
			rotatedImage = imaging.Rotate(smth, rotation, color.Gray{})
		} else {
			rotatedImage = smth
		}

		f, err := os.Create(locationImage + filepath.Base(tempFile.Name()))

		if err != nil {
			return "", errors.New("file system")
		}

		newSize := imaging.Thumbnail(rotatedImage, 256, 256, imaging.ResampleFilter{})

		err = jpeg.Encode(f, newSize, nil)

		if err != nil {
			return "", errors.New("file system")
		}

		fileName = filepath.Base(tempFile.Name())
	}
	r.ParseMultipartForm(0)

	return fileName, nil
}

func CreateWorkoutUploadImage(r *http.Request) (string, error) {
	// Image handling
	r.ParseMultipartForm(10 << 20)

	fileName := ""
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("image")

	if file != nil {
		defer file.Close()
	}

	locationImage, exists := os.LookupEnv("LocationWorkoutUploadImage")

	if !exists {
		return "", errors.New("enviroment variable is not set")
	}

	locationBigImages := locationImage + "bigImage/"
	if err != nil {
		if err.Error() != "http: no such file" {

			return "", err
		}
	} else {
		// Create a temporary file within our temp-images directory that follows
		tempFile, err := ioutil.TempFile(locationBigImages, "upload-*.jpeg")
		if err != nil {
			return "", err
		}
		defer tempFile.Close()
		// read all of the contents of our uploaded file into a
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}
		tempFile.Write(fileBytes)

		fileName = filepath.Base(tempFile.Name())
	}
	r.ParseMultipartForm(0)

	return fileName, nil
}
