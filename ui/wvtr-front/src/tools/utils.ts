import type { Ref } from 'vue'
import { EncounterState, HeroTakeDamageStatus, type CurrentStepRequestMessage, type ExpeditionStepResolveInfo, type FieldActionDesc, type Hero, type User, type Waifu } from './types';

class global {

    // connexion
    public static readonly REQ_AUTH = "/api/oidc/auth";

    //Request object by id
    public static readonly REQ_HERO = "/hero/{id}";
    public static readonly REQ_TEAM = "/teams/{id}";
    public static readonly REQ_EXPEDITIONREPORT = "/expeditionReport/{uid}";
    public static readonly REQ_USR = "/user/{id}";
    public static readonly REQ_AVAILABLEEXPEDITIONS = "/availableexpeditions/{id}"
    public static readonly REQ_CURRENTEXPEDITIONSTEP = "/currentexpeditionstep/";


    //request update objects
    public static readonly REQ_LAUNCHEXPEDITION = "/launchExpedition/{usr}/{expId}";
    public static readonly REQ_UPDATETEAM = "/updateTeam/";
    public static readonly REQ_SAVEUSER = "/saveUser/";
    public static readonly REQ_SAVEGAMESTATE = "/saveGameState/";

    //Create objects
    public static readonly REQ_CREATEHEROFROMWAIFU = "/createherofromwaifu/{id}"

    //nanapi requests 
    public static readonly REQ_USERWAIFUS = "/userwaifus/{id}"

    public static readonly NO_IMAGE = "/imgs/noimage.jpg";
    public static readonly EXPEDITION = "/imgs/expedition.png";
}


enum RequestType {
    Hero = 1,
    Team,
    ExpeditionReport,
    User,
    AvailableExpeditions,
    CurrentExpeditionStep,
    UserWaifus,
    CreateHeroFromWaifu,
    SaveUser,
    SaveGameState,

    LaunchExpedition,
    UpdateTeam,
}


function buildRequestPath(reqType: RequestType, pathParams: { id: string; value: string }[] | undefined = undefined): string {
    let request: string = ""
    switch (reqType) {
        case RequestType.Hero:
            request += global.REQ_HERO
            break
        case RequestType.Team:
            request += global.REQ_TEAM
            break
        case RequestType.ExpeditionReport:
            request += global.REQ_EXPEDITIONREPORT
            break
        case RequestType.User:
            request += global.REQ_USR
            break
        case RequestType.AvailableExpeditions:
            request += global.REQ_AVAILABLEEXPEDITIONS
            break
        case RequestType.CurrentExpeditionStep:
            request += global.REQ_CURRENTEXPEDITIONSTEP
            break
        case RequestType.UserWaifus:
            request += global.REQ_USERWAIFUS
            break
        case RequestType.CreateHeroFromWaifu:
            request += global.REQ_CREATEHEROFROMWAIFU
            break
        case RequestType.LaunchExpedition:
            request += global.REQ_LAUNCHEXPEDITION
            break
        case RequestType.UpdateTeam:
            request += global.REQ_UPDATETEAM
            break
        case RequestType.SaveUser:
            request += global.REQ_SAVEUSER
            break
        case RequestType.SaveGameState:
            request += global.REQ_SAVEGAMESTATE
            break
        default:
            request = ""
            break
    }

    if (pathParams) {
        for (const pathParam of pathParams) {
            request = request.replace(`{${pathParam.id}}`, pathParam.value)
        }
    }

    return request
}

async function fetchData<T>(target: Ref<T | undefined>, reqType: RequestType, pathParams: [{ id: string; value: string }] | undefined = undefined) {
    target.value = undefined;

    let request: string = buildRequestPath(reqType, pathParams)
    console.log(request)
    if (request !== "") {
        console.log("sending get request to : " + request)
        const res = await fetch(request)
        target.value = await res.json() as T
    }
}

async function postRequest<AnswerType, BodyType>(
    answer: Ref<AnswerType | undefined>,
    dataToSend: BodyType,
    requestType: RequestType,
    pathParams: [{ id: string; value: string }] | undefined = undefined) {

    answer.value = undefined;
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(dataToSend)
    };

    let request = buildRequestPath(requestType, pathParams)
    console.log(request)
    if (request !== "") {
        console.log("sending post request to : " + request)
        const res = await fetch(request, requestOptions)
        answer.value = await res.json() as AnswerType
    }
}

async function getCurrentExpeditionStepResolveInfo(answer: Ref<ExpeditionStepResolveInfo | undefined>, usreid: number) {
    let message: CurrentStepRequestMessage = {
        id: usreid,
        time: Date.now()
    }
    await postRequest<ExpeditionStepResolveInfo, CurrentStepRequestMessage>(answer, message, RequestType.CurrentExpeditionStep)
}

async function launchExpedition(target: Ref<ExpeditionStepResolveInfo | undefined>, user: User, expIdentifier: string) {
    target.value = undefined
    let request: string = buildRequestPath(RequestType.LaunchExpedition)
    request = request.replace(`{usr}`, String(user.id))
    request = request.replace(`{expId}`, expIdentifier)
    const response = await fetch(request);
    target.value = await response.json() as ExpeditionStepResolveInfo
    console.log(target.value)
    if (target.value) {
        user.state.state = target.value.stepState
    }
}

async function createAHeroFromAWaifu(target: Ref<Hero | undefined>, waifu: Waifu, user: User) {
    console.log(waifu)
    postRequest<Hero, Waifu>(target, waifu, RequestType.CreateHeroFromWaifu, [{ id: "id", value: `${user.id}` }])
}

function formatTextTimeFromTimeMS(timeMS: number) {
    let res = ""
    //console.log(distance)
    // Time calculations for days, hours, minutes and seconds
    var days = Math.floor(timeMS / (1000 * 60 * 60 * 24));
    var hours = Math.floor((timeMS % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    var minutes = Math.floor((timeMS % (1000 * 60 * 60)) / (1000 * 60));
    var seconds = Math.floor((timeMS % (1000 * 60)) / 1000);

    if (seconds > 0) {
        res = seconds + "s"
    }
    if (minutes > 0) {
        res = minutes + "m " + res
    }
    if (hours > 0) {
        res = hours + "h " + res
    }
    if (days > 0) {
        res = days + "d " + res
    }
    return res
}

function getStringFromFAD(fad: FieldActionDesc): string[] {
    let res = new Array<string>();
    let from = fad.fromH
    let fromname = (from && from.name ? from.name : "uknown")
    let target = fad.targetH
    let targetname = (target && target.name ? target.name : "uknown")
    let status = fad.targetStatus

    let critTxt = ""
    if (!!(status & HeroTakeDamageStatus.Crit)) {
        critTxt = "(crit)"
    }

    let targetPVChange = fad.targetPVChange
    if (!!(status & HeroTakeDamageStatus.TookDamage)) {
        res.push(fromname + " has inflicted " + targetPVChange.toFixed(2) + " dmg" + critTxt + " to " + targetname)
    }
    if (!!(status & HeroTakeDamageStatus.Died)) {
        res.push(targetname + " died.")
    }
    if (!!(status & HeroTakeDamageStatus.Dodged)) {
        res.push(targetname + " dodged.")
    }
    if (!!(status & HeroTakeDamageStatus.Blocked)) {
        res.push(targetname + " blocked.")
    }
    return res
}

function getEncounterStateString(state: EncounterState): string {
    switch (state) {
        case EncounterState.Home:
            return "Home"
            break
        case EncounterState.Travel:
            return "Travel"
            break
        case EncounterState.Fight:
            return "Fight"
            break
        case EncounterState.Neutral:
            return "Neutral"
            break
        case EncounterState.Report:
            return "Report"
            break
        case EncounterState.Error:
            return "Error"
            break
    }
    return ""
}

export {
    global,
    fetchData,
    postRequest,
    launchExpedition,
    getCurrentExpeditionStepResolveInfo,
    getEncounterStateString,
    formatTextTimeFromTimeMS,
    createAHeroFromAWaifu,
    getStringFromFAD,
    buildRequestPath,
    RequestType,
}
