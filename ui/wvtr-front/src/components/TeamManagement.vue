<script setup lang="ts">
    import { computed, inject, ref, watch } from 'vue'
    import type { Hero } from '../tools/types.ts'
    import InspectButton from './InspectButton.vue'
    import type { NavigationHandler } from '../tools/navigationHandler.ts'

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()

    const selectedHeroIds = ref<number[]>([])
    const currentTeamHeroIds = computed(() => user.value?.currentTeam?.heroes.filter(Boolean).map((hero) => hero.id) ?? [])
    const selectedCount = computed(() => selectedHeroIds.value.length)
    const hasPendingChanges = computed(() => {
        if (selectedHeroIds.value.length !== currentTeamHeroIds.value.length) {
            return true
        }

        return selectedHeroIds.value.some((heroId, idx) => heroId !== currentTeamHeroIds.value[idx])
    })

    watch(currentTeamHeroIds, (heroIds) => {
        selectedHeroIds.value = [...heroIds]
    }, { immediate: true })

    function isSelected(hero: Hero) {
        return selectedHeroIds.value.includes(hero.id)
    }

    function clickOnHero(hero: Hero) {
        if (isSelected(hero)) {
            selectedHeroIds.value = selectedHeroIds.value.filter((selectedHeroId) => selectedHeroId !== hero.id)
            return
        }

        if (selectedHeroIds.value.length < 3) {
            selectedHeroIds.value = [...selectedHeroIds.value, hero.id]
        }
    }

    async function saveTeam() {
        const selectedHeroes = selectedHeroIds.value
            .map((heroId) => user.value?.ownedHeroes.find((hero) => hero.id === heroId))
            .filter((hero): hero is Hero => Boolean(hero))

        await navigationHandler.setUserCurrentTeam(selectedHeroes)
    }
</script>

<template>
    <section class="selection-panel">
        <div class="panel-heading">
            <div>
                <p class="eyebrow">Team forge</p>
                <h2>Pick up to three heroes</h2>
                <p>Tap cards to build the lineup you want to send into the next run.</p>
            </div>

            <div class="selection-toolbar">
                <div class="stat-pill">
                    <span>Selected</span>
                    <strong>{{ selectedCount }}/3</strong>
                </div>
                <button class="primary-button" v-on:click="saveTeam()" :disabled="!hasPendingChanges">Save lineup</button>
            </div>
        </div>

        <div class="selection-grid hero-selection-grid">
            <article
                v-for="hero in user!.ownedHeroes"
                :key="hero.id"
                class="selectable-card hero-select-card"
                :class="{ 'is-selected': isSelected(hero) }"
                v-on:click="clickOnHero(hero)"
            >
                <div class="waifu-image-container selection-portrait">
                    <img :src="hero.imageUrl" :alt="hero.name"/>
                    <InspectButton :hero="hero"/>
                </div>

                <div class="selectable-copy">
                    <div>
                        <p class="eyebrow">Level {{ hero.attributes.level }}</p>
                        <h3>{{ hero.name }}</h3>
                    </div>

                    <div class="card-meta-row">
                        <span class="badge-pill">{{ hero.heroClass.name }}</span>
                        <span class="badge-pill badge-pill-soft">Rank {{ hero.rank }}</span>
                    </div>
                </div>
            </article>
        </div>
    </section>
</template>
