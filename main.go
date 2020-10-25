package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var fc *geojson.FeatureCollection
var rkiInformation = make(map[string]interface{})
var rkiInformationLock = &sync.Mutex{}

func main() {
	startRkiInformationUpdater()

	initFeatureCollection()

	setupRouter()
}

func setupRouter() {
	allowedOrigins := "http://localhost:8080" //os.Getenv("ORIGIN_ALLOWED")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{allowedOrigins})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./frontend/dist"))
	router.Handle("/", fs)
	router.PathPrefix("/css/").Handler(fs)
	router.PathPrefix("/img/").Handler(fs)
	router.PathPrefix("/js/").Handler(fs)
	router.Handle("/favicon.ico", fs)
	router.HandleFunc("/locate/{lat}/{lon}", locate)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func startRkiInformationUpdater() {
	updateRkiInformation()
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			updateRkiInformation()
		}
	}()
}

func updateRkiInformation() {
	rkiInformationLock.Lock()
	err := downloadRkiInformation("https://services7.arcgis.com/mOBPykOjAyBO2ZKk/arcgis/rest/services/RKI_Landkreisdaten/FeatureServer/0/query?where=1%3D1&outFields=*&outSR=4326&f=json", "kreis.json")
	if err != nil {
		log.Fatal(err)
	}
	initRkiInformation()
	rkiInformationLock.Unlock()
}

func downloadRkiInformation(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func initRkiInformation() {
	data, err := ioutil.ReadFile("kreis.json")
	if err != nil {
		log.Fatal("Couldn't read landkreis file")
	}
	err = json.Unmarshal(data, &rkiInformation)
	if err != nil {
		log.Fatal("Couldn't parse landkreis file")
	}
}

func initFeatureCollection() {
	data, err := ioutil.ReadFile("kreis.geojson")
	if err != nil {
		log.Fatal("Couldn't read landkreis file")
	}
	fc, err = geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		log.Fatal("Couldn't parse landkreis file")
	}
}

func findIncidence(nuts string) (float64, error) {
	rkiInformationLock.Lock()
	for _, feature := range rkiInformation["features"].([]interface{}) {
		if feature.(map[string]interface{})["attributes"].(map[string]interface{})["NUTS"] == nuts {
			return feature.(map[string]interface{})["attributes"].(map[string]interface{})["cases7_per_100k"].(float64), nil
		}
	}
	rkiInformationLock.Unlock()
	return 0, fmt.Errorf("landkreis not found")
}

func locateCoors(lat, lon float64) (LocateResult, error) {
	point := orb.Point{lon, lat}

	for _, feature := range fc.Features {
		// Try on a MultiPolygon to begin
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if !isMulti {
			log.Fatal("Something went terribly wrong, geojson was not containing a multipolygon!")
		}

		if planar.MultiPolygonContains(multiPoly, point) {
			nuts := feature.Properties.MustString("nuts")
			incidence, err := findIncidence(nuts)
			if err != nil {
				return LocateResult{}, err
			}

			return LocateResult{
				lat,
				lon,
				nuts,
				feature.Properties.MustString("gen"),
				feature.Properties.MustString("bundesland"),
				incidence,
			}, nil
		}
	}
	return LocateResult{}, fmt.Errorf("not found")
}

func locate(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	lat, err := strconv.ParseFloat(vars["lat"], 64)
	if err != nil {
		writer.Write([]byte("lat not formatted correctly!"))
		return
	}

	lon, err := strconv.ParseFloat(vars["lon"], 64)
	if err != nil {
		writer.Write([]byte("lon not formatted correctly!"))
		return
	}

	locateResult, err := locateCoors(lat, lon)

	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte("Something went wrong!"))
	} else {
		json.NewEncoder(writer).Encode(locateResult)
	}
}

type LocateResult struct {
	Lat       float64
	Lon       float64
	Nuts      string
	Name      string
	Land      string
	Incidence float64
}
