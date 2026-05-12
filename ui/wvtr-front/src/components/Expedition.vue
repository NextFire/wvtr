<script setup lang="ts">
    import { computed, inject, onMounted, onUnmounted, ref, type Ref } from 'vue'
    import { EncounterState, type ExpeditionStepResolveInfo, type Team } from '../tools/types.ts'
    import { formatTextTimeFromTimeMS } from '../tools/utils.ts'
    import Travel from './Travel.vue'
    import Neutral from './Neutral.vue'
    import Fight from './Fight.vue'
    import type { NavigationHandler } from '@/tools/navigationHandler.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()

    const timertxt = ref('Preparing encounter...')
    const answer = navigationHandler.getCurrentExpeditionStepResolveInfo()
    const eteam = ref<Team | undefined>(undefined)

    const phaseTitle = computed(() => {
        switch (user.value?.state.state) {
            case EncounterState.Travel:
                return 'Travel Phase'
            case EncounterState.Neutral:
                return 'Neutral Event'
            case EncounterState.Fight:
                return 'Battle Phase'
            default:
                return 'Expedition'
        }
    })

    const phaseCopy = computed(() => {
        switch (user.value?.state.state) {
            case EncounterState.Travel:
                return 'Your crew is crossing the map and the next encounter is already counting down.'
            case EncounterState.Neutral:
                return 'A narrative event is unfolding. Keep an eye on the timeline before the next state shift.'
            case EncounterState.Fight:
                return 'The enemy is on-screen. Compare both teams while the combat timeline advances.'
            default:
                return 'Tracking the current run state and preparing the next scene.'
        }
    })

    async function manageTimer(newtarget: Ref<ExpeditionStepResolveInfo | undefined>) {
        if (newtarget.value && newtarget.value.stepState) {
            eteam.value = newtarget.value.eTeam ? newtarget.value.eTeam : undefined
            user.value!.state.state = newtarget.value.stepState
            timer = launchTimer()
            return
        }

        stopTimer()
        await navigationHandler.fetchExpeditionReport()
        user.value!.state.state = EncounterState.Report
    }

    onMounted(async () => {
        await navigationHandler.fetchCurrentExpeditionStepResolveInfo(user.value!.id)
        await manageTimer(answer)
    })

    onUnmounted(() => {
        stopTimer()
    })

    async function tick() {
        if (!answer.value || user.value?.state.state == EncounterState.Report || !answer.value.timeline || !answer.value.timeline[answer.value.timeline.length - 1]) {
            stopTimer()
            return
        }

        const countDownDate = Date.parse(answer.value.timeline[answer.value.timeline.length - 1]!.when)
        const now = new Date().getTime()
        const distance = countDownDate - now

        navigationHandler.applyTimelineEventToTeam(user, answer, now)
        timertxt.value = formatTextTimeFromTimeMS(distance)

        if (distance < 0) {
            timertxt.value = 'Finished'
            await navigationHandler.fetchCurrentExpeditionStepResolveInfo(user.value!.id)
            stopTimer()
            await manageTimer(answer)
        }
    }

    let timer: number | undefined

    function launchTimer() {
        const endAt = answer.value?.timeline?.[answer.value.timeline.length - 1]?.when
        if (endAt) {
            tick()
            return setInterval(tick, 1000)
        }
    }

    function stopTimer() {
        clearInterval(timer)
        timer = undefined
    }
</script>

<template>
    <section v-if="answer" class="expedition-screen">
        <section class="panel panel-feature expedition-overview">
            <div>
                <p class="eyebrow">Active encounter</p>
                <h2>{{ phaseTitle }}</h2>
                <p>{{ phaseCopy }}</p>
            </div>

            <div class="countdown-card">
                <span>Next state shift</span>
                <strong>{{ timertxt }}</strong>
            </div>
        </section>

        <section class="expedition-stage">
            <Travel v-if="user!.state.state == EncounterState.Travel"/>
            <Neutral v-else-if="user!.state.state == EncounterState.Neutral"/>
            <Fight v-else-if="user!.state.state == EncounterState.Fight" :eteam="eteam"/>
        </section>
    </section>

    <section v-else class="panel expedition-loading">
        <p class="eyebrow">Tracking expedition</p>
        <h2>Trying to figure out where the party is...</h2>
    </section>
</template>
