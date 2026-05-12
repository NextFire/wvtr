package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"wvtrserv/data"
	"wvtrserv/logger"
	"wvtrserv/nanapi/config"
	"wvtrserv/utils"
)

type NanaClient struct {
	Config *config.NanapiConfig
	Header []string
}

func createClient() *NanaClient {
	conf := config.GetNanapiConfig()
	toEncode := conf.ClientUsername + ":" + conf.ClientSecret
	b64 := base64.StdEncoding.EncodeToString([]byte(toEncode))
	client := &NanaClient{
		Config: conf,
		Header: []string{
			"Authorization", "Basic " + b64,
			"Accept", "*/*",
			"Content-Type", "application/json",
		},
	}

	return client
}

var client *NanaClient = createClient()

func getAscendedWaifusFromDicordID(discordID string) *http.Response {
	methode := "GET"
	params := url.Values{}
	params.Add("discord_id", discordID)
	params.Add("client_id", client.Config.ClientId)
	params.Add("ascended", "1")
	params.Add("blooded", "0")
	//params.Add("level", "2")
	params.Add("exclude_custom_image", "1")

	reqPath := client.Config.NanapiDomain + "/prod/waicolle/waifus?" + params.Encode()
	logger.DumpLog.Println(reqPath)
	response := utils.Fetch(reqPath, methode, url.Values{}.Encode(), client.Header)

	if response == nil {
		return nil
	}

	return response
}

func fetchAnilistChar(wlist []*Waifu, user *data.User) []*JoinWC {
	doWlist := wlist

	methode := "POST"

	type IdsForWaifus struct {
		Ids []int `json:"ids_al"`
	}
	reqIds := IdsForWaifus{
		Ids: make([]int, len(wlist)),
	}

	for i, w := range doWlist {
		reqIds.Ids[i] = w.Charachter.IdAl
	}

	toSend, err1 := json.Marshal(reqIds)
	if err1 != nil {
		logger.ErrLog.Println("Can't marshal al char ids: ", err1)
		return nil
	}

	reqPath := client.Config.NanapiDomain + "/prod/anilist/charas/search"

	logger.DumpLog.Println("Request anilist characters to ", reqPath)
	var toSendStr string = string(toSend)
	response := utils.Fetch(reqPath, methode, toSendStr, client.Header)
	logger.DumpLog.Println("Received anilist characters response")

	var waifusAL []*CharachterAL = make([]*CharachterAL, 0)
	err := json.NewDecoder(response.Body).Decode(&waifusAL)

	if err != nil {
		logger.ErrLog.Println("Can't unmarshal waifus : ", err)
		a := ""
		json.NewDecoder(response.Body).Decode(&a)

		return nil
	}
	var mapCharAL map[int]*CharachterAL = make(map[int]*CharachterAL, 0)
	for _, cal := range waifusAL {
		mapCharAL[cal.IdAl] = cal
	}

	var res []*JoinWC = make([]*JoinWC, 0)

	for i, w := range doWlist {
		if user.GetOwnedHeroByWaifuID(doWlist[i].ID) == nil {
			res = append(res, &JoinWC{
				ID:                doWlist[i].ID,
				IdAl:              mapCharAL[w.Charachter.IdAl].IdAl,
				NameUserPreferred: mapCharAL[w.Charachter.IdAl].NameUserPreferred,
				ImageLarge:        mapCharAL[w.Charachter.IdAl].ImageLarge,
				Rank:              mapCharAL[w.Charachter.IdAl].Rank,
			})
		}
	}

	return res
}

func fetchAnilistCharBulk(wlist []*Waifu, bulksize int) ([]*Waifu, []*JoinWC) {
	rest := []*Waifu{}
	doWlist := wlist
	if len(wlist) > bulksize {
		rest = wlist[bulksize:]
		doWlist = wlist[:bulksize]
	}
	methode := "GET"
	params := url.Values{}
	ids := ""
	logger.DumpLog.Println(doWlist)
	for _, w := range doWlist {
		ids += fmt.Sprintf("%d,", w.Charachter.IdAl)
	}
	if ids != "" {
		params.Add("ids_al", ids[:len(ids)-1])
	}

	reqPath := client.Config.NanapiDomain + "/prod/anilist/charas?" + params.Encode()

	logger.DumpLog.Println("Request anilist charachters to ", reqPath)
	response := utils.Fetch(reqPath, methode, url.Values{}.Encode(), client.Header)
	logger.DumpLog.Println("Received anilist charachters response")

	var waifus []*CharachterAL = make([]*CharachterAL, 0)
	err := json.NewDecoder(response.Body).Decode(&waifus)

	if err != nil {
		logger.ErrLog.Println("Can't unmarshal waifus : ", err)
		return nil, nil
	}

	//logger.DumpLog.Println("Decoded : ", str)
	var res []*JoinWC = make([]*JoinWC, len(waifus))
	if len(waifus) != len(doWlist) {
		logger.ErrLog.Println("problem happened when getting the anilist characters char/waifu number miss match : ", len(waifus), "/", len(doWlist))
		return nil, nil
	}

	for i := 0; i < len(waifus); i++ {
		res[i] = &JoinWC{
			ID:                doWlist[i].ID,
			IdAl:              waifus[i].IdAl,
			NameUserPreferred: waifus[i].NameUserPreferred,
			ImageLarge:        waifus[i].ImageLarge,
			Rank:              waifus[i].Rank,
		}
	}

	return rest, res
}

func getAnilistChar(wlist []*Waifu, user *data.User) []*JoinWC {
	return fetchAnilistChar(wlist, user)
}

func getAnilistCharBulk(wlist []*Waifu) []*JoinWC {
	rest := wlist
	var res []*JoinWC
	bulkSize := 20
	rest, res = fetchAnilistCharBulk(wlist, bulkSize)
	for len(rest) > 0 {
		tmp1, tmp2 := fetchAnilistCharBulk(rest, bulkSize)
		rest = tmp1
		res = append(res, tmp2...)
	}
	return res
}

func GetAvailableWaifuToSendToWVTR(user *data.User) []*JoinWC {
	responseWaifu := getAscendedWaifusFromDicordID(user.DiscordID)
	if responseWaifu == nil {
		return nil
	}
	var waifus []*Waifu = make([]*Waifu, 0)
	//logger.DumpLog.Println(string(utils.ReadResponse(responseWaifu)))
	err := json.NewDecoder(responseWaifu.Body).Decode(&waifus) //json.Unmarshal(utils.ReadResponse(responseWaifu), &waifus)
	if err != nil {
		logger.ErrLog.Println("Can't Decode waifus : ", err)
		return nil
	}

	logger.DumpLog.Println("Received ascended waifus :", len(waifus))

	return getAnilistChar(waifus, user)
}
