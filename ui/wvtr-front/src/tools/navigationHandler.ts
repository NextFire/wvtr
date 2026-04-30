import { inject, ref } from "vue"
import type { Hero, User } from "./types"
import type { VueCookies } from "vue-cookies";
import { fetchData, global, RequestType } from "./utils";

enum HomeStatus {
    Noting = 1,
    Connexion,
    Home,
    ExpeditionManagement,
    TeamManagement,
    HeroMaker,
    InspectHero,
}

// connexion handling 
const authUrl = ref<string | undefined>(undefined)
const cookieKey = 'wvtrusrid'
//const $cookies = inject<VueCookies>('$cookies');

function isUserIdInRequestParams(): boolean {
    let urlParams = new URLSearchParams(window.location.search);
    return urlParams.has(cookieKey)
}

function isUserIdInCookies($cookies: VueCookies): boolean {
    return $cookies.get("wvtrusrid") != undefined
}

function getUserIdFromRequestParams(): string | null {
    let urlParams = new URLSearchParams(window.location.search);
    let res = urlParams.get(cookieKey)
    return !res ? null : res
}

function getUserIdFromCookies($cookies: VueCookies): string | null {
    console.log($cookies)
    return $cookies.get(cookieKey)
}

function getUserIDFromCookiesOrURLParams($cookies: VueCookies | undefined) {
    let wvtrusrid: string | null = null
    let urlParams = new URLSearchParams(window.location.search);
    if ($cookies != undefined && $cookies.get(cookieKey)) {
        wvtrusrid = $cookies.get(cookieKey)
        console.log("cookies uid : " + wvtrusrid)
    } else if (urlParams.has(cookieKey)) {
        wvtrusrid = urlParams.get(cookieKey)
        console.log("url param uid : " + wvtrusrid)
    }
    return wvtrusrid
}

async function setupCookieWithUserID($cookies: VueCookies, uid: number) {
    console.log("user.id = " + uid)
    // we got here only if cookies have a wvtrusrid and we fetched a user
    let uidfcookie = getUserIdFromCookies($cookies!)
    if (uidfcookie == uid + "") { // check integrity of the cookie and the user
        // all good
    } else {
        // client think it has the right uid but it is wrong auth again
        requestToAuth($cookies)
    }
}

async function requestToAuth($cookies: VueCookies) {
    const authServer = "https://auth.japan7.bde.enseeiht.fr";
    const client_id = ref<string>("japan7")

    const resp = await fetch(`${authServer}/.well-known/openid-configuration`);
    const config = await resp.json();
    console.log("uidfcookie = " + getUserIdFromCookies($cookies!))
    console.log(config);
    const params = new URLSearchParams();
    params.set("response_type", "code");
    params.set("client_id", client_id.value);
    params.set("redirect_uri", `${global.DOMAIN_NAME}/api/oidc/callback`);
    params.set("scope", "openid profile discord_id");
    authUrl.value = `${config.authorization_endpoint}?${params.toString()}`;
    window.location.replace(authUrl.value);
}

class NavigationHandler {
    // page status management
    homeStatus = ref(HomeStatus.Noting)
    heroToInspect = ref<Hero | undefined>(undefined)

    // the hendler is in charge of keeping it updated
    user = ref<User | undefined>(undefined)

    constructor() {
        this.homeStatus.value = HomeStatus.Noting
    }

    setHomeStatus(newHomeStatus: HomeStatus): void {
        this.homeStatus.value = newHomeStatus
    }

    getHomeStatus() {
        return this.homeStatus
    }

    setHeroToInspect(h: Hero) {
        this.heroToInspect.value = h
    }

    getHeroToInspect() {
        return this.heroToInspect
    }

    getUser() {
        return this.user
    }

    async setup($cookies: VueCookies) {
        if ($cookies && isUserIdInRequestParams()) { // we got here only after authantification has been done
            let uidstring = getUserIdFromRequestParams()
            $cookies.set(cookieKey, uidstring, '30d', undefined, undefined, true, "Strict")

            // redirect to the main page 
            window.location.replace(global.DOMAIN_NAME);
        } else if ($cookies && isUserIdInCookies($cookies)) { // client has been connected once, but we need to check if it matches the database
            await fetchData<User>(this.user, RequestType.User, [{ id: "id", value: `${getUserIdFromCookies($cookies)}` }])
            if (this.user.value != undefined) {
                setupCookieWithUserID($cookies, this.user.value.id)
            }
        } else { // we don't know the user and need to ask auth for it
            await requestToAuth($cookies)
        }
        return this.user
    }
}

export {
    NavigationHandler,
    HomeStatus,
}
