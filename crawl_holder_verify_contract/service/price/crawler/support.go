package crawler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"holder_contract/pkg/log"
	"holder_contract/pkg/server"
	"holder_contract/service/price/dao"
	"io"
	"io/ioutil"
	"net/http"
)

type StatusRes struct {
	Data struct {
		Status bool `json:"status"`
	} `json:"data"`
}

func checkContractVerifiedDB(cryptoId string) (bool, error) {
	return false, nil
	api := server.Config.GetString("DOMAIN_LOCAL") + fmt.Sprintf(server.Config.GetString("API_CHECK_CONTRACT_VERIFY"), cryptoId)
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	statusRes := &StatusRes{}
	err = json.Unmarshal(body, &statusRes)
	if err != nil {
		return false, err
	}
	if statusRes.Data.Status {
		return true, nil
	} else {
		return false, nil
	}
}

func updateHolderContract(repo dao.CryptoRepo) error {
	return nil
	api := server.Config.GetString("DOMAIN_LOCAL") + server.Config.GetString("API_UPDATE_HOLDERS_CONTRACT")
	jsonBody, err := json.Marshal(repo)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPatch, api, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return errors.New(res.Status + " " + string(resBody))
	}
	return nil

}

func checkContractETH(address string) (bool, error) {
	api := fmt.Sprintf("https://api.etherscan.io/api?module=contract&action=getsourcecode&address=%v&apikey=CASW5JBRADRZWQZVF6Z5X5ZFPCUFKHG333", address)
	resp, err := ethClient.Get(api)
	if err != nil {
		log.Println(log.LogLevelWarn, "checkContract", err)
	}

	if resp.Body != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(log.LogLevelWarn, "ioutil.ReadAll(resp.Body)", err)
		}
		defer resp.Body.Close()

		resCheckAPI := &ResCheckAPI{}
		err = json.Unmarshal(body, &resCheckAPI)
		if err != nil {
			log.Println(log.LogLevelWarn, "checkContract Unmarshal(body, &resCheckAPI)", err)
		}

		if err != nil {
			return false, nil
		}

		if resCheckAPI.Result[0].Abi == "Contract source code not verified" {
			return false, nil
		} else {
			return true, nil
		}
	}
	return false, nil
}

func checkContractBSC(address string) (bool, error) {
	api := fmt.Sprintf("https://api.bscscan.com/api?module=contract&action=getsourcecode&address=%v&apikey=QCECXQZ5GT5PP5W53HQX3MKQRBP8CW3VBC", address)
	resp, err := bscClient.Get(api)
	if err != nil {
		log.Println(log.LogLevelWarn, "checkContract", err)
	}

	if resp.Body != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(log.LogLevelWarn, "ioutil.ReadAll(resp.Body)", err)
		}
		defer resp.Body.Close()

		resCheckAPI := &ResCheckAPI{}
		err = json.Unmarshal(body, &resCheckAPI)
		if err != nil {
			log.Println(log.LogLevelWarn, "checkContract Unmarshal(body, &resCheckAPI)", err)
		}

		if err != nil {
			return false, nil
		}

		if resCheckAPI.Result[0].Abi == "Contract source code not verified" {
			return false, nil
		} else {
			return true, nil
		}
	}
	return false, nil
}
