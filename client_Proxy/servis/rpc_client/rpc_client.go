package rpcclient

import (
	"log"
	"net/rpc"
)

type GeoClient struct {
	client *rpc.Client
}

func NewGeoClient() *GeoClient {
	client, err := rpc.Dial("tcp", "server_rpc:1234")
	if err != nil {
		log.Fatal(err)
	}
    return &GeoClient{client: client}
}

type Address struct {
    GeoLat string `json:"lat"`
    GeoLon string `json:"lon"`
    Result string `json:"result"`
}

func (gss *GeoClient) SearchSer(query RequestAddressSearch) []Address {
	
	var result []Address                                             
	err := gss.client.Call("GeoService.AddressSearch", query, &result)


	if err != nil {
		log.Fatal(err)
	}

	log.Println("Результат поиска:", result)

	return result
}

func (gss *GeoClient) GeoCoder(geocode RequestAddressGeocode) []Address {

	var result []Address                                              // Инициализируйте переменную для результата
	err := gss.client.Call("GeoService.AddressGeoCode",geocode, &result) // Измените имя метода и передайте &result
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Результат поиска:", result)

	return result
}