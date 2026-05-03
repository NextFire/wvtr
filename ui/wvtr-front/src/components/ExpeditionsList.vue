<script setup lang="ts">
    import { inject, onMounted, ref, watch } from "vue"
    import type { ExpeditionStepResolveInfo, ExpToGetFromBack } from "../tools/types.ts"
    import { global, formatTextTimeFromTimeMS } from "../tools/utils.ts"
import type { NavigationHandler } from "@/tools/navigationHandler.ts"

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
        
    const expeditions = navigationHandler.getAvailableExpedition()
    const user = navigationHandler.getUser()

    let selectedExp = ref("")
    const errorMsg = ref("") 
    let selectionB = ref<Record<string,string>>({})
    onMounted(()=>{
        fillSelectionB(expeditions.value!)
    })

    function fillSelectionB(e: ExpToGetFromBack[]) {
        for (let i = 0; i < e.length; i++) {
            if (e[i]!.key === selectedExp.value) {
                selectionB.value[e[i]!.key] = "eselected"
            } else {
                selectionB.value[e[i]!.key] = "enotselected"
            }
        }
    }

    function clickOnExpedition(e: string) {
        if (expeditions.value) {
            if (selectedExp.value !== e) {
                selectedExp.value = e
                fillSelectionB(expeditions.value)
            } else {
                selectedExp.value = ""
                fillSelectionB(expeditions.value)
            }
        }
    }

    let expStepInfo = ref<ExpeditionStepResolveInfo|undefined>(undefined)
    async function onclick() {
        if (user.value!.currentTeam.heroes.length > 0) {
            await navigationHandler.launchExpedition(expStepInfo, selectedExp.value)
        } else {
            errorMsg.value = "You should have at least one hero in your team before starting an expedition."
        }
    }
    watch(expStepInfo, (newExpInfo) => {
        if (newExpInfo) {
            user.value!.state.state = newExpInfo.stepState
        }
    })

</script>

<template>
    <div v-if="expeditions">
        <h1>Select an Expedition</h1>
        <div style="display: flex; align-items: center; justify-content: center;">
            <div>
                <button v-on:click="onclick()">launch expedition</button>
                <p>{{ errorMsg }}</p>
            </div>
        </div>
        <div class="column">
            <div class="row"> 
                <div v-for="e in expeditions" v-on:click="clickOnExpedition(e.key)" :class="selectionB[e.key]">
                    <p style="text-align: center;">{{ e.key }}</p>
                    <p style="text-align: center;">time : {{ formatTextTimeFromTimeMS(e.duration/1000000) }}</p>
                    <img :src="e.imgURL" width="150px">
                </div>
            </div>
        </div>
    </div>
    <div v-else>
        <h1>Chargement...</h1>
    </div>
</template>


