package controllers

import (
	"encoding/json"
	"fmt"
	"gotrading/gotrading/app/models"
	"gotrading/gotrading/config"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var templates = template.Must(template.ParseFiles("app/views/chart.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "chart.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

var apiValidPath = regexp.MustCompile("^/api/candle/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func apiCandleHandler(w http.ResponseWriter, r *http.Request) {
	productCode := r.URL.Query().Get("product_code")
	if productCode == "" {
		APIError(w, "No product_code param", http.StatusBadRequest)
		return
	}
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 1000
	}

	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1m"
	}
	durationTime := config.Config.Durations[duration]

	df, _ := models.GetAllCandle(productCode, durationTime, limit)

	sma := r.URL.Query().Get("sma")
	if sma != "" {
		strSmaPeriod1 := r.URL.Query().Get("smaperiod1")
		strSmaPeriod2 := r.URL.Query().Get("smaperiod2")
		strSmaPeriod3 := r.URL.Query().Get("smaperiod3")

		period1, err := strconv.Atoi(strSmaPeriod1)
		if strSmaPeriod1 == "" || err != nil || period1 < 0 {
			period1 = 7
		}
		period2, err := strconv.Atoi(strSmaPeriod2)
		if strSmaPeriod2 == "" || err != nil || period2 < 0 {
			period2 = 14
		}
		period3, err := strconv.Atoi(strSmaPeriod3)
		if strSmaPeriod3 == "" || err != nil || period3 < 0 {
			period3 = 21
		}
		df.AddSMA(period1)
		df.AddSMA(period2)
		df.AddSMA(period3)

	}

	ema := r.URL.Query().Get("ema")
	if ema != "" {
		strEmaPeriod1 := r.URL.Query().Get("emaPeriod1")
		strEmaPeriod2 := r.URL.Query().Get("emaPeriod2")
		strEmaPeriod3 := r.URL.Query().Get("emaPeriod3")

		period1, err := strconv.Atoi(strEmaPeriod1)
		if strEmaPeriod1 == "" || err != nil || period1 < 0 {
			period1 = 7
		}

		period2, err := strconv.Atoi(strEmaPeriod2)
		if strEmaPeriod2 == "" || err != nil || period2 < 0 {
			period2 = 14
		}

		period3, err := strconv.Atoi(strEmaPeriod3)
		if strEmaPeriod3 == "" || err != nil || period3 < 0 {
			period3 = 21
		}

		df.AddEMA(period1)
		df.AddEMA(period2)
		df.AddEMA(period3)

	}

	bbands := r.URL.Query().Get("bbands")
	if bbands != "" {
		strN := r.URL.Query().Get("bbandsN")
		strK := r.URL.Query().Get("bbandsK")

		n, err := strconv.Atoi(strN)
		if strN == "" || err != nil || n < 0 {
			n = 20
		}

		k, err := strconv.Atoi(strK)
		if strK == "" || err != nil || k < 0 {
			k = 2
		}
		df.AddBBands(n, float64(k))

	}

	ichimoku := r.URL.Query().Get("ichimoku")
	if ichimoku != "" {
		df.AddIchimoku()
	}

	rsi := r.URL.Query().Get("rsi")
	if rsi != "" {
		strPeriod := r.URL.Query().Get("rsiPeriod")
		period, err := strconv.Atoi(strPeriod)
		if strPeriod == "" || err != nil || period < 0 {
			period = 14
		}
		df.AddRSI(period)
	}

	macd := r.URL.Query().Get("macd")
	if macd != "" {
		strMacdPeriod1 := r.URL.Query().Get("macdPeriod1")
		strMacdPeriod2 := r.URL.Query().Get("macdPeriod2")
		strMacdPeriod3 := r.URL.Query().Get("macdPeriod3")
		macdPeriod1, err := strconv.Atoi(strMacdPeriod1)
		if strMacdPeriod1 == "" || err != nil || macdPeriod1 < 0 {
			macdPeriod1 = 12
		}

		macdPeriod2, err := strconv.Atoi(strMacdPeriod2)
		if strMacdPeriod2 == "" || err != nil || macdPeriod2 < 0 {
			macdPeriod1 = 26
		}

		macdPeriod3, err := strconv.Atoi(strMacdPeriod3)
		if strMacdPeriod3 == "" || err != nil || macdPeriod3 < 0 {
			macdPeriod1 = 9
		}

		df.AddMacd(macdPeriod1, macdPeriod2, macdPeriod3)

	}

	hv := r.URL.Query().Get("hv")
	if hv != "" {
		strHvPeriod1 := r.URL.Query().Get("hvPeriod1")
		strHvPeriod2 := r.URL.Query().Get("hvPeriod2")
		strHvPeriod3 := r.URL.Query().Get("hvPeriod3")

		hvPeriod1, err := strconv.Atoi(strHvPeriod1)
		if strHvPeriod1 == "" || err != nil || hvPeriod1 < 0 {
			hvPeriod1 = 21
		}

		hvPeriod2, err := strconv.Atoi(strHvPeriod2)
		if strHvPeriod2 == "" || err != nil || hvPeriod2 < 0 {
			hvPeriod1 = 63
		}

		hvPeriod3, err := strconv.Atoi(strHvPeriod3)
		if strHvPeriod3 == "" || err != nil || hvPeriod3 < 0 {
			hvPeriod1 = 252
		}

		df.AddHv(hvPeriod1)
		df.AddHv(hvPeriod2)
		df.AddHv(hvPeriod3)

	}

	js, err := json.Marshal(df)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartWebServer() error {
	http.HandleFunc("/api/candle/", apiMakeHandler(apiCandleHandler))
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
