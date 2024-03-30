package services

import (
    "be_hiring_app/src/helper"
   experiencemodel "be_hiring_app/src/models/ExperienceModel"
   usermodel "be_hiring_app/src/models/UserModel"
   projectmodel "be_hiring_app/src/models/ProjectModel"
    "github.com/go-playground/validator/v10"
)

var (
    validate = validator.New()
)

type mediaUpload interface {
    FileUploadExperience(file experiencemodel.File) (string, error)
    FileUploadUser(file usermodel.File) (string, error)
    FileUploadProject(file projectmodel.File) (string, error)
    RemoteUpload(url experiencemodel.Experience) (string, error)
}

type media struct {}

func NewMediaUpload() mediaUpload {
    return &media{}
}

func (*media) FileUploadExperience(file experiencemodel.File) (string, error) {
    //validate
    err := validate.Struct(file)
    if err != nil {
        return "", err
    }

    //upload
    uploadUrl, err := helper.ImageUploadHelper(file.File)
    if err != nil {
        return "", err
    }
    return uploadUrl, nil
}

func (*media) FileUploadUser(file usermodel.File) (string, error) {
    //validate
    err := validate.Struct(file)
    if err != nil {
        return "", err
    }

    //upload
    uploadUrl, err := helper.ImageUploadHelper(file.File)
    if err != nil {
        return "", err
    }
    return uploadUrl, nil
}

func (*media) FileUploadProject(file projectmodel.File) (string, error) {
    //validate
    err := validate.Struct(file)
    if err != nil {
        return "", err
    }

    //upload
    uploadUrl, err := helper.ImageUploadHelper(file.File)
    if err != nil {
        return "", err
    }
    return uploadUrl, nil
}

func (*media) RemoteUpload(url experiencemodel.Experience) (string, error) {
    //validate
    err := validate.Struct(url)
    if err != nil {
        return "", err
    }

    //upload
    uploadUrl, errUrl := helper.ImageUploadHelper(url.Photo)
    if errUrl != nil {
        return "", err
    }
    return uploadUrl, nil
}