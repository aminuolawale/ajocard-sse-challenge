package helpers

import (
	"aminuolawale/ajocard/interfaces"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

const NIP_RATES_URL = "https://traf.nibss-plc.com.ng:7443/traf/ajax?command=website&action=detail&order=loadNIP&clientCode=NIBSS&txnSubCat=ALL"
const FAILURE_THRESHOLD = 0.5

func HandleError(err error) {
	if err !=nil{
		panic(err.Error())
	}
}

func ParsePercentage(val string) float64{
	floatPart := ""
	for i:=0; i < len(val); i++ {
		if string(val[i]) == "%"{
			break
		}
		floatPart += string(val[i])
	}
	parsedFloat, err:= strconv.ParseFloat(floatPart, 64)
	HandleError(err)
	return parsedFloat
}


func PanicHandler(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func(){
			error := recover()
			if error != nil {
				log.Println(error)
				resp := interfaces.ErrorResponse{Message: "Internal Server Error" }
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w,r)
	})
}

func WriteLogs(r *http.Request, response interfaces.StatusResponse){
	requestMethod  := r.Method
	requestPath  := r.URL.Path
	requestHeaders  := r.Header
	// responseBody := w.Body
	log.Print("Request Method: ", requestMethod)
	log.Print("Request URL: ", requestPath)
	log.Print("Request Headers: ", prepareHeader(requestHeaders))
	log.Print("Response Body: ", response)
}


func prepareHeader(m http.Header)string{
	out := ""
	filter := []string{"Connection", "Cookie", "Content-Length", "Content-Type", "User-Agent"}
	for k, v := range(m){
		if !sliceContains(filter, k){
			continue
		}
		out +=  k + ": " + joinSlice(v)
	}
	return out
}

func joinSlice(s []string) string {
	out := ""
	for _, l := range(s){
		out += l + ", "
	}
	return out
}

func sliceContains(s []string, val string) bool {
	for _, v := range(s) {
		if v == val {
			return true
		}
	}
	return false
}


func CheckStatus()interfaces.StatusResponse{
	res, err := http.Get(NIP_RATES_URL)
	HandleError(err)
	var decodedRes interfaces.Result
	err = json.NewDecoder(res.Body).Decode(&decodedRes)
	HandleError(err)
	inwardFailure := ParsePercentage(decodedRes.Msg.InwardFailure)
	outwardFailure := ParsePercentage(decodedRes.Msg.OutwardFailure)
	maxFailure := math.Max(inwardFailure, outwardFailure)
	var response interfaces.StatusResponse
	if maxFailure < FAILURE_THRESHOLD {
		response = interfaces.StatusResponse{Active: true, FailureRate: maxFailure, FailureThreshold: FAILURE_THRESHOLD}
	} else {
		response = interfaces.StatusResponse{Active: false, FailureRate: maxFailure, FailureThreshold: FAILURE_THRESHOLD}
	}
	return response
}


func SaveStatus(){
	fmt.Print("did this get called at all?")
	status := CheckStatus()
	file, err := json.MarshalIndent(status, "", " ")
	HandleError(err)
	err = ioutil.WriteFile("data/data.json", file, 0644)
	HandleError(err)
}

func ReadStatus() interfaces.StatusResponse{
	status:= interfaces.StatusResponse{}
	f,_ := ioutil.ReadFile("data/data.json") 
	_ = json.Unmarshal([]byte(f), &status)
	fmt.Println(status)
	return status
}