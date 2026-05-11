package main

//TODO: Make DB controller better stp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"wvtrserv/data"
	"wvtrserv/databasecontroller"
	"wvtrserv/gamedata"
	"wvtrserv/gamelogic/expedition"
	"wvtrserv/logger"
	"wvtrserv/nanapi/client"
	"wvtrserv/nanapi/config"
	"wvtrserv/utils"
)

// Main page ?
// func handler(w http.ResponseWriter, r *http.Request) {
// 	logger.DumpLog.Printf("req.Method: %s\n", r.Method)
// 	logger.DumpLog.Printf("req.URL.Path: %s\n", r.URL.Path)
// 	logger.DumpLog.Printf("req.ContentLength: %d\n", r.ContentLength)

// 	d := http.Dir("./ui/vu/UI/dist")
// 	f, err := d.Open("index.html")
// 	if err != nil {
// 		logger.ErrLog.Println(err)
// 	}

// 	defer f.Close()
// 	io.Copy(w, f)
// }

type AuthEndpoints struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndPoint         string `json:"token_endpoint"`
	UserInfoEndPoint      string `json:"userinfo_endpoint"`
}

type UserToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

var authEndPoints *AuthEndpoints = &AuthEndpoints{}

func fetchAuthEndpoints() {
	resp := utils.Fetch(config.GetNanapiConfig().OIDCURL+"/.well-known/openid-configuration", "GET", "", []string{"Content-Type", "application/json"})
	if resp == nil {
		return
	}
	//defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(authEndPoints)
	if err != nil {
		logger.ErrLog.Printf("Problem while trying to get the authentificator endoints\n")
	}
}

// Main page
// func handleMainPage(w http.ResponseWriter, r *http.Request) {

// }

type DiscordAccount struct {
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

func handlerAuth(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerAuth]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	// Example usage
	tokenEndpoint := authEndPoints.AuthorizationEndpoint
	logger.DumpLog.Println(tokenEndpoint)
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", config.GetNanapiConfig().OIDCCLientId)
	params.Add("redirect_uri", config.GetNanapiConfig().DomainName+"/api/oidc/callback")
	params.Add("scope", "openid profile discord_id")
	http.Redirect(w, r, tokenEndpoint+"?"+params.Encode(), 302)
}

// Connexion
func handlerConnexion(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerConnexion]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	// logger.DumpLog.Print(r)
	code := r.URL.Query().Get("code")
	logger.DumpLog.Println(code)
	clientId := config.GetNanapiConfig().OIDCCLientId
	clientSecret := config.GetNanapiConfig().OIDCCLientSecret

	// Example usage
	tokenEndpoint := authEndPoints.TokenEndPoint
	logger.DumpLog.Println(tokenEndpoint)
	methode := "POST"
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("code", code)
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("redirect_uri", config.GetNanapiConfig().DomainName+"/api/oidc/callback")

	header := []string{"Content-Type", "application/x-www-form-urlencoded"}
	logger.DumpLog.Println(params.Encode())
	tokenResp := utils.Fetch(tokenEndpoint, methode, params.Encode(), header)

	uToken := &UserToken{}
	err := json.NewDecoder(tokenResp.Body).Decode(uToken)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}

	// Read and print response
	//utils.ReadResponse(tokenResp)

	userResp := utils.Fetch(authEndPoints.UserInfoEndPoint, "", "", []string{"Authorization", "Bearer " + uToken.AccessToken})
	//readResponse(userResp)

	discordAccount := &DiscordAccount{}
	decodError := json.NewDecoder(userResp.Body).Decode(discordAccount)
	if decodError != nil {
		logger.ErrLog.Println(decodError)
		return
	}

	user := databasecontroller.GetUserByDiscordID(discordAccount.DiscordID)
	user = databasecontroller.GetUserGameState(user)

	// this means it's the first time the user arrive here.
	// and we need to create a new user based on the discord account info
	if user.DiscordID == "" {
		c := data.NewCurrencyOwned(databasecontroller.GetAllCurrencies())
		user = &data.User{
			Name: discordAccount.Name,
			CurrentTeam: &data.Team{
				Heroes: make([]*data.Hero, 0),
			},
			Inventory:   data.NewInventory(c),
			OwnedHeroes: make([]*data.Hero, 0),
			State: &data.GameState{
				State: data.Home,
			},
			DiscordID: discordAccount.DiscordID,
		}
		databasecontroller.CreateNewUser(user)
		logger.DumpLog.Println("Test")
		user = databasecontroller.GetUserByDiscordID(discordAccount.DiscordID)
	}

	redirectParams := url.Values{}
	redirectParams.Add("wvtrusrid", fmt.Sprintf("%d", user.ID))

	redReq, err := http.NewRequest("POST", config.GetNanapiConfig().DomainName, strings.NewReader(redirectParams.Encode()))
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	reqPath := config.GetNanapiConfig().DomainName + "?" + redirectParams.Encode()
	logger.DumpLog.Println(reqPath)
	http.Redirect(w, redReq, reqPath, http.StatusSeeOther)
}

func handleGetPlayerWaicolleAscendedWaifus(w http.ResponseWriter, r *http.Request) {
	functionS := "[handleGetPlayerWaicolleAscendedWaifus]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	user := databasecontroller.GetUserByID(uint(id))
	toSend := "{}"

	waifus := client.GetAvailableWaifuToSendToWVTR(user.DiscordID)
	if waifus == nil {
		logger.ErrLog.Printf("%s can't get response from nanpi with user [%d]", functionS, user.ID)
		fmt.Fprintf(w, "%s", toSend)
		return
	}

	strsend, err := json.Marshal(waifus)
	toSend = string(strsend)
	if err != nil {
		logger.ErrLog.Println("Can't decode waifu response")
		fmt.Fprintf(w, "%s", toSend)
	}
	logger.DumpLog.Println("Giving :", len(waifus), " waifus.")
	fmt.Fprintf(w, "%s", toSend)
}

func handlerCreateHeroForPlayer(w http.ResponseWriter, r *http.Request) {
	functionS := "[handleGetPlayerWaicolleAscendedWaifus]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)
	//user := databasecontroller.GetUserByID(uint(id))
	waifu := &client.JoinWC{}
	err := json.NewDecoder(r.Body).Decode(waifu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.DumpLog.Println(waifu)

	newH := gamedata.CreateNewHeroFromDBWaifuInfos(waifu, databasecontroller.GetHeroClasses(), databasecontroller.GetSkills())
	newH.UserID = uint(id)

	errReq := databasecontroller.CreateHero(newH)
	if errReq != nil {
		logger.ErrLog.Println("Can't Create new hero: ", errReq)
		fmt.Fprintf(w, "%s is already a hero", waifu.NameUserPreferred)
		return
	}
	stosend, err := json.Marshal(newH)

	if err != nil {
		logger.ErrLog.Println("Can't marshal new hero")
		return
	}

	fmt.Fprintf(w, "%s", stosend)
}

// Getters
func handlerHero(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerHero]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	hero := databasecontroller.GetHeroByID(uint(id))

	b, err := json.Marshal(hero)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerTeam(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerTeam]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	team := databasecontroller.GetTeamByID(uint(id))

	b, err := json.Marshal(team)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerExpeditionReport(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerExpeditionReport]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("uid")
	id, _ := strconv.Atoi(ids)

	user := databasecontroller.GetUserByID(uint(id))
	exp := user.State.CurrentExpedition
	user.GetReward(exp.ExpeditionRewards)
	databasecontroller.UpdateUser(user)

	b, err := json.Marshal(user.State.CurrentExpedition)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerAvailableExpeditions(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerAvailableExpeditions]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)
	user := databasecontroller.GetUserByID(uint(id))

	expeditions := gamedata.GetAvailableExpeditions(user)

	b, err := json.Marshal(expeditions)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerUser]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	user := databasecontroller.GetUserByID(uint(id))

	b, err := json.Marshal(user)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

// Updaters
func handlerSaveUser(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerSaveUser]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	if r.Method != http.MethodPost {
		s := fmt.Sprintf("Method not allowed (%s) POST expected.", r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}

	user := &data.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the user from the request body, got : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	databasecontroller.UpdateUser(user)
	w.WriteHeader(http.StatusCreated)
}

func handlerSaveGameState(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerSaveGameState]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	if r.Method != http.MethodPost {
		s := fmt.Sprintf("Method not allowed (%s) POST expected.", r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}

	state := &data.GameState{}
	err := json.NewDecoder(r.Body).Decode(state)
	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the user from the request body, got : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	databasecontroller.UpdateGameState(state)

	b, err := json.Marshal(state)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerUpdateTeam(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerUpdateTeam]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	if r.Method != http.MethodPost {
		s := fmt.Sprintf("%s Method not allowed (%s) POST expected.", functionS, r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}
	user := &data.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the user from the request body, got : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	databasecontroller.UpdateTeam(user.CurrentTeam)
	user = databasecontroller.GetUserByID(user.ID)
	b, err := json.Marshal(user.CurrentTeam)
	if err != nil {
		logger.ErrLog.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", string(b))
}

type CurrentStepRequestMessage struct {
	Uid  int   `json:"id"`
	Time int64 `json:"time"`
}

func handlerCurrentExpeditionStep(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerCurrentExpeditionStep]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	if r.Method != http.MethodPost {
		s := fmt.Sprintf("%s Method not allowed (%s) POST expected.", functionS, r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}
	var data CurrentStepRequestMessage
	err := json.NewDecoder(r.Body).Decode(&data)
	var t time.Time = time.Unix(0, data.Time*int64(time.Millisecond))

	user := databasecontroller.GetUserByID(uint(data.Uid))

	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the time : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := databasecontroller.UpdateGameStateWithTime(user.State, &t)
	databasecontroller.UpdateTeamWithExpAndTime(user.CurrentTeam, *user.State.CurrentExpedition, t)
	databasecontroller.UpdateGameState(user.State)

	resS := "{}"
	if res != nil {
		b, err := json.Marshal(res)

		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		resS = string(b)
	}

	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, resS)
	fmt.Fprintf(w, "%s", resS)
}

func handlerLaunchExpedition(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerLaunchExpedition]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("usr")
	id, _ := strconv.Atoi(ids)

	expIdentifier := r.PathValue("expId")

	user := databasecontroller.GetUserByID(uint(id))
	var exp expedition.Expedition = gamedata.Expeditions[expIdentifier]
	if exp.CanEnter(user) && user.Inventory.Remove(exp.Cost, exp.CostNumber) {
		databasecontroller.LaunchExpedition(user, exp.Solve(expIdentifier, user.CurrentTeam))
		//databasecontroller.SaveTeam(user.CurrentTeam)
		b, err := json.Marshal(user.State.CurrentExpedition.WhatHappened[0])
		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
		fmt.Fprintf(w, "%s", string(b))
		return
	}

	fmt.Fprintf(w, "Can't launch expedition")
}

func main() {
	fetchAuthEndpoints()
	databasecontroller.DBLogIn()

	// Get main page
	fs := http.FileServer(http.Dir("./ui/wvtr-front/dist"))
	http.Handle("/", fs)
	//http.HandleFunc("/", handleMainPage)

	// Connexion
	http.HandleFunc("/api/oidc/auth", handlerAuth)
	http.HandleFunc("/api/oidc/callback", handlerConnexion)

	// Request object by ID.
	//get
	http.HandleFunc("/hero/{id}", handlerHero)
	http.HandleFunc("/teams/{id}", handlerTeam)
	http.HandleFunc("/expeditionReport/{uid}", handlerExpeditionReport)
	http.HandleFunc("/user/{id}", handlerUser)
	http.HandleFunc("/availableexpeditions/{id}", handlerAvailableExpeditions)
	http.HandleFunc("/userwaifus/{id}", handleGetPlayerWaicolleAscendedWaifus)

	//post
	http.HandleFunc("/currentexpeditionstep/", handlerCurrentExpeditionStep)

	//Modify db
	//get
	http.HandleFunc("/launchExpedition/{usr}/{expId}", handlerLaunchExpedition)
	http.HandleFunc("/createherofromwaifu/{id}", handlerCreateHeroForPlayer)

	//post
	http.HandleFunc("/updateTeam/", handlerUpdateTeam)
	http.HandleFunc("/saveUser/", handlerSaveUser)
	http.HandleFunc("/saveGameState/", handlerSaveGameState)

	// Images handler
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs/"))))

	logger.DumpLog.Println("Listening on :4210...")
	err := http.ListenAndServe(":4210", nil)

	// testing std input
	var i int

	logger.DumpLog.Print("Type a number: ")
	fmt.Scan(&i)
	logger.DumpLog.Println("Your number is:", i)

	if err != nil {
		log.Fatal(err)
	}
}
