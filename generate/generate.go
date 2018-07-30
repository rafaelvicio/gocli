package generate

import (
	"errors"
	"fmt"
)

func Generate(typeFlag string, outputName string) error {

	if outputName == "" {
		outputName = "output.xml"
	}

	switch typeFlag {
	default:
		return errors.New("Tipo de Flag não reconhecida")
	case "audio":
		generateAudioBilhete()
	case "sms":
		generateSMSBilhete()
	}

	return nil
}

func generateAudioBilhete() {
	fmt.Println("Gerando bilhete...")
}

func generateSMSBilhete() {
	fmt.Println("Generate SMS...")
}
