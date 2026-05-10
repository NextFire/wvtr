<script setup lang="ts">
import type { NavigationHandler } from '@/tools/navigationHandler';
import { EncounterState, HeroTakeDamageStatus, type ExpeditionDB, type User } from '@/tools/types';
import { formatTextTimeFromTimeMS, getEncounterStateString, getStringFromFAD } from '@/tools/utils';
import { inject, ref } from 'vue';

    
    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const report = navigationHandler.getReport()
    const reportJson = JSON.stringify(report)
    async function onclick() {
        await navigationHandler.setGameState(EncounterState.Home)
        await navigationHandler.fetchTeam()
    }
</script>

<template>
    <div>
        <h1>report</h1>
        <div v-for="evR in report.whatHappened">
            <h2>{{ getEncounterStateString(evR.stepState) }}</h2>
            <div v-for="truc in evR.timeline">
                <div class="row">
                    <div>
                        {{  new Date(truc.when).toLocaleString() }} :
                    </div>
                    <div class="column">
                        <div>
                        {{ truc.what }}
                        </div>
                        <div v-if="truc.whatAction">
                            <div v-for="txt in getStringFromFAD(truc.whatAction)">
                                {{ txt }}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <button v-on:click="onclick()">ok</button>
    </div>
</template>