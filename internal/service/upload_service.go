package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// interface
type UploadService interface {
	/*
		|--------------------------------------------------------------------------
		| UploadLogo
		|--------------------------------------------------------------------------
		|
		| Bertugas:
		| - validasi file
		| - generate nama file
		| - simpan file
		| - mengembalikan nama file yang disimpan
		|
		|--------------------------------------------------------------------------
	*/
	UploadLogo(file *multipart.FileHeader) (string, error)
}

// struct implementasi
type uploadService struct {
}

// constructor
func NewUploadService() UploadService {
	return &uploadService{}
}

// struct method
func (s *uploadService) UploadLogo(file *multipart.FileHeader) (string, error) {
	/*
		|--------------------------------------------------------------------------
		| Validasi ukuran file
		|--------------------------------------------------------------------------
		| Maksimal 2 MB
		|--------------------------------------------------------------------------
	*/
	const maxSize = 2 * 1024 * 1024
	if file.Size > maxSize {
		return "", errors.New("maximum file size is 2 MB")
	}

	/*
		|--------------------------------------------------------------------------
		| Validasi extension
		|--------------------------------------------------------------------------
	*/
	ext := strings.ToLower(filepath.Ext(file.Filename))

	switch ext {
	case ".jpg", ".jpeg", ".png":
	default:
		return "", errors.New("only JPG and PNG are allowed")
	}

	/*
		|--------------------------------------------------------------------------
		| Generate nama file baru
		|--------------------------------------------------------------------------
	*/
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	/*
		|--------------------------------------------------------------------------
		| Pastikan folder storage/logo ada
		|--------------------------------------------------------------------------
	*/

	/*
		MkdirAll() akan:
		membuat folder jika belum ada ✅
		tidak error jika folder sudah ada ✅
		Ini adalah best practice.
	*/
	err := os.MkdirAll("storage/logo", os.ModePerm)
	if err != nil {
		return "", err
	}

	/*
		|--------------------------------------------------------------------------
		| Buka file upload
		|--------------------------------------------------------------------------
	*/
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	/*
		|--------------------------------------------------------------------------
		| Buat file tujuan
		|--------------------------------------------------------------------------
	*/
	dst, err := os.Create(filepath.Join("storage/logo", fileName))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	/*
		|--------------------------------------------------------------------------
		| Copy isi file
		|--------------------------------------------------------------------------
	*/
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
