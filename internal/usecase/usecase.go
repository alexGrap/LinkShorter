package usecase

import (
	"errors"
	"github.com/speps/go-hashids"
	"log"
	"ozon/internal/models"
	"ozon/internal/repository"
	redisRepository "ozon/internal/repository/redis"
	"strings"
	"time"
)

func CreateShortLink(fullLink string) (string, models.OwnError) {
	if !validationRequest(fullLink) {
		validErr := models.OwnError{Err: errors.New("not valid string"), Message: "Not valid string"}
		log.Print(validErr.Err)
		return "", validErr
	}
	result := "ozon.short/"

	sh, repErr := repository.CheckIfExist(fullLink)
	if repErr == nil {
		result += sh
		return result, models.OwnError{}
	}

	newLink, err := shortGeneration(fullLink)
	if err.Err != nil {
		log.Print(err.Err)
		return "", err
	}
	err = repository.CreateNewNode(newLink, fullLink)
	if err.Err != nil {
		log.Print(err.Err)
		return "", err
	}
	result += newLink
	return result, models.OwnError{}
}

func GetFullLink(shortLink string) (string, models.OwnError) {
	if !validationRequest(shortLink) {
		validErr := models.OwnError{Err: errors.New("not valid string"), Message: "Not valid string"}
		log.Print(validErr)
		return "", validErr
	}
	log.Print(shortLink)
	shortLink = shortLink[11:]
	fullLink, err := repository.GetFullLink(shortLink)
	if err.Err != nil {
		log.Print(err.Err)
		return "", err
	}
	return fullLink, models.OwnError{}

}

func shortGeneration(link string) (string, models.OwnError) {
	if strings.HasPrefix(link, "www.") {
		link = link[4:]
	}
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Print(err)
		return "", models.OwnError{Err: errors.New("server error"), Code: 500, Timestamp: time.Now()}
	}
	var arrForHash []int
	byteLink := []byte(link)
	for _, el := range byteLink {
		arrForHash = append(arrForHash, int(el))
	}
	hash, _ := h.Encode(arrForHash)
	result := hash[:4] + hash[len(hash)-6:]

	return result, models.OwnError{}
}

func CreationRedis(fullLink string) (string, models.OwnError) {
	if !validationRequest(fullLink) {
		validErr := models.OwnError{Err: errors.New("not valid string"), Message: "Not valid string"}
		log.Print(validErr.Err)
		return "", validErr
	}
	result := "ozon.short/"
	newLink, err := shortGeneration(fullLink)
	if err.Err != nil {
		log.Print(err.Err)
		return "", err
	}
	if redisRepository.CheckIfExist(newLink) {
		result += newLink
		return result, models.OwnError{}
	}
	err = redisRepository.CreateNewNode(newLink, fullLink)
	if err.Err != nil {
		log.Print(err.Err)
		return "", err
	}
	result += newLink
	return result, models.OwnError{}
}

func GetterRedis(shortLink string) (string, models.OwnError) {
	if !validationRequest(shortLink) {
		validErr := models.OwnError{Err: errors.New("not valid string"), Message: "Not valid string"}
		log.Print(validErr)
		return "", validErr
	}
	shortLink = shortLink[11:]
	fullLink, err := redisRepository.GetFullLink(shortLink)
	if err.Err != nil {
		log.Print(err.Err)
		return "", err
	}
	return fullLink, models.OwnError{}

}

func validationRequest(in string) bool {
	if in == "" {
		return false
	} else if !strings.Contains(in, ".") {
		return false
	}
	return true
}
