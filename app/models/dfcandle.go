package models

import (
	"time"

	"gotrading/gotrading/tradingalgorithm"

	"github.com/markcheno/go-talib"
)

// type Define
type DataFrameCandle struct {
	ProductCode   string         `json:"product_code"`
	Duration      time.Duration  `json:"duration"`
	Candles       []Candle       `json:"candles"`
	SMAs          []SMA          `json:"smas,omitempty"`
	EMAs          []EMA          `json:"emas,omitempty"`
	BBands        *BBands        `json:"bbands,omitempty"`
	IchimokuCloud *IchimokuCloud `json:"ichimoku,omitempty"`
	Rsi           *RSI           `json:"rsi,omitempty"`
	Macd          *Macd          `json:"macd,omitempty"`
	Hvs           []Hv           `json:"hvs,omitempty"`
	Events        *SignalEvents  `json:"events,omitempty"`
}

type SMA struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

type EMA struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

type BBands struct {
	N    int       `json:"n,omitempty"`
	K    float64   `json:"k,omitempty"`
	Up   []float64 `json:"up,omitempty"`
	Mid  []float64 `json:"mid,omitempty"`
	Down []float64 `json:"down,omitempty"`
}

type IchimokuCloud struct {
	Tenkan  []float64 `json:"tenkan,omitempty"`
	Kijun   []float64 `json:"kijun,omitempty"`
	SenkouA []float64 `json:"senkoua,omitempty"`
	SenkouB []float64 `json:"senkoub,omitempty"`
	Chikou  []float64 `json:"chikou,omitempty"`
}

type RSI struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

type Macd struct {
	FastPeriod   int       `json:"fast_period,omitempty"`
	SlowPeriod   int       `json:"slow_period,omitempty"`
	SignalPeriod int       `json:"signal_period,omitempty"`
	Macd         []float64 `json:"macd,omitempty"`
	MacdSignal   []float64 `json:"macd_signal,omitempty"`
	MacdHist     []float64 `json:"macd_hist,omitempty"`
}

type Hv struct {
	Period int       `json:"period,omitempty"`
	Values []float64 `json:"values,omitempty"`
}

// func Define
func (df *DataFrameCandle) Times() []time.Time {
	result := make([]time.Time, len(df.Candles))
	for i, candle := range df.Candles {
		result[i] = candle.Time
	}
	return result
}

func (df *DataFrameCandle) Opens() []float64 {
	result := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		result[i] = candle.Open
	}
	return result
}

func (df *DataFrameCandle) Closes() []float64 {
	result := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		result[i] = candle.Close
	}
	return result
}

func (df *DataFrameCandle) Highs() []float64 {
	result := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		result[i] = candle.High
	}
	return result
}

func (df *DataFrameCandle) Lows() []float64 {
	result := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		result[i] = candle.Low
	}
	return result
}

func (df *DataFrameCandle) Volumes() []float64 {
	result := make([]float64, len(df.Candles))
	for i, candle := range df.Candles {
		result[i] = candle.Volume
	}
	return result
}

// SMA
func (df *DataFrameCandle) AddSMA(period int) bool {
	if len(df.Candles) > period {
		df.SMAs = append(df.SMAs, SMA{
			Period: period,
			Values: talib.Sma(df.Closes(), period),
		})
		return true
	}
	return false
}

// EMA
func (df *DataFrameCandle) AddEMA(period int) bool {
	if len(df.Candles) > period {
		df.EMAs = append(df.EMAs, EMA{
			Period: period,
			Values: talib.Ema(df.Closes(), period),
		})
		return true
	}
	return false
}

// BBands
func (df *DataFrameCandle) AddBBands(n int, k float64) bool {
	if n <= len(df.Closes()) {
		up, mid, down := talib.BBands(df.Closes(), n, k, k, 0)
		df.BBands = &BBands{
			N:    n,
			K:    k,
			Up:   up,
			Mid:  mid,
			Down: down,
		}
		return true
	}
	return false
}

// Ichimoku
func (df *DataFrameCandle) AddIchimoku() bool {
	tenkanN := 9
	if len(df.Closes()) >= tenkanN {
		tenkan, kijun, senkouA, senkouB, chikou := tradingalgorithm.IchimokuCloud(df.Closes())
		df.IchimokuCloud = &IchimokuCloud{
			Tenkan:  tenkan,
			Kijun:   kijun,
			SenkouA: senkouA,
			SenkouB: senkouB,
			Chikou:  chikou,
		}
		return true
	}
	return false
}

// Rsi
func (df *DataFrameCandle) AddRSI(period int) bool {
	if len(df.Candles) > period {
		values := talib.Rsi(df.Closes(), period)
		df.Rsi = &RSI{
			Period: period,
			Values: values,
		}
		return true
	}
	return false
}

func (df *DataFrameCandle) AddMacd(inFastPeriod, inSlowPeriod, inSignalPeriod int) bool {
	if len(df.Candles) > 1 {
		outMACD, outMACDSignal, outMACDHist := talib.Macd(df.Closes(), inFastPeriod, inSlowPeriod, inSignalPeriod)
		df.Macd = &Macd{
			FastPeriod:   inFastPeriod,
			SlowPeriod:   inSlowPeriod,
			SignalPeriod: inSignalPeriod,
			Macd:         outMACD,
			MacdSignal:   outMACDSignal,
			MacdHist:     outMACDHist,
		}
		return true
	}
	return false
}

func (df *DataFrameCandle) AddHv(period int) bool {
	if len(df.Candles) >= period {
		df.Hvs = append(df.Hvs, Hv{
			Period: period,
			Values: tradingalgorithm.HistorycalVolatility(df.Closes(), period),
		})
		return true
	}
	return false
}

func (df *DataFrameCandle) AddEvents(timeTime time.Time) bool {
	SignalEvents := GetSignalEventsAfterTime(timeTime)
	if len(SignalEvents.Signals) > 0 {
		df.Events = SignalEvents
		return true
	}
	return false
}
