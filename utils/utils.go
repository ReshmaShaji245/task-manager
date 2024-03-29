package utils

import (
	"errors"
	"fmt"
	"log"

	"os"
	"strings"

	"github.com/natefinch/lumberjack"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(structVar interface{}) (string, error) {
	var errorResp []*ErrorResponse
	validate := validator.New()

	err := validate.Struct(structVar)
	if err != nil {
		msg := "Validation Failed: FieldNotFound/UnexpectedValue `%s`"

		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, err := range validationErrors {
				var element ErrorResponse

				element.FailedField = err.StructNamespace()
				element.Tag = err.Tag()
				element.Value = err.Param()

				errorResp = append(errorResp, &element)
				fmt.Printf("errorResp: %v\n", errorResp)

				msg = fmt.Sprintf(msg, strings.Split(element.FailedField, ".")[1])
			}
		} else {
			// Handle InvalidValidationError
			msg = fmt.Sprintf(msg, err.Error())
		}
		return msg, errors.New("validation failed")
	}

	return "", nil
}

type Logger struct {
	Info    *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
	Request *log.Logger
}

func InitiateLoggers(logFilePath *string) *Logger {

	infoLogger := log.New(os.Stderr, "INFO : ", log.Ldate|log.Lshortfile)
	errorLogger := log.New(os.Stderr, "EROR : ", log.Ldate|log.Lshortfile)
	debugLogger := log.New(os.Stderr, "DEBG : ", log.Ldate|log.Lshortfile)
	requestLogger := log.New(os.Stderr, "RQST : ", log.Ldate|log.Lshortfile)

	if logFilePath != nil && *logFilePath != "" {
		f, err := os.OpenFile(*logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		infoLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})

		errorLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})

		debugLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})

		requestLogger.SetOutput(&lumberjack.Logger{
			Filename:   *logFilePath,
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		})
	}

	return &Logger{
		Info:    infoLogger,
		Error:   errorLogger,
		Debug:   debugLogger,
		Request: requestLogger,
	}
}
