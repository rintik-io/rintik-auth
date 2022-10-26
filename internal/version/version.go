package version

import (
	"time"

	"github.com/fahmyabdul/golibs"
	"github.com/rintik-io/rintik-auth/app"
)

// ResponseVersion :
type ResponseVersion struct {
	Version   string `json:"version"`
	LastCheck string `json:"last_check"`
}

func GetVersion() ResponseVersion {
	// Set default response body
	responseStruct := ResponseVersion{
		Version:   app.CurrentVersion,
		LastCheck: time.Now().Format("2006-01-02 15:04:05"),
	}

	golibs.Log.Printf("| Current Version | %+v\n", &responseStruct)

	return responseStruct
}
