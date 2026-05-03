<script setup lang="ts">
    import { inject, ref } from "vue"
    import type { Team, User } from "../tools/types.ts"
    import type { Hero } from "../tools/types.ts"
    import { postRequest, RequestType } from "../tools/utils.ts"
    import InspectButton from "./InspectButton.vue";
    import type { NavigationHandler } from "../tools/navigationHandler.ts";

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()

    let selectedH = ref<Hero[]>([])
    let selectionB = ref(new Array(user.value!.ownedHeroes.length).fill(false))
    
    function clickOnHero(h: Hero) {
        const ownedHeroes = user.value!.ownedHeroes

        if (selectedH.value.includes(h)) { // already selected
            let idx = selectedH.value.indexOf(h)
            if (idx > -1) { // remove from selected
                selectedH.value.splice(idx, 1);
                selectionB.value[ownedHeroes.indexOf(h)] = false
            }
        } else if (selectedH.value.length < 3) { // Add to selected
            selectedH.value.push(h)
            selectionB.value[ownedHeroes.indexOf(h)] = true
        }
    }

    async function saveTeam() {
        // send request to modify current team of user
        await navigationHandler.setUserCurrentTeam(selectedH.value)
        selectedH = ref<Hero[]>([])
        selectionB.value.fill(false)
    }
</script>

<template>
    <div>
        <h1>Select You team (3 heroes maximum)</h1>
        <div style="display: flex; align-items: center; justify-content: center;">
            <div>
                <button v-on:click="saveTeam()">Save</button>
            </div>
        </div>
        <div class="row"> 
            <div v-for="h in user!.ownedHeroes">
                <div v-if="!selectionB[user!.ownedHeroes.indexOf(h)]" v-on:click="clickOnHero(h)" class="waifu-image-container">
                    <img class="hnotselected" :src="h.imageUrl">
                    <InspectButton :hero="h"/>
                </div >
                <div v-else v-on:click="clickOnHero(h)" class="waifu-image-container">
                    <img class="hselected" :src="h.imageUrl">
                    <InspectButton :hero="h"/>
                </div >
            </div>
        </div>
    </div>
</template>
