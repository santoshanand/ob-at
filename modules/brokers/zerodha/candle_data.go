package zerodha

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/santoshanand/at-kite/kite"
	"github.com/santoshanand/at/modules/common/utils"
)

var (
	profile     = "/oms/user/profile/full"
	margin      = "/oms/user/margins"
	buyRegular  = "/oms/orders/regular"
	marginOrder = "/oms/margins/orders"
	historyURL  = "/oms/instruments/historical"
	userOrders  = "/oms/orders"
	userTrades  = "/oms/trades"
)

func (z *zerodha) setToken() error {
	if utils.IsNotEmpty(z.token) {
		z.kConnect.SetAccessToken(z.token)
		return nil
	}
	token, err := z.dao.NewUserDao().GetSuperToken()
	if err != nil {
		z.log.Debug("error to get token: ", err.Error())
		return err
	}
	if token == nil {
		z.log.Debug("token is nil: ")
		return errors.New("token is nil")
	}
	z.token = *token
	z.kConnect.SetAccessToken(z.token)
	return nil
}

func (z *zerodha) handleNativeError(reader io.Reader) error {
	var dt ErrorResponse
	err := json.NewDecoder(reader).Decode(&dt)
	if err != nil {
		z.log.Debug("error decode:", err.Error())
		return err
	}
	return errors.New(dt.Message)
}

func (z *zerodha) getZerodha(token, url string, out interface{}) error {
	client := &http.Client{}
	finalURL := fmt.Sprintf("%s%s", z.cfg.KiteURL, url)
	req, err := http.NewRequest(http.MethodGet, finalURL, nil)
	if err != nil {
		return err
	}
	tokenString := fmt.Sprintf("enctoken %s", token)
	req.Header.Add("Authorization", tokenString)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.StatusCode != 200 {
		return z.handleNativeError(res.Body)
	}
	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return err
	}
	return nil
}

func (z *zerodha) getHistoryURL(instrumentToken int, interval string, fromDate time.Time, toDate time.Time) string {
	return fmt.Sprintf("%s/%d/%s?user_id=%s&oi=1&from=%s&to=%s", historyURL, instrumentToken, interval, z.cfg.GetUsername(), utils.FormatTimeToYYYYMMDDDate(fromDate), utils.FormatTimeToYYYYMMDDDate(toDate))
}

func (z *zerodha) formatHistoricalData(value DataCandle) ([]Candle, error) {
	var candles []Candle
	for _, v := range value.Candles {
		candle := ToCandle(v)
		candles = append(candles, candle)
	}
	return candles, nil
}

// DataByInterval implements IZerodha, interval could be minute, 5minute, 10minute, 15minute
func (z *zerodha) DataByInterval(instrumentToken int, interval string) ([]Candle, error) {
	err := z.setToken()
	if err != nil {
		return nil, err
	}
	toDate := utils.CurrentTime()
	fromDate := toDate.AddDate(0, 0, -5)

	url := z.getHistoryURL(instrumentToken, interval, fromDate, toDate)
	historyResponse := &HistoricalResponse{}
	err = z.getZerodha(z.token, url, historyResponse)
	if err != nil {
		return nil, err
	}
	candles, err := z.formatHistoricalData(historyResponse.Data)
	if err != nil {
		return nil, err
	}
	if err != nil {
		z.log.Debug("error HistoricalData: ", err)
		return nil, err
	}
	return candles, nil
}

func (z *zerodha) GetOHLC(instruments ...string) (kite.QuoteOHLC, error) {
	err := z.setToken()
	if err != nil {
		return nil, err
	}
	return z.kConnect.GetOHLC(instruments...)
}
