package rpcclient

type GeoSearchServicer interface {
	SearchSer(query string) string
	GeoCoder(query string) string
}