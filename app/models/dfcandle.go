package models

import (
	"time"

	"github.com/markcheno/go-talib"
)

type DataFrameCandle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Candles     []Candle      `json:"candles"`
	SMAs        []SMA         `json:"smas,omitempty"`
	EMAs        []EMA         `json:"emas,omitempty"`
	BBands      *BBands       `jason:"bbands,omitempty"`
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
