package template

import "encoding/xml"

type BilheteAudio struct {
	XMLName xml.Name `xml:"staff"`
}

func CreateAudioTemplate() {

}

// <?xml version="1.0" encoding="utf-8"?>
// <VIGIA>
// <CHAMADA IDEVENTO="1111111" IDCHAMADA="551199991234VIGIA328555" TIPO="CH_ATEND" NUM_DESVIO="9977771234" LIID="11.123456" NUM_MONITORADO="23132132323" NUM_B="26632662366" DH_EVENTO="2018-07-17 15:32:43" CELULA_A="" DH_VIGIA="10-08-2018 15:31:00" MULTIMIDIA="N" LATITUDE="-15.794904" LONGITUDE="-47.898137" CENTRAL="XYZ" OPERADORA="Nextel"/>
// </VIGIA>
