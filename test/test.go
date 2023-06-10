package testing

import (
	"log"
	"os"
	"ozon/internal/usecase"
)

func Test() {
	err := os.Remove("test.log")
	file, err := os.OpenFile("test.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal("Fatal in creation log.file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	redisWritingTest(file)
	redisReadingTesting(file)
	postgresWritingTest(file)
	postgresReadingTesting(file)
}

func redisWritingTest(file *os.File) {
	logger := log.New(file, "REDIS WRITING TEST: ", log.LstdFlags)
	counter := 0
	var testRequest string
	testRequest = "test.com"
	link, testErr := usecase.CreationRedis(testRequest)
	if testErr.Err != nil {
		logger.Println("Error with writing default link")
		counter++
	} else {
		if link != "ozon.short/ElMFf97uGp" {
			logger.Println("Hash was create with error")
			counter++
		}
		logger.Println("Default link was create without problem")
	}
	testRequest = "test"
	link, testErr = usecase.CreationRedis(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")
	} else {
		logger.Println("Not valid link isn't detected\n")
		counter++
	}
	testRequest = ""
	link, testErr = usecase.CreationRedis(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")

	} else {
		logger.Println("Not valid link isn't detected")
		counter++
	}
	testRequest = "test.com"
	link, testErr = usecase.CreationRedis(testRequest)
	if testErr.Err != nil {
		logger.Println("Redis: Error with post existing link")
		counter++
	} else {
		if link != "ozon.short/ElMFf97uGp" {
			logger.Println("Existing link isn't detected. New link was generated")
			counter++
		}
		logger.Println("Existing link was detected")
	}
	if counter != 0 {
		logger.Println("Losing tests")
	} else {
		logger.Println("Was correctly tested")
	}
}

func postgresWritingTest(file *os.File) {
	logger := log.New(file, "POSTGRES WRITING TEST: ", log.LstdFlags)
	counter := 0
	var testRequest string
	testRequest = "test.com"
	link, testErr := usecase.CreateShortLink(testRequest)
	if testErr.Err != nil {
		logger.Println("Error with writing default link")
		counter++
	} else {
		if link != "ozon.short/ElMFf97uGp" {
			logger.Println("Hash was create with error")
			counter++
		}
		logger.Println("Default link was create without problem")
	}
	testRequest = "test"
	link, testErr = usecase.CreateShortLink(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")
	} else {
		logger.Println("Not valid link isn't detected")
		counter++
	}
	testRequest = ""
	link, testErr = usecase.CreateShortLink(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")
	} else {
		logger.Println("Not valid link isn't detected")
		counter++
	}
	testRequest = "test.com"
	link, testErr = usecase.CreateShortLink(testRequest)
	if testErr.Err != nil {
		logger.Println("Error with post existing link")
		counter++
	} else {
		if link != "ozon.short/ElMFf97uGp" {
			logger.Println("Existing link isn't detected. New link was generated")
			counter++
		}
		logger.Println("Existing link was detected")
	}
	if counter != 0 {
		logger.Println("Losing tests")
	} else {
		logger.Println("as correctly tested")
	}
}

func redisReadingTesting(file *os.File) {
	logger := log.New(file, "REDIS READING TEST: ", log.LstdFlags)
	counter := 0
	var testRequest string
	testRequest = "ozon.short/ElMFf97uGp"
	link, testErr := usecase.GetterRedis(testRequest)
	if testErr.Err != nil {
		logger.Println("Error with getting default link")
		counter++
	} else {
		if link != "test.com" {
			logger.Println("Get was made with error")
			counter++
		}
		logger.Println("Default link got without problem")
	}
	testRequest = "test"
	link, testErr = usecase.GetterRedis(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")
	} else {
		logger.Println("Not valid link isn't detected")
		counter++
	}
	testRequest = ""
	link, testErr = usecase.GetterRedis(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")
	} else {
		logger.Println("Not valid link isn't detected")
		counter++
	}
	if counter != 0 {
		logger.Println("Losing tests")
	} else {
		logger.Println("as correctly tested")
	}
}

func postgresReadingTesting(file *os.File) {
	logger := log.New(file, "REDIS READING TEST: ", log.LstdFlags)
	counter := 0
	var testRequest string
	testRequest = "ozon.short/ElMFf97uGp"
	link, testErr := usecase.GetFullLink(testRequest)
	if testErr.Err != nil {
		logger.Println("Error with getting default link")
		counter++
	} else {
		if link != "test.com" {
			logger.Println("Link got with error")
			counter++
		}
		logger.Println("Default link got without problem")
	}
	testRequest = "test"
	link, testErr = usecase.GetFullLink(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")
	} else {
		logger.Println("Not valid link isn't detected")
		counter++
	}
	testRequest = ""
	link, testErr = usecase.GetFullLink(testRequest)
	if testErr.Err != nil {
		logger.Println("Not valid link was correctly detected")
	} else {
		logger.Println("Not valid link isn't detected")
		counter++
	}
	if counter != 0 {
		logger.Println("Losing tests")
	} else {
		logger.Println("as correctly tested")
	}
}
