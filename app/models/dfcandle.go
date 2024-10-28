package models

import "time"

type DataFrameCandle struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Candles     []Candle      `json:"candles"`
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